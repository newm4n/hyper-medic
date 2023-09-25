package code

import (
	"fmt"
	"strings"
)

type ActionReasonCode int

const (
	ActionReasonCodeOffPathway ActionReasonCode = iota
	ActionReasonCodeRiskAssessment
	ActionReasonCodeCareGapDetected
	ActionReasonCodeDrugFrugInteraction
	ActionReasonCodeQualityMeasure
)

func AllActionReasonCode() []ActionReasonCode {
	return []ActionReasonCode{
		ActionReasonCodeOffPathway,
		ActionReasonCodeRiskAssessment,
		ActionReasonCodeCareGapDetected,
		ActionReasonCodeDrugFrugInteraction,
		ActionReasonCodeQualityMeasure,
	}
}

func FindActionReasonCode(filter string) []ActionReasonCode {
	ret := make([]ActionReasonCode, 0)
	for _, at := range AllActionReasonCode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionReasonCode) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionReasonCode) String() string {
	switch cpt {
	case ActionReasonCodeOffPathway:
		return "Off pathway"
	case ActionReasonCodeRiskAssessment:
		return "Risk assessment"
	case ActionReasonCodeCareGapDetected:
		return "Risk assessment"
	case ActionReasonCodeDrugFrugInteraction:
		return "Drug-drug interaction"
	case ActionReasonCodeQualityMeasure:
		return "Quality measure"
	default:
		return "Unknown Action Required Behavior"
	}
}

/*
Off pathway
Risk assessment
Care gap detected
Drug-drug interaction
Quality measure
*/
