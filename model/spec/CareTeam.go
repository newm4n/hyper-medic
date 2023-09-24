package modelspec

import (
	"github.com/newm4n/hyper-medic/model/spec/code"
	"time"
)

type ICareTeam interface {
	IDomainResource
	GetIdentifiers() []string
	GetCareTeamStatus() code.CareTeamStatus
	GetCareTeamCateogy() code.CareTeamCategory
	GetName() string
	GetSubjectPatient() IPatient
	GetSubjectGroup() IGroup
	GetPeriod() IPeriod
	GetPractitionerParticipants() []ICareTeamPractitionerParticipant
	GetPractitionerRoleParticipants() []ICareTeamPractitionerRoleParticipant
	GetRelatedPersonParticipants() []ICareTeamRelatedPersonRoleParticipant
	GetPatientParticipants() []ICareTeamPetientParticipant
	GetOrganizationParticipants() []ICareTeamOrganizationParticipant
	GetOtherTeamParticipants() []ICareTeamOtherTeamParticipant
	GetReason() code.ClinicalFindings
	GetManagingOrganization() IOrganization
	GetTelecom() []IContactPoint
	Notes() []string
}

type ICareTeamPractitionerParticipant interface {
	GetCareTeam() ICareTeam
	GetRole() string
	GetPractitionerParticipant() IPractitioner
	GetOnBehalfOf() IOrganization
	GetCoverageTime() time.Time
	GetCoveragePriod() IPeriod
}

type ICareTeamPractitionerRoleParticipant interface {
	GetCareTeam() ICareTeam
	GetRole() string
	GetPractitionerRoleParticipant() IPractitionerRole
	GetOnBehalfOf() IOrganization
	GetCoverageTime() time.Time
	GetCoveragePriod() IPeriod
}

type ICareTeamRelatedPersonRoleParticipant interface {
	GetCareTeam() ICareTeam
	GetRole() string
	GetRelatedPersonParticipant() IRelatedPerson
	GetOnBehalfOf() IOrganization
	GetCoverageTime() time.Time
	GetCoveragePriod() IPeriod
}

type ICareTeamPetientParticipant interface {
	GetCareTeam() ICareTeam
	GetRole() string
	GetPatientParticipant() IPatient
	GetOnBehalfOf() IOrganization
	GetCoverageTime() time.Time
	GetCoveragePriod() IPeriod
}

type ICareTeamOrganizationParticipant interface {
	GetCareTeam() ICareTeam
	GetRole() string
	GetOrganizationParticipant() IOrganization
	GetOnBehalfOf() IOrganization
	GetCoverageTime() time.Time
	GetCoveragePriod() IPeriod
}

type ICareTeamOtherTeamParticipant interface {
	GetCareTeam() ICareTeam
	GetRole() string
	GetOtherTeamParticipant() ICareTeam
	GetOnBehalfOf() IOrganization
	GetCoverageTime() time.Time
	GetCoveragePriod() IPeriod
}
