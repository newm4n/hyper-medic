package code

import (
	"fmt"
	"strings"
)

type AssertResponseCodeTypes string

const (
	AssertResponseCodeTypesContinue                    AssertResponseCodeTypes = "continue"
	AssertResponseCodeTypesSwitchingProtocols          AssertResponseCodeTypes = "switchingProtocols"
	AssertResponseCodeTypesOkay                        AssertResponseCodeTypes = "okay"
	AssertResponseCodeTypesCreated                     AssertResponseCodeTypes = "created"
	AssertResponseCodeTypesAccepted                    AssertResponseCodeTypes = "accepted"
	AssertResponseCodeTypesNonAuthoritativeInformation AssertResponseCodeTypes = "nonAuthoritativeInformation"
	AssertResponseCodeTypesNoContent                   AssertResponseCodeTypes = "noContent"
	AssertResponseCodeTypesResetContent                AssertResponseCodeTypes = "resetContent"
	AssertResponseCodeTypesPartialContent              AssertResponseCodeTypes = "partialContent"
	AssertResponseCodeTypesMultipleChoices             AssertResponseCodeTypes = "multipleChoices"
	AssertResponseCodeTypesMovedPermanently            AssertResponseCodeTypes = "movedPermanently"
	AssertResponseCodeTypesFound                       AssertResponseCodeTypes = "found"
	AssertResponseCodeTypesSeeOther                    AssertResponseCodeTypes = "seeOther"
	AssertResponseCodeTypesNotModified                 AssertResponseCodeTypes = "notModified"
	AssertResponseCodeTypesUseProxy                    AssertResponseCodeTypes = "useProxy"
	AssertResponseCodeTypesTemporaryRedirect           AssertResponseCodeTypes = "temporaryRedirect"
	AssertResponseCodeTypesPermanentRedirect           AssertResponseCodeTypes = "permanentRedirect"
	AssertResponseCodeTypesBadRequest                  AssertResponseCodeTypes = "badRequest"
	AssertResponseCodeTypesUnauthorized                AssertResponseCodeTypes = "unauthorized"
	AssertResponseCodeTypesPaymentRequired             AssertResponseCodeTypes = "paymentRequired"
	AssertResponseCodeTypesForbidden                   AssertResponseCodeTypes = "forbidden"
	AssertResponseCodeTypesNotFound                    AssertResponseCodeTypes = "notFound"
	AssertResponseCodeTypesMethodNotAllowed            AssertResponseCodeTypes = "methodNotAllowed"
	AssertResponseCodeTypesNotAcceptable               AssertResponseCodeTypes = "notAcceptable"
	AssertResponseCodeTypesProxyAuthenticationRequired AssertResponseCodeTypes = "proxyAuthenticationRequired"
	AssertResponseCodeTypesRequestTimeout              AssertResponseCodeTypes = "requestTimeout"
	AssertResponseCodeTypesConflict                    AssertResponseCodeTypes = "conflict"
	AssertResponseCodeTypesGone                        AssertResponseCodeTypes = "gone"
	AssertResponseCodeTypesLengthRequired              AssertResponseCodeTypes = "lengthRequired"
	AssertResponseCodeTypesPreconditionFailed          AssertResponseCodeTypes = "preconditionFailed"
	AssertResponseCodeTypesContentTooLarge             AssertResponseCodeTypes = "contentTooLarge"
	AssertResponseCodeTypesUriTooLong                  AssertResponseCodeTypes = "uriTooLong"
	AssertResponseCodeTypesUnsupportedMediaType        AssertResponseCodeTypes = "unsupportedMediaType"
	AssertResponseCodeTypesRangeNotSatisfiable         AssertResponseCodeTypes = "rangeNotSatisfiable"
	AssertResponseCodeTypesExpectationFailed           AssertResponseCodeTypes = "expectationFailed"
	AssertResponseCodeTypesMisdirectedRequest          AssertResponseCodeTypes = "misdirectedRequest"
	AssertResponseCodeTypesUnprocessableContent        AssertResponseCodeTypes = "unprocessableContent"
	AssertResponseCodeTypesUpgradeRequired             AssertResponseCodeTypes = "upgradeRequired"
	AssertResponseCodeTypesInternalServerError         AssertResponseCodeTypes = "internalServerError"
	AssertResponseCodeTypesNotImplemented              AssertResponseCodeTypes = "notImplemented"
	AssertResponseCodeTypesBadGateway                  AssertResponseCodeTypes = "badGateway"
	AssertResponseCodeTypesServiceUnavailable          AssertResponseCodeTypes = "serviceUnavailable"
	AssertResponseCodeTypesGatewayTimeout              AssertResponseCodeTypes = "gatewayTimeout"
	AssertResponseCodeTypesHttpVersionNotSupported     AssertResponseCodeTypes = "httpVersionNotSupported"
)

func AllAssertResponseCodeTypes() []AssertResponseCodeTypes {
	return []AssertResponseCodeTypes{
		AssertResponseCodeTypesContinue,
		AssertResponseCodeTypesSwitchingProtocols,
		AssertResponseCodeTypesOkay,
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
		AssertResponseCodeTypesUriTooLong,
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
		AssertResponseCodeTypesHttpVersionNotSupported,
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
	case AssertResponseCodeTypesOkay:
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
	case AssertResponseCodeTypesUriTooLong:
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
	case AssertResponseCodeTypesHttpVersionNotSupported:
		return "HTTP Version Not Supported"

	default:
		return "Unknown Address Use"
	}
}
