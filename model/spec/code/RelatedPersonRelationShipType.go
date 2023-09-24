package code

import (
	"fmt"
	"strings"
)

type RelatedPersonRelationShipType int

const (
	RelatedPersonRelationShipTypeWitness RelatedPersonRelationShipType = iota
	RelatedPersonRelationShipTypeNotaryPublic
	RelatedPersonRelationShipTypeEmergencyContact
	RelatedPersonRelationShipTypeNextOfKin
	RelatedPersonRelationShipTypeGuardian
	RelatedPersonRelationShipTypeDependent
	RelatedPersonRelationShipTypeEmployee
	RelatedPersonRelationShipTypeGuarantor
	RelatedPersonRelationShipTypeCaregiver
	RelatedPersonRelationShipTypeEmployer
	RelatedPersonRelationShipTypeOther
	RelatedPersonRelationShipTypeUnknown
	RelatedPersonRelationShipTypeInterpreter
	RelatedPersonRelationShipTypePowerOfAttorney
	RelatedPersonRelationShipTypeDurablePowerOfAttorney
	RelatedPersonRelationShipTypeHealthcarePowerOfAttorney
	RelatedPersonRelationShipTypeSpecialPowerOfAttorney
	RelatedPersonRelationShipTypeBillingContact
)

func AllRelatedPersonRelationShipType() []RelatedPersonRelationShipType {
	return []RelatedPersonRelationShipType{
		RelatedPersonRelationShipTypeWitness,
		RelatedPersonRelationShipTypeNotaryPublic,
		RelatedPersonRelationShipTypeEmergencyContact,
		RelatedPersonRelationShipTypeNextOfKin,
		RelatedPersonRelationShipTypeGuardian,
		RelatedPersonRelationShipTypeDependent,
		RelatedPersonRelationShipTypeEmployee,
		RelatedPersonRelationShipTypeGuarantor,
		RelatedPersonRelationShipTypeCaregiver,
		RelatedPersonRelationShipTypeEmployer,
		RelatedPersonRelationShipTypeOther,
		RelatedPersonRelationShipTypeUnknown,
		RelatedPersonRelationShipTypeInterpreter,
		RelatedPersonRelationShipTypePowerOfAttorney,
		RelatedPersonRelationShipTypeDurablePowerOfAttorney,
		RelatedPersonRelationShipTypeHealthcarePowerOfAttorney,
		RelatedPersonRelationShipTypeSpecialPowerOfAttorney,
		RelatedPersonRelationShipTypeBillingContact,
	}
}

func FindRelatedPersonRelationShipType(filter string) []RelatedPersonRelationShipType {
	ret := make([]RelatedPersonRelationShipType, 0)
	for _, at := range AllRelatedPersonRelationShipType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt RelatedPersonRelationShipType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt RelatedPersonRelationShipType) String() string {
	switch cpt {
	case RelatedPersonRelationShipTypeWitness:
		return "witness"
	case RelatedPersonRelationShipTypeNotaryPublic:
		return "notary public"
	case RelatedPersonRelationShipTypeEmergencyContact:
		return "emergency contact"
	case RelatedPersonRelationShipTypeNextOfKin:
		return "next of kin"
	case RelatedPersonRelationShipTypeGuardian:
		return "guardian"
	case RelatedPersonRelationShipTypeDependent:
		return "dependent"
	case RelatedPersonRelationShipTypeEmployee:
		return "employee"
	case RelatedPersonRelationShipTypeGuarantor:
		return "guarantor"
	case RelatedPersonRelationShipTypeCaregiver:
		return "caregiver"
	case RelatedPersonRelationShipTypeEmployer:
		return "Employer"
	case RelatedPersonRelationShipTypeOther:
		return "Other"
	case RelatedPersonRelationShipTypeUnknown:
		return "Unknown"
	case RelatedPersonRelationShipTypeInterpreter:
		return "interpreter"
	case RelatedPersonRelationShipTypePowerOfAttorney:
		return "power of attorney"
	case RelatedPersonRelationShipTypeDurablePowerOfAttorney:
		return "durable power of attorney"
	case RelatedPersonRelationShipTypeHealthcarePowerOfAttorney:
		return "healthcare power of attorney"
	case RelatedPersonRelationShipTypeSpecialPowerOfAttorney:
		return "special power of attorney"
	case RelatedPersonRelationShipTypeBillingContact:
		return "Billing Contact"

	default:
		return "Unknown UDI Device Status"
	}
}

/*
witness
notary public
emergency contact
next of kin
guardian
dependent
employee
guarantor
caregiver
Employer
Other
Unknown
interpreter
power of attorney
durable power of attorney
healthcare power of attorney
special power of attorney
Billing Contact

*/
