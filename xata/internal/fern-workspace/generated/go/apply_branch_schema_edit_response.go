// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

type ApplyBranchSchemaEditResponse struct {
	// <span style="white-space: nowrap">`non-empty`</span>
	MigrationId       string          `json:"migrationID"`
	ParentMigrationId string          `json:"parentMigrationID"`
	Status            MigrationStatus `json:"status,omitempty"`
}
