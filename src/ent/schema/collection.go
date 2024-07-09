package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("New Collection"),
		field.Int("project_id"),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requests", Request.Type),
		edge.From("project", Project.Type).
			Ref("collections").
			Unique().
			Field("project_id").
			Required(),
	}
}
func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}

}
