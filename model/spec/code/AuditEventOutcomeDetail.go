package code

import (
	"fmt"
	"strings"
)

type AuditEventOutcomeDetail string

const (
	AuditEventOutcomeDetailDeleteMultipleMatches       AuditEventOutcomeDetail = "DELETE_MULTIPLE_MATCHES"
	AuditEventOutcomeDetailMsgAuthRequired             AuditEventOutcomeDetail = "MSG_AUTH_REQUIRED"
	AuditEventOutcomeDetailMsgBadFormat                AuditEventOutcomeDetail = "MSG_BAD_FORMAT"
	AuditEventOutcomeDetailMsgBadSyntax                AuditEventOutcomeDetail = "MSG_BAD_SYNTAX"
	AuditEventOutcomeDetailMsgCantParseContent         AuditEventOutcomeDetail = "MSG_CANT_PARSE_CONTENT"
	AuditEventOutcomeDetailMsgCantParseRoot            AuditEventOutcomeDetail = "MSG_CANT_PARSE_ROOT"
	AuditEventOutcomeDetailMsgCreated                  AuditEventOutcomeDetail = "MSG_CREATED"
	AuditEventOutcomeDetailMsgDateFormat               AuditEventOutcomeDetail = "MSG_DATE_FORMAT"
	AuditEventOutcomeDetailMsgDeleted                  AuditEventOutcomeDetail = "MSG_DELETED"
	AuditEventOutcomeDetailMsgDeletedDone              AuditEventOutcomeDetail = "MSG_DELETED_DONE"
	AuditEventOutcomeDetailMsgDeletedId                AuditEventOutcomeDetail = "MSG_DELETED_ID"
	AuditEventOutcomeDetailMsgDuplicateId              AuditEventOutcomeDetail = "MSG_DUPLICATE_ID"
	AuditEventOutcomeDetailMsgErrorParsing             AuditEventOutcomeDetail = "MSG_ERROR_PARSING"
	AuditEventOutcomeDetailMsgExternalFail             AuditEventOutcomeDetail = "MSG_EXTERNAL_FAIL"
	AuditEventOutcomeDetailMsgIdInvalid                AuditEventOutcomeDetail = "MSG_ID_INVALID"
	AuditEventOutcomeDetailMsgIdTooLong                AuditEventOutcomeDetail = "MSG_ID_TOO_LONG"
	AuditEventOutcomeDetailMsgInvalidId                AuditEventOutcomeDetail = "MSG_INVALID_ID"
	AuditEventOutcomeDetailMsgJsonObject               AuditEventOutcomeDetail = "MSG_JSON_OBJECT"
	AuditEventOutcomeDetailMsgLocalFail                AuditEventOutcomeDetail = "MSG_LOCAL_FAIL"
	AuditEventOutcomeDetailMsgNoExist                  AuditEventOutcomeDetail = "MSG_NO_EXIST"
	AuditEventOutcomeDetailMsgNoMatch                  AuditEventOutcomeDetail = "MSG_NO_MATCH"
	AuditEventOutcomeDetailMsgNoModule                 AuditEventOutcomeDetail = "MSG_NO_MODULE"
	AuditEventOutcomeDetailMsgNoSummary                AuditEventOutcomeDetail = "MSG_NO_SUMMARY"
	AuditEventOutcomeDetailMsgOpNotAllowed             AuditEventOutcomeDetail = "MSG_OP_NOT_ALLOWED"
	AuditEventOutcomeDetailMsgParamChained             AuditEventOutcomeDetail = "MSG_PARAM_CHAINED"
	AuditEventOutcomeDetailMsgParamInvalid             AuditEventOutcomeDetail = "MSG_PARAM_INVALID"
	AuditEventOutcomeDetailMsgParamModifierInvalid     AuditEventOutcomeDetail = "MSG_PARAM_MODIFIER_INVALID"
	AuditEventOutcomeDetailMsgParamNoRepeat            AuditEventOutcomeDetail = "MSG_PARAM_NO_REPEAT"
	AuditEventOutcomeDetailMsgParamUnknown             AuditEventOutcomeDetail = "MSG_PARAM_UNKNOWN"
	AuditEventOutcomeDetailMsgResourceExampleProtected AuditEventOutcomeDetail = "MSG_RESOURCE_EXAMPLE_PROTECTED"
	AuditEventOutcomeDetailMsgResourceIdFail           AuditEventOutcomeDetail = "MSG_RESOURCE_ID_FAIL"
	AuditEventOutcomeDetailMsgResourceIdMismatch       AuditEventOutcomeDetail = "MSG_RESOURCE_ID_MISMATCH"
	AuditEventOutcomeDetailMsgResourceIdMissing        AuditEventOutcomeDetail = "MSG_RESOURCE_ID_MISSING"
	AuditEventOutcomeDetailMsgResourceNotAllowed       AuditEventOutcomeDetail = "MSG_RESOURCE_NOT_ALLOWED"
	AuditEventOutcomeDetailMsgResourceRequired         AuditEventOutcomeDetail = "MSG_RESOURCE_REQUIRED"
	AuditEventOutcomeDetailMsgResourceTypeMismatch     AuditEventOutcomeDetail = "MSG_RESOURCE_TYPE_MISMATCH"
	AuditEventOutcomeDetailMsgSortUnknown              AuditEventOutcomeDetail = "MSG_SORT_UNKNOWN"
	AuditEventOutcomeDetailMsgTransactionDuplicateId   AuditEventOutcomeDetail = "MSG_TRANSACTION_DUPLICATE_ID"
	AuditEventOutcomeDetailMsgTransactionMissingId     AuditEventOutcomeDetail = "MSG_TRANSACTION_MISSING_ID"
	AuditEventOutcomeDetailMsgUnhandledNodeType        AuditEventOutcomeDetail = "MSG_UNHANDLED_NODE_TYPE"
	AuditEventOutcomeDetailMsgUnknownContent           AuditEventOutcomeDetail = "MSG_UNKNOWN_CONTENT"
	AuditEventOutcomeDetailMsgUnknownOperation         AuditEventOutcomeDetail = "MSG_UNKNOWN_OPERATION"
	AuditEventOutcomeDetailMsgUnknownType              AuditEventOutcomeDetail = "MSG_UNKNOWN_TYPE"
	AuditEventOutcomeDetailMsgUpdated                  AuditEventOutcomeDetail = "MSG_UPDATED"
	AuditEventOutcomeDetailMsgVersionAware             AuditEventOutcomeDetail = "MSG_VERSION_AWARE"
	AuditEventOutcomeDetailMsgVersionAwareConflict     AuditEventOutcomeDetail = "MSG_VERSION_AWARE_CONFLICT"
	AuditEventOutcomeDetailMsgVersionAwareUrl          AuditEventOutcomeDetail = "MSG_VERSION_AWARE_URL"
	AuditEventOutcomeDetailMsgWrongNs                  AuditEventOutcomeDetail = "MSG_WRONG_NS"
	AuditEventOutcomeDetailSearchMultiple              AuditEventOutcomeDetail = "SEARCH_MULTIPLE"
	AuditEventOutcomeDetailSearchNone                  AuditEventOutcomeDetail = "SEARCH_NONE"
	AuditEventOutcomeDetailUpdateMultipleMatches       AuditEventOutcomeDetail = "UPDATE_MULTIPLE_MATCHES"
)

func AllAuditEventOutcomeDetail() []AuditEventOutcomeDetail {
	return []AuditEventOutcomeDetail{
		AuditEventOutcomeDetailDeleteMultipleMatches,
		AuditEventOutcomeDetailMsgAuthRequired,
		AuditEventOutcomeDetailMsgBadFormat,
		AuditEventOutcomeDetailMsgBadSyntax,
		AuditEventOutcomeDetailMsgCantParseContent,
		AuditEventOutcomeDetailMsgCantParseRoot,
		AuditEventOutcomeDetailMsgCreated,
		AuditEventOutcomeDetailMsgDateFormat,
		AuditEventOutcomeDetailMsgDeleted,
		AuditEventOutcomeDetailMsgDeletedDone,
		AuditEventOutcomeDetailMsgDeletedId,
		AuditEventOutcomeDetailMsgDuplicateId,
		AuditEventOutcomeDetailMsgErrorParsing,
		AuditEventOutcomeDetailMsgExternalFail,
		AuditEventOutcomeDetailMsgIdInvalid,
		AuditEventOutcomeDetailMsgIdTooLong,
		AuditEventOutcomeDetailMsgInvalidId,
		AuditEventOutcomeDetailMsgJsonObject,
		AuditEventOutcomeDetailMsgLocalFail,
		AuditEventOutcomeDetailMsgNoExist,
		AuditEventOutcomeDetailMsgNoMatch,
		AuditEventOutcomeDetailMsgNoModule,
		AuditEventOutcomeDetailMsgNoSummary,
		AuditEventOutcomeDetailMsgOpNotAllowed,
		AuditEventOutcomeDetailMsgParamChained,
		AuditEventOutcomeDetailMsgParamInvalid,
		AuditEventOutcomeDetailMsgParamModifierInvalid,
		AuditEventOutcomeDetailMsgParamNoRepeat,
		AuditEventOutcomeDetailMsgParamUnknown,
		AuditEventOutcomeDetailMsgResourceExampleProtected,
		AuditEventOutcomeDetailMsgResourceIdFail,
		AuditEventOutcomeDetailMsgResourceIdMismatch,
		AuditEventOutcomeDetailMsgResourceIdMissing,
		AuditEventOutcomeDetailMsgResourceNotAllowed,
		AuditEventOutcomeDetailMsgResourceRequired,
		AuditEventOutcomeDetailMsgResourceTypeMismatch,
		AuditEventOutcomeDetailMsgSortUnknown,
		AuditEventOutcomeDetailMsgTransactionDuplicateId,
		AuditEventOutcomeDetailMsgTransactionMissingId,
		AuditEventOutcomeDetailMsgUnhandledNodeType,
		AuditEventOutcomeDetailMsgUnknownContent,
		AuditEventOutcomeDetailMsgUnknownOperation,
		AuditEventOutcomeDetailMsgUnknownType,
		AuditEventOutcomeDetailMsgUpdated,
		AuditEventOutcomeDetailMsgVersionAware,
		AuditEventOutcomeDetailMsgVersionAwareConflict,
		AuditEventOutcomeDetailMsgVersionAwareUrl,
		AuditEventOutcomeDetailMsgWrongNs,
		AuditEventOutcomeDetailSearchMultiple,
		AuditEventOutcomeDetailSearchNone,
		AuditEventOutcomeDetailUpdateMultipleMatches,
	}
}

func FindAuditEventOutcomeDetail(filter string) []AuditEventOutcomeDetail {
	ret := make([]AuditEventOutcomeDetail, 0)
	for _, at := range AllAuditEventOutcomeDetail() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AuditEventOutcomeDetail) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AuditEventOutcomeDetail) String() string {
	switch cpt {
	case AuditEventOutcomeDetailDeleteMultipleMatches:
		return "Error: Multiple matches exist for the conditional delete"
	case AuditEventOutcomeDetailMsgAuthRequired:
		return "You must authenticate before you can use this service"
	case AuditEventOutcomeDetailMsgBadFormat:
		return "Bad Syntax: \"%s\" must be a %s'"
	case AuditEventOutcomeDetailMsgBadSyntax:
		return "Bad Syntax in %s"
	case AuditEventOutcomeDetailMsgCantParseContent:
		return "Unable to parse feed (entry content type = \"%s\")"
	case AuditEventOutcomeDetailMsgCantParseRoot:
		return "Unable to parse feed (root element name = \"%s\")"
	case AuditEventOutcomeDetailMsgCreated:
		return "New resource created"
	case AuditEventOutcomeDetailMsgDateFormat:
		return "The Date value %s is not in the correct format (Xml Date Format required)"
	case AuditEventOutcomeDetailMsgDeleted:
		return "This resource has been deleted"
	case AuditEventOutcomeDetailMsgDeletedDone:
		return "Resource deleted"
	case AuditEventOutcomeDetailMsgDeletedId:
		return "The resource \"%s\" has been deleted"
	case AuditEventOutcomeDetailMsgDuplicateId:
		return "Duplicate Id %s for resource type %s"
	case AuditEventOutcomeDetailMsgErrorParsing:
		return "Error parsing resource Xml (%s)"
	case AuditEventOutcomeDetailMsgExternalFail:
		return "Unable to resolve external reference to resource %s"
	case AuditEventOutcomeDetailMsgIdInvalid:
		return "Id \"%s\" has an invalid character \"%s\""
	case AuditEventOutcomeDetailMsgIdTooLong:
		return "Id \"%s\" too long (length limit 36)"
	case AuditEventOutcomeDetailMsgInvalidId:
		return "Id not accepted"
	case AuditEventOutcomeDetailMsgJsonObject:
		return "Json Source for a resource should start with an object"
	case AuditEventOutcomeDetailMsgLocalFail:
		return "Unable to resolve local reference to resource %s"
	case AuditEventOutcomeDetailMsgNoExist:
		return "Resource Id \"%s\" does not exist"
	case AuditEventOutcomeDetailMsgNoMatch:
		return "No Resource found matching the query \"%s\""
	case AuditEventOutcomeDetailMsgNoModule:
		return "No module could be found to handle the request \"%s\""
	case AuditEventOutcomeDetailMsgNoSummary:
		return "No Summary for this resource"
	case AuditEventOutcomeDetailMsgOpNotAllowed:
		return "Operation %s not allowed for resource %s (due to local configuration)"
	case AuditEventOutcomeDetailMsgParamChained:
		return "Unknown chained parameter name \"%s\""
	case AuditEventOutcomeDetailMsgParamInvalid:
		return "Parameter \"%s\" content is invalid"
	case AuditEventOutcomeDetailMsgParamModifierInvalid:
		return "Parameter \"%s\" modifier is invalid"
	case AuditEventOutcomeDetailMsgParamNoRepeat:
		return "Parameter \"%s\" is not allowed to repeat"
	case AuditEventOutcomeDetailMsgParamUnknown:
		return "Parameter \"%s\" not understood"
	case AuditEventOutcomeDetailMsgResourceExampleProtected:
		return "Resources with identity \"example\" cannot be deleted (for testing/training purposes)"
	case AuditEventOutcomeDetailMsgResourceIdFail:
		return "unable to allocate resource id"
	case AuditEventOutcomeDetailMsgResourceIdMismatch:
		return "Resource Id Mismatch"
	case AuditEventOutcomeDetailMsgResourceIdMissing:
		return "Resource Id Missing"
	case AuditEventOutcomeDetailMsgResourceNotAllowed:
		return "Not allowed to submit a resource for this operation"
	case AuditEventOutcomeDetailMsgResourceRequired:
		return "A resource is required"
	case AuditEventOutcomeDetailMsgResourceTypeMismatch:
		return "Resource Type Mismatch"
	case AuditEventOutcomeDetailMsgSortUnknown:
		return "Unknown sort parameter name \"%s\""
	case AuditEventOutcomeDetailMsgTransactionDuplicateId:
		return "Duplicate Identifier in transaction: %s"
	case AuditEventOutcomeDetailMsgTransactionMissingId:
		return "Missing Identifier in transaction - an entry.id must be provided"
	case AuditEventOutcomeDetailMsgUnhandledNodeType:
		return "Unhandled xml node type \"%s\""
	case AuditEventOutcomeDetailMsgUnknownContent:
		return "Unknown Content (%s) at %s"
	case AuditEventOutcomeDetailMsgUnknownOperation:
		return "unknown FHIR http operation"
	case AuditEventOutcomeDetailMsgUnknownType:
		return "Resource Type \"%s\" not recognised"
	case AuditEventOutcomeDetailMsgUpdated:
		return "existing resource updated"
	case AuditEventOutcomeDetailMsgVersionAware:
		return "Version aware updates are required for this resource"
	case AuditEventOutcomeDetailMsgVersionAwareConflict:
		return "Update Conflict (server current version = \"%s\", client version referenced = \"%s\")"
	case AuditEventOutcomeDetailMsgVersionAwareUrl:
		return "Version specific URL not recognised"
	case AuditEventOutcomeDetailMsgWrongNs:
		return "This does not appear to be a FHIR element or resource (wrong namespace \"%s\")"
	case AuditEventOutcomeDetailSearchMultiple:
		return "Error: Multiple matches exist for %s search parameters \"%s\""
	case AuditEventOutcomeDetailSearchNone:
		return "Error: no processable search found for %s search parameters \"%s\""
	case AuditEventOutcomeDetailUpdateMultipleMatches:
		return "Error: Multiple matches exist for the conditional update"

	default:
		return "Unknown Action Grouping Behavior"
	}
}
