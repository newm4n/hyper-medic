package code

import (
	"fmt"
	"strings"
)

type ContactRelationship int

const (
	ContactRelationshipBilling ContactRelationship = iota
	ContactRelationshipContact
	ContactRelationshipEmergencyPerson
	ContactRelationshipPreparing
	ContactRelationshipEmployer
	ContactRelationshipEmergency
	ContactRelationshipFederalAgency
	ContactRelationshipInsurance
	ContactRelationshipNextOfKin
	ContactRelationshipStateAgency
	ContactRelationshipUnknown
)

func AllContactRelationships() []ContactRelationship {
	return []ContactRelationship{
		ContactRelationshipBilling,
		ContactRelationshipContact,
		ContactRelationshipEmergencyPerson,
		ContactRelationshipPreparing,
		ContactRelationshipEmployer,
		ContactRelationshipEmergency,
		ContactRelationshipFederalAgency,
		ContactRelationshipInsurance,
		ContactRelationshipNextOfKin,
		ContactRelationshipStateAgency,
		ContactRelationshipUnknown,
	}
}

func FindContactRelationships(filter string) []ContactRelationship {
	ret := make([]ContactRelationship, 0)
	for _, at := range AllContactRelationships() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (s ContactRelationship) ToString() {
	fmt.Println(s.String())
}
func (s ContactRelationship) String() string {
	switch s {
	case ContactRelationshipBilling:
		return "Billing Contact"
	case ContactRelationshipContact:
		return "Contact"
	case ContactRelationshipEmergencyPerson:
		return "Emergency Person"
	case ContactRelationshipPreparing:
		return "Perparing Person"
	case ContactRelationshipEmployer:
		return "Employer"
	case ContactRelationshipEmergency:
		return "Emergency"
	case ContactRelationshipFederalAgency:
		return "Federal Agency"
	case ContactRelationshipInsurance:
		return "Insurance"
	case ContactRelationshipNextOfKin:
		return "Next of Kin"
	case ContactRelationshipStateAgency:
		return "State Agency"
	case ContactRelationshipUnknown:
		return "Unknown"
	default:
		return "Unknown Contact Relation Ship"
	}
}
