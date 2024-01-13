package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mugraph/okidoks-server/utils"
)

type DataCiteTitle struct {
	Title *string `json:"title"`
}

type DataCiteCreator struct {
	Name                    *string                  `json:"name"`
	GivenName               *string                  `json:"givenName"`
	FamilyName              *string                  `json:"familyName"`
	DataCiteNameIdentifiers []DataCiteNameIdentifier `json:"nameIdentifiers,omitempty"`
	Affiliation             []DataCiteAffiliation    `json:"affiliations,omitempty"`
}

type DataCiteAffiliation struct {
	Items DataCiteAffiliationItems `json:"items"`
}

type DataCiteAffiliationItems struct {
	Properties DataCiteAffiliationProperties `json:"properties"`
}

type DataCiteAffiliationProperties struct {
	Name                        *string `json:"name"`
	SchemeURI                   string  `json:"schemeUri"`
	AffiliationIdentifier       *string `json:"affiliationIdentifier"`
	AffiliationIdentifierScheme *string `json:"affiliationIdentifierScheme"`
}

type DataCiteNameIdentifier struct {
	Items DataCiteNameIdentifierItems `json:"items"`
}

type DataCiteNameIdentifierItems struct {
	Properties DataCiteNameIdentifierProperties `json:"properties"`
}

type DataCiteNameIdentifierProperties struct {
	SchemeURI            string  `json:"schemeUri"`
	NameIdentifier       *string `json:"nameIdentifier"`
	NameIdentifierScheme *string `json:"nameIdentifierScheme"`
}

type DataCiteAttributes struct {
	ID              *string               `json:"id"`
	DOI             *string               `json:"doi"`
	Prefix          *string               `json:"prefix"`
	Titles          []DataCiteTitle       `json:"titles"`
	Creators        []DataCiteCreator     `json:"creators"`
	Publisher       string                `json:"publisher"`
	Container       DataCiteContainer     `json:"container,omitempty"`
	PublicationYear uint                  `json:"publicationYear"`
	Subjects        []DataCiteSubject     `json:"subjects"`
	Contributors    []DataCiteContributor `json:"contributors"`
	// Dates
	// Language
	Types *DataCiteTypes `json:"types"`
}

type DataCiteContributor struct {
	Items DataCiteContributorItems `json:"items"`
}

type DataCiteContributorItems struct {
	Properties DataCiteContributorItemsProperties `json:"properties"`
}

type DataCiteContributorItemsProperties struct {
	ContributorType *string                  `json:"contributorType"`
	Name            *string                  `json:"name"`
	NameType        *string                  `json:"nameType"`
	GivenName       *string                  `json:"givenName"`
	FamilyName      *string                  `json:"familyName"`
	NameIdentifiers []DataCiteNameIdentifier `json:"nameIdentifiers"`
	Affiliation
}

type DataCiteSubject struct {
	Items DataCiteSubjectItems `json:"items"`
}

type DataCiteSubjectItems struct {
	Properties DataCiteSubjectItemsProperties `json:"properties"`
}

type DataCiteSubjectItemsProperties struct {
	Subject       *string `json:"subject"`
	SubjectScheme *string `json:"SubjectScheme"`
	SchemeURI     string  `json:"schemeUri"`
	ValueURI      string  `json:"valueUri"`
	Lang          *string `json:"lang"`
}

type DataCiteTypes struct {
	Ris                 *string `json:"ris"`
	Bibtex              *string `json:"bibtex"`
	Citeproc            *string `json:"citeproc"`
	SchemaOrg           *string `json:"schemaOrg"`
	ResourceType        string  `json:"resourceType"`
	ResourceTypeGeneral string  `json:"resourceTypeGeneral"`
}

type DataCiteContainer struct {
	Properties DataCiteContainerProperties `json:"properties,omitempty"`
}

type DataCiteContainerProperties struct {
	Type      *string `json:"type,omitempty"`
	Name      *string `json:"name,omitempty"`
	FirstPage *string `json:"firstPage,omitempty"`
}

type DataCiteResource struct {
	ID         *string            `json:"id"`
	Attributes DataCiteAttributes `json:"attributes"`
}

type DataCiteData struct {
	DataCiteResource `json:"data"`
}

func GetDataCite(doi string) (attributes DataCiteAttributes, err error) {
	// Construct the API URL for fetching DOI metadata
	apiURL := "https://api.datacite.org/dois/" + doi

	// Make the HTTP GET request with NewRequest and DefaultClient.Do
	// for use with custom header
	request, _ := http.NewRequest("GET", apiURL, nil)
	request.Header.Add("accept", "application/vnd.api+json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("Failed to make HTTP request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("Failed to read response body: %v", err)
	}

	// Unmarshal JSON response
	var data DataCiteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Printf("\n%v\n\n", data.Attributes)

	return data.Attributes, nil
}

func ReadDataCite(attributes DataCiteAttributes) (resource Resource, err error) {
	meta := attributes

	id, err := utils.DOIAsURL(*meta.DOI)
	if err != nil {
		return Resource{}, fmt.Errorf("Could not get DOI as URL: %v", err)
	}

	resourceTypeGeneral := meta.Types.ResourceTypeGeneral
	resourceType := meta.Types.ResourceType

	typeGeneral := utils.DcToCmTranslations[resourceTypeGeneral]
	typeAdditional := utils.DcToCmTranslations[resourceType]

	if typeAdditional != "" {
		typeGeneral = typeAdditional
		typeAdditional = ""
	}

	// contributors := utils.GetAuthors(meta.)
	// contrib := utils.GetAuthors()
	// if contrib != "" {
	// 	contributers = contributors + contrib
	// }

	publisher := Publisher{Name: meta.Publisher}

	resource = Resource{
		ID:   id,
		Type: ResourceType(typeGeneral),
		// Add URL
		// Add Contributors
		Publisher: publisher,
		// Recommended and optional properties
		AdditionalType: &typeAdditional,
	}

	fmt.Printf("%v\n", resource)

	return resource, nil
}
