// Code generated by ent, DO NOT EDIT.

package report

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/zibbp/eros/internal/utils"
)

const (
	// Label holds the string label denoting the report type in the database.
	Label = "report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldS3File holds the string denoting the s3_file field in the database.
	FieldS3File = "s3_file"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeScript holds the string denoting the script edge name in mutations.
	EdgeScript = "script"
	// Table holds the table name of the report in the database.
	Table = "reports"
	// ScriptTable is the table that holds the script relation/edge.
	ScriptTable = "reports"
	// ScriptInverseTable is the table name for the Script entity.
	// It exists in this package in order to avoid circular dependency with the "script" package.
	ScriptInverseTable = "scripts"
	// ScriptColumn is the table column denoting the script relation/edge.
	ScriptColumn = "script_reports"
)

// Columns holds all SQL columns for report fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStatus,
	FieldS3File,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reports"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"script_reports",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s utils.ReportStatus) error {
	switch s {
	case "success", "failed":
		return nil
	default:
		return fmt.Errorf("report: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Report queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByS3File orders the results by the s3_file field.
func ByS3File(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldS3File, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByScriptField orders the results by script field.
func ByScriptField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScriptStep(), sql.OrderByField(field, opts...))
	}
}
func newScriptStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScriptInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ScriptTable, ScriptColumn),
	)
}
