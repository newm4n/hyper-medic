package code

import (
	"fmt"
	"strings"
)

type AllResourceType int

const (
	ResourceTypeAccount AllResourceType = iota
	ResourceTypeActivityDefinition
	ResourceTypeActorDefinition
	ResourceTypeAdministrableProductDefinition
	ResourceTypeAdverseEvent
	ResourceTypeAllergyIntolerance
	ResourceTypeAppointment
	ResourceTypeAppointmentResponse
	ResourceTypeArtifactAssessment
	ResourceTypeAuditEvent
	ResourceTypeBasic
	ResourceTypeBinary
	ResourceTypeBiologicallyDerivedProduct
	ResourceTypeBiologicallyDerivedProductDispense
	ResourceTypeBodyStructure
	ResourceTypeBundle
	ResourceTypeCanonicalResource
	ResourceTypeCapabilityStatement
	ResourceTypeCarePlan
	ResourceTypeCareTeam
	ResourceTypeChargeItem
	ResourceTypeChargeItemDefinition
	ResourceTypeCitation
	ResourceTypeClaim
	ResourceTypeClaimResponse
	ResourceTypeClinicalImpression
	ResourceTypeClinicalUseDefinition
	ResourceTypeCodeSystem
	ResourceTypeCommunication
	ResourceTypeCommunicationRequest
	ResourceTypeCompartmentDefinition
	ResourceTypeComposition
	ResourceTypeConceptMap
	ResourceTypeCondition
	ResourceTypeConditionDefinition
	ResourceTypeConsent
	ResourceTypeContract
	ResourceTypeCoverage
	ResourceTypeCoverageEligibilityRequest
	ResourceTypeCoverageEligibilityResponse
	ResourceTypeDetectedIssue
	ResourceTypeDevice
	ResourceTypeDeviceAssociation
	ResourceTypeDeviceDefinition
	ResourceTypeDeviceDispense
	ResourceTypeDeviceMetric
	ResourceTypeDeviceRequest
	ResourceTypeDeviceUsage
	ResourceTypeDiagnosticReport
	ResourceTypeDocumentReference
	ResourceTypeDomainResource
	ResourceTypeEncounter
	ResourceTypeEncounterHistory
	ResourceTypeEndpoint
	ResourceTypeEnrollmentRequest
	ResourceTypeEnrollmentResponse
	ResourceTypeEpisodeOfCare
	ResourceTypeEventDefinition
	ResourceTypeEvidence
	ResourceTypeEvidenceReport
	ResourceTypeEvidenceVariable
	ResourceTypeExampleScenario
	ResourceTypeExplanationOfBenefit
	ResourceTypeFamilyMemberHistory
	ResourceTypeFlag
	ResourceTypeFormularyItem
	ResourceTypeGenomicStudy
	ResourceTypeGoal
	ResourceTypeGraphDefinition
	ResourceTypeGroup
	ResourceTypeGuidanceResponse
	ResourceTypeHealthcareService
	ResourceTypeImagingSelection
	ResourceTypeImagingStudy
	ResourceTypeImmunization
	ResourceTypeImmunizationEvaluation
	ResourceTypeImmunizationRecommendation
	ResourceTypeImplementationGuide
	ResourceTypeIngredient
	ResourceTypeInsurancePlan
	ResourceTypeInventoryItem
	ResourceTypeInventoryReport
	ResourceTypeInvoice
	ResourceTypeLibrary
	ResourceTypeLinkage
	ResourceTypeList
	ResourceTypeLocation
	ResourceTypeManufacturedItemDefinition
	ResourceTypeMeasure
	ResourceTypeMeasureReport
	ResourceTypeMedication
	ResourceTypeMedicationAdministration
	ResourceTypeMedicationDispense
	ResourceTypeMedicationKnowledge
	ResourceTypeMedicationRequest
	ResourceTypeMedicationStatement
	ResourceTypeMedicinalProductDefinition
	ResourceTypeMessageDefinition
	ResourceTypeMessageHeader
	ResourceTypeMetadataResource
	ResourceTypeMolecularSequence
	ResourceTypeNamingSystem
	ResourceTypeNutritionIntake
	ResourceTypeNutritionOrder
	ResourceTypeNutritionProduct
	ResourceTypeObservation
	ResourceTypeObservationDefinition
	ResourceTypeOperationDefinition
	ResourceTypeOperationOutcome
	ResourceTypeOrganization
	ResourceTypeOrganizationAffiliation
	ResourceTypePackagedProductDefinition
	ResourceTypeParameters
	ResourceTypePatient
	ResourceTypePaymentNotice
	ResourceTypePaymentReconciliation
	ResourceTypePermission
	ResourceTypePerson
	ResourceTypePlanDefinition
	ResourceTypePractitioner
	ResourceTypePractitionerRole
	ResourceTypeProcedure
	ResourceTypeProvenance
	ResourceTypeQuestionnaire
	ResourceTypeQuestionnaireResponse
	ResourceTypeRegulatedAuthorization
	ResourceTypeRelatedPerson
	ResourceTypeRequestOrchestration
	ResourceTypeRequirements
	ResourceTypeResearchStudy
	ResourceTypeResearchSubject
	ResourceTypeResource
	ResourceTypeRiskAssessment
	ResourceTypeSchedule
	ResourceTypeSearchParameter
	ResourceTypeServiceRequest
	ResourceTypeSlot
	ResourceTypeSpecimen
	ResourceTypeSpecimenDefinition
	ResourceTypeStructureDefinition
	ResourceTypeStructureMap
	ResourceTypeSubscription
	ResourceTypeSubscriptionStatus
	ResourceTypeSubscriptionTopic
	ResourceTypeSubstance
	ResourceTypeSubstanceDefinition
	ResourceTypeSubstanceNucleicAcid
	ResourceTypeSubstancePolymer
	ResourceTypeSubstanceProtein
	ResourceTypeSubstanceReferenceInformation
	ResourceTypeSubstanceSourceMaterial
	ResourceTypeSupplyDelivery
	ResourceTypeSupplyRequest
	ResourceTypeTask
	ResourceTypeTerminologyCapabilities
	ResourceTypeTestPlan
	ResourceTypeTestReport
	ResourceTypeTestScript
	ResourceTypeTransport
	ResourceTypeValueSet
	ResourceTypeVerificationResult
	ResourceTypeVisionPrescription
)

func AllAllResourceType() []AllResourceType {
	return []AllResourceType{
		ResourceTypeAccount,
		ResourceTypeActivityDefinition,
		ResourceTypeActorDefinition,
		ResourceTypeAdministrableProductDefinition,
		ResourceTypeAdverseEvent,
		ResourceTypeAllergyIntolerance,
		ResourceTypeAppointment,
		ResourceTypeAppointmentResponse,
		ResourceTypeArtifactAssessment,
		ResourceTypeAuditEvent,
		ResourceTypeBasic,
		ResourceTypeBinary,
		ResourceTypeBiologicallyDerivedProduct,
		ResourceTypeBiologicallyDerivedProductDispense,
		ResourceTypeBodyStructure,
		ResourceTypeBundle,
		ResourceTypeCanonicalResource,
		ResourceTypeCapabilityStatement,
		ResourceTypeCarePlan,
		ResourceTypeCareTeam,
		ResourceTypeChargeItem,
		ResourceTypeChargeItemDefinition,
		ResourceTypeCitation,
		ResourceTypeClaim,
		ResourceTypeClaimResponse,
		ResourceTypeClinicalImpression,
		ResourceTypeClinicalUseDefinition,
		ResourceTypeCodeSystem,
		ResourceTypeCommunication,
		ResourceTypeCommunicationRequest,
		ResourceTypeCompartmentDefinition,
		ResourceTypeComposition,
		ResourceTypeConceptMap,
		ResourceTypeCondition,
		ResourceTypeConditionDefinition,
		ResourceTypeConsent,
		ResourceTypeContract,
		ResourceTypeCoverage,
		ResourceTypeCoverageEligibilityRequest,
		ResourceTypeCoverageEligibilityResponse,
		ResourceTypeDetectedIssue,
		ResourceTypeDevice,
		ResourceTypeDeviceAssociation,
		ResourceTypeDeviceDefinition,
		ResourceTypeDeviceDispense,
		ResourceTypeDeviceMetric,
		ResourceTypeDeviceRequest,
		ResourceTypeDeviceUsage,
		ResourceTypeDiagnosticReport,
		ResourceTypeDocumentReference,
		ResourceTypeDomainResource,
		ResourceTypeEncounter,
		ResourceTypeEncounterHistory,
		ResourceTypeEndpoint,
		ResourceTypeEnrollmentRequest,
		ResourceTypeEnrollmentResponse,
		ResourceTypeEpisodeOfCare,
		ResourceTypeEventDefinition,
		ResourceTypeEvidence,
		ResourceTypeEvidenceReport,
		ResourceTypeEvidenceVariable,
		ResourceTypeExampleScenario,
		ResourceTypeExplanationOfBenefit,
		ResourceTypeFamilyMemberHistory,
		ResourceTypeFlag,
		ResourceTypeFormularyItem,
		ResourceTypeGenomicStudy,
		ResourceTypeGoal,
		ResourceTypeGraphDefinition,
		ResourceTypeGroup,
		ResourceTypeGuidanceResponse,
		ResourceTypeHealthcareService,
		ResourceTypeImagingSelection,
		ResourceTypeImagingStudy,
		ResourceTypeImmunization,
		ResourceTypeImmunizationEvaluation,
		ResourceTypeImmunizationRecommendation,
		ResourceTypeImplementationGuide,
		ResourceTypeIngredient,
		ResourceTypeInsurancePlan,
		ResourceTypeInventoryItem,
		ResourceTypeInventoryReport,
		ResourceTypeInvoice,
		ResourceTypeLibrary,
		ResourceTypeLinkage,
		ResourceTypeList,
		ResourceTypeLocation,
		ResourceTypeManufacturedItemDefinition,
		ResourceTypeMeasure,
		ResourceTypeMeasureReport,
		ResourceTypeMedication,
		ResourceTypeMedicationAdministration,
		ResourceTypeMedicationDispense,
		ResourceTypeMedicationKnowledge,
		ResourceTypeMedicationRequest,
		ResourceTypeMedicationStatement,
		ResourceTypeMedicinalProductDefinition,
		ResourceTypeMessageDefinition,
		ResourceTypeMessageHeader,
		ResourceTypeMetadataResource,
		ResourceTypeMolecularSequence,
		ResourceTypeNamingSystem,
		ResourceTypeNutritionIntake,
		ResourceTypeNutritionOrder,
		ResourceTypeNutritionProduct,
		ResourceTypeObservation,
		ResourceTypeObservationDefinition,
		ResourceTypeOperationDefinition,
		ResourceTypeOperationOutcome,
		ResourceTypeOrganization,
		ResourceTypeOrganizationAffiliation,
		ResourceTypePackagedProductDefinition,
		ResourceTypeParameters,
		ResourceTypePatient,
		ResourceTypePaymentNotice,
		ResourceTypePaymentReconciliation,
		ResourceTypePermission,
		ResourceTypePerson,
		ResourceTypePlanDefinition,
		ResourceTypePractitioner,
		ResourceTypePractitionerRole,
		ResourceTypeProcedure,
		ResourceTypeProvenance,
		ResourceTypeQuestionnaire,
		ResourceTypeQuestionnaireResponse,
		ResourceTypeRegulatedAuthorization,
		ResourceTypeRelatedPerson,
		ResourceTypeRequestOrchestration,
		ResourceTypeRequirements,
		ResourceTypeResearchStudy,
		ResourceTypeResearchSubject,
		ResourceTypeResource,
		ResourceTypeRiskAssessment,
		ResourceTypeSchedule,
		ResourceTypeSearchParameter,
		ResourceTypeServiceRequest,
		ResourceTypeSlot,
		ResourceTypeSpecimen,
		ResourceTypeSpecimenDefinition,
		ResourceTypeStructureDefinition,
		ResourceTypeStructureMap,
		ResourceTypeSubscription,
		ResourceTypeSubscriptionStatus,
		ResourceTypeSubscriptionTopic,
		ResourceTypeSubstance,
		ResourceTypeSubstanceDefinition,
		ResourceTypeSubstanceNucleicAcid,
		ResourceTypeSubstancePolymer,
		ResourceTypeSubstanceProtein,
		ResourceTypeSubstanceReferenceInformation,
		ResourceTypeSubstanceSourceMaterial,
		ResourceTypeSupplyDelivery,
		ResourceTypeSupplyRequest,
		ResourceTypeTask,
		ResourceTypeTerminologyCapabilities,
		ResourceTypeTestPlan,
		ResourceTypeTestReport,
		ResourceTypeTestScript,
		ResourceTypeTransport,
		ResourceTypeValueSet,
		ResourceTypeVerificationResult,
		ResourceTypeVisionPrescription,
	}
}

func FindAllResourceType(filter string) []AllResourceType {
	ret := make([]AllResourceType, 0)
	for _, at := range AllAllResourceType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllResourceType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllResourceType) String() string {
	switch cpt {
	case ResourceTypeAccount:
		return "Account"
	case ResourceTypeActivityDefinition:
		return "ActivityDefinition"
	case ResourceTypeActorDefinition:
		return "ActorDefinition"
	case ResourceTypeAdministrableProductDefinition:
		return "AdministrableProductDefinition"
	case ResourceTypeAdverseEvent:
		return "AdverseEvent"
	case ResourceTypeAllergyIntolerance:
		return "AllergyIntolerance"
	case ResourceTypeAppointment:
		return "Appointment"
	case ResourceTypeAppointmentResponse:
		return "AppointmentResponse"
	case ResourceTypeArtifactAssessment:
		return "ArtifactAssessment"
	case ResourceTypeAuditEvent:
		return "AuditEvent"
	case ResourceTypeBasic:
		return "Basic"
	case ResourceTypeBinary:
		return "Binary"
	case ResourceTypeBiologicallyDerivedProduct:
		return "BiologicallyDerivedProduct"
	case ResourceTypeBiologicallyDerivedProductDispense:
		return "BiologicallyDerivedProductDispense"
	case ResourceTypeBodyStructure:
		return "BodyStructure"
	case ResourceTypeBundle:
		return "Bundle"
	case ResourceTypeCanonicalResource:
		return "CanonicalResource"
	case ResourceTypeCapabilityStatement:
		return "CapabilityStatement"
	case ResourceTypeCarePlan:
		return "CarePlan"
	case ResourceTypeCareTeam:
		return "CareTeam"
	case ResourceTypeChargeItem:
		return "ChargeItem"
	case ResourceTypeChargeItemDefinition:
		return "ChargeItemDefinition"
	case ResourceTypeCitation:
		return "Citation"
	case ResourceTypeClaim:
		return "Claim"
	case ResourceTypeClaimResponse:
		return "ClaimResponse"
	case ResourceTypeClinicalImpression:
		return "ClinicalImpression"
	case ResourceTypeClinicalUseDefinition:
		return "ClinicalUseDefinition"
	case ResourceTypeCodeSystem:
		return "CodeSystem"
	case ResourceTypeCommunication:
		return "Communication"
	case ResourceTypeCommunicationRequest:
		return "CommunicationRequest"
	case ResourceTypeCompartmentDefinition:
		return "CompartmentDefinition"
	case ResourceTypeComposition:
		return "Composition"
	case ResourceTypeConceptMap:
		return "ConceptMap"
	case ResourceTypeCondition:
		return "Condition"
	case ResourceTypeConditionDefinition:
		return "ConditionDefinition"
	case ResourceTypeConsent:
		return "Consent"
	case ResourceTypeContract:
		return "Contract"
	case ResourceTypeCoverage:
		return "Coverage"
	case ResourceTypeCoverageEligibilityRequest:
		return "CoverageEligibilityRequest"
	case ResourceTypeCoverageEligibilityResponse:
		return "CoverageEligibilityResponse"
	case ResourceTypeDetectedIssue:
		return "DetectedIssue"
	case ResourceTypeDevice:
		return "Device"
	case ResourceTypeDeviceAssociation:
		return "DeviceAssociation"
	case ResourceTypeDeviceDefinition:
		return "DeviceDefinition"
	case ResourceTypeDeviceDispense:
		return "DeviceDispense"
	case ResourceTypeDeviceMetric:
		return "DeviceMetric"
	case ResourceTypeDeviceRequest:
		return "DeviceRequest"
	case ResourceTypeDeviceUsage:
		return "DeviceUsage"
	case ResourceTypeDiagnosticReport:
		return "DiagnosticReport"
	case ResourceTypeDocumentReference:
		return "DocumentReference"
	case ResourceTypeDomainResource:
		return "DomainResource"
	case ResourceTypeEncounter:
		return "Encounter"
	case ResourceTypeEncounterHistory:
		return "EncounterHistory"
	case ResourceTypeEndpoint:
		return "Endpoint"
	case ResourceTypeEnrollmentRequest:
		return "EnrollmentRequest"
	case ResourceTypeEnrollmentResponse:
		return "EnrollmentResponse"
	case ResourceTypeEpisodeOfCare:
		return "EpisodeOfCare"
	case ResourceTypeEventDefinition:
		return "EventDefinition"
	case ResourceTypeEvidence:
		return "Evidence"
	case ResourceTypeEvidenceReport:
		return "EvidenceReport"
	case ResourceTypeEvidenceVariable:
		return "EvidenceVariable"
	case ResourceTypeExampleScenario:
		return "ExampleScenario"
	case ResourceTypeExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case ResourceTypeFamilyMemberHistory:
		return "FamilyMemberHistory"
	case ResourceTypeFlag:
		return "Flag"
	case ResourceTypeFormularyItem:
		return "FormularyItem"
	case ResourceTypeGenomicStudy:
		return "GenomicStudy"
	case ResourceTypeGoal:
		return "Goal"
	case ResourceTypeGraphDefinition:
		return "GraphDefinition"
	case ResourceTypeGroup:
		return "Group"
	case ResourceTypeGuidanceResponse:
		return "GuidanceResponse"
	case ResourceTypeHealthcareService:
		return "HealthcareService"
	case ResourceTypeImagingSelection:
		return "ImagingSelection"
	case ResourceTypeImagingStudy:
		return "ImagingStudy"
	case ResourceTypeImmunization:
		return "Immunization"
	case ResourceTypeImmunizationEvaluation:
		return "ImmunizationEvaluation"
	case ResourceTypeImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case ResourceTypeImplementationGuide:
		return "ImplementationGuide"
	case ResourceTypeIngredient:
		return "Ingredient"
	case ResourceTypeInsurancePlan:
		return "InsurancePlan"
	case ResourceTypeInventoryItem:
		return "InventoryItem"
	case ResourceTypeInventoryReport:
		return "InventoryReport"
	case ResourceTypeInvoice:
		return "Invoice"
	case ResourceTypeLibrary:
		return "Library"
	case ResourceTypeLinkage:
		return "Linkage"
	case ResourceTypeList:
		return "List"
	case ResourceTypeLocation:
		return "Location"
	case ResourceTypeManufacturedItemDefinition:
		return "ManufacturedItemDefinition"
	case ResourceTypeMeasure:
		return "Measure"
	case ResourceTypeMeasureReport:
		return "MeasureReport"
	case ResourceTypeMedication:
		return "Medication"
	case ResourceTypeMedicationAdministration:
		return "MedicationAdministration"
	case ResourceTypeMedicationDispense:
		return "MedicationDispense"
	case ResourceTypeMedicationKnowledge:
		return "MedicationKnowledge"
	case ResourceTypeMedicationRequest:
		return "MedicationRequest"
	case ResourceTypeMedicationStatement:
		return "MedicationStatement"
	case ResourceTypeMedicinalProductDefinition:
		return "MedicinalProductDefinition"
	case ResourceTypeMessageDefinition:
		return "MessageDefinition"
	case ResourceTypeMessageHeader:
		return "MessageHeader"
	case ResourceTypeMetadataResource:
		return "MetadataResource"
	case ResourceTypeMolecularSequence:
		return "MolecularSequence"
	case ResourceTypeNamingSystem:
		return "NamingSystem"
	case ResourceTypeNutritionIntake:
		return "NutritionIntake"
	case ResourceTypeNutritionOrder:
		return "NutritionOrder"
	case ResourceTypeNutritionProduct:
		return "NutritionProduct"
	case ResourceTypeObservation:
		return "Observation"
	case ResourceTypeObservationDefinition:
		return "ObservationDefinition"
	case ResourceTypeOperationDefinition:
		return "OperationDefinition"
	case ResourceTypeOperationOutcome:
		return "OperationOutcome"
	case ResourceTypeOrganization:
		return "Organization"
	case ResourceTypeOrganizationAffiliation:
		return "OrganizationAffiliation"
	case ResourceTypePackagedProductDefinition:
		return "PackagedProductDefinition"
	case ResourceTypeParameters:
		return "Parameters"
	case ResourceTypePatient:
		return "Patient"
	case ResourceTypePaymentNotice:
		return "PaymentNotice"
	case ResourceTypePaymentReconciliation:
		return "PaymentReconciliation"
	case ResourceTypePermission:
		return "Permission"
	case ResourceTypePerson:
		return "Person"
	case ResourceTypePlanDefinition:
		return "PlanDefinition"
	case ResourceTypePractitioner:
		return "Practitioner"
	case ResourceTypePractitionerRole:
		return "PractitionerRole"
	case ResourceTypeProcedure:
		return "Procedure"
	case ResourceTypeProvenance:
		return "Provenance"
	case ResourceTypeQuestionnaire:
		return "Questionnaire"
	case ResourceTypeQuestionnaireResponse:
		return "QuestionnaireResponse"
	case ResourceTypeRegulatedAuthorization:
		return "RegulatedAuthorization"
	case ResourceTypeRelatedPerson:
		return "RelatedPerson"
	case ResourceTypeRequestOrchestration:
		return "RequestOrchestration"
	case ResourceTypeRequirements:
		return "Requirements"
	case ResourceTypeResearchStudy:
		return "ResearchStudy"
	case ResourceTypeResearchSubject:
		return "ResearchSubject"
	case ResourceTypeResource:
		return "Resource"
	case ResourceTypeRiskAssessment:
		return "RiskAssessment"
	case ResourceTypeSchedule:
		return "Schedule"
	case ResourceTypeSearchParameter:
		return "SearchParameter"
	case ResourceTypeServiceRequest:
		return "ServiceRequest"
	case ResourceTypeSlot:
		return "Slot"
	case ResourceTypeSpecimen:
		return "Specimen"
	case ResourceTypeSpecimenDefinition:
		return "SpecimenDefinition"
	case ResourceTypeStructureDefinition:
		return "StructureDefinition"
	case ResourceTypeStructureMap:
		return "StructureMap"
	case ResourceTypeSubscription:
		return "Subscription"
	case ResourceTypeSubscriptionStatus:
		return "SubscriptionStatus"
	case ResourceTypeSubscriptionTopic:
		return "SubscriptionTopic"
	case ResourceTypeSubstance:
		return "Substance"
	case ResourceTypeSubstanceDefinition:
		return "SubstanceDefinition"
	case ResourceTypeSubstanceNucleicAcid:
		return "SubstanceNucleicAcid"
	case ResourceTypeSubstancePolymer:
		return "SubstancePolymer"
	case ResourceTypeSubstanceProtein:
		return "SubstanceProtein"
	case ResourceTypeSubstanceReferenceInformation:
		return "SubstanceReferenceInformation"
	case ResourceTypeSubstanceSourceMaterial:
		return "SubstanceSourceMaterial"
	case ResourceTypeSupplyDelivery:
		return "SupplyDelivery"
	case ResourceTypeSupplyRequest:
		return "SupplyRequest"
	case ResourceTypeTask:
		return "Task"
	case ResourceTypeTerminologyCapabilities:
		return "TerminologyCapabilities"
	case ResourceTypeTestPlan:
		return "TestPlan"
	case ResourceTypeTestReport:
		return "TestReport"
	case ResourceTypeTestScript:
		return "TestScript"
	case ResourceTypeTransport:
		return "Transport"
	case ResourceTypeValueSet:
		return "ValueSet"
	case ResourceTypeVerificationResult:
		return "VerificationResult"
	case ResourceTypeVisionPrescription:
		return "VisionPrescription"
	default:
		return "Unknown Appointment Recurrence Type"
	}
}

/*
Account
ActivityDefinition
ActorDefinition
AdministrableProductDefinition
AdverseEvent
AllergyIntolerance
Appointment
AppointmentResponse
ArtifactAssessment
AuditEvent
Basic
Binary
BiologicallyDerivedProduct
BiologicallyDerivedProductDispense
BodyStructure
Bundle
CanonicalResource
CapabilityStatement
CarePlan
CareTeam
ChargeItem
ChargeItemDefinition
Citation
Claim
ClaimResponse
ClinicalImpression
ClinicalUseDefinition
CodeSystem
Communication
CommunicationRequest
CompartmentDefinition
Composition
ConceptMap
Condition
ConditionDefinition
Consent
Contract
Coverage
CoverageEligibilityRequest
CoverageEligibilityResponse
DetectedIssue
Device
DeviceAssociation
DeviceDefinition
DeviceDispense
DeviceMetric
DeviceRequest
DeviceUsage
DiagnosticReport
DocumentReference
DomainResource
Encounter
EncounterHistory
Endpoint
EnrollmentRequest
EnrollmentResponse
EpisodeOfCare
EventDefinition
Evidence
EvidenceReport
EvidenceVariable
ExampleScenario
ExplanationOfBenefit
FamilyMemberHistory
Flag
FormularyItem
GenomicStudy
Goal
GraphDefinition
Group
GuidanceResponse
HealthcareService
ImagingSelection
ImagingStudy
Immunization
ImmunizationEvaluation
ImmunizationRecommendation
ImplementationGuide
Ingredient
InsurancePlan
InventoryItem
InventoryReport
Invoice
Library
Linkage
List
Location
ManufacturedItemDefinition
Measure
MeasureReport
Medication
MedicationAdministration
MedicationDispense
MedicationKnowledge
MedicationRequest
MedicationStatement
MedicinalProductDefinition
MessageDefinition
MessageHeader
MetadataResource
MolecularSequence
NamingSystem
NutritionIntake
NutritionOrder
NutritionProduct
Observation
ObservationDefinition
OperationDefinition
OperationOutcome
Organization
OrganizationAffiliation
PackagedProductDefinition
Parameters
Patient
PaymentNotice
PaymentReconciliation
Permission
Person
PlanDefinition
Practitioner
PractitionerRole
Procedure
Provenance
Questionnaire
QuestionnaireResponse
RegulatedAuthorization
RelatedPerson
RequestOrchestration
Requirements
ResearchStudy
ResearchSubject
Resource
RiskAssessment
Schedule
SearchParameter
ServiceRequest
Slot
Specimen
SpecimenDefinition
StructureDefinition
StructureMap
Subscription
SubscriptionStatus
SubscriptionTopic
Substance
SubstanceDefinition
SubstanceNucleicAcid
SubstancePolymer
SubstanceProtein
SubstanceReferenceInformation
SubstanceSourceMaterial
SupplyDelivery
SupplyRequest
Task
TerminologyCapabilities
TestPlan
TestReport
TestScript
Transport
ValueSet
VerificationResult
VisionPrescription

*/
