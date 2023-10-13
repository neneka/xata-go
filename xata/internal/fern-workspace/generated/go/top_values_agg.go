// This file was auto-generated by Fern from our API Definition.

package api

// Split data into buckets by the unique values in a column. Accepts sub-aggregations for each bucket.
// The top values as ordered by the number of records (`$count`) are returned.
type TopValuesAgg struct {
	// The column to use for bucketing. Accepted types are `string`, `email`, `int`, `float`, or `bool`.
	Column string            `json:"column"`
	Aggs   *AggExpressionMap `json:"aggs,omitempty"`
	// The maximum number of unique values to return.
	Size *int `json:"size,omitempty"`
}