package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Contributor struct {
	UUID             uuid.UUID         `json:"uuid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt        time.Time         `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt        time.Time         `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	Affiliation      []Affiliation     `json:"affiliation,omitempty"`
	ContributorRoles []ContributorRole `json:"contributorRoles" gorm:"foreignKey:ContributorUUID"` // List of roles assumed by the contributor when working on the resource.
	FamilyName       *string           `json:"familyName,omitempty"`                               // The family name of the contributor.
	GivenName        *string           `json:"givenName,omitempty"`                                // The given name of the contributor.
	ID               *string           `json:"id,omitempty"`                                       // The unique identifier for the contributor.
	Name             *string           `json:"name,omitempty"`                                     // The name of the contributor.
	Type             ContributorType   `json:"type"`                                               // The type of the contributor.
	Resources        []*Resource       `json:"resources" gorm:"many2many:resource2contributors;"`
}

type ContributorJSON struct {
	Affiliation      []Affiliation `json:"affiliation,omitempty"`
	ContributorRoles []string      `json:"contributorRoles,omitempty"`
	FamilyName       *string       `json:"familyName,omitempty"`
	GivenName        *string       `json:"givenName,omitempty"`
	ID               *string       `json:"id,omitempty"`
	Name             *string       `json:"name,omitempty"`
	Type             string        `json:"type,omitempty"`
}

// The type of contribution made by a contributor
type ContributorRole struct {
	UUID            uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt       time.Time `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt       time.Time `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	ContributorUUID uuid.UUID `json:"contributor_id"`
	Role            Role      `json:"role"`
}

type Role string

func (c *Contributor) MarshalJSON() ([]byte, error) {
	contributorJSON := ContributorJSON{
		ID:               c.ID,
		Type:             string(c.Type),
		ContributorRoles: []string{},
		Name:             c.Name,
		FamilyName:       c.FamilyName,
		GivenName:        c.GivenName,
		// Affiliation: c.Affiliation,
	}

	for _, cr := range c.ContributorRoles {
		contributorJSON.ContributorRoles = append(contributorJSON.ContributorRoles, string(cr.Role))
	}
	return json.Marshal(contributorJSON)
}

func getContributorRole() []Role {
	return []Role{
		"Author",
		"Editor",
		"Chair",
		"Reviewer",
		"ReviewAssistant",
		"StatsReviewer",
		"ReviewerExternal",
		"Reader",
		"Translator",
		"ContactPerson",
		"DataManager",
		"Distributor",
		"HostingInstitution",
		"Producer",
		"ProjectLeader",
		"ProjectManager",
		"ProjectMember",
		"RegistrationAgency",
		"RegistrationAuthority",
		"RelatedPerson",
		"ResearchGroup",
		"RightsHolder",
		"Researcher",
		"Sponsor",
		"WorkPackageLeader",
		"Conceptualization",
		"DataCuration",
		"FormalAnalysis",
		"FundingAcquisition",
		"Investigation",
		"Methodology",
		"ProjectAdministration",
		"Resources",
		"Software",
		"Supervision",
		"Validation",
		"Visualization",
		"WritingOriginalDraft",
		"WritingReviewEditing",
		"Maintainer",
		"Other",
	}
}
