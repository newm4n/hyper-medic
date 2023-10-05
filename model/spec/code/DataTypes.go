package code

import (
	"fmt"
	"strings"
)

type DataTypes string

const (
	DataTypesBase                               DataTypes = "Base"
	DataTypesElement                            DataTypes = "Element"
	DataTypesBackboneElement                    DataTypes = "BackboneElement"
	DataTypesDataType                           DataTypes = "DataType"
	DataTypesAddress                            DataTypes = "Address"
	DataTypesAnnotation                         DataTypes = "Annotation"
	DataTypesAttachment                         DataTypes = "Attachment"
	DataTypesAvailability                       DataTypes = "Availability"
	DataTypesBackboneType                       DataTypes = "BackboneType"
	DataTypesDosage                             DataTypes = "Dosage"
	DataTypesElementDefinition                  DataTypes = "ElementDefinition"
	DataTypesMarketingStatus                    DataTypes = "MarketingStatus"
	DataTypesProductShelfLife                   DataTypes = "ProductShelfLife"
	DataTypesTiming                             DataTypes = "Timing"
	DataTypesCodeableConcept                    DataTypes = "CodeableConcept"
	DataTypesCodeableReference                  DataTypes = "CodeableReference"
	DataTypesCoding                             DataTypes = "Coding"
	DataTypesContactDetail                      DataTypes = "ContactDetail"
	DataTypesContactPoint                       DataTypes = "ContactPoint"
	DataTypesContributor                        DataTypes = "Contributor"
	DataTypesDataRequirement                    DataTypes = "DataRequirement"
	DataTypesExpression                         DataTypes = "Expression"
	DataTypesExtendedContactDetail              DataTypes = "ExtendedContactDetail"
	DataTypesExtension                          DataTypes = "Extension"
	DataTypesHumanName                          DataTypes = "HumanName"
	DataTypesIdentifier                         DataTypes = "Identifier"
	DataTypesMeta                               DataTypes = "Meta"
	DataTypesMonetaryComponent                  DataTypes = "MonetaryComponent"
	DataTypesMoney                              DataTypes = "Money"
	DataTypesNarrative                          DataTypes = "Narrative"
	DataTypesParameterDefinition                DataTypes = "ParameterDefinition"
	DataTypesPeriod                             DataTypes = "Period"
	DataTypesPrimitiveType                      DataTypes = "PrimitiveType"
	DataTypesbase64Binary                       DataTypes = "base64Binary"
	DataTypesboolean                            DataTypes = "boolean"
	DataTypesdate                               DataTypes = "date"
	DataTypesdateTime                           DataTypes = "dateTime"
	DataTypesdecimal                            DataTypes = "decimal"
	DataTypesinstant                            DataTypes = "instant"
	DataTypesinteger                            DataTypes = "integer"
	DataTypespositiveInt                        DataTypes = "positiveInt"
	DataTypesunsignedInt                        DataTypes = "unsignedInt"
	DataTypesinteger64                          DataTypes = "integer64"
	DataTypesstring                             DataTypes = "string"
	DataTypescode                               DataTypes = "code"
	DataTypesid                                 DataTypes = "id"
	DataTypesmarkdown                           DataTypes = "markdown"
	DataTypestime                               DataTypes = "time"
	DataTypesuri                                DataTypes = "uri"
	DataTypescanonical                          DataTypes = "canonical"
	DataTypesoid                                DataTypes = "oid"
	DataTypesurl                                DataTypes = "url"
	DataTypesuuid                               DataTypes = "uuid"
	DataTypesQuantity                           DataTypes = "Quantity"
	DataTypesAge                                DataTypes = "Age"
	DataTypesCount                              DataTypes = "Count"
	DataTypesDistance                           DataTypes = "Distance"
	DataTypesDuration                           DataTypes = "Duration"
	DataTypesRange                              DataTypes = "Range"
	DataTypesRatio                              DataTypes = "Ratio"
	DataTypesRatioRange                         DataTypes = "RatioRange"
	DataTypesReference                          DataTypes = "Reference"
	DataTypesRelatedArtifact                    DataTypes = "RelatedArtifact"
	DataTypesSampledData                        DataTypes = "SampledData"
	DataTypesSignature                          DataTypes = "Signature"
	DataTypesTriggerDefinition                  DataTypes = "TriggerDefinition"
	DataTypesUsageContext                       DataTypes = "UsageContext"
	DataTypesVirtualServiceDetail               DataTypes = "VirtualServiceDetail"
	DataTypesxhtml                              DataTypes = "xhtml"
	DataTypesResource                           DataTypes = "Resource"
	DataTypesBinary                             DataTypes = "Binary"
	DataTypesBundle                             DataTypes = "Bundle"
	DataTypesDomainResource                     DataTypes = "DomainResource"
	DataTypesAccount                            DataTypes = "Account"
	DataTypesActivityDefinition                 DataTypes = "ActivityDefinition"
	DataTypesActorDefinition                    DataTypes = "ActorDefinition"
	DataTypesAdministrableProductDefinition     DataTypes = "AdministrableProductDefinition"
	DataTypesAdverseEvent                       DataTypes = "AdverseEvent"
	DataTypesAllergyIntolerance                 DataTypes = "AllergyIntolerance"
	DataTypesAppointment                        DataTypes = "Appointment"
	DataTypesAppointmentResponse                DataTypes = "AppointmentResponse"
	DataTypesArtifactAssessment                 DataTypes = "ArtifactAssessment"
	DataTypesAuditEvent                         DataTypes = "AuditEvent"
	DataTypesBasic                              DataTypes = "Basic"
	DataTypesBiologicallyDerivedProduct         DataTypes = "BiologicallyDerivedProduct"
	DataTypesBiologicallyDerivedProductDispense DataTypes = "BiologicallyDerivedProductDispense"
	DataTypesBodyStructure                      DataTypes = "BodyStructure"
	DataTypesCanonicalResource                  DataTypes = "CanonicalResource"
	DataTypesCapabilityStatement                DataTypes = "CapabilityStatement"
	DataTypesCarePlan                           DataTypes = "CarePlan"
	DataTypesCareTeam                           DataTypes = "CareTeam"
	DataTypesChargeItem                         DataTypes = "ChargeItem"
	DataTypesChargeItemDefinition               DataTypes = "ChargeItemDefinition"
	DataTypesCitation                           DataTypes = "Citation"
	DataTypesClaim                              DataTypes = "Claim"
	DataTypesClaimResponse                      DataTypes = "ClaimResponse"
	DataTypesClinicalImpression                 DataTypes = "ClinicalImpression"
	DataTypesClinicalUseDefinition              DataTypes = "ClinicalUseDefinition"
	DataTypesCodeSystem                         DataTypes = "CodeSystem"
	DataTypesCommunication                      DataTypes = "Communication"
	DataTypesCommunicationRequest               DataTypes = "CommunicationRequest"
	DataTypesCompartmentDefinition              DataTypes = "CompartmentDefinition"
	DataTypesComposition                        DataTypes = "Composition"
	DataTypesConceptMap                         DataTypes = "ConceptMap"
	DataTypesCondition                          DataTypes = "Condition"
	DataTypesConditionDefinition                DataTypes = "ConditionDefinition"
	DataTypesConsent                            DataTypes = "Consent"
	DataTypesContract                           DataTypes = "Contract"
	DataTypesCoverage                           DataTypes = "Coverage"
	DataTypesCoverageEligibilityRequest         DataTypes = "CoverageEligibilityRequest"
	DataTypesCoverageEligibilityResponse        DataTypes = "CoverageEligibilityResponse"
	DataTypesDetectedIssue                      DataTypes = "DetectedIssue"
	DataTypesDevice                             DataTypes = "Device"
	DataTypesDeviceAssociation                  DataTypes = "DeviceAssociation"
	DataTypesDeviceDefinition                   DataTypes = "DeviceDefinition"
	DataTypesDeviceDispense                     DataTypes = "DeviceDispense"
	DataTypesDeviceMetric                       DataTypes = "DeviceMetric"
	DataTypesDeviceRequest                      DataTypes = "DeviceRequest"
	DataTypesDeviceUsage                        DataTypes = "DeviceUsage"
	DataTypesDiagnosticReport                   DataTypes = "DiagnosticReport"
	DataTypesDocumentReference                  DataTypes = "DocumentReference"
	DataTypesEncounter                          DataTypes = "Encounter"
	DataTypesEncounterHistory                   DataTypes = "EncounterHistory"
	DataTypesEndpoint                           DataTypes = "Endpoint"
	DataTypesEnrollmentRequest                  DataTypes = "EnrollmentRequest"
	DataTypesEnrollmentResponse                 DataTypes = "EnrollmentResponse"
	DataTypesEpisodeOfCare                      DataTypes = "EpisodeOfCare"
	DataTypesEventDefinition                    DataTypes = "EventDefinition"
	DataTypesEvidence                           DataTypes = "Evidence"
	DataTypesEvidenceReport                     DataTypes = "EvidenceReport"
	DataTypesEvidenceVariable                   DataTypes = "EvidenceVariable"
	DataTypesExampleScenario                    DataTypes = "ExampleScenario"
	DataTypesExplanationOfBenefit               DataTypes = "ExplanationOfBenefit"
	DataTypesFamilyMemberHistory                DataTypes = "FamilyMemberHistory"
	DataTypesFlag                               DataTypes = "Flag"
	DataTypesFormularyItem                      DataTypes = "FormularyItem"
	DataTypesGenomicStudy                       DataTypes = "GenomicStudy"
	DataTypesGoal                               DataTypes = "Goal"
	DataTypesGraphDefinition                    DataTypes = "GraphDefinition"
	DataTypesGroup                              DataTypes = "Group"
	DataTypesGuidanceResponse                   DataTypes = "GuidanceResponse"
	DataTypesHealthcareService                  DataTypes = "HealthcareService"
	DataTypesImagingSelection                   DataTypes = "ImagingSelection"
	DataTypesImagingStudy                       DataTypes = "ImagingStudy"
	DataTypesImmunization                       DataTypes = "Immunization"
	DataTypesImmunizationEvaluation             DataTypes = "ImmunizationEvaluation"
	DataTypesImmunizationRecommendation         DataTypes = "ImmunizationRecommendation"
	DataTypesImplementationGuide                DataTypes = "ImplementationGuide"
	DataTypesIngredient                         DataTypes = "Ingredient"
	DataTypesInsurancePlan                      DataTypes = "InsurancePlan"
	DataTypesInventoryItem                      DataTypes = "InventoryItem"
	DataTypesInventoryReport                    DataTypes = "InventoryReport"
	DataTypesInvoice                            DataTypes = "Invoice"
	DataTypesLibrary                            DataTypes = "Library"
	DataTypesLinkage                            DataTypes = "Linkage"
	DataTypesList                               DataTypes = "List"
	DataTypesLocation                           DataTypes = "Location"
	DataTypesManufacturedItemDefinition         DataTypes = "ManufacturedItemDefinition"
	DataTypesMeasure                            DataTypes = "Measure"
	DataTypesMeasureReport                      DataTypes = "MeasureReport"
	DataTypesMedication                         DataTypes = "Medication"
	DataTypesMedicationAdministration           DataTypes = "MedicationAdministration"
	DataTypesMedicationDispense                 DataTypes = "MedicationDispense"
	DataTypesMedicationKnowledge                DataTypes = "MedicationKnowledge"
	DataTypesMedicationRequest                  DataTypes = "MedicationRequest"
	DataTypesMedicationStatement                DataTypes = "MedicationStatement"
	DataTypesMedicinalProductDefinition         DataTypes = "MedicinalProductDefinition"
	DataTypesMessageDefinition                  DataTypes = "MessageDefinition"
	DataTypesMessageHeader                      DataTypes = "MessageHeader"
	DataTypesMetadataResource                   DataTypes = "MetadataResource"
	DataTypesMolecularSequence                  DataTypes = "MolecularSequence"
	DataTypesNamingSystem                       DataTypes = "NamingSystem"
	DataTypesNutritionIntake                    DataTypes = "NutritionIntake"
	DataTypesNutritionOrder                     DataTypes = "NutritionOrder"
	DataTypesNutritionProduct                   DataTypes = "NutritionProduct"
	DataTypesObservation                        DataTypes = "Observation"
	DataTypesObservationDefinition              DataTypes = "ObservationDefinition"
	DataTypesOperationDefinition                DataTypes = "OperationDefinition"
	DataTypesOperationOutcome                   DataTypes = "OperationOutcome"
	DataTypesOrganization                       DataTypes = "Organization"
	DataTypesOrganizationAffiliation            DataTypes = "OrganizationAffiliation"
	DataTypesPackagedProductDefinition          DataTypes = "PackagedProductDefinition"
	DataTypesPatient                            DataTypes = "Patient"
	DataTypesPaymentNotice                      DataTypes = "PaymentNotice"
	DataTypesPaymentReconciliation              DataTypes = "PaymentReconciliation"
	DataTypesPermission                         DataTypes = "Permission"
	DataTypesPerson                             DataTypes = "Person"
	DataTypesPlanDefinition                     DataTypes = "PlanDefinition"
	DataTypesPractitioner                       DataTypes = "Practitioner"
	DataTypesPractitionerRole                   DataTypes = "PractitionerRole"
	DataTypesProcedure                          DataTypes = "Procedure"
	DataTypesProvenance                         DataTypes = "Provenance"
	DataTypesQuestionnaire                      DataTypes = "Questionnaire"
	DataTypesQuestionnaireResponse              DataTypes = "QuestionnaireResponse"
	DataTypesRegulatedAuthorization             DataTypes = "RegulatedAuthorization"
	DataTypesRelatedPerson                      DataTypes = "RelatedPerson"
	DataTypesRequestOrchestration               DataTypes = "RequestOrchestration"
	DataTypesRequirements                       DataTypes = "Requirements"
	DataTypesResearchStudy                      DataTypes = "ResearchStudy"
	DataTypesResearchSubject                    DataTypes = "ResearchSubject"
	DataTypesRiskAssessment                     DataTypes = "RiskAssessment"
	DataTypesSchedule                           DataTypes = "Schedule"
	DataTypesSearchParameter                    DataTypes = "SearchParameter"
	DataTypesServiceRequest                     DataTypes = "ServiceRequest"
	DataTypesSlot                               DataTypes = "Slot"
	DataTypesSpecimen                           DataTypes = "Specimen"
	DataTypesSpecimenDefinition                 DataTypes = "SpecimenDefinition"
	DataTypesStructureDefinition                DataTypes = "StructureDefinition"
	DataTypesStructureMap                       DataTypes = "StructureMap"
	DataTypesSubscription                       DataTypes = "Subscription"
	DataTypesSubscriptionStatus                 DataTypes = "SubscriptionStatus"
	DataTypesSubscriptionTopic                  DataTypes = "SubscriptionTopic"
	DataTypesSubstance                          DataTypes = "Substance"
	DataTypesSubstanceDefinition                DataTypes = "SubstanceDefinition"
	DataTypesSubstanceNucleicAcid               DataTypes = "SubstanceNucleicAcid"
	DataTypesSubstancePolymer                   DataTypes = "SubstancePolymer"
	DataTypesSubstanceProtein                   DataTypes = "SubstanceProtein"
	DataTypesSubstanceReferenceInformation      DataTypes = "SubstanceReferenceInformation"
	DataTypesSubstanceSourceMaterial            DataTypes = "SubstanceSourceMaterial"
	DataTypesSupplyDelivery                     DataTypes = "SupplyDelivery"
	DataTypesSupplyRequest                      DataTypes = "SupplyRequest"
	DataTypesTask                               DataTypes = "Task"
	DataTypesTerminologyCapabilities            DataTypes = "TerminologyCapabilities"
	DataTypesTestPlan                           DataTypes = "TestPlan"
	DataTypesTestReport                         DataTypes = "TestReport"
	DataTypesTestScript                         DataTypes = "TestScript"
	DataTypesTransport                          DataTypes = "Transport"
	DataTypesValueSet                           DataTypes = "ValueSet"
	DataTypesVerificationResult                 DataTypes = "VerificationResult"
	DataTypesVisionPrescription                 DataTypes = "VisionPrescription"
	DataTypesParameters                         DataTypes = "Parameters"
)

func AllDataTypes() []DataTypes {
	return []DataTypes{
		DataTypesBase,
		DataTypesElement,
		DataTypesBackboneElement,
		DataTypesDataType,
		DataTypesAddress,
		DataTypesAnnotation,
		DataTypesAttachment,
		DataTypesAvailability,
		DataTypesBackboneType,
		DataTypesDosage,
		DataTypesElementDefinition,
		DataTypesMarketingStatus,
		DataTypesProductShelfLife,
		DataTypesTiming,
		DataTypesCodeableConcept,
		DataTypesCodeableReference,
		DataTypesCoding,
		DataTypesContactDetail,
		DataTypesContactPoint,
		DataTypesContributor,
		DataTypesDataRequirement,
		DataTypesExpression,
		DataTypesExtendedContactDetail,
		DataTypesExtension,
		DataTypesHumanName,
		DataTypesIdentifier,
		DataTypesMeta,
		DataTypesMonetaryComponent,
		DataTypesMoney,
		DataTypesNarrative,
		DataTypesParameterDefinition,
		DataTypesPeriod,
		DataTypesPrimitiveType,
		DataTypesbase64Binary,
		DataTypesboolean,
		DataTypesdate,
		DataTypesdateTime,
		DataTypesdecimal,
		DataTypesinstant,
		DataTypesinteger,
		DataTypespositiveInt,
		DataTypesunsignedInt,
		DataTypesinteger64,
		DataTypesstring,
		DataTypescode,
		DataTypesid,
		DataTypesmarkdown,
		DataTypestime,
		DataTypesuri,
		DataTypescanonical,
		DataTypesoid,
		DataTypesurl,
		DataTypesuuid,
		DataTypesQuantity,
		DataTypesAge,
		DataTypesCount,
		DataTypesDistance,
		DataTypesDuration,
		DataTypesRange,
		DataTypesRatio,
		DataTypesRatioRange,
		DataTypesReference,
		DataTypesRelatedArtifact,
		DataTypesSampledData,
		DataTypesSignature,
		DataTypesTriggerDefinition,
		DataTypesUsageContext,
		DataTypesVirtualServiceDetail,
		DataTypesxhtml,
		DataTypesResource,
		DataTypesBinary,
		DataTypesBundle,
		DataTypesDomainResource,
		DataTypesAccount,
		DataTypesActivityDefinition,
		DataTypesActorDefinition,
		DataTypesAdministrableProductDefinition,
		DataTypesAdverseEvent,
		DataTypesAllergyIntolerance,
		DataTypesAppointment,
		DataTypesAppointmentResponse,
		DataTypesArtifactAssessment,
		DataTypesAuditEvent,
		DataTypesBasic,
		DataTypesBiologicallyDerivedProduct,
		DataTypesBiologicallyDerivedProductDispense,
		DataTypesBodyStructure,
		DataTypesCanonicalResource,
		DataTypesCapabilityStatement,
		DataTypesCarePlan,
		DataTypesCareTeam,
		DataTypesChargeItem,
		DataTypesChargeItemDefinition,
		DataTypesCitation,
		DataTypesClaim,
		DataTypesClaimResponse,
		DataTypesClinicalImpression,
		DataTypesClinicalUseDefinition,
		DataTypesCodeSystem,
		DataTypesCommunication,
		DataTypesCommunicationRequest,
		DataTypesCompartmentDefinition,
		DataTypesComposition,
		DataTypesConceptMap,
		DataTypesCondition,
		DataTypesConditionDefinition,
		DataTypesConsent,
		DataTypesContract,
		DataTypesCoverage,
		DataTypesCoverageEligibilityRequest,
		DataTypesCoverageEligibilityResponse,
		DataTypesDetectedIssue,
		DataTypesDevice,
		DataTypesDeviceAssociation,
		DataTypesDeviceDefinition,
		DataTypesDeviceDispense,
		DataTypesDeviceMetric,
		DataTypesDeviceRequest,
		DataTypesDeviceUsage,
		DataTypesDiagnosticReport,
		DataTypesDocumentReference,
		DataTypesEncounter,
		DataTypesEncounterHistory,
		DataTypesEndpoint,
		DataTypesEnrollmentRequest,
		DataTypesEnrollmentResponse,
		DataTypesEpisodeOfCare,
		DataTypesEventDefinition,
		DataTypesEvidence,
		DataTypesEvidenceReport,
		DataTypesEvidenceVariable,
		DataTypesExampleScenario,
		DataTypesExplanationOfBenefit,
		DataTypesFamilyMemberHistory,
		DataTypesFlag,
		DataTypesFormularyItem,
		DataTypesGenomicStudy,
		DataTypesGoal,
		DataTypesGraphDefinition,
		DataTypesGroup,
		DataTypesGuidanceResponse,
		DataTypesHealthcareService,
		DataTypesImagingSelection,
		DataTypesImagingStudy,
		DataTypesImmunization,
		DataTypesImmunizationEvaluation,
		DataTypesImmunizationRecommendation,
		DataTypesImplementationGuide,
		DataTypesIngredient,
		DataTypesInsurancePlan,
		DataTypesInventoryItem,
		DataTypesInventoryReport,
		DataTypesInvoice,
		DataTypesLibrary,
		DataTypesLinkage,
		DataTypesList,
		DataTypesLocation,
		DataTypesManufacturedItemDefinition,
		DataTypesMeasure,
		DataTypesMeasureReport,
		DataTypesMedication,
		DataTypesMedicationAdministration,
		DataTypesMedicationDispense,
		DataTypesMedicationKnowledge,
		DataTypesMedicationRequest,
		DataTypesMedicationStatement,
		DataTypesMedicinalProductDefinition,
		DataTypesMessageDefinition,
		DataTypesMessageHeader,
		DataTypesMetadataResource,
		DataTypesMolecularSequence,
		DataTypesNamingSystem,
		DataTypesNutritionIntake,
		DataTypesNutritionOrder,
		DataTypesNutritionProduct,
		DataTypesObservation,
		DataTypesObservationDefinition,
		DataTypesOperationDefinition,
		DataTypesOperationOutcome,
		DataTypesOrganization,
		DataTypesOrganizationAffiliation,
		DataTypesPackagedProductDefinition,
		DataTypesPatient,
		DataTypesPaymentNotice,
		DataTypesPaymentReconciliation,
		DataTypesPermission,
		DataTypesPerson,
		DataTypesPlanDefinition,
		DataTypesPractitioner,
		DataTypesPractitionerRole,
		DataTypesProcedure,
		DataTypesProvenance,
		DataTypesQuestionnaire,
		DataTypesQuestionnaireResponse,
		DataTypesRegulatedAuthorization,
		DataTypesRelatedPerson,
		DataTypesRequestOrchestration,
		DataTypesRequirements,
		DataTypesResearchStudy,
		DataTypesResearchSubject,
		DataTypesRiskAssessment,
		DataTypesSchedule,
		DataTypesSearchParameter,
		DataTypesServiceRequest,
		DataTypesSlot,
		DataTypesSpecimen,
		DataTypesSpecimenDefinition,
		DataTypesStructureDefinition,
		DataTypesStructureMap,
		DataTypesSubscription,
		DataTypesSubscriptionStatus,
		DataTypesSubscriptionTopic,
		DataTypesSubstance,
		DataTypesSubstanceDefinition,
		DataTypesSubstanceNucleicAcid,
		DataTypesSubstancePolymer,
		DataTypesSubstanceProtein,
		DataTypesSubstanceReferenceInformation,
		DataTypesSubstanceSourceMaterial,
		DataTypesSupplyDelivery,
		DataTypesSupplyRequest,
		DataTypesTask,
		DataTypesTerminologyCapabilities,
		DataTypesTestPlan,
		DataTypesTestReport,
		DataTypesTestScript,
		DataTypesTransport,
		DataTypesValueSet,
		DataTypesVerificationResult,
		DataTypesVisionPrescription,
		DataTypesParameters,
	}
}

func FindDataTypes(filter string) []DataTypes {
	ret := make([]DataTypes, 0)
	for _, at := range AllDataTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DataTypes) ToString() {
	fmt.Println(au.String())
}
func (au DataTypes) String() string {
	switch au {
	case DataTypesBase:
		return "Base"
	case DataTypesElement:
		return "Element"
	case DataTypesBackboneElement:
		return "BackboneElement"
	case DataTypesDataType:
		return "DataType"
	case DataTypesAddress:
		return "Address"
	case DataTypesAnnotation:
		return "Annotation"
	case DataTypesAttachment:
		return "Attachment"
	case DataTypesAvailability:
		return "Availability"
	case DataTypesBackboneType:
		return "BackboneType"
	case DataTypesDosage:
		return "Dosage"
	case DataTypesElementDefinition:
		return "ElementDefinition"
	case DataTypesMarketingStatus:
		return "MarketingStatus"
	case DataTypesProductShelfLife:
		return "ProductShelfLife"
	case DataTypesTiming:
		return "Timing"
	case DataTypesCodeableConcept:
		return "CodeableConcept"
	case DataTypesCodeableReference:
		return "CodeableReference"
	case DataTypesCoding:
		return "Coding"
	case DataTypesContactDetail:
		return "ContactDetail"
	case DataTypesContactPoint:
		return "ContactPoint"
	case DataTypesContributor:
		return "Contributor"
	case DataTypesDataRequirement:
		return "DataRequirement"
	case DataTypesExpression:
		return "Expression"
	case DataTypesExtendedContactDetail:
		return "ExtendedContactDetail"
	case DataTypesExtension:
		return "Extension"
	case DataTypesHumanName:
		return "HumanName"
	case DataTypesIdentifier:
		return "Identifier"
	case DataTypesMeta:
		return "Meta"
	case DataTypesMonetaryComponent:
		return "MonetaryComponent"
	case DataTypesMoney:
		return "Money"
	case DataTypesNarrative:
		return "Narrative"
	case DataTypesParameterDefinition:
		return "ParameterDefinition"
	case DataTypesPeriod:
		return "Period"
	case DataTypesPrimitiveType:
		return "PrimitiveType"
	case DataTypesbase64Binary:
		return "base64Binary"
	case DataTypesboolean:
		return "boolean"
	case DataTypesdate:
		return "date"
	case DataTypesdateTime:
		return "dateTime"
	case DataTypesdecimal:
		return "decimal"
	case DataTypesinstant:
		return "instant"
	case DataTypesinteger:
		return "integer"
	case DataTypespositiveInt:
		return "positiveInt"
	case DataTypesunsignedInt:
		return "unsignedInt"
	case DataTypesinteger64:
		return "integer64"
	case DataTypesstring:
		return "string"
	case DataTypescode:
		return "code"
	case DataTypesid:
		return "id"
	case DataTypesmarkdown:
		return "markdown"
	case DataTypestime:
		return "time"
	case DataTypesuri:
		return "uri"
	case DataTypescanonical:
		return "canonical"
	case DataTypesoid:
		return "oid"
	case DataTypesurl:
		return "url"
	case DataTypesuuid:
		return "uuid"
	case DataTypesQuantity:
		return "Quantity"
	case DataTypesAge:
		return "Age"
	case DataTypesCount:
		return "Count"
	case DataTypesDistance:
		return "Distance"
	case DataTypesDuration:
		return "Duration"
	case DataTypesRange:
		return "Range"
	case DataTypesRatio:
		return "Ratio"
	case DataTypesRatioRange:
		return "RatioRange"
	case DataTypesReference:
		return "Reference"
	case DataTypesRelatedArtifact:
		return "RelatedArtifact"
	case DataTypesSampledData:
		return "SampledData"
	case DataTypesSignature:
		return "Signature"
	case DataTypesTriggerDefinition:
		return "TriggerDefinition"
	case DataTypesUsageContext:
		return "UsageContext"
	case DataTypesVirtualServiceDetail:
		return "VirtualServiceDetail"
	case DataTypesxhtml:
		return "xhtml"
	case DataTypesResource:
		return "Resource"
	case DataTypesBinary:
		return "Binary"
	case DataTypesBundle:
		return "Bundle"
	case DataTypesDomainResource:
		return "DomainResource"
	case DataTypesAccount:
		return "Account"
	case DataTypesActivityDefinition:
		return "ActivityDefinition"
	case DataTypesActorDefinition:
		return "ActorDefinition"
	case DataTypesAdministrableProductDefinition:
		return "AdministrableProductDefinition"
	case DataTypesAdverseEvent:
		return "AdverseEvent"
	case DataTypesAllergyIntolerance:
		return "AllergyIntolerance"
	case DataTypesAppointment:
		return "Appointment"
	case DataTypesAppointmentResponse:
		return "AppointmentResponse"
	case DataTypesArtifactAssessment:
		return "ArtifactAssessment"
	case DataTypesAuditEvent:
		return "AuditEvent"
	case DataTypesBasic:
		return "Basic"
	case DataTypesBiologicallyDerivedProduct:
		return "BiologicallyDerivedProduct"
	case DataTypesBiologicallyDerivedProductDispense:
		return "BiologicallyDerivedProductDispense"
	case DataTypesBodyStructure:
		return "BodyStructure"
	case DataTypesCanonicalResource:
		return "CanonicalResource"
	case DataTypesCapabilityStatement:
		return "CapabilityStatement"
	case DataTypesCarePlan:
		return "CarePlan"
	case DataTypesCareTeam:
		return "CareTeam"
	case DataTypesChargeItem:
		return "ChargeItem"
	case DataTypesChargeItemDefinition:
		return "ChargeItemDefinition"
	case DataTypesCitation:
		return "Citation"
	case DataTypesClaim:
		return "Claim"
	case DataTypesClaimResponse:
		return "ClaimResponse"
	case DataTypesClinicalImpression:
		return "ClinicalImpression"
	case DataTypesClinicalUseDefinition:
		return "ClinicalUseDefinition"
	case DataTypesCodeSystem:
		return "CodeSystem"
	case DataTypesCommunication:
		return "Communication"
	case DataTypesCommunicationRequest:
		return "CommunicationRequest"
	case DataTypesCompartmentDefinition:
		return "CompartmentDefinition"
	case DataTypesComposition:
		return "Composition"
	case DataTypesConceptMap:
		return "ConceptMap"
	case DataTypesCondition:
		return "Condition"
	case DataTypesConditionDefinition:
		return "ConditionDefinition"
	case DataTypesConsent:
		return "Consent"
	case DataTypesContract:
		return "Contract"
	case DataTypesCoverage:
		return "Coverage"
	case DataTypesCoverageEligibilityRequest:
		return "CoverageEligibilityRequest"
	case DataTypesCoverageEligibilityResponse:
		return "CoverageEligibilityResponse"
	case DataTypesDetectedIssue:
		return "DetectedIssue"
	case DataTypesDevice:
		return "Device"
	case DataTypesDeviceAssociation:
		return "DeviceAssociation"
	case DataTypesDeviceDefinition:
		return "DeviceDefinition"
	case DataTypesDeviceDispense:
		return "DeviceDispense"
	case DataTypesDeviceMetric:
		return "DeviceMetric"
	case DataTypesDeviceRequest:
		return "DeviceRequest"
	case DataTypesDeviceUsage:
		return "DeviceUsage"
	case DataTypesDiagnosticReport:
		return "DiagnosticReport"
	case DataTypesDocumentReference:
		return "DocumentReference"
	case DataTypesEncounter:
		return "Encounter"
	case DataTypesEncounterHistory:
		return "EncounterHistory"
	case DataTypesEndpoint:
		return "Endpoint"
	case DataTypesEnrollmentRequest:
		return "EnrollmentRequest"
	case DataTypesEnrollmentResponse:
		return "EnrollmentResponse"
	case DataTypesEpisodeOfCare:
		return "EpisodeOfCare"
	case DataTypesEventDefinition:
		return "EventDefinition"
	case DataTypesEvidence:
		return "Evidence"
	case DataTypesEvidenceReport:
		return "EvidenceReport"
	case DataTypesEvidenceVariable:
		return "EvidenceVariable"
	case DataTypesExampleScenario:
		return "ExampleScenario"
	case DataTypesExplanationOfBenefit:
		return "ExplanationOfBenefit"
	case DataTypesFamilyMemberHistory:
		return "FamilyMemberHistory"
	case DataTypesFlag:
		return "Flag"
	case DataTypesFormularyItem:
		return "FormularyItem"
	case DataTypesGenomicStudy:
		return "GenomicStudy"
	case DataTypesGoal:
		return "Goal"
	case DataTypesGraphDefinition:
		return "GraphDefinition"
	case DataTypesGroup:
		return "Group"
	case DataTypesGuidanceResponse:
		return "GuidanceResponse"
	case DataTypesHealthcareService:
		return "HealthcareService"
	case DataTypesImagingSelection:
		return "ImagingSelection"
	case DataTypesImagingStudy:
		return "ImagingStudy"
	case DataTypesImmunization:
		return "Immunization"
	case DataTypesImmunizationEvaluation:
		return "ImmunizationEvaluation"
	case DataTypesImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case DataTypesImplementationGuide:
		return "ImplementationGuide"
	case DataTypesIngredient:
		return "Ingredient"
	case DataTypesInsurancePlan:
		return "InsurancePlan"
	case DataTypesInventoryItem:
		return "InventoryItem"
	case DataTypesInventoryReport:
		return "InventoryReport"
	case DataTypesInvoice:
		return "Invoice"
	case DataTypesLibrary:
		return "Library"
	case DataTypesLinkage:
		return "Linkage"
	case DataTypesList:
		return "List"
	case DataTypesLocation:
		return "Location"
	case DataTypesManufacturedItemDefinition:
		return "ManufacturedItemDefinition"
	case DataTypesMeasure:
		return "Measure"
	case DataTypesMeasureReport:
		return "MeasureReport"
	case DataTypesMedication:
		return "Medication"
	case DataTypesMedicationAdministration:
		return "MedicationAdministration"
	case DataTypesMedicationDispense:
		return "MedicationDispense"
	case DataTypesMedicationKnowledge:
		return "MedicationKnowledge"
	case DataTypesMedicationRequest:
		return "MedicationRequest"
	case DataTypesMedicationStatement:
		return "MedicationStatement"
	case DataTypesMedicinalProductDefinition:
		return "MedicinalProductDefinition"
	case DataTypesMessageDefinition:
		return "MessageDefinition"
	case DataTypesMessageHeader:
		return "MessageHeader"
	case DataTypesMetadataResource:
		return "MetadataResource"
	case DataTypesMolecularSequence:
		return "MolecularSequence"
	case DataTypesNamingSystem:
		return "NamingSystem"
	case DataTypesNutritionIntake:
		return "NutritionIntake"
	case DataTypesNutritionOrder:
		return "NutritionOrder"
	case DataTypesNutritionProduct:
		return "NutritionProduct"
	case DataTypesObservation:
		return "Observation"
	case DataTypesObservationDefinition:
		return "ObservationDefinition"
	case DataTypesOperationDefinition:
		return "OperationDefinition"
	case DataTypesOperationOutcome:
		return "OperationOutcome"
	case DataTypesOrganization:
		return "Organization"
	case DataTypesOrganizationAffiliation:
		return "OrganizationAffiliation"
	case DataTypesPackagedProductDefinition:
		return "PackagedProductDefinition"
	case DataTypesPatient:
		return "Patient"
	case DataTypesPaymentNotice:
		return "PaymentNotice"
	case DataTypesPaymentReconciliation:
		return "PaymentReconciliation"
	case DataTypesPermission:
		return "Permission"
	case DataTypesPerson:
		return "Person"
	case DataTypesPlanDefinition:
		return "PlanDefinition"
	case DataTypesPractitioner:
		return "Practitioner"
	case DataTypesPractitionerRole:
		return "PractitionerRole"
	case DataTypesProcedure:
		return "Procedure"
	case DataTypesProvenance:
		return "Provenance"
	case DataTypesQuestionnaire:
		return "Questionnaire"
	case DataTypesQuestionnaireResponse:
		return "QuestionnaireResponse"
	case DataTypesRegulatedAuthorization:
		return "RegulatedAuthorization"
	case DataTypesRelatedPerson:
		return "RelatedPerson"
	case DataTypesRequestOrchestration:
		return "RequestOrchestration"
	case DataTypesRequirements:
		return "Requirements"
	case DataTypesResearchStudy:
		return "ResearchStudy"
	case DataTypesResearchSubject:
		return "ResearchSubject"
	case DataTypesRiskAssessment:
		return "RiskAssessment"
	case DataTypesSchedule:
		return "Schedule"
	case DataTypesSearchParameter:
		return "SearchParameter"
	case DataTypesServiceRequest:
		return "ServiceRequest"
	case DataTypesSlot:
		return "Slot"
	case DataTypesSpecimen:
		return "Specimen"
	case DataTypesSpecimenDefinition:
		return "SpecimenDefinition"
	case DataTypesStructureDefinition:
		return "StructureDefinition"
	case DataTypesStructureMap:
		return "StructureMap"
	case DataTypesSubscription:
		return "Subscription"
	case DataTypesSubscriptionStatus:
		return "SubscriptionStatus"
	case DataTypesSubscriptionTopic:
		return "SubscriptionTopic"
	case DataTypesSubstance:
		return "Substance"
	case DataTypesSubstanceDefinition:
		return "SubstanceDefinition"
	case DataTypesSubstanceNucleicAcid:
		return "SubstanceNucleicAcid"
	case DataTypesSubstancePolymer:
		return "SubstancePolymer"
	case DataTypesSubstanceProtein:
		return "SubstanceProtein"
	case DataTypesSubstanceReferenceInformation:
		return "SubstanceReferenceInformation"
	case DataTypesSubstanceSourceMaterial:
		return "SubstanceSourceMaterial"
	case DataTypesSupplyDelivery:
		return "SupplyDelivery"
	case DataTypesSupplyRequest:
		return "SupplyRequest"
	case DataTypesTask:
		return "Task"
	case DataTypesTerminologyCapabilities:
		return "TerminologyCapabilities"
	case DataTypesTestPlan:
		return "TestPlan"
	case DataTypesTestReport:
		return "TestReport"
	case DataTypesTestScript:
		return "TestScript"
	case DataTypesTransport:
		return "Transport"
	case DataTypesValueSet:
		return "ValueSet"
	case DataTypesVerificationResult:
		return "VerificationResult"
	case DataTypesVisionPrescription:
		return "VisionPrescription"
	case DataTypesParameters:
		return "Parameters"

	default:
		return "Unknown Data Types"
	}
}
