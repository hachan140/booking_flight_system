package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("email").Unique().NotEmpty(),
		field.String("password").Sensitive(),
		field.String("phone_number").NotEmpty().Unique(),
		field.String("full_name").NotEmpty(),
		field.Time("dob").Optional(),
		field.String("cid").Unique(),
		field.Enum("member_type").NamedValues("Admin", "ADMIN", "Member", "MEMBER").Default("MEMBER").Optional(),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("has_customer", Customer.Type).Unique(),
	}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
