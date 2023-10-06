package code

import (
	"fmt"
	"strings"
)

type AuditEventSubType string

const (
	AuditEventSubTypeApplicationStart                AuditEventSubType = "110120"
	AuditEventSubTypeApplicationStop                 AuditEventSubType = "110121"
	AuditEventSubTypeLogin                           AuditEventSubType = "110122"
	AuditEventSubTypeLogout                          AuditEventSubType = "110123"
	AuditEventSubTypeAttach                          AuditEventSubType = "110124"
	AuditEventSubTypeDetach                          AuditEventSubType = "110125"
	AuditEventSubTypeNodeAuthentication              AuditEventSubType = "110126"
	AuditEventSubTypeEmergencyOverrideStarted        AuditEventSubType = "110127"
	AuditEventSubTypeNetworkConfiguration            AuditEventSubType = "110128"
	AuditEventSubTypeSecurityConfiguration           AuditEventSubType = "110129"
	AuditEventSubTypeHardwareConfiguration           AuditEventSubType = "110130"
	AuditEventSubTypeSoftwareConfiguration           AuditEventSubType = "110131"
	AuditEventSubTypeUseOfRestrictedFunction         AuditEventSubType = "110132"
	AuditEventSubTypeAuditRecordingStopped           AuditEventSubType = "110133"
	AuditEventSubTypeAuditRecordingStarted           AuditEventSubType = "110134"
	AuditEventSubTypeObjectSecurityAttributesChanged AuditEventSubType = "110135"
	AuditEventSubTypeSecurityRolesChanged            AuditEventSubType = "110136"
	AuditEventSubTypeUserSecurityAttributesChanged   AuditEventSubType = "110137"
	AuditEventSubTypeEmergencyOverrideStopped        AuditEventSubType = "110138"
	AuditEventSubTypeRemoteServiceOperationStarted   AuditEventSubType = "110139"
	AuditEventSubTypeRemoteServiceOperationStopped   AuditEventSubType = "110140"
	AuditEventSubTypeLocalServiceOperationStarted    AuditEventSubType = "110141"
	AuditEventSubTypeLocalServiceOperationStopped    AuditEventSubType = "110142"
	AuditEventSubTypeAuthenticationDecision          AuditEventSubType = "110143"
	AuditEventSubTypeAuthorizationDecision           AuditEventSubType = "110144"
	AuditEventSubTypeSessionStart                    AuditEventSubType = "110145"
	AuditEventSubTypeSessionStop                     AuditEventSubType = "110146"
	AuditEventSubTypeAccessControlDecision           AuditEventSubType = "110147"
	AuditEventSubTypeRead                            AuditEventSubType = "read"
	AuditEventSubTypeVread                           AuditEventSubType = "vread"
	AuditEventSubTypeUpdate                          AuditEventSubType = "update"
	AuditEventSubTypePatch                           AuditEventSubType = "patch"
	AuditEventSubTypeDelete                          AuditEventSubType = "delete"
	AuditEventSubTypeHistory                         AuditEventSubType = "history"
	AuditEventSubTypeHistoryInstance                 AuditEventSubType = "history-instance"
	AuditEventSubTypeHistoryType                     AuditEventSubType = "history-type"
	AuditEventSubTypeHistorySystem                   AuditEventSubType = "history-system"
	AuditEventSubTypeCreate                          AuditEventSubType = "create"
	AuditEventSubTypeSearch                          AuditEventSubType = "search"
	AuditEventSubTypeSearchType                      AuditEventSubType = "search-type"
	AuditEventSubTypeSearchSystem                    AuditEventSubType = "search-system"
	AuditEventSubTypeSearchCompartment               AuditEventSubType = "search-compartment"
	AuditEventSubTypeCapabilities                    AuditEventSubType = "capabilities"
	AuditEventSubTypeTransaction                     AuditEventSubType = "transaction"
	AuditEventSubTypeBatch                           AuditEventSubType = "batch"
	AuditEventSubTypeOperation                       AuditEventSubType = "operation"
)

func AllAuditEventSubType() []AuditEventSubType {
	return []AuditEventSubType{
		AuditEventSubTypeApplicationStart,
		AuditEventSubTypeApplicationStop,
		AuditEventSubTypeLogin,
		AuditEventSubTypeLogout,
		AuditEventSubTypeAttach,
		AuditEventSubTypeDetach,
		AuditEventSubTypeNodeAuthentication,
		AuditEventSubTypeEmergencyOverrideStarted,
		AuditEventSubTypeNetworkConfiguration,
		AuditEventSubTypeSecurityConfiguration,
		AuditEventSubTypeHardwareConfiguration,
		AuditEventSubTypeSoftwareConfiguration,
		AuditEventSubTypeUseOfRestrictedFunction,
		AuditEventSubTypeAuditRecordingStopped,
		AuditEventSubTypeAuditRecordingStarted,
		AuditEventSubTypeObjectSecurityAttributesChanged,
		AuditEventSubTypeSecurityRolesChanged,
		AuditEventSubTypeUserSecurityAttributesChanged,
		AuditEventSubTypeEmergencyOverrideStopped,
		AuditEventSubTypeRemoteServiceOperationStarted,
		AuditEventSubTypeRemoteServiceOperationStopped,
		AuditEventSubTypeLocalServiceOperationStarted,
		AuditEventSubTypeLocalServiceOperationStopped,
		AuditEventSubTypeAuthenticationDecision,
		AuditEventSubTypeAuthorizationDecision,
		AuditEventSubTypeSessionStart,
		AuditEventSubTypeSessionStop,
		AuditEventSubTypeAccessControlDecision,
		AuditEventSubTypeRead,
		AuditEventSubTypeVread,
		AuditEventSubTypeUpdate,
		AuditEventSubTypePatch,
		AuditEventSubTypeDelete,
		AuditEventSubTypeHistory,
		AuditEventSubTypeHistoryInstance,
		AuditEventSubTypeHistoryType,
		AuditEventSubTypeHistorySystem,
		AuditEventSubTypeCreate,
		AuditEventSubTypeSearch,
		AuditEventSubTypeSearchType,
		AuditEventSubTypeSearchSystem,
		AuditEventSubTypeSearchCompartment,
		AuditEventSubTypeCapabilities,
		AuditEventSubTypeTransaction,
		AuditEventSubTypeBatch,
		AuditEventSubTypeOperation,
	}
}

func FindAuditEventSubType(filter string) []AuditEventSubType {
	ret := make([]AuditEventSubType, 0)
	for _, at := range AllAuditEventSubType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AuditEventSubType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AuditEventSubType) String() string {
	switch cpt {
	case AuditEventSubTypeApplicationStart:
		return "Application Start"
	case AuditEventSubTypeApplicationStop:
		return "Application Stop"
	case AuditEventSubTypeLogin:
		return "Login"
	case AuditEventSubTypeLogout:
		return "Logout"
	case AuditEventSubTypeAttach:
		return "Attach"
	case AuditEventSubTypeDetach:
		return "Detach"
	case AuditEventSubTypeNodeAuthentication:
		return "Node Authentication"
	case AuditEventSubTypeEmergencyOverrideStarted:
		return "Emergency Override Started"
	case AuditEventSubTypeNetworkConfiguration:
		return "Network Configuration"
	case AuditEventSubTypeSecurityConfiguration:
		return "Security Configuration"
	case AuditEventSubTypeHardwareConfiguration:
		return "Hardware Configuration"
	case AuditEventSubTypeSoftwareConfiguration:
		return "Software Configuration"
	case AuditEventSubTypeUseOfRestrictedFunction:
		return "Use of Restricted Function"
	case AuditEventSubTypeAuditRecordingStopped:
		return "Audit Recording Stopped"
	case AuditEventSubTypeAuditRecordingStarted:
		return "Audit Recording Started"
	case AuditEventSubTypeObjectSecurityAttributesChanged:
		return "Object Security Attributes Changed"
	case AuditEventSubTypeSecurityRolesChanged:
		return "Security Roles Changed"
	case AuditEventSubTypeUserSecurityAttributesChanged:
		return "User security Attributes Changed"
	case AuditEventSubTypeEmergencyOverrideStopped:
		return "Emergency Override Stopped"
	case AuditEventSubTypeRemoteServiceOperationStarted:
		return "Remote Service Operation Started"
	case AuditEventSubTypeRemoteServiceOperationStopped:
		return "Remote Service Operation Stopped"
	case AuditEventSubTypeLocalServiceOperationStarted:
		return "Local Service Operation Started"
	case AuditEventSubTypeLocalServiceOperationStopped:
		return "Local Service Operation Stopped"
	case AuditEventSubTypeAuthenticationDecision:
		return "Authentication Decision"
	case AuditEventSubTypeAuthorizationDecision:
		return "Authorization Decision"
	case AuditEventSubTypeSessionStart:
		return "Session start"
	case AuditEventSubTypeSessionStop:
		return "Session stop"
	case AuditEventSubTypeAccessControlDecision:
		return "Access Control Decision"
	case AuditEventSubTypeRead:
		return "read"
	case AuditEventSubTypeVread:
		return "vread"
	case AuditEventSubTypeUpdate:
		return "update"
	case AuditEventSubTypePatch:
		return "patch"
	case AuditEventSubTypeDelete:
		return "delete"
	case AuditEventSubTypeHistory:
		return "history"
	case AuditEventSubTypeHistoryInstance:
		return "history-instance"
	case AuditEventSubTypeHistoryType:
		return "history-type"
	case AuditEventSubTypeHistorySystem:
		return "history-system"
	case AuditEventSubTypeCreate:
		return "create"
	case AuditEventSubTypeSearch:
		return "search"
	case AuditEventSubTypeSearchType:
		return "search-type"
	case AuditEventSubTypeSearchSystem:
		return "search-system"
	case AuditEventSubTypeSearchCompartment:
		return "search-compartment"
	case AuditEventSubTypeCapabilities:
		return "capabilities"
	case AuditEventSubTypeTransaction:
		return "transaction"
	case AuditEventSubTypeBatch:
		return "batch"
	case AuditEventSubTypeOperation:
		return "operation"

	default:
		return "Unknown Action Grouping Behavior"
	}
}
