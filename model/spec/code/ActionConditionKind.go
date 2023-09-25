package code

import (
	"fmt"
	"strings"
)

type ActionConditionKind int

const (
	ActionConditionKindApplicability ActionConditionKind = iota
	ActionConditionKindStart
	ActionConditionKindStop
)

func AllActionConditionKind() []ActionConditionKind {
	return []ActionConditionKind{
		ActionConditionKindApplicability,
		ActionConditionKindStart,
		ActionConditionKindStop,
	}
}

func FindActionConditionKind(filter string) []ActionConditionKind {
	ret := make([]ActionConditionKind, 0)
	for _, at := range AllActionConditionKind() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionConditionKind) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionConditionKind) String() string {
	switch cpt {
	case ActionConditionKindApplicability:
		return "Applicability"
	case ActionConditionKindStart:
		return "Start"
	case ActionConditionKindStop:
		return "Stop"
	default:
		return "Unknown Action Grouping Behavior"
	}
}

/*
Applicability
Start
Stop
*/
