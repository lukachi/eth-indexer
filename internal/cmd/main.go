package main

import (
	"log"
	"lukachi/eth-indexer/internal/api"
	"lukachi/eth-indexer/internal/config"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/indexer"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cfg := config.Load()

	database, err := db.ConnectDB(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	if err := database.CreateTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	ethClient, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	go indexer.StartIndexer(ethClient, database)
	api.StartServer(database)
}
