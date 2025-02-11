package api

import (
	"github.com/gin-gonic/gin"
	"lukachi/eth-indexer/internal/api/handlers/blocks"
	"lukachi/eth-indexer/internal/api/handlers/transactions"
	"lukachi/eth-indexer/internal/db"
	// Import the generated code from resources.
	openapi "lukachi/eth-indexer/resources"
)

type Server struct {
	DB *db.DB
}

func (s *Server) TransactionsGetTransactions(c *gin.Context, params openapi.TransactionsGetTransactionsParams) {
	transactions.GetTransactions(c, s.DB, params)
}

func (s *Server) BlocksGetBlocks(c *gin.Context, params openapi.BlocksGetBlocksParams) {
	blocks.GetBlocks(c, s.DB, params)
}

func (s *Server) BlocksGetBlock(c *gin.Context, blockId string) {
	blocks.GetBlock(c, s.DB, blockId)
}

func (s *Server) TransactionsGetTransaction(c *gin.Context, transactionId string) {
	transactions.GetTransaction(c, s.DB, transactionId)
}

func StartServer(database *db.DB) {
	s := &Server{DB: database}
	router := gin.Default()
	openapi.RegisterHandlers(router, s)
	err := router.Run(":8089")
	if err != nil {
		panic(err)
	}
}
