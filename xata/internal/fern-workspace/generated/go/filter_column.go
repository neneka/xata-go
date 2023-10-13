// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type FilterColumn struct {
	typeName             string
	FilterColumnIncludes *FilterColumnIncludes
	FilterPredicate      *FilterPredicate
	FilterList           *FilterList
}

func NewFilterColumnFromFilterColumnIncludes(value *FilterColumnIncludes) *FilterColumn {
	return &FilterColumn{typeName: "filterColumnIncludes", FilterColumnIncludes: value}
}

func NewFilterColumnFromFilterPredicate(value *FilterPredicate) *FilterColumn {
	return &FilterColumn{typeName: "filterPredicate", FilterPredicate: value}
}

func NewFilterColumnFromFilterList(value *FilterList) *FilterColumn {
	return &FilterColumn{typeName: "filterList", FilterList: value}
}

func (f *FilterColumn) UnmarshalJSON(data []byte) error {
	valueFilterColumnIncludes := new(FilterColumnIncludes)
	if err := json.Unmarshal(data, &valueFilterColumnIncludes); err == nil {
		f.typeName = "filterColumnIncludes"
		f.FilterColumnIncludes = valueFilterColumnIncludes
		return nil
	}
	valueFilterPredicate := new(FilterPredicate)
	if err := json.Unmarshal(data, &valueFilterPredicate); err == nil {
		f.typeName = "filterPredicate"
		f.FilterPredicate = valueFilterPredicate
		return nil
	}
	valueFilterList := new(FilterList)
	if err := json.Unmarshal(data, &valueFilterList); err == nil {
		f.typeName = "filterList"
		f.FilterList = valueFilterList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, f)
}

func (f FilterColumn) MarshalJSON() ([]byte, error) {
	switch f.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", f.typeName, f)
	case "filterColumnIncludes":
		return json.Marshal(f.FilterColumnIncludes)
	case "filterPredicate":
		return json.Marshal(f.FilterPredicate)
	case "filterList":
		return json.Marshal(f.FilterList)
	}
}

type FilterColumnVisitor interface {
	VisitFilterColumnIncludes(*FilterColumnIncludes) error
	VisitFilterPredicate(*FilterPredicate) error
	VisitFilterList(*FilterList) error
}

func (f *FilterColumn) Accept(v FilterColumnVisitor) error {
	switch f.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", f.typeName, f)
	case "filterColumnIncludes":
		return v.VisitFilterColumnIncludes(f.FilterColumnIncludes)
	case "filterPredicate":
		return v.VisitFilterPredicate(f.FilterPredicate)
	case "filterList":
		return v.VisitFilterList(f.FilterList)
	}
}