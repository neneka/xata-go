// This file was auto-generated by Fern from our API Definition.

package api

// SearchBranchRequest is an in-lined request used by the SearchBranch endpoint.
type SearchBranchRequest struct {
	// An array with the tables in which to search. By default, all tables are included. Optionally, filters can be included that apply to each table.
	Tables *[]*SearchBranchRequestTablesItem `json:"tables,omitempty"`
	// The query string. <span style="white-space: nowrap">`non-empty`</span>
	Query     string               `json:"query"`
	Fuzziness *FuzzinessExpression `json:"fuzziness,omitempty"`
	Prefix    *PrefixExpression    `json:"prefix,omitempty"`
	Highlight *HighlightExpression `json:"highlight,omitempty"`
	Page      *SearchPageConfig    `json:"page,omitempty"`
}