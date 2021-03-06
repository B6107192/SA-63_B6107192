package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{

		field.String("CLUBE_LOCATION_NAME").NotEmpty(),
		field.String("CLUBE_LOCATION_ADDRESS").NotEmpty(),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("club", Club.Type).StorageKey(edge.Column("location_id")),
	}
}
