package blocks

import (
	"context"
	"github.com/rs/zerolog/log"
	"lukachi/eth-indexer/internal/api/helpers"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	openapi "lukachi/eth-indexer/resources"
	"net/http"
	"strconv"
)

func GetBlock(w http.ResponseWriter, r *http.Request, DB *db.DB, blockId string) {
	// Convert the blockId (string) to an int64
	number, err := strconv.ParseInt(blockId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		helpers.RenderErr(w, http.StatusBadRequest, openapi.InternalServerError{
			Code:    "INVALID_PARAMS",
			Message: "Invalid block number format",
		})
		return
	}

	// Query the database for the block using xo-generated model function.
	block, err := models.BlockByNumber(context.Background(), DB.Conn, number)
	if err != nil {
		log.Error().Msg(err.Error())
		// TODO: parse error for case with internal errors
		helpers.RenderErr(w, http.StatusNotFound, openapi.BlockNotFoundError{
			Code:    "SQL_SCAN_ERROR",
			Message: "Failed to scan query",
		})
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

	helpers.Render(w, http.StatusOK, resBlock)
}
