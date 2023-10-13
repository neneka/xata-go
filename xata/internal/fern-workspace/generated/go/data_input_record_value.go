// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type DataInputRecordValue struct {
	typeName       string
	RecordId       RecordId
	String         string
	Boolean        bool
	Double         float64
	StringList     []string
	DoubleList     []float64
	DateTime       DateTime
	InputFileArray InputFileArray
	InputFile      *InputFile
}

func NewDataInputRecordValueFromRecordId(value RecordId) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "recordId", RecordId: value}
}

func NewDataInputRecordValueFromString(value string) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "string", String: value}
}

func NewDataInputRecordValueFromBoolean(value bool) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "boolean", Boolean: value}
}

func NewDataInputRecordValueFromDouble(value float64) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "double", Double: value}
}

func NewDataInputRecordValueFromStringList(value []string) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "stringList", StringList: value}
}

func NewDataInputRecordValueFromDoubleList(value []float64) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "doubleList", DoubleList: value}
}

func NewDataInputRecordValueFromDateTime(value DateTime) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "dateTime", DateTime: value}
}

func NewDataInputRecordValueFromInputFileArray(value InputFileArray) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "inputFileArray", InputFileArray: value}
}

func NewDataInputRecordValueFromInputFile(value *InputFile) *DataInputRecordValue {
	return &DataInputRecordValue{typeName: "inputFile", InputFile: value}
}

func (d *DataInputRecordValue) UnmarshalJSON(data []byte) error {
	var valueRecordId RecordId
	if err := json.Unmarshal(data, &valueRecordId); err == nil {
		d.typeName = "recordId"
		d.RecordId = valueRecordId
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		d.typeName = "string"
		d.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		d.typeName = "boolean"
		d.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		d.typeName = "double"
		d.Double = valueDouble
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		d.typeName = "stringList"
		d.StringList = valueStringList
		return nil
	}
	var valueDoubleList []float64
	if err := json.Unmarshal(data, &valueDoubleList); err == nil {
		d.typeName = "doubleList"
		d.DoubleList = valueDoubleList
		return nil
	}
	var valueDateTime DateTime
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		d.typeName = "dateTime"
		d.DateTime = valueDateTime
		return nil
	}
	var valueInputFileArray InputFileArray
	if err := json.Unmarshal(data, &valueInputFileArray); err == nil {
		d.typeName = "inputFileArray"
		d.InputFileArray = valueInputFileArray
		return nil
	}
	valueInputFile := new(InputFile)
	if err := json.Unmarshal(data, &valueInputFile); err == nil {
		d.typeName = "inputFile"
		d.InputFile = valueInputFile
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DataInputRecordValue) MarshalJSON() ([]byte, error) {
	switch d.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", d.typeName, d)
	case "recordId":
		return json.Marshal(d.RecordId)
	case "string":
		return json.Marshal(d.String)
	case "boolean":
		return json.Marshal(d.Boolean)
	case "double":
		return json.Marshal(d.Double)
	case "stringList":
		return json.Marshal(d.StringList)
	case "doubleList":
		return json.Marshal(d.DoubleList)
	case "dateTime":
		return json.Marshal(d.DateTime)
	case "inputFileArray":
		return json.Marshal(d.InputFileArray)
	case "inputFile":
		return json.Marshal(d.InputFile)
	}
}

type DataInputRecordValueVisitor interface {
	VisitRecordId(RecordId) error
	VisitString(string) error
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitStringList([]string) error
	VisitDoubleList([]float64) error
	VisitDateTime(DateTime) error
	VisitInputFileArray(InputFileArray) error
	VisitInputFile(*InputFile) error
}

func (d *DataInputRecordValue) Accept(v DataInputRecordValueVisitor) error {
	switch d.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", d.typeName, d)
	case "recordId":
		return v.VisitRecordId(d.RecordId)
	case "string":
		return v.VisitString(d.String)
	case "boolean":
		return v.VisitBoolean(d.Boolean)
	case "double":
		return v.VisitDouble(d.Double)
	case "stringList":
		return v.VisitStringList(d.StringList)
	case "doubleList":
		return v.VisitDoubleList(d.DoubleList)
	case "dateTime":
		return v.VisitDateTime(d.DateTime)
	case "inputFileArray":
		return v.VisitInputFileArray(d.InputFileArray)
	case "inputFile":
		return v.VisitInputFile(d.InputFile)
	}
}