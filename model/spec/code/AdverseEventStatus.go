package code

import (
	"fmt"
	"strings"
)

type AdverseEventStatus int

const (
	AdverseEventStatusInProgress AdverseEventStatus = iota
	AdverseEventStatusCompleted
	AdverseEventStatusEnteredInError
	AdverseEventStatusUnknown
)

func AllAdverseEventStatus() []AdverseEventStatus {
	return []AdverseEventStatus{
		AdverseEventStatusInProgress,
		AdverseEventStatusCompleted,
		AdverseEventStatusEnteredInError,
		AdverseEventStatusUnknown,
	}
}

func FindAdverseEventStatus(filter string) []AdverseEventStatus {
	ret := make([]AdverseEventStatus, 0)
	for _, at := range AllAdverseEventStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventStatus) String() string {
	switch cpt {
	case AdverseEventStatusInProgress:
		return "In Progress"
	case AdverseEventStatusCompleted:
		return "Completed"
	case AdverseEventStatusEnteredInError:
		return "Entered in Error"
	case AdverseEventStatusUnknown:
		return "Unknown"
	default:
		return "Unknown Allergy Intolerance Clinical"
	}
}

/**
In Progress
Completed
Entered in Error
Unknown
*/
