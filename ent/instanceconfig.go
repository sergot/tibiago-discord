// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sergot/tibiago/ent/instance"
	"github.com/sergot/tibiago/ent/instanceconfig"
)

// InstanceConfig is the model entity for the InstanceConfig schema.
type InstanceConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstanceConfigQuery when eager-loading is set.
	Edges           InstanceConfigEdges `json:"edges"`
	instance_config *uuid.UUID
}

// InstanceConfigEdges holds the relations/edges for other nodes in the graph.
type InstanceConfigEdges struct {
	// Instance holds the value of the instance edge.
	Instance *Instance `json:"instance,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceConfigEdges) InstanceOrErr() (*Instance, error) {
	if e.loadedTypes[0] {
		if e.Instance == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: instance.Label}
		}
		return e.Instance, nil
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InstanceConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case instanceconfig.FieldKey, instanceconfig.FieldValue:
			values[i] = new(sql.NullString)
		case instanceconfig.FieldID:
			values[i] = new(uuid.UUID)
		case instanceconfig.ForeignKeys[0]: // instance_config
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type InstanceConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InstanceConfig fields.
func (ic *InstanceConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case instanceconfig.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ic.ID = *value
			}
		case instanceconfig.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				ic.Key = value.String
			}
		case instanceconfig.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ic.Value = value.String
			}
		case instanceconfig.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field instance_config", values[i])
			} else if value.Valid {
				ic.instance_config = new(uuid.UUID)
				*ic.instance_config = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryInstance queries the "instance" edge of the InstanceConfig entity.
func (ic *InstanceConfig) QueryInstance() *InstanceQuery {
	return (&InstanceConfigClient{config: ic.config}).QueryInstance(ic)
}

// Update returns a builder for updating this InstanceConfig.
// Note that you need to call InstanceConfig.Unwrap() before calling this method if this InstanceConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *InstanceConfig) Update() *InstanceConfigUpdateOne {
	return (&InstanceConfigClient{config: ic.config}).UpdateOne(ic)
}

// Unwrap unwraps the InstanceConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *InstanceConfig) Unwrap() *InstanceConfig {
	_tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: InstanceConfig is not a transactional entity")
	}
	ic.config.driver = _tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *InstanceConfig) String() string {
	var builder strings.Builder
	builder.WriteString("InstanceConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ic.ID))
	builder.WriteString("key=")
	builder.WriteString(ic.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(ic.Value)
	builder.WriteByte(')')
	return builder.String()
}

// InstanceConfigs is a parsable slice of InstanceConfig.
type InstanceConfigs []*InstanceConfig

func (ic InstanceConfigs) config(cfg config) {
	for _i := range ic {
		ic[_i].config = cfg
	}
}