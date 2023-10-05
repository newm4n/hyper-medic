package code

import (
	"fmt"
	"strings"
)

type RelatedpersonRelationshipType string

const (
	RelatedpersonRelationshipTypeWIT       RelatedpersonRelationshipType = "WIT"
	RelatedpersonRelationshipTypeNOT       RelatedpersonRelationshipType = "NOT"
	RelatedpersonRelationshipTypeECON      RelatedpersonRelationshipType = "ECON"
	RelatedpersonRelationshipTypeNOK       RelatedpersonRelationshipType = "NOK"
	RelatedpersonRelationshipTypeGUARD     RelatedpersonRelationshipType = "GUARD"
	RelatedpersonRelationshipTypeDEPEN     RelatedpersonRelationshipType = "DEPEN"
	RelatedpersonRelationshipTypeEMP       RelatedpersonRelationshipType = "EMP"
	RelatedpersonRelationshipTypeGUAR      RelatedpersonRelationshipType = "GUAR"
	RelatedpersonRelationshipTypeCAREGIVER RelatedpersonRelationshipType = "CAREGIVER"
	RelatedpersonRelationshipTypeE         RelatedpersonRelationshipType = "E"
	RelatedpersonRelationshipTypeO         RelatedpersonRelationshipType = "O"
	RelatedpersonRelationshipTypeU         RelatedpersonRelationshipType = "U"
	RelatedpersonRelationshipTypeINTPRTER  RelatedpersonRelationshipType = "INTPRTER"
	RelatedpersonRelationshipTypePOWATT    RelatedpersonRelationshipType = "POWATT"
	RelatedpersonRelationshipTypeDPOWATT   RelatedpersonRelationshipType = "DPOWATT"
	RelatedpersonRelationshipTypeHPOWATT   RelatedpersonRelationshipType = "HPOWATT"
	RelatedpersonRelationshipTypeSPOWATT   RelatedpersonRelationshipType = "SPOWATT"
	RelatedpersonRelationshipTypeBILL      RelatedpersonRelationshipType = "BILL"
)

func AllRelatedpersonRelationshipType() []RelatedpersonRelationshipType {
	return []RelatedpersonRelationshipType{
		RelatedpersonRelationshipTypeWIT,
		RelatedpersonRelationshipTypeNOT,
		RelatedpersonRelationshipTypeECON,
		RelatedpersonRelationshipTypeNOK,
		RelatedpersonRelationshipTypeGUARD,
		RelatedpersonRelationshipTypeDEPEN,
		RelatedpersonRelationshipTypeEMP,
		RelatedpersonRelationshipTypeGUAR,
		RelatedpersonRelationshipTypeCAREGIVER,
		RelatedpersonRelationshipTypeE,
		RelatedpersonRelationshipTypeO,
		RelatedpersonRelationshipTypeU,
		RelatedpersonRelationshipTypeINTPRTER,
		RelatedpersonRelationshipTypePOWATT,
		RelatedpersonRelationshipTypeDPOWATT,
		RelatedpersonRelationshipTypeHPOWATT,
		RelatedpersonRelationshipTypeSPOWATT,
		RelatedpersonRelationshipTypeBILL,
	}
}

func FindRelatedpersonRelationshipType(filter string) []RelatedpersonRelationshipType {
	ret := make([]RelatedpersonRelationshipType, 0)
	for _, at := range AllRelatedpersonRelationshipType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au RelatedpersonRelationshipType) ToString() {
	fmt.Println(au.String())
}
func (au RelatedpersonRelationshipType) String() string {
	switch au {
	case RelatedpersonRelationshipTypeWIT:
		return "witness"
	case RelatedpersonRelationshipTypeNOT:
		return "notary public"
	case RelatedpersonRelationshipTypeECON:
		return "emergency contact"
	case RelatedpersonRelationshipTypeNOK:
		return "next of kin"
	case RelatedpersonRelationshipTypeGUARD:
		return "guardian"
	case RelatedpersonRelationshipTypeDEPEN:
		return "dependent"
	case RelatedpersonRelationshipTypeEMP:
		return "employee"
	case RelatedpersonRelationshipTypeGUAR:
		return "guarantor"
	case RelatedpersonRelationshipTypeCAREGIVER:
		return "caregiver"
	case RelatedpersonRelationshipTypeE:
		return "Employer"
	case RelatedpersonRelationshipTypeO:
		return "Other"
	case RelatedpersonRelationshipTypeU:
		return "Unknown"
	case RelatedpersonRelationshipTypeINTPRTER:
		return "interpreter"
	case RelatedpersonRelationshipTypePOWATT:
		return "power of attorney"
	case RelatedpersonRelationshipTypeDPOWATT:
		return "durable power of attorney"
	case RelatedpersonRelationshipTypeHPOWATT:
		return "healthcare power of attorney"
	case RelatedpersonRelationshipTypeSPOWATT:
		return "special power of attorney"
	case RelatedpersonRelationshipTypeBILL:
		return "Billing Contact"
	default:
		return "Unknown Constraint Severity"
	}
}
