package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Airport holds the schema definition for the Airport entity.
type Airport struct {
	ent.Schema
}

// Fields of the Airport.
func (Airport) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("name").NotEmpty(),
		field.Float("lat").Positive().Nillable(),
		field.Float("long").Positive().Nillable(),
	}
}

// Edges of the Airport.
func (Airport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("has_Flight", Flight.Type),
	}
}
