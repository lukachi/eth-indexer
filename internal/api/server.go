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

// Implement the generated handler interfaces.
func (s *Server) GetBlock(c *gin.Context, params openapi.GetBlockParams) {
	// Dummy response
	c.JSON(http.StatusOK, openapi.Block{
		Id:   "12341234",
		Type: "block",
		Attributes: openapi.BlockAttributes{
			Hash:       "0x8758765f78657f65",
			Number:     1234,
			ParentHash: "0x781234617236",
			Timestamp:  123412341234,
		},
	})
}

func (s *Server) GetTransaction(c *gin.Context, params openapi.GetTransactionParams) {
	// Dummy response
	c.JSON(http.StatusOK, openapi.Transaction{
		Id:   "12341234",
		Type: "transaction",
		Attributes: openapi.TransactionAttributes{
			BlockNumber: 1234,
			From:        "0x12341234",
			Hash:        "0x12341234",
			Timestamp:   12341234,
			To:          "0x12341234",
			Value:       "12341234",
		},
	})
}

func StartServer(database *db.DB) {
	s := &Server{DB: database}
	router := gin.Default()
	openapi.RegisterHandlers(router, s)
	router.Run(":8089")
}
