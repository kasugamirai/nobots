// Code generated by ent, DO NOT EDIT.

package core

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"freefrom.space/nobot/core/nostrnotesfetchhistory"
)

// NostrNotesFetchHistory is the model entity for the NostrNotesFetchHistory schema.
type NostrNotesFetchHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// post time
	PublishAt time.Time `json:"publish_at,omitempty"`
	// public key
	PublicKeyHex string `json:"public_key_hex,omitempty"`
	// nostr context
	Context      string `json:"context,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NostrNotesFetchHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case nostrnotesfetchhistory.FieldID, nostrnotesfetchhistory.FieldPublicKeyHex, nostrnotesfetchhistory.FieldContext:
			values[i] = new(sql.NullString)
		case nostrnotesfetchhistory.FieldPublishAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NostrNotesFetchHistory fields.
func (nnfh *NostrNotesFetchHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case nostrnotesfetchhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				nnfh.ID = value.String
			}
		case nostrnotesfetchhistory.FieldPublishAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field publish_at", values[i])
			} else if value.Valid {
				nnfh.PublishAt = value.Time
			}
		case nostrnotesfetchhistory.FieldPublicKeyHex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_key_hex", values[i])
			} else if value.Valid {
				nnfh.PublicKeyHex = value.String
			}
		case nostrnotesfetchhistory.FieldContext:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field context", values[i])
			} else if value.Valid {
				nnfh.Context = value.String
			}
		default:
			nnfh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NostrNotesFetchHistory.
// This includes values selected through modifiers, order, etc.
func (nnfh *NostrNotesFetchHistory) Value(name string) (ent.Value, error) {
	return nnfh.selectValues.Get(name)
}

// Update returns a builder for updating this NostrNotesFetchHistory.
// Note that you need to call NostrNotesFetchHistory.Unwrap() before calling this method if this NostrNotesFetchHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (nnfh *NostrNotesFetchHistory) Update() *NostrNotesFetchHistoryUpdateOne {
	return NewNostrNotesFetchHistoryClient(nnfh.config).UpdateOne(nnfh)
}

// Unwrap unwraps the NostrNotesFetchHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nnfh *NostrNotesFetchHistory) Unwrap() *NostrNotesFetchHistory {
	_tx, ok := nnfh.config.driver.(*txDriver)
	if !ok {
		panic("core: NostrNotesFetchHistory is not a transactional entity")
	}
	nnfh.config.driver = _tx.drv
	return nnfh
}

// String implements the fmt.Stringer.
func (nnfh *NostrNotesFetchHistory) String() string {
	var builder strings.Builder
	builder.WriteString("NostrNotesFetchHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nnfh.ID))
	builder.WriteString("publish_at=")
	builder.WriteString(nnfh.PublishAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("public_key_hex=")
	builder.WriteString(nnfh.PublicKeyHex)
	builder.WriteString(", ")
	builder.WriteString("context=")
	builder.WriteString(nnfh.Context)
	builder.WriteByte(')')
	return builder.String()
}

// NostrNotesFetchHistories is a parsable slice of NostrNotesFetchHistory.
type NostrNotesFetchHistories []*NostrNotesFetchHistory
