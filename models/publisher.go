package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// DB Model
// The publisher of many resources.
type Publisher struct {
	UUID      uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	ID        string    `json:"id,omitempty"`            // The identifier for the publisher.
	Name      string    `json:"name" gorm:"uniqueIndex"` // The name of the publisher.
}

// JSON Model
type PublisherJSON struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (p *Publisher) ToJSONModel() PublisherJSON {
	return PublisherJSON{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p *Publisher) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.ToJSONModel())
}
