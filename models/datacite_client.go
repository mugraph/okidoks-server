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
	Publisher          DataCitePublisher     `json:"publisher"`
	Container          DataCiteContainer     `json:"container,omitempty"`
	PublicationYear    uint                  `json:"publicationYear"`
	Subjects           []DataCiteSubject     `json:"subjects"`
	Contributors       []DataCiteContributor `json:"contributors"`
	Dates              []DataCiteDate        `json:"dates"`
	Language           string                `json:"language"`
	Types              *DataCiteTypes        `json:"types"`
	RelatedIdentifiers []RelatedIdentifier   `json:"relatedIdentifier"`
	URL                string                `json:"url"`
	RightsList         []DataCiteRightsList  `json:"rightsList"`
}

type DataCiteTitle struct {
	Title *string `json:"title"`
}

type DataCitePublisher struct {
	Name                      *string `json:"name"`
	SchemeURI                 *string `json:"schemeUri"`
	PublisherIdentifier       *string `json:"publisherIdentifier"`
	PublisherIdentifierScheme *string `json:"publisherIdentifierScheme"`
	Lang                      *string `json:"lang"`
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

type DataCiteRightsList struct {
	Lang                   *string `json:"lang,omitempty"`
	Rights                 *string `json:"rights,omitempty"`
	RightsIdentifier       *string `json:"rightsIdentifier,omitempty"`
	RightsIdentifierScheme *string `json:"rightsIdentifierScheme,omitempty"`
	RightsURI              *string `json:"rightsUri,omitempty"`
	SchemeURI              *string `json:"schemeUri,omitempty"`
}

func GetDataCite(doi string) (attributes DataCiteAttributes, err error) {
	doi, err = utils.ValidateDOI(doi)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("failed to validate DOI: %v", err)
	}

	// Construct the API URL for fetching DOI attrdata
	apiURL := "https://api.datacite.org/dois/" + doi + "?publisher=true&affiliation=true"

	// Make the HTTP GET request with NewRequest and DefaultClient.Do
	// for use with custom header
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("accept", "application/vnd.api+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Return on non-200 status codes
	// DataCite API complies with JSON:API v1.1 where a server MUST respond with
	// 200 OK for a successful request:
	// https://jsonapi.org/format/#fetching-resources-responses
	if resp.StatusCode != http.StatusOK {
		return DataCiteAttributes{}, fmt.Errorf("HTTP response status: %v %v", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal JSON response
	var data DataCiteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return DataCiteAttributes{}, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return data.Attributes, nil
}

func ReadDataCite(attr DataCiteAttributes) (resource Resource, err error) {

	id, err := utils.DOIAsURL(*attr.DOI)
	if err != nil {
		return Resource{}, fmt.Errorf("could not get DOI as URL: %v", err)
	}

	// Define resource type according to Commonattr
	resourceTypeGeneral := attr.Types.ResourceTypeGeneral
	resourceType := attr.Types.ResourceType

	typeGeneral := utils.DcToCmTranslations[resourceTypeGeneral]
	typeAdditional := utils.DcToCmTranslations[resourceType]

	if typeAdditional != "" {
		typeGeneral = typeAdditional
		typeAdditional = ""
	}

	publisher := Publisher{Name: *attr.Publisher.Name}

	contributors := []Contributor{}
	for _, v := range attr.Contributors {
		contributorRoles := []ContributorRole{}
		contributorRoles = append(contributorRoles, ContributorRole{Role: Role(*v.ContributorType)})

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

	license := License{}
	if len(attr.RightsList) > 0 {
		license = License{
			URL: attr.RightsList[0].RightsURI,
		}
	}

	resource = Resource{
		ID:           id,
		Type:         ResourceType(typeGeneral),
		URL:          utils.NormalizeURL(attr.URL, true, true),
		Contributors: contributors,
		Publisher:    publisher,
		// Recommended and optional properties
		AdditionalType: &typeAdditional,
		License:        &license,
	}

	return resource, nil
}
