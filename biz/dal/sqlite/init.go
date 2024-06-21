package sqlite

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"freefrom.space/nobot/pkl"

	"freefrom.space/nobot/core"
	_ "github.com/mattn/go-sqlite3"
)

// Client is a global variable to access the ent.Client
var Client *core.Client

func Init() {
	var err error
	conf := pkl.GetConf()
	// Get the configuration
	fmt.Println(conf.NewConfig.GetSqlite().GetDsn())
	Client, err = core.Open(dialect.SQLite, conf.NewConfig.GetSqlite().GetDsn())
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}
}
