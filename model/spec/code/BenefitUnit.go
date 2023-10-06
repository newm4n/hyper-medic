package code

import (
	"fmt"
	"strings"
)

type BenefitUnit string

const (
	BenefitUnitIndividual BenefitUnit = "individual"
	BenefitUnitFamily     BenefitUnit = "family"
)

func AllBenefitUnit() []BenefitUnit {
	return []BenefitUnit{
		BenefitUnitIndividual,
		BenefitUnitFamily,
	}
}

func FindBenefitUnit(filter string) []BenefitUnit {
	ret := make([]BenefitUnit, 0)
	for _, at := range AllBenefitUnit() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BenefitUnit) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BenefitUnit) String() string {
	switch cpt {
	case BenefitUnitIndividual:
		return "Individual"
	case BenefitUnitFamily:
		return "Family"
	default:
		return "Unknown Benefit Unit"
	}
}
