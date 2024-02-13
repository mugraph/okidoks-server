package commonmeta

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// DB model
// The license for the resource. Use one of the SPDX license identifiers.
type License struct {
	UUID      uuid.UUID `json:"uuid"          gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at"                                                          example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at"                                                          example:"2024-01-05T22:00:00.000000+01:00"`
	// attributes
	ID  *string `json:"id,omitempty"`
	URL string  `json:"url,omitempty" gorm:"uniqueIndex"`
}

// JSON model
type LicenseJSON struct {
	ID  *string `json:"id,omitempty"`
	URL string  `json:"url,omitempty"` // format uri
}

func (l *License) ToJSONModel() LicenseJSON {
	return LicenseJSON{
		ID:  l.ID,
		URL: l.URL,
	}
}

func (l *License) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToJSONModel())
}
