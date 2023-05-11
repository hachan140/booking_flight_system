// Code generated by ent, DO NOT EDIT.

package customer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldDob holds the string denoting the dob field in the database.
	FieldDob = "dob"
	// FieldCid holds the string denoting the cid field in the database.
	FieldCid = "cid"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// EdgeHasMember holds the string denoting the has_member edge name in mutations.
	EdgeHasMember = "has_member"
	// EdgeHasFlight holds the string denoting the has_flight edge name in mutations.
	EdgeHasFlight = "has_flight"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// HasMemberTable is the table that holds the has_member relation/edge.
	HasMemberTable = "customers"
	// HasMemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	HasMemberInverseTable = "members"
	// HasMemberColumn is the table column denoting the has_member relation/edge.
	HasMemberColumn = "member_id"
	// HasFlightTable is the table that holds the has_flight relation/edge.
	HasFlightTable = "flights"
	// HasFlightInverseTable is the table name for the Flight entity.
	// It exists in this package in order to avoid circular dependency with the "flight" package.
	HasFlightInverseTable = "flights"
	// HasFlightColumn is the table column denoting the has_flight relation/edge.
	HasFlightColumn = "customer_id"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmail,
	FieldPhoneNumber,
	FieldFullName,
	FieldDob,
	FieldCid,
	FieldMemberID,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	FullNameValidator func(string) error
	// CidValidator is a validator for the "cid" field. It is called by the builders before save.
	CidValidator func(string) error
)

// OrderOption defines the ordering options for the Customer queries.
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

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByDob orders the results by the dob field.
func ByDob(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDob, opts...).ToFunc()
}

// ByCid orders the results by the cid field.
func ByCid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCid, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByHasMemberField orders the results by has_member field.
func ByHasMemberField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasMemberStep(), sql.OrderByField(field, opts...))
	}
}

// ByHasFlightCount orders the results by has_flight count.
func ByHasFlightCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHasFlightStep(), opts...)
	}
}

// ByHasFlight orders the results by has_flight terms.
func ByHasFlight(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasFlightStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHasMemberStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasMemberInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, HasMemberTable, HasMemberColumn),
	)
}
func newHasFlightStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasFlightInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HasFlightTable, HasFlightColumn),
	)
}
