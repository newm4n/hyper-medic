package code

import (
	"fmt"
	"strings"
)

type AuditEventType string

const (
	AuditEventTypeApplicationActivity             AuditEventType = "110100"
	AuditEventTypeAuditLogUsed                    AuditEventType = "110101"
	AuditEventTypeBeginTransferringDICOMInstances AuditEventType = "110102"
	AuditEventTypeDICOMInstancesAccessed          AuditEventType = "110103"
	AuditEventTypeDICOMInstancesTransferred       AuditEventType = "110104"
	AuditEventTypeDICOMStudyDeleted               AuditEventType = "110105"
	AuditEventTypeExport                          AuditEventType = "110106"
	AuditEventTypeImport                          AuditEventType = "110107"
	AuditEventTypeNetworkEntry                    AuditEventType = "110108"
	AuditEventTypeOrderRecord                     AuditEventType = "110109"
	AuditEventTypePatientRecord                   AuditEventType = "110110"
	AuditEventTypeProcedureRecord                 AuditEventType = "110111"
	AuditEventTypeQuery                           AuditEventType = "110112"
	AuditEventTypeSecurityAlert                   AuditEventType = "110113"
	AuditEventTypeUserAuthentication              AuditEventType = "110114"
	AuditEventTypeRESTfulOperation                AuditEventType = "rest"
	AuditEventTypeHL7v2Operation                  AuditEventType = "hl7-v2"
	AuditEventTypeHL7v3Operation                  AuditEventType = "hl7-v3"
	AuditEventTypeADocumentOperation              AuditEventType = "document"
	AuditEventTypeAnOperationonotherObjects       AuditEventType = "object"
)

func AllAuditEventType() []AuditEventType {
	return []AuditEventType{
		AuditEventTypeApplicationActivity,
		AuditEventTypeAuditLogUsed,
		AuditEventTypeBeginTransferringDICOMInstances,
		AuditEventTypeDICOMInstancesAccessed,
		AuditEventTypeDICOMInstancesTransferred,
		AuditEventTypeDICOMStudyDeleted,
		AuditEventTypeExport,
		AuditEventTypeImport,
		AuditEventTypeNetworkEntry,
		AuditEventTypeOrderRecord,
		AuditEventTypePatientRecord,
		AuditEventTypeProcedureRecord,
		AuditEventTypeQuery,
		AuditEventTypeSecurityAlert,
		AuditEventTypeUserAuthentication,
		AuditEventTypeRESTfulOperation,
		AuditEventTypeHL7v2Operation,
		AuditEventTypeHL7v3Operation,
		AuditEventTypeADocumentOperation,
		AuditEventTypeAnOperationonotherObjects,
	}
}

func FindAuditEventType(filter string) []AuditEventType {
	ret := make([]AuditEventType, 0)
	for _, at := range AllAuditEventType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AuditEventType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AuditEventType) String() string {
	switch cpt {
	case AuditEventTypeApplicationActivity:
		return "Application Activity"
	case AuditEventTypeAuditLogUsed:
		return "Audit Log Used"
	case AuditEventTypeBeginTransferringDICOMInstances:
		return "Begin Transferring DICOM Instances"
	case AuditEventTypeDICOMInstancesAccessed:
		return "DICOM Instances Accessed"
	case AuditEventTypeDICOMInstancesTransferred:
		return "DICOM Instances Transferred"
	case AuditEventTypeDICOMStudyDeleted:
		return "DICOM Study Deleted"
	case AuditEventTypeExport:
		return "Export"
	case AuditEventTypeImport:
		return "Import"
	case AuditEventTypeNetworkEntry:
		return "Network Entry"
	case AuditEventTypeOrderRecord:
		return "Order Record"
	case AuditEventTypePatientRecord:
		return "Patient Record"
	case AuditEventTypeProcedureRecord:
		return "Procedure Record"
	case AuditEventTypeQuery:
		return "Query"
	case AuditEventTypeSecurityAlert:
		return "Security Alert"
	case AuditEventTypeUserAuthentication:
		return "User Authentication"
	case AuditEventTypeRESTfulOperation:
		return "RESTful Operation"
	case AuditEventTypeHL7v2Operation:
		return "HL7 v2 Operation"
	case AuditEventTypeHL7v3Operation:
		return "HL7 v3 Operation"
	case AuditEventTypeADocumentOperation:
		return "A Document Operation"
	case AuditEventTypeAnOperationonotherObjects:
		return "An Operation on other Objects"

	default:
		return "Unknown Action Grouping Behavior"
	}
}
