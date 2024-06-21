// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

import "github.com/apple/pkl-go/pkl"

func init() {
	pkl.RegisterMapping("nobot.AppConfig", AppConfig{})
	pkl.RegisterMapping("nobot.AppConfig#Config", ConfigImpl{})
	pkl.RegisterMapping("nobot.AppConfig#Postgres", PostgresImpl{})
	pkl.RegisterMapping("nobot.AppConfig#Relays", RelaysImpl{})
	pkl.RegisterMapping("nobot.AppConfig#Sqlite", SqliteImpl{})
	pkl.RegisterMapping("nobot.AppConfig#Twitter", TwitterImpl{})
	pkl.RegisterMapping("nobot.AppConfig#Report", ReportImpl{})
	pkl.RegisterMapping("nobot.AppConfig#nostr", NostrImpl{})
}
