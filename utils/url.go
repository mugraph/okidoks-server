package utils

import (
	"net/url"
	"regexp"
)

// ValidateURL checks whether URL is DOI, URL or ISSN and returns a string
func ValidateURL(s string) string {
	if s == "" {
		return ""
	}
	// Check for DOI
	if _, err := ValidateDOI(s); err != nil {
		log.Warn("could not validate DOI", "error", err)
	} else {
		return "DOI"
	}
	// Check for URL
	if _, err := url.ParseRequestURI(s); err != nil {
		log.Warn("could not validate URL", "error", err)
	} else {
		return "URL"
	}
	// Check for ISSN
	pattern := `(ISSN|eISSN) (\d{4}-\d{3}[0-9X]+)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Warn("failed to compile regex", err)
	}

	match := re.FindStringSubmatch(s)
	if len(match) >= 1 {
		return "ISSN"
	}

	return ""
}
