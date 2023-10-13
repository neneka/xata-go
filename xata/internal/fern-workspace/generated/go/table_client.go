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
)

type TableClient interface {
	CreateTable(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*CreateTableResponse, error)
	UpdateTable(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *UpdateTableRequest) (*UpdateTableResponse, error)
	DeleteTable(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*DeleteTableResponse, error)
	GetTableSchema(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*GetTableSchemaResponse, error)
	SetTableSchema(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *SetTableSchemaRequest) (*SetTableSchemaResponse, error)
	GetTableColumns(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*GetTableColumnsResponse, error)
	AddTableColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *Column) (*AddTableColumnResponse, error)
	GetColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, columnName ColumnName) (*Column, error)
	UpdateColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, columnName ColumnName, request *UpdateColumnRequest) (*UpdateColumnResponse, error)
	DeleteColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, columnName ColumnName) (*DeleteColumnResponse, error)
}

func NewTableClient(opts ...core.ClientOption) TableClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &tableClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type tableClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Creates a new table with the given name. Returns 422 if a table with the same name already exists.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) CreateTable(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*CreateTableResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v", dbBranchName, tableName)

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

	var response *CreateTableResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Update table. Currently there is only one update operation supported: renaming the table by providing a new name.
//
// In the example below, we rename a table from “users” to “people”:
//
// ```json
// // PATCH /db/test:main/tables/users
//
//	{
//	  "name": "people"
//	}
//
// ```
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) UpdateTable(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *UpdateTableRequest) (*UpdateTableResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v", dbBranchName, tableName)

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

	var response *UpdateTableResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Deletes the table with the given name.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) DeleteTable(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*DeleteTableResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v", dbBranchName, tableName)

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

	var response *DeleteTableResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) GetTableSchema(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*GetTableSchemaResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/schema", dbBranchName, tableName)

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

	var response *GetTableSchemaResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) SetTableSchema(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *SetTableSchemaRequest) (*SetTableSchemaResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/schema", dbBranchName, tableName)

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
		case 409:
			value := new(ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *SetTableSchemaResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Retrieves the list of table columns and their definition. This endpoint returns the column list with object columns being reported with their
// full dot-separated path (flattened).
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) GetTableColumns(ctx context.Context, dbBranchName DbBranchName, tableName TableName) (*GetTableColumnsResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/columns", dbBranchName, tableName)

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

	var response *GetTableColumnsResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Adds a new column to the table. The body of the request should contain the column definition.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
func (t *tableClient) AddTableColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, request *Column) (*AddTableColumnResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/columns", dbBranchName, tableName)

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

	var response *AddTableColumnResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Get the definition of a single column.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Column name
func (t *tableClient) GetColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, columnName ColumnName) (*Column, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/columns/%v", dbBranchName, tableName, columnName)

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

	var response *Column
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Update column with partial data. Can be used for renaming the column by providing a new "name" field.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Column name
func (t *tableClient) UpdateColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, columnName ColumnName, request *UpdateColumnRequest) (*UpdateColumnResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/columns/%v", dbBranchName, tableName, columnName)

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

	var response *UpdateColumnResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Deletes the specified column.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Column name
func (t *tableClient) DeleteColumn(ctx context.Context, dbBranchName DbBranchName, tableName TableName, columnName ColumnName) (*DeleteColumnResponse, error) {
	baseURL := "/"
	if t.baseURL != "" {
		baseURL = t.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/columns/%v", dbBranchName, tableName, columnName)

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

	var response *DeleteColumnResponse
	if err := core.DoRequest(
		ctx,
		t.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		&response,
		false,
		t.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}