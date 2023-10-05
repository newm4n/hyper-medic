package code

import (
	"fmt"
	"strings"
)

/* TODO change type to string */
type AssertResponseCodeTypes int

const (
	AssertResponseCodeTypesContinue AssertResponseCodeTypes = iota
	AssertResponseCodeTypesSwitchingProtocols
	AssertResponseCodeTypesOK
	AssertResponseCodeTypesCreated
	AssertResponseCodeTypesAccepted
	AssertResponseCodeTypesNonAuthoritativeInformation
	AssertResponseCodeTypesNoContent
	AssertResponseCodeTypesResetContent
	AssertResponseCodeTypesPartialContent
	AssertResponseCodeTypesMultipleChoices
	AssertResponseCodeTypesMovedPermanently
	AssertResponseCodeTypesFound
	AssertResponseCodeTypesSeeOther
	AssertResponseCodeTypesNotModified
	AssertResponseCodeTypesUseProxy
	AssertResponseCodeTypesTemporaryRedirect
	AssertResponseCodeTypesPermanentRedirect
	AssertResponseCodeTypesBadRequest
	AssertResponseCodeTypesUnauthorized
	AssertResponseCodeTypesPaymentRequired
	AssertResponseCodeTypesForbidden
	AssertResponseCodeTypesNotFound
	AssertResponseCodeTypesMethodNotAllowed
	AssertResponseCodeTypesNotAcceptable
	AssertResponseCodeTypesProxyAuthenticationRequired
	AssertResponseCodeTypesRequestTimeout
	AssertResponseCodeTypesConflict
	AssertResponseCodeTypesGone
	AssertResponseCodeTypesLengthRequired
	AssertResponseCodeTypesPreconditionFailed
	AssertResponseCodeTypesContentTooLarge
	AssertResponseCodeTypesURITooLong
	AssertResponseCodeTypesUnsupportedMediaType
	AssertResponseCodeTypesRangeNotSatisfiable
	AssertResponseCodeTypesExpectationFailed
	AssertResponseCodeTypesMisdirectedRequest
	AssertResponseCodeTypesUnprocessableContent
	AssertResponseCodeTypesUpgradeRequired
	AssertResponseCodeTypesInternalServerError
	AssertResponseCodeTypesNotImplemented
	AssertResponseCodeTypesBadGateway
	AssertResponseCodeTypesServiceUnavailable
	AssertResponseCodeTypesGatewayTimeout
	AssertResponseCodeTypesHTTPVersionNotSupported
)

func AllAssertResponseCodeTypes() []AssertResponseCodeTypes {
	return []AssertResponseCodeTypes{
		AssertResponseCodeTypesContinue,
		AssertResponseCodeTypesSwitchingProtocols,
		AssertResponseCodeTypesOK,
		AssertResponseCodeTypesCreated,
		AssertResponseCodeTypesAccepted,
		AssertResponseCodeTypesNonAuthoritativeInformation,
		AssertResponseCodeTypesNoContent,
		AssertResponseCodeTypesResetContent,
		AssertResponseCodeTypesPartialContent,
		AssertResponseCodeTypesMultipleChoices,
		AssertResponseCodeTypesMovedPermanently,
		AssertResponseCodeTypesFound,
		AssertResponseCodeTypesSeeOther,
		AssertResponseCodeTypesNotModified,
		AssertResponseCodeTypesUseProxy,
		AssertResponseCodeTypesTemporaryRedirect,
		AssertResponseCodeTypesPermanentRedirect,
		AssertResponseCodeTypesBadRequest,
		AssertResponseCodeTypesUnauthorized,
		AssertResponseCodeTypesPaymentRequired,
		AssertResponseCodeTypesForbidden,
		AssertResponseCodeTypesNotFound,
		AssertResponseCodeTypesMethodNotAllowed,
		AssertResponseCodeTypesNotAcceptable,
		AssertResponseCodeTypesProxyAuthenticationRequired,
		AssertResponseCodeTypesRequestTimeout,
		AssertResponseCodeTypesConflict,
		AssertResponseCodeTypesGone,
		AssertResponseCodeTypesLengthRequired,
		AssertResponseCodeTypesPreconditionFailed,
		AssertResponseCodeTypesContentTooLarge,
		AssertResponseCodeTypesURITooLong,
		AssertResponseCodeTypesUnsupportedMediaType,
		AssertResponseCodeTypesRangeNotSatisfiable,
		AssertResponseCodeTypesExpectationFailed,
		AssertResponseCodeTypesMisdirectedRequest,
		AssertResponseCodeTypesUnprocessableContent,
		AssertResponseCodeTypesUpgradeRequired,
		AssertResponseCodeTypesInternalServerError,
		AssertResponseCodeTypesNotImplemented,
		AssertResponseCodeTypesBadGateway,
		AssertResponseCodeTypesServiceUnavailable,
		AssertResponseCodeTypesGatewayTimeout,
		AssertResponseCodeTypesHTTPVersionNotSupported,
	}
}

func FindAssertResponseCodeTypes(filter string) []AssertResponseCodeTypes {
	ret := make([]AssertResponseCodeTypes, 0)
	for _, at := range AllAssertResponseCodeTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au AssertResponseCodeTypes) ToString() {
	fmt.Println(au.String())
}
func (au AssertResponseCodeTypes) String() string {
	switch au {
	case AssertResponseCodeTypesContinue:
		return "Continue"
	case AssertResponseCodeTypesSwitchingProtocols:
		return "Switching Protocols"
	case AssertResponseCodeTypesOK:
		return "OK"
	case AssertResponseCodeTypesCreated:
		return "Created"
	case AssertResponseCodeTypesAccepted:
		return "Accepted"
	case AssertResponseCodeTypesNonAuthoritativeInformation:
		return "Non-Authoritative Information"
	case AssertResponseCodeTypesNoContent:
		return "No Content"
	case AssertResponseCodeTypesResetContent:
		return "Reset Content"
	case AssertResponseCodeTypesPartialContent:
		return "Partial Content"
	case AssertResponseCodeTypesMultipleChoices:
		return "Multiple Choices"
	case AssertResponseCodeTypesMovedPermanently:
		return "Moved Permanently"
	case AssertResponseCodeTypesFound:
		return "Found"
	case AssertResponseCodeTypesSeeOther:
		return "See Other"
	case AssertResponseCodeTypesNotModified:
		return "Not Modified"
	case AssertResponseCodeTypesUseProxy:
		return "Use Proxy"
	case AssertResponseCodeTypesTemporaryRedirect:
		return "Temporary Redirect"
	case AssertResponseCodeTypesPermanentRedirect:
		return "Permanent Redirect"
	case AssertResponseCodeTypesBadRequest:
		return "Bad Request"
	case AssertResponseCodeTypesUnauthorized:
		return "Unauthorized"
	case AssertResponseCodeTypesPaymentRequired:
		return "Payment Required"
	case AssertResponseCodeTypesForbidden:
		return "Forbidden"
	case AssertResponseCodeTypesNotFound:
		return "Not Found"
	case AssertResponseCodeTypesMethodNotAllowed:
		return "Method Not Allowed"
	case AssertResponseCodeTypesNotAcceptable:
		return "Not Acceptable"
	case AssertResponseCodeTypesProxyAuthenticationRequired:
		return "Proxy Authentication Required"
	case AssertResponseCodeTypesRequestTimeout:
		return "Request Timeout"
	case AssertResponseCodeTypesConflict:
		return "Conflict"
	case AssertResponseCodeTypesGone:
		return "Gone"
	case AssertResponseCodeTypesLengthRequired:
		return "Length Required"
	case AssertResponseCodeTypesPreconditionFailed:
		return "Precondition Failed"
	case AssertResponseCodeTypesContentTooLarge:
		return "Content Too Large"
	case AssertResponseCodeTypesURITooLong:
		return "URI Too Long"
	case AssertResponseCodeTypesUnsupportedMediaType:
		return "Unsupported Media Type"
	case AssertResponseCodeTypesRangeNotSatisfiable:
		return "Range Not Satisfiable"
	case AssertResponseCodeTypesExpectationFailed:
		return "Expectation Failed"
	case AssertResponseCodeTypesMisdirectedRequest:
		return "Misdirected Request"
	case AssertResponseCodeTypesUnprocessableContent:
		return "Unprocessable Content"
	case AssertResponseCodeTypesUpgradeRequired:
		return "Upgrade Required"
	case AssertResponseCodeTypesInternalServerError:
		return "Internal Server Error"
	case AssertResponseCodeTypesNotImplemented:
		return "Not Implemented"
	case AssertResponseCodeTypesBadGateway:
		return "Bad Gateway"
	case AssertResponseCodeTypesServiceUnavailable:
		return "Service Unavailable"
	case AssertResponseCodeTypesGatewayTimeout:
		return "Gateway Timeout"
	case AssertResponseCodeTypesHTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	default:
		return "Unknown Address Use"
	}
}

/*
Continue
Switching Protocols
OK
Created
Accepted
Non-Authoritative Information
No Content
Reset Content
Partial Content
Multiple Choices
Moved Permanently
Found
See Other
Not Modified
Use Proxy
Temporary Redirect
Permanent Redirect
Bad Request
Unauthorized
Payment Required
Forbidden
Not Found
Method Not Allowed
Not Acceptable
Proxy Authentication Required
Request Timeout
Conflict
Gone
Length Required
Precondition Failed
Content Too Large
URI Too Long
Unsupported Media Type
Range Not Satisfiable
Expectation Failed
Misdirected Request
Unprocessable Content
Upgrade Required
Internal Server Error
Not Implemented
Bad Gateway
Service Unavailable
Gateway Timeout
HTTP Version Not Supported

*/
