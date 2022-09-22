// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/karashiiro/velcro/ent/message"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// SourceAddress holds the value of the "source_address" field.
	SourceAddress string `json:"source_address,omitempty"`
	// SourcePort holds the value of the "source_port" field.
	SourcePort int `json:"source_port,omitempty"`
	// DestinationAddress holds the value of the "destination_address" field.
	DestinationAddress string `json:"destination_address,omitempty"`
	// DestinationPort holds the value of the "destination_port" field.
	DestinationPort int `json:"destination_port,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldData:
			values[i] = new([]byte)
		case message.FieldID, message.FieldVersion, message.FieldSourcePort, message.FieldDestinationPort:
			values[i] = new(sql.NullInt64)
		case message.FieldSourceAddress, message.FieldDestinationAddress:
			values[i] = new(sql.NullString)
		case message.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Message", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				m.Timestamp = value.Time
			}
		case message.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				m.Version = int(value.Int64)
			}
		case message.FieldSourceAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_address", values[i])
			} else if value.Valid {
				m.SourceAddress = value.String
			}
		case message.FieldSourcePort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_port", values[i])
			} else if value.Valid {
				m.SourcePort = int(value.Int64)
			}
		case message.FieldDestinationAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination_address", values[i])
			} else if value.Valid {
				m.DestinationAddress = value.String
			}
		case message.FieldDestinationPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field destination_port", values[i])
			} else if value.Valid {
				m.DestinationPort = int(value.Int64)
			}
		case message.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				m.Data = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return (&MessageClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("timestamp=")
	builder.WriteString(m.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", m.Version))
	builder.WriteString(", ")
	builder.WriteString("source_address=")
	builder.WriteString(m.SourceAddress)
	builder.WriteString(", ")
	builder.WriteString("source_port=")
	builder.WriteString(fmt.Sprintf("%v", m.SourcePort))
	builder.WriteString(", ")
	builder.WriteString("destination_address=")
	builder.WriteString(m.DestinationAddress)
	builder.WriteString(", ")
	builder.WriteString("destination_port=")
	builder.WriteString(fmt.Sprintf("%v", m.DestinationPort))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", m.Data))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
