package pkl

import (
	"context"
	"sync"

	"freefrom.space/nobot/appConfig"
)

var (
	conf *appconfig.AppConfig
	once sync.Once
)

func GetConf() *appconfig.AppConfig {
	once.Do(initConf)
	return conf
}

func initConf() {
	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/local/nobotConfig.pkl")
	if err != nil {
		panic(err)
	}
	conf = cfg
}
