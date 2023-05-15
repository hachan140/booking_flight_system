package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Plane holds the schema definition for the Plane entity.
type Plane struct {
	ent.Schema
}

// Fields of the Plane.
func (Plane) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("name").NotEmpty(),
		field.Int64("economy_class_slots").Default(0),
		field.Int64("business_class_slots").Default(0),
		field.Enum("status").NamedValues("booked", "free").Default("free"),
	}
}

// Edges of the Plane.
func (Plane) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("flights", Flight.Type),
	}
}
func (Plane) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
