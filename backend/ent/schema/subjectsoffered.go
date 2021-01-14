package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// SubjectsOffered holds the schema definition for the SubjectsOffered entity.
type SubjectsOffered struct {
	ent.Schema
}

// Fields of the SubjectsOffered.
func (SubjectsOffered) Fields() []ent.Field {
	return []ent.Field{

		field.String("AMOUNT").NotEmpty(),
		field.String("STATUS").NotEmpty(),
	}
}

// Edges of the SubjectsOffered.
func (SubjectsOffered) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Subject", Subject.Type).
			Ref("SubjectsOffered").Unique(),
		edge.From("EducationLevel", EducationLevel.Type).
			Ref("SubjectsOffered").Unique(),
		edge.From("Year", Year.Type).
			Ref("SubjectsOffered").Unique(),
	}
}
