/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type TransactionsApi interface {

	/*
	AddMetadataOnTransaction Set the metadata of a transaction by its ID.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @param txid Transaction ID.
	 @return ApiAddMetadataOnTransactionRequest
	*/
	AddMetadataOnTransaction(ctx _context.Context, ledger string, txid int32) ApiAddMetadataOnTransactionRequest

	// AddMetadataOnTransactionExecute executes the request
	AddMetadataOnTransactionExecute(r ApiAddMetadataOnTransactionRequest) (*_nethttp.Response, error)

	/*
	CountTransactions Count the transactions from a ledger.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiCountTransactionsRequest
	*/
	CountTransactions(ctx _context.Context, ledger string) ApiCountTransactionsRequest

	// CountTransactionsExecute executes the request
	CountTransactionsExecute(r ApiCountTransactionsRequest) (*_nethttp.Response, error)

	/*
	CreateTransaction Create a new transaction to a ledger.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiCreateTransactionRequest
	*/
	CreateTransaction(ctx _context.Context, ledger string) ApiCreateTransactionRequest

	// CreateTransactionExecute executes the request
	//  @return TransactionsResponse
	CreateTransactionExecute(r ApiCreateTransactionRequest) (TransactionsResponse, *_nethttp.Response, error)

	/*
	CreateTransactions Create a new batch of transactions to a ledger.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiCreateTransactionsRequest
	*/
	CreateTransactions(ctx _context.Context, ledger string) ApiCreateTransactionsRequest

	// CreateTransactionsExecute executes the request
	//  @return TransactionsResponse
	CreateTransactionsExecute(r ApiCreateTransactionsRequest) (TransactionsResponse, *_nethttp.Response, error)

	/*
	GetTransaction Get transaction from a ledger by its ID.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @param txid Transaction ID.
	 @return ApiGetTransactionRequest
	*/
	GetTransaction(ctx _context.Context, ledger string, txid int32) ApiGetTransactionRequest

	// GetTransactionExecute executes the request
	//  @return TransactionResponse
	GetTransactionExecute(r ApiGetTransactionRequest) (TransactionResponse, *_nethttp.Response, error)

	/*
	ListTransactions List transactions from a ledger.

	List transactions from a ledger, sorted by txid in descending order.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiListTransactionsRequest
	*/
	ListTransactions(ctx _context.Context, ledger string) ApiListTransactionsRequest

	// ListTransactionsExecute executes the request
	//  @return ListTransactions200Response
	ListTransactionsExecute(r ApiListTransactionsRequest) (ListTransactions200Response, *_nethttp.Response, error)

	/*
	RevertTransaction Revert a ledger transaction by its ID.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @param txid Transaction ID.
	 @return ApiRevertTransactionRequest
	*/
	RevertTransaction(ctx _context.Context, ledger string, txid int32) ApiRevertTransactionRequest

	// RevertTransactionExecute executes the request
	//  @return TransactionResponse
	RevertTransactionExecute(r ApiRevertTransactionRequest) (TransactionResponse, *_nethttp.Response, error)
}

// TransactionsApiService TransactionsApi service
type TransactionsApiService service

type ApiAddMetadataOnTransactionRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	txid int32
	requestBody *map[string]interface{}
}

// metadata
func (r ApiAddMetadataOnTransactionRequest) RequestBody(requestBody map[string]interface{}) ApiAddMetadataOnTransactionRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiAddMetadataOnTransactionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AddMetadataOnTransactionExecute(r)
}

/*
AddMetadataOnTransaction Set the metadata of a transaction by its ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param txid Transaction ID.
 @return ApiAddMetadataOnTransactionRequest
*/
func (a *TransactionsApiService) AddMetadataOnTransaction(ctx _context.Context, ledger string, txid int32) ApiAddMetadataOnTransactionRequest {
	return ApiAddMetadataOnTransactionRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		txid: txid,
	}
}

// Execute executes the request
func (a *TransactionsApiService) AddMetadataOnTransactionExecute(r ApiAddMetadataOnTransactionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.AddMetadataOnTransaction")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions/{txid}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"txid"+"}", _neturl.PathEscape(parameterToString(r.txid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetTransaction400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetTransaction404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCountTransactionsRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	reference *string
	account *string
	source *string
	destination *string
	metadata *map[string]interface{}
}

// Filter transactions by reference field.
func (r ApiCountTransactionsRequest) Reference(reference string) ApiCountTransactionsRequest {
	r.reference = &reference
	return r
}
// Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).
func (r ApiCountTransactionsRequest) Account(account string) ApiCountTransactionsRequest {
	r.account = &account
	return r
}
// Filter transactions with postings involving given account at source (regular expression placed between ^ and $).
func (r ApiCountTransactionsRequest) Source(source string) ApiCountTransactionsRequest {
	r.source = &source
	return r
}
// Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).
func (r ApiCountTransactionsRequest) Destination(destination string) ApiCountTransactionsRequest {
	r.destination = &destination
	return r
}
// Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
func (r ApiCountTransactionsRequest) Metadata(metadata map[string]interface{}) ApiCountTransactionsRequest {
	r.metadata = &metadata
	return r
}

func (r ApiCountTransactionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CountTransactionsExecute(r)
}

/*
CountTransactions Count the transactions from a ledger.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiCountTransactionsRequest
*/
func (a *TransactionsApiService) CountTransactions(ctx _context.Context, ledger string) ApiCountTransactionsRequest {
	return ApiCountTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
func (a *TransactionsApiService) CountTransactionsExecute(r ApiCountTransactionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.CountTransactions")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.reference != nil {
		localVarQueryParams.Add("reference", parameterToString(*r.reference, ""))
	}
	if r.account != nil {
		localVarQueryParams.Add("account", parameterToString(*r.account, ""))
	}
	if r.source != nil {
		localVarQueryParams.Add("source", parameterToString(*r.source, ""))
	}
	if r.destination != nil {
		localVarQueryParams.Add("destination", parameterToString(*r.destination, ""))
	}
	if r.metadata != nil {
		localVarQueryParams.Add("metadata", parameterToString(*r.metadata, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateTransactionRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	transactionData *TransactionData
	preview *bool
}

func (r ApiCreateTransactionRequest) TransactionData(transactionData TransactionData) ApiCreateTransactionRequest {
	r.transactionData = &transactionData
	return r
}
// Set the preview mode. Preview mode doesn&#39;t add the logs to the database or publish a message to the message broker.
func (r ApiCreateTransactionRequest) Preview(preview bool) ApiCreateTransactionRequest {
	r.preview = &preview
	return r
}

func (r ApiCreateTransactionRequest) Execute() (TransactionsResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateTransactionExecute(r)
}

/*
CreateTransaction Create a new transaction to a ledger.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiCreateTransactionRequest
*/
func (a *TransactionsApiService) CreateTransaction(ctx _context.Context, ledger string) ApiCreateTransactionRequest {
	return ApiCreateTransactionRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return TransactionsResponse
func (a *TransactionsApiService) CreateTransactionExecute(r ApiCreateTransactionRequest) (TransactionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  TransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.CreateTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.transactionData == nil {
		return localVarReturnValue, nil, reportError("transactionData is required and must be specified")
	}

	if r.preview != nil {
		localVarQueryParams.Add("preview", parameterToString(*r.preview, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transactionData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 304 {
			var v TransactionsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CreateTransaction400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v CreateTransaction409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateTransactionsRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	transactions *Transactions
}

func (r ApiCreateTransactionsRequest) Transactions(transactions Transactions) ApiCreateTransactionsRequest {
	r.transactions = &transactions
	return r
}

func (r ApiCreateTransactionsRequest) Execute() (TransactionsResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateTransactionsExecute(r)
}

/*
CreateTransactions Create a new batch of transactions to a ledger.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiCreateTransactionsRequest
*/
func (a *TransactionsApiService) CreateTransactions(ctx _context.Context, ledger string) ApiCreateTransactionsRequest {
	return ApiCreateTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return TransactionsResponse
func (a *TransactionsApiService) CreateTransactionsExecute(r ApiCreateTransactionsRequest) (TransactionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  TransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.CreateTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions/batch"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.transactions == nil {
		return localVarReturnValue, nil, reportError("transactions is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transactions
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CreateTransactions400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v CreateTransaction409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTransactionRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	txid int32
}


func (r ApiGetTransactionRequest) Execute() (TransactionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetTransactionExecute(r)
}

/*
GetTransaction Get transaction from a ledger by its ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param txid Transaction ID.
 @return ApiGetTransactionRequest
*/
func (a *TransactionsApiService) GetTransaction(ctx _context.Context, ledger string, txid int32) ApiGetTransactionRequest {
	return ApiGetTransactionRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		txid: txid,
	}
}

// Execute executes the request
//  @return TransactionResponse
func (a *TransactionsApiService) GetTransactionExecute(r ApiGetTransactionRequest) (TransactionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  TransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GetTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions/{txid}"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"txid"+"}", _neturl.PathEscape(parameterToString(r.txid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetTransaction400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetTransaction404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListTransactionsRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	pageSize *int32
	after *string
	reference *string
	account *string
	source *string
	destination *string
	startTime *string
	endTime *string
	paginationToken *string
	metadata *map[string]interface{}
}

// The maximum number of results to return per page
func (r ApiListTransactionsRequest) PageSize(pageSize int32) ApiListTransactionsRequest {
	r.pageSize = &pageSize
	return r
}
// Pagination cursor, will return transactions after given txid (in descending order).
func (r ApiListTransactionsRequest) After(after string) ApiListTransactionsRequest {
	r.after = &after
	return r
}
// Find transactions by reference field.
func (r ApiListTransactionsRequest) Reference(reference string) ApiListTransactionsRequest {
	r.reference = &reference
	return r
}
// Find transactions with postings involving given account, either as source or destination.
func (r ApiListTransactionsRequest) Account(account string) ApiListTransactionsRequest {
	r.account = &account
	return r
}
// Find transactions with postings involving given account at source.
func (r ApiListTransactionsRequest) Source(source string) ApiListTransactionsRequest {
	r.source = &source
	return r
}
// Find transactions with postings involving given account at destination.
func (r ApiListTransactionsRequest) Destination(destination string) ApiListTransactionsRequest {
	r.destination = &destination
	return r
}
// Filter transactions that occurred after this timestamp. The format is RFC3339 and is inclusive (for example, 12:00:01 includes the first second of the minute). 
func (r ApiListTransactionsRequest) StartTime(startTime string) ApiListTransactionsRequest {
	r.startTime = &startTime
	return r
}
// Filter transactions that occurred before this timestamp. The format is RFC3339 and is exclusive (for example, 12:00:01 excludes the first second of the minute). 
func (r ApiListTransactionsRequest) EndTime(endTime string) ApiListTransactionsRequest {
	r.endTime = &endTime
	return r
}
// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results.  Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set. 
func (r ApiListTransactionsRequest) PaginationToken(paginationToken string) ApiListTransactionsRequest {
	r.paginationToken = &paginationToken
	return r
}
// Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
func (r ApiListTransactionsRequest) Metadata(metadata map[string]interface{}) ApiListTransactionsRequest {
	r.metadata = &metadata
	return r
}

func (r ApiListTransactionsRequest) Execute() (ListTransactions200Response, *_nethttp.Response, error) {
	return r.ApiService.ListTransactionsExecute(r)
}

/*
ListTransactions List transactions from a ledger.

List transactions from a ledger, sorted by txid in descending order.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiListTransactionsRequest
*/
func (a *TransactionsApiService) ListTransactions(ctx _context.Context, ledger string) ApiListTransactionsRequest {
	return ApiListTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return ListTransactions200Response
func (a *TransactionsApiService) ListTransactionsExecute(r ApiListTransactionsRequest) (ListTransactions200Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ListTransactions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.ListTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.reference != nil {
		localVarQueryParams.Add("reference", parameterToString(*r.reference, ""))
	}
	if r.account != nil {
		localVarQueryParams.Add("account", parameterToString(*r.account, ""))
	}
	if r.source != nil {
		localVarQueryParams.Add("source", parameterToString(*r.source, ""))
	}
	if r.destination != nil {
		localVarQueryParams.Add("destination", parameterToString(*r.destination, ""))
	}
	if r.startTime != nil {
		localVarQueryParams.Add("start_time", parameterToString(*r.startTime, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("end_time", parameterToString(*r.endTime, ""))
	}
	if r.paginationToken != nil {
		localVarQueryParams.Add("pagination_token", parameterToString(*r.paginationToken, ""))
	}
	if r.metadata != nil {
		localVarQueryParams.Add("metadata", parameterToString(*r.metadata, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ListAccounts400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRevertTransactionRequest struct {
	ctx _context.Context
	ApiService TransactionsApi
	ledger string
	txid int32
}


func (r ApiRevertTransactionRequest) Execute() (TransactionResponse, *_nethttp.Response, error) {
	return r.ApiService.RevertTransactionExecute(r)
}

/*
RevertTransaction Revert a ledger transaction by its ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param txid Transaction ID.
 @return ApiRevertTransactionRequest
*/
func (a *TransactionsApiService) RevertTransaction(ctx _context.Context, ledger string, txid int32) ApiRevertTransactionRequest {
	return ApiRevertTransactionRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		txid: txid,
	}
}

// Execute executes the request
//  @return TransactionResponse
func (a *TransactionsApiService) RevertTransactionExecute(r ApiRevertTransactionRequest) (TransactionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  TransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.RevertTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/transactions/{txid}/revert"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"txid"+"}", _neturl.PathEscape(parameterToString(r.txid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetTransaction400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetTransaction404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
