package db

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbHandler *gorm.DB

func Init() *gorm.DB {
	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.name")
	port := viper.GetString("database.port")
	password := viper.GetString("database.password")

	dsn := fmt.Sprintf(`
		host = %s
		user = %s
		password = %s
		dbname = %s
		port = %s
		sslmode = disable 
		TimeZone=America/Mexico_City
		`, host, user, dbname, password, port)

	dbHandler, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sql, err := dbHandler.DB()
	if err != nil {
		log.Fatal(err)
	}

	sql.SetMaxIdleConns(99)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxIdleTime(time.Second * 5)
	sql.SetConnMaxLifetime(time.Minute)

	return dbHandler
}

func GetDB() *gorm.DB {
	return dbHandler
}
