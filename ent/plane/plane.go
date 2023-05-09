// Code generated by ent, DO NOT EDIT.

package plane

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the plane type in the database.
	Label = "plane"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEconomyClassSlots holds the string denoting the economy_class_slots field in the database.
	FieldEconomyClassSlots = "economy_class_slots"
	// FieldBusinessClassSlots holds the string denoting the business_class_slots field in the database.
	FieldBusinessClassSlots = "business_class_slots"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeFlights holds the string denoting the flights edge name in mutations.
	EdgeFlights = "flights"
	// Table holds the table name of the plane in the database.
	Table = "planes"
	// FlightsTable is the table that holds the flights relation/edge.
	FlightsTable = "flights"
	// FlightsInverseTable is the table name for the Flight entity.
	// It exists in this package in order to avoid circular dependency with the "flight" package.
	FlightsInverseTable = "flights"
	// FlightsColumn is the table column denoting the flights relation/edge.
	FlightsColumn = "plane_id"
)

// Columns holds all SQL columns for plane fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEconomyClassSlots,
	FieldBusinessClassSlots,
	FieldStatus,
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
)

// Status defines the type for the "status" enum field.
type Status string

// StatusBooked is the default value of the Status enum.
const DefaultStatus = StatusBooked

// Status values.
const (
	StatusBooked Status = "free"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusBooked:
		return nil
	default:
		return fmt.Errorf("plane: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Plane queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEconomyClassSlots orders the results by the economy_class_slots field.
func ByEconomyClassSlots(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEconomyClassSlots, opts...).ToFunc()
}

// ByBusinessClassSlots orders the results by the business_class_slots field.
func ByBusinessClassSlots(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessClassSlots, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByFlightsCount orders the results by flights count.
func ByFlightsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFlightsStep(), opts...)
	}
}

// ByFlights orders the results by flights terms.
func ByFlights(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFlightsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFlightsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FlightsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FlightsTable, FlightsColumn),
	)
}
