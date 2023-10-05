package code

import (
	"fmt"
	"strings"
)

type IssueType string

const (
	IssueTypeInvalid         IssueType = "invalid"
	IssueTypeStructure       IssueType = "structure"
	IssueTypeRequired        IssueType = "required"
	IssueTypeValue           IssueType = "value"
	IssueTypeInvariant       IssueType = "invariant"
	IssueTypeSecurity        IssueType = "security"
	IssueTypeLogin           IssueType = "login"
	IssueTypeUnknown         IssueType = "unknown"
	IssueTypeExpired         IssueType = "expired"
	IssueTypeForbidden       IssueType = "forbidden"
	IssueTypeSuppressed      IssueType = "suppressed"
	IssueTypeProcessing      IssueType = "processing"
	IssueTypeNotSupported    IssueType = "not-supported"
	IssueTypeDuplicate       IssueType = "duplicate"
	IssueTypeMultipleMatches IssueType = "multiple-matches"
	IssueTypeNotFound        IssueType = "not-found"
	IssueTypeDeleted         IssueType = "deleted"
	IssueTypeTooLong         IssueType = "too-long"
	IssueTypeCodeInvalid     IssueType = "code-invalid"
	IssueTypeExtension       IssueType = "extension"
	IssueTypeTooCostly       IssueType = "too-costly"
	IssueTypeBusinessRule    IssueType = "business-rule"
	IssueTypeConflict        IssueType = "conflict"
	IssueTypeLimitedFilter   IssueType = "limited-filter"
	IssueTypeTransient       IssueType = "transient"
	IssueTypeLockError       IssueType = "lock-error"
	IssueTypeNoStore         IssueType = "no-store"
	IssueTypeException       IssueType = "exception"
	IssueTypeTimeout         IssueType = "timeout"
	IssueTypeIncomplete      IssueType = "incomplete"
	IssueTypeThrottled       IssueType = "throttled"
	IssueTypeInformational   IssueType = "informational"
	IssueTypeSuccess         IssueType = "success"
)

func AllIssueType() []IssueType {
	return []IssueType{
		IssueTypeInvalid,
		IssueTypeStructure,
		IssueTypeRequired,
		IssueTypeValue,
		IssueTypeInvariant,
		IssueTypeSecurity,
		IssueTypeLogin,
		IssueTypeUnknown,
		IssueTypeExpired,
		IssueTypeForbidden,
		IssueTypeSuppressed,
		IssueTypeProcessing,
		IssueTypeNotSupported,
		IssueTypeDuplicate,
		IssueTypeMultipleMatches,
		IssueTypeNotFound,
		IssueTypeDeleted,
		IssueTypeTooLong,
		IssueTypeCodeInvalid,
		IssueTypeExtension,
		IssueTypeTooCostly,
		IssueTypeBusinessRule,
		IssueTypeConflict,
		IssueTypeLimitedFilter,
		IssueTypeTransient,
		IssueTypeLockError,
		IssueTypeNoStore,
		IssueTypeException,
		IssueTypeTimeout,
		IssueTypeIncomplete,
		IssueTypeThrottled,
		IssueTypeInformational,
		IssueTypeSuccess,
	}
}

func FindIssueType(filter string) []IssueType {
	ret := make([]IssueType, 0)
	for _, at := range AllIssueType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au IssueType) ToString() {
	fmt.Println(au.String())
}
func (au IssueType) String() string {
	switch au {
	case IssueTypeInvalid:
		return "Invalid Content"
	case IssueTypeStructure:
		return "Structural Issue"
	case IssueTypeRequired:
		return "Required element missing"
	case IssueTypeValue:
		return "Element value invalid"
	case IssueTypeInvariant:
		return "Validation rule failed"
	case IssueTypeSecurity:
		return "Security Problem"
	case IssueTypeLogin:
		return "Login Required"
	case IssueTypeUnknown:
		return "Unknown User"
	case IssueTypeExpired:
		return "Session Expired"
	case IssueTypeForbidden:
		return "Forbidden"
	case IssueTypeSuppressed:
		return "Information Suppressed"
	case IssueTypeProcessing:
		return "Processing Failure"
	case IssueTypeNotSupported:
		return "Content not supported"
	case IssueTypeDuplicate:
		return "Duplicate"
	case IssueTypeMultipleMatches:
		return "Multiple Matches"
	case IssueTypeNotFound:
		return "Not Found"
	case IssueTypeDeleted:
		return "Deleted"
	case IssueTypeTooLong:
		return "Content Too Long"
	case IssueTypeCodeInvalid:
		return "Invalid Code"
	case IssueTypeExtension:
		return "Unacceptable Extension"
	case IssueTypeTooCostly:
		return "Operation Too Costly"
	case IssueTypeBusinessRule:
		return "Business Rule Violation"
	case IssueTypeConflict:
		return "Edit Version Conflict"
	case IssueTypeLimitedFilter:
		return "Limited Filter Application"
	case IssueTypeTransient:
		return "Transient Issue"
	case IssueTypeLockError:
		return "Lock Error"
	case IssueTypeNoStore:
		return "No Store Available"
	case IssueTypeException:
		return "Exception"
	case IssueTypeTimeout:
		return "Timeout"
	case IssueTypeIncomplete:
		return "Incomplete Results"
	case IssueTypeThrottled:
		return "Throttled"
	case IssueTypeInformational:
		return "Informational Note"
	case IssueTypeSuccess:
		return "Operation Successful"

	default:
		return "Unknown Issue Type"
	}
}
