package code

import (
	"fmt"
	"strings"
)

type AuditEventSeverity string

const (
	AuditEventSeverityEmergency     AuditEventSeverity = "emergency"
	AuditEventSeverityAlert         AuditEventSeverity = "alert"
	AuditEventSeverityCritical      AuditEventSeverity = "critical"
	AuditEventSeverityError         AuditEventSeverity = "error"
	AuditEventSeverityWarning       AuditEventSeverity = "warning"
	AuditEventSeverityNotice        AuditEventSeverity = "notice"
	AuditEventSeverityInformational AuditEventSeverity = "informational"
	AuditEventSeverityDebug         AuditEventSeverity = "debug"
)

func AllAuditEventSeverity() []AuditEventSeverity {
	return []AuditEventSeverity{
		AuditEventSeverityEmergency,
		AuditEventSeverityAlert,
		AuditEventSeverityCritical,
		AuditEventSeverityError,
		AuditEventSeverityWarning,
		AuditEventSeverityNotice,
		AuditEventSeverityInformational,
		AuditEventSeverityDebug,
	}
}

func FindAuditEventSeverity(filter string) []AuditEventSeverity {
	ret := make([]AuditEventSeverity, 0)
	for _, at := range AllAuditEventSeverity() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AuditEventSeverity) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AuditEventSeverity) String() string {
	switch cpt {
	case AuditEventSeverityEmergency:
		return "Emergency"
	case AuditEventSeverityAlert:
		return "Alert"
	case AuditEventSeverityCritical:
		return "Critical"
	case AuditEventSeverityError:
		return "Error"
	case AuditEventSeverityWarning:
		return "Warning"
	case AuditEventSeverityNotice:
		return "Notice"
	case AuditEventSeverityInformational:
		return "Informational"
	case AuditEventSeverityDebug:
		return "Debug"
	default:
		return "Unknown Audit Event Severity"
	}
}
