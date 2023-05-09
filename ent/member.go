// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-sytem/ent/customer"
	"booking-flight-sytem/ent/member"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// Dob holds the value of the "dob" field.
	Dob time.Time `json:"dob,omitempty"`
	// Cid holds the value of the "cid" field.
	Cid string `json:"cid,omitempty"`
	// Role holds the value of the "role" field.
	Role int `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberQuery when eager-loading is set.
	Edges            MemberEdges `json:"edges"`
	member_member_id *int
	selectValues     sql.SelectValues
}

// MemberEdges holds the relations/edges for other nodes in the graph.
type MemberEdges struct {
	// MemberID holds the value of the member_id edge.
	MemberID *Customer `json:"member_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MemberIDOrErr returns the MemberID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberEdges) MemberIDOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.MemberID == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.MemberID, nil
	}
	return nil, &NotLoadedError{edge: "member_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldID, member.FieldRole:
			values[i] = new(sql.NullInt64)
		case member.FieldEmail, member.FieldPassword, member.FieldPhoneNumber, member.FieldFullName, member.FieldCid:
			values[i] = new(sql.NullString)
		case member.FieldCreatedAt, member.FieldUpdatedAt, member.FieldDob:
			values[i] = new(sql.NullTime)
		case member.ForeignKeys[0]: // member_member_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case member.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case member.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case member.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				m.Email = value.String
			}
		case member.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				m.Password = value.String
			}
		case member.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				m.PhoneNumber = value.String
			}
		case member.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				m.FullName = value.String
			}
		case member.FieldDob:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dob", values[i])
			} else if value.Valid {
				m.Dob = value.Time
			}
		case member.FieldCid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cid", values[i])
			} else if value.Valid {
				m.Cid = value.String
			}
		case member.FieldRole:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				m.Role = int(value.Int64)
			}
		case member.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field member_member_id", value)
			} else if value.Valid {
				m.member_member_id = new(int)
				*m.member_member_id = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Member.
// This includes values selected through modifiers, order, etc.
func (m *Member) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMemberID queries the "member_id" edge of the Member entity.
func (m *Member) QueryMemberID() *CustomerQuery {
	return NewMemberClient(m.config).QueryMemberID(m)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return NewMemberClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(m.Email)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(m.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(m.FullName)
	builder.WriteString(", ")
	builder.WriteString("dob=")
	builder.WriteString(m.Dob.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("cid=")
	builder.WriteString(m.Cid)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", m.Role))
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member
