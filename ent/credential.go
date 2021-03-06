// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/target"
)

// Credential is the model entity for the Credential schema.
type Credential struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Principal holds the value of the "principal" field.
	Principal string `json:"principal,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"secret,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind credential.Kind `json:"kind,omitempty"`
	// Fails holds the value of the "fails" field.
	Fails int `json:"fails,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CredentialQuery when eager-loading is set.
	Edges              CredentialEdges `json:"edges"`
	target_credentials *int
}

// CredentialEdges holds the relations/edges for other nodes in the graph.
type CredentialEdges struct {
	// Target holds the value of the target edge.
	Target *Target
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TargetOrErr returns the Target value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CredentialEdges) TargetOrErr() (*Target, error) {
	if e.loadedTypes[0] {
		if e.Target == nil {
			// The edge target was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: target.Label}
		}
		return e.Target, nil
	}
	return nil, &NotLoadedError{edge: "target"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Credential) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // principal
		&sql.NullString{}, // secret
		&sql.NullString{}, // kind
		&sql.NullInt64{},  // fails
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Credential) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // target_credentials
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Credential fields.
func (c *Credential) assignValues(values ...interface{}) error {
	if m, n := len(values), len(credential.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field principal", values[0])
	} else if value.Valid {
		c.Principal = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field secret", values[1])
	} else if value.Valid {
		c.Secret = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field kind", values[2])
	} else if value.Valid {
		c.Kind = credential.Kind(value.String)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field fails", values[3])
	} else if value.Valid {
		c.Fails = int(value.Int64)
	}
	values = values[4:]
	if len(values) == len(credential.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field target_credentials", value)
		} else if value.Valid {
			c.target_credentials = new(int)
			*c.target_credentials = int(value.Int64)
		}
	}
	return nil
}

// QueryTarget queries the target edge of the Credential.
func (c *Credential) QueryTarget() *TargetQuery {
	return (&CredentialClient{c.config}).QueryTarget(c)
}

// Update returns a builder for updating this Credential.
// Note that, you need to call Credential.Unwrap() before calling this method, if this Credential
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Credential) Update() *CredentialUpdateOne {
	return (&CredentialClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Credential) Unwrap() *Credential {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Credential is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Credential) String() string {
	var builder strings.Builder
	builder.WriteString("Credential(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", principal=")
	builder.WriteString(c.Principal)
	builder.WriteString(", secret=")
	builder.WriteString(c.Secret)
	builder.WriteString(", kind=")
	builder.WriteString(fmt.Sprintf("%v", c.Kind))
	builder.WriteString(", fails=")
	builder.WriteString(fmt.Sprintf("%v", c.Fails))
	builder.WriteByte(')')
	return builder.String()
}

// Credentials is a parsable slice of Credential.
type Credentials []*Credential

func (c Credentials) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
