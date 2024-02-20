// Version: v0.11
// URL:     https://commonmeta.org/commonmeta_v0.11.json

package commonmeta

import (
	"encoding/json"

	"github.com/google/uuid"
)

// a Go function declaration
func UnmarshalCommonMetaJSON(data []byte) (r Resource, err error) {
	err = json.Unmarshal(data, &r)
	return r, err
}

// a Go method declaration
func (r *Resource) MarshalCommonMetaResource() ([]byte, error) {
	return json.Marshal(r)
}

type AlternateIdentifier struct {
	AlternateIdentifier     string `json:"alternateIdentifier"`
	AlternateIdentifierType string `json:"alternateIdentifierType"`
	ResourceID              string
}

// The container of the resource.
type Container struct {
	FirstPage      *string        `json:"firstPage,omitempty"`      // The firstPage of the resource.
	Identifier     *string        `json:"identifier,omitempty"`     // The identifier for the container.
	IdentifierType *string        `json:"identifierType,omitempty"` // The identifierType for the container.
	Issue          *string        `json:"issue,omitempty"`          // The issue of the resource.
	LastPage       *string        `json:"lastPage,omitempty"`       // The last page of the resource.
	Title          *string        `json:"title,omitempty"`          // The title of the container.
	Type           *ContainerType `json:"type,omitempty"`           // The type of the container.
	Volume         *string        `json:"volume,omitempty"`         // The volume of the resource.
	ResourceID     string
}

type Affiliation struct {
	ID            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	ContributorID uuid.UUID
}

type Description struct {
	Description string           `json:"description"`        // The description of the resource.
	Language    *string          `json:"language,omitempty"` // The language of the description. Use one of the language codes from the IETF BCP 47 standard.
	Type        *DescriptionType `json:"type,omitempty"`     // The type of the description.
	ResourceID  string
}

type File struct {
	Bucket     *string `json:"bucket,omitempty"`
	Checksum   *string `json:"checksum,omitempty"`
	Key        *string `json:"key,omitempty"`
	MIMEType   *string `json:"mimeType,omitempty"`
	Size       *int64  `json:"size,omitempty"`
	URL        string  `json:"url"`
	ResourceID string
}

type FundingReference struct {
	AwardURI             *string               `json:"award_uri,omitempty"`
	AwardNumber          *string               `json:"awardNumber,omitempty"`
	FunderIdentifier     *string               `json:"funderIdentifier,omitempty"`
	FunderIdentifierType *FunderIdentifierType `json:"funderIdentifierType,omitempty"`
	FunderName           string                `json:"funderName"`
	ResourceID           string
}

type GeoLocation struct {
	UUID                uuid.UUID `json:"uuid"                          gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	GeoLocationBoxID    int
	GeoLocationBox      *GeoLocationBox      `json:"geoLocationBox,omitempty"`
	GeoLocationPlace    *string              `json:"geoLocationPlace,omitempty"`
	GeoLocationPoint    *GeoLocationPoint    `json:"geoLocationPoint,omitempty"`
	GeoLocationPolygons []GeoLocationPolygon `json:"geoLocationPolygons,omitempty"`
	ResourceID          string
}

type GeoLocationBox struct {
	ID                 int
	EastBoundLongitude float64 `json:"eastBoundLongitude"`
	NorthBoundLatitude float64 `json:"northBoundLatitude"`
	SouthBoundLatitude float64 `json:"southBoundLatitude"`
	WestBoundLongitude float64 `json:"westBoundLongitude"`
}

type GeoLocationPoint struct {
	ID                   int
	PointLatitude        float64 `json:"pointLatitude"`
	PointLongitude       float64 `json:"pointLongitude"`
	GeoLocationID        uuid.UUID
	GeoLocationPolygonID uuid.UUID
}

type GeoLocationPolygon struct {
	UUID             uuid.UUID `json:"uuid"                     gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	InPolygonPointID int
	InPolygonPoint   *GeoLocationPoint  `json:"inPolygonPoint,omitempty"`
	PolygonPoints    []GeoLocationPoint `json:"polygonPoints"`
	GeoLocationID    uuid.UUID
}

type Reference struct {
	ContainerTitle  *string `json:"containerTitle,omitempty"`
	Contributor     *string `json:"contributor,omitempty"`
	Doi             *string `json:"doi,omitempty"`
	Edition         *string `json:"edition,omitempty"`
	FirstPage       *string `json:"firstPage,omitempty"`
	Issue           *string `json:"issue,omitempty"`
	Key             string  `json:"key"`
	LastPage        *string `json:"lastPage,omitempty"`
	PublicationYear *string `json:"publicationYear,omitempty"`
	Publisher       *string `json:"publisher,omitempty"`
	Title           *string `json:"title,omitempty"`
	Unstructured    *string `json:"unstructured,omitempty"`
	Volume          *string `json:"volume,omitempty"`
	ResourceID      string
}

type RelatedIdentifier struct {
	ID         string                `json:"id"`
	Type       RelatedIdentifierType `json:"type"`
	ResourceID string
}

type Subject struct {
	Subject    string `json:"subject"`
	ResourceID string
}

// The location where content is archived.
type ArchiveLocation string

// The type of the container.
type ContainerType string

// The type of the description.
type DescriptionType string

type FunderIdentifierType string

// The provider of the resource. This can be a DOI registration agency or a repository.
type Provider string

type RelatedIdentifierType string

// The schema version of the resource.
type SchemaVersion string

// The state of the resource.
type State string
