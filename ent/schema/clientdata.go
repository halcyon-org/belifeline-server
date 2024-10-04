package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// ClientData holds the schema definition for the ClientData entity.
type ClientData struct {
	ent.Schema
}

// Fields of the ClientData.
func (ClientData) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(pulid.ID("")).Unique().Immutable().DefaultFunc(func() pulid.ID { return pulid.MustNew("US") }),
		field.String("username").NotEmpty(),
		field.String("api_key"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("last_used_at").Default(time.Now),
		field.Time("last_updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ClientData.
func (ClientData) Edges() []ent.Edge {
	return nil
}
