// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Sqlite interface {
	GetDsn() string
}

var _ Sqlite = (*SqliteImpl)(nil)

// The configuration for the SQLite database.
type SqliteImpl struct {
	// The path to the SQLite database file.
	Dsn string `pkl:"dsn"`
}

// The path to the SQLite database file.
func (rcv *SqliteImpl) GetDsn() string {
	return rcv.Dsn
}
