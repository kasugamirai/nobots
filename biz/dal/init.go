package dal

import (
	"freefrom.space/nobot/biz/dal/postgres"
	"freefrom.space/nobot/biz/dal/sqlite"
)

func Init() {
	sqlite.Init()
	postgres.Init()
}
