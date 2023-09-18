// models/setup.go

package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=okidoks_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[setup.go] Failed to connect to database!", err)
	}
	// database.AutoMigrate(&Tour{}, &Chapter{}, &Properties{})

	DB = database
}
