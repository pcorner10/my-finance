package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Mexico_City", host, user, password, dbname, port)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

}
