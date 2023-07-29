// Code generated by ent, DO NOT EDIT.

package script

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the script type in the database.
	Label = "script"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldNotify holds the string denoting the notify field in the database.
	FieldNotify = "notify"
	// FieldLastRun holds the string denoting the last_run field in the database.
	FieldLastRun = "last_run"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeReports holds the string denoting the reports edge name in mutations.
	EdgeReports = "reports"
	// Table holds the table name of the script in the database.
	Table = "scripts"
	// ReportsTable is the table that holds the reports relation/edge.
	ReportsTable = "reports"
	// ReportsInverseTable is the table name for the Report entity.
	// It exists in this package in order to avoid circular dependency with the "report" package.
	ReportsInverseTable = "reports"
	// ReportsColumn is the table column denoting the reports relation/edge.
	ReportsColumn = "script_reports"
)

// Columns holds all SQL columns for script fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldHostname,
	FieldNotify,
	FieldLastRun,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	HostnameValidator func(string) error
	// DefaultNotify holds the default value on creation for the "notify" field.
	DefaultNotify bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Script queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHostname orders the results by the hostname field.
func ByHostname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostname, opts...).ToFunc()
}

// ByNotify orders the results by the notify field.
func ByNotify(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotify, opts...).ToFunc()
}

// ByLastRun orders the results by the last_run field.
func ByLastRun(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastRun, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByReportsCount orders the results by reports count.
func ByReportsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReportsStep(), opts...)
	}
}

// ByReports orders the results by reports terms.
func ByReports(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReportsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newReportsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReportsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReportsTable, ReportsColumn),
	)
}
