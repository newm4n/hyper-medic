package code

import (
	"fmt"
	"strings"
)

const (
	OrganizationTypeProvider OrganizationType = iota
	OrganizationTypeHospitalDept
	OrganizationTypeOrganizationalTeam
	OrganizationTypeGovernment
	OrganizationTypeInsuranceCompany
	OrganizationTypePayer
	OrganizationTypeEducationalInstitute
	OrganizationTypeReligiousInstitute
	OrganizationTypeClinicalResearch
	OrganizationTypeCommunityGroup
	OrganizationTypeBusinessCorp
	OrganizationTypeOther
)

func AllOrganizationTypes() []OrganizationType {
	return []OrganizationType{
		OrganizationTypeProvider,
		OrganizationTypeHospitalDept,
		OrganizationTypeOrganizationalTeam,
		OrganizationTypeGovernment,
		OrganizationTypeInsuranceCompany,
		OrganizationTypePayer,
		OrganizationTypeEducationalInstitute,
		OrganizationTypeReligiousInstitute,
		OrganizationTypeClinicalResearch,
		OrganizationTypeCommunityGroup,
		OrganizationTypeBusinessCorp,
		OrganizationTypeOther,
	}
}

func FindOrganizationType(filter string) []OrganizationType {
	ret := make([]OrganizationType, 0)
	for _, at := range AllOrganizationTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

type OrganizationType int

func (ot OrganizationType) ToString() {
	fmt.Println(ot.String())
}
func (ot OrganizationType) String() string {
	switch ot {
	case OrganizationTypeProvider:
		return "Provider"
	case OrganizationTypeHospitalDept:
		return "Hospital Department"
	case OrganizationTypeOrganizationalTeam:
		return "Organizational Team"
	case OrganizationTypeGovernment:
		return "Government"
	case OrganizationTypeInsuranceCompany:
		return "Insurance Company"
	case OrganizationTypePayer:
		return "Payer"
	case OrganizationTypeEducationalInstitute:
		return "Education institute"
	case OrganizationTypeReligiousInstitute:
		return "Religious Institution"
	case OrganizationTypeClinicalResearch:
		return "Clinical Research"
	case OrganizationTypeCommunityGroup:
		return "Community Group"
	case OrganizationTypeBusinessCorp:
		return "Business Corporations"
	case OrganizationTypeOther:
		return "Other Organization"
	default:
		return "Unknown Organization Type"
	}
}
