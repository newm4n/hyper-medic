package code

type ActionReasonCode int

const (
	ActionReasonCodeOffPathway ActionReasonCode = iota
	ActionReasonCodeRiskAssessment
	ActionReasonCodeCareGapDetected
	ActionReasonCodeDrugFrugInteraction
	ActionReasonCodeQualityMeasure
)

/*
Off pathway
Risk assessment
Care gap detected
Drug-drug interaction
Quality measure
*/
