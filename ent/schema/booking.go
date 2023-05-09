package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("code").NotEmpty().Unique(),
		field.Enum("status").Values("success", "fail", "cancel"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return nil
}
