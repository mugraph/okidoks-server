package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type DOIRAObj struct {
	DOI string
	RA  string
}

// Return the DOI registration agency for a given DOI
func GetDOIRA(doi string) (ra string, err error) {

	// Get prefix
	prefix, err := ValidatePrefix(doi)
	if err != nil || prefix == "" {
		return "", fmt.Errorf("failed to validate Prefix: %v", err)
	}

	// The doi API string
	apiURL := "https://doi.org/ra/" + prefix

	// Make the HTTP GET request
	response, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal JSON response
	var obj []DOIRAObj
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	ra = obj[0].RA

	return ra, nil
}

// Validate a DOI
func ValidateDOI(doi string) (vDoi string, err error) {

	// Define regex pattern
	pattern := `(?:(http|https):\/(\/)?(dx\.)?(doi\.org|handle\.stage\.datacite\.org|handle\.test\.datacite\.org)\/)?(doi:)?(10\.\d{4,5}\/.+)`

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile regex: %v", err)
	}

	// Find macthes in the DOI string
	match := re.FindStringSubmatch(doi)
	if err != nil {
		return "", fmt.Errorf("failed to find string submatch: %v", err)
	} else if len(match) > 6 {
		vDoi = match[6]
	} else {
		// Handle the case when there is no match[6]
		return "", fmt.Errorf("not a valid DOI: %s", doi)
	}

	return vDoi, nil
}

// Validate a DOI prefix for a given DOI
func ValidatePrefix(doi string) (prefix string, err error) {

	// Define regex pattern
	pattern := `^(?:(http|https):\/(\/)?(dx\.)?(doi\.org|handle\.stage\.datacite\.org|handle\.test\.datacite\.org)\/)?(doi:)?(10\.\d{4,5}).*$`

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile regex: %v", err)
	}

	// Find macthes in the DOI string
	match := re.FindStringSubmatch(doi)
	if err != nil {
		return "", fmt.Errorf("failed to find string submatch: %v", err)
	}

	// Make sure match has enough elements
	if len(match) > 6 {
		prefix = match[6]
	} else {
		// Handle the case when there is no match[6]
		err := fmt.Errorf("no prefix found in the match")
		return "", err
	}

	return prefix, nil
}

// Return a DOI from a URL
func DOIFromURL(url string) (doi string, err error) {

	// Define regex pattern
	pattern := `^(?:(http|https)://(dx\.)?(doi\.org|handle\.stage\.datacite\.org|handle\.test\.datacite\.org)/)?(doi:)?(10\.\d{4,5}/.+)$`

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile regex: %v", err)
	}

	// Find macthes in the DOI string
	match := re.FindStringSubmatch(url)
	if err != nil {
		return "", fmt.Errorf("failed to find string submatch: %v", err)
	}

	// Make sure match has enough elements
	if len(match) > 5 {
		doi = match[5]
		doi = strings.ToLower(doi)
	} else {
		// Handle the case when there is no match[6]
		err := fmt.Errorf("no DOI found in the match")
		return "", err
	}

	return doi, nil
}

// Return URL from DOI
func DOIAsURL(doi string) (url string, err error) {
	doi, err = ValidateDOI(doi)
	if err != nil {
		return "", err
	}
	url = "https://doi.org/" + doi
	return url, err
}
