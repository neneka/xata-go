// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type MigrationTableOp struct {
	typeName                    string
	MigrationTableOpAddTable    *MigrationTableOpAddTable
	MigrationTableOpRemoveTable *MigrationTableOpRemoveTable
	MigrationTableOpRenameTable *MigrationTableOpRenameTable
}

func NewMigrationTableOpFromMigrationTableOpAddTable(value *MigrationTableOpAddTable) *MigrationTableOp {
	return &MigrationTableOp{typeName: "migrationTableOpAddTable", MigrationTableOpAddTable: value}
}

func NewMigrationTableOpFromMigrationTableOpRemoveTable(value *MigrationTableOpRemoveTable) *MigrationTableOp {
	return &MigrationTableOp{typeName: "migrationTableOpRemoveTable", MigrationTableOpRemoveTable: value}
}

func NewMigrationTableOpFromMigrationTableOpRenameTable(value *MigrationTableOpRenameTable) *MigrationTableOp {
	return &MigrationTableOp{typeName: "migrationTableOpRenameTable", MigrationTableOpRenameTable: value}
}

func (m *MigrationTableOp) UnmarshalJSON(data []byte) error {
	valueMigrationTableOpAddTable := new(MigrationTableOpAddTable)
	if err := json.Unmarshal(data, &valueMigrationTableOpAddTable); err == nil {
		m.typeName = "migrationTableOpAddTable"
		m.MigrationTableOpAddTable = valueMigrationTableOpAddTable
		return nil
	}
	valueMigrationTableOpRemoveTable := new(MigrationTableOpRemoveTable)
	if err := json.Unmarshal(data, &valueMigrationTableOpRemoveTable); err == nil {
		m.typeName = "migrationTableOpRemoveTable"
		m.MigrationTableOpRemoveTable = valueMigrationTableOpRemoveTable
		return nil
	}
	valueMigrationTableOpRenameTable := new(MigrationTableOpRenameTable)
	if err := json.Unmarshal(data, &valueMigrationTableOpRenameTable); err == nil {
		m.typeName = "migrationTableOpRenameTable"
		m.MigrationTableOpRenameTable = valueMigrationTableOpRenameTable
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MigrationTableOp) MarshalJSON() ([]byte, error) {
	switch m.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", m.typeName, m)
	case "migrationTableOpAddTable":
		return json.Marshal(m.MigrationTableOpAddTable)
	case "migrationTableOpRemoveTable":
		return json.Marshal(m.MigrationTableOpRemoveTable)
	case "migrationTableOpRenameTable":
		return json.Marshal(m.MigrationTableOpRenameTable)
	}
}

type MigrationTableOpVisitor interface {
	VisitMigrationTableOpAddTable(*MigrationTableOpAddTable) error
	VisitMigrationTableOpRemoveTable(*MigrationTableOpRemoveTable) error
	VisitMigrationTableOpRenameTable(*MigrationTableOpRenameTable) error
}

func (m *MigrationTableOp) Accept(v MigrationTableOpVisitor) error {
	switch m.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", m.typeName, m)
	case "migrationTableOpAddTable":
		return v.VisitMigrationTableOpAddTable(m.MigrationTableOpAddTable)
	case "migrationTableOpRemoveTable":
		return v.VisitMigrationTableOpRemoveTable(m.MigrationTableOpRemoveTable)
	case "migrationTableOpRenameTable":
		return v.VisitMigrationTableOpRenameTable(m.MigrationTableOpRenameTable)
	}
}
