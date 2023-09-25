package code

import (
	"fmt"
	"strings"
)

type ActionCardinalityBehavior int

const (
	ActionCardinalityBehaviorSingle ActionCardinalityBehavior = iota
	ActionCardinalityBehaviorMultiple
)

func AllActionCardinalityBehavior() []ActionCardinalityBehavior {
	return []ActionCardinalityBehavior{
		ActionCardinalityBehaviorSingle,
		ActionCardinalityBehaviorMultiple,
	}
}

func FindActionCardinalityBehavior(filter string) []ActionCardinalityBehavior {
	ret := make([]ActionCardinalityBehavior, 0)
	for _, at := range AllActionConditionKind() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionCardinalityBehavior) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionCardinalityBehavior) String() string {
	switch cpt {
	case ActionCardinalityBehaviorSingle:
		return "Single"
	case ActionCardinalityBehaviorMultiple:
		return "Multiple"
	default:
		return "Unknown Action Grouping Behavior"
	}
}

//Single
//Multiple
