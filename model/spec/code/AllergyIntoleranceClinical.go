package code

import (
	"fmt"
	"strings"
)

type AllergyIntoleranceClinical string

const (
	AllergyIntoleranceClinicalActive   AllergyIntoleranceClinical = "active"
	AllergyIntoleranceClinicalInactive AllergyIntoleranceClinical = "inactive"
	AllergyIntoleranceClinicalResolved AllergyIntoleranceClinical = "resolved"
)

func AllAllergyIntoleranceClinical() []AllergyIntoleranceClinical {
	return []AllergyIntoleranceClinical{
		AllergyIntoleranceClinicalActive,
		AllergyIntoleranceClinicalInactive,
		AllergyIntoleranceClinicalResolved,
	}
}

func FindAllergyIntoleranceClinical(filter string) []AllergyIntoleranceClinical {
	ret := make([]AllergyIntoleranceClinical, 0)
	for _, at := range AllAllergyIntoleranceClinical() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllergyIntoleranceClinical) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllergyIntoleranceClinical) String() string {
	switch cpt {
	case AllergyIntoleranceClinicalActive:
		return "Active"
	case AllergyIntoleranceClinicalInactive:
		return "Inactive"
	case AllergyIntoleranceClinicalResolved:
		return "Resolved"
	default:
		return "Unknown Allergy Intolerance Clinical"
	}
}

/*
Active
Inactive
Resolved
*/
