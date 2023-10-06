package code

import (
	"fmt"
	"strings"
)

type SubscriptionTypes string

const (
	SubscriptionTypesReference                          SubscriptionTypes = "Reference"
	SubscriptionTypesAccount                            SubscriptionTypes = "Account"
	SubscriptionTypesActivityDefinition                 SubscriptionTypes = "ActivityDefinition"
	SubscriptionTypesActorDefinition                    SubscriptionTypes = "ActorDefinition"
	SubscriptionTypesAdministrableProductDefinition     SubscriptionTypes = "AdministrableProductDefinition"
	SubscriptionTypesAdverseEvent                       SubscriptionTypes = "AdverseEvent"
	SubscriptionTypesAllergyIntolerance                 SubscriptionTypes = "AllergyIntolerance"
	SubscriptionTypesAppointment                        SubscriptionTypes = "Appointment"
	SubscriptionTypesAppointmentResponse                SubscriptionTypes = "AppointmentResponse"
	SubscriptionTypesArtifactAssessment                 SubscriptionTypes = "ArtifactAssessment"
	SubscriptionTypesAuditEvent                         SubscriptionTypes = "AuditEvent"
	SubscriptionTypesBasic                              SubscriptionTypes = "Basic"
	SubscriptionTypesBinary                             SubscriptionTypes = "Binary"
	SubscriptionTypesBiologicallyDerivedProduct         SubscriptionTypes = "BiologicallyDerivedProduct"
	SubscriptionTypesBiologicallyDerivedProductDispense SubscriptionTypes = "BiologicallyDerivedProductDispense"
	SubscriptionTypesBodyStructure                      SubscriptionTypes = "BodyStructure"
	SubscriptionTypesBundle                             SubscriptionTypes = "Bundle"
	SubscriptionTypesCapabilityStatement                SubscriptionTypes = "CapabilityStatement"
	SubscriptionTypesCarePlan                           SubscriptionTypes = "CarePlan"
	SubscriptionTypesCareTeam                           SubscriptionTypes = "CareTeam"
	SubscriptionTypesChargeItem                         SubscriptionTypes = "ChargeItem"
	SubscriptionTypesChargeItemDefinition               SubscriptionTypes = "ChargeItemDefinition"
	SubscriptionTypesCitation                           SubscriptionTypes = "Citation"
	SubscriptionTypesClaim                              SubscriptionTypes = "Claim"
	SubscriptionTypesClaimResponse                      SubscriptionTypes = "ClaimResponse"
	SubscriptionTypesClinicalImpression                 SubscriptionTypes = "ClinicalImpression"
	SubscriptionTypesClinicalUseDefinition              SubscriptionTypes = "ClinicalUseDefinition"
	SubscriptionTypesCodeSystem                         SubscriptionTypes = "CodeSystem"
	SubscriptionTypesCommunication                      SubscriptionTypes = "Communication"
	SubscriptionTypesCommunicationRequest               SubscriptionTypes = "CommunicationRequest"
	SubscriptionTypesCompartmentDefinition              SubscriptionTypes = "CompartmentDefinition"
	SubscriptionTypesComposition                        SubscriptionTypes = "Composition"
	SubscriptionTypesConceptMap                         SubscriptionTypes = "ConceptMap"
	SubscriptionTypesCondition                          SubscriptionTypes = "Condition"
	SubscriptionTypesConditionDefinition                SubscriptionTypes = "ConditionDefinition"
	SubscriptionTypesConsent                            SubscriptionTypes = "Consent"
	SubscriptionTypesContract                           SubscriptionTypes = "Contract"
	SubscriptionTypesCoverage                           SubscriptionTypes = "Coverage"
	SubscriptionTypesCoverageEligibilityRequest         SubscriptionTypes = "CoverageEligibilityRequest"
	SubscriptionTypesCoverageEligibilityResponse        SubscriptionTypes = "CoverageEligibilityResponse"
	SubscriptionTypesDetectedIssue                      SubscriptionTypes = "DetectedIssue"
	SubscriptionTypesDevice                             SubscriptionTypes = "Device"
	SubscriptionTypesDeviceAssociation                  SubscriptionTypes = "DeviceAssociation"
	SubscriptionTypesDeviceDefinition                   SubscriptionTypes = "DeviceDefinition"
	SubscriptionTypesDeviceDispense                     SubscriptionTypes = "DeviceDispense"
	SubscriptionTypesDeviceMetric                       SubscriptionTypes = "DeviceMetric"
	SubscriptionTypesDeviceRequest                      SubscriptionTypes = "DeviceRequest"
	SubscriptionTypesDeviceUsage                        SubscriptionTypes = "DeviceUsage"
	SubscriptionTypesDiagnosticReport                   SubscriptionTypes = "DiagnosticReport"
	SubscriptionTypesDocumentReference                  SubscriptionTypes = "DocumentReference"
	SubscriptionTypesEncounter                          SubscriptionTypes = "Encounter"
	SubscriptionTypesEncounterHistory                   SubscriptionTypes = "EncounterHistory"
	SubscriptionTypesEndpoint                           SubscriptionTypes = "Endpoint"
	SubscriptionTypesEnrollmentRequest                  SubscriptionTypes = "EnrollmentRequest"
	SubscriptionTypesEnrollmentResponse                 SubscriptionTypes = "EnrollmentResponse"
	SubscriptionTypesEpisodeOfCare                      SubscriptionTypes = "EpisodeOfCare"
	SubscriptionTypesEventDefinition                    SubscriptionTypes = "EventDefinition"
	SubscriptionTypesEvidence                           SubscriptionTypes = "Evidence"
	SubscriptionTypesEvidenceReport                     SubscriptionTypes = "EvidenceReport"
	SubscriptionTypesEvidenceVariable                   SubscriptionTypes = "EvidenceVariable"
	SubscriptionTypesExampleScenario                    SubscriptionTypes = "ExampleScenario"
	SubscriptionTypesExplanationOfBenefit               SubscriptionTypes = "ExplanationOfBenefit"
	SubscriptionTypesFamilyMemberHistory                SubscriptionTypes = "FamilyMemberHistory"
	SubscriptionTypesFlag                               SubscriptionTypes = "Flag"
	SubscriptionTypesFormularyItem                      SubscriptionTypes = "FormularyItem"
	SubscriptionTypesGenomicStudy                       SubscriptionTypes = "GenomicStudy"
	SubscriptionTypesGoal                               SubscriptionTypes = "Goal"
	SubscriptionTypesGraphDefinition                    SubscriptionTypes = "GraphDefinition"
	SubscriptionTypesGroup                              SubscriptionTypes = "Group"
	SubscriptionTypesGuidanceResponse                   SubscriptionTypes = "GuidanceResponse"
	SubscriptionTypesHealthcareService                  SubscriptionTypes = "HealthcareService"
	SubscriptionTypesImagingSelection                   SubscriptionTypes = "ImagingSelection"
	SubscriptionTypesImagingStudy                       SubscriptionTypes = "ImagingStudy"
	SubscriptionTypesImmunization                       SubscriptionTypes = "Immunization"
	SubscriptionTypesImmunizationEvaluation             SubscriptionTypes = "ImmunizationEvaluation"
	SubscriptionTypesImmunizationRecommendation         SubscriptionTypes = "ImmunizationRecommendation"
	SubscriptionTypesImplementationGuide                SubscriptionTypes = "ImplementationGuide"
	SubscriptionTypesIngredient                         SubscriptionTypes = "Ingredient"
	SubscriptionTypesInsurancePlan                      SubscriptionTypes = "InsurancePlan"
	SubscriptionTypesInventoryItem                      SubscriptionTypes = "InventoryItem"
	SubscriptionTypesInventoryReport                    SubscriptionTypes = "InventoryReport"
	SubscriptionTypesInvoice                            SubscriptionTypes = "Invoice"
	SubscriptionTypesLibrary                            SubscriptionTypes = "Library"
	SubscriptionTypesLinkage                            SubscriptionTypes = "Linkage"
	SubscriptionTypesList                               SubscriptionTypes = "List"
	SubscriptionTypesLocation                           SubscriptionTypes = "Location"
	SubscriptionTypesManufacturedItemDefinition         SubscriptionTypes = "ManufacturedItemDefinition"
	SubscriptionTypesMeasure                            SubscriptionTypes = "Measure"
	SubscriptionTypesMeasureReport                      SubscriptionTypes = "MeasureReport"
	SubscriptionTypesMedication                         SubscriptionTypes = "Medication"
	SubscriptionTypesMedicationAdministration           SubscriptionTypes = "MedicationAdministration"
	SubscriptionTypesMedicationDispense                 SubscriptionTypes = "MedicationDispense"
	SubscriptionTypesMedicationKnowledge                SubscriptionTypes = "MedicationKnowledge"
	SubscriptionTypesMedicationRequest                  SubscriptionTypes = "MedicationRequest"
	SubscriptionTypesMedicationStatement                SubscriptionTypes = "MedicationStatement"
	SubscriptionTypesMedicinalProductDefinition         SubscriptionTypes = "MedicinalProductDefinition"
	SubscriptionTypesMessageDefinition                  SubscriptionTypes = "MessageDefinition"
	SubscriptionTypesMessageHeader                      SubscriptionTypes = "MessageHeader"
	SubscriptionTypesMolecularSequence                  SubscriptionTypes = "MolecularSequence"
	SubscriptionTypesNamingSystem                       SubscriptionTypes = "NamingSystem"
	SubscriptionTypesNutritionIntake                    SubscriptionTypes = "NutritionIntake"
	SubscriptionTypesNutritionOrder                     SubscriptionTypes = "NutritionOrder"
	SubscriptionTypesNutritionProduct                   SubscriptionTypes = "NutritionProduct"
	SubscriptionTypesObservation                        SubscriptionTypes = "Observation"
	SubscriptionTypesObservationDefinition              SubscriptionTypes = "ObservationDefinition"
	SubscriptionTypesOperationDefinition                SubscriptionTypes = "OperationDefinition"
	SubscriptionTypesOperationOutcome                   SubscriptionTypes = "OperationOutcome"
	SubscriptionTypesOrganization                       SubscriptionTypes = "Organization"
	SubscriptionTypesOrganizationAffiliation            SubscriptionTypes = "OrganizationAffiliation"
	SubscriptionTypesPackagedProductDefinition          SubscriptionTypes = "PackagedProductDefinition"
	SubscriptionTypesParameters                         SubscriptionTypes = "Parameters"
	SubscriptionTypesPatient                            SubscriptionTypes = "Patient"
	SubscriptionTypesPaymentNotice                      SubscriptionTypes = "PaymentNotice"
	SubscriptionTypesPaymentReconciliation              SubscriptionTypes = "PaymentReconciliation"
	SubscriptionTypesPermission                         SubscriptionTypes = "Permission"
	SubscriptionTypesPerson                             SubscriptionTypes = "Person"
	SubscriptionTypesPlanDefinition                     SubscriptionTypes = "PlanDefinition"
	SubscriptionTypesPractitioner                       SubscriptionTypes = "Practitioner"
	SubscriptionTypesPractitionerRole                   SubscriptionTypes = "PractitionerRole"
	SubscriptionTypesProcedure                          SubscriptionTypes = "Procedure"
	SubscriptionTypesProvenance                         SubscriptionTypes = "Provenance"
	SubscriptionTypesQuestionnaire                      SubscriptionTypes = "Questionnaire"
	SubscriptionTypesQuestionnaireResponse              SubscriptionTypes = "QuestionnaireResponse"
	SubscriptionTypesRegulatedAuthorization             SubscriptionTypes = "RegulatedAuthorization"
	SubscriptionTypesRelatedPerson                      SubscriptionTypes = "RelatedPerson"
	SubscriptionTypesRequestOrchestration               SubscriptionTypes = "RequestOrchestration"
	SubscriptionTypesRequirements                       SubscriptionTypes = "Requirements"
	SubscriptionTypesResearchStudy                      SubscriptionTypes = "ResearchStudy"
	SubscriptionTypesResearchSubject                    SubscriptionTypes = "ResearchSubject"
	SubscriptionTypesRiskAssessment                     SubscriptionTypes = "RiskAssessment"
	SubscriptionTypesSchedule                           SubscriptionTypes = "Schedule"
	SubscriptionTypesSearchParameter                    SubscriptionTypes = "SearchParameter"
	SubscriptionTypesServiceRequest                     SubscriptionTypes = "ServiceRequest"
	SubscriptionTypesSlot                               SubscriptionTypes = "Slot"
	SubscriptionTypesSpecimen                           SubscriptionTypes = "Specimen"
	SubscriptionTypesSpecimenDefinition                 SubscriptionTypes = "SpecimenDefinition"
	SubscriptionTypesStructureDefinition                SubscriptionTypes = "StructureDefinition"
	SubscriptionTypesStructureMap                       SubscriptionTypes = "StructureMap"
	SubscriptionTypesSubscription                       SubscriptionTypes = "Subscription"
	SubscriptionTypesSubscriptionStatus                 SubscriptionTypes = "SubscriptionStatus"
	SubscriptionTypesSubscriptionTopic                  SubscriptionTypes = "SubscriptionTopic"
	SubscriptionTypesSubstance                          SubscriptionTypes = "Substance"
	SubscriptionTypesSubstanceDefinition                SubscriptionTypes = "SubstanceDefinition"
	SubscriptionTypesSubstanceNucleicAcid               SubscriptionTypes = "SubstanceNucleicAcid"
	SubscriptionTypesSubstancePolymer                   SubscriptionTypes = "SubstancePolymer"
	SubscriptionTypesSubstanceProtein                   SubscriptionTypes = "SubstanceProtein"
	SubscriptionTypesSubstanceReferenceInformation      SubscriptionTypes = "SubstanceReferenceInformation"
	SubscriptionTypesSubstanceSourceMaterial            SubscriptionTypes = "SubstanceSourceMaterial"
	SubscriptionTypesSupplyDelivery                     SubscriptionTypes = "SupplyDelivery"
	SubscriptionTypesSupplyRequest                      SubscriptionTypes = "SupplyRequest"
	SubscriptionTypesTask                               SubscriptionTypes = "Task"
	SubscriptionTypesTerminologyCapabilities            SubscriptionTypes = "TerminologyCapabilities"
	SubscriptionTypesTestPlan                           SubscriptionTypes = "TestPlan"
	SubscriptionTypesTestReport                         SubscriptionTypes = "TestReport"
	SubscriptionTypesTestScript                         SubscriptionTypes = "TestScript"
	SubscriptionTypesTransport                          SubscriptionTypes = "Transport"
	SubscriptionTypesValueSet                           SubscriptionTypes = "ValueSet"
	SubscriptionTypesVerificationResult                 SubscriptionTypes = "VerificationResult"
	SubscriptionTypesVisionPrescription                 SubscriptionTypes = "VisionPrescription"
	SubscriptionTypesBodySite                           SubscriptionTypes = "BodySite"
	SubscriptionTypesCatalogEntry                       SubscriptionTypes = "CatalogEntry"
	SubscriptionTypesConformance                        SubscriptionTypes = "Conformance"
	SubscriptionTypesDataElement                        SubscriptionTypes = "DataElement"
	SubscriptionTypesDeviceComponent                    SubscriptionTypes = "DeviceComponent"
	SubscriptionTypesDeviceUseRequest                   SubscriptionTypes = "DeviceUseRequest"
	SubscriptionTypesDeviceUseStatement                 SubscriptionTypes = "DeviceUseStatement"
	SubscriptionTypesDiagnosticOrder                    SubscriptionTypes = "DiagnosticOrder"
	SubscriptionTypesDocumentManifest                   SubscriptionTypes = "DocumentManifest"
	SubscriptionTypesEffectEvidenceSynthesis            SubscriptionTypes = "EffectEvidenceSynthesis"
	SubscriptionTypesEligibilityRequest                 SubscriptionTypes = "EligibilityRequest"
	SubscriptionTypesEligibilityResponse                SubscriptionTypes = "EligibilityResponse"
	SubscriptionTypesExpansionProfile                   SubscriptionTypes = "ExpansionProfile"
	SubscriptionTypesImagingManifest                    SubscriptionTypes = "ImagingManifest"
	SubscriptionTypesImagingObjectSelection             SubscriptionTypes = "ImagingObjectSelection"
	SubscriptionTypesMedia                              SubscriptionTypes = "Media"
	SubscriptionTypesMedicationOrder                    SubscriptionTypes = "MedicationOrder"
	SubscriptionTypesMedicationUsage                    SubscriptionTypes = "MedicationUsage"
	SubscriptionTypesMedicinalProduct                   SubscriptionTypes = "MedicinalProduct"
	SubscriptionTypesMedicinalProductAuthorization      SubscriptionTypes = "MedicinalProductAuthorization"
	SubscriptionTypesMedicinalProductContraindication   SubscriptionTypes = "MedicinalProductContraindication"
	SubscriptionTypesMedicinalProductIndication         SubscriptionTypes = "MedicinalProductIndication"
	SubscriptionTypesMedicinalProductIngredient         SubscriptionTypes = "MedicinalProductIngredient"
	SubscriptionTypesMedicinalProductInteraction        SubscriptionTypes = "MedicinalProductInteraction"
	SubscriptionTypesMedicinalProductManufactured       SubscriptionTypes = "MedicinalProductManufactured"
	SubscriptionTypesMedicinalProductPackaged           SubscriptionTypes = "MedicinalProductPackaged"
	SubscriptionTypesMedicinalProductPharmaceutical     SubscriptionTypes = "MedicinalProductPharmaceutical"
	SubscriptionTypesMedicinalProductUndesirableEffect  SubscriptionTypes = "MedicinalProductUndesirableEffect"
	SubscriptionTypesOrder                              SubscriptionTypes = "Order"
	SubscriptionTypesOrderResponse                      SubscriptionTypes = "OrderResponse"
	SubscriptionTypesProcedureRequest                   SubscriptionTypes = "ProcedureRequest"
	SubscriptionTypesProcessRequest                     SubscriptionTypes = "ProcessRequest"
	SubscriptionTypesProcessResponse                    SubscriptionTypes = "ProcessResponse"
	SubscriptionTypesReferralRequest                    SubscriptionTypes = "ReferralRequest"
	SubscriptionTypesRequestGroup                       SubscriptionTypes = "RequestGroup"
	SubscriptionTypesResearchDefinition                 SubscriptionTypes = "ResearchDefinition"
	SubscriptionTypesResearchElementDefinition          SubscriptionTypes = "ResearchElementDefinition"
	SubscriptionTypesRiskEvidenceSynthesis              SubscriptionTypes = "RiskEvidenceSynthesis"
	SubscriptionTypesSequence                           SubscriptionTypes = "Sequence"
	SubscriptionTypesServiceDefinition                  SubscriptionTypes = "ServiceDefinition"
	SubscriptionTypesSubstanceSpecification             SubscriptionTypes = "SubstanceSpecification"
)

func AllSubscriptionTypes() []SubscriptionTypes {
	return []SubscriptionTypes{
		SubscriptionTypesReference,
		SubscriptionTypesAccount,
		SubscriptionTypesActivityDefinition,
		SubscriptionTypesActorDefinition,
		SubscriptionTypesAdministrableProductDefinition,
		SubscriptionTypesAdverseEvent,
		SubscriptionTypesAllergyIntolerance,
		SubscriptionTypesAppointment,
		SubscriptionTypesAppointmentResponse,
		SubscriptionTypesArtifactAssessment,
		SubscriptionTypesAuditEvent,
		SubscriptionTypesBasic,
		SubscriptionTypesBinary,
		SubscriptionTypesBiologicallyDerivedProduct,
		SubscriptionTypesBiologicallyDerivedProductDispense,
		SubscriptionTypesBodyStructure,
		SubscriptionTypesBundle,
		SubscriptionTypesCapabilityStatement,
		SubscriptionTypesCarePlan,
		SubscriptionTypesCareTeam,
		SubscriptionTypesChargeItem,
		SubscriptionTypesChargeItemDefinition,
		SubscriptionTypesCitation,
		SubscriptionTypesClaim,
		SubscriptionTypesClaimResponse,
		SubscriptionTypesClinicalImpression,
		SubscriptionTypesClinicalUseDefinition,
		SubscriptionTypesCodeSystem,
		SubscriptionTypesCommunication,
		SubscriptionTypesCommunicationRequest,
		SubscriptionTypesCompartmentDefinition,
		SubscriptionTypesComposition,
		SubscriptionTypesConceptMap,
		SubscriptionTypesCondition,
		SubscriptionTypesConditionDefinition,
		SubscriptionTypesConsent,
		SubscriptionTypesContract,
		SubscriptionTypesCoverage,
		SubscriptionTypesCoverageEligibilityRequest,
		SubscriptionTypesCoverageEligibilityResponse,
		SubscriptionTypesDetectedIssue,
		SubscriptionTypesDevice,
		SubscriptionTypesDeviceAssociation,
		SubscriptionTypesDeviceDefinition,
		SubscriptionTypesDeviceDispense,
		SubscriptionTypesDeviceMetric,
		SubscriptionTypesDeviceRequest,
		SubscriptionTypesDeviceUsage,
		SubscriptionTypesDiagnosticReport,
		SubscriptionTypesDocumentReference,
		SubscriptionTypesEncounter,
		SubscriptionTypesEncounterHistory,
		SubscriptionTypesEndpoint,
		SubscriptionTypesEnrollmentRequest,
		SubscriptionTypesEnrollmentResponse,
		SubscriptionTypesEpisodeOfCare,
		SubscriptionTypesEventDefinition,
		SubscriptionTypesEvidence,
		SubscriptionTypesEvidenceReport,
		SubscriptionTypesEvidenceVariable,
		SubscriptionTypesExampleScenario,
		SubscriptionTypesExplanationOfBenefit,
		SubscriptionTypesFamilyMemberHistory,
		SubscriptionTypesFlag,
		SubscriptionTypesFormularyItem,
		SubscriptionTypesGenomicStudy,
		SubscriptionTypesGoal,
		SubscriptionTypesGraphDefinition,
		SubscriptionTypesGroup,
		SubscriptionTypesGuidanceResponse,
		SubscriptionTypesHealthcareService,
		SubscriptionTypesImagingSelection,
		SubscriptionTypesImagingStudy,
		SubscriptionTypesImmunization,
		SubscriptionTypesImmunizationEvaluation,
		SubscriptionTypesImmunizationRecommendation,
		SubscriptionTypesImplementationGuide,
		SubscriptionTypesIngredient,
		SubscriptionTypesInsurancePlan,
		SubscriptionTypesInventoryItem,
		SubscriptionTypesInventoryReport,
		SubscriptionTypesInvoice,
		SubscriptionTypesLibrary,
		SubscriptionTypesLinkage,
		SubscriptionTypesList,
		SubscriptionTypesLocation,
		SubscriptionTypesManufacturedItemDefinition,
		SubscriptionTypesMeasure,
		SubscriptionTypesMeasureReport,
		SubscriptionTypesMedication,
		SubscriptionTypesMedicationAdministration,
		SubscriptionTypesMedicationDispense,
		SubscriptionTypesMedicationKnowledge,
		SubscriptionTypesMedicationRequest,
		SubscriptionTypesMedicationStatement,
		SubscriptionTypesMedicinalProductDefinition,
		SubscriptionTypesMessageDefinition,
		SubscriptionTypesMessageHeader,
		SubscriptionTypesMolecularSequence,
		SubscriptionTypesNamingSystem,
		SubscriptionTypesNutritionIntake,
		SubscriptionTypesNutritionOrder,
		SubscriptionTypesNutritionProduct,
		SubscriptionTypesObservation,
		SubscriptionTypesObservationDefinition,
		SubscriptionTypesOperationDefinition,
		SubscriptionTypesOperationOutcome,
		SubscriptionTypesOrganization,
		SubscriptionTypesOrganizationAffiliation,
		SubscriptionTypesPackagedProductDefinition,
		SubscriptionTypesParameters,
		SubscriptionTypesPatient,
		SubscriptionTypesPaymentNotice,
		SubscriptionTypesPaymentReconciliation,
		SubscriptionTypesPermission,
		SubscriptionTypesPerson,
		SubscriptionTypesPlanDefinition,
		SubscriptionTypesPractitioner,
		SubscriptionTypesPractitionerRole,
		SubscriptionTypesProcedure,
		SubscriptionTypesProvenance,
		SubscriptionTypesQuestionnaire,
		SubscriptionTypesQuestionnaireResponse,
		SubscriptionTypesRegulatedAuthorization,
		SubscriptionTypesRelatedPerson,
		SubscriptionTypesRequestOrchestration,
		SubscriptionTypesRequirements,
		SubscriptionTypesResearchStudy,
		SubscriptionTypesResearchSubject,
		SubscriptionTypesRiskAssessment,
		SubscriptionTypesSchedule,
		SubscriptionTypesSearchParameter,
		SubscriptionTypesServiceRequest,
		SubscriptionTypesSlot,
		SubscriptionTypesSpecimen,
		SubscriptionTypesSpecimenDefinition,
		SubscriptionTypesStructureDefinition,
		SubscriptionTypesStructureMap,
		SubscriptionTypesSubscription,
		SubscriptionTypesSubscriptionStatus,
		SubscriptionTypesSubscriptionTopic,
		SubscriptionTypesSubstance,
		SubscriptionTypesSubstanceDefinition,
		SubscriptionTypesSubstanceNucleicAcid,
		SubscriptionTypesSubstancePolymer,
		SubscriptionTypesSubstanceProtein,
		SubscriptionTypesSubstanceReferenceInformation,
		SubscriptionTypesSubstanceSourceMaterial,
		SubscriptionTypesSupplyDelivery,
		SubscriptionTypesSupplyRequest,
		SubscriptionTypesTask,
		SubscriptionTypesTerminologyCapabilities,
		SubscriptionTypesTestPlan,
		SubscriptionTypesTestReport,
		SubscriptionTypesTestScript,
		SubscriptionTypesTransport,
		SubscriptionTypesValueSet,
		SubscriptionTypesVerificationResult,
		SubscriptionTypesVisionPrescription,
		SubscriptionTypesBodySite,
		SubscriptionTypesCatalogEntry,
		SubscriptionTypesConformance,
		SubscriptionTypesDataElement,
		SubscriptionTypesDeviceComponent,
		SubscriptionTypesDeviceUseRequest,
		SubscriptionTypesDeviceUseStatement,
		SubscriptionTypesDiagnosticOrder,
		SubscriptionTypesDocumentManifest,
		SubscriptionTypesEffectEvidenceSynthesis,
		SubscriptionTypesEligibilityRequest,
		SubscriptionTypesEligibilityResponse,
		SubscriptionTypesExpansionProfile,
		SubscriptionTypesImagingManifest,
		SubscriptionTypesImagingObjectSelection,
		SubscriptionTypesMedia,
		SubscriptionTypesMedicationOrder,
		SubscriptionTypesMedicationUsage,
		SubscriptionTypesMedicinalProduct,
		SubscriptionTypesMedicinalProductAuthorization,
		SubscriptionTypesMedicinalProductContraindication,
		SubscriptionTypesMedicinalProductIndication,
		SubscriptionTypesMedicinalProductIngredient,
		SubscriptionTypesMedicinalProductInteraction,
		SubscriptionTypesMedicinalProductManufactured,
		SubscriptionTypesMedicinalProductPackaged,
		SubscriptionTypesMedicinalProductPharmaceutical,
		SubscriptionTypesMedicinalProductUndesirableEffect,
		SubscriptionTypesOrder,
		SubscriptionTypesOrderResponse,
		SubscriptionTypesProcedureRequest,
		SubscriptionTypesProcessRequest,
		SubscriptionTypesProcessResponse,
		SubscriptionTypesReferralRequest,
		SubscriptionTypesRequestGroup,
		SubscriptionTypesResearchDefinition,
		SubscriptionTypesResearchElementDefinition,
		SubscriptionTypesRiskEvidenceSynthesis,
		SubscriptionTypesSequence,
		SubscriptionTypesServiceDefinition,
		SubscriptionTypesSubstanceSpecification,
	}
}

func FindSubscriptionTypes(filter string) []SubscriptionTypes {
	ret := make([]SubscriptionTypes, 0)
	for _, at := range AllSubscriptionTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au SubscriptionTypes) ToString() {
	fmt.Println(au.String())
}
func (au SubscriptionTypes) String() string {
	switch au {
	case SubscriptionTypesReference:
		return "Reference"
	case SubscriptionTypesAccount:
		return "Account"
	case SubscriptionTypesActivityDefinition:
		return "ActivityDefinition"
	case SubscriptionTypesActorDefinition:
		return "ActorDefinition"
	case SubscriptionTypesAdministrableProductDefinition:
		return "AdministrableProductDefinition"
	case SubscriptionTypesAdverseEvent:
		return "AdverseEvent"
	case SubscriptionTypesAllergyIntolerance:
		return "AllergyIntolerance"
	case SubscriptionTypesAppointment:
		return "Appointment"
	case SubscriptionTypesAppointmentResponse:
		return "AppointmentResponse"
	case SubscriptionTypesArtifactAssessment:
		return "ArtifactAssessment"
	case SubscriptionTypesAuditEvent:
		return "AuditEvent"
	case SubscriptionTypesBasic:
		return "Basic"
	case SubscriptionTypesBinary:
		return "Binary"
	case SubscriptionTypesBiologicallyDerivedProduct:
		return "BiologicallyDerivedProduct"
	case SubscriptionTypesBiologicallyDerivedProductDispense:
		return "BiologicallyDerivedProductDispense"
	case SubscriptionTypesBodyStructure:
		return "BodyStructure"
	case SubscriptionTypesBundle:
		return "Bundle"
	case SubscriptionTypesCapabilityStatement:
		return "CapabilityStatement"
	case SubscriptionTypesCarePlan:
		return "CarePlan"
	case SubscriptionTypesCareTeam:
		return "CareTeam"
	case SubscriptionTypesChargeItem:
		return "ChargeItem"
	case SubscriptionTypesChargeItemDefinition:
		return "ChargeItemDefinition"
	case SubscriptionTypesCitation:
		return "Citation"
	case SubscriptionTypesClaim:
		return "Claim"
	case SubscriptionTypesClaimResponse:
		return "ClaimResponse"
	case SubscriptionTypesClinicalImpression:
		return "ClinicalImpression"
	case SubscriptionTypesClinicalUseDefinition:
		return "ClinicalUseDefinition"
	case SubscriptionTypesCodeSystem:
		return "CodeSystem"
	case SubscriptionTypesCommunication:
		return "Communication"
	case SubscriptionTypesCommunicationRequest:
		return "CommunicationRequest"
	case SubscriptionTypesCompartmentDefinition:
		return "CompartmentDefinition"
	case SubscriptionTypesComposition:
		return "Composition"
	case SubscriptionTypesConceptMap:
		return "ConceptMap"
	case SubscriptionTypesCondition:
		return "Condition"
	case SubscriptionTypesConditionDefinition:
		return "ConditionDefinition"
	case SubscriptionTypesConsent:
		return "Consent"
	case SubscriptionTypesContract:
		return "Contract"
	case SubscriptionTypesCoverage:
		return "Coverage"
	case SubscriptionTypesCoverageEligibilityRequest:
		return "CoverageEligibilityRequest"
	case SubscriptionTypesCoverageEligibilityResponse:
		return "CoverageEligibilityResponse"
	case SubscriptionTypesDetectedIssue:
		return "DetectedIssue"
	case SubscriptionTypesDevice:
		return "Device"
	case SubscriptionTypesDeviceAssociation:
		return "DeviceAssociation"
	case SubscriptionTypesDeviceDefinition:
		return "DeviceDefinition"
	case SubscriptionTypesDeviceDispense:
		return "DeviceDispense"
	case SubscriptionTypesDeviceMetric:
		return "DeviceMetric"
	case SubscriptionTypesDeviceRequest:
		return "DeviceRequest"
	case SubscriptionTypesDeviceUsage:
		return "DeviceUsage"
	case SubscriptionTypesDiagnosticReport:
		return "DiagnosticReport"
	case SubscriptionTypesDocumentReference:
		return "DocumentReference"
	case SubscriptionTypesEncounter:
		return "Encounter"
	case SubscriptionTypesEncounterHistory:
		return "EncounterHistory"
	case SubscriptionTypesEndpoint:
		return "Endpoint"
	case SubscriptionTypesEnrollmentRequest:
		return "EnrollmentRequest"
	case SubscriptionTypesEnrollmentResponse:
		return "EnrollmentResponse"
	case SubscriptionTypesEpisodeOfCare:
		return "EpisodeOfCare"
	case SubscriptionTypesEventDefinition:
		return "EventDefinition"
	case SubscriptionTypesEvidence:
		return "Evidence"
	case SubscriptionTypesEvidenceReport:
		return "EvidenceReport"
	case SubscriptionTypesEvidenceVariable:
		return "EvidenceVariable"
	case SubscriptionTypesExampleScenario:
		return "ExampleScenario"
	case SubscriptionTypesExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case SubscriptionTypesFamilyMemberHistory:
		return "FamilyMemberHistory"
	case SubscriptionTypesFlag:
		return "Flag"
	case SubscriptionTypesFormularyItem:
		return "FormularyItem"
	case SubscriptionTypesGenomicStudy:
		return "GenomicStudy"
	case SubscriptionTypesGoal:
		return "Goal"
	case SubscriptionTypesGraphDefinition:
		return "GraphDefinition"
	case SubscriptionTypesGroup:
		return "Group"
	case SubscriptionTypesGuidanceResponse:
		return "GuidanceResponse"
	case SubscriptionTypesHealthcareService:
		return "HealthcareService"
	case SubscriptionTypesImagingSelection:
		return "ImagingSelection"
	case SubscriptionTypesImagingStudy:
		return "ImagingStudy"
	case SubscriptionTypesImmunization:
		return "Immunization"
	case SubscriptionTypesImmunizationEvaluation:
		return "ImmunizationEvaluation"
	case SubscriptionTypesImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case SubscriptionTypesImplementationGuide:
		return "ImplementationGuide"
	case SubscriptionTypesIngredient:
		return "Ingredient"
	case SubscriptionTypesInsurancePlan:
		return "InsurancePlan"
	case SubscriptionTypesInventoryItem:
		return "InventoryItem"
	case SubscriptionTypesInventoryReport:
		return "InventoryReport"
	case SubscriptionTypesInvoice:
		return "Invoice"
	case SubscriptionTypesLibrary:
		return "Library"
	case SubscriptionTypesLinkage:
		return "Linkage"
	case SubscriptionTypesList:
		return "List"
	case SubscriptionTypesLocation:
		return "Location"
	case SubscriptionTypesManufacturedItemDefinition:
		return "ManufacturedItemDefinition"
	case SubscriptionTypesMeasure:
		return "Measure"
	case SubscriptionTypesMeasureReport:
		return "MeasureReport"
	case SubscriptionTypesMedication:
		return "Medication"
	case SubscriptionTypesMedicationAdministration:
		return "MedicationAdministration"
	case SubscriptionTypesMedicationDispense:
		return "MedicationDispense"
	case SubscriptionTypesMedicationKnowledge:
		return "MedicationKnowledge"
	case SubscriptionTypesMedicationRequest:
		return "MedicationRequest"
	case SubscriptionTypesMedicationStatement:
		return "MedicationStatement"
	case SubscriptionTypesMedicinalProductDefinition:
		return "MedicinalProductDefinition"
	case SubscriptionTypesMessageDefinition:
		return "MessageDefinition"
	case SubscriptionTypesMessageHeader:
		return "MessageHeader"
	case SubscriptionTypesMolecularSequence:
		return "MolecularSequence"
	case SubscriptionTypesNamingSystem:
		return "NamingSystem"
	case SubscriptionTypesNutritionIntake:
		return "NutritionIntake"
	case SubscriptionTypesNutritionOrder:
		return "NutritionOrder"
	case SubscriptionTypesNutritionProduct:
		return "NutritionProduct"
	case SubscriptionTypesObservation:
		return "Observation"
	case SubscriptionTypesObservationDefinition:
		return "ObservationDefinition"
	case SubscriptionTypesOperationDefinition:
		return "OperationDefinition"
	case SubscriptionTypesOperationOutcome:
		return "OperationOutcome"
	case SubscriptionTypesOrganization:
		return "Organization"
	case SubscriptionTypesOrganizationAffiliation:
		return "OrganizationAffiliation"
	case SubscriptionTypesPackagedProductDefinition:
		return "PackagedProductDefinition"
	case SubscriptionTypesParameters:
		return "Parameters"
	case SubscriptionTypesPatient:
		return "Patient"
	case SubscriptionTypesPaymentNotice:
		return "PaymentNotice"
	case SubscriptionTypesPaymentReconciliation:
		return "PaymentReconciliation"
	case SubscriptionTypesPermission:
		return "Permission"
	case SubscriptionTypesPerson:
		return "Person"
	case SubscriptionTypesPlanDefinition:
		return "PlanDefinition"
	case SubscriptionTypesPractitioner:
		return "Practitioner"
	case SubscriptionTypesPractitionerRole:
		return "PractitionerRole"
	case SubscriptionTypesProcedure:
		return "Procedure"
	case SubscriptionTypesProvenance:
		return "Provenance"
	case SubscriptionTypesQuestionnaire:
		return "Questionnaire"
	case SubscriptionTypesQuestionnaireResponse:
		return "QuestionnaireResponse"
	case SubscriptionTypesRegulatedAuthorization:
		return "RegulatedAuthorization"
	case SubscriptionTypesRelatedPerson:
		return "RelatedPerson"
	case SubscriptionTypesRequestOrchestration:
		return "RequestOrchestration"
	case SubscriptionTypesRequirements:
		return "Requirements"
	case SubscriptionTypesResearchStudy:
		return "ResearchStudy"
	case SubscriptionTypesResearchSubject:
		return "ResearchSubject"
	case SubscriptionTypesRiskAssessment:
		return "RiskAssessment"
	case SubscriptionTypesSchedule:
		return "Schedule"
	case SubscriptionTypesSearchParameter:
		return "SearchParameter"
	case SubscriptionTypesServiceRequest:
		return "ServiceRequest"
	case SubscriptionTypesSlot:
		return "Slot"
	case SubscriptionTypesSpecimen:
		return "Specimen"
	case SubscriptionTypesSpecimenDefinition:
		return "SpecimenDefinition"
	case SubscriptionTypesStructureDefinition:
		return "StructureDefinition"
	case SubscriptionTypesStructureMap:
		return "StructureMap"
	case SubscriptionTypesSubscription:
		return "Subscription"
	case SubscriptionTypesSubscriptionStatus:
		return "SubscriptionStatus"
	case SubscriptionTypesSubscriptionTopic:
		return "SubscriptionTopic"
	case SubscriptionTypesSubstance:
		return "Substance"
	case SubscriptionTypesSubstanceDefinition:
		return "SubstanceDefinition"
	case SubscriptionTypesSubstanceNucleicAcid:
		return "SubstanceNucleicAcid"
	case SubscriptionTypesSubstancePolymer:
		return "SubstancePolymer"
	case SubscriptionTypesSubstanceProtein:
		return "SubstanceProtein"
	case SubscriptionTypesSubstanceReferenceInformation:
		return "SubstanceReferenceInformation"
	case SubscriptionTypesSubstanceSourceMaterial:
		return "SubstanceSourceMaterial"
	case SubscriptionTypesSupplyDelivery:
		return "SupplyDelivery"
	case SubscriptionTypesSupplyRequest:
		return "SupplyRequest"
	case SubscriptionTypesTask:
		return "Task"
	case SubscriptionTypesTerminologyCapabilities:
		return "TerminologyCapabilities"
	case SubscriptionTypesTestPlan:
		return "TestPlan"
	case SubscriptionTypesTestReport:
		return "TestReport"
	case SubscriptionTypesTestScript:
		return "TestScript"
	case SubscriptionTypesTransport:
		return "Transport"
	case SubscriptionTypesValueSet:
		return "ValueSet"
	case SubscriptionTypesVerificationResult:
		return "VerificationResult"
	case SubscriptionTypesVisionPrescription:
		return "VisionPrescription"
	case SubscriptionTypesBodySite:
		return "BodySite"
	case SubscriptionTypesCatalogEntry:
		return "CatalogEntry"
	case SubscriptionTypesConformance:
		return "Conformance"
	case SubscriptionTypesDataElement:
		return "DataElement"
	case SubscriptionTypesDeviceComponent:
		return "DeviceComponent"
	case SubscriptionTypesDeviceUseRequest:
		return "DeviceUseRequest"
	case SubscriptionTypesDeviceUseStatement:
		return "DeviceUseStatement"
	case SubscriptionTypesDiagnosticOrder:
		return "DiagnosticOrder"
	case SubscriptionTypesDocumentManifest:
		return "DocumentManifest"
	case SubscriptionTypesEffectEvidenceSynthesis:
		return "EffectEvidenceSynthesis"
	case SubscriptionTypesEligibilityRequest:
		return "EligibilityRequest"
	case SubscriptionTypesEligibilityResponse:
		return "EligibilityResponse"
	case SubscriptionTypesExpansionProfile:
		return "ExpansionProfile"
	case SubscriptionTypesImagingManifest:
		return "ImagingManifest"
	case SubscriptionTypesImagingObjectSelection:
		return "ImagingObjectSelection"
	case SubscriptionTypesMedia:
		return "Media"
	case SubscriptionTypesMedicationOrder:
		return "MedicationOrder"
	case SubscriptionTypesMedicationUsage:
		return "MedicationUsage"
	case SubscriptionTypesMedicinalProduct:
		return "MedicinalProduct"
	case SubscriptionTypesMedicinalProductAuthorization:
		return "MedicinalProductAuthorization"
	case SubscriptionTypesMedicinalProductContraindication:
		return "MedicinalProductContraindication"
	case SubscriptionTypesMedicinalProductIndication:
		return "MedicinalProductIndication"
	case SubscriptionTypesMedicinalProductIngredient:
		return "MedicinalProductIngredient"
	case SubscriptionTypesMedicinalProductInteraction:
		return "MedicinalProductInteraction"
	case SubscriptionTypesMedicinalProductManufactured:
		return "MedicinalProductManufactured"
	case SubscriptionTypesMedicinalProductPackaged:
		return "MedicinalProductPackaged"
	case SubscriptionTypesMedicinalProductPharmaceutical:
		return "MedicinalProductPharmaceutical"
	case SubscriptionTypesMedicinalProductUndesirableEffect:
		return "MedicinalProductUndesirableEffect"
	case SubscriptionTypesOrder:
		return "Order"
	case SubscriptionTypesOrderResponse:
		return "OrderResponse"
	case SubscriptionTypesProcedureRequest:
		return "ProcedureRequest"
	case SubscriptionTypesProcessRequest:
		return "ProcessRequest"
	case SubscriptionTypesProcessResponse:
		return "ProcessResponse"
	case SubscriptionTypesReferralRequest:
		return "ReferralRequest"
	case SubscriptionTypesRequestGroup:
		return "RequestGroup"
	case SubscriptionTypesResearchDefinition:
		return "ResearchDefinition"
	case SubscriptionTypesResearchElementDefinition:
		return "ResearchElementDefinition"
	case SubscriptionTypesRiskEvidenceSynthesis:
		return "RiskEvidenceSynthesis"
	case SubscriptionTypesSequence:
		return "Sequence"
	case SubscriptionTypesServiceDefinition:
		return "ServiceDefinition"
	case SubscriptionTypesSubstanceSpecification:
		return "SubstanceSpecification"
	default:
		return "Unknown Subscription Types"
	}
}
