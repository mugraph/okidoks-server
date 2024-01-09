package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID          uuid.UUID     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt   time.Time     `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt   time.Time     `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	Given       string        `json:"given"`
	Family      string        `json:"family"`
	Sequence    string        `json:"sequence"`
	ORCID       string        `json:"ORCID" gorm:"unique_index"`
	Affiliation []Affiliation `gorm:"many2many:author_affiliation"`
	Resources   []Resource    `gorm:"many2many:resource_author"`
}

func AddAuthorsToDB(authors []Author) {
	fmt.Printf("%v\n", authors)
	// Loop over Authors and create them if they don't exist
	for i := range authors {
		// Add Affiliations
		// AddAffiliationsToDB(authors[i].Affiliation)

		DB.Create(&authors[i])
	}
}
