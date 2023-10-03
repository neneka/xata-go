// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// The description of a single aggregation operation. It is an object with only one key-value pair.
// The key represents the aggregation type, while the value is an object with the configuration of
// the aggregation.
type AggExpression struct {
	typeName                      string
	AggExpressionCount            *AggExpressionCount
	AggExpressionSum              *AggExpressionSum
	AggExpressionMax              *AggExpressionMax
	AggExpressionMin              *AggExpressionMin
	AggExpressionAverage          *AggExpressionAverage
	AggExpressionUniqueCount      *AggExpressionUniqueCount
	AggExpressionDateHistogram    *AggExpressionDateHistogram
	AggExpressionTopValues        *AggExpressionTopValues
	AggExpressionNumericHistogram *AggExpressionNumericHistogram
}

func NewAggExpressionFromAggExpressionCount(value *AggExpressionCount) *AggExpression {
	return &AggExpression{typeName: "aggExpressionCount", AggExpressionCount: value}
}

func NewAggExpressionFromAggExpressionSum(value *AggExpressionSum) *AggExpression {
	return &AggExpression{typeName: "aggExpressionSum", AggExpressionSum: value}
}

func NewAggExpressionFromAggExpressionMax(value *AggExpressionMax) *AggExpression {
	return &AggExpression{typeName: "aggExpressionMax", AggExpressionMax: value}
}

func NewAggExpressionFromAggExpressionMin(value *AggExpressionMin) *AggExpression {
	return &AggExpression{typeName: "aggExpressionMin", AggExpressionMin: value}
}

func NewAggExpressionFromAggExpressionAverage(value *AggExpressionAverage) *AggExpression {
	return &AggExpression{typeName: "aggExpressionAverage", AggExpressionAverage: value}
}

func NewAggExpressionFromAggExpressionUniqueCount(value *AggExpressionUniqueCount) *AggExpression {
	return &AggExpression{typeName: "aggExpressionUniqueCount", AggExpressionUniqueCount: value}
}

func NewAggExpressionFromAggExpressionDateHistogram(value *AggExpressionDateHistogram) *AggExpression {
	return &AggExpression{typeName: "aggExpressionDateHistogram", AggExpressionDateHistogram: value}
}

func NewAggExpressionFromAggExpressionTopValues(value *AggExpressionTopValues) *AggExpression {
	return &AggExpression{typeName: "aggExpressionTopValues", AggExpressionTopValues: value}
}

func NewAggExpressionFromAggExpressionNumericHistogram(value *AggExpressionNumericHistogram) *AggExpression {
	return &AggExpression{typeName: "aggExpressionNumericHistogram", AggExpressionNumericHistogram: value}
}

func (a *AggExpression) UnmarshalJSON(data []byte) error {
	valueAggExpressionCount := new(AggExpressionCount)
	if err := json.Unmarshal(data, &valueAggExpressionCount); err == nil {
		a.typeName = "aggExpressionCount"
		a.AggExpressionCount = valueAggExpressionCount
		return nil
	}
	valueAggExpressionSum := new(AggExpressionSum)
	if err := json.Unmarshal(data, &valueAggExpressionSum); err == nil {
		a.typeName = "aggExpressionSum"
		a.AggExpressionSum = valueAggExpressionSum
		return nil
	}
	valueAggExpressionMax := new(AggExpressionMax)
	if err := json.Unmarshal(data, &valueAggExpressionMax); err == nil {
		a.typeName = "aggExpressionMax"
		a.AggExpressionMax = valueAggExpressionMax
		return nil
	}
	valueAggExpressionMin := new(AggExpressionMin)
	if err := json.Unmarshal(data, &valueAggExpressionMin); err == nil {
		a.typeName = "aggExpressionMin"
		a.AggExpressionMin = valueAggExpressionMin
		return nil
	}
	valueAggExpressionAverage := new(AggExpressionAverage)
	if err := json.Unmarshal(data, &valueAggExpressionAverage); err == nil {
		a.typeName = "aggExpressionAverage"
		a.AggExpressionAverage = valueAggExpressionAverage
		return nil
	}
	valueAggExpressionUniqueCount := new(AggExpressionUniqueCount)
	if err := json.Unmarshal(data, &valueAggExpressionUniqueCount); err == nil {
		a.typeName = "aggExpressionUniqueCount"
		a.AggExpressionUniqueCount = valueAggExpressionUniqueCount
		return nil
	}
	valueAggExpressionDateHistogram := new(AggExpressionDateHistogram)
	if err := json.Unmarshal(data, &valueAggExpressionDateHistogram); err == nil {
		a.typeName = "aggExpressionDateHistogram"
		a.AggExpressionDateHistogram = valueAggExpressionDateHistogram
		return nil
	}
	valueAggExpressionTopValues := new(AggExpressionTopValues)
	if err := json.Unmarshal(data, &valueAggExpressionTopValues); err == nil {
		a.typeName = "aggExpressionTopValues"
		a.AggExpressionTopValues = valueAggExpressionTopValues
		return nil
	}
	valueAggExpressionNumericHistogram := new(AggExpressionNumericHistogram)
	if err := json.Unmarshal(data, &valueAggExpressionNumericHistogram); err == nil {
		a.typeName = "aggExpressionNumericHistogram"
		a.AggExpressionNumericHistogram = valueAggExpressionNumericHistogram
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AggExpression) MarshalJSON() ([]byte, error) {
	switch a.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.typeName, a)
	case "aggExpressionCount":
		return json.Marshal(a.AggExpressionCount)
	case "aggExpressionSum":
		return json.Marshal(a.AggExpressionSum)
	case "aggExpressionMax":
		return json.Marshal(a.AggExpressionMax)
	case "aggExpressionMin":
		return json.Marshal(a.AggExpressionMin)
	case "aggExpressionAverage":
		return json.Marshal(a.AggExpressionAverage)
	case "aggExpressionUniqueCount":
		return json.Marshal(a.AggExpressionUniqueCount)
	case "aggExpressionDateHistogram":
		return json.Marshal(a.AggExpressionDateHistogram)
	case "aggExpressionTopValues":
		return json.Marshal(a.AggExpressionTopValues)
	case "aggExpressionNumericHistogram":
		return json.Marshal(a.AggExpressionNumericHistogram)
	}
}

type AggExpressionVisitor interface {
	VisitAggExpressionCount(*AggExpressionCount) error
	VisitAggExpressionSum(*AggExpressionSum) error
	VisitAggExpressionMax(*AggExpressionMax) error
	VisitAggExpressionMin(*AggExpressionMin) error
	VisitAggExpressionAverage(*AggExpressionAverage) error
	VisitAggExpressionUniqueCount(*AggExpressionUniqueCount) error
	VisitAggExpressionDateHistogram(*AggExpressionDateHistogram) error
	VisitAggExpressionTopValues(*AggExpressionTopValues) error
	VisitAggExpressionNumericHistogram(*AggExpressionNumericHistogram) error
}

func (a *AggExpression) Accept(v AggExpressionVisitor) error {
	switch a.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", a.typeName, a)
	case "aggExpressionCount":
		return v.VisitAggExpressionCount(a.AggExpressionCount)
	case "aggExpressionSum":
		return v.VisitAggExpressionSum(a.AggExpressionSum)
	case "aggExpressionMax":
		return v.VisitAggExpressionMax(a.AggExpressionMax)
	case "aggExpressionMin":
		return v.VisitAggExpressionMin(a.AggExpressionMin)
	case "aggExpressionAverage":
		return v.VisitAggExpressionAverage(a.AggExpressionAverage)
	case "aggExpressionUniqueCount":
		return v.VisitAggExpressionUniqueCount(a.AggExpressionUniqueCount)
	case "aggExpressionDateHistogram":
		return v.VisitAggExpressionDateHistogram(a.AggExpressionDateHistogram)
	case "aggExpressionTopValues":
		return v.VisitAggExpressionTopValues(a.AggExpressionTopValues)
	case "aggExpressionNumericHistogram":
		return v.VisitAggExpressionNumericHistogram(a.AggExpressionNumericHistogram)
	}
}
