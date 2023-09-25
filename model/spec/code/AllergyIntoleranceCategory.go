package code

import (
	"fmt"
	"strings"
)

type AllergyIntoleranceCategory int

const (
	AllergyIntoleranceCategoryFood AllergyIntoleranceCategory = iota
	AllergyIntoleranceCategoryMedication
	AllergyIntoleranceCategoryEnvironment
	AllergyIntoleranceCategoryBiologic
)

func AllAllergyIntoleranceCategory() []AllergyIntoleranceCategory {
	return []AllergyIntoleranceCategory{
		AllergyIntoleranceCategoryFood,
		AllergyIntoleranceCategoryMedication,
		AllergyIntoleranceCategoryEnvironment,
		AllergyIntoleranceCategoryBiologic,
	}
}

func FindAllergyIntoleranceCategory(filter string) []AllergyIntoleranceCategory {
	ret := make([]AllergyIntoleranceCategory, 0)
	for _, at := range AllAllergyIntoleranceCategory() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllergyIntoleranceCategory) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllergyIntoleranceCategory) String() string {
	switch cpt {
	case AllergyIntoleranceCategoryFood:
		return "Food"
	case AllergyIntoleranceCategoryMedication:
		return "Medication"
	case AllergyIntoleranceCategoryEnvironment:
		return "Environment"
	case AllergyIntoleranceCategoryBiologic:
		return "Biologic"
	default:
		return "Unknown Allergy Intolerance Clinical"
	}
}

/*
Food
Medication
Environment
Biologic
*/
