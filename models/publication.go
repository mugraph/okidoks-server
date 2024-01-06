// models/publication.go

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Publication struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt  time.Time      `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt  time.Time      `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	Name       string         `json:"name"`
	Attributes datatypes.JSON `json:"attributes"`
}
