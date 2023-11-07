// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

type DbBranch struct {
	DatabaseName    DbName               `json:"databaseName"`
	BranchName      BranchName           `json:"branchName"`
	CreatedAt       DateTime             `json:"createdAt"`
	Id              string               `json:"id"`
	Version         float64              `json:"version"`
	LastMigrationId string               `json:"lastMigrationID"`
	Metadata        *BranchMetadata      `json:"metadata,omitempty"`
	StartedFrom     *StartedFromMetadata `json:"startedFrom,omitempty"`
	Schema          *Schema              `json:"schema,omitempty"`
}
