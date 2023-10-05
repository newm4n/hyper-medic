package code

import (
	"fmt"
	"strings"
)

type FilterOperator string

const (
	FilterOperatorEquals                    FilterOperator = "="
	FilterOperatorIsABySubsumption          FilterOperator = "is-a"
	FilterOperatorDescendentOfBySubsumption FilterOperator = "descendent-of"
	FilterOperatorNotIsABySubsumption       FilterOperator = "is-not-a"
	FilterOperatorRegularExpression         FilterOperator = "regex"
	FilterOperatorInSet                     FilterOperator = "in"
	FilterOperatorNotInSet                  FilterOperator = "not-in"
	FilterOperatorGeneralizesBySubsumption  FilterOperator = "generalizes"
	FilterOperatorChildOf                   FilterOperator = "child-of"
	FilterOperatorDescendentLeaf            FilterOperator = "descendent-leaf"
	FilterOperatorExists                    FilterOperator = "exists"
)

func AllFilterOperator() []FilterOperator {
	return []FilterOperator{
		FilterOperatorEquals,
		FilterOperatorIsABySubsumption,
		FilterOperatorDescendentOfBySubsumption,
		FilterOperatorNotIsABySubsumption,
		FilterOperatorRegularExpression,
		FilterOperatorInSet,
		FilterOperatorNotInSet,
		FilterOperatorGeneralizesBySubsumption,
		FilterOperatorChildOf,
		FilterOperatorDescendentLeaf,
		FilterOperatorExists,
	}
}

func FindFilterOperator(filter string) []FilterOperator {
	ret := make([]FilterOperator, 0)
	for _, at := range AllFilterOperator() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au FilterOperator) ToString() {
	fmt.Println(au.String())
}
func (au FilterOperator) String() string {
	switch au {
	case FilterOperatorEquals:
		return "Equals"
	case FilterOperatorIsABySubsumption:
		return "Is A (by subsumption)"
	case FilterOperatorDescendentOfBySubsumption:
		return "Descendent Of (by subsumption)"
	case FilterOperatorNotIsABySubsumption:
		return "Not (Is A) (by subsumption)"
	case FilterOperatorRegularExpression:
		return "Regular Expression"
	case FilterOperatorInSet:
		return "In Set"
	case FilterOperatorNotInSet:
		return "Not in Set"
	case FilterOperatorGeneralizesBySubsumption:
		return "Generalizes (by Subsumption)"
	case FilterOperatorChildOf:
		return "Child Of"
	case FilterOperatorDescendentLeaf:
		return "Descendent Leaf"
	case FilterOperatorExists:
		return "Exists"
	default:
		return "Unknown Constraint Severity"
	}
}
