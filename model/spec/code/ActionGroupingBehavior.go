package code

import (
	"fmt"
	"strings"
)

type ActionGroupingBehavior string

const (
	ActionGroupingBehaviorVisualGroup   ActionGroupingBehavior = "visual-group"
	ActionGroupingBehaviorLogicalGroup  ActionGroupingBehavior = "logical-group"
	ActionGroupingBehaviorSentenceGroup ActionGroupingBehavior = "sentence-group"
)

func AllActionGroupingBehavior() []ActionGroupingBehavior {
	return []ActionGroupingBehavior{
		ActionGroupingBehaviorVisualGroup,
		ActionGroupingBehaviorLogicalGroup,
		ActionGroupingBehaviorSentenceGroup,
	}
}

func FindActionGroupingBehavior(filter string) []ActionGroupingBehavior {
	ret := make([]ActionGroupingBehavior, 0)
	for _, at := range AllActionGroupingBehavior() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionGroupingBehavior) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionGroupingBehavior) String() string {
	switch cpt {
	case ActionGroupingBehaviorVisualGroup:
		return "Visual Group"
	case ActionGroupingBehaviorLogicalGroup:
		return "Logical Group"
	case ActionGroupingBehaviorSentenceGroup:
		return "Sentence Group"
	default:
		return "Unknown Action Grouping Behavior"
	}
}

/*
Visual Group
Logical Group
Sentence Group

*/
