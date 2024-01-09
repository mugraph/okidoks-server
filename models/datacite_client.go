package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataCiteTitle struct {
	Title string `json:"title"`
}

type DataCiteAuthor struct {
	Given                   string                    `json:"givenName"`
	Family                  string                    `json:"familyName"`
	DataCiteNameIdentifiers []DataCiteNameIdentifiers `json:"nameIdentifiers"`
	// Affiliation []Affiliation `gorm:"many2many:author_affiliation"`
}

type DataCiteNameIdentifiers struct {
	SchemeURI            string `json:"schemeUri"`
	NameIdentifier       string `json:"nameIdentifier"`
	NameIdentifierScheme string `json:"nameIdentifierScheme"`
}

type DataCiteAttributes struct {
	ID      string           `json:"id"`
	DOI     string           `json:"doi"`
	Prefix  string           `json:"prefix"`
	Titles  []DataCiteTitle  `json:"titles"`
	Authors []DataCiteAuthor `json:"creators"`
}

type DataCiteResource struct {
	ID string `json:"id"`
	// Use struct embedding with anonymous field
	DataCiteAttributes `json:"attributes"`
}

type DataCiteData struct {
	DataCiteResource `json:"data"`
}

func GetDOIMetadataFromDataCite(doi string) (resource Resource, err error) {
	// Construct the API URL for fetching DOI metadata
	apiURL := "https://api.datacite.org/dois/" + doi

	// Make the HTTP GET request with NewRequest and DefaultClient.Do
	// for use with custom header
	request, _ := http.NewRequest("GET", apiURL, nil)
	request.Header.Add("accept", "application/vnd.api+json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Resource{}, fmt.Errorf("Failed to make HTTP request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Resource{}, fmt.Errorf("Failed to read response body: %v", err)
	}

	// // Prettify the response body
	// prettyBody, err := utils.PrettyJSON(body)
	// if err != nil {
	// 	return Resource{}, fmt.Errorf("Failed to prettify response data: %v", err)
	// }

	// // Print the prettified response body
	// fmt.Println(prettyBody + "\n")

	// Unmarshal JSON response
	var data DataCiteData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Resource{}, fmt.Errorf("Failed to unmarshal JSON: %v", err)
	}

	// Save titles in slice of strings
	var titles []string
	for _, item := range data.Titles {
		titles = append(titles, item.Title)
	}

	// Save Authors in slice of Author type
	var authors []Author
	for _, item := range data.Authors {
		ORCID := ""
		for _, identifier := range item.DataCiteNameIdentifiers {
			if identifier.NameIdentifierScheme == "ORCID" {
				ORCID = identifier.NameIdentifier
			}
		}
		author := Author{
			Given:  item.Given,
			Family: item.Family,
			ORCID:  ORCID,
		}
		authors = append(authors, author)
	}

	// Map to Resource type
	resource = Resource{
		DOIAgency: "datacite",
		DOI:       data.DOI,
		Prefix:    data.Prefix,
		Title:     titles,
		Authors:   authors,
	}

	return resource, nil
}
