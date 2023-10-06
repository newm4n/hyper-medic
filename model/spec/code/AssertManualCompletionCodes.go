package code

import (
	"fmt"
	"strings"
)

type AssertManualCompletionCodes string

const (
	AssertManualCompletionCodesFail AssertManualCompletionCodes = "fail"
	AssertManualCompletionCodesPass AssertManualCompletionCodes = "pass"
	AssertManualCompletionCodesSkip AssertManualCompletionCodes = "skip"
	AssertManualCompletionCodesStop AssertManualCompletionCodes = "stop"
)

func AllAssertManualCompletionCodes() []AssertManualCompletionCodes {
	return []AssertManualCompletionCodes{
		AssertManualCompletionCodesFail,
		AssertManualCompletionCodesPass,
		AssertManualCompletionCodesSkip,
		AssertManualCompletionCodesStop,
	}
}

func FindAssertManualCompletionCodes(filter string) []AssertManualCompletionCodes {
	ret := make([]AssertManualCompletionCodes, 0)
	for _, at := range AllAssertManualCompletionCodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AssertManualCompletionCodes) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AssertManualCompletionCodes) String() string {
	switch cpt {
	case AssertManualCompletionCodesFail:
		return "Fail"
	case AssertManualCompletionCodesPass:
		return "Pass"
	case AssertManualCompletionCodesSkip:
		return "Skip"
	case AssertManualCompletionCodesStop:
		return "Stop"
	default:
		return "Unknown Assert Manual Completion Codes"
	}
}
