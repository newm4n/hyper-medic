package code

import (
	"fmt"
	"strings"
)

type ObservationStatus string

const (
	ObservationStatusRegistered     ObservationStatus = "registered"
	ObservationStatusPreliminary    ObservationStatus = "preliminary"
	ObservationStatusFinal          ObservationStatus = "final"
	ObservationStatusAmended        ObservationStatus = "amended"
	ObservationStatusCorrected      ObservationStatus = "corrected"
	ObservationStatusCancelled      ObservationStatus = "cancelled"
	ObservationStatusEnteredInError ObservationStatus = "entered"
	ObservationStatusUnknown        ObservationStatus = "unknown"
)

func AllObservationStatus() []ObservationStatus {
	return []ObservationStatus{
		ObservationStatusRegistered,
		ObservationStatusPreliminary,
		ObservationStatusFinal,
		ObservationStatusAmended,
		ObservationStatusCorrected,
		ObservationStatusCancelled,
		ObservationStatusEnteredInError,
		ObservationStatusUnknown,
	}
}

func FindObservationStatus(filter string) []ObservationStatus {
	ret := make([]ObservationStatus, 0)
	for _, at := range AllObservationStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ObservationStatus) ToString() {
	fmt.Println(au.String())
}
func (au ObservationStatus) String() string {
	switch au {
	case ObservationStatusRegistered:
		return "Registered"
	case ObservationStatusPreliminary:
		return "Preliminary"
	case ObservationStatusFinal:
		return "Final"
	case ObservationStatusAmended:
		return "Amended"
	case ObservationStatusCorrected:
		return "Corrected"
	case ObservationStatusCancelled:
		return "Cancelled"
	case ObservationStatusEnteredInError:
		return "Entered in Error"
	case ObservationStatusUnknown:
		return "Unknown"
	default:
		return "Unknown Constraint Severity"
	}
}
