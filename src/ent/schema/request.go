package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("New Request"),
		field.Int("collection_id"),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("collection", Collection.Type).
			Ref("requests").
			Unique().
			Field("collection_id").
			Required(),
	}
}

func (Request) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
