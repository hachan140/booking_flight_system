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
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
		field.String("email").Unique().NotEmpty().Annotations(entgql.OrderField("EMAIL")),
		field.String("password").Sensitive().Annotations(entgql.OrderField("PASSWORD")),
		field.String("phone_number").NotEmpty().Unique().Annotations(entgql.OrderField("PHONE_NUMBER")),
		field.String("full_name").NotEmpty().Annotations(entgql.OrderField("FULL_NAME")),
		field.Time("dob").Optional().Annotations(entgql.OrderField("DOB")),
		field.String("cid").Unique().Annotations(entgql.OrderField("CID")),
		field.Enum("member_type").NamedValues("Admin", "ADMIN", "Member", "MEMBER").Default("MEMBER").Optional().Annotations(entgql.OrderField("MEMBER_TYPE")),
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
