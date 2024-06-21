package main

import (
	"fmt"

	"freefrom.space/nobot/pkl"
)

func main() {
	fmt.Printf("I'm running on host %v\n", pkl.GetConf().NewConfig.GetSqlite())
}
