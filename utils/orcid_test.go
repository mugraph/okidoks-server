package utils

import "testing"

func TestValidateORCID(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{
			"http://orcid.org/0000-0002-2590-225X should be 0000-0002-2590-225X",
			"http://orcid.org/0000-0002-2590-225X",
			"0000-0002-2590-225X",
		},
		{
			"https://orcid.org/0000-0002-2590-225X should be 0000-0002-2590-225X",
			"https://orcid.org/0000-0002-2590-225X",
			"0000-0002-2590-225X",
		},
		{
			"0000-0002-2590-225X should be 0000-0002-2590-225X",
			"0000-0002-2590-225X",
			"0000-0002-2590-225X",
		},
		{
			"https://www.orcid.org/0000-0002-2590-225X should be 0000-0002-2590-225X",
			"https://www.orcid.org/0000-0002-2590-225X",
			"0000-0002-2590-225X",
		},
		{
			"0000 0002 1394 3097 should be 0000-0002-1394-3097",
			"0000 0002 1394 3097",
			"0000-0002-1394-3097",
		},
		{
			"http://sandbox.orcid.org/0000-0002-2590-225X should be 0000-0002-2590-225X",
			"http://sandbox.orcid.org/0000-0002-2590-225X",
			"0000-0002-2590-225X",
		},
		{
			"https://sandbox.orcid.org/0000-0002-2590-225X should be 0000-0002-2590-225X",
			"https://sandbox.orcid.org/0000-0002-2590-225X",
			"0000-0002-2590-225X",
		},
		{
			"0000-0002-1394-309 should be empty",
			"0000-0002-1394-309",
			"",
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := ValidateORCID(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
