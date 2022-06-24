// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/testifysec/archivist/ent/dsse"
	"github.com/testifysec/archivist/ent/signature"
)

// Signature is the model entity for the Signature schema.
type Signature struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// KeyID holds the value of the "key_id" field.
	KeyID string `json:"key_id,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature string `json:"signature,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SignatureQuery when eager-loading is set.
	Edges           SignatureEdges `json:"edges"`
	dsse_signatures *int
}

// SignatureEdges holds the relations/edges for other nodes in the graph.
type SignatureEdges struct {
	// Dsse holds the value of the dsse edge.
	Dsse *Dsse `json:"dsse,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DsseOrErr returns the Dsse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SignatureEdges) DsseOrErr() (*Dsse, error) {
	if e.loadedTypes[0] {
		if e.Dsse == nil {
			// The edge dsse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dsse.Label}
		}
		return e.Dsse, nil
	}
	return nil, &NotLoadedError{edge: "dsse"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Signature) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case signature.FieldID:
			values[i] = new(sql.NullInt64)
		case signature.FieldKeyID, signature.FieldSignature:
			values[i] = new(sql.NullString)
		case signature.ForeignKeys[0]: // dsse_signatures
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Signature", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Signature fields.
func (s *Signature) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case signature.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case signature.FieldKeyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key_id", values[i])
			} else if value.Valid {
				s.KeyID = value.String
			}
		case signature.FieldSignature:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value.Valid {
				s.Signature = value.String
			}
		case signature.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field dsse_signatures", value)
			} else if value.Valid {
				s.dsse_signatures = new(int)
				*s.dsse_signatures = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDsse queries the "dsse" edge of the Signature entity.
func (s *Signature) QueryDsse() *DsseQuery {
	return (&SignatureClient{config: s.config}).QueryDsse(s)
}

// Update returns a builder for updating this Signature.
// Note that you need to call Signature.Unwrap() before calling this method if this Signature
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Signature) Update() *SignatureUpdateOne {
	return (&SignatureClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Signature entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Signature) Unwrap() *Signature {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Signature is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Signature) String() string {
	var builder strings.Builder
	builder.WriteString("Signature(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", key_id=")
	builder.WriteString(s.KeyID)
	builder.WriteString(", signature=")
	builder.WriteString(s.Signature)
	builder.WriteByte(')')
	return builder.String()
}

// Signatures is a parsable slice of Signature.
type Signatures []*Signature

func (s Signatures) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}