// This file was auto-generated by Fern from our API Definition.

package api

type SchemaEditScript struct {
	SourceMigrationId *string        `json:"sourceMigrationID,omitempty"`
	TargetMigrationId *string        `json:"targetMigrationID,omitempty"`
	Operations        []*MigrationOp `json:"operations,omitempty"`
}