package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Club holds the schema definition for the Club entity.
type Club struct {
	ent.Schema
}

// Fields of the Club.
func (Club) Fields() []ent.Field {
	return []ent.Field{
		
		field.String("CLUB_NAME").NotEmpty(),
		
	}
}

// Edges of the Club.
func (Club) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("club").Unique(),
		edge.From("clubtypes", ClubTypes.Type).
			Ref("club").Unique(),
		edge.From("user", User.Type).
			Ref("club").Unique(),
		edge.From("activity", Activity.Type).
			Ref("club").Unique(),
	}
}
