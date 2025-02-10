package api

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
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

func (s *Server) BlocksGetBlocks(c *gin.Context, params openapi.BlocksGetBlocksParams) {
	pageNumber := 1
	if params.PageNumber != nil {
		pageNumber = *params.PageNumber
	}

	pageSize := 20
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}

	builder := s.DB.SqlBuilder.Select("number", "hash", "parent_hash", "timestamp").From("public.blocks")
	countBuilder := s.DB.SqlBuilder.Select("count(*)").From("public.blocks")

	if params.FilterNumber != nil {
		builder = builder.Where(squirrel.Eq{"number": *params.FilterNumber})
		countBuilder = countBuilder.Where(squirrel.Eq{"number": *params.FilterNumber})
	}

	if params.FilterHash != nil {
		builder = builder.Where(squirrel.Eq{"hash": *params.FilterHash})
		countBuilder = countBuilder.Where(squirrel.Eq{"hash": *params.FilterHash})
	}

	if params.FilterTimestamp != nil {
		builder = builder.Where(squirrel.Eq{"timestamp": *params.FilterTimestamp})
		countBuilder = countBuilder.Where(squirrel.Eq{"timestamp": *params.FilterTimestamp})
	}

	builder = builder.OrderBy("number ASC")

	offset := (pageNumber - 1) * pageSize
	builder = builder.Offset(uint64(offset)).Limit(uint64(pageSize))

	sqlString, args, err := builder.ToSql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_BUILD_ERROR",
			Message: "Failed to build sql string",
		})
		return
	}

	rows, err := s.DB.Conn.QueryContext(context.Background(), sqlString, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_EXEC_ERROR",
			Message: "Failed to exec sql string",
		})
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
				Code:    "SQL_CLOSE_ERROR",
				Message: "Failed to close rows",
			})
		}
	}(rows)

	var blocks []models.Block
	for rows.Next() {
		var block models.Block
		if err := rows.Scan(&block.Number, &block.Hash, &block.ParentHash, &block.Timestamp); err != nil {
			c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
				Code:    "SQL_SCAN_ERROR",
				Message: "failed scanning row",
			})
		}
		blocks = append(blocks, block)
	}

	countSqlString, countArgs, err := countBuilder.ToSql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_EXEC_ERROR",
			Message: "Failed to build count sql string",
		})
	}

	var totalCount int64
	if err := s.DB.Conn.QueryRowContext(context.Background(), countSqlString, countArgs...).Scan(&totalCount); err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_EXEC_ERROR",
			Message: "Failed to query and scan count sql string",
		})
	}

	var resBlocks []openapi.Block
	for _, block := range blocks {
		resBlock := openapi.Block{
			Id:   strconv.FormatInt(block.Number, 10),
			Type: "block",
			Attributes: openapi.BlockAttributes{
				Hash:       block.Hash,
				Number:     int(block.Number),
				ParentHash: block.ParentHash,
				Timestamp:  int(block.Timestamp),
			},
		}
		resBlocks = append(resBlocks, resBlock)
	}
	if resBlocks == nil {
		resBlocks = []openapi.Block{}
	}

	baseUrl := c.Request.URL.Path

	firstPage := baseUrl + "?page[number]=1&page[size]=20"
	lastPageNumber := (totalCount + int64(pageSize) - 1) / int64(pageSize)
	lastPage := baseUrl + "?page[number]=" + strconv.Itoa(int(lastPageNumber)) + "&page[size]=" + strconv.Itoa(pageSize)

	var prevPage, nextPage string
	if pageNumber > 1 {
		prevPage = baseUrl + "?page[number]=" + strconv.Itoa(pageNumber-1) + "&page[size]=" + strconv.Itoa(pageSize)
	}
	if int64(pageNumber) < lastPageNumber {
		nextPage = baseUrl + "?page[number]=" + strconv.Itoa(pageNumber+1) + "&page[size]=" + strconv.Itoa(pageSize)
	}

	links := openapi.Links{
		First: &firstPage,
		Last:  &lastPage,
		Next:  &nextPage,
		Prev:  &prevPage,
	}

	meta := openapi.BlocksMeta{
		TotalCount: func(v int) *int { return &v }(int(totalCount)),
	}

	c.JSON(http.StatusOK, openapi.GetBlocksResponse{
		Data:  resBlocks,
		Links: &links,
		Meta:  &meta,
	})
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
