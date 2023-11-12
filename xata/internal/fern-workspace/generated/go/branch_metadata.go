// This file was auto-generated by Fern from our API Definition.

package api

type BranchMetadata struct {
	Branch *BranchName `json:"branch,omitempty"`
	Labels *[]string   `json:"labels,omitempty"`
	// <span style="white-space: nowrap">`non-empty`</span>
	Repository *string `json:"repository,omitempty"`
	// <span style="white-space: nowrap">`non-empty`</span>
	Stage *string `json:"stage,omitempty"`
}
