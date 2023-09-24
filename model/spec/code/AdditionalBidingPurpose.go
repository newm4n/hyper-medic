package code

type AdditionalBindingPurpose int

const (
	AdditionalBindingPurposeMaximum AdditionalBindingPurpose = iota
	AdditionalBindingPurposeMinimum
	AdditionalBindingPurposeRequired
	AdditionalBindingPurposeConformance
	AdditionalBindingPurposeCandidate
	AdditionalBindingPurposeCurrent
	AdditionalBindingPurposePreferred
	AdditionalBindingPurposeUISuggested
	AdditionalBindingPurposeStarter
	AdditionalBindingPurposeComponent
)

/*
Maximum Binding
Minimum Binding
Required Binding
Conformance Binding
Candidate Binding
Current Binding
Preferred Binding
UI Suggested Binding
Starter Binding
Component Binding
*/
