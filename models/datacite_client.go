package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mugraph/okidoks-server/utils"
)

type DataCiteData struct {
	DataCiteResource `json:"data"`
}

type DataCiteResource struct {
	ID         *string            `json:"id"`
	Attributes DataCiteAttributes `json:"attributes"`
}

type DataCiteAttributes struct {
	ID                 *string               `json:"id"`
	DOI                *string               `json:"doi"`
	Prefix             *string               `json:"prefix"`
	Titles             []DataCiteTitle       `json:"titles"`
	Creators           []DataCiteCreator     `json:"creators"`
	Publisher          string                `json:"publisher"`
	Container          DataCiteContainer     `json:"container,omitempty"`
	PublicationYear    uint                  `json:"publicationYear"`
	Subjects           []DataCiteSubject     `json:"subjects"`
	Contributors       []DataCiteContributor `json:"contributors"`
	Dates              []DataCiteDate        `json:"dates"`
	Language           string                `json:"language"`
	Types              *DataCiteTypes        `json:"types"`
	RelatedIdentifiers []RelatedIdentifier   `json:"relatedIdentifier"`
	URL                string                `json:"url"`
}

type DataCiteTitle struct {
	Title *string `json:"title"`
}

type DataCiteTypes struct {
	Ris                 *string `json:"ris"`
	Bibtex              *string `json:"bibtex"`
	Citeproc            *string `json:"citeproc"`
	SchemaOrg           *string `json:"schemaOrg"`
	ResourceType        string  `json:"resourceType"`
	ResourceTypeGeneral string  `json:"resourceTypeGeneral"`
}

type DataCiteCreator struct {
	Name            *string                  `json:"name"`
	GivenName       *string                  `json:"givenName"`
	FamilyName      *string                  `json:"familyName"`
	NameIdentifiers []DataCiteNameIdentifier `json:"nameIdentifiers,omitempty"`
	Affiliation     []DataCiteAffiliation    `json:"affiliations,omitempty"`
}

type DataCiteAffiliation struct {
	Name                        *string `json:"name"`
	SchemeURI                   string  `json:"schemeUri"`
	AffiliationIdentifier       *string `json:"affiliationIdentifier"`
	AffiliationIdentifierScheme *string `json:"affiliationIdentifierScheme"`
}

type DataCiteContainer struct {
	Properties DataCiteContainerProperties `json:"properties,omitempty"`
}

type DataCiteContainerProperties struct {
	Type      *string `json:"type,omitempty"`
	Name      *string `json:"name,omitempty"`
	FirstPage *string `json:"firstPage,omitempty"`
}

type DataCiteSubject struct {
	Subject       *string `json:"subject"`
	SubjectScheme *string `json:"SubjectScheme"`
	SchemeURI     string  `json:"schemeUri"`
	ValueURI      string  `json:"valueUri"`
	Lang          *string `json:"lang"`
}

type DataCiteNameIdentifier struct {
	SchemeURI            string  `json:"schemeUri"`
	NameIdentifier       *string `json:"nameIdentifier"`
	NameIdentifierScheme *string `json:"nameIdentifierScheme"`
}

type DataCiteContributor struct {
	ContributorType *string                  `json:"contributorType"`
	Name            *string                  `json:"name"`
	NameType        *string                  `json:"nameType"`
	GivenName       *string                  `json:"givenName"`
	FamilyName      *string                  `json:"familyName"`
	NameIdentifiers []DataCiteNameIdentifier `json:"nameIdentifiers"`
	Affiliation
}

type DataCiteDate struct {
	Date     string `json:"date"`
	DateType string `json:"dateType"`
}

type DataCiteRelatedIdentifier struct {
	RelationType          string `json:"relationType"`
	RelatedIdentifier     string `json:"relatedIdentifier"`
	RelationTypeGeneral   string `json:"relationTypeGeneral"`
	RelatedIdentifierType string `json:"relatedIdentifierType"`
}

func GetDataCite(doi string) (attributes DataCiteAttributes, err error) {
	doi, err = utils.ValidateDOI(doi)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("Failed to validate DOI: %v", err)
	}

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

	return data.Attributes, nil
}

func ReadDataCite(attributes DataCiteAttributes) (resource Resource, err error) {
	meta := attributes

	id, err := utils.DOIAsURL(*meta.DOI)
	if err != nil {
		return Resource{}, fmt.Errorf("Could not get DOI as URL: %v", err)
	}

	// Define resource type according to Commonmeta
	resourceTypeGeneral := meta.Types.ResourceTypeGeneral
	resourceType := meta.Types.ResourceType

	typeGeneral := utils.DcToCmTranslations[resourceTypeGeneral]
	typeAdditional := utils.DcToCmTranslations[resourceType]

	if typeAdditional != "" {
		typeGeneral = typeAdditional
		typeAdditional = ""
	}

	publisher := Publisher{Name: meta.Publisher}

	contributors := []Contributor{}
	for _, v := range meta.Contributors {
		contributorRoles := []ContributorRole{}
		contributorRoles = append(contributorRoles, ContributorRole{Role: Role(*v.ContributorType)})
		fmt.Printf("%v\n", contributorRoles)

		// if v.ContributorType is Person
		contributors = append(contributors, Contributor{
			GivenName:        v.GivenName,
			FamilyName:       v.FamilyName,
			ContributorRoles: contributorRoles,
		})
		// else if is Organization
		// contributors = append(contributors, Contributor{
		// 	Name:  v.Name,
		// })
	}
	contributorsJSON, err := json.MarshalIndent(contributors, "", "  ")
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Printf("MarshalIndent Contributors: %s\n", string(contributorsJSON))

	resource = Resource{
		ID:           id,
		Type:         ResourceType(typeGeneral),
		URL:          utils.NormalizeURL(meta.URL, true, true),
		Contributors: contributors,
		Publisher:    publisher,
		// Recommended and optional properties
		AdditionalType: &typeAdditional,
	}

	// fmt.Printf("%v\n", resource)

	return resource, nil
}
