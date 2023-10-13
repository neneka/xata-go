// This file was auto-generated by Fern from our API Definition.

package api

type MigrationObject struct {
	Title      *string        `json:"title,omitempty"`
	Message    *string        `json:"message,omitempty"`
	Id         string         `json:"id"`
	ParentId   *string        `json:"parentID,omitempty"`
	Checksum   string         `json:"checksum"`
	Operations []*MigrationOp `json:"operations,omitempty"`
}