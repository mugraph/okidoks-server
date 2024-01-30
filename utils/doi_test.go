package utils

import "testing"

func TestGetDOIRA(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{"10.1371/journal.pone.0042793 should be Crossref", "10.1371/journal.pone.0042793", "Crossref"},
		{"https://doi.org/10.5061/dryad.8515 should be DataCite", "https://doi.org/10.5061/dryad.8515", "DataCite"},
		{"https://doi.org/10.1392/roma081203 should be mEDRA", "https://doi.org/10.1392/roma081203", "mEDRA"},
		{"https://doi.org/10.5012/bkcs.2013.34.10.2889 should be KISTI", "https://doi.org/10.5012/bkcs.2013.34.10.2889", "KISTI"},
		{"https://doi.org/10.11367/grsj1979.12.283 should be JaLC", "https://doi.org/10.11367/grsj1979.12.283", "JaLC"},
		{"https://doi.org/10.2903/j.efsa.2018.5239 should be OP", "https://doi.org/10.2903/j.efsa.2018.5239", "OP"},
		// Not a valid prefix
		{"https://doi.org/10.a/dryad.8515x should be ''", "https://doi.org/10.a/dryad.8515x", ""},
		// Not found
		{"https://doi.org/10.99999/dryad.8515x should be ''", "https://doi.org/10.99999/dryad.8515x", ""},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := GetDOIRA(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestValidateDOI(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{"10.1371/journal.pone.0042793 should be 10.1371/journal.pone.0042793", "10.1371/journal.pone.0042793", "10.1371/journal.pone.0042793"},
		{"https://doi.org/10.1371/journal.pone.0042793 should be 10.1371/journal.pone.0042793", "https://doi.org/10.1371/journal.pone.0042793", "10.1371/journal.pone.0042793"},
		{"journal.pone.0042793 should be ''", "journal.pone.0042793", ""},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := ValidateDOI(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestValidaterePrefix(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{"10.1371/journal.pone.0042793 should be 10.1371", "10.1371/journal.pone.0042793", "10.1371"},
		{"doi:10.1371/journal.pone.0042793 should be 10.1371", "doi:10.1371/journal.pone.0042793", "10.1371"},
		{"http://doi.org/10.1371/journal.pone.0042793 should be 10.1371", "http://doi.org/10.1371/journal.pone.0042793", "10.1371"},
		{"10.1371 should be 10.1371", "10.1371", "10.1371"},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := ValidatePrefix(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestDOIfromURL(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{"https://doi.org/10.5061/dryad.8515 should be 10.5071/dryad.8515", "https://doi.org/10.5061/dryad.8515", "10.5061/dryad.8515"},
		{"10.5061/dryad.8515 should be 10.5071/dryad.8515", "10.5061/dryad.8515", "10.5061/dryad.8515"},
		{"10.5067/terra+aqua/ceres/cldtyphist_l3.004 should be 10.5067/terra+aqua/ceres/cldtyphist_l3.004", "10.5067/terra+aqua/ceres/cldtyphist_l3.004", "10.5067/terra+aqua/ceres/cldtyphist_l3.004"},
		{"doi:10.1371/journal.pone.0042793 should be 10.1371/journal.pone.0042793", "doi:10.1371/journal.pone.0042793", "10.1371/journal.pone.0042793"},
		{"https://handle.stage.datacite.org/10.5438/55e5-t5c0 should be 10.5438/55e5-t5c0", "https://handle.stage.datacite.org/10.5438/55e5-t5c0", "10.5438/55e5-t5c0"},
		{"https://doi.org/10.5061 should be ''", "https://doi.org/10.5061", ""},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := DOIFromURL(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestDOIAsURL(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The test cases
		{"10.5061/dryad.8515 should be https://doi.org/10.5061/dryad.8515", "10.5061/dryad.8515", "https://doi.org/10.5061/dryad.8515"},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := DOIAsURL(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
