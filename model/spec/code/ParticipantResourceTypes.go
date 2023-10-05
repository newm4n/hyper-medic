package code

import (
	"fmt"
	"strings"
)

type ParticipantResourceTypes string

const (
	ParticipantResourceTypesCareTeam          ParticipantResourceTypes = "CareTeam"
	ParticipantResourceTypesDevice            ParticipantResourceTypes = "Device"
	ParticipantResourceTypesGroup             ParticipantResourceTypes = "Group"
	ParticipantResourceTypesHealthcareService ParticipantResourceTypes = "HealthcareService"
	ParticipantResourceTypesLocation          ParticipantResourceTypes = "Location"
	ParticipantResourceTypesOrganization      ParticipantResourceTypes = "Organization"
	ParticipantResourceTypesPatient           ParticipantResourceTypes = "Patient"
	ParticipantResourceTypesPractitioner      ParticipantResourceTypes = "Practitioner"
	ParticipantResourceTypesPractitionerRole  ParticipantResourceTypes = "PractitionerRole"
	ParticipantResourceTypesRelatedPerson     ParticipantResourceTypes = "RelatedPerson"
)

func AllParticipantResourceTypes() []ParticipantResourceTypes {
	return []ParticipantResourceTypes{
		ParticipantResourceTypesCareTeam,
		ParticipantResourceTypesDevice,
		ParticipantResourceTypesGroup,
		ParticipantResourceTypesHealthcareService,
		ParticipantResourceTypesLocation,
		ParticipantResourceTypesOrganization,
		ParticipantResourceTypesPatient,
		ParticipantResourceTypesPractitioner,
		ParticipantResourceTypesPractitionerRole,
		ParticipantResourceTypesRelatedPerson,
	}
}

func FindParticipantResourceTypes(filter string) []ParticipantResourceTypes {
	ret := make([]ParticipantResourceTypes, 0)
	for _, at := range AllParticipantResourceTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ParticipantResourceTypes) ToString() {
	fmt.Println(au.String())
}
func (au ParticipantResourceTypes) String() string {
	switch au {
	case ParticipantResourceTypesCareTeam:
		return "CareTeam"
	case ParticipantResourceTypesDevice:
		return "Device"
	case ParticipantResourceTypesGroup:
		return "Group"
	case ParticipantResourceTypesHealthcareService:
		return "HealthcareService"
	case ParticipantResourceTypesLocation:
		return "Location"
	case ParticipantResourceTypesOrganization:
		return "Organization"
	case ParticipantResourceTypesPatient:
		return "Patient"
	case ParticipantResourceTypesPractitioner:
		return "Practitioner"
	case ParticipantResourceTypesPractitionerRole:
		return "PractitionerRole"
	case ParticipantResourceTypesRelatedPerson:
		return "RelatedPerson"
	default:
		return "Unknown Participant Resource Types"
	}
}
