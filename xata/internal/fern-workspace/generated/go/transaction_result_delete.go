// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

// A result from a delete operation.
type TransactionResultDelete struct {
	// The number of deleted rows
	Rows      int                       `json:"rows"`
	Columns   *TransactionResultColumns `json:"columns,omitempty"`
	operation string
}

func (t *TransactionResultDelete) Operation() string {
	return t.operation
}

func (t *TransactionResultDelete) UnmarshalJSON(data []byte) error {
	type unmarshaler TransactionResultDelete
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TransactionResultDelete(value)
	t.operation = "delete"
	return nil
}

func (t *TransactionResultDelete) MarshalJSON() ([]byte, error) {
	type embed TransactionResultDelete
	var marshaler = struct {
		embed
		Operation string `json:"operation"`
	}{
		embed:     embed(*t),
		Operation: "delete",
	}
	return json.Marshal(marshaler)
}
