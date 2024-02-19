// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    coordinate, err := UnmarshalResource(bytes)
//    bytes, err = coordinate.Marshal()

package datacite

import "encoding/json"

func UnmarshalResource(data []byte) (Resource, error) {
	var d DataCiteData
	err := json.Unmarshal(data, &d)
	return d.Attributes, err
}

func (r *Resource) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DataCiteData struct {
	DataCiteResource `json:"data"`
}

type DataCiteResource struct {
	ID         *string  `json:"id"`
	Attributes Resource `json:"attributes"`
}

// JSON representation of the DataCite v4.5 schema.
type Resource struct {
	AlternateIdentifiers []AlternateIdentifier      `json:"alternateIdentifiers,omitempty"`
	Container            *Container                 `json:"container,omitempty"`
	Contributors         []ResourceContributor      `json:"contributors,omitempty"`
	Creators             []ResourceCreator          `json:"creators"`
	Dates                []Date                     `json:"dates,omitempty"`
	Descriptions         []Description              `json:"descriptions,omitempty"`
	Doi                  string                     `json:"doi,omitempty"`
	Formats              []string                   `json:"formats,omitempty"`
	FundingReferences    []FundingReference         `json:"fundingReferences,omitempty"`
	GeoLocations         []GeoLocation              `json:"geoLocations,omitempty"`
	ID                   *string                    `json:"id,omitempty"`
	Language             *string                    `json:"language,omitempty"`
	PublicationYear      uint                       `json:"publicationYear"`
	Publisher            Publisher                  `json:"publisher"`
	RelatedIdentifiers   []RelatedIdentifierElement `json:"relatedIdentifiers,omitempty"`
	RelatedItems         []RelatedItemElement       `json:"relatedItems,omitempty"`
	RightsList           []RightsList               `json:"rightsList,omitempty"`
	SchemaVersion        SchemaVersion              `json:"schemaVersion"`
	Sizes                []string                   `json:"sizes,omitempty"`
	Subjects             []Subject                  `json:"subjects,omitempty"`
	Titles               []ResourceTitle            `json:"titles"`
	Types                Types                      `json:"types"`
	URL                  *string                    `json:"url,omitempty"`
	Version              *string                    `json:"version,omitempty"`
}

type AlternateIdentifier struct {
	AlternateIdentifier     string `json:"alternateIdentifier"`
	AlternateIdentifierType string `json:"alternateIdentifierType"`
}

type Container struct {
	FirstPage *string `json:"firstPage,omitempty"`
	Title     *string `json:"title,omitempty"`
	Type      *string `json:"type,omitempty"`
}

type ResourceContributor struct {
	ContributorType ContributorType  `json:"contributorType"`
	Name            string           `json:"name"`
	Affiliation     []Affiliation    `json:"affiliation,omitempty"`
	FamilyName      *string          `json:"familyName,omitempty"`
	GivenName       *string          `json:"givenName,omitempty"`
	Lang            *string          `json:"lang,omitempty"`
	NameIdentifiers []NameIdentifier `json:"nameIdentifiers,omitempty"`
	NameType        *NameType        `json:"nameType,omitempty"`
}

type Affiliation struct {
	AffiliationIdentifier       *string `json:"affiliationIdentifier,omitempty"`
	AffiliationIdentifierScheme *string `json:"affiliationIdentifierScheme,omitempty"`
	Name                        string  `json:"name"`
	SchemeURI                   *string `json:"schemeUri,omitempty"`
}

type NameIdentifier struct {
	NameIdentifier       string  `json:"nameIdentifier"`
	NameIdentifierScheme string  `json:"nameIdentifierScheme"`
	SchemeURI            *string `json:"schemeUri,omitempty"`
}

type ResourceCreator struct {
	Name            string           `json:"name"`
	Affiliation     []Affiliation    `json:"affiliation,omitempty"`
	FamilyName      *string          `json:"familyName,omitempty"`
	GivenName       *string          `json:"givenName,omitempty"`
	Lang            *string          `json:"lang,omitempty"`
	NameIdentifiers []NameIdentifier `json:"nameIdentifiers,omitempty"`
	NameType        *NameType        `json:"nameType,omitempty"`
}

type Date struct {
	Date            string   `json:"date"`
	DateInformation *string  `json:"dateInformation,omitempty"`
	DateType        DateType `json:"dateType"`
}

type Description struct {
	Description     string          `json:"description"`
	DescriptionType DescriptionType `json:"descriptionType"`
	Lang            *string         `json:"lang,omitempty"`
}

type FundingReference struct {
	AwardNumber          *string               `json:"awardNumber,omitempty"`
	AwardTitle           *string               `json:"awardTitle,omitempty"`
	AwardURI             *string               `json:"awardUri,omitempty"`
	FunderIdentifier     *string               `json:"funderIdentifier,omitempty"`
	FunderIdentifierType *FunderIdentifierType `json:"funderIdentifierType,omitempty"`
	FunderName           string                `json:"funderName"`
}

type GeoLocation struct {
	GeoLocationBox      *GeoLocationBox      `json:"geoLocationBox,omitempty"`
	GeoLocationPlace    *string              `json:"geoLocationPlace,omitempty"`
	GeoLocationPoint    *GeoLocationPoint    `json:"geoLocationPoint,omitempty"`
	GeoLocationPolygons []GeoLocationPolygon `json:"geoLocationPolygons,omitempty"`
}

type GeoLocationBox struct {
	EastBoundLongitude float64 `json:"eastBoundLongitude"`
	NorthBoundLatitude float64 `json:"northBoundLatitude"`
	SouthBoundLatitude float64 `json:"southBoundLatitude"`
	WestBoundLongitude float64 `json:"westBoundLongitude"`
}

type GeoLocationPoint struct {
	PointLatitude  float64 `json:"pointLatitude"`
	PointLongitude float64 `json:"pointLongitude"`
}

type GeoLocationPolygon struct {
	InPolygonPoint *GeoLocationPoint  `json:"inPolygonPoint,omitempty"`
	PolygonPoints  []GeoLocationPoint `json:"polygonPoints"`
}

type Publisher struct {
	Lang                      *string `json:"lang,omitempty"`
	Name                      string  `json:"name"`
	PublisherIdentifier       *string `json:"publisherIdentifier,omitempty"`
	PublisherIdentifierScheme *string `json:"publisherIdentifierScheme,omitempty"`
	SchemeURI                 *string `json:"schemeURI,omitempty"`
}

type RelatedIdentifierElement struct {
	RelatedIdentifier     string                `json:"relatedIdentifier"`
	RelatedIdentifierType RelatedIdentifierType `json:"relatedIdentifierType"`
	RelationType          RelationType          `json:"relationType"`
	RelatedMetadataScheme *string               `json:"relatedMetadataScheme,omitempty"`
	ResourceTypeGeneral   *ResourceTypeGeneral  `json:"resourceTypeGeneral,omitempty"`
	SchemeType            *string               `json:"schemeType,omitempty"`
	SchemeURI             *string               `json:"schemeUri,omitempty"`
}

type RelatedItemElement struct {
	Contributors          []RelatedItemContributor `json:"contributors,omitempty"`
	Creators              []RelatedItemCreator     `json:"creators,omitempty"`
	Edition               *string                  `json:"edition,omitempty"`
	FirstPage             *string                  `json:"firstPage,omitempty"`
	Issue                 *string                  `json:"issue,omitempty"`
	LastPage              *string                  `json:"lastPage,omitempty"`
	Number                *string                  `json:"number,omitempty"`
	NumberType            *NumberType              `json:"numberType,omitempty"`
	PublicationYear       *string                  `json:"publicationYear,omitempty"`
	Publisher             *string                  `json:"publisher,omitempty"`
	RelatedItemIdentifier *RelatedItemIdentifier   `json:"relatedItemIdentifier,omitempty"`
	RelatedItemType       ResourceTypeGeneral      `json:"relatedItemType"`
	Titles                []RelatedItemTitle       `json:"titles"`
	Volume                *string                  `json:"volume,omitempty"`
	RelationType          RelationType             `json:"relationType"`
	RelatedMetadataScheme *string                  `json:"relatedMetadataScheme,omitempty"`
	ResourceTypeGeneral   *ResourceTypeGeneral     `json:"resourceTypeGeneral,omitempty"`
	SchemeType            *string                  `json:"schemeType,omitempty"`
	SchemeURI             *string                  `json:"schemeUri,omitempty"`
}

type RelatedItemContributor struct {
	ContributorType ContributorType  `json:"contributorType"`
	Name            string           `json:"name"`
	Affiliation     []Affiliation    `json:"affiliation,omitempty"`
	FamilyName      *string          `json:"familyName,omitempty"`
	GivenName       *string          `json:"givenName,omitempty"`
	Lang            *string          `json:"lang,omitempty"`
	NameIdentifiers []NameIdentifier `json:"nameIdentifiers,omitempty"`
	NameType        *NameType        `json:"nameType,omitempty"`
}

type RelatedItemCreator struct {
	Name            string           `json:"name"`
	Affiliation     []Affiliation    `json:"affiliation,omitempty"`
	FamilyName      *string          `json:"familyName,omitempty"`
	GivenName       *string          `json:"givenName,omitempty"`
	Lang            *string          `json:"lang,omitempty"`
	NameIdentifiers []NameIdentifier `json:"nameIdentifiers,omitempty"`
	NameType        *NameType        `json:"nameType,omitempty"`
}

type RelatedItemIdentifier struct {
	RelatedItemIdentifier     string                `json:"relatedItemIdentifier"`
	RelatedItemIdentifierType RelatedIdentifierType `json:"relatedItemIdentifierType"`
}

type RelatedItemTitle struct {
	Lang      *string    `json:"lang,omitempty"`
	Title     string     `json:"title"`
	TitleType *TitleType `json:"titleType,omitempty"`
}

type RightsList struct {
	Lang                   *string `json:"lang,omitempty"`
	Rights                 *string `json:"rights,omitempty"`
	RightsIdentifier       *string `json:"rightsIdentifier,omitempty"`
	RightsIdentifierScheme *string `json:"rightsIdentifierScheme,omitempty"`
	RightsURI              *string `json:"rightsUri,omitempty"`
	SchemeURI              *string `json:"schemeUri,omitempty"`
}

type Subject struct {
	ClassificationCode *string `json:"classificationCode,omitempty"`
	Lang               *string `json:"lang,omitempty"`
	SchemeURI          *string `json:"schemeUri,omitempty"`
	Subject            string  `json:"subject"`
	SubjectScheme      *string `json:"subjectScheme,omitempty"`
	ValueURI           *string `json:"valueUri,omitempty"`
}

type ResourceTitle struct {
	Lang      *string    `json:"lang,omitempty"`
	Title     string     `json:"title"`
	TitleType *TitleType `json:"titleType,omitempty"`
}

type Types struct {
	ResourceType        string              `json:"resourceType,omitempty"`
	ResourceTypeGeneral ResourceTypeGeneral `json:"resourceTypeGeneral"`
}

type ContributorType string

const (
	ContactPerson         ContributorType = "ContactPerson"
	ContributorTypeOther  ContributorType = "Other"
	DataCollector         ContributorType = "DataCollector"
	DataCurator           ContributorType = "DataCurator"
	DataManager           ContributorType = "DataManager"
	Distributor           ContributorType = "Distributor"
	Editor                ContributorType = "Editor"
	HostingInstitution    ContributorType = "HostingInstitution"
	Producer              ContributorType = "Producer"
	ProjectLeader         ContributorType = "ProjectLeader"
	ProjectManager        ContributorType = "ProjectManager"
	ProjectMember         ContributorType = "ProjectMember"
	RegistrationAgency    ContributorType = "RegistrationAgency"
	RegistrationAuthority ContributorType = "RegistrationAuthority"
	RelatedPerson         ContributorType = "RelatedPerson"
	ResearchGroup         ContributorType = "ResearchGroup"
	Researcher            ContributorType = "Researcher"
	RightsHolder          ContributorType = "RightsHolder"
	Sponsor               ContributorType = "Sponsor"
	Supervisor            ContributorType = "Supervisor"
	WorkPackageLeader     ContributorType = "WorkPackageLeader"
)

type NameType string

const (
	Organizational NameType = "Organizational"
	Personal       NameType = "Personal"
)

type DateType string

const (
	Accepted      DateType = "Accepted"
	Available     DateType = "Available"
	Collected     DateType = "Collected"
	Copyrighted   DateType = "Copyrighted"
	Created       DateType = "Created"
	DateTypeOther DateType = "Other"
	Issued        DateType = "Issued"
	Submitted     DateType = "Submitted"
	Updated       DateType = "Updated"
	Valid         DateType = "Valid"
	Withdrawn     DateType = "Withdrawn"
)

type DescriptionType string

const (
	Abstract             DescriptionType = "Abstract"
	DescriptionTypeOther DescriptionType = "Other"
	Methods              DescriptionType = "Methods"
	SeriesInformation    DescriptionType = "SeriesInformation"
	TableOfContents      DescriptionType = "TableOfContents"
	TechnicalInfo        DescriptionType = "TechnicalInfo"
)

type FunderIdentifierType string

const (
	CrossrefFunderID          FunderIdentifierType = "Crossref Funder ID"
	FunderIdentifierTypeOther FunderIdentifierType = "Other"
	Grid                      FunderIdentifierType = "GRID"
	Isni                      FunderIdentifierType = "ISNI"
	Ror                       FunderIdentifierType = "ROR"
)

type RelatedIdentifierType string

const (
	ArXiv   RelatedIdentifierType = "arXiv"
	Ark     RelatedIdentifierType = "ARK"
	Bibcode RelatedIdentifierType = "bibcode"
	Doi     RelatedIdentifierType = "DOI"
	Ean13   RelatedIdentifierType = "EAN13"
	Eissn   RelatedIdentifierType = "EISSN"
	Handle  RelatedIdentifierType = "Handle"
	Igsn    RelatedIdentifierType = "IGSN"
	Isbn    RelatedIdentifierType = "ISBN"
	Issn    RelatedIdentifierType = "ISSN"
	Istc    RelatedIdentifierType = "ISTC"
	Lissn   RelatedIdentifierType = "LISSN"
	Lsid    RelatedIdentifierType = "LSID"
	Pmid    RelatedIdentifierType = "PMID"
	Purl    RelatedIdentifierType = "PURL"
	URL     RelatedIdentifierType = "URL"
	Upc     RelatedIdentifierType = "UPC"
	Urn     RelatedIdentifierType = "URN"
	W3ID    RelatedIdentifierType = "w3id"
)

type RelationType string

const (
	Cites               RelationType = "Cites"
	Collects            RelationType = "Collects"
	Compiles            RelationType = "Compiles"
	Continues           RelationType = "Continues"
	Describes           RelationType = "Describes"
	Documents           RelationType = "Documents"
	HasMetadata         RelationType = "HasMetadata"
	HasPart             RelationType = "HasPart"
	HasVersion          RelationType = "HasVersion"
	IsCitedBy           RelationType = "IsCitedBy"
	IsCollectedBy       RelationType = "IsCollectedBy"
	IsCompiledBy        RelationType = "IsCompiledBy"
	IsContinuedBy       RelationType = "IsContinuedBy"
	IsDerivedFrom       RelationType = "IsDerivedFrom"
	IsDescribedBy       RelationType = "IsDescribedBy"
	IsDocumentedBy      RelationType = "IsDocumentedBy"
	IsIdenticalTo       RelationType = "IsIdenticalTo"
	IsMetadataFor       RelationType = "IsMetadataFor"
	IsNewVersionOf      RelationType = "IsNewVersionOf"
	IsObsoletedBy       RelationType = "IsObsoletedBy"
	IsOriginalFormOf    RelationType = "IsOriginalFormOf"
	IsPartOf            RelationType = "IsPartOf"
	IsPreviousVersionOf RelationType = "IsPreviousVersionOf"
	IsPublishedIn       RelationType = "IsPublishedIn"
	IsReferencedBy      RelationType = "IsReferencedBy"
	IsRequiredBy        RelationType = "IsRequiredBy"
	IsReviewedBy        RelationType = "IsReviewedBy"
	IsSourceOf          RelationType = "IsSourceOf"
	IsSupplementTo      RelationType = "IsSupplementTo"
	IsSupplementedBy    RelationType = "IsSupplementedBy"
	IsVariantFormOf     RelationType = "IsVariantFormOf"
	IsVersionOf         RelationType = "IsVersionOf"
	Obsoletes           RelationType = "Obsoletes"
	References          RelationType = "References"
	Requires            RelationType = "Requires"
	Reviews             RelationType = "Reviews"
)

type ResourceTypeGeneral string

const (
	Audiovisual               ResourceTypeGeneral = "Audiovisual"
	Book                      ResourceTypeGeneral = "Book"
	BookChapter               ResourceTypeGeneral = "BookChapter"
	Collection                ResourceTypeGeneral = "Collection"
	ComputationalNotebook     ResourceTypeGeneral = "ComputationalNotebook"
	ConferencePaper           ResourceTypeGeneral = "ConferencePaper"
	ConferenceProceeding      ResourceTypeGeneral = "ConferenceProceeding"
	DataPaper                 ResourceTypeGeneral = "DataPaper"
	Dataset                   ResourceTypeGeneral = "Dataset"
	Dissertation              ResourceTypeGeneral = "Dissertation"
	Event                     ResourceTypeGeneral = "Event"
	Image                     ResourceTypeGeneral = "Image"
	Instrument                ResourceTypeGeneral = "Instrument"
	InteractiveResource       ResourceTypeGeneral = "InteractiveResource"
	Journal                   ResourceTypeGeneral = "Journal"
	JournalArticle            ResourceTypeGeneral = "JournalArticle"
	Model                     ResourceTypeGeneral = "Model"
	OutputManagementPlan      ResourceTypeGeneral = "OutputManagementPlan"
	PeerReview                ResourceTypeGeneral = "PeerReview"
	PhysicalObject            ResourceTypeGeneral = "PhysicalObject"
	Preprint                  ResourceTypeGeneral = "Preprint"
	ResourceTypeGeneralOther  ResourceTypeGeneral = "Other"
	ResourceTypeGeneralReport ResourceTypeGeneral = "Report"
	Service                   ResourceTypeGeneral = "Service"
	Software                  ResourceTypeGeneral = "Software"
	Sound                     ResourceTypeGeneral = "Sound"
	Standard                  ResourceTypeGeneral = "Standard"
	StudyRegistration         ResourceTypeGeneral = "StudyRegistration"
	Text                      ResourceTypeGeneral = "Text"
	Workflow                  ResourceTypeGeneral = "Workflow"
)

type NumberType string

const (
	Article          NumberType = "Article"
	Chapter          NumberType = "Chapter"
	NumberTypeOther  NumberType = "Other"
	NumberTypeReport NumberType = "Report"
)

type TitleType string

const (
	AlternativeTitle TitleType = "AlternativeTitle"
	Subtitle         TitleType = "Subtitle"
	TitleTypeOther   TitleType = "Other"
	TranslatedTitle  TitleType = "TranslatedTitle"
)

type SchemaVersion string

const (
	HTTPDataciteOrgSchemaKernel4 SchemaVersion = "http://datacite.org/schema/kernel-4"
)
