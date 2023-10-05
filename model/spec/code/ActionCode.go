package code

import (
	"fmt"
	"strings"
)

type ActionCode string

const (
	ActionCodeSendMessage             ActionCode = "send-message"
	ActionCodeCollectInformation      ActionCode = "collect-information"
	ActionCodePrescribeMedication     ActionCode = "prescribe-medication"
	ActionCodeRecommendAnImmunization ActionCode = "recommend-immunization"
	ActionCodeOrderService            ActionCode = "order-service"
	ActionCodeProposeDiagnosis        ActionCode = "propose-diagnosis"
	ActionCodeRecordDetectedIssue     ActionCode = "record-detected-issue"
	ActionCodeRecordInference         ActionCode = "record-inference"
	ActionCodeReportFlag              ActionCode = "report-flag"
)

func AllActionCode() []ActionCode {
	return []ActionCode{
		ActionCodeSendMessage,
		ActionCodeCollectInformation,
		ActionCodePrescribeMedication,
		ActionCodeRecommendAnImmunization,
		ActionCodeOrderService,
		ActionCodeProposeDiagnosis,
		ActionCodeRecordDetectedIssue,
		ActionCodeRecordInference,
		ActionCodeReportFlag,
	}
}

func FindActionCode(filter string) []ActionCode {
	ret := make([]ActionCode, 0)
	for _, at := range AllActionCode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionCode) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionCode) String() string {
	switch cpt {
	case ActionCodeSendMessage:
		return "Send a message"
	case ActionCodeCollectInformation:
		return "Collect information"
	case ActionCodePrescribeMedication:
		return "Prescribe a medication"
	case ActionCodeRecommendAnImmunization:
		return "Recommend an immunization"
	case ActionCodeOrderService:
		return "Order a service"
	case ActionCodeProposeDiagnosis:
		return "Propose a diagnosis"
	case ActionCodeRecordDetectedIssue:
		return "Record a detected issue"
	case ActionCodeRecordInference:
		return "Record an inference"
	case ActionCodeReportFlag:
		return "Report a flag"
	default:
		return "Unknown Action Code"
	}
}

/*
Send a message
Collect information
Prescribe a medication
Recommend an immunization
Order a service
Propose a diagnosis
Record a detected issue
Record an inference
Report a flag

*/
