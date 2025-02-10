package indexer

import (
	"context"
	"log"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
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
	ctx := context.Background()
	blockNumber := big.NewInt(0).SetUint64(number)

	// Fetch the block from the Ethereum node.
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Printf("Failed to fetch block %d: %v\n", number, err)
		return
	}

	// Extract fields from the Ethereum block.
	// (You might want to use more robust parsing depending on your use case)
	num := block.NumberU64()
	hash := block.Hash().Hex()
	parentHash := block.ParentHash().Hex()
	timestamp := block.Time() // UNIX timestamp

	// Create an instance of your xo-generated Block model.
	newBlock := models.Block{
		Number:     int64(num),
		Hash:       hash,
		ParentHash: parentHash,
		Timestamp:  int64(timestamp),
	}

	// Use Upsert (or Save) to insert/update the block in the DB.
	if err := newBlock.Upsert(ctx, database.Conn); err != nil {
		log.Printf("Failed to upsert block %d: %v\n", number, err)
		return
	}

	log.Printf("Processed and saved block %d\n", number)

	// Optionally, if you want to process transactions in the block:
	for _, tx := range block.Transactions() {
		// Extract tx fields and create a new models.Transaction
		newTx := models.Transaction{
			Hash: tx.Hash().Hex(),
			// Parse other fields (From, To, Value, etc.) as needed.
			BlockNumber: int64(num),
		}
		if err := newTx.Upsert(ctx, database.Conn); err != nil {
			log.Printf("Failed to upsert transaction %s: %v\n", tx.Hash().Hex(), err)
			continue
		}
	}
}
