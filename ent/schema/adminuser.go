package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
	ent.Schema
}

// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(pulid.ID("")).Unique().Immutable().DefaultFunc(func() pulid.ID { return pulid.MustNew("US") }),
		field.String("name").NotEmpty(),
		field.String("api_key"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("last_used_at").Default(time.Now),
		field.Time("last_updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
	return nil
}
