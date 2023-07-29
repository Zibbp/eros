package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/eros/internal/utils"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.Enum("status").GoType(utils.ReportStatus("")),
		field.String("s3_file").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("script", Script.Type).Ref("reports").Unique().Required(),
	}
}
