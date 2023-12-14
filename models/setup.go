// models/setup.go

package models

import (
	"log"

	"gorm.io/datatypes"
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

	database.AutoMigrate(&Publication{})

	database.Create(&Publication{
		Name:       "publication-1",
		Attributes: datatypes.JSON([]byte(`{"title": "The Surprising Power of Liberating Structures: Simple Rules to Unleash A Culture of Innovation", "creator": ["Henri Lipmanovicz", "Keith MacCandless"], "subtitle": "-", "publisher": "Liberating Structures Press", "abstract": "Smart leaders know that they would greatly increase productivity and innovation if only they could get everyone fully engaged. So do professors, facilitators and all changemakers. The challenge is how. Liberating Structures are novel, practical and no-nonsense methods to help you accomplish this goal with groups of any size. Prepare to be surprised by how simple and easy they are for anyone to use. This book shows you how with detailed descriptions for putting them into practice plus tips on how to get started and traps to avoid. It takes the design and facilitation methods experts use and puts them within reach of anyone in any organization or initiative, from the frontline to the C-suite."}`)),
	})

	DB = database
}
