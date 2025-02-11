package main

import (
	"context"
	"log"
	"lukachi/eth-indexer/internal/api"
	"lukachi/eth-indexer/internal/config"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/indexer"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cfg := config.Load()

	database, err := db.ConnectDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	// Determine the starting block number
	startBlock := 0

	if strings.ToLower(cfg.Continue) == "true" {
		// Otherwise, if CONTINUE is true, query the database for the last block.
		err = database.Conn.QueryRowContext(context.Background(),
			"SELECT COALESCE(MAX(number), 0) FROM public.blocks").Scan(&startBlock)
		if err != nil {
			log.Fatalf("Failed to query last processed block: %v", err)
		}
		// Continue from the next block
		startBlock = startBlock + 1
		log.Printf("Continuing indexing from block %d (last saved block + 1)", startBlock)
	}

	if cfg.FromBlock != "" {
		// If FROM_BLOCK is specified, use it.
		n, err := strconv.ParseUint(cfg.FromBlock, 10, 64)
		if err != nil {
			log.Fatalf("Invalid FROM_BLOCK value: %v", err)
		}
		startBlock = int(n)
		log.Printf("Starting indexing from block %d (FROM_BLOCK specified)", startBlock)
	}

	// Start the indexer with the determined starting block.
	go indexer.StartIndexer(ethClient, database, uint64(startBlock))

	api.StartServer(database)
}
