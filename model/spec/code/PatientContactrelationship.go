package code

import (
	"fmt"
	"strings"
)

type PatientContactrelationship string

const (
	PatientContactrelationshipBilling           PatientContactrelationship = "BP"
	PatientContactrelationshipContact           PatientContactrelationship = "CP"
	PatientContactrelationshipEmergencyPerson   PatientContactrelationship = "EF"
	PatientContactrelationshipPreparingReferral PatientContactrelationship = "PR"
	PatientContactrelationshipEmployer          PatientContactrelationship = "E"
	PatientContactrelationshipEmergency         PatientContactrelationship = "C"
	PatientContactrelationshipFederalAgency     PatientContactrelationship = "F"
	PatientContactrelationshipInsurance         PatientContactrelationship = "I"
	PatientContactrelationshipNextOfKin         PatientContactrelationship = "N"
	PatientContactrelationshipStateAgency       PatientContactrelationship = "S"
	PatientContactrelationshipUnknown           PatientContactrelationship = "U"
)

func AllPatientContactrelationship() []PatientContactrelationship {
	return []PatientContactrelationship{
		PatientContactrelationshipBilling,
		PatientContactrelationshipContact,
		PatientContactrelationshipEmergencyPerson,
		PatientContactrelationshipPreparingReferral,
		PatientContactrelationshipEmployer,
		PatientContactrelationshipEmergency,
		PatientContactrelationshipFederalAgency,
		PatientContactrelationshipInsurance,
		PatientContactrelationshipNextOfKin,
		PatientContactrelationshipStateAgency,
		PatientContactrelationshipUnknown,
	}
}

func FindPatientContactrelationship(filter string) []PatientContactrelationship {
	ret := make([]PatientContactrelationship, 0)
	for _, at := range AllPatientContactrelationship() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au PatientContactrelationship) ToString() {
	fmt.Println(au.String())
}
func (au PatientContactrelationship) String() string {
	switch au {
	case PatientContactrelationshipBilling:
		return "Billing contact person"
	case PatientContactrelationshipContact:
		return "Contact person"
	case PatientContactrelationshipEmergencyPerson:
		return "Emergency contact person"
	case PatientContactrelationshipPreparingReferral:
		return "Person preparing referral"
	case PatientContactrelationshipEmployer:
		return "Employer"
	case PatientContactrelationshipEmergency:
		return "Emergency Contact"
	case PatientContactrelationshipFederalAgency:
		return "Federal Agency"
	case PatientContactrelationshipInsurance:
		return "Insurance Company"
	case PatientContactrelationshipNextOfKin:
		return "Next-of-Kin"
	case PatientContactrelationshipStateAgency:
		return "State Agency"
	case PatientContactrelationshipUnknown:
		return "Unknown"
	default:
		return "Unknown Constraint Severity"
	}
}
