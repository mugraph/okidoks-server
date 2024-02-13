// Version: v0.11
// URL:     https://commonmeta.org/commonmeta_v0.11.json
package commonmeta

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// The type of the title.
type TitleType string

type Title struct {
	UUID      uuid.UUID `json:"uuid"               gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at"                                                               example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at"                                                               example:"2024-01-05T22:00:00.000000+01:00"`
	// attributes
	Title      string     `json:"title"`              // The title of the resource.
	Type       *TitleType `json:"type,omitempty"`     // The type of the title.
	Language   string     `json:"language,omitempty"` // The language of the title. Use one of the language codes from the IETF BCP 47 standard.
	ResourceID string     `json:"-"`
}

type TitleJSON struct {
	Title    string `json:"title"`              // The title of the resource.
	Type     string `json:"type,omitempty"`     // The type of the title.
	Language string `json:"language,omitempty"` // The language of the title. Use one of the language codes from the IETF BCP 47 standard.
}

func (t *Title) ToJSONModel() TitleJSON {
	tj := TitleJSON{
		Title:    t.Title,
		Type:     string(*t.Type),
		Language: t.Language,
	}

	return tj
}

func (t *Title) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToJSONModel())
}
