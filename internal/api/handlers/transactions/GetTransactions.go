package transactions

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	openapi "lukachi/eth-indexer/resources"
	"net/http"
	"strconv"
)

func GetTransactions(c *gin.Context, DB *db.DB, params openapi.TransactionsGetTransactionsParams) {
	pageNumber := 1
	if params.PageNumber != nil {
		pageNumber = *params.PageNumber
	}

	pageSize := 20
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}

	builder := DB.SqlBuilder.Select("*").From("public.transactions")
	countBuilder := DB.SqlBuilder.Select("count(*)").From("public.transactions")

	if params.FilterFrom != nil {
		builder = builder.Where(squirrel.Eq{"from_address": *params.FilterFrom})
		countBuilder = countBuilder.Where(squirrel.Eq{"from_address": *params.FilterFrom})
	}

	if params.FilterTo != nil {
		builder = builder.Where(squirrel.Eq{"to_address": *params.FilterTo})
		countBuilder = countBuilder.Where(squirrel.Eq{"to_address": *params.FilterTo})
	}

	if params.FilterBlockNumber != nil {
		builder = builder.Where(squirrel.Eq{"block_number": *params.FilterBlockNumber})
		countBuilder = countBuilder.Where(squirrel.Eq{"block_number": *params.FilterBlockNumber})
	}

	builder = builder.OrderBy("timestamp DESC")

	offset := (pageNumber - 1) * pageSize
	builder = builder.Offset(uint64(offset)).Limit(uint64(pageSize))

	sqlString, args, err := builder.ToSql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_BUILD_ERROR",
			Message: "Failed to build SQL query string",
		})
		return
	}

	rows, err := DB.Conn.QueryContext(context.Background(), sqlString, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_EXEC_ERROR",
			Message: "Failed to execute SQL query",
		})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.Hash, &transaction.Hash, &transaction.From, &transaction.To, &transaction.BlockNumber, &transaction.Value, &transaction.Timestamp); err != nil {
			c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
				Code:    "SQL_SCAN_ERROR",
				Message: "Failed to scan transaction row",
			})
			return
		}
		transactions = append(transactions, transaction)
	}

	countSqlString, countArgs, err := countBuilder.ToSql()
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_COUNT_BUILD_ERROR",
			Message: "Failed to build SQL count query",
		})
		return
	}

	var totalCount int64
	if err := DB.Conn.QueryRowContext(context.Background(), countSqlString, countArgs...).Scan(&totalCount); err != nil {
		c.JSON(http.StatusInternalServerError, openapi.InternalServerError{
			Code:    "SQL_COUNT_EXEC_ERROR",
			Message: "Failed to execute SQL count query",
		})
		return
	}

	var resTransactions []openapi.Transaction
	for _, transaction := range transactions {
		resTransaction := openapi.Transaction{
			Id:   transaction.Hash,
			Type: "transaction",
			Attributes: openapi.TransactionAttributes{
				Hash:        transaction.Hash,
				From:        transaction.From,
				To:          transaction.To,
				BlockNumber: int(transaction.BlockNumber),
				Value:       transaction.Value,
				Timestamp:   int(transaction.Timestamp),
			},
		}
		resTransactions = append(resTransactions, resTransaction)
	}
	if resTransactions == nil {
		resTransactions = []openapi.Transaction{}
	}

	baseUrl := c.Request.URL.Path

	// Generate the current page number with query parameters for consistency

	selfPage := baseUrl + "?page[number]=" + strconv.Itoa(pageNumber) + "&page[size]=" + strconv.Itoa(pageSize)
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
		Self:  &selfPage,
		First: &firstPage,
		Last:  &lastPage,
		Next:  &nextPage,
		Prev:  &prevPage,
	}

	meta := openapi.TransactionsMeta{
		TotalCount: func(v int) *int { return &v }(int(totalCount)),
	}

	c.JSON(http.StatusOK, openapi.GetTransactionsResponse{
		Data:  resTransactions,
		Links: &links,
		Meta:  &meta,
	})
}
