// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ColumnType uint8

const (
	ColumnTypeBool ColumnType = iota + 1
	ColumnTypeInt
	ColumnTypeFloat
	ColumnTypeString
	ColumnTypeText
	ColumnTypeEmail
	ColumnTypeMultiple
	ColumnTypeLink
	ColumnTypeObject
	ColumnTypeDatetime
	ColumnTypeVector
	ColumnTypeFile
	ColumnTypeFileMap
	ColumnTypeJson
)

func (c ColumnType) String() string {
	switch c {
	default:
		return strconv.Itoa(int(c))
	case ColumnTypeBool:
		return "bool"
	case ColumnTypeInt:
		return "int"
	case ColumnTypeFloat:
		return "float"
	case ColumnTypeString:
		return "string"
	case ColumnTypeText:
		return "text"
	case ColumnTypeEmail:
		return "email"
	case ColumnTypeMultiple:
		return "multiple"
	case ColumnTypeLink:
		return "link"
	case ColumnTypeObject:
		return "object"
	case ColumnTypeDatetime:
		return "datetime"
	case ColumnTypeVector:
		return "vector"
	case ColumnTypeFile:
		return "file"
	case ColumnTypeFileMap:
		return "file[]"
	case ColumnTypeJson:
		return "json"
	}
}

func (c ColumnType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", c.String())), nil
}

func (c *ColumnType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "bool":
		value := ColumnTypeBool
		*c = value
	case "int":
		value := ColumnTypeInt
		*c = value
	case "float":
		value := ColumnTypeFloat
		*c = value
	case "string":
		value := ColumnTypeString
		*c = value
	case "text":
		value := ColumnTypeText
		*c = value
	case "email":
		value := ColumnTypeEmail
		*c = value
	case "multiple":
		value := ColumnTypeMultiple
		*c = value
	case "link":
		value := ColumnTypeLink
		*c = value
	case "object":
		value := ColumnTypeObject
		*c = value
	case "datetime":
		value := ColumnTypeDatetime
		*c = value
	case "vector":
		value := ColumnTypeVector
		*c = value
	case "file":
		value := ColumnTypeFile
		*c = value
	case "fileMap":
		value := ColumnTypeFileMap
		*c = value
	case "json":
		value := ColumnTypeJson
		*c = value
	}
	return nil
}
