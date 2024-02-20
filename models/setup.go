// models/setup.go

package models

import (
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models/commonmeta"
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
		log.Error("failed to connect to database", err)
	}

	database.AutoMigrate(
		&commonmeta.Resource{},
		&commonmeta.ContributorRole{},
		&commonmeta.Contributor{},
		&commonmeta.License{},
		&commonmeta.Publisher{},
		&commonmeta.Date{},
		&commonmeta.Title{},
	)

	DB = database
}
