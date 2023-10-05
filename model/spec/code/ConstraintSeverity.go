package code

import (
	"fmt"
	"strings"
)

type ConstraintSeverity string

const (
	ConstraintSeverityError   ConstraintSeverity = "error"
	ConstraintSeverityWarning ConstraintSeverity = "warning"
)

func AllConstraintSeverity() []ConstraintSeverity {
	return []ConstraintSeverity{
		ConstraintSeverityError,
		ConstraintSeverityWarning,
	}
}

func FindConstraintSeverity(filter string) []ConstraintSeverity {
	ret := make([]ConstraintSeverity, 0)
	for _, at := range AllConstraintSeverity() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ConstraintSeverity) ToString() {
	fmt.Println(au.String())
}
func (au ConstraintSeverity) String() string {
	switch au {
	case ConstraintSeverityError:
		return "Error"
	case ConstraintSeverityWarning:
		return "Warning"
	default:
		return "Unknown Constraint Severity"
	}
}
