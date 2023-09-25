package code

import (
	"fmt"
	"strings"
)

type ActionParticipantType int

const (
	ActionParticipantTypeCareTeam ActionParticipantType = iota
	ActionParticipantTypeDevice
	ActionParticipantTypeGroup
	ActionParticipantTypeHealthcareService
	ActionParticipantTypeLocation
	ActionParticipantTypeOrganization
	ActionParticipantTypePatient
	ActionParticipantTypePractitioner
	ActionParticipantTypePractitionerRole
	ActionParticipantTypeRelatedPerson
)

func AllActionParticipantType() []ActionParticipantType {
	return []ActionParticipantType{
		ActionParticipantTypeCareTeam,
		ActionParticipantTypeDevice,
		ActionParticipantTypeGroup,
		ActionParticipantTypeHealthcareService,
		ActionParticipantTypeLocation,
		ActionParticipantTypeOrganization,
		ActionParticipantTypePatient,
		ActionParticipantTypePractitioner,
		ActionParticipantTypePractitionerRole,
		ActionParticipantTypeRelatedPerson,
	}
}

func FindActionParticipantType(filter string) []ActionParticipantType {
	ret := make([]ActionParticipantType, 0)
	for _, at := range AllActionParticipantType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionParticipantType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionParticipantType) String() string {
	switch cpt {
	case ActionParticipantTypeCareTeam:
		return "CareTeam"
	case ActionParticipantTypeDevice:
		return "Device"
	case ActionParticipantTypeGroup:
		return "Group"
	case ActionParticipantTypeHealthcareService:
		return "HealthcareService"
	case ActionParticipantTypeLocation:
		return "Location"
	case ActionParticipantTypeOrganization:
		return "Organization"
	case ActionParticipantTypePatient:
		return "Patient"
	case ActionParticipantTypePractitioner:
		return "Practitioner"
	case ActionParticipantTypePractitionerRole:
		return "PractitionerRole"
	case ActionParticipantTypeRelatedPerson:
		return "RelatedPerson"
	default:
		return "Unknown Action Participant Type"
	}
}

/*
CareTeam
Device
Group
HealthcareService
Location
Organization
Patient
Practitioner
PractitionerRole
RelatedPerson
*/
