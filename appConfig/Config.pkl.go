// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Config interface {
	GetPostgres() Postgres

	GetRelays() Relays

	GetSqlite() Sqlite

	GetTwitter() Twitter

	GetReport() Report

	GetNostr() Nostr
}

var _ Config = (*ConfigImpl)(nil)

type ConfigImpl struct {
	Postgres Postgres `pkl:"postgres"`

	Relays Relays `pkl:"relays"`

	Sqlite Sqlite `pkl:"sqlite"`

	Twitter Twitter `pkl:"twitter"`

	Report Report `pkl:"report"`

	Nostr Nostr `pkl:"nostr"`
}

func (rcv *ConfigImpl) GetPostgres() Postgres {
	return rcv.Postgres
}

func (rcv *ConfigImpl) GetRelays() Relays {
	return rcv.Relays
}

func (rcv *ConfigImpl) GetSqlite() Sqlite {
	return rcv.Sqlite
}

func (rcv *ConfigImpl) GetTwitter() Twitter {
	return rcv.Twitter
}

func (rcv *ConfigImpl) GetReport() Report {
	return rcv.Report
}

func (rcv *ConfigImpl) GetNostr() Nostr {
	return rcv.Nostr
}
