// This file was auto-generated by Fern from our API Definition.

package api

type GetBranchSchemaHistoryResponse struct {
	Meta *GetBranchSchemaHistoryResponseMeta `json:"meta,omitempty"`
	Logs []*Commit                           `json:"logs,omitempty"`
}