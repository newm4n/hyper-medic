package code

import (
	"fmt"
	"strings"
)

type ResourceSlicingRules string

const (
	ResourceSlicingRulesClosed    ResourceSlicingRules = "closed"
	ResourceSlicingRulesOpen      ResourceSlicingRules = "open"
	ResourceSlicingRulesOpenAtEnd ResourceSlicingRules = "openAtEnd"
)

func AllCResourceSlicingRules() []ResourceSlicingRules {
	return []ResourceSlicingRules{
		ResourceSlicingRulesClosed,
		ResourceSlicingRulesOpen,
		ResourceSlicingRulesOpenAtEnd,
	}
}

func FindResourceSlicingRules(filter string) []ResourceSlicingRules {
	ret := make([]ResourceSlicingRules, 0)
	for _, at := range AllCResourceSlicingRules() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ResourceSlicingRules) ToString() {
	fmt.Println(au.String())
}
func (au ResourceSlicingRules) String() string {
	switch au {
	case ResourceSlicingRulesClosed:
		return "Closed"
	case ResourceSlicingRulesOpen:
		return "Open"
	case ResourceSlicingRulesOpenAtEnd:
		return "Open at End"
	default:
		return "Unknown Resource Slicing Rules"
	}
}
