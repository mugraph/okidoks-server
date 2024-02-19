package commonmeta

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/mugraph/okidoks-server/utils"
)

type Contributor struct {
	UUID      uuid.UUID `json:"uuid"                  gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at"                                                                  example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at"                                                                  example:"2024-01-05T22:00:00.000000+01:00"`
	// attributes
	Affiliation      []Affiliation     `json:"affiliation,omitempty"`
	ContributorRoles []ContributorRole `json:"contributorRoles"      gorm:"foreignKey:ContributorUUID"` // List of roles assumed by the contributor when working on the resource.
	FamilyName       *string           `json:"familyName,omitempty"`                                    // The family name of the contributor.
	GivenName        *string           `json:"givenName,omitempty"`                                     // The given name of the contributor.
	ID               *string           `json:"id,omitempty"`                                            // The unique identifier for the contributor.
	Name             *string           `json:"name,omitempty"`                                          // The name of the contributor.
	Type             ContributorType   `json:"type"`                                                    // The type of the contributor.
	Resources        []Resource        `json:"resources"             gorm:"many2many:resource2contributors;"`
}

// The type of contribution made by a contributor
type ContributorRole struct {
	UUID      uuid.UUID `json:"uuid"           gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt time.Time `json:"created_at"                                                           example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt time.Time `json:"updated_at"                                                           example:"2024-01-05T22:00:00.000000+01:00"`
	// attributes
	ContributorUUID uuid.UUID `json:"contributor_id"`
	Role            Role      `json:"role"`
}

type Role string

// The type of the contributor.
type ContributorType string

type ContributorJSON struct {
	Affiliation      []Affiliation `json:"affiliation,omitempty"`
	ContributorRoles []string      `json:"contributorRoles,omitempty"`
	FamilyName       *string       `json:"familyName,omitempty"`
	GivenName        *string       `json:"givenName,omitempty"`
	ID               *string       `json:"id,omitempty"`
	Name             *string       `json:"name,omitempty"`
	Type             *string       `json:"type,omitempty"`
}

func (c *Contributor) ToJSONModel() ContributorJSON {
	t := string(c.Type)
	var roles []string
	cj := ContributorJSON{
		ID:               c.ID,
		Type:             NilOrPtrToString(&t),
		ContributorRoles: roles,
		Name:             c.Name,
		FamilyName:       c.FamilyName,
		GivenName:        c.GivenName,
		// Affiliation: c.Affiliation,
	}

	for _, cr := range c.ContributorRoles {
		cj.ContributorRoles = append(cj.ContributorRoles, utils.ContributorRoleMap.GetVal(string(cr.Role), true, true))
	}
	return cj
}

func (c *Contributor) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToJSONModel())
}
