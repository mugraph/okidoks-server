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

func GetDataCite(doi string) (attr DataCiteAttributes, err error) {
	doi, err = utils.ValidateDOI(doi)
	if err != nil {
		return attr, fmt.Errorf("failed to validate DOI: %v", err)
	}

	// Construct the API URL for fetching DOI attrdata
	apiURL := "https://api.datacite.org/dois/" + doi +
		"?publisher=true&affiliation=true"

	// Make the HTTP GET request with NewRequest and DefaultClient.Do
	// for use with custom header
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("accept", "application/vnd.api+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return attr, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Return on non-200 status codes
	// DataCite API complies with JSON:API v1.1 where a server MUST respond with
	// 200 OK for a successful request:
	// https://jsonapi.org/format/#fetching-resources-responses
	if resp.StatusCode != http.StatusOK {
		return attr, fmt.Errorf("HTTP response status: %v %v",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return attr, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal JSON response
	var data DataCiteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return attr, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return data.Attributes, nil
}

func ReadDataCite(attr DataCiteAttributes) (rs Resource, err error) {

	id, err := utils.DOIAsURL(*attr.DOI)
	if err != nil {
		return Resource{}, fmt.Errorf("could not get DOI as URL: %v", err)
	}

	// Define resource type according to Commonattr
	rtg := attr.Types.ResourceTypeGeneral
	rt := attr.Types.ResourceType
	general := utils.DataCiteToCommonmeta[rtg]
	additional := utils.DataCiteToCommonmeta[rt]

	if additional != "" {
		general = additional
		additional = ""
	}

	p := Publisher{Name: *attr.Publisher.Name}

	c := []Contributor{}
	for _, v := range attr.Contributors {
		roles := []ContributorRole{}
		roles = append(roles, ContributorRole{Role: Role(*v.ContributorType)})

		// if v.ContributorType is Person
		c = append(c, Contributor{
			GivenName:        v.GivenName,
			FamilyName:       v.FamilyName,
			ContributorRoles: roles,
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

	rs = Resource{
		ID:           id,
		Type:         ResourceType(general),
		URL:          utils.NormalizeURL(attr.URL, true, true),
		Contributors: c,
		Publisher:    p,
		// Recommended and optional properties
		AdditionalType: &additional,
		License:        &license,
	}

	return rs, nil
}
