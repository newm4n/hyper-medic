package code

import (
	"fmt"
	"strings"
)

type DefinitionResourceTypes string

const (
	DefinitionResourceTypesActivityDefinition    DefinitionResourceTypes = "ActivityDefinition"
	DefinitionResourceTypesChargeItemDefinition  DefinitionResourceTypes = "ChargeItemDefinition"
	DefinitionResourceTypesClinicalUseDefinition DefinitionResourceTypes = "ClinicalUseDefinition"
	DefinitionResourceTypesEventDefinition       DefinitionResourceTypes = "EventDefinition"
	DefinitionResourceTypesMeasure               DefinitionResourceTypes = "Measure"
	DefinitionResourceTypesMessageDefinition     DefinitionResourceTypes = "MessageDefinition"
	DefinitionResourceTypesObservationDefinition DefinitionResourceTypes = "ObservationDefinition"
	DefinitionResourceTypesOperationDefinition   DefinitionResourceTypes = "OperationDefinition"
	DefinitionResourceTypesPlanDefinition        DefinitionResourceTypes = "PlanDefinition"
	DefinitionResourceTypesQuestionnaire         DefinitionResourceTypes = "Questionnaire"
	DefinitionResourceTypesRequirements          DefinitionResourceTypes = "Requirements"
	DefinitionResourceTypesSubscriptionTopic     DefinitionResourceTypes = "SubscriptionTopic"
	DefinitionResourceTypesTestPlan              DefinitionResourceTypes = "TestPlan"
	DefinitionResourceTypesTestScript            DefinitionResourceTypes = "TestScript"
)

func AllDefinitionResourceTypes() []DefinitionResourceTypes {
	return []DefinitionResourceTypes{
		DefinitionResourceTypesActivityDefinition,
		DefinitionResourceTypesChargeItemDefinition,
		DefinitionResourceTypesClinicalUseDefinition,
		DefinitionResourceTypesEventDefinition,
		DefinitionResourceTypesMeasure,
		DefinitionResourceTypesMessageDefinition,
		DefinitionResourceTypesObservationDefinition,
		DefinitionResourceTypesOperationDefinition,
		DefinitionResourceTypesPlanDefinition,
		DefinitionResourceTypesQuestionnaire,
		DefinitionResourceTypesRequirements,
		DefinitionResourceTypesSubscriptionTopic,
		DefinitionResourceTypesTestPlan,
		DefinitionResourceTypesTestScript,
	}
}

func FinDefinitionResourceTypes(filter string) []DefinitionResourceTypes {
	ret := make([]DefinitionResourceTypes, 0)
	for _, at := range AllDefinitionResourceTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DefinitionResourceTypes) ToString() {
	fmt.Println(au.String())
}
func (au DefinitionResourceTypes) String() string {
	switch au {
	case DefinitionResourceTypesActivityDefinition:
		return "ActivityDefinition"
	case DefinitionResourceTypesChargeItemDefinition:
		return "ChargeItemDefinition"
	case DefinitionResourceTypesClinicalUseDefinition:
		return "ClinicalUseDefinition"
	case DefinitionResourceTypesEventDefinition:
		return "EventDefinition"
	case DefinitionResourceTypesMeasure:
		return "Measure"
	case DefinitionResourceTypesMessageDefinition:
		return "MessageDefinition"
	case DefinitionResourceTypesObservationDefinition:
		return "ObservationDefinition"
	case DefinitionResourceTypesOperationDefinition:
		return "OperationDefinition"
	case DefinitionResourceTypesPlanDefinition:
		return "PlanDefinition"
	case DefinitionResourceTypesQuestionnaire:
		return "Questionnaire"
	case DefinitionResourceTypesRequirements:
		return "Requirements"
	case DefinitionResourceTypesSubscriptionTopic:
		return "SubscriptionTopic"
	case DefinitionResourceTypesTestPlan:
		return "TestPlan"
	case DefinitionResourceTypesTestScript:
		return "TestScript"
	default:
		return "Unknown Definition Resource Types"
	}
}
