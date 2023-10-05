package code

import (
	"fmt"
	"strings"
)

type ActionPrecheckBehavior string

const (
	ActionPrecheckBehaviorYes ActionPrecheckBehavior = "yes"
	ActionPrecheckBehaviorNo  ActionPrecheckBehavior = "no"
)

func AllActionPrecheckBehavior() []ActionPrecheckBehavior {
	return []ActionPrecheckBehavior{
		ActionPrecheckBehaviorYes,
		ActionPrecheckBehaviorNo,
	}
}

func FindActionPrecheckBehavior(filter string) []ActionPrecheckBehavior {
	ret := make([]ActionPrecheckBehavior, 0)
	for _, at := range AllActionPrecheckBehavior() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionPrecheckBehavior) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionPrecheckBehavior) String() string {
	switch cpt {
	case ActionPrecheckBehaviorYes:
		return "Yes"
	case ActionPrecheckBehaviorNo:
		return "No"
	default:
		return "Unknown Action Precheck Behavior"
	}
}

/*
Yes
No
*/
