package twt

import (
	"context"

	"freefrom.space/nobot/biz/dal/sqlite"
	"freefrom.space/nobot/core"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var InitCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "update db schema",
	Flags:   []cli.Flag{},
	Action: func(cliCtx *cli.Context) error {
		dbClient := sqlite.Client
		dbClient = dbClient.Debug()

		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)

		if err := dbClient.Schema.Create(context.Background()); err != nil {
			log.Err(err).Msg("migrate db failed")
		} else {
			log.Info().Msg("migrate db ok")
		}
		return nil
	},
}
