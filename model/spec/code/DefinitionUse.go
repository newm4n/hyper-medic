package code

import (
	"fmt"
	"strings"
)

type DefinitionUse string

const (
	DefinitionUseFHIRStructure       DefinitionUse = "fhir-structure"
	DefinitionUseCustomResource      DefinitionUse = "custom-resource"
	DefinitionUseDomainAnalysisModel DefinitionUse = "dam"
	DefinitionUseWireFormat          DefinitionUse = "wire-format"
	DefinitionUseArchetype           DefinitionUse = "archetype"
	DefinitionUseTemplate            DefinitionUse = "template"
)

func AllDefinitionUse() []DefinitionUse {
	return []DefinitionUse{
		DefinitionUseFHIRStructure,
		DefinitionUseCustomResource,
		DefinitionUseDomainAnalysisModel,
		DefinitionUseWireFormat,
		DefinitionUseArchetype,
		DefinitionUseTemplate,
	}
}

func FindDefinitionUse(filter string) []DefinitionUse {
	ret := make([]DefinitionUse, 0)
	for _, at := range AllDefinitionUse() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DefinitionUse) ToString() {
	fmt.Println(au.String())
}
func (au DefinitionUse) String() string {
	switch au {
	case DefinitionUseFHIRStructure:
		return "FHIR Structure"
	case DefinitionUseCustomResource:
		return "Custom Resource"
	case DefinitionUseDomainAnalysisModel:
		return "Domain Analysis Model (DAM)"
	case DefinitionUseWireFormat:
		return "Wire Format"
	case DefinitionUseArchetype:
		return "Domain Analysis Model"
	case DefinitionUseTemplate:
		return "Template"
	default:
		return "Unknown Constraint Severity"
	}
}
