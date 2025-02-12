package indexer

import (
	"context"
	"github.com/rs/zerolog/log"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// StartIndexer now accepts a starting block number.
func StartIndexer(client *ethclient.Client, database *db.DB, startBlock uint64) {
	lastProcessed := startBlock
	log.Printf("Indexer starting from block %d", lastProcessed)
	for {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Error().Msg("Error fetching latest header:")
			time.Sleep(5 * time.Second)
			continue
		}
		latest := header.Number.Uint64()
		for i := lastProcessed; i <= latest; i++ {
			processBlock(i, client, database)
			lastProcessed = i + 1
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
		log.Error().Msgf("Failed to fetch block by number %d: %v", number, err)
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
		log.Error().Msgf("Failed to upsert block %d: %v", number, err)
		return
	}

	log.Info().Msgf("Processed and saved block %d", number)

	// Optionally, if you want to process transactions in the block:
	for _, tx := range block.Transactions() {
		// Extract tx fields and create a new models.Transaction
		newTx := models.Transaction{
			Hash: tx.Hash().Hex(),
			// Parse other fields (From, To, Value, etc.) as needed.
			BlockNumber: int64(num),
		}
		if err := newTx.Upsert(ctx, database.Conn); err != nil {
			log.Error().Msgf("Failed to upsert transaction %s: %v\n", tx.Hash().Hex(), err)
			continue
		}
	}
}
