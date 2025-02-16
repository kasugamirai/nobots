// Code generated by ent, DO NOT EDIT.

package core

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"freefrom.space/nobot/core/twitternotesfetchhistory"
)

// TwitterNotesFetchHistory is the model entity for the TwitterNotesFetchHistory schema.
type TwitterNotesFetchHistory struct {
	config `json:"-"`
	// ID of the ent.
	// auto inc id
	ID int64 `json:"id,omitempty"`
	// twitter user name
	UserID string `json:"user_id,omitempty"`
	// twitter user name. 冗余字段,只是为了方便没有工具时查看, 减少join 操作.
	Username string `json:"username,omitempty"`
	// twitter user name. 冗余字段,只是为了方便没有工具时查看, 减少join 操作.
	Name string `json:"name,omitempty"`
	// fetched time
	FetchedAt time.Time `json:"fetched_at,omitempty"`
	// fetched count
	FetchCount int `json:"fetch_count,omitempty"`
	// new count
	NewCount     int `json:"new_count,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TwitterNotesFetchHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case twitternotesfetchhistory.FieldID, twitternotesfetchhistory.FieldFetchCount, twitternotesfetchhistory.FieldNewCount:
			values[i] = new(sql.NullInt64)
		case twitternotesfetchhistory.FieldUserID, twitternotesfetchhistory.FieldUsername, twitternotesfetchhistory.FieldName:
			values[i] = new(sql.NullString)
		case twitternotesfetchhistory.FieldFetchedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TwitterNotesFetchHistory fields.
func (tnfh *TwitterNotesFetchHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case twitternotesfetchhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tnfh.ID = int64(value.Int64)
		case twitternotesfetchhistory.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tnfh.UserID = value.String
			}
		case twitternotesfetchhistory.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				tnfh.Username = value.String
			}
		case twitternotesfetchhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tnfh.Name = value.String
			}
		case twitternotesfetchhistory.FieldFetchedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field fetched_at", values[i])
			} else if value.Valid {
				tnfh.FetchedAt = value.Time
			}
		case twitternotesfetchhistory.FieldFetchCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fetch_count", values[i])
			} else if value.Valid {
				tnfh.FetchCount = int(value.Int64)
			}
		case twitternotesfetchhistory.FieldNewCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field new_count", values[i])
			} else if value.Valid {
				tnfh.NewCount = int(value.Int64)
			}
		default:
			tnfh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TwitterNotesFetchHistory.
// This includes values selected through modifiers, order, etc.
func (tnfh *TwitterNotesFetchHistory) Value(name string) (ent.Value, error) {
	return tnfh.selectValues.Get(name)
}

// Update returns a builder for updating this TwitterNotesFetchHistory.
// Note that you need to call TwitterNotesFetchHistory.Unwrap() before calling this method if this TwitterNotesFetchHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (tnfh *TwitterNotesFetchHistory) Update() *TwitterNotesFetchHistoryUpdateOne {
	return NewTwitterNotesFetchHistoryClient(tnfh.config).UpdateOne(tnfh)
}

// Unwrap unwraps the TwitterNotesFetchHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tnfh *TwitterNotesFetchHistory) Unwrap() *TwitterNotesFetchHistory {
	_tx, ok := tnfh.config.driver.(*txDriver)
	if !ok {
		panic("core: TwitterNotesFetchHistory is not a transactional entity")
	}
	tnfh.config.driver = _tx.drv
	return tnfh
}

// String implements the fmt.Stringer.
func (tnfh *TwitterNotesFetchHistory) String() string {
	var builder strings.Builder
	builder.WriteString("TwitterNotesFetchHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tnfh.ID))
	builder.WriteString("user_id=")
	builder.WriteString(tnfh.UserID)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(tnfh.Username)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tnfh.Name)
	builder.WriteString(", ")
	builder.WriteString("fetched_at=")
	builder.WriteString(tnfh.FetchedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fetch_count=")
	builder.WriteString(fmt.Sprintf("%v", tnfh.FetchCount))
	builder.WriteString(", ")
	builder.WriteString("new_count=")
	builder.WriteString(fmt.Sprintf("%v", tnfh.NewCount))
	builder.WriteByte(')')
	return builder.String()
}

// TwitterNotesFetchHistories is a parsable slice of TwitterNotesFetchHistory.
type TwitterNotesFetchHistories []*TwitterNotesFetchHistory
