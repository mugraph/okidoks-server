// models/setup.go

package models

import (
	"github.com/mugraph/okidoks-server/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var log = logger.Log

func ConnectDatabase() {
	// Data Source Name
	dsn := "host=localhost user=postgres dbname=okidoks_db port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database!", err)
	}

	database.AutoMigrate(&Publisher{})
	database.AutoMigrate(&Resource{}, &ContributorRole{}, &Contributor{}, &License{})

	DB = database
}
