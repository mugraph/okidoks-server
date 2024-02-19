package utils

import "testing"

func TestNormalizeCCURL(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{
			"https://creativecommons.org/licenses/by/4.0/ should be https://creativecommons.org/licenses/by/4.0/legalcode",
			"https://creativecommons.org/licenses/by/4.0/",
			"https://creativecommons.org/licenses/by/4.0/legalcode",
		},
		{
			"https://creativecommons.org/publicdomain/zero/1.0 should be https://creativecommons.org/publicdomain/zero/1.0/legalcode",
			"https://creativecommons.org/publicdomain/zero/1.0",
			"https://creativecommons.org/publicdomain/zero/1.0/legalcode",
		},
		{
			"http://creativecommons.org/publicdomain/zero/1.0 should be https://creativecommons.org/publicdomain/zero/1.0/legalcode",
			"http://creativecommons.org/publicdomain/zero/1.0",
			"https://creativecommons.org/publicdomain/zero/1.0/legalcode",
		},
		{
			"https://creativecommons.org/publicdomain/zero/1.0/legalcode should be https://creativecommons.org/publicdomain/zero/1.0/legalcode",
			"https://creativecommons.org/publicdomain/zero/1.0/legalcode",
			"https://creativecommons.org/publicdomain/zero/1.0/legalcode",
		},
		{
			"empty should be empty",
			"",
			"",
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NormalizeCCURL(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestNormalizeORCID(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{
			"http://orcid.org/0000-0002-2590-225X should be https://orcid.org/0000-0002-2590-225X",
			"http://orcid.org/0000-0002-2590-225X",
			"https://orcid.org/0000-0002-2590-225X",
		},
		{
			"https://orcid.org/0000-0002-2590-225X should be https://orcid.org/0000-0002-2590-225X",
			"https://orcid.org/0000-0002-2590-225X",
			"https://orcid.org/0000-0002-2590-225X",
		},
		{
			"0000-0002-2590-225X should be https://orcid.org/0000-0002-2590-225X",
			"0000-0002-2590-225X",
			"https://orcid.org/0000-0002-2590-225X",
		},
		{
			"0002-2590-225X should be empty",
			"0002-2590-225X",
			"",
		},
		{
			"empty should be empty",
			"",
			"",
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NormalizeORCID(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestNormalizeROR(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{
			"http://ror.org/0342dzm54 should be https://ror.org/0342dzm54",
			"http://ror.org/0342dzm54",
			"https://ror.org/0342dzm54",
		},
		{
			"empty should be empty",
			"",
			"",
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NormalizeROR(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
