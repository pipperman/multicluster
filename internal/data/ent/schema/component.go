package schema

import "entgo.io/ent"

// Component holds the schema definition for the Component entity.
type Component struct {
	ent.Schema
}

// Fields of the Component.
func (Component) Fields() []ent.Field {
	return nil
}

// Edges of the Component.
func (Component) Edges() []ent.Edge {
	return nil
}
