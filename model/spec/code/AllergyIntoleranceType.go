package code

import (
	"fmt"
	"strings"
)

type AllergyIntoleranceType string

const (
	AllergyIntoleranceTypeAllergy     AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeIntolerance AllergyIntoleranceType = "intolerance"
)

func AllAllergyIntoleranceType() []AllergyIntoleranceType {
	return []AllergyIntoleranceType{
		AllergyIntoleranceTypeAllergy,
		AllergyIntoleranceTypeIntolerance,
	}
}

func FindAllergyIntoleranceType(filter string) []AllergyIntoleranceType {
	ret := make([]AllergyIntoleranceType, 0)
	for _, at := range AllAllergyIntoleranceType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllergyIntoleranceType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllergyIntoleranceType) String() string {
	switch cpt {
	case AllergyIntoleranceTypeAllergy:
		return "Allergy"
	case AllergyIntoleranceTypeIntolerance:
		return "Intolerance"
	default:
		return "Unknown Allergy Intolerance Type"
	}
}

/*
Allergy
Intolerance
*/
