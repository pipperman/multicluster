package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Cluster holds the schema definition for the Cluster entity.
type Cluster struct {
	ent.Schema
}

// Fields of the Cluster.
func (Cluster) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("cluster_id"),
		field.String("name"),
		field.String("cluster_type").
			Default("managed"),
		field.String("cluster_spec"),
		field.String("version"),
		field.String("profile").
			Default("default"),
		field.String("region_id"),
		field.String("vpc_id"),
		field.String("zone_id"),
		field.Bool("enable_deletion_protection").
			Default(true),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Cluster.
func (Cluster) Edges() []ent.Edge {
	return []ent.Edge{}
}
