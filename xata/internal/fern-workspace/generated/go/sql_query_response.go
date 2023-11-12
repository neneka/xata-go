// This file was auto-generated by Fern from our API Definition.

package api

type SqlQueryResponse struct {
	Columns *map[string]any `json:"columns,omitempty"`
	Records *[]SqlRecord    `json:"records,omitempty"`
	// Number of selected columns
	Total   *int    `json:"total,omitempty"`
	Warning *string `json:"warning,omitempty"`
}
