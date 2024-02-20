package commonmeta

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// The dates for the resource.
type Date struct {
	UUID      uuid.UUID `json:"uuid"         gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at"                                                         example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at"                                                         example:"2024-01-05T22:00:00.000000+01:00"`
	// attributes
	Accepted   string `json:"accepted,omitempty"`  // The date the resource was accepted.
	Available  string `json:"available,omitempty"` // The date the resource was made available.
	Created    string `json:"created,omitempty"`   // The date the resource was created.
	Published  string `json:"published,omitempty"` // The date the resource was published.
	Submitted  string `json:"submitted,omitempty"` // The date the resource was submitted.
	Updated    string `json:"updated,omitempty"`   // The date the resource was updated.
	Withdrawn  string `json:"withdrawn,omitempty"` // The date the resource was withdrawn.
	ResourceID uuid.UUID  `json:"-"`
}

type DateJSON struct {
	Accepted   string `json:"accepted,omitempty"`  // The date the resource was accepted.
	Available  string `json:"available,omitempty"` // The date the resource was made available.
	Created    string `json:"created,omitempty"`   // The date the resource was created.
	Published  string `json:"published,omitempty"` // The date the resource was published.
	Submitted  string `json:"submitted,omitempty"` // The date the resource was submitted.
	Updated    string `json:"updated,omitempty"`   // The date the resource was updated.
	Withdrawn  string `json:"withdrawn,omitempty"` // The date the resource was withdrawn.
}

func (d *Date) ToJSONModel() DateJSON {
	dj := DateJSON{
		Accepted:    d.Accepted,
		Available: d.Available,
		Created: d.Created,
		Published: d.Published,
		Submitted: d.Submitted,
		Updated: d.Updated,
		Withdrawn: d.Withdrawn,
	}

	return dj
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToJSONModel())
}