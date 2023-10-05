package code

import (
	"fmt"
	"strings"
)

type AdjudicationReason string

const (
	AdjudicationReasonNotCovered       AdjudicationReason = "ar001"
	AdjudicationReasonPlanLimitReached AdjudicationReason = "ar002"
)

func AllAdjudicationReason() []AdjudicationReason {
	return []AdjudicationReason{
		AdjudicationReasonNotCovered,
		AdjudicationReasonPlanLimitReached,
	}
}

func FindAdjudicationReason(filter string) []AdjudicationReason {
	ret := make([]AdjudicationReason, 0)
	for _, at := range AllAdjudicationReason() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdjudicationReason) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdjudicationReason) String() string {
	switch cpt {
	case AdjudicationReasonNotCovered:
		return "Not covered"
	case AdjudicationReasonPlanLimitReached:
		return "Plan Limit Reached"
	default:
		return "Unknown Adjudication Reason"
	}
}

/*
Not covered
Plan Limit Reached
*/
