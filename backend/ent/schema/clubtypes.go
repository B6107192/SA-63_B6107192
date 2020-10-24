package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubTypes holds the schema definition for the ClubTypes entity.
type ClubTypes struct {
	ent.Schema
}

// Fields of the ClubTypes.
func (ClubTypes) Fields() []ent.Field {
	return []ent.Field{

		field.String("CLUBE_TYPE_NAME").NotEmpty(),
	}
}

// Edges of the ClubTypes.
func (ClubTypes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("club", Club.Type).StorageKey(edge.Column("clubtypes_id")),
	}
}
