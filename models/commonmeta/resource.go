// Version: v0.11
// URL:     https://commonmeta.org/commonmeta_v0.11.json
package commonmeta

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// DB representation of a Commonmeta Resource.
type Resource struct {
	UUID      uuid.UUID `json:"uuid"                            gorm:"primaryKey;type:uuid;default:gen_random_uuid()"              example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at"                                                                                         example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at"                                                                                         example:"2024-01-05T22:00:00.000000+01:00"`
	// required attributes (sorted by schema order)
	ID   string        `json:"id"                              gorm:"uniqueIndex"` // The unique identifier for the resource.
	Type *ResourceType `json:"type"`                                               // The type of the resource.
	// optional attributes (sorted alphabetically)
	AdditionalType       *ResourceType          `json:"additional_type,omitempty"`                                                                          // The additional type of the resource.
	AlternateIdentifiers []*AlternateIdentifier `json:"alternate_identifiers,omitempty"`                                                                    // Alternate identifiers for the resource.
	ArchiveLocations     []*ArchiveLocation     `json:"archive_locations,omitempty"     gorm:"type:text[]"`                                                 // The location where content is archived.
	Container            *Container             `json:"container,omitempty"`                                                                                // The container of the resource.
	Contributors         []*Contributor         `json:"contributors,omitempty"          gorm:"many2many:resource2contributors;constraint:OnDelete:CASCADE"` // The contributors to the resource.
	Date                 *Date                  `json:"date,omitempty"                  gorm:"foreignKey:UUID"`                                             // The dates for the resource.
	Descriptions         []*Description         `json:"descriptions,omitempty"`                                                                             // The descriptions of the resource.
	Files                []*File                `json:"files,omitempty"`                                                                                    // The downloadable files for the resource.
	Formats              []*string              `json:"formats,omitempty"               gorm:"type:text[]"`                                                 // The formats of the resource.
	FundingReferences    []*FundingReference    `json:"funding_references,omitempty"`                                                                       // The funding references for the resource.
	GeoLocations         []*GeoLocation         `json:"geo_locations,omitempty"`                                                                            // The geolocations for the resource.
	Language             string                 `json:"language,omitempty"`                                                                                 // The language of the resource. Use one of the language codes from the IETF BCP 47 standard.
	License              *License               `json:"license,omitempty"               gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`                // The license for the resource. Use one of the SPDX license identifiers.
	LicenseID            *uuid.UUID             `json:"-"                               gorm:"type:uuid"`                                                   // foreignKey
	Provider             *Provider              `json:"provider,omitempty"              gorm:"type:text[]"`                                                 // The provider of the resource. This can be a DOI registration agency or a repository.
	Publisher            *Publisher             `json:"publisher,omitempty"             gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`                // The publisher of the resource.
	PublisherID          *uuid.UUID             `json:"-"                               gorm:"type:uuid"`                                                   // foreignKey
	References           []*Reference           `json:"references,omitempty"`                                                                               // The references of the resource.
	RelatedIdentifiers   []*RelatedIdentifier   `json:"related_identifiers,omitempty"`                                                                      // Other resolvable persistent unique IDs related to the resource.
	SchemaVersion        *SchemaVersion         `json:"schema_version,omitempty"`                                                                           // The schema version of the resource.
	Sizes                []*string              `json:"sizes,omitempty"                 gorm:"type:text[]"`                                                 // The sizes of the resource.
	State                *State                 `json:"state,omitempty"`                                                                                    // The state of the resource.
	Subjects             []*Subject             `json:"subjects,omitempty"`                                                                                 // The subjects of the resource.
	Titles               []*Title               `json:"titles,omitempty"                gorm:"foreignKey:UUID"`                                             // The titles of the resource.
	URL                  string                 `json:"url,omitempty"`                                                                                      // The URL of the resource.
	Version              string                 `json:"version,omitempty"`                                                                                  // The version of the resource.
}

// The type of the resource.
type ResourceType string

// JSON representation of a Commonmeta Resource.
type ResourceJSON struct {
	// required attributes (sorted by schema order)
	ID   string        `json:"id"`
	Type *ResourceType `json:"type"`
	// optional attributes (sorted alphabetically)
	AdditionalType       *ResourceType          `json:"additional_type,omitempty"`
	AlternateIdentifiers []*AlternateIdentifier `json:"alternate_identifiers,omitempty"`
	ArchiveLocations     []*ArchiveLocation     `json:"archive_locations,omitempty"`
	Container            *Container             `json:"container,omitempty"`
	Contributors         []*Contributor         `json:"contributors,omitempty"`
	Date                 *Date                  `json:"date,omitempty"`
	Descriptions         []*Description         `json:"descriptions,omitempty"`
	Files                []*File                `json:"files,omitempty"`
	Formats              []*string              `json:"formats,omitempty"`
	FundingReferences    []*FundingReference    `json:"funding_references,omitempty"`
	GeoLocations         []*GeoLocation         `json:"geo_locations,omitempty"`
	Language             string                 `json:"language,omitempty"`
	License              *License               `json:"license,omitempty"`
	Provider             *Provider              `json:"provider,omitempty"`
	Publisher            *Publisher             `json:"publisher,omitempty"`
	References           []*Reference           `json:"references,omitempty"`
	RelatedIdentifiers   []*RelatedIdentifier   `json:"related_identifiers,omitempty"`
	SchemaVersion        *SchemaVersion         `json:"schema_version,omitempty"`
	Sizes                []*string              `json:"sizes,omitempty"`
	State                *State                 `json:"state,omitempty"`
	Subjects             []*Subject             `json:"subjects,omitempty"`
	Titles               []*Title               `json:"titles,omitempty"`
	URL                  string                 `json:"url,omitempty"`
	Version              string                 `json:"version,omitempty"`
}

// NilOrNotEmpty takes a pointer to a string type and returns nil if the value
// pointed to is empty. In this case the pointer itself is not nil.
// Otherwise it returns a pointer to a value.
func NilOrPtrToString[T ~string](s *T) *T {
	if s != nil && *s == "" {
		return nil
	}
	return s
}

func (r *Resource) ToJSONModel() ResourceJSON {
	return ResourceJSON{
		// required attributes (sorted by schema order)
		ID:   r.ID,
		Type: r.Type,
		// optional attributes (sorted alphabetically)
		AdditionalType:       NilOrPtrToString(r.AdditionalType),
		AlternateIdentifiers: r.AlternateIdentifiers,
		ArchiveLocations:     r.ArchiveLocations,
		Container:            r.Container,
		Contributors:         r.Contributors,
		Date:                 r.Date,
		Descriptions:         r.Descriptions,
		Files:                r.Files,
		Formats:              r.Formats,
		FundingReferences:    r.FundingReferences,
		GeoLocations:         r.GeoLocations,
		Language:             r.Language,
		License:              r.License,
		Provider:             r.Provider,
		Publisher:            r.Publisher,
		References:           r.References,
		RelatedIdentifiers:   r.RelatedIdentifiers,
		SchemaVersion:        r.SchemaVersion,
		Sizes:                r.Sizes,
		State:                r.State,
		Subjects:             r.Subjects,
		Titles:               r.Titles,
		URL:                  r.URL,
		Version:              r.Version,
	}
}

func (r *Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ToJSONModel())
}
