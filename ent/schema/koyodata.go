package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// KoyoData holds the schema definition for the KoyoData entity.
type KoyoData struct {
	ent.Schema
}

// Fields of the KoyoData.
func (KoyoData) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(pulid.ID("")).
			Unique().
			Immutable().
			DefaultFunc(func() pulid.ID { return pulid.MustNew("US") }),
		field.String("koyo_id").GoType(pulid.ID("")).Unique(),
		field.Float("scale"),
		field.JSON("params", map[string]string{}),
		field.String("version").NotEmpty(),
		field.Time("entry_at").Default(time.Now).Immutable(),
		field.Time("target_at").Immutable(),
	}
}

// Edges of the KoyoData.
func (KoyoData) Edges() []ent.Edge {
	return nil
}
