package api

import (
	"lukachi/eth-indexer/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"

	// Import the generated code from resources.
	openapi "lukachi/eth-indexer/resources"
)

type Server struct {
	DB *db.DB
}

// For MVP, convert number parameter to integer.
func ptr(s string) *string {
	return &s
}

// Implement the generated handler interfaces.
func (s *Server) GetBlock(c *gin.Context, params openapi.GetBlockParams) {
	// Example: fetch block from DB (dummy response for now)
	number := params.Number
	timestamp := 1630000000

	c.JSON(http.StatusOK, openapi.Block{
		Number:     &number,
		Hash:       ptr("0xdummyhash"),
		ParentHash: ptr("0xdummyparent"),
		Timestamp:  &timestamp,
	})
}

func (s *Server) GetTransaction(c *gin.Context, params openapi.GetTransactionParams) {
	blockNumber := 12345678
	// Dummy response using params.Hash.
	c.JSON(http.StatusOK, openapi.Transaction{
		Hash:        &params.Hash,
		From:        ptr("0xfromaddress"),
		To:          ptr("0xtoaddress"),
		Value:       ptr("1000000000000000000"),
		BlockNumber: &blockNumber,
	})
}

func StartServer(database *db.DB) {
	s := &Server{DB: database}
	router := gin.Default()
	openapi.RegisterHandlers(router, s)
	router.Run(":7000")
}
