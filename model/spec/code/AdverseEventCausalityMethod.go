package code

import (
	"fmt"
	"strings"
)

type AdverseEventCausalityMethod int

const (
	AdverseEventCausalityMethodProbabilityScale AdverseEventCausalityMethod = iota
	AdverseEventCausalityMethodBayesian
	AdverseEventCausalityMethodChecklist
)

func AllAdverseEventCausalityMethod() []AdverseEventCausalityMethod {
	return []AdverseEventCausalityMethod{
		AdverseEventCausalityMethodProbabilityScale,
		AdverseEventCausalityMethodBayesian,
		AdverseEventCausalityMethodChecklist,
	}
}

func FindAdverseEventCausalityMethod(filter string) []AdverseEventCausalityMethod {
	ret := make([]AdverseEventCausalityMethod, 0)
	for _, at := range AllAdverseEventCausalityMethod() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventCausalityMethod) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventCausalityMethod) String() string {
	switch cpt {
	case AdverseEventCausalityMethodProbabilityScale:
		return "Probability Scale"
	case AdverseEventCausalityMethodBayesian:
		return "Bayesian"
	case AdverseEventCausalityMethodChecklist:
		return "Checklist"
	default:
		return "Unknown Adverse Event Outcome"
	}
}

/*
Probability Scale
Bayesian
Checklist
*/
