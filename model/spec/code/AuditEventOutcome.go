package code

import (
	"fmt"
	"strings"
)

type AuditEventOutcome string

const (
	AuditEventOutcomeFatal               AuditEventOutcome = "fatal"
	AuditEventOutcomeError               AuditEventOutcome = "error"
	AuditEventOutcomeWarning             AuditEventOutcome = "warning"
	AuditEventOutcomeInformation         AuditEventOutcome = "information"
	AuditEventOutcomeOperationSuccessful AuditEventOutcome = "success"
)

func AllAuditEventOutcome() []AuditEventOutcome {
	return []AuditEventOutcome{
		AuditEventOutcomeFatal,
		AuditEventOutcomeError,
		AuditEventOutcomeWarning,
		AuditEventOutcomeInformation,
		AuditEventOutcomeOperationSuccessful,
	}
}

func FindAuditEventOutcome(filter string) []AuditEventOutcome {
	ret := make([]AuditEventOutcome, 0)
	for _, at := range AllAuditEventOutcome() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AuditEventOutcome) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AuditEventOutcome) String() string {
	switch cpt {
	case AuditEventOutcomeFatal:
		return "Fatal"
	case AuditEventOutcomeError:
		return "Error"
	case AuditEventOutcomeWarning:
		return "Warning"
	case AuditEventOutcomeInformation:
		return "Information"
	case AuditEventOutcomeOperationSuccessful:
		return "Operation Successful"
	default:
		return "Unknown Audit Event Outcome"
	}
}
