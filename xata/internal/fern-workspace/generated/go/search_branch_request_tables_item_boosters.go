// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

type SearchBranchRequestTablesItemBoosters struct {
	// The name of the table.
	Table    string                `json:"table"`
	Filter   *FilterExpression     `json:"filter,omitempty"`
	Target   *TargetExpression     `json:"target,omitempty"`
	Boosters *[]*BoosterExpression `json:"boosters,omitempty"`
}
