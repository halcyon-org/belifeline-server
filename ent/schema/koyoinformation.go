package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// KoyoInformation holds the schema definition for the KoyoInformation entity.
type KoyoInformation struct {
	ent.Schema
}

// Fields of the KoyoInformation.
func (KoyoInformation) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(pulid.ID("")).Unique().Immutable().DefaultFunc(func() pulid.ID { return pulid.MustNew("US") }),
		field.String("name").NotEmpty(),
		field.String("description"),
		field.JSON("params", map[string]string{}),
		field.JSON("scales", []float64{}),
		field.String("version").NotEmpty(),
		field.String("license").NotEmpty(),
		field.Enum("data_type").Values("unspecified", "image", "csv", "json"),
		field.Time("first_entry_at").Default(time.Now).Immutable(),
		field.Time("last_entry_at"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the KoyoInformation.
func (KoyoInformation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("externals", ExternalInformation.Type),
		edge.To("data", KoyoData.Type),
	}
}
