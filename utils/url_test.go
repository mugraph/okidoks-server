package utils

import "testing"

func TestValidateURL(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{
			"https://doi.org/10.5438/0000-00ss should be DOI",
			"https://doi.org/10.5438/0000-00ss",
			"DOI",
		},
		{
			"https://blog.datacite.org/eating-your-own-dog-food should be URL",
			"https://blog.datacite.org/eating-your-own-dog-food",
			"URL",
		},
		{
			"ISSN 2050-084X should be ISSN",
			"ISSN 2050-084X",
			"ISSN",
		},
		{
			"eating-your-own-dog-food should be empty",
			"eating-your-own-dog-food",
			"",
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ValidateURL(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
