package indexer

import (
	"context"
	"log"
	"lukachi/eth-indexer/internal/db"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func StartIndexer(client *ethclient.Client, database *db.DB) {
	var lastProcessed uint64 = 0
	for {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Println("Error fetching latest header:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		latest := header.Number.Uint64()
		for i := lastProcessed + 1; i <= latest; i++ {
			processBlock(i, client, database)
			lastProcessed = i
		}
		time.Sleep(3 * time.Second)
	}
}

func processBlock(number uint64, client *ethclient.Client, database *db.DB) {
	blockNumber := big.NewInt(0).SetUint64(number)
	_, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Printf("Failed to fetch block %d: %v\n", number, err)
		return
	}
	// Extract fields (number, hash, parent hash, timestamp) and insert into DB.
	// (InsertBlock(database, block) - implementation omitted for brevity)
	log.Printf("Processed block %d\n", number)
}
