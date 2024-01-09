package models

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mugraph/okidoks-server/utils"
)

type Resource struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time      `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	DOIAgency string         `json:"DOI_agency"`
	Type      string         `json:"type"`
	DOI       string         `json:"DOI"`
	Prefix    string         `json:"prefix"`
	Title     pq.StringArray `json:"title" gorm:"type:text[]"`
	Authors   []Author       `json:"author" gorm:"many2many:resource_author"`
}

func ResourceFromDOI(doi string) (resource Resource, err error) {
	// Get registration agency from DOI
	ra, err := utils.GetDOIRA(doi)
	if err != nil {
		return Resource{}, fmt.Errorf("Could not get DOI registration agency: %w", err)
	}
	fmt.Printf("ra:  %v\n", ra)

	URL, err := utils.DOIAsURL(doi)
	if err != nil {
		return Resource{}, fmt.Errorf("Could not get URL from DOI: %w", err)
	}
	fmt.Printf("URL: %v\n\n", URL)

	escaped_doi := url.QueryEscape(doi)

	switch id := ra; id {
	case "Crossref":
		// GET Metadata from Crossref
		resource, err = GetDOIMetadataFromCrossref(escaped_doi)
		if err != nil {
			return Resource{}, fmt.Errorf("Could not get Resource from Crossref DOI: %v", err)
		}
	case "Datacite":
		// GET Metadata from DataCite
		resource, err = GetDOIMetadataFromDataCite(escaped_doi)
		if err != nil {
			return Resource{}, fmt.Errorf("Could not get Resource from DataCite DOI: %v", err)
		}
	}
	return resource, nil
}

func AddResourceToDB(resource Resource) {

	// Add Authors
	// AddAuthorsToDB(resource.Authors)

	// Create new Resource in DB, if DOI doesn't exists
	query := Resource{
		DOI: resource.DOI,
	}
	DB.Where(query).FirstOrCreate(&resource)
}
