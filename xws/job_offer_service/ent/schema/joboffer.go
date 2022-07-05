package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JobOffer holds the schema definition for the JobOffer entity.
type JobOffer struct {
	ent.Schema
}

// Fields of the JobOffer.
func (JobOffer) Fields() []ent.Field {
	return []ent.Field{
		field.String("profile_id"),
		field.String("company"),
		field.String("position"),
		field.String("description"),
		field.String("criteria"),
		field.Time("created_at"),
	}
}

// Edges of the JobOffer.
func (JobOffer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requires", Skill.Type),
	}
}
