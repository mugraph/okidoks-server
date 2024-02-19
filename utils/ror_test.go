package utils

import "testing"

func TestValidateROR(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{
			"https://ror.org/0342dzm54 should be 0342dzm54",
			"https://ror.org/0342dzm54",
			"0342dzm54",
		},
		{
			"ror.org/0343dzm54 should be 0342dzm54",
			"ror.org/0342dzm54",
			"0342dzm54",
		},
		{
			"ror.org/0342dzm should be empty",
			"ror.org/0342dzm",
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
			ans, _ := ValidateROR(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
