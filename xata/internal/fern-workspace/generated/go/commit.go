// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

type Commit struct {
	Title         *string        `json:"title,omitempty"`
	Message       *string        `json:"message,omitempty"`
	Id            string         `json:"id"`
	ParentId      *string        `json:"parentID,omitempty"`
	Checksum      string         `json:"checksum"`
	MergeParentId *string        `json:"mergeParentID,omitempty"`
	CreatedAt     DateTime       `json:"createdAt"`
	Operations    []*MigrationOp `json:"operations,omitempty"`
}
