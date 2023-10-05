package code

import (
	"fmt"
	"strings"
)

type AdditionalBindingPurpose string

const (
	AdditionalBindingPurposeMaximum     AdditionalBindingPurpose = "maximum"
	AdditionalBindingPurposeMinimum     AdditionalBindingPurpose = "minimum"
	AdditionalBindingPurposeRequired    AdditionalBindingPurpose = "required"
	AdditionalBindingPurposeConformance AdditionalBindingPurpose = "extensible"
	AdditionalBindingPurposeCandidate   AdditionalBindingPurpose = "candidate"
	AdditionalBindingPurposeCurrent     AdditionalBindingPurpose = "current"
	AdditionalBindingPurposePreferred   AdditionalBindingPurpose = "preferred"
	AdditionalBindingPurposeUISuggested AdditionalBindingPurpose = "ui"
	AdditionalBindingPurposeStarter     AdditionalBindingPurpose = "started"
	AdditionalBindingPurposeComponent   AdditionalBindingPurpose = "component"
)

func AllAdditionalBindingPurpose() []AdditionalBindingPurpose {
	return []AdditionalBindingPurpose{
		AdditionalBindingPurposeMaximum,
		AdditionalBindingPurposeMinimum,
		AdditionalBindingPurposeRequired,
		AdditionalBindingPurposeConformance,
		AdditionalBindingPurposeCandidate,
		AdditionalBindingPurposeCurrent,
		AdditionalBindingPurposePreferred,
		AdditionalBindingPurposeUISuggested,
		AdditionalBindingPurposeStarter,
		AdditionalBindingPurposeComponent,
	}
}

func FindAdditionalBindingPurpose(filter string) []AdditionalBindingPurpose {
	ret := make([]AdditionalBindingPurpose, 0)
	for _, at := range AllAdditionalBindingPurpose() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdditionalBindingPurpose) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdditionalBindingPurpose) String() string {
	switch cpt {
	case AdditionalBindingPurposeMaximum:
		return "Maximum Binding"
	case AdditionalBindingPurposeMinimum:
		return "Minimum Binding"
	case AdditionalBindingPurposeRequired:
		return "Required Binding"
	case AdditionalBindingPurposeConformance:
		return "Conformance Binding"
	case AdditionalBindingPurposeCandidate:
		return "Candidate Binding"
	case AdditionalBindingPurposeCurrent:
		return "Current Binding"
	case AdditionalBindingPurposePreferred:
		return "Preferred Binding"
	case AdditionalBindingPurposeUISuggested:
		return "UI Suggested Binding"
	case AdditionalBindingPurposeStarter:
		return "Starter Binding"
	case AdditionalBindingPurposeComponent:
		return "Component Binding"
	default:
		return "Unknown Adverse Event Actuality"
	}
}

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
