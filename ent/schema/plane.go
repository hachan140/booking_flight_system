package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Plane holds the schema definition for the Plane entity.
type Plane struct {
	ent.Schema
}

// Fields of the Plane.
func (Plane) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int64("economy_class_slots"),
		field.Int64("business_class_slots"),
		field.Enum("status").NamedValues("booked", "free").Default("free"),
	}
}

// Edges of the Plane.
func (Plane) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("flights", Flight.Type),
	}
}
