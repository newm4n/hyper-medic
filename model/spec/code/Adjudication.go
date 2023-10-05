package code

import (
	"fmt"
	"strings"
)

type Adjudication string

const (
	AdjudicationSubmittedAmount       Adjudication = "submitted"
	AdjudicationCoPay                 Adjudication = "copay"
	AdjudicationEligibleAmount        Adjudication = "eligible"
	AdjudicationDeductible            Adjudication = "deductible"
	AdjudicationUnallocatedDeductible Adjudication = "unallocdeduct"
	AdjudicationEligiblePercentage    Adjudication = "eligpercent"
	Adjudicationax                    Adjudication = "tax"
	AdjudicationBenefitAmount         Adjudication = "benefit"
)

func AllAdjudication() []Adjudication {
	return []Adjudication{
		AdjudicationSubmittedAmount,
		AdjudicationCoPay,
		AdjudicationEligibleAmount,
		AdjudicationDeductible,
		AdjudicationUnallocatedDeductible,
		AdjudicationEligiblePercentage,
		Adjudicationax,
		AdjudicationBenefitAmount,
	}
}

func FindAdjudication(filter string) []Adjudication {
	ret := make([]Adjudication, 0)
	for _, at := range AllAdjudication() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt Adjudication) ToString() {
	fmt.Println(cpt.String())
}

func (cpt Adjudication) String() string {
	switch cpt {
	case AdjudicationSubmittedAmount:
		return "Submitted Amount"
	case AdjudicationCoPay:
		return "CoPay"
	case AdjudicationEligibleAmount:
		return "Eligible Amount"
	case AdjudicationDeductible:
		return "Deductible"
	case AdjudicationUnallocatedDeductible:
		return "Unallocated Deductible"
	case AdjudicationEligiblePercentage:
		return "Eligible %"
	case Adjudicationax:
		return "Tax"
	case AdjudicationBenefitAmount:
		return "Benefit Amount"
	default:
		return "Unknown Adjudication"
	}
}

/*
Submitted Amount
CoPay
Eligible Amount
Deductible
Unallocated Deductible
Eligible %
Tax
Benefit Amount
*/
