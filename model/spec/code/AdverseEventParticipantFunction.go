package code

import (
	"fmt"
	"strings"
)

type AdverseEventParticipantFunction string

const (
	AdverseEventParticipantFunctionInformant        AdverseEventParticipantFunction = "INF"
	AdverseEventParticipantFunctionParticipation    AdverseEventParticipantFunction = "PART"
	AdverseEventParticipantFunctionWitness          AdverseEventParticipantFunction = "WIT"
	AdverseEventParticipantFunctionAuthorOriginator AdverseEventParticipantFunction = "AUT"
)

func AllAAdverseEventParticipantFunction() []AdverseEventParticipantFunction {
	return []AdverseEventParticipantFunction{
		AdverseEventParticipantFunctionInformant,
		AdverseEventParticipantFunctionParticipation,
		AdverseEventParticipantFunctionWitness,
		AdverseEventParticipantFunctionAuthorOriginator,
	}
}

func FindAdverseEventParticipantFunction(filter string) []AdverseEventParticipantFunction {
	ret := make([]AdverseEventParticipantFunction, 0)
	for _, at := range AllAAdverseEventParticipantFunction() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventParticipantFunction) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventParticipantFunction) String() string {
	switch cpt {
	case AdverseEventParticipantFunctionInformant:
		return "informant"
	case AdverseEventParticipantFunctionParticipation:
		return "participation"
	case AdverseEventParticipantFunctionWitness:
		return "witness"
	case AdverseEventParticipantFunctionAuthorOriginator:
		return "author (originator)"
	default:
		return "Unknown Adverse Event Participant Function"
	}
}

/**
informant
Participation
witness
author (originator)
*/
