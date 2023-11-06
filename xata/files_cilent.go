package xata

import (
	"context"
	"fmt"

	xatagenworkspace "github.com/xataio/xata-go/xata/internal/fern-workspace/generated/go"
	xatagenclient "github.com/xataio/xata-go/xata/internal/fern-workspace/generated/go/core"
)

type FilesClient interface {
	// GetFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) error
	// PutFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) (*FileResponse, error)
	// DeleteFileItem(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName, fileId FileItemId) (*FileResponse, error)
	// GetFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) error
	// PutFile(ctx context.Context, dbBranchName DbBranchName, tableName TableName, recordId RecordId, columnName ColumnName) (*FileResponse, error)
	Delete(ctx context.Context, request DeleteFileRequest) (*xatagenworkspace.FileResponse, error)
}

type filesClient struct {
	generated  xatagenworkspace.FilesClient
	dbName     string
	branchName string
}

func (f filesClient) dbBranchName(request BranchRequestOptional) (string, error) {
	if request.DatabaseName == nil {
		if f.dbName == "" {
			return "", fmt.Errorf("database name cannot be empty")
		}
		request.DatabaseName = String(f.dbName)
	}

	if request.BranchName == nil {
		if f.branchName == "" {
			return "", fmt.Errorf("branch name cannot be empty")
		}
		request.BranchName = String(f.branchName)
	}

	return fmt.Sprintf("%s:%s", *request.DatabaseName, *request.BranchName), nil
}

type DeleteFileRequest struct {
	BranchRequestOptional
	TableName  string
	RecordId   string
	ColumnName string
}

func (f filesClient) Delete(ctx context.Context, request DeleteFileRequest) (*xatagenworkspace.FileResponse, error) {
	dbBranchName, err := f.dbBranchName(request.BranchRequestOptional)
	if err != nil {
		return nil, err
	}

	return f.generated.DeleteFile(ctx, dbBranchName, request.TableName, request.RecordId, request.ColumnName)
}

func NewFilesClient(opts ...ClientOption) (FilesClient, error) {
	cliOpts, dbCfg, err := consolidateClientOptionsForWorkspace(opts...)
	if err != nil {
		return nil, err
	}

	return filesClient{
			generated: xatagenworkspace.NewFilesClient(
				func(options *xatagenclient.ClientOptions) {
					options.HTTPClient = cliOpts.HTTPClient
					options.BaseURL = cliOpts.BaseURL
					options.Bearer = cliOpts.Bearer
				}),
			dbName:     dbCfg.dbName,
			branchName: dbCfg.branchName,
		},
		nil
}