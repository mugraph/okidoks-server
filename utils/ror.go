package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateROR
func ValidateROR(ror string) (string, error) {
	pattern := `(?:(?:http|https):\/\/ror\.org\/)?([0-9a-z]{7}\d{2})`

	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(ror)
	if len(match) >= 1 {
		ror = match[1]
		ror = strings.ReplaceAll(ror, " ", "-")
	} else {
		// Handle the case when there is no match[6]
		return "", fmt.Errorf("not a valid ROR: %v", ror)
	}
	return ror, nil
}
