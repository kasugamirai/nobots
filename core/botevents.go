// Code generated by ent, DO NOT EDIT.

package core

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"freefrom.space/nobot/core/botevents"
)

// BotEvents is the model entity for the BotEvents schema.
type BotEvents struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EventID holds the value of the "event_id" field.
	EventID string `json:"event_id,omitempty"`
	// InsertedAt holds the value of the "inserted_at" field.
	InsertedAt   time.Time `json:"inserted_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BotEvents) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case botevents.FieldID:
			values[i] = new(sql.NullInt64)
		case botevents.FieldEventID:
			values[i] = new(sql.NullString)
		case botevents.FieldInsertedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BotEvents fields.
func (be *BotEvents) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case botevents.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			be.ID = int(value.Int64)
		case botevents.FieldEventID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value.Valid {
				be.EventID = value.String
			}
		case botevents.FieldInsertedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field inserted_at", values[i])
			} else if value.Valid {
				be.InsertedAt = value.Time
			}
		default:
			be.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BotEvents.
// This includes values selected through modifiers, order, etc.
func (be *BotEvents) Value(name string) (ent.Value, error) {
	return be.selectValues.Get(name)
}

// Update returns a builder for updating this BotEvents.
// Note that you need to call BotEvents.Unwrap() before calling this method if this BotEvents
// was returned from a transaction, and the transaction was committed or rolled back.
func (be *BotEvents) Update() *BotEventsUpdateOne {
	return NewBotEventsClient(be.config).UpdateOne(be)
}

// Unwrap unwraps the BotEvents entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (be *BotEvents) Unwrap() *BotEvents {
	_tx, ok := be.config.driver.(*txDriver)
	if !ok {
		panic("core: BotEvents is not a transactional entity")
	}
	be.config.driver = _tx.drv
	return be
}

// String implements the fmt.Stringer.
func (be *BotEvents) String() string {
	var builder strings.Builder
	builder.WriteString("BotEvents(")
	builder.WriteString(fmt.Sprintf("id=%v, ", be.ID))
	builder.WriteString("event_id=")
	builder.WriteString(be.EventID)
	builder.WriteString(", ")
	builder.WriteString("inserted_at=")
	builder.WriteString(be.InsertedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BotEventsSlice is a parsable slice of BotEvents.
type BotEventsSlice []*BotEvents
