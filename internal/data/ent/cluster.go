// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"multicluster/internal/data/ent/cluster"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Cluster is the model entity for the Cluster schema.
type Cluster struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cluster) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cluster.FieldID:
			values[i] = new(sql.NullInt64)
		case cluster.FieldName:
			values[i] = new(sql.NullString)
		case cluster.FieldCreatedAt, cluster.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cluster", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cluster fields.
func (c *Cluster) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cluster.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case cluster.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case cluster.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cluster.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Cluster.
// Note that you need to call Cluster.Unwrap() before calling this method if this Cluster
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cluster) Update() *ClusterUpdateOne {
	return (&ClusterClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cluster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cluster) Unwrap() *Cluster {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cluster is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cluster) String() string {
	var builder strings.Builder
	builder.WriteString("Cluster(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Clusters is a parsable slice of Cluster.
type Clusters []*Cluster

func (c Clusters) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
