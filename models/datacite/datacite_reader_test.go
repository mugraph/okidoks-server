package datacite

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mugraph/okidoks-server/models/commonmeta"
)

func TestTypes(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input Types
		want1 commonmeta.ResourceType
		want2 commonmeta.ResourceType
	}{
		// The Test Cases
		{
			name: "Type should be Article. AdditionalType should be empty",
			input: Types{
				ResourceType:        "Article",
				ResourceTypeGeneral: "Preprint",
			},
			want1: commonmeta.ResourceType("Article"),
			want2: commonmeta.ResourceType(""),
		},
		{
			name: "Type should be JournalArticle. AdditionalType should be empty",
			input: Types{
				ResourceType:        "JournalArticle",
				ResourceTypeGeneral: "Text",
			},
			want1: commonmeta.ResourceType("JournalArticle"),
			want2: commonmeta.ResourceType(""),
		},
		// {
		// 	name: "Type should be Document. AdditionalType should be Master Thesis",
		// 	input: Types{
		// 		ResourceType:        "Master Thesis",
		// 		ResourceTypeGeneral: "Text",
		// 	},
		// 	want1: commonmeta.ResourceType("Document"),
		// 	want2: commonmeta.ResourceType("Master Thesis"),
		// },
		{
			name: "Type should be Other. AdditionalType should be empty",
			input: Types{
				ResourceType:        "",
				ResourceTypeGeneral: "",
			},
			want1: commonmeta.ResourceType("Other"),
			want2: commonmeta.ResourceType(""),
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans1, ans2 := types(tt.input)
			if ans1 != tt.want1 {
				t.Errorf("got %v, want %v", ans1, tt.want1)
			}
			if ans2 != tt.want2 {
				t.Errorf("got %v, want %v", ans2, tt.want2)
			}
		})
	}
}

func TestIsPersonalName(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		// The Test Cases
		{
			name:  "OrganisationWithoutSpaces should be false",
			input: "OrganisationWithoutSpaces",
			want:  false,
		},
		{
			name:  "OrganisationWith; Semicolon should be false",
			input: "OrganisationWith; Semicolon",
			want:  false,
		},
		{
			name:  "Some University should be false",
			input: "Some University",
			want:  false,
		},
		{
			name:  "Givenname Surname, PhD should be true",
			input: "Givenname Surname, PhD",
			want:  true,
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isPersonalName(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestCleanupName(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The Test Cases
		{
			name:  "John Smith is John Smith",
			input: "John Smith",
			want:  "John Smith",
		},
		{
			name:  "Smith, John should be Smith, John",
			input: "Smith, John",
			want:  "Smith, John",
		},
		{
			name:  "Smith, J. should be Smith, J.",
			input: "Smith, J.",
			want:  "Smith, J.",
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := cleanupName(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestAffiliations(t *testing.T) {
	// Define table columns
	var nilSlice []commonmeta.Affiliation
	var tests = []struct {
		name  string
		input []Affiliation
		want  []commonmeta.Affiliation
	}{
		// The Test Cases
		{
			name:  "Empty array returns empty",
			input: []Affiliation{},
			want:  nilSlice,
		},
	}
	// Execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := affiliations(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %#v, want %#v", ans, tt.want)
			}
		})
	}
}

func TestAffiliationsAgain(t *testing.T) {
	// Define table columns
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// The Test Cases
		{
			name:  "Empty array returns empty",
			input: `[]`,
			want:  ``,
		},
		{
			name:  "Name returns name",
			input: `[{"name": "University of California, Santa Barbara"}]`,
			want:  `[{"name":"University of California, Santa Barbara"}]`,
		},
		{
			name: "Complicated DataCite returns Name and ID",
			input: `[
						{
							"affiliationIdentifier": "02t274463",
							"affiliationIdentifierScheme": "ROR",
							"name": "University of California, Santa Barbara",
							"schemeURI": "https://ror.org/"
						}
					]`,
			want: `[
						{
							"id": "https://ror.org/02t274463",
							"name": "University of California, Santa Barbara"
						}
					]`,
		},
	}
	// Execution loop
	for _, tt := range tests {
		var in []Affiliation
		t.Run(tt.name, func(t *testing.T) {
			// Unmarshal input
			err := json.Unmarshal([]byte(tt.input), &in)
			if err != nil {
				t.Logf("could not unmarschal input: %v", tt.input)
			}

			ans := affiliations(in)

			out, err := json.Marshal(ans)
			if err != nil {
				t.Logf("could not marschal ans: %v", ans)
			}

			str := string(out)

			var d interface{}
			err = json.Unmarshal([]byte(tt.want), &d)
			if err != nil {
				t.Logf("could not unmarshal tt.want: %v", err)
			}

			want, err := json.Marshal(d)
			if err != nil {
				t.Logf("could not re-marshall unmarshalled tt.want: %v", err)
			}

			strWant := string(want)

			if str != strWant {
				t.Errorf("got %#v, want %#v", str, strWant)
			}
		})
	}
}
