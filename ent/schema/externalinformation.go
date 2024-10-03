package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// ExternalInformation holds the schema definition for the ExternalInformation entity.
type ExternalInformation struct {
	ent.Schema
}

// Fields of the ExternalInformation.
func (ExternalInformation) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(pulid.ID("")).Unique().Immutable().DefaultFunc(func() pulid.ID { return pulid.MustNew("US") }),
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("license").NotEmpty(),
		field.String("license_description"),
		field.Time("first_entry_at").Default(time.Now).Immutable(),
		field.Time("last_updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ExternalInformation.
func (ExternalInformation) Edges() []ent.Edge {
	return nil
}
