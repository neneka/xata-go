// This file was auto-generated by Fern from our API Definition.

package api

type GetBranchSchemaHistoryResponseMeta struct {
	// last record id
	Cursor string `json:"cursor"`
	// true if more records can be fetch
	More bool `json:"more"`
}