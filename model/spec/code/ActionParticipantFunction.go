package code

import (
	"fmt"
	"strings"
)

type ActionParticipantFunction string

const (
	ActionParticipantFunctionPerformer ActionParticipantFunction = "performer"
	ActionParticipantFunctionAuthor    ActionParticipantFunction = "author"
	ActionParticipantFunctionReviewer  ActionParticipantFunction = "reviewer"
	ActionParticipantFunctionWitness   ActionParticipantFunction = "witness"
)

func AllActionParticipantFunction() []ActionParticipantFunction {
	return []ActionParticipantFunction{
		ActionParticipantFunctionPerformer,
		ActionParticipantFunctionAuthor,
		ActionParticipantFunctionReviewer,
		ActionParticipantFunctionWitness,
	}
}

func FindActionParticipantFunction(filter string) []ActionParticipantFunction {
	ret := make([]ActionParticipantFunction, 0)
	for _, at := range AllActionParticipantFunction() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionParticipantFunction) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionParticipantFunction) String() string {
	switch cpt {
	case ActionParticipantFunctionPerformer:
		return "Performer"
	case ActionParticipantFunctionAuthor:
		return "Author"
	case ActionParticipantFunctionReviewer:
		return "Reviewer"
	case ActionParticipantFunctionWitness:
		return "Witness"
	default:
		return "Unknown Action Participant Function"
	}
}

/*
Performer
Author
Reviewer
Witness
*/
