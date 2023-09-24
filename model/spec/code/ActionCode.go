package code

type ActionCode int

const (
	ActionCodeSendMessage ActionCode = iota
	ActionCodeCollectInformation
	ActionCodePrescribeMedication
	ActionCodeRecommendAnImmunization
	ActionCodeOrderService
	ActionCodeProposeDiagnosis
	ActionCodeRecordDetectedIssue
	ActionCodeRecordInference
	ActionCodeReportFlag
)

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
