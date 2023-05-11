// Code generated by ent, DO NOT EDIT.

package flight

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the flight type in the database.
	Label = "flight"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDepartAt holds the string denoting the depart_at field in the database.
	FieldDepartAt = "depart_at"
	// FieldLandAt holds the string denoting the land_at field in the database.
	FieldLandAt = "land_at"
	// FieldAvailableEcSlot holds the string denoting the available_ec_slot field in the database.
	FieldAvailableEcSlot = "available_ec_slot"
	// FieldAvailableBcSlot holds the string denoting the available_bc_slot field in the database.
	FieldAvailableBcSlot = "available_bc_slot"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPlaneID holds the string denoting the plane_id field in the database.
	FieldPlaneID = "plane_id"
	// FieldAirportID holds the string denoting the airport_id field in the database.
	FieldAirportID = "airport_id"
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// EdgeHasPlane holds the string denoting the has_plane edge name in mutations.
	EdgeHasPlane = "has_plane"
	// EdgeHasBooking holds the string denoting the has_booking edge name in mutations.
	EdgeHasBooking = "has_booking"
	// EdgeHasAirport holds the string denoting the has_airport edge name in mutations.
	EdgeHasAirport = "has_Airport"
	// EdgeHasCustomer holds the string denoting the has_customer edge name in mutations.
	EdgeHasCustomer = "has_Customer"
	// Table holds the table name of the flight in the database.
	Table = "flights"
	// HasPlaneTable is the table that holds the has_plane relation/edge.
	HasPlaneTable = "flights"
	// HasPlaneInverseTable is the table name for the Plane entity.
	// It exists in this package in order to avoid circular dependency with the "plane" package.
	HasPlaneInverseTable = "planes"
	// HasPlaneColumn is the table column denoting the has_plane relation/edge.
	HasPlaneColumn = "plane_id"
	// HasBookingTable is the table that holds the has_booking relation/edge.
	HasBookingTable = "bookings"
	// HasBookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	HasBookingInverseTable = "bookings"
	// HasBookingColumn is the table column denoting the has_booking relation/edge.
	HasBookingColumn = "flight_id"
	// HasAirportTable is the table that holds the has_Airport relation/edge.
	HasAirportTable = "flights"
	// HasAirportInverseTable is the table name for the Airport entity.
	// It exists in this package in order to avoid circular dependency with the "airport" package.
	HasAirportInverseTable = "airports"
	// HasAirportColumn is the table column denoting the has_Airport relation/edge.
	HasAirportColumn = "airport_id"
	// HasCustomerTable is the table that holds the has_Customer relation/edge.
	HasCustomerTable = "flights"
	// HasCustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	HasCustomerInverseTable = "customers"
	// HasCustomerColumn is the table column denoting the has_Customer relation/edge.
	HasCustomerColumn = "customer_id"
)

// Columns holds all SQL columns for flight fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDepartAt,
	FieldLandAt,
	FieldAvailableEcSlot,
	FieldAvailableBcSlot,
	FieldStatus,
	FieldPlaneID,
	FieldAirportID,
	FieldCustomerID,
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
)

// Status defines the type for the "status" enum field.
type Status string

// StatusCanceled is the default value of the Status enum.
const DefaultStatus = StatusCanceled

// Status values.
const (
	StatusFlying   Status = "scheduled"
	StatusCanceled Status = "landed"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusFlying, StatusCanceled:
		return nil
	default:
		return fmt.Errorf("flight: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Flight queries.
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

// ByDepartAt orders the results by the depart_at field.
func ByDepartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartAt, opts...).ToFunc()
}

// ByLandAt orders the results by the land_at field.
func ByLandAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLandAt, opts...).ToFunc()
}

// ByAvailableEcSlot orders the results by the available_ec_slot field.
func ByAvailableEcSlot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailableEcSlot, opts...).ToFunc()
}

// ByAvailableBcSlot orders the results by the available_bc_slot field.
func ByAvailableBcSlot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailableBcSlot, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPlaneID orders the results by the plane_id field.
func ByPlaneID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlaneID, opts...).ToFunc()
}

// ByAirportID orders the results by the airport_id field.
func ByAirportID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAirportID, opts...).ToFunc()
}

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByHasPlaneField orders the results by has_plane field.
func ByHasPlaneField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasPlaneStep(), sql.OrderByField(field, opts...))
	}
}

// ByHasBookingCount orders the results by has_booking count.
func ByHasBookingCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHasBookingStep(), opts...)
	}
}

// ByHasBooking orders the results by has_booking terms.
func ByHasBooking(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasBookingStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHasAirportField orders the results by has_Airport field.
func ByHasAirportField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasAirportStep(), sql.OrderByField(field, opts...))
	}
}

// ByHasCustomerField orders the results by has_Customer field.
func ByHasCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newHasPlaneStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasPlaneInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HasPlaneTable, HasPlaneColumn),
	)
}
func newHasBookingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasBookingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HasBookingTable, HasBookingColumn),
	)
}
func newHasAirportStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasAirportInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HasAirportTable, HasAirportColumn),
	)
}
func newHasCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasCustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, HasCustomerTable, HasCustomerColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
