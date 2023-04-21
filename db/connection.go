package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=postgres dbname=gorm port=5432"

var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Db Connected")

	}
}
