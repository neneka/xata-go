// This file was auto-generated by Fern from our API Definition.

package api

// CompareBranchWithUserSchemaRequest is an in-lined request used by the CompareBranchWithUserSchema endpoint.
type CompareBranchWithUserSchemaRequest struct {
	Schema           *Schema         `json:"schema,omitempty"`
	SchemaOperations *[]*MigrationOp `json:"schemaOperations,omitempty"`
	BranchOperations *[]*MigrationOp `json:"branchOperations,omitempty"`
}