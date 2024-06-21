package main

import (
	"flag"
	"os"
	"time"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	BaseEntPackageName = "freefrom.space/nobot"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	})

	flag.Bool("h", false, "show help")
	flag.Parse()
	svc := flag.Arg(0)

	if svc == "" {
		// flag.CommandLine.Usage()
		// return
		svc = "core"
	}

	log.Info().Str("service", svc).Msg("begin generate code")
	switch svc {
	case "core":
		entgen("core")
	case "eve":
		entgen("eve")
	case "chat":
		entgen("chat")
	default:
		log.Warn().Str("service", svc).Msg("unknown service")
	}
}

func entgen(target string) {
	var schemaPath, genSrcDir, pkgName string
	schemaPath = "./" + target + "/schema"
	genSrcDir = "./" + target
	pkgName = BaseEntPackageName + "/" + target

	// gqlext, err := entgql.NewExtension(
	// 	entgql.WithSchemaGenerator(),
	// 	entgql.WithSchemaPath("graph/ent.graphqls"),
	// )
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("new ent gql extension failed")
	// 	return
	// }

	t1 := time.Now()
	err := entc.Generate(schemaPath,
		&gen.Config{
			Target:  genSrcDir,
			Package: pkgName,
			Features: []gen.Feature{
				gen.FeatureUpsert,
				gen.FeatureLock,
				gen.FeatureModifier,
				gen.FeaturePrivacy,
				gen.FeatureExecQuery,
				gen.FeatureSchemaConfig,
			},
		},
		// entc.Extensions(gqlext),
	)
	t2 := time.Now()
	if err != nil {
		log.Fatal().Err(err).Msg("generate ent code failed.")
		return
	}
	log.Info().TimeDiff("time usage", t2, t1).Msg("generate ent code ok.")
}
