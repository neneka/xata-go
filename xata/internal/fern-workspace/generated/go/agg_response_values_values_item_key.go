// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type AggResponseValuesValuesItemKey struct {
	typeName string
	String   string
	Double   float64
}

func NewAggResponseValuesValuesItemKeyFromString(value string) *AggResponseValuesValuesItemKey {
	return &AggResponseValuesValuesItemKey{typeName: "string", String: value}
}

func NewAggResponseValuesValuesItemKeyFromDouble(value float64) *AggResponseValuesValuesItemKey {
	return &AggResponseValuesValuesItemKey{typeName: "double", Double: value}
}

func (a *AggResponseValuesValuesItemKey) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typeName = "string"
		a.String = valueString
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typeName = "double"
		a.Double = valueDouble
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AggResponseValuesValuesItemKey) MarshalJSON() ([]byte, error) {
	switch a.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.typeName, a)
	case "string":
		return json.Marshal(a.String)
	case "double":
		return json.Marshal(a.Double)
	}
}

type AggResponseValuesValuesItemKeyVisitor interface {
	VisitString(string) error
	VisitDouble(float64) error
}

func (a *AggResponseValuesValuesItemKey) Accept(v AggResponseValuesValuesItemKeyVisitor) error {
	switch a.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", a.typeName, a)
	case "string":
		return v.VisitString(a.String)
	case "double":
		return v.VisitDouble(a.Double)
	}
}
