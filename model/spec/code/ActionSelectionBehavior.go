package code

import (
	"fmt"
	"strings"
)

type ActionSelectionBehavior string

const (
	ActionSelectionBehaviorAny        ActionSelectionBehavior = "any"
	ActionSelectionBehaviorAll        ActionSelectionBehavior = "all"
	ActionSelectionBehaviorAllOrNone  ActionSelectionBehavior = "all-or-none"
	ActionSelectionBehaviorExactlyOne ActionSelectionBehavior = "exactly-one"
	ActionSelectionBehaviorAtMostOne  ActionSelectionBehavior = "at-most-one"
	ActionSelectionBehaviorOneOrMore  ActionSelectionBehavior = "one-or-more"
)

func AllActionSelectionBehavior() []ActionSelectionBehavior {
	return []ActionSelectionBehavior{
		ActionSelectionBehaviorAny,
		ActionSelectionBehaviorAll,
		ActionSelectionBehaviorAllOrNone,
		ActionSelectionBehaviorExactlyOne,
		ActionSelectionBehaviorAtMostOne,
		ActionSelectionBehaviorOneOrMore,
	}
}

func FindActionSelectionBehavior(filter string) []ActionSelectionBehavior {
	ret := make([]ActionSelectionBehavior, 0)
	for _, at := range AllActionSelectionBehavior() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionSelectionBehavior) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionSelectionBehavior) String() string {
	switch cpt {
	case ActionSelectionBehaviorAny:
		return "Any"
	case ActionSelectionBehaviorAll:
		return "All"
	case ActionSelectionBehaviorAllOrNone:
		return "All Or None"
	case ActionSelectionBehaviorExactlyOne:
		return "Exactly One"
	case ActionSelectionBehaviorAtMostOne:
		return "At Most One"
	case ActionSelectionBehaviorOneOrMore:
		return "One Or More"
	default:
		return "Unknown Action Type"
	}
}

/*
Any
All
All Or None
Exactly One
At Most One
One Or More
*/
