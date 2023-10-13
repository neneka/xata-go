// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

// InsertRecordWithIdRequest is an in-lined request used by the InsertRecordWithId endpoint.
type InsertRecordWithIdRequest struct {
	CreateOnly *bool `json:"-"`
	IfVersion  *int  `json:"-"`
	// Column filters
	Columns []*string                        `json:"-"`
	Body    map[string]*DataInputRecordValue `json:"-"`
}

func (i *InsertRecordWithIdRequest) UnmarshalJSON(data []byte) error {
	var body map[string]*DataInputRecordValue
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.Body = body
	return nil
}

func (i *InsertRecordWithIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}