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

type FilesClient interface {
	GetFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) error
	PutFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) (*FileResponse, error)
	DeleteFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) (*FileResponse, error)
	GetFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) error
	PutFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) (*FileResponse, error)
	DeleteFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) (*FileResponse, error)
}

func NewFilesClient(opts ...core.ClientOption) FilesClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &filesClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type filesClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Retrieves file content from an array by file ID
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
// The Column name
// The File Identifier
func (f *filesClient) GetFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) error {
	baseURL := "/"
	if f.baseURL != "" {
		baseURL = f.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v/column/%v/file/%v", dbBranchName, tableName, recordId, columnName, fileId)

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

	if err := core.DoRequest(
		ctx,
		f.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		nil,
		false,
		f.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

// Uploads the file content to an array given the file ID
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
// The Column name
// The File Identifier
func (f *filesClient) PutFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) (*FileResponse, error) {
	baseURL := "/"
	if f.baseURL != "" {
		baseURL = f.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v/column/%v/file/%v", dbBranchName, tableName, recordId, columnName, fileId)

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

	var response *FileResponse
	if err := core.DoRequest(
		ctx,
		f.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		f.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Deletes an item from an file array column given the file ID
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
// The Column name
// The File Identifier
func (f *filesClient) DeleteFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) (*FileResponse, error) {
	baseURL := "/"
	if f.baseURL != "" {
		baseURL = f.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v/column/%v/file/%v", dbBranchName, tableName, recordId, columnName, fileId)

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

	var response *FileResponse
	if err := core.DoRequest(
		ctx,
		f.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		&response,
		false,
		f.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Retrieves the file content from a file column
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
// The Column name
func (f *filesClient) GetFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) error {
	baseURL := "/"
	if f.baseURL != "" {
		baseURL = f.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v/column/%v/file", dbBranchName, tableName, recordId, columnName)

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

	if err := core.DoRequest(
		ctx,
		f.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		nil,
		false,
		f.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

// Uploads the file content to the given file column
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
// The Column name
func (f *filesClient) PutFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) (*FileResponse, error) {
	baseURL := "/"
	if f.baseURL != "" {
		baseURL = f.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v/column/%v/file", dbBranchName, tableName, recordId, columnName)

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

	var response *FileResponse
	if err := core.DoRequest(
		ctx,
		f.httpClient,
		endpointURL,
		http.MethodPut,
		nil,
		&response,
		false,
		f.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Deletes a file referred in a file column
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Table name
// The Record name
// The Column name
func (f *filesClient) DeleteFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) (*FileResponse, error) {
	baseURL := "/"
	if f.baseURL != "" {
		baseURL = f.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/tables/%v/data/%v/column/%v/file", dbBranchName, tableName, recordId, columnName)

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

	var response *FileResponse
	if err := core.DoRequest(
		ctx,
		f.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		&response,
		false,
		f.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}