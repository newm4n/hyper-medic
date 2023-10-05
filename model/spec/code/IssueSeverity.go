package code

import (
	"fmt"
	"strings"
)

type IssueSeverity string

const (
	IssueSeverityFatal               IssueSeverity = "fatal"
	IssueSeverityError               IssueSeverity = "error"
	IssueSeverityWarning             IssueSeverity = "warning"
	IssueSeverityInformation         IssueSeverity = "information"
	IssueSeverityOperationSuccessful IssueSeverity = "success"
)

func AllIssueSeverity() []IssueSeverity {
	return []IssueSeverity{
		IssueSeverityFatal,
		IssueSeverityError,
		IssueSeverityWarning,
		IssueSeverityInformation,
		IssueSeverityOperationSuccessful,
	}
}

func FindIssueSeverity(filter string) []IssueSeverity {
	ret := make([]IssueSeverity, 0)
	for _, at := range AllIssueSeverity() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au IssueSeverity) ToString() {
	fmt.Println(au.String())
}
func (au IssueSeverity) String() string {
	switch au {
	case IssueSeverityFatal:
		return "Fatal"
	case IssueSeverityError:
		return "Error"
	case IssueSeverityWarning:
		return "Warning"
	case IssueSeverityInformation:
		return "Information"
	case IssueSeverityOperationSuccessful:
		return "Operation Successful"
	default:
		return "Unknown Issue Severity"
	}
}
