package code

import (
	"fmt"
	"strings"
)

type ActionType string

const (
	ActionTypeCreate    ActionType = "create"
	ActionTypeUpdate    ActionType = "update"
	ActionTypeRemove    ActionType = "move"
	ActionTypeFireEvent ActionType = "fire-event"
)

func AllActionType() []ActionType {
	return []ActionType{
		ActionTypeCreate,
		ActionTypeUpdate,
		ActionTypeRemove,
		ActionTypeFireEvent,
	}
}

func FindActionType(filter string) []ActionType {
	ret := make([]ActionType, 0)
	for _, at := range AllActionType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionType) String() string {
	switch cpt {
	case ActionTypeCreate:
		return "Create"
	case ActionTypeUpdate:
		return "Update"
	case ActionTypeRemove:
		return "Remove"
	case ActionTypeFireEvent:
		return "Fire Event"
	default:
		return "Unknown Action Type"
	}
}

/*
Create
Update
Remove
Fire Event
*/
