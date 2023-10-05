package code

import (
	"fmt"
	"strings"
)

type ElementdefinitionTypes string

const (
	ElementdefinitionTypesStringURL                          ElementdefinitionTypes = "http://hl7.org/fhirpath/System.String"
	ElementdefinitionTypesBooleanURL                         ElementdefinitionTypes = "http://hl7.org/fhirpath/System.Boolean"
	ElementdefinitionTypesDateURL                            ElementdefinitionTypes = "http://hl7.org/fhirpath/System.Date"
	ElementdefinitionTypesDatetimeURL                        ElementdefinitionTypes = "http://hl7.org/fhirpath/System.DateTime"
	ElementdefinitionTypesDecimalURL                         ElementdefinitionTypes = "http://hl7.org/fhirpath/System.Decimal"
	ElementdefinitionTypesIntegerURL                         ElementdefinitionTypes = "http://hl7.org/fhirpath/System.Integer"
	ElementdefinitionTypesTimeURL                            ElementdefinitionTypes = "http://hl7.org/fhirpath/System.Time"
	ElementdefinitionTypesBase                               ElementdefinitionTypes = "Base"
	ElementdefinitionTypesElement                            ElementdefinitionTypes = "Element"
	ElementdefinitionTypesBackboneelement                    ElementdefinitionTypes = "BackboneElement"
	ElementdefinitionTypesDatatype                           ElementdefinitionTypes = "DataType"
	ElementdefinitionTypesAddress                            ElementdefinitionTypes = "Address"
	ElementdefinitionTypesAnnotation                         ElementdefinitionTypes = "Annotation"
	ElementdefinitionTypesAttachment                         ElementdefinitionTypes = "Attachment"
	ElementdefinitionTypesAvailability                       ElementdefinitionTypes = "Availability"
	ElementdefinitionTypesBackbonetype                       ElementdefinitionTypes = "BackboneType"
	ElementdefinitionTypesDosage                             ElementdefinitionTypes = "Dosage"
	ElementdefinitionTypesElementdefinition                  ElementdefinitionTypes = "ElementDefinition"
	ElementdefinitionTypesMarketingstatus                    ElementdefinitionTypes = "MarketingStatus"
	ElementdefinitionTypesProductshelflife                   ElementdefinitionTypes = "ProductShelfLife"
	ElementdefinitionTypesTiming                             ElementdefinitionTypes = "Timing"
	ElementdefinitionTypesCodeableconcept                    ElementdefinitionTypes = "CodeableConcept"
	ElementdefinitionTypesCodeablereference                  ElementdefinitionTypes = "CodeableReference"
	ElementdefinitionTypesCoding                             ElementdefinitionTypes = "Coding"
	ElementdefinitionTypesContactdetail                      ElementdefinitionTypes = "ContactDetail"
	ElementdefinitionTypesContactpoint                       ElementdefinitionTypes = "ContactPoint"
	ElementdefinitionTypesContributor                        ElementdefinitionTypes = "Contributor"
	ElementdefinitionTypesDatarequirement                    ElementdefinitionTypes = "DataRequirement"
	ElementdefinitionTypesExpression                         ElementdefinitionTypes = "Expression"
	ElementdefinitionTypesExtendedcontactdetail              ElementdefinitionTypes = "ExtendedContactDetail"
	ElementdefinitionTypesExtension                          ElementdefinitionTypes = "Extension"
	ElementdefinitionTypesHumanname                          ElementdefinitionTypes = "HumanName"
	ElementdefinitionTypesIdentifier                         ElementdefinitionTypes = "Identifier"
	ElementdefinitionTypesMeta                               ElementdefinitionTypes = "Meta"
	ElementdefinitionTypesMonetarycomponent                  ElementdefinitionTypes = "MonetaryComponent"
	ElementdefinitionTypesMoney                              ElementdefinitionTypes = "Money"
	ElementdefinitionTypesNarrative                          ElementdefinitionTypes = "Narrative"
	ElementdefinitionTypesParameterdefinition                ElementdefinitionTypes = "ParameterDefinition"
	ElementdefinitionTypesPeriod                             ElementdefinitionTypes = "Period"
	ElementdefinitionTypesPrimitivetype                      ElementdefinitionTypes = "PrimitiveType"
	ElementdefinitionTypesBase64Binary                       ElementdefinitionTypes = "base64Binary"
	ElementdefinitionTypesBoolean                            ElementdefinitionTypes = "boolean"
	ElementdefinitionTypesDate                               ElementdefinitionTypes = "date"
	ElementdefinitionTypesDatetime                           ElementdefinitionTypes = "dateTime"
	ElementdefinitionTypesDecimal                            ElementdefinitionTypes = "decimal"
	ElementdefinitionTypesInstant                            ElementdefinitionTypes = "instant"
	ElementdefinitionTypesInteger                            ElementdefinitionTypes = "integer"
	ElementdefinitionTypesPositiveint                        ElementdefinitionTypes = "positiveInt"
	ElementdefinitionTypesUnsignedint                        ElementdefinitionTypes = "unsignedInt"
	ElementdefinitionTypesInteger64                          ElementdefinitionTypes = "integer64"
	ElementdefinitionTypesString                             ElementdefinitionTypes = "string"
	ElementdefinitionTypesCode                               ElementdefinitionTypes = "code"
	ElementdefinitionTypesId                                 ElementdefinitionTypes = "id"
	ElementdefinitionTypesMarkdown                           ElementdefinitionTypes = "markdown"
	ElementdefinitionTypesTime                               ElementdefinitionTypes = "time"
	ElementdefinitionTypesUri                                ElementdefinitionTypes = "uri"
	ElementdefinitionTypesCanonical                          ElementdefinitionTypes = "canonical"
	ElementdefinitionTypesOid                                ElementdefinitionTypes = "oid"
	ElementdefinitionTypesUrl                                ElementdefinitionTypes = "url"
	ElementdefinitionTypesUuid                               ElementdefinitionTypes = "uuid"
	ElementdefinitionTypesQuantity                           ElementdefinitionTypes = "Quantity"
	ElementdefinitionTypesAge                                ElementdefinitionTypes = "Age"
	ElementdefinitionTypesCount                              ElementdefinitionTypes = "Count"
	ElementdefinitionTypesDistance                           ElementdefinitionTypes = "Distance"
	ElementdefinitionTypesDuration                           ElementdefinitionTypes = "Duration"
	ElementdefinitionTypesRange                              ElementdefinitionTypes = "Range"
	ElementdefinitionTypesRatio                              ElementdefinitionTypes = "Ratio"
	ElementdefinitionTypesRatiorange                         ElementdefinitionTypes = "RatioRange"
	ElementdefinitionTypesReference                          ElementdefinitionTypes = "Reference"
	ElementdefinitionTypesRelatedartifact                    ElementdefinitionTypes = "RelatedArtifact"
	ElementdefinitionTypesSampleddata                        ElementdefinitionTypes = "SampledData"
	ElementdefinitionTypesSignature                          ElementdefinitionTypes = "Signature"
	ElementdefinitionTypesTriggerdefinition                  ElementdefinitionTypes = "TriggerDefinition"
	ElementdefinitionTypesUsagecontext                       ElementdefinitionTypes = "UsageContext"
	ElementdefinitionTypesVirtualservicedetail               ElementdefinitionTypes = "VirtualServiceDetail"
	ElementdefinitionTypesXhtml                              ElementdefinitionTypes = "xhtml"
	ElementdefinitionTypesResource                           ElementdefinitionTypes = "Resource"
	ElementdefinitionTypesBinary                             ElementdefinitionTypes = "Binary"
	ElementdefinitionTypesBundle                             ElementdefinitionTypes = "Bundle"
	ElementdefinitionTypesDomainresource                     ElementdefinitionTypes = "DomainResource"
	ElementdefinitionTypesAccount                            ElementdefinitionTypes = "Account"
	ElementdefinitionTypesActivitydefinition                 ElementdefinitionTypes = "ActivityDefinition"
	ElementdefinitionTypesActordefinition                    ElementdefinitionTypes = "ActorDefinition"
	ElementdefinitionTypesAdministrableproductdefinition     ElementdefinitionTypes = "AdministrableProductDefinition"
	ElementdefinitionTypesAdverseevent                       ElementdefinitionTypes = "AdverseEvent"
	ElementdefinitionTypesAllergyintolerance                 ElementdefinitionTypes = "AllergyIntolerance"
	ElementdefinitionTypesAppointment                        ElementdefinitionTypes = "Appointment"
	ElementdefinitionTypesAppointmentresponse                ElementdefinitionTypes = "AppointmentResponse"
	ElementdefinitionTypesArtifactassessment                 ElementdefinitionTypes = "ArtifactAssessment"
	ElementdefinitionTypesAuditevent                         ElementdefinitionTypes = "AuditEvent"
	ElementdefinitionTypesBasic                              ElementdefinitionTypes = "Basic"
	ElementdefinitionTypesBiologicallyderivedproduct         ElementdefinitionTypes = "BiologicallyDerivedProduct"
	ElementdefinitionTypesBiologicallyderivedproductdispense ElementdefinitionTypes = "BiologicallyDerivedProductDispense"
	ElementdefinitionTypesBodystructure                      ElementdefinitionTypes = "BodyStructure"
	ElementdefinitionTypesCanonicalresource                  ElementdefinitionTypes = "CanonicalResource"
	ElementdefinitionTypesCapabilitystatement                ElementdefinitionTypes = "CapabilityStatement"
	ElementdefinitionTypesCareplan                           ElementdefinitionTypes = "CarePlan"
	ElementdefinitionTypesCareteam                           ElementdefinitionTypes = "CareTeam"
	ElementdefinitionTypesChargeitem                         ElementdefinitionTypes = "ChargeItem"
	ElementdefinitionTypesChargeitemdefinition               ElementdefinitionTypes = "ChargeItemDefinition"
	ElementdefinitionTypesCitation                           ElementdefinitionTypes = "Citation"
	ElementdefinitionTypesClaim                              ElementdefinitionTypes = "Claim"
	ElementdefinitionTypesClaimresponse                      ElementdefinitionTypes = "ClaimResponse"
	ElementdefinitionTypesClinicalimpression                 ElementdefinitionTypes = "ClinicalImpression"
	ElementdefinitionTypesClinicalusedefinition              ElementdefinitionTypes = "ClinicalUseDefinition"
	ElementdefinitionTypesCodesystem                         ElementdefinitionTypes = "CodeSystem"
	ElementdefinitionTypesCommunication                      ElementdefinitionTypes = "Communication"
	ElementdefinitionTypesCommunicationrequest               ElementdefinitionTypes = "CommunicationRequest"
	ElementdefinitionTypesCompartmentdefinition              ElementdefinitionTypes = "CompartmentDefinition"
	ElementdefinitionTypesComposition                        ElementdefinitionTypes = "Composition"
	ElementdefinitionTypesConceptmap                         ElementdefinitionTypes = "ConceptMap"
	ElementdefinitionTypesCondition                          ElementdefinitionTypes = "Condition"
	ElementdefinitionTypesConditiondefinition                ElementdefinitionTypes = "ConditionDefinition"
	ElementdefinitionTypesConsent                            ElementdefinitionTypes = "Consent"
	ElementdefinitionTypesContract                           ElementdefinitionTypes = "Contract"
	ElementdefinitionTypesCoverage                           ElementdefinitionTypes = "Coverage"
	ElementdefinitionTypesCoverageeligibilityrequest         ElementdefinitionTypes = "CoverageEligibilityRequest"
	ElementdefinitionTypesCoverageeligibilityresponse        ElementdefinitionTypes = "CoverageEligibilityResponse"
	ElementdefinitionTypesDetectedissue                      ElementdefinitionTypes = "DetectedIssue"
	ElementdefinitionTypesDevice                             ElementdefinitionTypes = "Device"
	ElementdefinitionTypesDeviceassociation                  ElementdefinitionTypes = "DeviceAssociation"
	ElementdefinitionTypesDevicedefinition                   ElementdefinitionTypes = "DeviceDefinition"
	ElementdefinitionTypesDevicedispense                     ElementdefinitionTypes = "DeviceDispense"
	ElementdefinitionTypesDevicemetric                       ElementdefinitionTypes = "DeviceMetric"
	ElementdefinitionTypesDevicerequest                      ElementdefinitionTypes = "DeviceRequest"
	ElementdefinitionTypesDeviceusage                        ElementdefinitionTypes = "DeviceUsage"
	ElementdefinitionTypesDiagnosticreport                   ElementdefinitionTypes = "DiagnosticReport"
	ElementdefinitionTypesDocumentreference                  ElementdefinitionTypes = "DocumentReference"
	ElementdefinitionTypesEncounter                          ElementdefinitionTypes = "Encounter"
	ElementdefinitionTypesEncounterhistory                   ElementdefinitionTypes = "EncounterHistory"
	ElementdefinitionTypesEndpoint                           ElementdefinitionTypes = "Endpoint"
	ElementdefinitionTypesEnrollmentrequest                  ElementdefinitionTypes = "EnrollmentRequest"
	ElementdefinitionTypesEnrollmentresponse                 ElementdefinitionTypes = "EnrollmentResponse"
	ElementdefinitionTypesEpisodeofcare                      ElementdefinitionTypes = "EpisodeOfCare"
	ElementdefinitionTypesEventdefinition                    ElementdefinitionTypes = "EventDefinition"
	ElementdefinitionTypesEvidence                           ElementdefinitionTypes = "Evidence"
	ElementdefinitionTypesEvidencereport                     ElementdefinitionTypes = "EvidenceReport"
	ElementdefinitionTypesEvidencevariable                   ElementdefinitionTypes = "EvidenceVariable"
	ElementdefinitionTypesExamplescenario                    ElementdefinitionTypes = "ExampleScenario"
	ElementdefinitionTypesExplanationofbenefit               ElementdefinitionTypes = "ExplanationOfBenefit"
	ElementdefinitionTypesFamilymemberhistory                ElementdefinitionTypes = "FamilyMemberHistory"
	ElementdefinitionTypesFlag                               ElementdefinitionTypes = "Flag"
	ElementdefinitionTypesFormularyitem                      ElementdefinitionTypes = "FormularyItem"
	ElementdefinitionTypesGenomicstudy                       ElementdefinitionTypes = "GenomicStudy"
	ElementdefinitionTypesGoal                               ElementdefinitionTypes = "Goal"
	ElementdefinitionTypesGraphdefinition                    ElementdefinitionTypes = "GraphDefinition"
	ElementdefinitionTypesGroup                              ElementdefinitionTypes = "Group"
	ElementdefinitionTypesGuidanceresponse                   ElementdefinitionTypes = "GuidanceResponse"
	ElementdefinitionTypesHealthcareservice                  ElementdefinitionTypes = "HealthcareService"
	ElementdefinitionTypesImagingselection                   ElementdefinitionTypes = "ImagingSelection"
	ElementdefinitionTypesImagingstudy                       ElementdefinitionTypes = "ImagingStudy"
	ElementdefinitionTypesImmunization                       ElementdefinitionTypes = "Immunization"
	ElementdefinitionTypesImmunizationevaluation             ElementdefinitionTypes = "ImmunizationEvaluation"
	ElementdefinitionTypesImmunizationrecommendation         ElementdefinitionTypes = "ImmunizationRecommendation"
	ElementdefinitionTypesImplementationguide                ElementdefinitionTypes = "ImplementationGuide"
	ElementdefinitionTypesIngredient                         ElementdefinitionTypes = "Ingredient"
	ElementdefinitionTypesInsuranceplan                      ElementdefinitionTypes = "InsurancePlan"
	ElementdefinitionTypesInventoryitem                      ElementdefinitionTypes = "InventoryItem"
	ElementdefinitionTypesInventoryreport                    ElementdefinitionTypes = "InventoryReport"
	ElementdefinitionTypesInvoice                            ElementdefinitionTypes = "Invoice"
	ElementdefinitionTypesLibrary                            ElementdefinitionTypes = "Library"
	ElementdefinitionTypesLinkage                            ElementdefinitionTypes = "Linkage"
	ElementdefinitionTypesList                               ElementdefinitionTypes = "List"
	ElementdefinitionTypesLocation                           ElementdefinitionTypes = "Location"
	ElementdefinitionTypesManufactureditemdefinition         ElementdefinitionTypes = "ManufacturedItemDefinition"
	ElementdefinitionTypesMeasure                            ElementdefinitionTypes = "Measure"
	ElementdefinitionTypesMeasurereport                      ElementdefinitionTypes = "MeasureReport"
	ElementdefinitionTypesMedication                         ElementdefinitionTypes = "Medication"
	ElementdefinitionTypesMedicationadministration           ElementdefinitionTypes = "MedicationAdministration"
	ElementdefinitionTypesMedicationdispense                 ElementdefinitionTypes = "MedicationDispense"
	ElementdefinitionTypesMedicationknowledge                ElementdefinitionTypes = "MedicationKnowledge"
	ElementdefinitionTypesMedicationrequest                  ElementdefinitionTypes = "MedicationRequest"
	ElementdefinitionTypesMedicationstatement                ElementdefinitionTypes = "MedicationStatement"
	ElementdefinitionTypesMedicinalproductdefinition         ElementdefinitionTypes = "MedicinalProductDefinition"
	ElementdefinitionTypesMessagedefinition                  ElementdefinitionTypes = "MessageDefinition"
	ElementdefinitionTypesMessageheader                      ElementdefinitionTypes = "MessageHeader"
	ElementdefinitionTypesMetadataresource                   ElementdefinitionTypes = "MetadataResource"
	ElementdefinitionTypesMolecularsequence                  ElementdefinitionTypes = "MolecularSequence"
	ElementdefinitionTypesNamingsystem                       ElementdefinitionTypes = "NamingSystem"
	ElementdefinitionTypesNutritionintake                    ElementdefinitionTypes = "NutritionIntake"
	ElementdefinitionTypesNutritionorder                     ElementdefinitionTypes = "NutritionOrder"
	ElementdefinitionTypesNutritionproduct                   ElementdefinitionTypes = "NutritionProduct"
	ElementdefinitionTypesObservation                        ElementdefinitionTypes = "Observation"
	ElementdefinitionTypesObservationdefinition              ElementdefinitionTypes = "ObservationDefinition"
	ElementdefinitionTypesOperationdefinition                ElementdefinitionTypes = "OperationDefinition"
	ElementdefinitionTypesOperationoutcome                   ElementdefinitionTypes = "OperationOutcome"
	ElementdefinitionTypesOrganization                       ElementdefinitionTypes = "Organization"
	ElementdefinitionTypesOrganizationaffiliation            ElementdefinitionTypes = "OrganizationAffiliation"
	ElementdefinitionTypesPackagedproductdefinition          ElementdefinitionTypes = "PackagedProductDefinition"
	ElementdefinitionTypesPatient                            ElementdefinitionTypes = "Patient"
	ElementdefinitionTypesPaymentnotice                      ElementdefinitionTypes = "PaymentNotice"
	ElementdefinitionTypesPaymentreconciliation              ElementdefinitionTypes = "PaymentReconciliation"
	ElementdefinitionTypesPermission                         ElementdefinitionTypes = "Permission"
	ElementdefinitionTypesPerson                             ElementdefinitionTypes = "Person"
	ElementdefinitionTypesPlandefinition                     ElementdefinitionTypes = "PlanDefinition"
	ElementdefinitionTypesPractitioner                       ElementdefinitionTypes = "Practitioner"
	ElementdefinitionTypesPractitionerrole                   ElementdefinitionTypes = "PractitionerRole"
	ElementdefinitionTypesProcedure                          ElementdefinitionTypes = "Procedure"
	ElementdefinitionTypesProvenance                         ElementdefinitionTypes = "Provenance"
	ElementdefinitionTypesQuestionnaire                      ElementdefinitionTypes = "Questionnaire"
	ElementdefinitionTypesQuestionnaireresponse              ElementdefinitionTypes = "QuestionnaireResponse"
	ElementdefinitionTypesRegulatedauthorization             ElementdefinitionTypes = "RegulatedAuthorization"
	ElementdefinitionTypesRelatedperson                      ElementdefinitionTypes = "RelatedPerson"
	ElementdefinitionTypesRequestorchestration               ElementdefinitionTypes = "RequestOrchestration"
	ElementdefinitionTypesRequirements                       ElementdefinitionTypes = "Requirements"
	ElementdefinitionTypesResearchstudy                      ElementdefinitionTypes = "ResearchStudy"
	ElementdefinitionTypesResearchsubject                    ElementdefinitionTypes = "ResearchSubject"
	ElementdefinitionTypesRiskassessment                     ElementdefinitionTypes = "RiskAssessment"
	ElementdefinitionTypesSchedule                           ElementdefinitionTypes = "Schedule"
	ElementdefinitionTypesSearchparameter                    ElementdefinitionTypes = "SearchParameter"
	ElementdefinitionTypesServicerequest                     ElementdefinitionTypes = "ServiceRequest"
	ElementdefinitionTypesSlot                               ElementdefinitionTypes = "Slot"
	ElementdefinitionTypesSpecimen                           ElementdefinitionTypes = "Specimen"
	ElementdefinitionTypesSpecimendefinition                 ElementdefinitionTypes = "SpecimenDefinition"
	ElementdefinitionTypesStructuredefinition                ElementdefinitionTypes = "StructureDefinition"
	ElementdefinitionTypesStructuremap                       ElementdefinitionTypes = "StructureMap"
	ElementdefinitionTypesSubscription                       ElementdefinitionTypes = "Subscription"
	ElementdefinitionTypesSubscriptionstatus                 ElementdefinitionTypes = "SubscriptionStatus"
	ElementdefinitionTypesSubscriptiontopic                  ElementdefinitionTypes = "SubscriptionTopic"
	ElementdefinitionTypesSubstance                          ElementdefinitionTypes = "Substance"
	ElementdefinitionTypesSubstancedefinition                ElementdefinitionTypes = "SubstanceDefinition"
	ElementdefinitionTypesSubstancenucleicacid               ElementdefinitionTypes = "SubstanceNucleicAcid"
	ElementdefinitionTypesSubstancepolymer                   ElementdefinitionTypes = "SubstancePolymer"
	ElementdefinitionTypesSubstanceprotein                   ElementdefinitionTypes = "SubstanceProtein"
	ElementdefinitionTypesSubstancereferenceinformation      ElementdefinitionTypes = "SubstanceReferenceInformation"
	ElementdefinitionTypesSubstancesourcematerial            ElementdefinitionTypes = "SubstanceSourceMaterial"
	ElementdefinitionTypesSupplydelivery                     ElementdefinitionTypes = "SupplyDelivery"
	ElementdefinitionTypesSupplyrequest                      ElementdefinitionTypes = "SupplyRequest"
	ElementdefinitionTypesTask                               ElementdefinitionTypes = "Task"
	ElementdefinitionTypesTerminologycapabilities            ElementdefinitionTypes = "TerminologyCapabilities"
	ElementdefinitionTypesTestplan                           ElementdefinitionTypes = "TestPlan"
	ElementdefinitionTypesTestreport                         ElementdefinitionTypes = "TestReport"
	ElementdefinitionTypesTestscript                         ElementdefinitionTypes = "TestScript"
	ElementdefinitionTypesTransport                          ElementdefinitionTypes = "Transport"
	ElementdefinitionTypesValueset                           ElementdefinitionTypes = "ValueSet"
	ElementdefinitionTypesVerificationresult                 ElementdefinitionTypes = "VerificationResult"
	ElementdefinitionTypesVisionprescription                 ElementdefinitionTypes = "VisionPrescription"
	ElementdefinitionTypesParameters                         ElementdefinitionTypes = "Parameters"
)

func AllElementdefinitionTypes() []ElementdefinitionTypes {
	return []ElementdefinitionTypes{
		ElementdefinitionTypesStringURL,
		ElementdefinitionTypesBooleanURL,
		ElementdefinitionTypesDateURL,
		ElementdefinitionTypesDatetimeURL,
		ElementdefinitionTypesDecimalURL,
		ElementdefinitionTypesIntegerURL,
		ElementdefinitionTypesTimeURL,
		ElementdefinitionTypesBase,
		ElementdefinitionTypesElement,
		ElementdefinitionTypesBackboneelement,
		ElementdefinitionTypesDatatype,
		ElementdefinitionTypesAddress,
		ElementdefinitionTypesAnnotation,
		ElementdefinitionTypesAttachment,
		ElementdefinitionTypesAvailability,
		ElementdefinitionTypesBackbonetype,
		ElementdefinitionTypesDosage,
		ElementdefinitionTypesElementdefinition,
		ElementdefinitionTypesMarketingstatus,
		ElementdefinitionTypesProductshelflife,
		ElementdefinitionTypesTiming,
		ElementdefinitionTypesCodeableconcept,
		ElementdefinitionTypesCodeablereference,
		ElementdefinitionTypesCoding,
		ElementdefinitionTypesContactdetail,
		ElementdefinitionTypesContactpoint,
		ElementdefinitionTypesContributor,
		ElementdefinitionTypesDatarequirement,
		ElementdefinitionTypesExpression,
		ElementdefinitionTypesExtendedcontactdetail,
		ElementdefinitionTypesExtension,
		ElementdefinitionTypesHumanname,
		ElementdefinitionTypesIdentifier,
		ElementdefinitionTypesMeta,
		ElementdefinitionTypesMonetarycomponent,
		ElementdefinitionTypesMoney,
		ElementdefinitionTypesNarrative,
		ElementdefinitionTypesParameterdefinition,
		ElementdefinitionTypesPeriod,
		ElementdefinitionTypesPrimitivetype,
		ElementdefinitionTypesBase64Binary,
		ElementdefinitionTypesBoolean,
		ElementdefinitionTypesDate,
		ElementdefinitionTypesDatetime,
		ElementdefinitionTypesDecimal,
		ElementdefinitionTypesInstant,
		ElementdefinitionTypesInteger,
		ElementdefinitionTypesPositiveint,
		ElementdefinitionTypesUnsignedint,
		ElementdefinitionTypesInteger64,
		ElementdefinitionTypesString,
		ElementdefinitionTypesCode,
		ElementdefinitionTypesId,
		ElementdefinitionTypesMarkdown,
		ElementdefinitionTypesTime,
		ElementdefinitionTypesUri,
		ElementdefinitionTypesCanonical,
		ElementdefinitionTypesOid,
		ElementdefinitionTypesUrl,
		ElementdefinitionTypesUuid,
		ElementdefinitionTypesQuantity,
		ElementdefinitionTypesAge,
		ElementdefinitionTypesCount,
		ElementdefinitionTypesDistance,
		ElementdefinitionTypesDuration,
		ElementdefinitionTypesRange,
		ElementdefinitionTypesRatio,
		ElementdefinitionTypesRatiorange,
		ElementdefinitionTypesReference,
		ElementdefinitionTypesRelatedartifact,
		ElementdefinitionTypesSampleddata,
		ElementdefinitionTypesSignature,
		ElementdefinitionTypesTriggerdefinition,
		ElementdefinitionTypesUsagecontext,
		ElementdefinitionTypesVirtualservicedetail,
		ElementdefinitionTypesXhtml,
		ElementdefinitionTypesResource,
		ElementdefinitionTypesBinary,
		ElementdefinitionTypesBundle,
		ElementdefinitionTypesDomainresource,
		ElementdefinitionTypesAccount,
		ElementdefinitionTypesActivitydefinition,
		ElementdefinitionTypesActordefinition,
		ElementdefinitionTypesAdministrableproductdefinition,
		ElementdefinitionTypesAdverseevent,
		ElementdefinitionTypesAllergyintolerance,
		ElementdefinitionTypesAppointment,
		ElementdefinitionTypesAppointmentresponse,
		ElementdefinitionTypesArtifactassessment,
		ElementdefinitionTypesAuditevent,
		ElementdefinitionTypesBasic,
		ElementdefinitionTypesBiologicallyderivedproduct,
		ElementdefinitionTypesBiologicallyderivedproductdispense,
		ElementdefinitionTypesBodystructure,
		ElementdefinitionTypesCanonicalresource,
		ElementdefinitionTypesCapabilitystatement,
		ElementdefinitionTypesCareplan,
		ElementdefinitionTypesCareteam,
		ElementdefinitionTypesChargeitem,
		ElementdefinitionTypesChargeitemdefinition,
		ElementdefinitionTypesCitation,
		ElementdefinitionTypesClaim,
		ElementdefinitionTypesClaimresponse,
		ElementdefinitionTypesClinicalimpression,
		ElementdefinitionTypesClinicalusedefinition,
		ElementdefinitionTypesCodesystem,
		ElementdefinitionTypesCommunication,
		ElementdefinitionTypesCommunicationrequest,
		ElementdefinitionTypesCompartmentdefinition,
		ElementdefinitionTypesComposition,
		ElementdefinitionTypesConceptmap,
		ElementdefinitionTypesCondition,
		ElementdefinitionTypesConditiondefinition,
		ElementdefinitionTypesConsent,
		ElementdefinitionTypesContract,
		ElementdefinitionTypesCoverage,
		ElementdefinitionTypesCoverageeligibilityrequest,
		ElementdefinitionTypesCoverageeligibilityresponse,
		ElementdefinitionTypesDetectedissue,
		ElementdefinitionTypesDevice,
		ElementdefinitionTypesDeviceassociation,
		ElementdefinitionTypesDevicedefinition,
		ElementdefinitionTypesDevicedispense,
		ElementdefinitionTypesDevicemetric,
		ElementdefinitionTypesDevicerequest,
		ElementdefinitionTypesDeviceusage,
		ElementdefinitionTypesDiagnosticreport,
		ElementdefinitionTypesDocumentreference,
		ElementdefinitionTypesEncounter,
		ElementdefinitionTypesEncounterhistory,
		ElementdefinitionTypesEndpoint,
		ElementdefinitionTypesEnrollmentrequest,
		ElementdefinitionTypesEnrollmentresponse,
		ElementdefinitionTypesEpisodeofcare,
		ElementdefinitionTypesEventdefinition,
		ElementdefinitionTypesEvidence,
		ElementdefinitionTypesEvidencereport,
		ElementdefinitionTypesEvidencevariable,
		ElementdefinitionTypesExamplescenario,
		ElementdefinitionTypesExplanationofbenefit,
		ElementdefinitionTypesFamilymemberhistory,
		ElementdefinitionTypesFlag,
		ElementdefinitionTypesFormularyitem,
		ElementdefinitionTypesGenomicstudy,
		ElementdefinitionTypesGoal,
		ElementdefinitionTypesGraphdefinition,
		ElementdefinitionTypesGroup,
		ElementdefinitionTypesGuidanceresponse,
		ElementdefinitionTypesHealthcareservice,
		ElementdefinitionTypesImagingselection,
		ElementdefinitionTypesImagingstudy,
		ElementdefinitionTypesImmunization,
		ElementdefinitionTypesImmunizationevaluation,
		ElementdefinitionTypesImmunizationrecommendation,
		ElementdefinitionTypesImplementationguide,
		ElementdefinitionTypesIngredient,
		ElementdefinitionTypesInsuranceplan,
		ElementdefinitionTypesInventoryitem,
		ElementdefinitionTypesInventoryreport,
		ElementdefinitionTypesInvoice,
		ElementdefinitionTypesLibrary,
		ElementdefinitionTypesLinkage,
		ElementdefinitionTypesList,
		ElementdefinitionTypesLocation,
		ElementdefinitionTypesManufactureditemdefinition,
		ElementdefinitionTypesMeasure,
		ElementdefinitionTypesMeasurereport,
		ElementdefinitionTypesMedication,
		ElementdefinitionTypesMedicationadministration,
		ElementdefinitionTypesMedicationdispense,
		ElementdefinitionTypesMedicationknowledge,
		ElementdefinitionTypesMedicationrequest,
		ElementdefinitionTypesMedicationstatement,
		ElementdefinitionTypesMedicinalproductdefinition,
		ElementdefinitionTypesMessagedefinition,
		ElementdefinitionTypesMessageheader,
		ElementdefinitionTypesMetadataresource,
		ElementdefinitionTypesMolecularsequence,
		ElementdefinitionTypesNamingsystem,
		ElementdefinitionTypesNutritionintake,
		ElementdefinitionTypesNutritionorder,
		ElementdefinitionTypesNutritionproduct,
		ElementdefinitionTypesObservation,
		ElementdefinitionTypesObservationdefinition,
		ElementdefinitionTypesOperationdefinition,
		ElementdefinitionTypesOperationoutcome,
		ElementdefinitionTypesOrganization,
		ElementdefinitionTypesOrganizationaffiliation,
		ElementdefinitionTypesPackagedproductdefinition,
		ElementdefinitionTypesPatient,
		ElementdefinitionTypesPaymentnotice,
		ElementdefinitionTypesPaymentreconciliation,
		ElementdefinitionTypesPermission,
		ElementdefinitionTypesPerson,
		ElementdefinitionTypesPlandefinition,
		ElementdefinitionTypesPractitioner,
		ElementdefinitionTypesPractitionerrole,
		ElementdefinitionTypesProcedure,
		ElementdefinitionTypesProvenance,
		ElementdefinitionTypesQuestionnaire,
		ElementdefinitionTypesQuestionnaireresponse,
		ElementdefinitionTypesRegulatedauthorization,
		ElementdefinitionTypesRelatedperson,
		ElementdefinitionTypesRequestorchestration,
		ElementdefinitionTypesRequirements,
		ElementdefinitionTypesResearchstudy,
		ElementdefinitionTypesResearchsubject,
		ElementdefinitionTypesRiskassessment,
		ElementdefinitionTypesSchedule,
		ElementdefinitionTypesSearchparameter,
		ElementdefinitionTypesServicerequest,
		ElementdefinitionTypesSlot,
		ElementdefinitionTypesSpecimen,
		ElementdefinitionTypesSpecimendefinition,
		ElementdefinitionTypesStructuredefinition,
		ElementdefinitionTypesStructuremap,
		ElementdefinitionTypesSubscription,
		ElementdefinitionTypesSubscriptionstatus,
		ElementdefinitionTypesSubscriptiontopic,
		ElementdefinitionTypesSubstance,
		ElementdefinitionTypesSubstancedefinition,
		ElementdefinitionTypesSubstancenucleicacid,
		ElementdefinitionTypesSubstancepolymer,
		ElementdefinitionTypesSubstanceprotein,
		ElementdefinitionTypesSubstancereferenceinformation,
		ElementdefinitionTypesSubstancesourcematerial,
		ElementdefinitionTypesSupplydelivery,
		ElementdefinitionTypesSupplyrequest,
		ElementdefinitionTypesTask,
		ElementdefinitionTypesTerminologycapabilities,
		ElementdefinitionTypesTestplan,
		ElementdefinitionTypesTestreport,
		ElementdefinitionTypesTestscript,
		ElementdefinitionTypesTransport,
		ElementdefinitionTypesValueset,
		ElementdefinitionTypesVerificationresult,
		ElementdefinitionTypesVisionprescription,
		ElementdefinitionTypesParameters,
	}
}

func FindElementdefinitionTypes(filter string) []ElementdefinitionTypes {
	ret := make([]ElementdefinitionTypes, 0)
	for _, at := range AllElementdefinitionTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ElementdefinitionTypes) ToString() {
	fmt.Println(au.String())
}
func (au ElementdefinitionTypes) String() string {
	switch au {
	case ElementdefinitionTypesStringURL:
		return "String"
	case ElementdefinitionTypesBooleanURL:
		return "Boolean"
	case ElementdefinitionTypesDateURL:
		return "Date"
	case ElementdefinitionTypesDatetimeURL:
		return "DateTime"
	case ElementdefinitionTypesDecimalURL:
		return "Decimal"
	case ElementdefinitionTypesIntegerURL:
		return "Integer"
	case ElementdefinitionTypesTimeURL:
		return "Time"
	case ElementdefinitionTypesBase:
		return "Base"
	case ElementdefinitionTypesElement:
		return "Element"
	case ElementdefinitionTypesBackboneelement:
		return "BackboneElement"
	case ElementdefinitionTypesDatatype:
		return "DataType"
	case ElementdefinitionTypesAddress:
		return "Address"
	case ElementdefinitionTypesAnnotation:
		return "Annotation"
	case ElementdefinitionTypesAttachment:
		return "Attachment"
	case ElementdefinitionTypesAvailability:
		return "Availability"
	case ElementdefinitionTypesBackbonetype:
		return "BackboneType"
	case ElementdefinitionTypesDosage:
		return "Dosage"
	case ElementdefinitionTypesElementdefinition:
		return "ElementDefinition"
	case ElementdefinitionTypesMarketingstatus:
		return "MarketingStatus"
	case ElementdefinitionTypesProductshelflife:
		return "ProductShelfLife"
	case ElementdefinitionTypesTiming:
		return "Timing"
	case ElementdefinitionTypesCodeableconcept:
		return "CodeableConcept"
	case ElementdefinitionTypesCodeablereference:
		return "CodeableReference"
	case ElementdefinitionTypesCoding:
		return "Coding"
	case ElementdefinitionTypesContactdetail:
		return "ContactDetail"
	case ElementdefinitionTypesContactpoint:
		return "ContactPoint"
	case ElementdefinitionTypesContributor:
		return "Contributor"
	case ElementdefinitionTypesDatarequirement:
		return "DataRequirement"
	case ElementdefinitionTypesExpression:
		return "Expression"
	case ElementdefinitionTypesExtendedcontactdetail:
		return "ExtendedContactDetail"
	case ElementdefinitionTypesExtension:
		return "Extension"
	case ElementdefinitionTypesHumanname:
		return "HumanName"
	case ElementdefinitionTypesIdentifier:
		return "Identifier"
	case ElementdefinitionTypesMeta:
		return "Meta"
	case ElementdefinitionTypesMonetarycomponent:
		return "MonetaryComponent"
	case ElementdefinitionTypesMoney:
		return "Money"
	case ElementdefinitionTypesNarrative:
		return "Narrative"
	case ElementdefinitionTypesParameterdefinition:
		return "ParameterDefinition"
	case ElementdefinitionTypesPeriod:
		return "Period"
	case ElementdefinitionTypesPrimitivetype:
		return "PrimitiveType"
	case ElementdefinitionTypesBase64Binary:
		return "base64Binary"
	case ElementdefinitionTypesBoolean:
		return "boolean"
	case ElementdefinitionTypesDate:
		return "date"
	case ElementdefinitionTypesDatetime:
		return "dateTime"
	case ElementdefinitionTypesDecimal:
		return "decimal"
	case ElementdefinitionTypesInstant:
		return "instant"
	case ElementdefinitionTypesInteger:
		return "integer"
	case ElementdefinitionTypesPositiveint:
		return "positiveInt"
	case ElementdefinitionTypesUnsignedint:
		return "unsignedInt"
	case ElementdefinitionTypesInteger64:
		return "integer64"
	case ElementdefinitionTypesString:
		return "string"
	case ElementdefinitionTypesCode:
		return "code"
	case ElementdefinitionTypesId:
		return "id"
	case ElementdefinitionTypesMarkdown:
		return "markdown"
	case ElementdefinitionTypesTime:
		return "time"
	case ElementdefinitionTypesUri:
		return "uri"
	case ElementdefinitionTypesCanonical:
		return "canonical"
	case ElementdefinitionTypesOid:
		return "oid"
	case ElementdefinitionTypesUrl:
		return "url"
	case ElementdefinitionTypesUuid:
		return "uuid"
	case ElementdefinitionTypesQuantity:
		return "Quantity"
	case ElementdefinitionTypesAge:
		return "Age"
	case ElementdefinitionTypesCount:
		return "Count"
	case ElementdefinitionTypesDistance:
		return "Distance"
	case ElementdefinitionTypesDuration:
		return "Duration"
	case ElementdefinitionTypesRange:
		return "Range"
	case ElementdefinitionTypesRatio:
		return "Ratio"
	case ElementdefinitionTypesRatiorange:
		return "RatioRange"
	case ElementdefinitionTypesReference:
		return "Reference"
	case ElementdefinitionTypesRelatedartifact:
		return "RelatedArtifact"
	case ElementdefinitionTypesSampleddata:
		return "SampledData"
	case ElementdefinitionTypesSignature:
		return "Signature"
	case ElementdefinitionTypesTriggerdefinition:
		return "TriggerDefinition"
	case ElementdefinitionTypesUsagecontext:
		return "UsageContext"
	case ElementdefinitionTypesVirtualservicedetail:
		return "VirtualServiceDetail"
	case ElementdefinitionTypesXhtml:
		return "xhtml"
	case ElementdefinitionTypesResource:
		return "Resource"
	case ElementdefinitionTypesBinary:
		return "Binary"
	case ElementdefinitionTypesBundle:
		return "Bundle"
	case ElementdefinitionTypesDomainresource:
		return "DomainResource"
	case ElementdefinitionTypesAccount:
		return "Account"
	case ElementdefinitionTypesActivitydefinition:
		return "ActivityDefinition"
	case ElementdefinitionTypesActordefinition:
		return "ActorDefinition"
	case ElementdefinitionTypesAdministrableproductdefinition:
		return "AdministrableProductDefinition"
	case ElementdefinitionTypesAdverseevent:
		return "AdverseEvent"
	case ElementdefinitionTypesAllergyintolerance:
		return "AllergyIntolerance"
	case ElementdefinitionTypesAppointment:
		return "Appointment"
	case ElementdefinitionTypesAppointmentresponse:
		return "AppointmentResponse"
	case ElementdefinitionTypesArtifactassessment:
		return "ArtifactAssessment"
	case ElementdefinitionTypesAuditevent:
		return "AuditEvent"
	case ElementdefinitionTypesBasic:
		return "Basic"
	case ElementdefinitionTypesBiologicallyderivedproduct:
		return "BiologicallyDerivedProduct"
	case ElementdefinitionTypesBiologicallyderivedproductdispense:
		return "BiologicallyDerivedProductDispense"
	case ElementdefinitionTypesBodystructure:
		return "BodyStructure"
	case ElementdefinitionTypesCanonicalresource:
		return "CanonicalResource"
	case ElementdefinitionTypesCapabilitystatement:
		return "CapabilityStatement"
	case ElementdefinitionTypesCareplan:
		return "CarePlan"
	case ElementdefinitionTypesCareteam:
		return "CareTeam"
	case ElementdefinitionTypesChargeitem:
		return "ChargeItem"
	case ElementdefinitionTypesChargeitemdefinition:
		return "ChargeItemDefinition"
	case ElementdefinitionTypesCitation:
		return "Citation"
	case ElementdefinitionTypesClaim:
		return "Claim"
	case ElementdefinitionTypesClaimresponse:
		return "ClaimResponse"
	case ElementdefinitionTypesClinicalimpression:
		return "ClinicalImpression"
	case ElementdefinitionTypesClinicalusedefinition:
		return "ClinicalUseDefinition"
	case ElementdefinitionTypesCodesystem:
		return "CodeSystem"
	case ElementdefinitionTypesCommunication:
		return "Communication"
	case ElementdefinitionTypesCommunicationrequest:
		return "CommunicationRequest"
	case ElementdefinitionTypesCompartmentdefinition:
		return "CompartmentDefinition"
	case ElementdefinitionTypesComposition:
		return "Composition"
	case ElementdefinitionTypesConceptmap:
		return "ConceptMap"
	case ElementdefinitionTypesCondition:
		return "Condition"
	case ElementdefinitionTypesConditiondefinition:
		return "ConditionDefinition"
	case ElementdefinitionTypesConsent:
		return "Consent"
	case ElementdefinitionTypesContract:
		return "Contract"
	case ElementdefinitionTypesCoverage:
		return "Coverage"
	case ElementdefinitionTypesCoverageeligibilityrequest:
		return "CoverageEligibilityRequest"
	case ElementdefinitionTypesCoverageeligibilityresponse:
		return "CoverageEligibilityResponse"
	case ElementdefinitionTypesDetectedissue:
		return "DetectedIssue"
	case ElementdefinitionTypesDevice:
		return "Device"
	case ElementdefinitionTypesDeviceassociation:
		return "DeviceAssociation"
	case ElementdefinitionTypesDevicedefinition:
		return "DeviceDefinition"
	case ElementdefinitionTypesDevicedispense:
		return "DeviceDispense"
	case ElementdefinitionTypesDevicemetric:
		return "DeviceMetric"
	case ElementdefinitionTypesDevicerequest:
		return "DeviceRequest"
	case ElementdefinitionTypesDeviceusage:
		return "DeviceUsage"
	case ElementdefinitionTypesDiagnosticreport:
		return "DiagnosticReport"
	case ElementdefinitionTypesDocumentreference:
		return "DocumentReference"
	case ElementdefinitionTypesEncounter:
		return "Encounter"
	case ElementdefinitionTypesEncounterhistory:
		return "EncounterHistory"
	case ElementdefinitionTypesEndpoint:
		return "Endpoint"
	case ElementdefinitionTypesEnrollmentrequest:
		return "EnrollmentRequest"
	case ElementdefinitionTypesEnrollmentresponse:
		return "EnrollmentResponse"
	case ElementdefinitionTypesEpisodeofcare:
		return "EpisodeOfCare"
	case ElementdefinitionTypesEventdefinition:
		return "EventDefinition"
	case ElementdefinitionTypesEvidence:
		return "Evidence"
	case ElementdefinitionTypesEvidencereport:
		return "EvidenceReport"
	case ElementdefinitionTypesEvidencevariable:
		return "EvidenceVariable"
	case ElementdefinitionTypesExamplescenario:
		return "ExampleScenario"
	case ElementdefinitionTypesExplanationofbenefit:
		return "ExplanationOfBenefit"
	case ElementdefinitionTypesFamilymemberhistory:
		return "FamilyMemberHistory"
	case ElementdefinitionTypesFlag:
		return "Flag"
	case ElementdefinitionTypesFormularyitem:
		return "FormularyItem"
	case ElementdefinitionTypesGenomicstudy:
		return "GenomicStudy"
	case ElementdefinitionTypesGoal:
		return "Goal"
	case ElementdefinitionTypesGraphdefinition:
		return "GraphDefinition"
	case ElementdefinitionTypesGroup:
		return "Group"
	case ElementdefinitionTypesGuidanceresponse:
		return "GuidanceResponse"
	case ElementdefinitionTypesHealthcareservice:
		return "HealthcareService"
	case ElementdefinitionTypesImagingselection:
		return "ImagingSelection"
	case ElementdefinitionTypesImagingstudy:
		return "ImagingStudy"
	case ElementdefinitionTypesImmunization:
		return "Immunization"
	case ElementdefinitionTypesImmunizationevaluation:
		return "ImmunizationEvaluation"
	case ElementdefinitionTypesImmunizationrecommendation:
		return "ImmunizationRecommendation"
	case ElementdefinitionTypesImplementationguide:
		return "ImplementationGuide"
	case ElementdefinitionTypesIngredient:
		return "Ingredient"
	case ElementdefinitionTypesInsuranceplan:
		return "InsurancePlan"
	case ElementdefinitionTypesInventoryitem:
		return "InventoryItem"
	case ElementdefinitionTypesInventoryreport:
		return "InventoryReport"
	case ElementdefinitionTypesInvoice:
		return "Invoice"
	case ElementdefinitionTypesLibrary:
		return "Library"
	case ElementdefinitionTypesLinkage:
		return "Linkage"
	case ElementdefinitionTypesList:
		return "List"
	case ElementdefinitionTypesLocation:
		return "Location"
	case ElementdefinitionTypesManufactureditemdefinition:
		return "ManufacturedItemDefinition"
	case ElementdefinitionTypesMeasure:
		return "Measure"
	case ElementdefinitionTypesMeasurereport:
		return "MeasureReport"
	case ElementdefinitionTypesMedication:
		return "Medication"
	case ElementdefinitionTypesMedicationadministration:
		return "MedicationAdministration"
	case ElementdefinitionTypesMedicationdispense:
		return "MedicationDispense"
	case ElementdefinitionTypesMedicationknowledge:
		return "MedicationKnowledge"
	case ElementdefinitionTypesMedicationrequest:
		return "MedicationRequest"
	case ElementdefinitionTypesMedicationstatement:
		return "MedicationStatement"
	case ElementdefinitionTypesMedicinalproductdefinition:
		return "MedicinalProductDefinition"
	case ElementdefinitionTypesMessagedefinition:
		return "MessageDefinition"
	case ElementdefinitionTypesMessageheader:
		return "MessageHeader"
	case ElementdefinitionTypesMetadataresource:
		return "MetadataResource"
	case ElementdefinitionTypesMolecularsequence:
		return "MolecularSequence"
	case ElementdefinitionTypesNamingsystem:
		return "NamingSystem"
	case ElementdefinitionTypesNutritionintake:
		return "NutritionIntake"
	case ElementdefinitionTypesNutritionorder:
		return "NutritionOrder"
	case ElementdefinitionTypesNutritionproduct:
		return "NutritionProduct"
	case ElementdefinitionTypesObservation:
		return "Observation"
	case ElementdefinitionTypesObservationdefinition:
		return "ObservationDefinition"
	case ElementdefinitionTypesOperationdefinition:
		return "OperationDefinition"
	case ElementdefinitionTypesOperationoutcome:
		return "OperationOutcome"
	case ElementdefinitionTypesOrganization:
		return "Organization"
	case ElementdefinitionTypesOrganizationaffiliation:
		return "OrganizationAffiliation"
	case ElementdefinitionTypesPackagedproductdefinition:
		return "PackagedProductDefinition"
	case ElementdefinitionTypesPatient:
		return "Patient"
	case ElementdefinitionTypesPaymentnotice:
		return "PaymentNotice"
	case ElementdefinitionTypesPaymentreconciliation:
		return "PaymentReconciliation"
	case ElementdefinitionTypesPermission:
		return "Permission"
	case ElementdefinitionTypesPerson:
		return "Person"
	case ElementdefinitionTypesPlandefinition:
		return "PlanDefinition"
	case ElementdefinitionTypesPractitioner:
		return "Practitioner"
	case ElementdefinitionTypesPractitionerrole:
		return "PractitionerRole"
	case ElementdefinitionTypesProcedure:
		return "Procedure"
	case ElementdefinitionTypesProvenance:
		return "Provenance"
	case ElementdefinitionTypesQuestionnaire:
		return "Questionnaire"
	case ElementdefinitionTypesQuestionnaireresponse:
		return "QuestionnaireResponse"
	case ElementdefinitionTypesRegulatedauthorization:
		return "RegulatedAuthorization"
	case ElementdefinitionTypesRelatedperson:
		return "RelatedPerson"
	case ElementdefinitionTypesRequestorchestration:
		return "RequestOrchestration"
	case ElementdefinitionTypesRequirements:
		return "Requirements"
	case ElementdefinitionTypesResearchstudy:
		return "ResearchStudy"
	case ElementdefinitionTypesResearchsubject:
		return "ResearchSubject"
	case ElementdefinitionTypesRiskassessment:
		return "RiskAssessment"
	case ElementdefinitionTypesSchedule:
		return "Schedule"
	case ElementdefinitionTypesSearchparameter:
		return "SearchParameter"
	case ElementdefinitionTypesServicerequest:
		return "ServiceRequest"
	case ElementdefinitionTypesSlot:
		return "Slot"
	case ElementdefinitionTypesSpecimen:
		return "Specimen"
	case ElementdefinitionTypesSpecimendefinition:
		return "SpecimenDefinition"
	case ElementdefinitionTypesStructuredefinition:
		return "StructureDefinition"
	case ElementdefinitionTypesStructuremap:
		return "StructureMap"
	case ElementdefinitionTypesSubscription:
		return "Subscription"
	case ElementdefinitionTypesSubscriptionstatus:
		return "SubscriptionStatus"
	case ElementdefinitionTypesSubscriptiontopic:
		return "SubscriptionTopic"
	case ElementdefinitionTypesSubstance:
		return "Substance"
	case ElementdefinitionTypesSubstancedefinition:
		return "SubstanceDefinition"
	case ElementdefinitionTypesSubstancenucleicacid:
		return "SubstanceNucleicAcid"
	case ElementdefinitionTypesSubstancepolymer:
		return "SubstancePolymer"
	case ElementdefinitionTypesSubstanceprotein:
		return "SubstanceProtein"
	case ElementdefinitionTypesSubstancereferenceinformation:
		return "SubstanceReferenceInformation"
	case ElementdefinitionTypesSubstancesourcematerial:
		return "SubstanceSourceMaterial"
	case ElementdefinitionTypesSupplydelivery:
		return "SupplyDelivery"
	case ElementdefinitionTypesSupplyrequest:
		return "SupplyRequest"
	case ElementdefinitionTypesTask:
		return "Task"
	case ElementdefinitionTypesTerminologycapabilities:
		return "TerminologyCapabilities"
	case ElementdefinitionTypesTestplan:
		return "TestPlan"
	case ElementdefinitionTypesTestreport:
		return "TestReport"
	case ElementdefinitionTypesTestscript:
		return "TestScript"
	case ElementdefinitionTypesTransport:
		return "Transport"
	case ElementdefinitionTypesValueset:
		return "ValueSet"
	case ElementdefinitionTypesVerificationresult:
		return "VerificationResult"
	case ElementdefinitionTypesVisionprescription:
		return "VisionPrescription"
	case ElementdefinitionTypesParameters:
		return "Parameters"

	default:
		return "Unknown Constraint Severity"
	}
}
