package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Script holds the schema definition for the Script entity.
type Script struct {
	ent.Schema
}

// Fields of the Script.
func (Script) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("hostname").NotEmpty(),
		field.Bool("notify").Default(false),
		field.Time("last_run").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Script.
func (Script) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reports", Report.Type),
	}
}
