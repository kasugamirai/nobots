package postgres

import (
	"context"
	"fmt"
	"freefrom.space/nobot/pkl"

	"entgo.io/ent/dialect"

	"freefrom.space/nobot/core"
	_ "github.com/lib/pq"
)

// Client is a global variable to access the ent.Client
var Client *core.Client

func Init() {
	var err error
	// Get the configuration
	c := pkl.GetConf().NewConfig.GetPostgres()
	print(c.GetPort())

	// Create the Postgres connection string from the configuration
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		c.GetHost(),
		c.GetPort(),
		c.GetUser(),
		c.GetDbname(),
		c.GetPassword(),
		c.GetSslmode(),
		"eve",
	)

	Client, err = core.Open(dialect.Postgres, connStr)
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to postgres: %v", err))
	}
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}
}
