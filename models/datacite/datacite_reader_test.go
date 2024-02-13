package datacite

import (
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
			name: "Type should be Article. AdditionalType should be Article",
			input: Types{
				ResourceType: "Article",
				ResourceTypeGeneral: "Preprint",
			},
			want1: commonmeta.ResourceType("Article"),
			want2: commonmeta.ResourceType("Article"),
		},
		{
			name: "Type should be JournalArticle. AdditionalType should be empty",
			input: Types{
				ResourceType: "JournalArticle",
				ResourceTypeGeneral: "Text",
			},
			want1: commonmeta.ResourceType("JournalArticle"),
			want2: commonmeta.ResourceType(""),
		},
		{
			name: "Type should be Document. AdditionalType should be Master Thesis",
			input: Types{
				ResourceType: "Master Thesis",
				ResourceTypeGeneral: "Text",
			},
			want1: commonmeta.ResourceType("Document"),
			want2: commonmeta.ResourceType("Master Thesis"),
		},
		{
			name: "Type should be Other. AdditionalType should be empty",
			input: Types{
				ResourceType: "",
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