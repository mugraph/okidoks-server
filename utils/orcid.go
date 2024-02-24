package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateORCID
func ValidateORCID(orcid string) (string, error) {
	pattern := `(?:(?:http|https):\/\/(?:(?:www|sandbox)?\.)?orcid\.org\/)?(\d{4}[ -]\d{4}[ -]\d{4}[ -]\d{3}[0-9X]+)`

	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(orcid)
	if len(match) >= 1 {
		orcid = match[1]
		orcid = strings.ReplaceAll(orcid, " ", "-")
	} else {
		// Handle the case when there is no match[6]
		return "", fmt.Errorf("not a valid ORCID: %v", orcid)
	}
	return orcid, nil
}
