package models

import (
	"time"

	"github.com/google/uuid"
)

type Contributor struct {
	UUID             uuid.UUID         `json:"uuid" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" example:"5f410ec2-8eb8-4afd-b1f1-5a76114cc53e"`
	CreatedAt        time.Time         `json:"created_at" example:"2024-01-05T19:00:00.000000+01:00"`
	UpdatedAt        time.Time         `json:"updated_at" example:"2024-01-05T22:00:00.000000+01:00"`
	Affiliation      []Affiliation     `json:"affiliation,omitempty"`
	ContributorRoles []ContributorRole `json:"contributorRoles" gorm:"type:text[]"` // List of roles assumed by the contributor when working on the resource.
	FamilyName       *string           `json:"familyName,omitempty"`                // The family name of the contributor.
	GivenName        *string           `json:"givenName,omitempty"`                 // The given name of the contributor.
	ID               *string           `json:"id,omitempty"`                        // The unique identifier for the contributor.
	Name             *string           `json:"name,omitempty"`                      // The name of the contributor.
	Type             ContributorType   `json:"type"`                                // The type of the contributor.
	Resources        []*Resource       `json:"resources" gorm:"many2many:resource_contributors;"`
}

// The type of contribution made by a contributor
type ContributorRole string

const (
	Author                  ContributorRole = "Author"
	Chair                   ContributorRole = "Chair"
	Conceptualization       ContributorRole = "Conceptualization"
	ContactPerson           ContributorRole = "ContactPerson"
	ContributorRoleOther    ContributorRole = "Other"
	ContributorRoleSoftware ContributorRole = "Software"
	DataCuration            ContributorRole = "DataCuration"
	DataManager             ContributorRole = "DataManager"
	Distributor             ContributorRole = "Distributor"
	Editor                  ContributorRole = "Editor"
	FormalAnalysis          ContributorRole = "FormalAnalysis"
	FundingAcquisition      ContributorRole = "FundingAcquisition"
	HostingInstitution      ContributorRole = "HostingInstitution"
	Investigation           ContributorRole = "Investigation"
	Maintainer              ContributorRole = "Maintainer"
	Methodology             ContributorRole = "Methodology"
	Producer                ContributorRole = "Producer"
	ProjectAdministration   ContributorRole = "ProjectAdministration"
	ProjectLeader           ContributorRole = "ProjectLeader"
	ProjectManager          ContributorRole = "ProjectManager"
	ProjectMember           ContributorRole = "ProjectMember"
	Reader                  ContributorRole = "Reader"
	RegistrationAgency      ContributorRole = "RegistrationAgency"
	RegistrationAuthority   ContributorRole = "RegistrationAuthority"
	RelatedPerson           ContributorRole = "RelatedPerson"
	ResearchGroup           ContributorRole = "ResearchGroup"
	Researcher              ContributorRole = "Researcher"
	Resources               ContributorRole = "Resources"
	ReviewAssistant         ContributorRole = "ReviewAssistant"
	Reviewer                ContributorRole = "Reviewer"
	ReviewerExternal        ContributorRole = "ReviewerExternal"
	RightsHolder            ContributorRole = "RightsHolder"
	Sponsor                 ContributorRole = "Sponsor"
	StatsReviewer           ContributorRole = "StatsReviewer"
	Supervision             ContributorRole = "Supervision"
	Translator              ContributorRole = "Translator"
	Validation              ContributorRole = "Validation"
	Visualization           ContributorRole = "Visualization"
	WorkPackageLeader       ContributorRole = "WorkPackageLeader"
	WritingOriginalDraft    ContributorRole = "WritingOriginalDraft"
	WritingReviewEditing    ContributorRole = "WritingReviewEditing"
)
