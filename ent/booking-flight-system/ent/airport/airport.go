// Code generated by ent, DO NOT EDIT.

package airport

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the airport type in the database.
	Label = "airport"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLat holds the string denoting the lat field in the database.
	FieldLat = "lat"
	// FieldLong holds the string denoting the long field in the database.
	FieldLong = "long"
	// EdgeHasFlight holds the string denoting the has_flight edge name in mutations.
	EdgeHasFlight = "has_Flight"
	// Table holds the table name of the airport in the database.
	Table = "airports"
	// HasFlightTable is the table that holds the has_Flight relation/edge.
	HasFlightTable = "flights"
	// HasFlightInverseTable is the table name for the Flight entity.
	// It exists in this package in order to avoid circular dependency with the "flight" package.
	HasFlightInverseTable = "flights"
	// HasFlightColumn is the table column denoting the has_Flight relation/edge.
	HasFlightColumn = "airport_id"
)

// Columns holds all SQL columns for airport fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldLat,
	FieldLong,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LatValidator is a validator for the "lat" field. It is called by the builders before save.
	LatValidator func(float64) error
	// LongValidator is a validator for the "long" field. It is called by the builders before save.
	LongValidator func(float64) error
)

// OrderOption defines the ordering options for the Airport queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLat orders the results by the lat field.
func ByLat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLat, opts...).ToFunc()
}

// ByLong orders the results by the long field.
func ByLong(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLong, opts...).ToFunc()
}

// ByHasFlightCount orders the results by has_Flight count.
func ByHasFlightCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHasFlightStep(), opts...)
	}
}

// ByHasFlight orders the results by has_Flight terms.
func ByHasFlight(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasFlightStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHasFlightStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasFlightInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HasFlightTable, HasFlightColumn),
	)
}
