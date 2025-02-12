// Package resources provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package resources

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// Defines values for BlockNotFoundErrorCode.
const (
	BLOCKNOTFOUND BlockNotFoundErrorCode = "BLOCK_NOT_FOUND"
)

// Defines values for InternalServerErrorCode.
const (
	INTERNALSERVERERROR InternalServerErrorCode = "INTERNAL_SERVER_ERROR"
)

// Defines values for TransactionNotFoundErrorCode.
const (
	TRANSACTIONNOTFOUND TransactionNotFoundErrorCode = "TRANSACTION_NOT_FOUND"
)

// Block defines model for Block.
type Block struct {
	Attributes BlockAttributes `json:"attributes"`
	Id         string          `json:"id"`
	Type       string          `json:"type"`
}

// BlockAttributes defines model for BlockAttributes.
type BlockAttributes struct {
	Hash       string `json:"hash"`
	Number     int    `json:"number"`
	ParentHash string `json:"parentHash"`
	Timestamp  int    `json:"timestamp"`
}

// BlockNotFoundError defines model for BlockNotFoundError.
type BlockNotFoundError struct {
	Code    BlockNotFoundErrorCode `json:"code"`
	Message string                 `json:"message"`
}

// BlockNotFoundErrorCode defines model for BlockNotFoundError.Code.
type BlockNotFoundErrorCode string

// BlocksMeta defines model for BlocksMeta.
type BlocksMeta struct {
	TotalCount *int `json:"totalCount,omitempty"`
}

// GetBlocksResponse defines model for GetBlocksResponse.
type GetBlocksResponse struct {
	Data  []Block     `json:"data"`
	Links *Links      `json:"links,omitempty"`
	Meta  *BlocksMeta `json:"meta,omitempty"`
}

// GetTransactionsResponse defines model for GetTransactionsResponse.
type GetTransactionsResponse struct {
	Data  []Transaction     `json:"data"`
	Links *Links            `json:"links,omitempty"`
	Meta  *TransactionsMeta `json:"meta,omitempty"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError struct {
	Code    InternalServerErrorCode `json:"code"`
	Message string                  `json:"message"`
}

// InternalServerErrorCode defines model for InternalServerError.Code.
type InternalServerErrorCode string

// Links defines model for Links.
type Links struct {
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
	Self  *string `json:"self,omitempty"`
}

// Transaction defines model for Transaction.
type Transaction struct {
	Attributes TransactionAttributes `json:"attributes"`
	Id         string                `json:"id"`
	Type       string                `json:"type"`
}

// TransactionAttributes defines model for TransactionAttributes.
type TransactionAttributes struct {
	BlockNumber int    `json:"blockNumber"`
	From        string `json:"from"`
	Hash        string `json:"hash"`
	Timestamp   int    `json:"timestamp"`
	To          string `json:"to"`
	Value       string `json:"value"`
}

// TransactionNotFoundError defines model for TransactionNotFoundError.
type TransactionNotFoundError struct {
	Code    TransactionNotFoundErrorCode `json:"code"`
	Message string                       `json:"message"`
}

// TransactionNotFoundErrorCode defines model for TransactionNotFoundError.Code.
type TransactionNotFoundErrorCode string

// TransactionsMeta defines model for TransactionsMeta.
type TransactionsMeta struct {
	TotalCount *int `json:"totalCount,omitempty"`
}

// BlocksGetBlocksParams defines parameters for BlocksGetBlocks.
type BlocksGetBlocksParams struct {
	FilterNumber    *int    `form:"filter[number],omitempty" json:"filter[number],omitempty"`
	FilterHash      *string `form:"filter[hash],omitempty" json:"filter[hash],omitempty"`
	FilterTimestamp *int    `form:"filter[timestamp],omitempty" json:"filter[timestamp],omitempty"`
	PageNumber      *int    `form:"page[number],omitempty" json:"page[number],omitempty"`
	PageSize        *int    `form:"page[size],omitempty" json:"page[size],omitempty"`
}

// TransactionsGetTransactionsParams defines parameters for TransactionsGetTransactions.
type TransactionsGetTransactionsParams struct {
	FilterFrom        *string `form:"filter[from],omitempty" json:"filter[from],omitempty"`
	FilterTo          *string `form:"filter[to],omitempty" json:"filter[to],omitempty"`
	FilterBlockNumber *int    `form:"filter[block_number],omitempty" json:"filter[block_number],omitempty"`
	FilterTimestamp   *int    `form:"filter[timestamp],omitempty" json:"filter[timestamp],omitempty"`
	PageNumber        *int    `form:"page[number],omitempty" json:"page[number],omitempty"`
	PageSize          *int    `form:"page[size],omitempty" json:"page[size],omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /blocks)
	BlocksGetBlocks(w http.ResponseWriter, r *http.Request, params BlocksGetBlocksParams)

	// (GET /blocks/{blockId})
	BlocksGetBlock(w http.ResponseWriter, r *http.Request, blockId string)

	// (GET /transactions)
	TransactionsGetTransactions(w http.ResponseWriter, r *http.Request, params TransactionsGetTransactionsParams)

	// (GET /transactions/{transactionHash})
	TransactionsGetTransaction(w http.ResponseWriter, r *http.Request, transactionHash string)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /blocks)
func (_ Unimplemented) BlocksGetBlocks(w http.ResponseWriter, r *http.Request, params BlocksGetBlocksParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /blocks/{blockId})
func (_ Unimplemented) BlocksGetBlock(w http.ResponseWriter, r *http.Request, blockId string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /transactions)
func (_ Unimplemented) TransactionsGetTransactions(w http.ResponseWriter, r *http.Request, params TransactionsGetTransactionsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /transactions/{transactionHash})
func (_ Unimplemented) TransactionsGetTransaction(w http.ResponseWriter, r *http.Request, transactionHash string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// BlocksGetBlocks operation middleware
func (siw *ServerInterfaceWrapper) BlocksGetBlocks(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params BlocksGetBlocksParams

	// ------------- Optional query parameter "filter[number]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[number]", r.URL.Query(), &params.FilterNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[number]", Err: err})
		return
	}

	// ------------- Optional query parameter "filter[hash]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[hash]", r.URL.Query(), &params.FilterHash)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[hash]", Err: err})
		return
	}

	// ------------- Optional query parameter "filter[timestamp]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[timestamp]", r.URL.Query(), &params.FilterTimestamp)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[timestamp]", Err: err})
		return
	}

	// ------------- Optional query parameter "page[number]" -------------

	err = runtime.BindQueryParameter("form", false, false, "page[number]", r.URL.Query(), &params.PageNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page[number]", Err: err})
		return
	}

	// ------------- Optional query parameter "page[size]" -------------

	err = runtime.BindQueryParameter("form", false, false, "page[size]", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page[size]", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.BlocksGetBlocks(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// BlocksGetBlock operation middleware
func (siw *ServerInterfaceWrapper) BlocksGetBlock(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "blockId" -------------
	var blockId string

	err = runtime.BindStyledParameterWithOptions("simple", "blockId", chi.URLParam(r, "blockId"), &blockId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "blockId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.BlocksGetBlock(w, r, blockId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// TransactionsGetTransactions operation middleware
func (siw *ServerInterfaceWrapper) TransactionsGetTransactions(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params TransactionsGetTransactionsParams

	// ------------- Optional query parameter "filter[from]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[from]", r.URL.Query(), &params.FilterFrom)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[from]", Err: err})
		return
	}

	// ------------- Optional query parameter "filter[to]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[to]", r.URL.Query(), &params.FilterTo)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[to]", Err: err})
		return
	}

	// ------------- Optional query parameter "filter[block_number]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[block_number]", r.URL.Query(), &params.FilterBlockNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[block_number]", Err: err})
		return
	}

	// ------------- Optional query parameter "filter[timestamp]" -------------

	err = runtime.BindQueryParameter("form", false, false, "filter[timestamp]", r.URL.Query(), &params.FilterTimestamp)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter[timestamp]", Err: err})
		return
	}

	// ------------- Optional query parameter "page[number]" -------------

	err = runtime.BindQueryParameter("form", false, false, "page[number]", r.URL.Query(), &params.PageNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page[number]", Err: err})
		return
	}

	// ------------- Optional query parameter "page[size]" -------------

	err = runtime.BindQueryParameter("form", false, false, "page[size]", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page[size]", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TransactionsGetTransactions(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// TransactionsGetTransaction operation middleware
func (siw *ServerInterfaceWrapper) TransactionsGetTransaction(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "transactionHash" -------------
	var transactionHash string

	err = runtime.BindStyledParameterWithOptions("simple", "transactionHash", chi.URLParam(r, "transactionHash"), &transactionHash, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "transactionHash", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TransactionsGetTransaction(w, r, transactionHash)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/blocks", wrapper.BlocksGetBlocks)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/blocks/{blockId}", wrapper.BlocksGetBlock)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/transactions", wrapper.TransactionsGetTransactions)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/transactions/{transactionHash}", wrapper.TransactionsGetTransaction)
	})

	return r
}
