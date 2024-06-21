package internal

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func ConfigZeroLog() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerolog.TimestampFieldName = strings.ToLower("timestamp")
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
}

type (
	Config struct {
		AppMode string `env:"APP_MODE,default=dev"`
		Verbose bool   `env:"VERBOSE,default=false"`
		DbURL   string `env:"DB_URL,default=postgresql://postgres:postgres@127.0.0.1:5432/postgres?search_path=eve&connect_timeout=3"`
	}
)

var (
	neverAgain   sync.Once
	globalConfig Config
)

func GlobalConfig() *Config {
	neverAgain.Do(func() {
		config := &globalConfig
		err := envdecode.Decode(config)
		if err != nil {
			log.Fatal().Err(err).Msg("load config failed")
		}

		// dev模式使用从环境变量中读取的值
		// 其他模式就使用预定义的服务器URL了
		if config.AppMode == "dev" {
		} else if config.AppMode == "local" {
		} else if config.AppMode == "stg" {
			config.DbURL = "postgresql://postgres:postgres@db.ccjehbluyfnt.ap-northeast-1.rds.amazonaws.com:5432/postgres?search_path=eve&connect_timeout=3"
		} else if config.AppMode == "prod" {
			config.DbURL = "postgresql://postgres:postgres@db.ccjehbluyfnt.ap-northeast-1.rds.amazonaws.com:5432/postgres?search_path=eve&connect_timeout=3"
		} else {
			log.Fatal().Msg("unknown app mode")
		}

		log.Info().
			Str("app mode", config.AppMode).
			Msg("config load ok")
	})
	return &globalConfig
}
