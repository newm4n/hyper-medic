package code

import (
	"fmt"
	"strings"
)

type AdjudicationError string

const (
	AdjudicationErrorMissingIdentifier   AdjudicationError = "a001"
	AdjudicationErrorMissingCreationDate AdjudicationError = "a002"
)

func AllAdjudicationError() []AdjudicationError {
	return []AdjudicationError{
		AdjudicationErrorMissingIdentifier,
		AdjudicationErrorMissingCreationDate,
	}
}

func FindAdjudicationError(filter string) []AdjudicationError {
	ret := make([]AdjudicationError, 0)
	for _, at := range AllAdjudicationError() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdjudicationError) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdjudicationError) String() string {
	switch cpt {
	case AdjudicationErrorMissingIdentifier:
		return "Missing Identifier"
	case AdjudicationErrorMissingCreationDate:
		return "Missing Creation Date"
	default:
		return "Unknown Adjudication Error"
	}
}

/**
Missing Identifier
Missing Creation Date
*/
