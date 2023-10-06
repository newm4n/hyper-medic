package code

import (
	"fmt"
	"strings"
)

type BenefitTerm string

const (
	BenefitTermAnnual   BenefitTerm = "annual"
	BenefitTermDay      BenefitTerm = "day"
	BenefitTermLifetime BenefitTerm = "lifetime"
)

func AllBenefitTerm() []BenefitTerm {
	return []BenefitTerm{
		BenefitTermAnnual,
		BenefitTermDay,
		BenefitTermLifetime,
	}
}

func FindBenefitTerm(filter string) []BenefitTerm {
	ret := make([]BenefitTerm, 0)
	for _, at := range AllBenefitTerm() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BenefitTerm) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BenefitTerm) String() string {
	switch cpt {
	case BenefitTermAnnual:
		return "Annual"
	case BenefitTermDay:
		return "Day"
	case BenefitTermLifetime:
		return "Lifetime"
	default:
		return "Unknown Benefit Term"
	}
}
