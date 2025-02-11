package blocks

import (
	"context"
	"github.com/gin-gonic/gin"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	openapi "lukachi/eth-indexer/resources"
	"net/http"
	"strconv"
)

func GetBlock(c *gin.Context, DB *db.DB, blockId string) {
	// Convert the blockId (string) to an int64
	number, err := strconv.ParseInt(blockId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid block number format"})
		return
	}

	// Query the database for the block using xo-generated model function.
	block, err := models.BlockByNumber(context.Background(), DB.Conn, number)
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
