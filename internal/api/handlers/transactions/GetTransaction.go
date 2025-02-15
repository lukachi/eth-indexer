package transactions

import (
	"context"
	"github.com/rs/zerolog/log"
	context2 "lukachi/eth-indexer/internal/api/context"
	"lukachi/eth-indexer/internal/api/helpers"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	openapi "lukachi/eth-indexer/resources"
	"net/http"
)

func GetTransaction(w http.ResponseWriter, r *http.Request, transactionId string) {
	dbCtx := r.Context().Value(context2.DBCtxKey).(db.DB)
	// Query the database for the transaction by hash using xo-generated function.
	tx, err := models.TransactionByHash(context.Background(), dbCtx.Conn, transactionId)
	if err != nil {
		log.Error().Msg(err.Error())
		helpers.RenderErr(w, http.StatusNotFound, err)
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

	helpers.Render(w, http.StatusOK, resTx)
}
