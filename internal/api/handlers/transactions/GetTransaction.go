package transactions

import (
	"context"
	"github.com/gin-gonic/gin"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	openapi "lukachi/eth-indexer/resources"
	"net/http"
)

func GetTransaction(c *gin.Context, DB *db.DB, transactionId string) {
	// Query the database for the transaction by hash using xo-generated function.
	tx, err := models.TransactionByHash(context.Background(), DB.Conn, transactionId)
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
