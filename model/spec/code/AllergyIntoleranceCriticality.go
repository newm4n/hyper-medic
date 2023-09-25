package code

import (
	"fmt"
	"strings"
)

type AllergyIntoleranceCriticality int

const (
	AllergyIntoleranceCriticalityLowRisk AllergyIntoleranceCriticality = iota
	AllergyIntoleranceCriticalityHighRisk
	AllergyIntoleranceCriticalityUnableToAssessRisk
)

func AllAllergyIntoleranceCriticality() []AllergyIntoleranceCriticality {
	return []AllergyIntoleranceCriticality{
		AllergyIntoleranceCriticalityLowRisk,
		AllergyIntoleranceCriticalityHighRisk,
		AllergyIntoleranceCriticalityUnableToAssessRisk,
	}
}

func FindAllAllergyIntoleranceVerification(filter string) []AllergyIntoleranceCriticality {
	ret := make([]AllergyIntoleranceCriticality, 0)
	for _, at := range AllAllergyIntoleranceCriticality() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllergyIntoleranceCriticality) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllergyIntoleranceCriticality) String() string {
	switch cpt {
	case AllergyIntoleranceCriticalityLowRisk:
		return "Low Risk"
	case AllergyIntoleranceCriticalityHighRisk:
		return "High Risk"
	case AllergyIntoleranceCriticalityUnableToAssessRisk:
		return "Unable to Assess Risk"
	default:
		return "Unknown Allergy Intolerance Criticality"
	}
}

/*
Low Risk
High Risk
Unable to Assess Risk
*/
