// This file was auto-generated by Fern from our API Definition.

package api

type CreateTableResponse struct {
	BranchName string `json:"branchName"`
	// <span style="white-space: nowrap">`non-empty`</span>
	TableName string          `json:"tableName"`
	Status    MigrationStatus `json:"status,omitempty"`
}