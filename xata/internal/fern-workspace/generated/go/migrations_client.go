// SPDX-License-Identifier: Apache-2.0

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

type MigrationsClient interface {
	ApplyBranchSchemaEdit(ctx context.Context, dbBranchName DbBranchName, request *ApplyBranchSchemaEditRequest) (*ApplyBranchSchemaEditResponse, error)
	CompareBranchWithUserSchema(ctx context.Context, dbBranchName DbBranchName, request *CompareBranchWithUserSchemaRequest) (*CompareBranchWithUserSchemaResponse, error)
	CompareBranchSchemas(ctx context.Context, dbBranchName DbBranchName, branchName BranchName, request *CompareBranchSchemasRequest) (*CompareBranchSchemasResponse, error)
	GetBranchSchemaHistory(ctx context.Context, dbBranchName DbBranchName, request *GetBranchSchemaHistoryRequest) (*GetBranchSchemaHistoryResponse, error)
	PreviewBranchSchemaEdit(ctx context.Context, dbBranchName DbBranchName, request *PreviewBranchSchemaEditRequest) (*PreviewBranchSchemaEditResponse, error)
	PushBranchMigrations(ctx context.Context, dbBranchName DbBranchName, request *PushBranchMigrationsRequest) (*PushBranchMigrationsResponse, error)
	UpdateBranchSchema(ctx context.Context, dbBranchName DbBranchName, request *Migration) (*UpdateBranchSchemaResponse, error)
}

func NewMigrationsClient(opts ...core.ClientOption) MigrationsClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &migrationsClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type migrationsClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (m *migrationsClient) ApplyBranchSchemaEdit(ctx context.Context, dbBranchName DbBranchName, request *ApplyBranchSchemaEditRequest) (*ApplyBranchSchemaEditResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/apply", dbBranchName)

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

	var response *ApplyBranchSchemaEditResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (m *migrationsClient) CompareBranchWithUserSchema(ctx context.Context, dbBranchName DbBranchName, request *CompareBranchWithUserSchemaRequest) (*CompareBranchWithUserSchemaResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/compare", dbBranchName)

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

	var response *CompareBranchWithUserSchemaResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
//
// The Database Name
func (m *migrationsClient) CompareBranchSchemas(ctx context.Context, dbBranchName DbBranchName, branchName BranchName, request *CompareBranchSchemasRequest) (*CompareBranchSchemasResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/compare/%v", dbBranchName, branchName)

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

	var response *CompareBranchSchemasResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (m *migrationsClient) GetBranchSchemaHistory(ctx context.Context, dbBranchName DbBranchName, request *GetBranchSchemaHistoryRequest) (*GetBranchSchemaHistoryResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/history", dbBranchName)

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

	var response *GetBranchSchemaHistoryResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (m *migrationsClient) PreviewBranchSchemaEdit(ctx context.Context, dbBranchName DbBranchName, request *PreviewBranchSchemaEditRequest) (*PreviewBranchSchemaEditResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/preview", dbBranchName)

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

	var response *PreviewBranchSchemaEditResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The `schema/push` API accepts a list of migrations to be applied to the
// current branch. A list of applicable migrations can be fetched using
// the `schema/history` API from another branch or database.
//
// The most recent migration must be part of the list or referenced (via
// `parentID`) by the first migration in the list of migrations to be pushed.
//
// Each migration in the list has an `id`, `parentID`, and `checksum`. The
// checksum for migrations are generated and verified by xata. The
// operation fails if any migration in the list has an invalid checksum.
//
// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (m *migrationsClient) PushBranchMigrations(ctx context.Context, dbBranchName DbBranchName, request *PushBranchMigrationsRequest) (*PushBranchMigrationsResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/push", dbBranchName)

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

	var response *PushBranchMigrationsResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// The DBBranchName matches the pattern `{db_name}:{branch_name}`.
func (m *migrationsClient) UpdateBranchSchema(ctx context.Context, dbBranchName DbBranchName, request *Migration) (*UpdateBranchSchemaResponse, error) {
	baseURL := "/"
	if m.baseURL != "" {
		baseURL = m.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"db/%v/schema/update", dbBranchName)

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

	var response *UpdateBranchSchemaResponse
	if err := core.DoRequest(
		ctx,
		m.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		m.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
