// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type FilterList struct {
	typeName             string
	FilterExpression     *FilterExpression
	FilterExpressionList []*FilterExpression
}

func NewFilterListFromFilterExpression(value *FilterExpression) *FilterList {
	return &FilterList{typeName: "filterExpression", FilterExpression: value}
}

func NewFilterListFromFilterExpressionList(value []*FilterExpression) *FilterList {
	return &FilterList{typeName: "filterExpressionList", FilterExpressionList: value}
}

func (f *FilterList) UnmarshalJSON(data []byte) error {
	valueFilterExpression := new(FilterExpression)
	if err := json.Unmarshal(data, &valueFilterExpression); err == nil {
		f.typeName = "filterExpression"
		f.FilterExpression = valueFilterExpression
		return nil
	}
	var valueFilterExpressionList []*FilterExpression
	if err := json.Unmarshal(data, &valueFilterExpressionList); err == nil {
		f.typeName = "filterExpressionList"
		f.FilterExpressionList = valueFilterExpressionList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, f)
}

func (f FilterList) MarshalJSON() ([]byte, error) {
	switch f.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", f.typeName, f)
	case "filterExpression":
		return json.Marshal(f.FilterExpression)
	case "filterExpressionList":
		return json.Marshal(f.FilterExpressionList)
	}
}

type FilterListVisitor interface {
	VisitFilterExpression(*FilterExpression) error
	VisitFilterExpressionList([]*FilterExpression) error
}

func (f *FilterList) Accept(v FilterListVisitor) error {
	switch f.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", f.typeName, f)
	case "filterExpression":
		return v.VisitFilterExpression(f.FilterExpression)
	case "filterExpressionList":
		return v.VisitFilterExpressionList(f.FilterExpressionList)
	}
}