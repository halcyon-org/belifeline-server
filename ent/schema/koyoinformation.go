package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
	"go.starlark.net/lib/time"
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
		// TODO: repeated ULID need_external
		field.JSON("params", map[string]string{}),
		field.JSON("scales", []float64{}),
		// TODO: repeated ULID koyo_data_ids
		field.String("version").NotEmpty(),
		field.String("license").NotEmpty(),
		// TODO: repeated string ext_licenses
		field.Enum("data_type").Values("unspecified", "image", "csv", "json"),
		field.Time("first_entry_at").Default(time.Now).Immutable(),
		field.Time("last_entry_at"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the KoyoInformation.
func (KoyoInformation) Edges() []ent.Edge {
	return nil
}
