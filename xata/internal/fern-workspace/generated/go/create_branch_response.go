// This file was auto-generated by Fern from our API Definition.

package api

type CreateBranchResponse struct {
	BranchName string `json:"branchName"`
	// <span style="white-space: nowrap">`non-empty`</span>
	DatabaseName string          `json:"databaseName"`
	Status       MigrationStatus `json:"status,omitempty"`
}
