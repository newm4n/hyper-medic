package code

import (
	"fmt"
	"strings"
)

type AssertOperatorCodes string

const (
	AssertOperatorCodesEquals         AssertOperatorCodes = "equals"
	AssertOperatorCodesNotEquals      AssertOperatorCodes = "notEquals"
	AssertOperatorCodesIn             AssertOperatorCodes = "in"
	AssertOperatorCodesNotIn          AssertOperatorCodes = "notIn"
	AssertOperatorCodesGreaterThan    AssertOperatorCodes = "greaterThan"
	AssertOperatorCodesLessThan       AssertOperatorCodes = "lessThan"
	AssertOperatorCodesEmpty          AssertOperatorCodes = "empty"
	AssertOperatorCodesNotEmpty       AssertOperatorCodes = "notEmpty"
	AssertOperatorCodesContains       AssertOperatorCodes = "contains"
	AssertOperatorCodesNotContains    AssertOperatorCodes = "notContains"
	AssertOperatorCodesEvaluate       AssertOperatorCodes = "eval"
	AssertOperatorCodesManualEvaluate AssertOperatorCodes = "manualEval"
)

func AllAssertOperatorCodes() []AssertOperatorCodes {
	return []AssertOperatorCodes{
		AssertOperatorCodesEquals,
		AssertOperatorCodesNotEquals,
		AssertOperatorCodesIn,
		AssertOperatorCodesNotIn,
		AssertOperatorCodesGreaterThan,
		AssertOperatorCodesLessThan,
		AssertOperatorCodesEmpty,
		AssertOperatorCodesNotEmpty,
		AssertOperatorCodesContains,
		AssertOperatorCodesNotContains,
		AssertOperatorCodesEvaluate,
		AssertOperatorCodesManualEvaluate,
	}
}

func FindAssertOperatorCodes(filter string) []AssertOperatorCodes {
	ret := make([]AssertOperatorCodes, 0)
	for _, at := range AllAssertOperatorCodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AssertOperatorCodes) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AssertOperatorCodes) String() string {
	switch cpt {
	case AssertOperatorCodesEquals:
		return "equals"
	case AssertOperatorCodesNotEquals:
		return "notEquals"
	case AssertOperatorCodesIn:
		return "in"
	case AssertOperatorCodesNotIn:
		return "notIn"
	case AssertOperatorCodesGreaterThan:
		return "greaterThan"
	case AssertOperatorCodesLessThan:
		return "lessThan"
	case AssertOperatorCodesEmpty:
		return "empty"
	case AssertOperatorCodesNotEmpty:
		return "notEmpty"
	case AssertOperatorCodesContains:
		return "contains"
	case AssertOperatorCodesNotContains:
		return "notContains"
	case AssertOperatorCodesEvaluate:
		return "eval"
	case AssertOperatorCodesManualEvaluate:
		return "manualEval"

	default:
		return "Unknown Assert Operator Codes"
	}
}
