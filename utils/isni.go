package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateISNI
func ValidateISNI(isni string) (string, error) {
	pattern := `(?:(?:http|https):\/\/isni\.org\/isni\/)?(\d{4}([ -])?\d{4}([ -])?\d{4}([ -])?\d{3}[0-9X]+)`

	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(isni)
	if len(match) >= 1 {
		isni = match[1]
		isni = strings.ReplaceAll(isni, " ", "")
	} else {
		// Handle the case when there is no match[6]
		return "", fmt.Errorf("not a valid ISNI: %v", isni)
	}
	return isni, nil
}
