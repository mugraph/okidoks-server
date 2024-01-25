// Version: v0.10.5
// URL:     https://commonmeta.org/commonmeta_v0.10.5.json
package models

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
	ID         *string        `json:"id,omitempty"`    // The identifier for the container.
	Title      *string        `json:"title,omitempty"` // The title of the container.
	Type       *ContainerType `json:"type,omitempty"`  // The type of the container.
	ResourceID string
}

type Affiliation struct {
	ID            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	ContributorID uuid.UUID
}

// The dates for the resource.
type Date struct {
	Accepted   *string `json:"accepted,omitempty"`  // The date the resource was accepted.
	Available  *string `json:"available,omitempty"` // The date the resource was made available.
	Created    *string `json:"created,omitempty"`   // The date the resource was created.
	Published  *string `json:"published,omitempty"` // The date the resource was published.
	Submitted  *string `json:"submitted,omitempty"` // The date the resource was submitted.
	Updated    *string `json:"updated,omitempty"`   // The date the resource was updated.
	Withdrawn  *string `json:"withdrawn,omitempty"` // The date the resource was withdrawn.
	ResourceID string  `json:"-"`
}

type DescriptionElement struct {
	Description string           `json:"description"`    // The description of the resource.
	Type        *DescriptionType `json:"type,omitempty"` // The type of the description.
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
	UUID                uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
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
	UUID             uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
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

type Title struct {
	Title      string     `json:"title"`          // The title of the resource.
	Type       *TitleType `json:"type,omitempty"` // The type of the title.
	ResourceID string
}

type ArchiveLocation string

const (
	Clockss         ArchiveLocation = "CLOCKSS"
	Dwt             ArchiveLocation = "DWT"
	InternetArchive ArchiveLocation = "Internet Archive"
	KB              ArchiveLocation = "KB"
	Lockss          ArchiveLocation = "LOCKSS"
	Portico         ArchiveLocation = "Portico"
)

// The type of the container.
type ContainerType string

const (
	DataCatalog             ContainerType = "DataCatalog"
	Periodical              ContainerType = "Periodical"
	PurpleBook              ContainerType = "Book"
	PurpleBookSeries        ContainerType = "BookSeries"
	PurpleJournal           ContainerType = "Journal"
	PurpleProceedingsSeries ContainerType = "ProceedingsSeries"
	Repository              ContainerType = "Repository"
	Series                  ContainerType = "Series"
)

// The type of the contributor.
type ContributorType string

const (
	Organization ContributorType = "Organization"
	Person       ContributorType = "Person"
)

// The type of the description.
type DescriptionType string

const (
	Abstract    DescriptionType = "Abstract"
	Description DescriptionType = "Description"
	Summary     DescriptionType = "Summary"
)

type FunderIdentifierType string

const (
	CrossrefFunderID          FunderIdentifierType = "Crossref Funder ID"
	FunderIdentifierTypeOther FunderIdentifierType = "Other"
	Grid                      FunderIdentifierType = "GRID"
	Isni                      FunderIdentifierType = "ISNI"
	Ringgold                  FunderIdentifierType = "Ringgold"
	Ror                       FunderIdentifierType = "ROR"
)

// The provider of the resource. This can be a DOI registration agency or a repository.
type Provider string

const (
	Crossref Provider = "Crossref"
	DataCite Provider = "DataCite"
	GitHub   Provider = "GitHub"
	JaLC     Provider = "JaLC"
	Kisti    Provider = "KISTI"
	MEDRA    Provider = "mEDRA"
	Op       Provider = "OP"
)

type RelatedIdentifierType string

const (
	HasPart             RelatedIdentifierType = "HasPart"
	HasPreprint         RelatedIdentifierType = "HasPreprint"
	HasVersion          RelatedIdentifierType = "HasVersion"
	IsIdenticalTo       RelatedIdentifierType = "IsIdenticalTo"
	IsNewVersionOf      RelatedIdentifierType = "IsNewVersionOf"
	IsOriginalFormOf    RelatedIdentifierType = "IsOriginalFormOf"
	IsPartOf            RelatedIdentifierType = "IsPartOf"
	IsPreprintOf        RelatedIdentifierType = "IsPreprintOf"
	IsPreviousVersionOf RelatedIdentifierType = "IsPreviousVersionOf"
	IsReviewedBy        RelatedIdentifierType = "IsReviewedBy"
	IsSupplementTo      RelatedIdentifierType = "isSupplementTo"
	IsTranslationOf     RelatedIdentifierType = "IsTranslationOf"
	IsVariantFormOf     RelatedIdentifierType = "IsVariantFormOf"
	IsVersionOf         RelatedIdentifierType = "IsVersionOf"
	Reviews             RelatedIdentifierType = "Reviews"
)

// The schema version of the resource.
type SchemaVersion string

const (
	HTTPDataciteOrgSchemaKernel3          SchemaVersion = "http://datacite.org/schema/kernel-3"
	HTTPDataciteOrgSchemaKernel4          SchemaVersion = "http://datacite.org/schema/kernel-4"
	HTTPSCommonmetaOrgCommonmetaV0105JSON SchemaVersion = "https://commonmeta.org/commonmeta_v0.10.5.json"
)

// The state of the resource.
type State string

const (
	Findable State = "findable"
	NotFound State = "not_found"
)

// The type of the title.
type TitleType string

const (
	AlternativeTitle TitleType = "AlternativeTitle"
	Subtitle         TitleType = "Subtitle"
	TranslatedTitle  TitleType = "TranslatedTitle"
)
