// Package resources provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package resources

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// Defines values for BlockNotFoundErrorCode.
const (
	BLOCKNOTFOUND BlockNotFoundErrorCode = "BLOCK_NOT_FOUND"
)

// Defines values for TransactionNotFoundErrorCode.
const (
	TRANSACTIONNOTFOUND TransactionNotFoundErrorCode = "TRANSACTION_NOT_FOUND"
)

// Block defines model for Block.
type Block struct {
	Attributes struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
	} `json:"attributes"`
	Id   string `json:"id"`
	Type string `json:"type"`
}

// BlockNotFoundError defines model for BlockNotFoundError.
type BlockNotFoundError struct {
	Code    BlockNotFoundErrorCode `json:"code"`
	Message string                 `json:"message"`
}

// BlockNotFoundErrorCode defines model for BlockNotFoundError.Code.
type BlockNotFoundErrorCode string

// Transaction defines model for Transaction.
type Transaction struct {
	Attributes struct {
		BlockNumber int    `json:"blockNumber"`
		From        string `json:"from"`
		Hash        string `json:"hash"`
		Timestamp   int    `json:"timestamp"`
		To          string `json:"to"`
		Value       string `json:"value"`
	} `json:"attributes"`
	Id   string `json:"id"`
	Type string `json:"type"`
}

// TransactionNotFoundError defines model for TransactionNotFoundError.
type TransactionNotFoundError struct {
	Code    TransactionNotFoundErrorCode `json:"code"`
	Message string                       `json:"message"`
}

// TransactionNotFoundErrorCode defines model for TransactionNotFoundError.Code.
type TransactionNotFoundErrorCode string

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /blocks/{blockId})
	BlocksGetBlock(c *gin.Context, blockId string)

	// (GET /transactions/{transactionId})
	TransactionsGetTransaction(c *gin.Context, transactionId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// BlocksGetBlock operation middleware
func (siw *ServerInterfaceWrapper) BlocksGetBlock(c *gin.Context) {

	var err error

	// ------------- Path parameter "blockId" -------------
	var blockId string

	err = runtime.BindStyledParameterWithOptions("simple", "blockId", c.Param("blockId"), &blockId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter blockId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.BlocksGetBlock(c, blockId)
}

// TransactionsGetTransaction operation middleware
func (siw *ServerInterfaceWrapper) TransactionsGetTransaction(c *gin.Context) {

	var err error

	// ------------- Path parameter "transactionId" -------------
	var transactionId string

	err = runtime.BindStyledParameterWithOptions("simple", "transactionId", c.Param("transactionId"), &transactionId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter transactionId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.TransactionsGetTransaction(c, transactionId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/blocks/:blockId", wrapper.BlocksGetBlock)
	router.GET(options.BaseURL+"/transactions/:transactionId", wrapper.TransactionsGetTransaction)
}
