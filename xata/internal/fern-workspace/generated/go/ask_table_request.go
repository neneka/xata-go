// This file was auto-generated by Fern from our API Definition.

package api

// AskTableRequest is an in-lined request used by the AskTable endpoint.
type AskTableRequest struct {
	// The question you'd like to ask.
	Question string `json:"question"`
	// The type of search to use. If set to `keyword` (the default), the search can be configured by passing
	// a `search` object with the following fields. For more details about each, see the Search endpoint documentation.
	// All fields are optional.
	//   - fuzziness  - typo tolerance
	//   - target - columns to search into, and weights.
	//   - prefix - prefix search type.
	//   - filter - pre-filter before searching.
	//   - boosters - control relevancy.
	//
	// If set to `vector`, a `vectorSearch` object must be passed, with the following parameters. For more details, see the Vector
	// Search endpoint documentation. The `column` and `contentColumn` parameters are required.
	//   - column - the vector column containing the embeddings.
	//   - contentColumn - the column that contains the text from which the embeddings where computed.
	//   - filter - pre-filter before searching.
	SearchType   *AskTableRequestSearchType   `json:"searchType,omitempty"`
	Search       *AskTableRequestSearch       `json:"search,omitempty"`
	VectorSearch *AskTableRequestVectorSearch `json:"vectorSearch,omitempty"`
	Rules        *[]string                    `json:"rules,omitempty"`
}