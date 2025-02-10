package api

import (
	"context"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	// Import the generated code from resources.
	openapi "lukachi/eth-indexer/resources"
)

type Server struct {
	DB *db.DB
}

func (s *Server) BlocksGetBlock(c *gin.Context, blockId string) {
	// Convert the blockId (string) to an int64
	number, err := strconv.ParseInt(blockId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid block number format"})
		return
	}

	// Query the database for the block using xo-generated model function.
	block, err := models.BlockByNumber(context.Background(), s.DB.Conn, number)
	if err != nil {
		// Depending on the error you could return 404 if no row found.
		c.JSON(http.StatusNotFound, gin.H{"msg": "Block not found"})
		return
	}

	// Map the xo model to the OpenAPI response model.
	resBlock := openapi.Block{
		Id:   strconv.FormatInt(block.Number, 10), // converting number to string for API
		Type: "block",
		Attributes: openapi.BlockAttributes{
			Hash:       block.Hash,
			Number:     int(block.Number), // if the API expects an int
			ParentHash: block.ParentHash,
			Timestamp:  int(block.Timestamp), // adjust types as needed
		},
	}

	c.JSON(http.StatusOK, resBlock)
}

func (s *Server) TransactionsGetTransaction(c *gin.Context, transactionId string) {
	// Query the database for the transaction by hash using xo-generated function.
	tx, err := models.TransactionByHash(context.Background(), s.DB.Conn, transactionId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Transaction not found"})
		return
	}

	// Map the xo model to the OpenAPI response model.
	resTx := openapi.Transaction{
		Id:   tx.Hash, // assuming the transaction ID is the hash
		Type: "transaction",
		Attributes: openapi.TransactionAttributes{
			BlockNumber: int(tx.BlockNumber), // adjust type conversion as needed
			From:        tx.From,
			Hash:        tx.Hash,
			Timestamp:   int(tx.Timestamp), // if available; otherwise, adjust accordingly
			To:          tx.To,
			Value:       tx.Value,
		},
	}

	c.JSON(http.StatusOK, resTx)
}

func StartServer(database *db.DB) {
	s := &Server{DB: database}
	router := gin.Default()
	openapi.RegisterHandlers(router, s)
	router.Run(":8089")
}
