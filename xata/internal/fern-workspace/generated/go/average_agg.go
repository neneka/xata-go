// This file was auto-generated by Fern from our API Definition.

package api

// The average of the numeric values in a particular column.
type AverageAgg struct {
	// The column on which to compute the average. Must be a numeric type.
	Column string `json:"column"`
}