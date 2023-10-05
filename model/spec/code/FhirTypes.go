package code

import (
	"fmt"
	"strings"
)

type FhirTypes string

const (
	FhirTypesBase                               FhirTypes = "Base"
	FhirTypesElement                            FhirTypes = "Element"
	FhirTypesBackboneElement                    FhirTypes = "BackboneElement"
	FhirTypesDataType                           FhirTypes = "DataType"
	FhirTypesAddress                            FhirTypes = "Address"
	FhirTypesAnnotation                         FhirTypes = "Annotation"
	FhirTypesAttachment                         FhirTypes = "Attachment"
	FhirTypesAvailability                       FhirTypes = "Availability"
	FhirTypesBackboneType                       FhirTypes = "BackboneType"
	FhirTypesDosage                             FhirTypes = "Dosage"
	FhirTypesElementDefinition                  FhirTypes = "ElementDefinition"
	FhirTypesMarketingStatus                    FhirTypes = "MarketingStatus"
	FhirTypesProductShelfLife                   FhirTypes = "ProductShelfLife"
	FhirTypesTiming                             FhirTypes = "Timing"
	FhirTypesCodeableConcept                    FhirTypes = "CodeableConcept"
	FhirTypesCodeableReference                  FhirTypes = "CodeableReference"
	FhirTypesCoding                             FhirTypes = "Coding"
	FhirTypesContactDetail                      FhirTypes = "ContactDetail"
	FhirTypesContactPoint                       FhirTypes = "ContactPoint"
	FhirTypesContributor                        FhirTypes = "Contributor"
	FhirTypesDataRequirement                    FhirTypes = "DataRequirement"
	FhirTypesExpression                         FhirTypes = "Expression"
	FhirTypesExtendedContactDetail              FhirTypes = "ExtendedContactDetail"
	FhirTypesExtension                          FhirTypes = "Extension"
	FhirTypesHumanName                          FhirTypes = "HumanName"
	FhirTypesIdentifier                         FhirTypes = "Identifier"
	FhirTypesMeta                               FhirTypes = "Meta"
	FhirTypesMonetaryComponent                  FhirTypes = "MonetaryComponent"
	FhirTypesMoney                              FhirTypes = "Money"
	FhirTypesNarrative                          FhirTypes = "Narrative"
	FhirTypesParameterDefinition                FhirTypes = "ParameterDefinition"
	FhirTypesPeriod                             FhirTypes = "Period"
	FhirTypesPrimitiveType                      FhirTypes = "PrimitiveType"
	FhirTypesbase64Binary                       FhirTypes = "base64Binary"
	FhirTypesboolean                            FhirTypes = "boolean"
	FhirTypesdate                               FhirTypes = "date"
	FhirTypesdateTime                           FhirTypes = "dateTime"
	FhirTypesdecimal                            FhirTypes = "decimal"
	FhirTypesinstant                            FhirTypes = "instant"
	FhirTypesinteger                            FhirTypes = "integer"
	FhirTypespositiveInt                        FhirTypes = "positiveInt"
	FhirTypesunsignedInt                        FhirTypes = "unsignedInt"
	FhirTypesinteger64                          FhirTypes = "integer64"
	FhirTypesstring                             FhirTypes = "string"
	FhirTypescode                               FhirTypes = "code"
	FhirTypesid                                 FhirTypes = "id"
	FhirTypesmarkdown                           FhirTypes = "markdown"
	FhirTypestime                               FhirTypes = "time"
	FhirTypesuri                                FhirTypes = "uri"
	FhirTypescanonical                          FhirTypes = "canonical"
	FhirTypesoid                                FhirTypes = "oid"
	FhirTypesurl                                FhirTypes = "url"
	FhirTypesuuid                               FhirTypes = "uuid"
	FhirTypesQuantity                           FhirTypes = "Quantity"
	FhirTypesAge                                FhirTypes = "Age"
	FhirTypesCount                              FhirTypes = "Count"
	FhirTypesDistance                           FhirTypes = "Distance"
	FhirTypesDuration                           FhirTypes = "Duration"
	FhirTypesRange                              FhirTypes = "Range"
	FhirTypesRatio                              FhirTypes = "Ratio"
	FhirTypesRatioRange                         FhirTypes = "RatioRange"
	FhirTypesReference                          FhirTypes = "Reference"
	FhirTypesRelatedArtifact                    FhirTypes = "RelatedArtifact"
	FhirTypesSampledData                        FhirTypes = "SampledData"
	FhirTypesSignature                          FhirTypes = "Signature"
	FhirTypesTriggerDefinition                  FhirTypes = "TriggerDefinition"
	FhirTypesUsageContext                       FhirTypes = "UsageContext"
	FhirTypesVirtualServiceDetail               FhirTypes = "VirtualServiceDetail"
	FhirTypesxhtml                              FhirTypes = "xhtml"
	FhirTypesResource                           FhirTypes = "Resource"
	FhirTypesBinary                             FhirTypes = "Binary"
	FhirTypesBundle                             FhirTypes = "Bundle"
	FhirTypesDomainResource                     FhirTypes = "DomainResource"
	FhirTypesAccount                            FhirTypes = "Account"
	FhirTypesActivityDefinition                 FhirTypes = "ActivityDefinition"
	FhirTypesActorDefinition                    FhirTypes = "ActorDefinition"
	FhirTypesAdministrableProductDefinition     FhirTypes = "AdministrableProductDefinition"
	FhirTypesAdverseEvent                       FhirTypes = "AdverseEvent"
	FhirTypesAllergyIntolerance                 FhirTypes = "AllergyIntolerance"
	FhirTypesAppointment                        FhirTypes = "Appointment"
	FhirTypesAppointmentResponse                FhirTypes = "AppointmentResponse"
	FhirTypesArtifactAssessment                 FhirTypes = "ArtifactAssessment"
	FhirTypesAuditEvent                         FhirTypes = "AuditEvent"
	FhirTypesBasic                              FhirTypes = "Basic"
	FhirTypesBiologicallyDerivedProduct         FhirTypes = "BiologicallyDerivedProduct"
	FhirTypesBiologicallyDerivedProductDispense FhirTypes = "BiologicallyDerivedProductDispense"
	FhirTypesBodyStructure                      FhirTypes = "BodyStructure"
	FhirTypesCanonicalResource                  FhirTypes = "CanonicalResource"
	FhirTypesCapabilityStatement                FhirTypes = "CapabilityStatement"
	FhirTypesCarePlan                           FhirTypes = "CarePlan"
	FhirTypesCareTeam                           FhirTypes = "CareTeam"
	FhirTypesChargeItem                         FhirTypes = "ChargeItem"
	FhirTypesChargeItemDefinition               FhirTypes = "ChargeItemDefinition"
	FhirTypesCitation                           FhirTypes = "Citation"
	FhirTypesClaim                              FhirTypes = "Claim"
	FhirTypesClaimResponse                      FhirTypes = "ClaimResponse"
	FhirTypesClinicalImpression                 FhirTypes = "ClinicalImpression"
	FhirTypesClinicalUseDefinition              FhirTypes = "ClinicalUseDefinition"
	FhirTypesCodeSystem                         FhirTypes = "CodeSystem"
	FhirTypesCommunication                      FhirTypes = "Communication"
	FhirTypesCommunicationRequest               FhirTypes = "CommunicationRequest"
	FhirTypesCompartmentDefinition              FhirTypes = "CompartmentDefinition"
	FhirTypesComposition                        FhirTypes = "Composition"
	FhirTypesConceptMap                         FhirTypes = "ConceptMap"
	FhirTypesCondition                          FhirTypes = "Condition"
	FhirTypesConditionDefinition                FhirTypes = "ConditionDefinition"
	FhirTypesConsent                            FhirTypes = "Consent"
	FhirTypesContract                           FhirTypes = "Contract"
	FhirTypesCoverage                           FhirTypes = "Coverage"
	FhirTypesCoverageEligibilityRequest         FhirTypes = "CoverageEligibilityRequest"
	FhirTypesCoverageEligibilityResponse        FhirTypes = "CoverageEligibilityResponse"
	FhirTypesDetectedIssue                      FhirTypes = "DetectedIssue"
	FhirTypesDevice                             FhirTypes = "Device"
	FhirTypesDeviceAssociation                  FhirTypes = "DeviceAssociation"
	FhirTypesDeviceDefinition                   FhirTypes = "DeviceDefinition"
	FhirTypesDeviceDispense                     FhirTypes = "DeviceDispense"
	FhirTypesDeviceMetric                       FhirTypes = "DeviceMetric"
	FhirTypesDeviceRequest                      FhirTypes = "DeviceRequest"
	FhirTypesDeviceUsage                        FhirTypes = "DeviceUsage"
	FhirTypesDiagnosticReport                   FhirTypes = "DiagnosticReport"
	FhirTypesDocumentReference                  FhirTypes = "DocumentReference"
	FhirTypesEncounter                          FhirTypes = "Encounter"
	FhirTypesEncounterHistory                   FhirTypes = "EncounterHistory"
	FhirTypesEndpoint                           FhirTypes = "Endpoint"
	FhirTypesEnrollmentRequest                  FhirTypes = "EnrollmentRequest"
	FhirTypesEnrollmentResponse                 FhirTypes = "EnrollmentResponse"
	FhirTypesEpisodeOfCare                      FhirTypes = "EpisodeOfCare"
	FhirTypesEventDefinition                    FhirTypes = "EventDefinition"
	FhirTypesEvidence                           FhirTypes = "Evidence"
	FhirTypesEvidenceReport                     FhirTypes = "EvidenceReport"
	FhirTypesEvidenceVariable                   FhirTypes = "EvidenceVariable"
	FhirTypesExampleScenario                    FhirTypes = "ExampleScenario"
	FhirTypesExplanationOfBenefit               FhirTypes = "ExplanationOfBenefit"
	FhirTypesFamilyMemberHistory                FhirTypes = "FamilyMemberHistory"
	FhirTypesFlag                               FhirTypes = "Flag"
	FhirTypesFormularyItem                      FhirTypes = "FormularyItem"
	FhirTypesGenomicStudy                       FhirTypes = "GenomicStudy"
	FhirTypesGoal                               FhirTypes = "Goal"
	FhirTypesGraphDefinition                    FhirTypes = "GraphDefinition"
	FhirTypesGroup                              FhirTypes = "Group"
	FhirTypesGuidanceResponse                   FhirTypes = "GuidanceResponse"
	FhirTypesHealthcareService                  FhirTypes = "HealthcareService"
	FhirTypesImagingSelection                   FhirTypes = "ImagingSelection"
	FhirTypesImagingStudy                       FhirTypes = "ImagingStudy"
	FhirTypesImmunization                       FhirTypes = "Immunization"
	FhirTypesImmunizationEvaluation             FhirTypes = "ImmunizationEvaluation"
	FhirTypesImmunizationRecommendation         FhirTypes = "ImmunizationRecommendation"
	FhirTypesImplementationGuide                FhirTypes = "ImplementationGuide"
	FhirTypesIngredient                         FhirTypes = "Ingredient"
	FhirTypesInsurancePlan                      FhirTypes = "InsurancePlan"
	FhirTypesInventoryItem                      FhirTypes = "InventoryItem"
	FhirTypesInventoryReport                    FhirTypes = "InventoryReport"
	FhirTypesInvoice                            FhirTypes = "Invoice"
	FhirTypesLibrary                            FhirTypes = "Library"
	FhirTypesLinkage                            FhirTypes = "Linkage"
	FhirTypesList                               FhirTypes = "List"
	FhirTypesLocation                           FhirTypes = "Location"
	FhirTypesManufacturedItemDefinition         FhirTypes = "ManufacturedItemDefinition"
	FhirTypesMeasure                            FhirTypes = "Measure"
	FhirTypesMeasureReport                      FhirTypes = "MeasureReport"
	FhirTypesMedication                         FhirTypes = "Medication"
	FhirTypesMedicationAdministration           FhirTypes = "MedicationAdministration"
	FhirTypesMedicationDispense                 FhirTypes = "MedicationDispense"
	FhirTypesMedicationKnowledge                FhirTypes = "MedicationKnowledge"
	FhirTypesMedicationRequest                  FhirTypes = "MedicationRequest"
	FhirTypesMedicationStatement                FhirTypes = "MedicationStatement"
	//FhirTypes	FhirTypes = ""
	FhirTypesMedicinalProductDefinition    FhirTypes = "MedicinalProductDefinition"
	FhirTypesMessageDefinition             FhirTypes = "MessageDefinition"
	FhirTypesMessageHeader                 FhirTypes = "MessageHeader"
	FhirTypesMetadataResource              FhirTypes = "MetadataResource"
	FhirTypesMolecularSequence             FhirTypes = "MolecularSequence"
	FhirTypesNamingSystem                  FhirTypes = "NamingSystem"
	FhirTypesNutritionIntake               FhirTypes = "NutritionIntake"
	FhirTypesNutritionOrder                FhirTypes = "NutritionOrder"
	FhirTypesNutritionProduct              FhirTypes = "NutritionProduct"
	FhirTypesObservation                   FhirTypes = "Observation"
	FhirTypesObservationDefinition         FhirTypes = "ObservationDefinition"
	FhirTypesOperationDefinition           FhirTypes = "OperationDefinition"
	FhirTypesOperationOutcome              FhirTypes = "OperationOutcome"
	FhirTypesOrganization                  FhirTypes = "Organization"
	FhirTypesOrganizationAffiliation       FhirTypes = "OrganizationAffiliation"
	FhirTypesPackagedProductDefinition     FhirTypes = "PackagedProductDefinition"
	FhirTypesPatient                       FhirTypes = "Patient"
	FhirTypesPaymentNotice                 FhirTypes = "PaymentNotice"
	FhirTypesPaymentReconciliation         FhirTypes = "PaymentReconciliation"
	FhirTypesPermission                    FhirTypes = "Permission"
	FhirTypesPerson                        FhirTypes = "Person"
	FhirTypesPlanDefinition                FhirTypes = "PlanDefinition"
	FhirTypesPractitioner                  FhirTypes = "Practitioner"
	FhirTypesPractitionerRole              FhirTypes = "PractitionerRole"
	FhirTypesProcedure                     FhirTypes = "Procedure"
	FhirTypesProvenance                    FhirTypes = "Provenance"
	FhirTypesQuestionnaire                 FhirTypes = "Questionnaire"
	FhirTypesQuestionnaireResponse         FhirTypes = "QuestionnaireResponse"
	FhirTypesRegulatedAuthorization        FhirTypes = "RegulatedAuthorization"
	FhirTypesRelatedPerson                 FhirTypes = "RelatedPerson"
	FhirTypesRequestOrchestration          FhirTypes = "RequestOrchestration"
	FhirTypesRequirements                  FhirTypes = "Requirements"
	FhirTypesResearchStudy                 FhirTypes = "ResearchStudy"
	FhirTypesResearchSubject               FhirTypes = "ResearchSubject"
	FhirTypesRiskAssessment                FhirTypes = "RiskAssessment"
	FhirTypesSchedule                      FhirTypes = "Schedule"
	FhirTypesSearchParameter               FhirTypes = "SearchParameter"
	FhirTypesServiceRequest                FhirTypes = "ServiceRequest"
	FhirTypesSlot                          FhirTypes = "Slot"
	FhirTypesSpecimen                      FhirTypes = "Specimen"
	FhirTypesSpecimenDefinition            FhirTypes = "SpecimenDefinition"
	FhirTypesStructureDefinition           FhirTypes = "StructureDefinition"
	FhirTypesStructureMap                  FhirTypes = "StructureMap"
	FhirTypesSubscription                  FhirTypes = "Subscription"
	FhirTypesSubscriptionStatus            FhirTypes = "SubscriptionStatus"
	FhirTypesSubscriptionTopic             FhirTypes = "SubscriptionTopic"
	FhirTypesSubstance                     FhirTypes = "Substance"
	FhirTypesSubstanceDefinition           FhirTypes = "SubstanceDefinition"
	FhirTypesSubstanceNucleicAcid          FhirTypes = "SubstanceNucleicAcid"
	FhirTypesSubstancePolymer              FhirTypes = "SubstancePolymer"
	FhirTypesSubstanceProtein              FhirTypes = "SubstanceProtein"
	FhirTypesSubstanceReferenceInformation FhirTypes = "SubstanceReferenceInformation"
	FhirTypesSubstanceSourceMaterial       FhirTypes = "SubstanceSourceMaterial"
	FhirTypesSupplyDelivery                FhirTypes = "SupplyDelivery"
	FhirTypesSupplyRequest                 FhirTypes = "SupplyRequest"
	FhirTypesTask                          FhirTypes = "Task"
	FhirTypesTerminologyCapabilities       FhirTypes = "TerminologyCapabilities"
	FhirTypesTestPlan                      FhirTypes = "TestPlan"
	FhirTypesTestReport                    FhirTypes = "TestReport"
	FhirTypesTestScript                    FhirTypes = "TestScript"
	FhirTypesTransport                     FhirTypes = "Transport"
	FhirTypesValueSet                      FhirTypes = "ValueSet"
	FhirTypesVerificationResult            FhirTypes = "VerificationResult"
	FhirTypesVisionPrescription            FhirTypes = "VisionPrescription"
	FhirTypesParameters                    FhirTypes = "Parameters"
)

func AllFhirTypes() []FhirTypes {
	return []FhirTypes{
		FhirTypesBase,
		FhirTypesElement,
		FhirTypesBackboneElement,
		FhirTypesDataType,
		FhirTypesAddress,
		FhirTypesAnnotation,
		FhirTypesAttachment,
		FhirTypesAvailability,
		FhirTypesBackboneType,
		FhirTypesDosage,
		FhirTypesElementDefinition,
		FhirTypesMarketingStatus,
		FhirTypesProductShelfLife,
		FhirTypesTiming,
		FhirTypesCodeableConcept,
		FhirTypesCodeableReference,
		FhirTypesCoding,
		FhirTypesContactDetail,
		FhirTypesContactPoint,
		FhirTypesContributor,
		FhirTypesDataRequirement,
		FhirTypesExpression,
		FhirTypesExtendedContactDetail,
		FhirTypesExtension,
		FhirTypesHumanName,
		FhirTypesIdentifier,
		FhirTypesMeta,
		FhirTypesMonetaryComponent,
		FhirTypesMoney,
		FhirTypesNarrative,
		FhirTypesParameterDefinition,
		FhirTypesPeriod,
		FhirTypesPrimitiveType,
		FhirTypesbase64Binary,
		FhirTypesboolean,
		FhirTypesdate,
		FhirTypesdateTime,
		FhirTypesdecimal,
		FhirTypesinstant,
		FhirTypesinteger,
		FhirTypespositiveInt,
		FhirTypesunsignedInt,
		FhirTypesinteger64,
		FhirTypesstring,
		FhirTypescode,
		FhirTypesid,
		FhirTypesmarkdown,
		FhirTypestime,
		FhirTypesuri,
		FhirTypescanonical,
		FhirTypesoid,
		FhirTypesurl,
		FhirTypesuuid,
		FhirTypesQuantity,
		FhirTypesAge,
		FhirTypesCount,
		FhirTypesDistance,
		FhirTypesDuration,
		FhirTypesRange,
		FhirTypesRatio,
		FhirTypesRatioRange,
		FhirTypesReference,
		FhirTypesRelatedArtifact,
		FhirTypesSampledData,
		FhirTypesSignature,
		FhirTypesTriggerDefinition,
		FhirTypesUsageContext,
		FhirTypesVirtualServiceDetail,
		FhirTypesxhtml,
		FhirTypesResource,
		FhirTypesBinary,
		FhirTypesBundle,
		FhirTypesDomainResource,
		FhirTypesAccount,
		FhirTypesActivityDefinition,
		FhirTypesActorDefinition,
		FhirTypesAdministrableProductDefinition,
		FhirTypesAdverseEvent,
		FhirTypesAllergyIntolerance,
		FhirTypesAppointment,
		FhirTypesAppointmentResponse,
		FhirTypesArtifactAssessment,
		FhirTypesAuditEvent,
		FhirTypesBasic,
		FhirTypesBiologicallyDerivedProduct,
		FhirTypesBiologicallyDerivedProductDispense,
		FhirTypesBodyStructure,
		FhirTypesCanonicalResource,
		FhirTypesCapabilityStatement,
		FhirTypesCarePlan,
		FhirTypesCareTeam,
		FhirTypesChargeItem,
		FhirTypesChargeItemDefinition,
		FhirTypesCitation,
		FhirTypesClaim,
		FhirTypesClaimResponse,
		FhirTypesClinicalImpression,
		FhirTypesClinicalUseDefinition,
		FhirTypesCodeSystem,
		FhirTypesCommunication,
		FhirTypesCommunicationRequest,
		FhirTypesCompartmentDefinition,
		FhirTypesComposition,
		FhirTypesConceptMap,
		FhirTypesCondition,
		FhirTypesConditionDefinition,
		FhirTypesConsent,
		FhirTypesContract,
		FhirTypesCoverage,
		FhirTypesCoverageEligibilityRequest,
		FhirTypesCoverageEligibilityResponse,
		FhirTypesDetectedIssue,
		FhirTypesDevice,
		FhirTypesDeviceAssociation,
		FhirTypesDeviceDefinition,
		FhirTypesDeviceDispense,
		FhirTypesDeviceMetric,
		FhirTypesDeviceRequest,
		FhirTypesDeviceUsage,
		FhirTypesDiagnosticReport,
		FhirTypesDocumentReference,
		FhirTypesEncounter,
		FhirTypesEncounterHistory,
		FhirTypesEndpoint,
		FhirTypesEnrollmentRequest,
		FhirTypesEnrollmentResponse,
		FhirTypesEpisodeOfCare,
		FhirTypesEventDefinition,
		FhirTypesEvidence,
		FhirTypesEvidenceReport,
		FhirTypesEvidenceVariable,
		FhirTypesExampleScenario,
		FhirTypesExplanationOfBenefit,
		FhirTypesFamilyMemberHistory,
		FhirTypesFlag,
		FhirTypesFormularyItem,
		FhirTypesGenomicStudy,
		FhirTypesGoal,
		FhirTypesGraphDefinition,
		FhirTypesGroup,
		FhirTypesGuidanceResponse,
		FhirTypesHealthcareService,
		FhirTypesImagingSelection,
		FhirTypesImagingStudy,
		FhirTypesImmunization,
		FhirTypesImmunizationEvaluation,
		FhirTypesImmunizationRecommendation,
		FhirTypesImplementationGuide,
		FhirTypesIngredient,
		FhirTypesInsurancePlan,
		FhirTypesInventoryItem,
		FhirTypesInventoryReport,
		FhirTypesInvoice,
		FhirTypesLibrary,
		FhirTypesLinkage,
		FhirTypesList,
		FhirTypesLocation,
		FhirTypesManufacturedItemDefinition,
		FhirTypesMeasure,
		FhirTypesMeasureReport,
		FhirTypesMedication,
		FhirTypesMedicationAdministration,
		FhirTypesMedicationDispense,
		FhirTypesMedicationKnowledge,
		FhirTypesMedicationRequest,
		FhirTypesMedicationStatement,
		//FhirTypes	,
		FhirTypesMedicinalProductDefinition,
		FhirTypesMessageDefinition,
		FhirTypesMessageHeader,
		FhirTypesMetadataResource,
		FhirTypesMolecularSequence,
		FhirTypesNamingSystem,
		FhirTypesNutritionIntake,
		FhirTypesNutritionOrder,
		FhirTypesNutritionProduct,
		FhirTypesObservation,
		FhirTypesObservationDefinition,
		FhirTypesOperationDefinition,
		FhirTypesOperationOutcome,
		FhirTypesOrganization,
		FhirTypesOrganizationAffiliation,
		FhirTypesPackagedProductDefinition,
		FhirTypesPatient,
		FhirTypesPaymentNotice,
		FhirTypesPaymentReconciliation,
		FhirTypesPermission,
		FhirTypesPerson,
		FhirTypesPlanDefinition,
		FhirTypesPractitioner,
		FhirTypesPractitionerRole,
		FhirTypesProcedure,
		FhirTypesProvenance,
		FhirTypesQuestionnaire,
		FhirTypesQuestionnaireResponse,
		FhirTypesRegulatedAuthorization,
		FhirTypesRelatedPerson,
		FhirTypesRequestOrchestration,
		FhirTypesRequirements,
		FhirTypesResearchStudy,
		FhirTypesResearchSubject,
		FhirTypesRiskAssessment,
		FhirTypesSchedule,
		FhirTypesSearchParameter,
		FhirTypesServiceRequest,
		FhirTypesSlot,
		FhirTypesSpecimen,
		FhirTypesSpecimenDefinition,
		FhirTypesStructureDefinition,
		FhirTypesStructureMap,
		FhirTypesSubscription,
		FhirTypesSubscriptionStatus,
		FhirTypesSubscriptionTopic,
		FhirTypesSubstance,
		FhirTypesSubstanceDefinition,
		FhirTypesSubstanceNucleicAcid,
		FhirTypesSubstancePolymer,
		FhirTypesSubstanceProtein,
		FhirTypesSubstanceReferenceInformation,
		FhirTypesSubstanceSourceMaterial,
		FhirTypesSupplyDelivery,
		FhirTypesSupplyRequest,
		FhirTypesTask,
		FhirTypesTerminologyCapabilities,
		FhirTypesTestPlan,
		FhirTypesTestReport,
		FhirTypesTestScript,
		FhirTypesTransport,
		FhirTypesValueSet,
		FhirTypesVerificationResult,
		FhirTypesVisionPrescription,
		FhirTypesParameters,
	}
}

func FindFhirTypes(filter string) []FhirTypes {
	ret := make([]FhirTypes, 0)
	for _, at := range AllFhirTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au FhirTypes) ToString() {
	fmt.Println(au.String())
}
func (au FhirTypes) String() string {
	switch au {
	case FhirTypesBase:
		return "Base"
	case FhirTypesElement:
		return "Element"
	case FhirTypesBackboneElement:
		return "BackboneElement"
	case FhirTypesDataType:
		return "DataType"
	case FhirTypesAddress:
		return "Address"
	case FhirTypesAnnotation:
		return "Annotation"
	case FhirTypesAttachment:
		return "Attachment"
	case FhirTypesAvailability:
		return "Availability"
	case FhirTypesBackboneType:
		return "BackboneType"
	case FhirTypesDosage:
		return "Dosage"
	case FhirTypesElementDefinition:
		return "ElementDefinition"
	case FhirTypesMarketingStatus:
		return "MarketingStatus"
	case FhirTypesProductShelfLife:
		return "ProductShelfLife"
	case FhirTypesTiming:
		return "Timing"
	case FhirTypesCodeableConcept:
		return "CodeableConcept"
	case FhirTypesCodeableReference:
		return "CodeableReference"
	case FhirTypesCoding:
		return "Coding"
	case FhirTypesContactDetail:
		return "ContactDetail"
	case FhirTypesContactPoint:
		return "ContactPoint"
	case FhirTypesContributor:
		return "Contributor"
	case FhirTypesDataRequirement:
		return "DataRequirement"
	case FhirTypesExpression:
		return "Expression"
	case FhirTypesExtendedContactDetail:
		return "ExtendedContactDetail"
	case FhirTypesExtension:
		return "Extension"
	case FhirTypesHumanName:
		return "HumanName"
	case FhirTypesIdentifier:
		return "Identifier"
	case FhirTypesMeta:
		return "Meta"
	case FhirTypesMonetaryComponent:
		return "MonetaryComponent"
	case FhirTypesMoney:
		return "Money"
	case FhirTypesNarrative:
		return "Narrative"
	case FhirTypesParameterDefinition:
		return "ParameterDefinition"
	case FhirTypesPeriod:
		return "Period"
	case FhirTypesPrimitiveType:
		return "PrimitiveType"
	case FhirTypesbase64Binary:
		return "base64Binary"
	case FhirTypesboolean:
		return "boolean"
	case FhirTypesdate:
		return "date"
	case FhirTypesdateTime:
		return "dateTime"
	case FhirTypesdecimal:
		return "decimal"
	case FhirTypesinstant:
		return "instant"
	case FhirTypesinteger:
		return "integer"
	case FhirTypespositiveInt:
		return "positiveInt"
	case FhirTypesunsignedInt:
		return "unsignedInt"
	case FhirTypesinteger64:
		return "integer64"
	case FhirTypesstring:
		return "string"
	case FhirTypescode:
		return "code"
	case FhirTypesid:
		return "id"
	case FhirTypesmarkdown:
		return "markdown"
	case FhirTypestime:
		return "time"
	case FhirTypesuri:
		return "uri"
	case FhirTypescanonical:
		return "canonical"
	case FhirTypesoid:
		return "oid"
	case FhirTypesurl:
		return "url"
	case FhirTypesuuid:
		return "uuid"
	case FhirTypesQuantity:
		return "Quantity"
	case FhirTypesAge:
		return "Age"
	case FhirTypesCount:
		return "Count"
	case FhirTypesDistance:
		return "Distance"
	case FhirTypesDuration:
		return "Duration"
	case FhirTypesRange:
		return "Range"
	case FhirTypesRatio:
		return "Ratio"
	case FhirTypesRatioRange:
		return "RatioRange"
	case FhirTypesReference:
		return "Reference"
	case FhirTypesRelatedArtifact:
		return "RelatedArtifact"
	case FhirTypesSampledData:
		return "SampledData"
	case FhirTypesSignature:
		return "Signature"
	case FhirTypesTriggerDefinition:
		return "TriggerDefinition"
	case FhirTypesUsageContext:
		return "UsageContext"
	case FhirTypesVirtualServiceDetail:
		return "VirtualServiceDetail"
	case FhirTypesxhtml:
		return "xhtml"
	case FhirTypesResource:
		return "Resource"
	case FhirTypesBinary:
		return "Binary"
	case FhirTypesBundle:
		return "Bundle"
	case FhirTypesDomainResource:
		return "DomainResource"
	case FhirTypesAccount:
		return "Account"
	case FhirTypesActivityDefinition:
		return "ActivityDefinition"
	case FhirTypesActorDefinition:
		return "ActorDefinition"
	case FhirTypesAdministrableProductDefinition:
		return "AdministrableProductDefinition"
	case FhirTypesAdverseEvent:
		return "AdverseEvent"
	case FhirTypesAllergyIntolerance:
		return "AllergyIntolerance"
	case FhirTypesAppointment:
		return "Appointment"
	case FhirTypesAppointmentResponse:
		return "AppointmentResponse"
	case FhirTypesArtifactAssessment:
		return "ArtifactAssessment"
	case FhirTypesAuditEvent:
		return "AuditEvent"
	case FhirTypesBasic:
		return "Basic"
	case FhirTypesBiologicallyDerivedProduct:
		return "BiologicallyDerivedProduct"
	case FhirTypesBiologicallyDerivedProductDispense:
		return "BiologicallyDerivedProductDispense"
	case FhirTypesBodyStructure:
		return "BodyStructure"
	case FhirTypesCanonicalResource:
		return "CanonicalResource"
	case FhirTypesCapabilityStatement:
		return "CapabilityStatement"
	case FhirTypesCarePlan:
		return "CarePlan"
	case FhirTypesCareTeam:
		return "CareTeam"
	case FhirTypesChargeItem:
		return "ChargeItem"
	case FhirTypesChargeItemDefinition:
		return "ChargeItemDefinition"
	case FhirTypesCitation:
		return "Citation"
	case FhirTypesClaim:
		return "Claim"
	case FhirTypesClaimResponse:
		return "ClaimResponse"
	case FhirTypesClinicalImpression:
		return "ClinicalImpression"
	case FhirTypesClinicalUseDefinition:
		return "ClinicalUseDefinition"
	case FhirTypesCodeSystem:
		return "CodeSystem"
	case FhirTypesCommunication:
		return "Communication"
	case FhirTypesCommunicationRequest:
		return "CommunicationRequest"
	case FhirTypesCompartmentDefinition:
		return "CompartmentDefinition"
	case FhirTypesComposition:
		return "Composition"
	case FhirTypesConceptMap:
		return "ConceptMap"
	case FhirTypesCondition:
		return "Condition"
	case FhirTypesConditionDefinition:
		return "ConditionDefinition"
	case FhirTypesConsent:
		return "Consent"
	case FhirTypesContract:
		return "Contract"
	case FhirTypesCoverage:
		return "Coverage"
	case FhirTypesCoverageEligibilityRequest:
		return "CoverageEligibilityRequest"
	case FhirTypesCoverageEligibilityResponse:
		return "CoverageEligibilityResponse"
	case FhirTypesDetectedIssue:
		return "DetectedIssue"
	case FhirTypesDevice:
		return "Device"
	case FhirTypesDeviceAssociation:
		return "DeviceAssociation"
	case FhirTypesDeviceDefinition:
		return "DeviceDefinition"
	case FhirTypesDeviceDispense:
		return "DeviceDispense"
	case FhirTypesDeviceMetric:
		return "DeviceMetric"
	case FhirTypesDeviceRequest:
		return "DeviceRequest"
	case FhirTypesDeviceUsage:
		return "DeviceUsage"
	case FhirTypesDiagnosticReport:
		return "DiagnosticReport"
	case FhirTypesDocumentReference:
		return "DocumentReference"
	case FhirTypesEncounter:
		return "Encounter"
	case FhirTypesEncounterHistory:
		return "EncounterHistory"
	case FhirTypesEndpoint:
		return "Endpoint"
	case FhirTypesEnrollmentRequest:
		return "EnrollmentRequest"
	case FhirTypesEnrollmentResponse:
		return "EnrollmentResponse"
	case FhirTypesEpisodeOfCare:
		return "EpisodeOfCare"
	case FhirTypesEventDefinition:
		return "EventDefinition"
	case FhirTypesEvidence:
		return "Evidence"
	case FhirTypesEvidenceReport:
		return "EvidenceReport"
	case FhirTypesEvidenceVariable:
		return "EvidenceVariable"
	case FhirTypesExampleScenario:
		return "ExampleScenario"
	case FhirTypesExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case FhirTypesFamilyMemberHistory:
		return "FamilyMemberHistory"
	case FhirTypesFlag:
		return "Flag"
	case FhirTypesFormularyItem:
		return "FormularyItem"
	case FhirTypesGenomicStudy:
		return "GenomicStudy"
	case FhirTypesGoal:
		return "Goal"
	case FhirTypesGraphDefinition:
		return "GraphDefinition"
	case FhirTypesGroup:
		return "Group"
	case FhirTypesGuidanceResponse:
		return "GuidanceResponse"
	case FhirTypesHealthcareService:
		return "HealthcareService"
	case FhirTypesImagingSelection:
		return "ImagingSelection"
	case FhirTypesImagingStudy:
		return "ImagingStudy"
	case FhirTypesImmunization:
		return "Immunization"
	case FhirTypesImmunizationEvaluation:
		return "ImmunizationEvaluation"
	case FhirTypesImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case FhirTypesImplementationGuide:
		return "ImplementationGuide"
	case FhirTypesIngredient:
		return "Ingredient"
	case FhirTypesInsurancePlan:
		return "InsurancePlan"
	case FhirTypesInventoryItem:
		return "InventoryItem"
	case FhirTypesInventoryReport:
		return "InventoryReport"
	case FhirTypesInvoice:
		return "Invoice"
	case FhirTypesLibrary:
		return "Library"
	case FhirTypesLinkage:
		return "Linkage"
	case FhirTypesList:
		return "List"
	case FhirTypesLocation:
		return "Location"
	case FhirTypesManufacturedItemDefinition:
		return "ManufacturedItemDefinition"
	case FhirTypesMeasure:
		return "Measure"
	case FhirTypesMeasureReport:
		return "MeasureReport"
	case FhirTypesMedication:
		return "Medication"
	case FhirTypesMedicationAdministration:
		return "MedicationAdministration"
	case FhirTypesMedicationDispense:
		return "MedicationDispense"
	case FhirTypesMedicationKnowledge:
		return "MedicationKnowledge"
	case FhirTypesMedicationRequest:
		return "MedicationRequest"
	case FhirTypesMedicationStatement:
		return "MedicationStatement"
	case FhirTypesMedicinalProductDefinition:
		return "MedicinalProductDefinition"
	case FhirTypesMessageDefinition:
		return "MessageDefinition"
	case FhirTypesMessageHeader:
		return "MessageHeader"
	case FhirTypesMetadataResource:
		return "MetadataResource"
	case FhirTypesMolecularSequence:
		return "MolecularSequence"
	case FhirTypesNamingSystem:
		return "NamingSystem"
	case FhirTypesNutritionIntake:
		return "NutritionIntake"
	case FhirTypesNutritionOrder:
		return "NutritionOrder"
	case FhirTypesNutritionProduct:
		return "NutritionProduct"
	case FhirTypesObservation:
		return "Observation"
	case FhirTypesObservationDefinition:
		return "ObservationDefinition"
	case FhirTypesOperationDefinition:
		return "OperationDefinition"
	case FhirTypesOperationOutcome:
		return "OperationOutcome"
	case FhirTypesOrganization:
		return "Organization"
	case FhirTypesOrganizationAffiliation:
		return "OrganizationAffiliation"
	case FhirTypesPackagedProductDefinition:
		return "PackagedProductDefinition"
	case FhirTypesPatient:
		return "Patient"
	case FhirTypesPaymentNotice:
		return "PaymentNotice"
	case FhirTypesPaymentReconciliation:
		return "PaymentReconciliation"
	case FhirTypesPermission:
		return "Permission"
	case FhirTypesPerson:
		return "Person"
	case FhirTypesPlanDefinition:
		return "PlanDefinition"
	case FhirTypesPractitioner:
		return "Practitioner"
	case FhirTypesPractitionerRole:
		return "PractitionerRole"
	case FhirTypesProcedure:
		return "Procedure"
	case FhirTypesProvenance:
		return "Provenance"
	case FhirTypesQuestionnaire:
		return "Questionnaire"
	case FhirTypesQuestionnaireResponse:
		return "QuestionnaireResponse"
	case FhirTypesRegulatedAuthorization:
		return "RegulatedAuthorization"
	case FhirTypesRelatedPerson:
		return "RelatedPerson"
	case FhirTypesRequestOrchestration:
		return "RequestOrchestration"
	case FhirTypesRequirements:
		return "Requirements"
	case FhirTypesResearchStudy:
		return "ResearchStudy"
	case FhirTypesResearchSubject:
		return "ResearchSubject"
	case FhirTypesRiskAssessment:
		return "RiskAssessment"
	case FhirTypesSchedule:
		return "Schedule"
	case FhirTypesSearchParameter:
		return "SearchParameter"
	case FhirTypesServiceRequest:
		return "ServiceRequest"
	case FhirTypesSlot:
		return "Slot"
	case FhirTypesSpecimen:
		return "Specimen"
	case FhirTypesSpecimenDefinition:
		return "SpecimenDefinition"
	case FhirTypesStructureDefinition:
		return "StructureDefinition"
	case FhirTypesStructureMap:
		return "StructureMap"
	case FhirTypesSubscription:
		return "Subscription"
	case FhirTypesSubscriptionStatus:
		return "SubscriptionStatus"
	case FhirTypesSubscriptionTopic:
		return "SubscriptionTopic"
	case FhirTypesSubstance:
		return "Substance"
	case FhirTypesSubstanceDefinition:
		return "SubstanceDefinition"
	case FhirTypesSubstanceNucleicAcid:
		return "SubstanceNucleicAcid"
	case FhirTypesSubstancePolymer:
		return "SubstancePolymer"
	case FhirTypesSubstanceProtein:
		return "SubstanceProtein"
	case FhirTypesSubstanceReferenceInformation:
		return "SubstanceReferenceInformation"
	case FhirTypesSubstanceSourceMaterial:
		return "SubstanceSourceMaterial"
	case FhirTypesSupplyDelivery:
		return "SupplyDelivery"
	case FhirTypesSupplyRequest:
		return "SupplyRequest"
	case FhirTypesTask:
		return "Task"
	case FhirTypesTerminologyCapabilities:
		return "TerminologyCapabilities"
	case FhirTypesTestPlan:
		return "TestPlan"
	case FhirTypesTestReport:
		return "TestReport"
	case FhirTypesTestScript:
		return "TestScript"
	case FhirTypesTransport:
		return "Transport"
	case FhirTypesValueSet:
		return "ValueSet"
	case FhirTypesVerificationResult:
		return "VerificationResult"
	case FhirTypesVisionPrescription:
		return "VisionPrescription"
	case FhirTypesParameters:
		return "Parameters"
	default:
		return "Unknown Fhir Types"
	}
}
