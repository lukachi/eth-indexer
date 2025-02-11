package blocks

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"lukachi/eth-indexer/internal/db"
	"lukachi/eth-indexer/internal/db/models"
	openapi "lukachi/eth-indexer/resources"
	"net/http"
	"strconv"
)

func GetBlocks(c *gin.Context, DB *db.DB, params openapi.BlocksGetBlocksParams) {
	pageNumber := 1
	if params.PageNumber != nil {
		pageNumber = *params.PageNumber
	}

	pageSize := 20
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}

	builder := DB.SqlBuilder.Select("number", "hash", "parent_hash", "timestamp").From("public.blocks")
	countBuilder := DB.SqlBuilder.Select("count(*)").From("public.blocks")

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

	rows, err := DB.Conn.QueryContext(context.Background(), sqlString, args...)
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
	if err := DB.Conn.QueryRowContext(context.Background(), countSqlString, countArgs...).Scan(&totalCount); err != nil {
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

	meta := openapi.BlocksMeta{
		TotalCount: func(v int) *int { return &v }(int(totalCount)),
	}

	c.JSON(http.StatusOK, openapi.GetBlocksResponse{
		Data:  resBlocks,
		Links: &links,
		Meta:  &meta,
	})
}
