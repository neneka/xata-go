// This file was auto-generated by Fern from our API Definition.

package api

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	core "github.com/xataio/xata-go/xata/internal/fern-workspace/generated/go/core"
	io "io"
	http "net/http"
	url "net/url"
)

type RecordsClient interface {
	BranchTransaction(ctx context.Context, dbBranchName DbBranchName, request *BranchTransactionRequest) (*TransactionSuccess, error)
	InsertRecord(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *InsertRecordRequest) (*InsertRecordResponse, error)
	GetRecord(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *GetRecordRequest) (*Record, error)
	UpsertRecordWithId(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *UpsertRecordWithIdRequest) (*UpsertRecordWithIdResponse, error)
	InsertRecordWithId(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *InsertRecordWithIdRequest) (*InsertRecordWithIdResponse, error)
	UpdateRecordWithId(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *UpdateRecordWithIdRequest) (*UpdateRecordWithIdResponse, error)
	DeleteRecord(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId) error
	BulkInsertTableRecords(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *BulkInsertTableRecordsRequest) (*BulkInsertTableRecordsResponse, error)
}

func NewRecordsClient(opts ...core.ClientOption) RecordsClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &recordsClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type recordsClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (r *recordsClient) BranchTransaction(ctx context.Context, dbBranchName DbBranchName, request *BranchTransactionRequest) (*TransactionSuccess, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/transaction", dbBranchName)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 429:
			value := new(TooManyRequestsError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *TransactionSuccess
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Insert a new Record into the Table
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (r *recordsClient) InsertRecord(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *InsertRecordRequest) (*InsertRecordResponse, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data", dbBranchName, tableName)

	queryParams := make(url.Values)
	for _, value := range request.Columns {
		queryParams.Add("columns", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *InsertRecordResponse
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Retrieve record by ID
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
func (r *recordsClient) GetRecord(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *GetRecordRequest) (*Record, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v", dbBranchName, tableName, recordId)

	queryParams := make(url.Values)
	for _, value := range request.Columns {
		queryParams.Add("columns", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *Record
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
func (r *recordsClient) UpsertRecordWithId(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *UpsertRecordWithIdRequest) (*UpsertRecordWithIdResponse, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v", dbBranchName, tableName, recordId)

	queryParams := make(url.Values)
	if request.IfVersion != nil {
		queryParams.Add("ifVersion", fmt.Sprintf("%v", *request.IfVersion))
	}
	for _, value := range request.Columns {
		queryParams.Add("columns", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *UpsertRecordWithIdResponse
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// By default, IDs are auto-generated when data is insterted into Xata. Sending a request to this endpoint allows us to insert a record with a pre-existing ID, bypassing the default automatic ID generation.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
func (r *recordsClient) InsertRecordWithId(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *InsertRecordWithIdRequest) (*InsertRecordWithIdResponse, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v", dbBranchName, tableName, recordId)

	queryParams := make(url.Values)
	if request.CreateOnly != nil {
		queryParams.Add("createOnly", fmt.Sprintf("%v", *request.CreateOnly))
	}
	if request.IfVersion != nil {
		queryParams.Add("ifVersion", fmt.Sprintf("%v", *request.IfVersion))
	}
	for _, value := range request.Columns {
		queryParams.Add("columns", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *InsertRecordWithIdResponse
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
func (r *recordsClient) UpdateRecordWithId(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, request *UpdateRecordWithIdRequest) (*UpdateRecordWithIdResponse, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v", dbBranchName, tableName, recordId)

	queryParams := make(url.Values)
	if request.IfVersion != nil {
		queryParams.Add("ifVersion", fmt.Sprintf("%v", *request.IfVersion))
	}
	for _, value := range request.Columns {
		queryParams.Add("columns", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *UpdateRecordWithIdResponse
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
func (r *recordsClient) DeleteRecord(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId) error {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v", dbBranchName, tableName, recordId)

	queryParams := make(url.Values)
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	//var response *Record
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		nil,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

// Bulk insert records
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (r *recordsClient) BulkInsertTableRecords(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *BulkInsertTableRecordsRequest) (*BulkInsertTableRecordsResponse, error) {
	baseURL := "/"
	if r.baseURL != "" {
		baseURL = r.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/bulk", dbBranchName, tableName)

	queryParams := make(url.Values)
	for _, value := range request.Columns {
		queryParams.Add("columns", fmt.Sprintf("%v", *value))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 404:
			value := new(NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *BulkInsertTableRecordsResponse
	if err := core.DoRequest(
		ctx,
		r.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		r.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
