package api

import (
	"github.com/Lavalier/zchi"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"lukachi/eth-indexer/internal/api/handlers/blocks"
	"lukachi/eth-indexer/internal/api/handlers/transactions"
	"lukachi/eth-indexer/internal/db"
	"net/http"

	// Import the generated code from resources.
	openapi "lukachi/eth-indexer/resources"
)

type Server struct {
	DB *db.DB
}

func (s *Server) TransactionsGetTransactions(w http.ResponseWriter, r *http.Request, params openapi.TransactionsGetTransactionsParams) {
	transactions.GetTransactions(w, r, s.DB, params)
}

func (s *Server) BlocksGetBlocks(w http.ResponseWriter, r *http.Request, params openapi.BlocksGetBlocksParams) {
	blocks.GetBlocks(w, r, s.DB, params)
}

func (s *Server) BlocksGetBlock(w http.ResponseWriter, r *http.Request, blockId string) {
	blocks.GetBlock(w, r, s.DB, blockId)
}

func (s *Server) TransactionsGetTransaction(w http.ResponseWriter, r *http.Request, transactionId string) {
	transactions.GetTransaction(w, r, s.DB, transactionId)
}

func StartServer(database *db.DB) {
	s := &Server{DB: database}
	router := chi.NewRouter()

	middlewares := make([]openapi.MiddlewareFunc, 0)
	middlewares = append(middlewares, zchi.Logger(log.Logger))

	handler := openapi.HandlerWithOptions(s, openapi.ChiServerOptions{
		BaseRouter:  router,
		Middlewares: middlewares,
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	})

	err := http.ListenAndServe(":8089", handler)
	if err != nil {
		return
	}
}
