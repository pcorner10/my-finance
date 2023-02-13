package main

import (
	"github.com/pcorner10/my-finance/config"
	"github.com/pcorner10/my-finance/db"
	"github.com/spf13/viper"
)

func main() {

	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		seed.Migrate(dbHandler)
	}
}
