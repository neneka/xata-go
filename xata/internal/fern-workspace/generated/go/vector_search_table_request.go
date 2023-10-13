// This file was auto-generated by Fern from our API Definition.

package api

// VectorSearchTableRequest is an in-lined request used by the VectorSearchTable endpoint.
type VectorSearchTableRequest struct {
	// The vector to search for similarities. Must have the same dimension as
	// the vector column used.
	QueryVector []float64 `json:"queryVector,omitempty"`
	// The vector column in which to search. It must be of type `vector`.
	Column string `json:"column"`
	// The function used to measure the distance between two points. Can be one of:
	// `cosineSimilarity`, `l1`, `l2`. The default is `cosineSimilarity`.
	SimilarityFunction *string `json:"similarityFunction,omitempty"`
	// Number of results to return.
	Size   *int              `json:"size,omitempty"`
	Filter *FilterExpression `json:"filter,omitempty"`
}