// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sergot/tibiago/ent/instance"
)

// Instance is the model entity for the Instance schema.
type Instance struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SessionID holds the value of the "session_id" field.
	SessionID string `json:"session_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstanceQuery when eager-loading is set.
	Edges InstanceEdges `json:"edges"`
}

// InstanceEdges holds the relations/edges for other nodes in the graph.
type InstanceEdges struct {
	// Config holds the value of the config edge.
	Config []*InstanceConfig `json:"config,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ConfigOrErr returns the Config value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceEdges) ConfigOrErr() ([]*InstanceConfig, error) {
	if e.loadedTypes[0] {
		return e.Config, nil
	}
	return nil, &NotLoadedError{edge: "config"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Instance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case instance.FieldSessionID:
			values[i] = new(sql.NullString)
		case instance.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Instance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Instance fields.
func (i *Instance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case instance.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case instance.FieldSessionID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field session_id", values[j])
			} else if value.Valid {
				i.SessionID = value.String
			}
		}
	}
	return nil
}

// QueryConfig queries the "config" edge of the Instance entity.
func (i *Instance) QueryConfig() *InstanceConfigQuery {
	return (&InstanceClient{config: i.config}).QueryConfig(i)
}

// Update returns a builder for updating this Instance.
// Note that you need to call Instance.Unwrap() before calling this method if this Instance
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Instance) Update() *InstanceUpdateOne {
	return (&InstanceClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Instance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Instance) Unwrap() *Instance {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Instance is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Instance) String() string {
	var builder strings.Builder
	builder.WriteString("Instance(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("session_id=")
	builder.WriteString(i.SessionID)
	builder.WriteByte(')')
	return builder.String()
}

// Instances is a parsable slice of Instance.
type Instances []*Instance

func (i Instances) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
