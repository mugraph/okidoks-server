package models

import "github.com/google/uuid"

type Affiliation struct {
	ID      uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	Name    string    `json:"name"`
	Authors []Author  `gorm:"many2many:author_affiliation"`
}

func AddAffiliationsToDB(affiliations []Affiliation) {
	// Loop over Affiliations and create them if they don't exist
	for i := range affiliations {
		query := Affiliation{
			Name: affiliations[i].Name,
		}
		DB.Where(query).FirstOrCreate(&affiliations[i])
	}
}
