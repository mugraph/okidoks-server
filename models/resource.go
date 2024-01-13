package models

import (
	"time"

	"github.com/google/uuid"
)

// JSON representation of the Commonmeta schema.
type Resource struct {
	ID                   string                `json:"id" gorm:"primaryKey"` // The unique identifier for the resource.
	UUID                 uuid.UUID             `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt            time.Time             `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt            time.Time             `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	AdditionalType       *string               `json:"additional_type,omitempty"`                            // The additional type of the resource.
	AlternateIdentifiers []AlternateIdentifier `json:"alternate_identifiers,omitempty"`                      // Alternate identifiers for the resource.
	ArchiveLocations     []ArchiveLocation     `json:"archive_locations,omitempty" gorm:"type:text[]"`       // The location where content is archived.
	Container            *Container            `json:"container,omitempty"`                                  // The container of the resource.
	Contributors         []*Contributor        `json:"contributors" gorm:"many2many:resource_contributors;"` // The contributors to the resource.
	Date                 Date                  `json:"date"`                                                 // The dates for the resource.
	Descriptions         []DescriptionElement  `json:"descriptions,omitempty"`                               // The descriptions of the resource.
	Files                []File                `json:"files,omitempty"`                                      // The downloadable files for the resource.
	Formats              []string              `json:"formats,omitempty" gorm:"type:text[]"`                 // The formats of the resource.
	FundingReferences    []FundingReference    `json:"funding_references,omitempty"`                         // The funding references for the resource.
	GeoLocations         []GeoLocation         `json:"geo_locations,omitempty"`                              // The geolocations for the resource.
	Language             *string               `json:"language,omitempty"`                                   // The language of the resource. Use one of the language codes from the IETF BCP 47 standard.
	// LicenseID            uuid.UUID
	// License              *License  `json:"license,omitempty"`                     // The license for the resource. Use one of the SPDX license identifiers.
	Provider           *Provider           `json:"provider,omitempty" gorm:"type:text[]"` // The provider of the resource. This can be a DOI registration agency or a repository.
	Publisher          Publisher           `json:"publisher"`                             // The publisher of the resource.
	References         []Reference         `json:"references,omitempty"`                  // The references of the resource.
	RelatedIdentifiers []RelatedIdentifier `json:"related_identifiers,omitempty"`         // Other resolvable persistent unique IDs related to the resource.
	SchemaVersion      *SchemaVersion      `json:"schema_version,omitempty"`              // The schema version of the resource.
	Sizes              []string            `json:"sizes,omitempty" gorm:"type:text[]"`    // The sizes of the resource.
	State              *State              `json:"state,omitempty"`                       // The state of the resource.
	Subjects           []Subject           `json:"subjects,omitempty"`                    // The subjects of the resource.
	Titles             []Title             `json:"titles"`                                // The titles of the resource.
	Type               ResourceType        `json:"type"`                                  // The type of the resource.
	URL                string              `json:"url"`                                   // The URL of the resource.
	Version            *string             `json:"version,omitempty"`                     // The version of the resource.
}

// The type of the resource.
type ResourceType string

const (
	Article                 ResourceType = "Article"
	Audiovisual             ResourceType = "Audiovisual"
	BookChapter             ResourceType = "BookChapter"
	Component               ResourceType = "Component"
	Dataset                 ResourceType = "Dataset"
	Dissertation            ResourceType = "Dissertation"
	Document                ResourceType = "Document"
	FluffyBook              ResourceType = "Book"
	FluffyBookSeries        ResourceType = "BookSeries"
	FluffyJournal           ResourceType = "Journal"
	FluffyProceedingsSeries ResourceType = "ProceedingsSeries"
	Grant                   ResourceType = "Grant"
	Instrument              ResourceType = "Instrument"
	JournalArticle          ResourceType = "JournalArticle"
	JournalIssue            ResourceType = "JournalIssue"
	JournalVolume           ResourceType = "JournalVolume"
	PeerReview              ResourceType = "PeerReview"
	PhysicalObject          ResourceType = "PhysicalObject"
	Proceedings             ResourceType = "Proceedings"
	ProceedingsArticle      ResourceType = "ProceedingsArticle"
	Report                  ResourceType = "Report"
	ReportComponent         ResourceType = "ReportComponent"
	ReportSeries            ResourceType = "ReportSeries"
	TypeOther               ResourceType = "Other"
	TypeSoftware            ResourceType = "Software"
)
