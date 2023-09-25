package code

import (
	"fmt"
	"strings"
)

type ActionRequiredBehavior int

const (
	ActionRequiredBehaviorMust ActionRequiredBehavior = iota
	ActionRequiredBehaviorCould
	ActionRequiredBehaviorMustUnlessDocumented
)

func AllActionRequiredBehavior() []ActionRequiredBehavior {
	return []ActionRequiredBehavior{
		ActionRequiredBehaviorMust,
		ActionRequiredBehaviorCould,
		ActionRequiredBehaviorMustUnlessDocumented,
	}
}

func FindActionRequiredBehavior(filter string) []ActionRequiredBehavior {
	ret := make([]ActionRequiredBehavior, 0)
	for _, at := range AllActionRequiredBehavior() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionRequiredBehavior) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionRequiredBehavior) String() string {
	switch cpt {
	case ActionRequiredBehaviorMust:
		return "Must"
	case ActionRequiredBehaviorCould:
		return "Could"
	case ActionRequiredBehaviorMustUnlessDocumented:
		return "Must Unless Documented"
	default:
		return "Unknown Action Required Behavior"
	}
}

/*
Must
Could
Must Unless Documented
*/
