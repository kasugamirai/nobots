// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Postgres interface {
	GetHost() string

	GetPort() int16

	GetUser() string

	GetPassword() string

	GetDbname() string

	GetSslmode() string

	GetTimezone() string
}

var _ Postgres = (*PostgresImpl)(nil)

type PostgresImpl struct {
	// The connection string for the Postgres database.
	Host string `pkl:"host"`

	Port int16 `pkl:"port"`

	User string `pkl:"user"`

	Password string `pkl:"password"`

	Dbname string `pkl:"dbname"`

	Sslmode string `pkl:"sslmode"`

	Timezone string `pkl:"timezone"`
}

// The connection string for the Postgres database.
func (rcv *PostgresImpl) GetHost() string {
	return rcv.Host
}

func (rcv *PostgresImpl) GetPort() int16 {
	return rcv.Port
}

func (rcv *PostgresImpl) GetUser() string {
	return rcv.User
}

func (rcv *PostgresImpl) GetPassword() string {
	return rcv.Password
}

func (rcv *PostgresImpl) GetDbname() string {
	return rcv.Dbname
}

func (rcv *PostgresImpl) GetSslmode() string {
	return rcv.Sslmode
}

func (rcv *PostgresImpl) GetTimezone() string {
	return rcv.Timezone
}
