package code

import (
	"fmt"
	"strings"
)

type ResourceTypes string

const (
	ResourceTypesAccount                            ResourceTypes = "Account"
	ResourceTypesActivityDefinition                 ResourceTypes = "ActivityDefinition"
	ResourceTypesActorDefinition                    ResourceTypes = "ActorDefinition"
	ResourceTypesAdministrableProductDefinition     ResourceTypes = "AdministrableProductDefinition"
	ResourceTypesAdverseEvent                       ResourceTypes = "AdverseEvent"
	ResourceTypesAllergyIntolerance                 ResourceTypes = "AllergyIntolerance"
	ResourceTypesAppointment                        ResourceTypes = "Appointment"
	ResourceTypesAppointmentResponse                ResourceTypes = "AppointmentResponse"
	ResourceTypesArtifactAssessment                 ResourceTypes = "ArtifactAssessment"
	ResourceTypesAuditEvent                         ResourceTypes = "AuditEvent"
	ResourceTypesBasic                              ResourceTypes = "Basic"
	ResourceTypesBinary                             ResourceTypes = "Binary"
	ResourceTypesBiologicallyDerivedProduct         ResourceTypes = "BiologicallyDerivedProduct"
	ResourceTypesBiologicallyDerivedProductDispense ResourceTypes = "BiologicallyDerivedProductDispense"
	ResourceTypesBodyStructure                      ResourceTypes = "BodyStructure"
	ResourceTypesBundle                             ResourceTypes = "Bundle"
	ResourceTypesCapabilityStatement                ResourceTypes = "CapabilityStatement"
	ResourceTypesCarePlan                           ResourceTypes = "CarePlan"
	ResourceTypesCareTeam                           ResourceTypes = "CareTeam"
	ResourceTypesChargeItem                         ResourceTypes = "ChargeItem"
	ResourceTypesChargeItemDefinition               ResourceTypes = "ChargeItemDefinition"
	ResourceTypesCitation                           ResourceTypes = "Citation"
	ResourceTypesClaim                              ResourceTypes = "Claim"
	ResourceTypesClaimResponse                      ResourceTypes = "ClaimResponse"
	ResourceTypesClinicalImpression                 ResourceTypes = "ClinicalImpression"
	ResourceTypesClinicalUseDefinition              ResourceTypes = "ClinicalUseDefinition"
	ResourceTypesCodeSystem                         ResourceTypes = "CodeSystem"
	ResourceTypesCommunication                      ResourceTypes = "Communication"
	ResourceTypesCommunicationRequest               ResourceTypes = "CommunicationRequest"
	ResourceTypesCompartmentDefinition              ResourceTypes = "CompartmentDefinition"
	ResourceTypesComposition                        ResourceTypes = "Composition"
	ResourceTypesConceptMap                         ResourceTypes = "ConceptMap"
	ResourceTypesCondition                          ResourceTypes = "Condition"
	ResourceTypesConditionDefinition                ResourceTypes = "ConditionDefinition"
	ResourceTypesConsent                            ResourceTypes = "Consent"
	ResourceTypesContract                           ResourceTypes = "Contract"
	ResourceTypesCoverage                           ResourceTypes = "Coverage"
	ResourceTypesCoverageEligibilityRequest         ResourceTypes = "CoverageEligibilityRequest"
	ResourceTypesCoverageEligibilityResponse        ResourceTypes = "CoverageEligibilityResponse"
	ResourceTypesDetectedIssue                      ResourceTypes = "DetectedIssue"
	ResourceTypesDevice                             ResourceTypes = "Device"
	ResourceTypesDeviceAssociation                  ResourceTypes = "DeviceAssociation"
	ResourceTypesDeviceDefinition                   ResourceTypes = "DeviceDefinition"
	ResourceTypesDeviceDispense                     ResourceTypes = "DeviceDispense"
	ResourceTypesDeviceMetric                       ResourceTypes = "DeviceMetric"
	ResourceTypesDeviceRequest                      ResourceTypes = "DeviceRequest"
	ResourceTypesDeviceUsage                        ResourceTypes = "DeviceUsage"
	ResourceTypesDiagnosticReport                   ResourceTypes = "DiagnosticReport"
	ResourceTypesDocumentReference                  ResourceTypes = "DocumentReference"
	ResourceTypesEncounter                          ResourceTypes = "Encounter"
	ResourceTypesEncounterHistory                   ResourceTypes = "EncounterHistory"
	ResourceTypesEndpoint                           ResourceTypes = "Endpoint"
	ResourceTypesEnrollmentRequest                  ResourceTypes = "EnrollmentRequest"
	ResourceTypesEnrollmentResponse                 ResourceTypes = "EnrollmentResponse"
	ResourceTypesEpisodeOfCare                      ResourceTypes = "EpisodeOfCare"
	ResourceTypesEventDefinition                    ResourceTypes = "EventDefinition"
	ResourceTypesEvidence                           ResourceTypes = "Evidence"
	ResourceTypesEvidenceReport                     ResourceTypes = "EvidenceReport"
	ResourceTypesEvidenceVariable                   ResourceTypes = "EvidenceVariable"
	ResourceTypesExampleScenario                    ResourceTypes = "ExampleScenario"
	ResourceTypesExplanationOfBenefit               ResourceTypes = "ExplanationOfBenefit"
	ResourceTypesFamilyMemberHistory                ResourceTypes = "FamilyMemberHistory"
	ResourceTypesFlag                               ResourceTypes = "Flag"
	ResourceTypesFormularyItem                      ResourceTypes = "FormularyItem"
	ResourceTypesGenomicStudy                       ResourceTypes = "GenomicStudy"
	ResourceTypesGoal                               ResourceTypes = "Goal"
	ResourceTypesGraphDefinition                    ResourceTypes = "GraphDefinition"
	ResourceTypesGroup                              ResourceTypes = "Group"
	ResourceTypesGuidanceResponse                   ResourceTypes = "GuidanceResponse"
	ResourceTypesHealthcareService                  ResourceTypes = "HealthcareService"
	ResourceTypesImagingSelection                   ResourceTypes = "ImagingSelection"
	ResourceTypesImagingStudy                       ResourceTypes = "ImagingStudy"
	ResourceTypesImmunization                       ResourceTypes = "Immunization"
	ResourceTypesImmunizationEvaluation             ResourceTypes = "ImmunizationEvaluation"
	ResourceTypesImmunizationRecommendation         ResourceTypes = "ImmunizationRecommendation"
	ResourceTypesImplementationGuide                ResourceTypes = "ImplementationGuide"
	ResourceTypesIngredient                         ResourceTypes = "Ingredient"
	ResourceTypesInsurancePlan                      ResourceTypes = "InsurancePlan"
	ResourceTypesInventoryItem                      ResourceTypes = "InventoryItem"
	ResourceTypesInventoryReport                    ResourceTypes = "InventoryReport"
	ResourceTypesInvoice                            ResourceTypes = "Invoice"
	ResourceTypesLibrary                            ResourceTypes = "Library"
	ResourceTypesLinkage                            ResourceTypes = "Linkage"
	ResourceTypesList                               ResourceTypes = "List"
	ResourceTypesLocation                           ResourceTypes = "Location"
	ResourceTypesManufacturedItemDefinition         ResourceTypes = "ManufacturedItemDefinition"
	ResourceTypesMeasure                            ResourceTypes = "Measure"
	ResourceTypesMeasureReport                      ResourceTypes = "MeasureReport"
	ResourceTypesMedication                         ResourceTypes = "Medication"
	ResourceTypesMedicationAdministration           ResourceTypes = "MedicationAdministration"
	ResourceTypesMedicationDispense                 ResourceTypes = "MedicationDispense"
	ResourceTypesMedicationKnowledge                ResourceTypes = "MedicationKnowledge"
	ResourceTypesMedicationRequest                  ResourceTypes = "MedicationRequest"
	ResourceTypesMedicationStatement                ResourceTypes = "MedicationStatement"
	ResourceTypesMedicinalProductDefinition         ResourceTypes = "MedicinalProductDefinition"
	ResourceTypesMessageDefinition                  ResourceTypes = "MessageDefinition"
	ResourceTypesMessageHeader                      ResourceTypes = "MessageHeader"
	ResourceTypesMolecularSequence                  ResourceTypes = "MolecularSequence"
	ResourceTypesNamingSystem                       ResourceTypes = "NamingSystem"
	ResourceTypesNutritionIntake                    ResourceTypes = "NutritionIntake"
	ResourceTypesNutritionOrder                     ResourceTypes = "NutritionOrder"
	ResourceTypesNutritionProduct                   ResourceTypes = "NutritionProduct"
	ResourceTypesObservation                        ResourceTypes = "Observation"
	ResourceTypesObservationDefinition              ResourceTypes = "ObservationDefinition"
	ResourceTypesOperationDefinition                ResourceTypes = "OperationDefinition"
	ResourceTypesOperationOutcome                   ResourceTypes = "OperationOutcome"
	ResourceTypesOrganization                       ResourceTypes = "Organization"
	ResourceTypesOrganizationAffiliation            ResourceTypes = "OrganizationAffiliation"
	ResourceTypesPackagedProductDefinition          ResourceTypes = "PackagedProductDefinition"
	ResourceTypesParameters                         ResourceTypes = "Parameters"
	ResourceTypesPatient                            ResourceTypes = "Patient"
	ResourceTypesPaymentNotice                      ResourceTypes = "PaymentNotice"
	ResourceTypesPaymentReconciliation              ResourceTypes = "PaymentReconciliation"
	ResourceTypesPermission                         ResourceTypes = "Permission"
	ResourceTypesPerson                             ResourceTypes = "Person"
	ResourceTypesPlanDefinition                     ResourceTypes = "PlanDefinition"
	ResourceTypesPractitioner                       ResourceTypes = "Practitioner"
	ResourceTypesPractitionerRole                   ResourceTypes = "PractitionerRole"
	ResourceTypesProcedure                          ResourceTypes = "Procedure"
	ResourceTypesProvenance                         ResourceTypes = "Provenance"
	ResourceTypesQuestionnaire                      ResourceTypes = "Questionnaire"
	ResourceTypesQuestionnaireResponse              ResourceTypes = "QuestionnaireResponse"
	ResourceTypesRegulatedAuthorization             ResourceTypes = "RegulatedAuthorization"
	ResourceTypesRelatedPerson                      ResourceTypes = "RelatedPerson"
	ResourceTypesRequestOrchestration               ResourceTypes = "RequestOrchestration"
	ResourceTypesRequirements                       ResourceTypes = "Requirements"
	ResourceTypesResearchStudy                      ResourceTypes = "ResearchStudy"
	ResourceTypesResearchSubject                    ResourceTypes = "ResearchSubject"
	ResourceTypesRiskAssessment                     ResourceTypes = "RiskAssessment"
	ResourceTypesSchedule                           ResourceTypes = "Schedule"
	ResourceTypesSearchParameter                    ResourceTypes = "SearchParameter"
	ResourceTypesServiceRequest                     ResourceTypes = "ServiceRequest"
	ResourceTypesSlot                               ResourceTypes = "Slot"
	ResourceTypesSpecimen                           ResourceTypes = "Specimen"
	ResourceTypesSpecimenDefinition                 ResourceTypes = "SpecimenDefinition"
	ResourceTypesStructureDefinition                ResourceTypes = "StructureDefinition"
	ResourceTypesStructureMap                       ResourceTypes = "StructureMap"
	ResourceTypesSubscription                       ResourceTypes = "Subscription"
	ResourceTypesSubscriptionStatus                 ResourceTypes = "SubscriptionStatus"
	ResourceTypesSubscriptionTopic                  ResourceTypes = "SubscriptionTopic"
	ResourceTypesSubstance                          ResourceTypes = "Substance"
	ResourceTypesSubstanceDefinition                ResourceTypes = "SubstanceDefinition"
	ResourceTypesSubstanceNucleicAcid               ResourceTypes = "SubstanceNucleicAcid"
	ResourceTypesSubstancePolymer                   ResourceTypes = "SubstancePolymer"
	ResourceTypesSubstanceProtein                   ResourceTypes = "SubstanceProtein"
	ResourceTypesSubstanceReferenceInformation      ResourceTypes = "SubstanceReferenceInformation"
	ResourceTypesSubstanceSourceMaterial            ResourceTypes = "SubstanceSourceMaterial"
	ResourceTypesSupplyDelivery                     ResourceTypes = "SupplyDelivery"
	ResourceTypesSupplyRequest                      ResourceTypes = "SupplyRequest"
	ResourceTypesTask                               ResourceTypes = "Task"
	ResourceTypesTerminologyCapabilities            ResourceTypes = "TerminologyCapabilities"
	ResourceTypesTestPlan                           ResourceTypes = "TestPlan"
	ResourceTypesTestReport                         ResourceTypes = "TestReport"
	ResourceTypesTestScript                         ResourceTypes = "TestScript"
	ResourceTypesTransport                          ResourceTypes = "Transport"
	ResourceTypesValueSet                           ResourceTypes = "ValueSet"
	ResourceTypesVerificationResult                 ResourceTypes = "VerificationResult"
	ResourceTypesVisionPrescription                 ResourceTypes = "VisionPrescription"
)

func AllResourceTypes() []ResourceTypes {
	return []ResourceTypes{
		ResourceTypesAccount,
		ResourceTypesActivityDefinition,
		ResourceTypesActorDefinition,
		ResourceTypesAdministrableProductDefinition,
		ResourceTypesAdverseEvent,
		ResourceTypesAllergyIntolerance,
		ResourceTypesAppointment,
		ResourceTypesAppointmentResponse,
		ResourceTypesArtifactAssessment,
		ResourceTypesAuditEvent,
		ResourceTypesBasic,
		ResourceTypesBinary,
		ResourceTypesBiologicallyDerivedProduct,
		ResourceTypesBiologicallyDerivedProductDispense,
		ResourceTypesBodyStructure,
		ResourceTypesBundle,
		ResourceTypesCapabilityStatement,
		ResourceTypesCarePlan,
		ResourceTypesCareTeam,
		ResourceTypesChargeItem,
		ResourceTypesChargeItemDefinition,
		ResourceTypesCitation,
		ResourceTypesClaim,
		ResourceTypesClaimResponse,
		ResourceTypesClinicalImpression,
		ResourceTypesClinicalUseDefinition,
		ResourceTypesCodeSystem,
		ResourceTypesCommunication,
		ResourceTypesCommunicationRequest,
		ResourceTypesCompartmentDefinition,
		ResourceTypesComposition,
		ResourceTypesConceptMap,
		ResourceTypesCondition,
		ResourceTypesConditionDefinition,
		ResourceTypesConsent,
		ResourceTypesContract,
		ResourceTypesCoverage,
		ResourceTypesCoverageEligibilityRequest,
		ResourceTypesCoverageEligibilityResponse,
		ResourceTypesDetectedIssue,
		ResourceTypesDevice,
		ResourceTypesDeviceAssociation,
		ResourceTypesDeviceDefinition,
		ResourceTypesDeviceDispense,
		ResourceTypesDeviceMetric,
		ResourceTypesDeviceRequest,
		ResourceTypesDeviceUsage,
		ResourceTypesDiagnosticReport,
		ResourceTypesDocumentReference,
		ResourceTypesEncounter,
		ResourceTypesEncounterHistory,
		ResourceTypesEndpoint,
		ResourceTypesEnrollmentRequest,
		ResourceTypesEnrollmentResponse,
		ResourceTypesEpisodeOfCare,
		ResourceTypesEventDefinition,
		ResourceTypesEvidence,
		ResourceTypesEvidenceReport,
		ResourceTypesEvidenceVariable,
		ResourceTypesExampleScenario,
		ResourceTypesExplanationOfBenefit,
		ResourceTypesFamilyMemberHistory,
		ResourceTypesFlag,
		ResourceTypesFormularyItem,
		ResourceTypesGenomicStudy,
		ResourceTypesGoal,
		ResourceTypesGraphDefinition,
		ResourceTypesGroup,
		ResourceTypesGuidanceResponse,
		ResourceTypesHealthcareService,
		ResourceTypesImagingSelection,
		ResourceTypesImagingStudy,
		ResourceTypesImmunization,
		ResourceTypesImmunizationEvaluation,
		ResourceTypesImmunizationRecommendation,
		ResourceTypesImplementationGuide,
		ResourceTypesIngredient,
		ResourceTypesInsurancePlan,
		ResourceTypesInventoryItem,
		ResourceTypesInventoryReport,
		ResourceTypesInvoice,
		ResourceTypesLibrary,
		ResourceTypesLinkage,
		ResourceTypesList,
		ResourceTypesLocation,
		ResourceTypesManufacturedItemDefinition,
		ResourceTypesMeasure,
		ResourceTypesMeasureReport,
		ResourceTypesMedication,
		ResourceTypesMedicationAdministration,
		ResourceTypesMedicationDispense,
		ResourceTypesMedicationKnowledge,
		ResourceTypesMedicationRequest,
		ResourceTypesMedicationStatement,
		ResourceTypesMedicinalProductDefinition,
		ResourceTypesMessageDefinition,
		ResourceTypesMessageHeader,
		ResourceTypesMolecularSequence,
		ResourceTypesNamingSystem,
		ResourceTypesNutritionIntake,
		ResourceTypesNutritionOrder,
		ResourceTypesNutritionProduct,
		ResourceTypesObservation,
		ResourceTypesObservationDefinition,
		ResourceTypesOperationDefinition,
		ResourceTypesOperationOutcome,
		ResourceTypesOrganization,
		ResourceTypesOrganizationAffiliation,
		ResourceTypesPackagedProductDefinition,
		ResourceTypesParameters,
		ResourceTypesPatient,
		ResourceTypesPaymentNotice,
		ResourceTypesPaymentReconciliation,
		ResourceTypesPermission,
		ResourceTypesPerson,
		ResourceTypesPlanDefinition,
		ResourceTypesPractitioner,
		ResourceTypesPractitionerRole,
		ResourceTypesProcedure,
		ResourceTypesProvenance,
		ResourceTypesQuestionnaire,
		ResourceTypesQuestionnaireResponse,
		ResourceTypesRegulatedAuthorization,
		ResourceTypesRelatedPerson,
		ResourceTypesRequestOrchestration,
		ResourceTypesRequirements,
		ResourceTypesResearchStudy,
		ResourceTypesResearchSubject,
		ResourceTypesRiskAssessment,
		ResourceTypesSchedule,
		ResourceTypesSearchParameter,
		ResourceTypesServiceRequest,
		ResourceTypesSlot,
		ResourceTypesSpecimen,
		ResourceTypesSpecimenDefinition,
		ResourceTypesStructureDefinition,
		ResourceTypesStructureMap,
		ResourceTypesSubscription,
		ResourceTypesSubscriptionStatus,
		ResourceTypesSubscriptionTopic,
		ResourceTypesSubstance,
		ResourceTypesSubstanceDefinition,
		ResourceTypesSubstanceNucleicAcid,
		ResourceTypesSubstancePolymer,
		ResourceTypesSubstanceProtein,
		ResourceTypesSubstanceReferenceInformation,
		ResourceTypesSubstanceSourceMaterial,
		ResourceTypesSupplyDelivery,
		ResourceTypesSupplyRequest,
		ResourceTypesTask,
		ResourceTypesTerminologyCapabilities,
		ResourceTypesTestPlan,
		ResourceTypesTestReport,
		ResourceTypesTestScript,
		ResourceTypesTransport,
		ResourceTypesValueSet,
		ResourceTypesVerificationResult,
		ResourceTypesVisionPrescription,
	}
}

func FindResourceTypes(filter string) []ResourceTypes {
	ret := make([]ResourceTypes, 0)
	for _, at := range AllResourceTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ResourceTypes) ToString() {
	fmt.Println(au.String())
}
func (au ResourceTypes) String() string {
	switch au {
	case ResourceTypesAccount:
		return "Account"
	case ResourceTypesActivityDefinition:
		return "ActivityDefinition"
	case ResourceTypesActorDefinition:
		return "ActorDefinition"
	case ResourceTypesAdministrableProductDefinition:
		return "AdministrableProductDefinition"
	case ResourceTypesAdverseEvent:
		return "AdverseEvent"
	case ResourceTypesAllergyIntolerance:
		return "AllergyIntolerance"
	case ResourceTypesAppointment:
		return "Appointment"
	case ResourceTypesAppointmentResponse:
		return "AppointmentResponse"
	case ResourceTypesArtifactAssessment:
		return "ArtifactAssessment"
	case ResourceTypesAuditEvent:
		return "AuditEvent"
	case ResourceTypesBasic:
		return "Basic"
	case ResourceTypesBinary:
		return "Binary"
	case ResourceTypesBiologicallyDerivedProduct:
		return "BiologicallyDerivedProduct"
	case ResourceTypesBiologicallyDerivedProductDispense:
		return "BiologicallyDerivedProductDispense"
	case ResourceTypesBodyStructure:
		return "BodyStructure"
	case ResourceTypesBundle:
		return "Bundle"
	case ResourceTypesCapabilityStatement:
		return "CapabilityStatement"
	case ResourceTypesCarePlan:
		return "CarePlan"
	case ResourceTypesCareTeam:
		return "CareTeam"
	case ResourceTypesChargeItem:
		return "ChargeItem"
	case ResourceTypesChargeItemDefinition:
		return "ChargeItemDefinition"
	case ResourceTypesCitation:
		return "Citation"
	case ResourceTypesClaim:
		return "Claim"
	case ResourceTypesClaimResponse:
		return "ClaimResponse"
	case ResourceTypesClinicalImpression:
		return "ClinicalImpression"
	case ResourceTypesClinicalUseDefinition:
		return "ClinicalUseDefinition"
	case ResourceTypesCodeSystem:
		return "CodeSystem"
	case ResourceTypesCommunication:
		return "Communication"
	case ResourceTypesCommunicationRequest:
		return "CommunicationRequest"
	case ResourceTypesCompartmentDefinition:
		return "CompartmentDefinition"
	case ResourceTypesComposition:
		return "Composition"
	case ResourceTypesConceptMap:
		return "ConceptMap"
	case ResourceTypesCondition:
		return "Condition"
	case ResourceTypesConditionDefinition:
		return "ConditionDefinition"
	case ResourceTypesConsent:
		return "Consent"
	case ResourceTypesContract:
		return "Contract"
	case ResourceTypesCoverage:
		return "Coverage"
	case ResourceTypesCoverageEligibilityRequest:
		return "CoverageEligibilityRequest"
	case ResourceTypesCoverageEligibilityResponse:
		return "CoverageEligibilityResponse"
	case ResourceTypesDetectedIssue:
		return "DetectedIssue"
	case ResourceTypesDevice:
		return "Device"
	case ResourceTypesDeviceAssociation:
		return "DeviceAssociation"
	case ResourceTypesDeviceDefinition:
		return "DeviceDefinition"
	case ResourceTypesDeviceDispense:
		return "DeviceDispense"
	case ResourceTypesDeviceMetric:
		return "DeviceMetric"
	case ResourceTypesDeviceRequest:
		return "DeviceRequest"
	case ResourceTypesDeviceUsage:
		return "DeviceUsage"
	case ResourceTypesDiagnosticReport:
		return "DiagnosticReport"
	case ResourceTypesDocumentReference:
		return "DocumentReference"
	case ResourceTypesEncounter:
		return "Encounter"
	case ResourceTypesEncounterHistory:
		return "EncounterHistory"
	case ResourceTypesEndpoint:
		return "Endpoint"
	case ResourceTypesEnrollmentRequest:
		return "EnrollmentRequest"
	case ResourceTypesEnrollmentResponse:
		return "EnrollmentResponse"
	case ResourceTypesEpisodeOfCare:
		return "EpisodeOfCare"
	case ResourceTypesEventDefinition:
		return "EventDefinition"
	case ResourceTypesEvidence:
		return "Evidence"
	case ResourceTypesEvidenceReport:
		return "EvidenceReport"
	case ResourceTypesEvidenceVariable:
		return "EvidenceVariable"
	case ResourceTypesExampleScenario:
		return "ExampleScenario"
	case ResourceTypesExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case ResourceTypesFamilyMemberHistory:
		return "FamilyMemberHistory"
	case ResourceTypesFlag:
		return "Flag"
	case ResourceTypesFormularyItem:
		return "FormularyItem"
	case ResourceTypesGenomicStudy:
		return "GenomicStudy"
	case ResourceTypesGoal:
		return "Goal"
	case ResourceTypesGraphDefinition:
		return "GraphDefinition"
	case ResourceTypesGroup:
		return "Group"
	case ResourceTypesGuidanceResponse:
		return "GuidanceResponse"
	case ResourceTypesHealthcareService:
		return "HealthcareService"
	case ResourceTypesImagingSelection:
		return "ImagingSelection"
	case ResourceTypesImagingStudy:
		return "ImagingStudy"
	case ResourceTypesImmunization:
		return "Immunization"
	case ResourceTypesImmunizationEvaluation:
		return "ImmunizationEvaluation"
	case ResourceTypesImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case ResourceTypesImplementationGuide:
		return "ImplementationGuide"
	case ResourceTypesIngredient:
		return "Ingredient"
	case ResourceTypesInsurancePlan:
		return "InsurancePlan"
	case ResourceTypesInventoryItem:
		return "InventoryItem"
	case ResourceTypesInventoryReport:
		return "InventoryReport"
	case ResourceTypesInvoice:
		return "Invoice"
	case ResourceTypesLibrary:
		return "Library"
	case ResourceTypesLinkage:
		return "Linkage"
	case ResourceTypesList:
		return "List"
	case ResourceTypesLocation:
		return "Location"
	case ResourceTypesManufacturedItemDefinition:
		return "ManufacturedItemDefinition"
	case ResourceTypesMeasure:
		return "Measure"
	case ResourceTypesMeasureReport:
		return "MeasureReport"
	case ResourceTypesMedication:
		return "Medication"
	case ResourceTypesMedicationAdministration:
		return "MedicationAdministration"
	case ResourceTypesMedicationDispense:
		return "MedicationDispense"
	case ResourceTypesMedicationKnowledge:
		return "MedicationKnowledge"
	case ResourceTypesMedicationRequest:
		return "MedicationRequest"
	case ResourceTypesMedicationStatement:
		return "MedicationStatement"
	case ResourceTypesMedicinalProductDefinition:
		return "MedicinalProductDefinition"
	case ResourceTypesMessageDefinition:
		return "MessageDefinition"
	case ResourceTypesMessageHeader:
		return "MessageHeader"
	case ResourceTypesMolecularSequence:
		return "MolecularSequence"
	case ResourceTypesNamingSystem:
		return "NamingSystem"
	case ResourceTypesNutritionIntake:
		return "NutritionIntake"
	case ResourceTypesNutritionOrder:
		return "NutritionOrder"
	case ResourceTypesNutritionProduct:
		return "NutritionProduct"
	case ResourceTypesObservation:
		return "Observation"
	case ResourceTypesObservationDefinition:
		return "ObservationDefinition"
	case ResourceTypesOperationDefinition:
		return "OperationDefinition"
	case ResourceTypesOperationOutcome:
		return "OperationOutcome"
	case ResourceTypesOrganization:
		return "Organization"
	case ResourceTypesOrganizationAffiliation:
		return "OrganizationAffiliation"
	case ResourceTypesPackagedProductDefinition:
		return "PackagedProductDefinition"
	case ResourceTypesParameters:
		return "Parameters"
	case ResourceTypesPatient:
		return "Patient"
	case ResourceTypesPaymentNotice:
		return "PaymentNotice"
	case ResourceTypesPaymentReconciliation:
		return "PaymentReconciliation"
	case ResourceTypesPermission:
		return "Permission"
	case ResourceTypesPerson:
		return "Person"
	case ResourceTypesPlanDefinition:
		return "PlanDefinition"
	case ResourceTypesPractitioner:
		return "Practitioner"
	case ResourceTypesPractitionerRole:
		return "PractitionerRole"
	case ResourceTypesProcedure:
		return "Procedure"
	case ResourceTypesProvenance:
		return "Provenance"
	case ResourceTypesQuestionnaire:
		return "Questionnaire"
	case ResourceTypesQuestionnaireResponse:
		return "QuestionnaireResponse"
	case ResourceTypesRegulatedAuthorization:
		return "RegulatedAuthorization"
	case ResourceTypesRelatedPerson:
		return "RelatedPerson"
	case ResourceTypesRequestOrchestration:
		return "RequestOrchestration"
	case ResourceTypesRequirements:
		return "Requirements"
	case ResourceTypesResearchStudy:
		return "ResearchStudy"
	case ResourceTypesResearchSubject:
		return "ResearchSubject"
	case ResourceTypesRiskAssessment:
		return "RiskAssessment"
	case ResourceTypesSchedule:
		return "Schedule"
	case ResourceTypesSearchParameter:
		return "SearchParameter"
	case ResourceTypesServiceRequest:
		return "ServiceRequest"
	case ResourceTypesSlot:
		return "Slot"
	case ResourceTypesSpecimen:
		return "Specimen"
	case ResourceTypesSpecimenDefinition:
		return "SpecimenDefinition"
	case ResourceTypesStructureDefinition:
		return "StructureDefinition"
	case ResourceTypesStructureMap:
		return "StructureMap"
	case ResourceTypesSubscription:
		return "Subscription"
	case ResourceTypesSubscriptionStatus:
		return "SubscriptionStatus"
	case ResourceTypesSubscriptionTopic:
		return "SubscriptionTopic"
	case ResourceTypesSubstance:
		return "Substance"
	case ResourceTypesSubstanceDefinition:
		return "SubstanceDefinition"
	case ResourceTypesSubstanceNucleicAcid:
		return "SubstanceNucleicAcid"
	case ResourceTypesSubstancePolymer:
		return "SubstancePolymer"
	case ResourceTypesSubstanceProtein:
		return "SubstanceProtein"
	case ResourceTypesSubstanceReferenceInformation:
		return "SubstanceReferenceInformation"
	case ResourceTypesSubstanceSourceMaterial:
		return "SubstanceSourceMaterial"
	case ResourceTypesSupplyDelivery:
		return "SupplyDelivery"
	case ResourceTypesSupplyRequest:
		return "SupplyRequest"
	case ResourceTypesTask:
		return "Task"
	case ResourceTypesTerminologyCapabilities:
		return "TerminologyCapabilities"
	case ResourceTypesTestPlan:
		return "TestPlan"
	case ResourceTypesTestReport:
		return "TestReport"
	case ResourceTypesTestScript:
		return "TestScript"
	case ResourceTypesTransport:
		return "Transport"
	case ResourceTypesValueSet:
		return "ValueSet"
	case ResourceTypesVerificationResult:
		return "VerificationResult"
	case ResourceTypesVisionPrescription:
		return "VisionPrescription"

	default:
		return "Unknown Constraint Severity"
	}
}
