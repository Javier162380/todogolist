package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("TaskName").NotEmpty(),
		field.String("TaskLabel").Default("unknown"),
		field.Time("TaskDate").Default(time.Now),
		field.Bool("RecurrentTask").Default(false),
		field.Int("TaskTimeInvested").Nillable().Optional(),
		field.String("Periodicity").Nillable().Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
