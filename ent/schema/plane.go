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
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Int64("economy_class_slots").Default(0).Annotations(entgql.OrderField("ECONOMY_CLASS_SLOTS")),
		field.Int64("business_class_slots").Default(0).Annotations(entgql.OrderField("BUSINESS_CLASS_SLOTS")),
		field.Enum("status").NamedValues("booked", "BOOKED", "free", "FREE").Default("FREE").Annotations(entgql.OrderField("STATUS")),
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
