package code

import (
	"fmt"
	"strings"
)

type AdverseEventCausalityAssess int

const (
	AdverseEventCausalityAssessCertain AdverseEventCausalityAssess = iota
	AdverseEventCausalityAssessProbablyLikely
	AdverseEventCausalityAssessPossible
	AdverseEventCausalityAssessUnlikely
	AdverseEventCausalityAssessConditionalClassified
	AdverseEventCausalityAssessUnassessableUnclassifiable
)

func AllAdverseEventCausalityAssess() []AdverseEventCausalityAssess {
	return []AdverseEventCausalityAssess{
		AdverseEventCausalityAssessCertain,
		AdverseEventCausalityAssessProbablyLikely,
		AdverseEventCausalityAssessPossible,
		AdverseEventCausalityAssessUnlikely,
		AdverseEventCausalityAssessConditionalClassified,
		AdverseEventCausalityAssessUnassessableUnclassifiable,
	}
}

func FindAdverseEventCausalityAssess(filter string) []AdverseEventCausalityAssess {
	ret := make([]AdverseEventCausalityAssess, 0)
	for _, at := range AllAdverseEventCausalityAssess() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventCausalityAssess) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventCausalityAssess) String() string {
	switch cpt {
	case AdverseEventCausalityAssessCertain:
		return "Certain"
	case AdverseEventCausalityAssessProbablyLikely:
		return "Probably/Likely"
	case AdverseEventCausalityAssessPossible:
		return "Possible"
	case AdverseEventCausalityAssessUnlikely:
		return "Unlikely"
	case AdverseEventCausalityAssessConditionalClassified:
		return "Conditional/Classified"
	case AdverseEventCausalityAssessUnassessableUnclassifiable:
		return "Unassessable/Unclassifiable"
	default:
		return "Unknown Adverse Event Causality Assess"
	}
}

/*
Certain
Probably/Likely
Possible
Unlikely
Conditional/Classified
Unassessable/Unclassifiable
*/
