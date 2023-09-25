package code

import (
	"fmt"
	"strings"
)

// https://www.hl7.org/fhir/valueset-c80-practice-codes.html

type PractitionerSpeciality int

const (
	PractitionerSpecialityAdultMentalIllness PractitionerSpeciality = iota
	PractitionerSpecialityAnesthetics
	PractitionerSpecialityAudiologicalMedicine
	PractitionerSpecialityBloodBankingAndTransfusionMedicine
	PractitionerSpecialityBurnsCare
	PractitionerSpecialityCardiology
	PractitionerSpecialityClinicalCytogeneticsAndMolecularGenetics
	PractitionerSpecialityClinicalGenetics
	PractitionerSpecialityClinicalHematology
	PractitionerSpecialityClinicalImmunology
	PractitionerSpecialityClinicalMicrobiology
	PractitionerSpecialityClinicalNeuroPhysiology
	PractitionerSpecialityClinicalOncology
	PractitionerSpecialityClinicalPharmacology
	PractitionerSpecialityClinicalPhysiology
	PractitionerSpecialityCommunityMedicine
	PractitionerSpecialityCriticalCareMedicine
	PractitionerSpecialityDentalMedicineSpecialties
	PractitionerSpecialityDentalGeneralDentalPractice
	PractitionerSpecialityDermatology
	PractitionerSpecialityDiabeticMedicine
	PractitionerSpecialityDiveMedicine
	PractitionerSpecialityEndocrinology
	PractitionerSpecialityFamilyPractice
	PractitionerSpecialityGastroenterology
	PractitionerSpecialityGeneralMedicalPractice
	PractitionerSpecialityGeneralMedicine
	PractitionerSpecialityGeneralPathology
	PractitionerSpecialityGeneralPractice
	PractitionerSpecialityGenitoUrinaryMedicine
	PractitionerSpecialityGeriatricMedicine
	PractitionerSpecialityGynecologicalOncology
	PractitionerSpecialityGynecology
	PractitionerSpecialityHematopathology
	PractitionerSpecialityHepatology
	PractitionerSpecialityHistopathology
	PractitionerSpecialityImmunopathology
	PractitionerSpecialityInfectiousDiseases
	PractitionerSpecialityInternalMedicine
	PractitionerSpecialityLearningDisability
	PractitionerSpecialityMedicalOncology
	PractitionerSpecialityMedicalOphthalmology
	PractitionerSpecialityMilitaryMedicine
	PractitionerSpecialityNephrology
	PractitionerSpecialityNeurology
	PractitionerSpecialityNeuropathology
	PractitionerSpecialityNuclearMedicine
	PractitionerSpecialityObstetrics
	PractitionerSpecialityObstetricsAndGynecology
	PractitionerSpecialityOccupationalMedicine
	PractitionerSpecialityOphthalmicSurgery
	PractitionerSpecialityOphthalmology
	PractitionerSpecialityOsteopathicManipulativeMedicine
	PractitionerSpecialityOtolaryngology
	PractitionerSpecialityPainManagement
	PractitionerSpecialityPalliativeMedicine
	PractitionerSpecialityPediatricChildAndAdolescentPsychiatry
	PractitionerSpecialityPediatricCardiology
	PractitionerSpecialityPediatricDentistry
	PractitionerSpecialityPediatricEndocrinology
	PractitionerSpecialityPediatricGastroenterology
	PractitionerSpecialityPediatricGenetics
	PractitionerSpecialityPediatricHematology
	PractitionerSpecialityPediatricImmunology
	PractitionerSpecialityPediatricInfectiousDiseases
	PractitionerSpecialityPediatricNephrology
	PractitionerSpecialityPediatricOncology
	PractitionerSpecialityPediatricOphthalmology
	PractitionerSpecialityPediatricPulmonology
	PractitionerSpecialityPediatricRheumatology
	PractitionerSpecialityPediatricSurgery
	PractitionerSpecialityPediatricSurgeryBoneMarrowTransplantation
	PractitionerSpecialityPreventiveMedicine
	PractitionerSpecialityPsychiatry
	PractitionerSpecialityPsychotherapy
	PractitionerSpecialityPublicHealthMedicine
	PractitionerSpecialityPulmonaryMedicine
	PractitionerSpecialityRadiationOncology
	PractitionerSpecialityRadiology
	PractitionerSpecialityRadiologyInterventionalRadiology
	PractitionerSpecialityRehabilitation
	PractitionerSpecialityRespiteCare
	PractitionerSpecialityRheumatology
	PractitionerSpecialitySleepStudies
	PractitionerSpecialitySurgeryBoneAndMarrowTransplantation
	PractitionerSpecialitySurgeryBreastSurgery
	PractitionerSpecialitySurgeryCardiacSurgery
	PractitionerSpecialitySurgeryCardiothoracicTransplantation
	PractitionerSpecialitySurgeryColorectalSurgery
	PractitionerSpecialitySurgeryDentalEndodontics
	PractitionerSpecialitySurgeryDentalOralAndMaxillofacialSurgery
	PractitionerSpecialitySurgeryDentalOralSurgery
	PractitionerSpecialitySurgeryDentalOrthodontics
	PractitionerSpecialitySurgeryDentalPeriodontalSurgery
	PractitionerSpecialitySurgeryDentalProstheticDentistryProsthodontics
	PractitionerSpecialitySurgeryDentalSurgicalProsthodontics
	PractitionerSpecialitySurgeryDentistryRestorativeDentistry
	PractitionerSpecialitySurgeryDentistrySurgical
	PractitionerSpecialitySurgeryDentistrySurgicalOrthodontics
	PractitionerSpecialitySurgeryDermatologicSurgery
	PractitionerSpecialitySurgeryEarNoseAndThroatSurgery
	PractitionerSpecialitySurgeryGeneral
	PractitionerSpecialitySurgeryHepatobiliaryAndPancreaticSurgery
	PractitionerSpecialitySurgeryNeurosurgery
	PractitionerSpecialitySurgeryPlasticSurgery
	PractitionerSpecialitySurgeryTransplantationSurgery
	PractitionerSpecialitySurgeryTraumaAndOrthopedics
	PractitionerSpecialitySurgeryVascular
	PractitionerSpecialitySurgicalOncology
	PractitionerSpecialitySurgicalAccidentEmergency
	PractitionerSpecialityThoracicMedicine
	PractitionerSpecialityToxicology
	PractitionerSpecialityTropicalMedicine
	PractitionerSpecialityUrologicalOncology
	PractitionerSpecialityUrology
	PractitionerSpecialityOtherMedicalSpecialty
	PractitionerSpecialityOtherSurgicalSpecialty
)

func AllPractitionerSpeciality() []PractitionerSpeciality {
	return []PractitionerSpeciality{
		PractitionerSpecialityAdultMentalIllness,
		PractitionerSpecialityAnesthetics,
		PractitionerSpecialityAudiologicalMedicine,
		PractitionerSpecialityBloodBankingAndTransfusionMedicine,
		PractitionerSpecialityBurnsCare,
		PractitionerSpecialityCardiology,
		PractitionerSpecialityClinicalCytogeneticsAndMolecularGenetics,
		PractitionerSpecialityClinicalGenetics,
		PractitionerSpecialityClinicalHematology,
		PractitionerSpecialityClinicalImmunology,
		PractitionerSpecialityClinicalMicrobiology,
		PractitionerSpecialityClinicalNeuroPhysiology,
		PractitionerSpecialityClinicalOncology,
		PractitionerSpecialityClinicalPharmacology,
		PractitionerSpecialityClinicalPhysiology,
		PractitionerSpecialityCommunityMedicine,
		PractitionerSpecialityCriticalCareMedicine,
		PractitionerSpecialityDentalMedicineSpecialties,
		PractitionerSpecialityDentalGeneralDentalPractice,
		PractitionerSpecialityDermatology,
		PractitionerSpecialityDiabeticMedicine,
		PractitionerSpecialityDiveMedicine,
		PractitionerSpecialityEndocrinology,
		PractitionerSpecialityFamilyPractice,
		PractitionerSpecialityGastroenterology,
		PractitionerSpecialityGeneralMedicalPractice,
		PractitionerSpecialityGeneralMedicine,
		PractitionerSpecialityGeneralPathology,
		PractitionerSpecialityGeneralPractice,
		PractitionerSpecialityGenitoUrinaryMedicine,
		PractitionerSpecialityGeriatricMedicine,
		PractitionerSpecialityGynecologicalOncology,
		PractitionerSpecialityGynecology,
		PractitionerSpecialityHematopathology,
		PractitionerSpecialityHepatology,
		PractitionerSpecialityHistopathology,
		PractitionerSpecialityImmunopathology,
		PractitionerSpecialityInfectiousDiseases,
		PractitionerSpecialityInternalMedicine,
		PractitionerSpecialityLearningDisability,
		PractitionerSpecialityMedicalOncology,
		PractitionerSpecialityMedicalOphthalmology,
		PractitionerSpecialityMilitaryMedicine,
		PractitionerSpecialityNephrology,
		PractitionerSpecialityNeurology,
		PractitionerSpecialityNeuropathology,
		PractitionerSpecialityNuclearMedicine,
		PractitionerSpecialityObstetrics,
		PractitionerSpecialityObstetricsAndGynecology,
		PractitionerSpecialityOccupationalMedicine,
		PractitionerSpecialityOphthalmicSurgery,
		PractitionerSpecialityOphthalmology,
		PractitionerSpecialityOsteopathicManipulativeMedicine,
		PractitionerSpecialityOtolaryngology,
		PractitionerSpecialityPainManagement,
		PractitionerSpecialityPalliativeMedicine,
		PractitionerSpecialityPediatricChildAndAdolescentPsychiatry,
		PractitionerSpecialityPediatricCardiology,
		PractitionerSpecialityPediatricDentistry,
		PractitionerSpecialityPediatricEndocrinology,
		PractitionerSpecialityPediatricGastroenterology,
		PractitionerSpecialityPediatricGenetics,
		PractitionerSpecialityPediatricHematology,
		PractitionerSpecialityPediatricImmunology,
		PractitionerSpecialityPediatricInfectiousDiseases,
		PractitionerSpecialityPediatricNephrology,
		PractitionerSpecialityPediatricOncology,
		PractitionerSpecialityPediatricOphthalmology,
		PractitionerSpecialityPediatricPulmonology,
		PractitionerSpecialityPediatricRheumatology,
		PractitionerSpecialityPediatricSurgery,
		PractitionerSpecialityPediatricSurgeryBoneMarrowTransplantation,
		PractitionerSpecialityPreventiveMedicine,
		PractitionerSpecialityPsychiatry,
		PractitionerSpecialityPsychotherapy,
		PractitionerSpecialityPublicHealthMedicine,
		PractitionerSpecialityPulmonaryMedicine,
		PractitionerSpecialityRadiationOncology,
		PractitionerSpecialityRadiology,
		PractitionerSpecialityRadiologyInterventionalRadiology,
		PractitionerSpecialityRehabilitation,
		PractitionerSpecialityRespiteCare,
		PractitionerSpecialityRheumatology,
		PractitionerSpecialitySleepStudies,
		PractitionerSpecialitySurgeryBoneAndMarrowTransplantation,
		PractitionerSpecialitySurgeryBreastSurgery,
		PractitionerSpecialitySurgeryCardiacSurgery,
		PractitionerSpecialitySurgeryCardiothoracicTransplantation,
		PractitionerSpecialitySurgeryColorectalSurgery,
		PractitionerSpecialitySurgeryDentalEndodontics,
		PractitionerSpecialitySurgeryDentalOralAndMaxillofacialSurgery,
		PractitionerSpecialitySurgeryDentalOralSurgery,
		PractitionerSpecialitySurgeryDentalOrthodontics,
		PractitionerSpecialitySurgeryDentalPeriodontalSurgery,
		PractitionerSpecialitySurgeryDentalProstheticDentistryProsthodontics,
		PractitionerSpecialitySurgeryDentalSurgicalProsthodontics,
		PractitionerSpecialitySurgeryDentistryRestorativeDentistry,
		PractitionerSpecialitySurgeryDentistrySurgical,
		PractitionerSpecialitySurgeryDentistrySurgicalOrthodontics,
		PractitionerSpecialitySurgeryDermatologicSurgery,
		PractitionerSpecialitySurgeryEarNoseAndThroatSurgery,
		PractitionerSpecialitySurgeryGeneral,
		PractitionerSpecialitySurgeryHepatobiliaryAndPancreaticSurgery,
		PractitionerSpecialitySurgeryNeurosurgery,
		PractitionerSpecialitySurgeryPlasticSurgery,
		PractitionerSpecialitySurgeryTransplantationSurgery,
		PractitionerSpecialitySurgeryTraumaAndOrthopedics,
		PractitionerSpecialitySurgeryVascular,
		PractitionerSpecialitySurgicalOncology,
		PractitionerSpecialitySurgicalAccidentEmergency,
		PractitionerSpecialityThoracicMedicine,
		PractitionerSpecialityToxicology,
		PractitionerSpecialityTropicalMedicine,
		PractitionerSpecialityUrologicalOncology,
		PractitionerSpecialityUrology,
		PractitionerSpecialityOtherMedicalSpecialty,
		PractitionerSpecialityOtherSurgicalSpecialty,
	}
}

func FindPractitionerSpeciality(filter string) []PractitionerSpeciality {
	ret := make([]PractitionerSpeciality, 0)
	for _, at := range AllPractitionerSpeciality() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt PractitionerSpeciality) ToString() {
	fmt.Println(cpt.String())
}

func (cpt PractitionerSpeciality) String() string {
	switch cpt {
	case PractitionerSpecialityAdultMentalIllness:
		return "Adult mental illness"
	case PractitionerSpecialityAnesthetics:
		return "Anesthetics"
	case PractitionerSpecialityAudiologicalMedicine:
		return "Audiological medicine"
	case PractitionerSpecialityBloodBankingAndTransfusionMedicine:
		return "Blood banking and transfusion medicine"
	case PractitionerSpecialityBurnsCare:
		return "Burns care"
	case PractitionerSpecialityCardiology:
		return "Cardiology"
	case PractitionerSpecialityClinicalCytogeneticsAndMolecularGenetics:
		return "Clinical cytogenetics and molecular genetics"
	case PractitionerSpecialityClinicalGenetics:
		return "Clinical genetics"
	case PractitionerSpecialityClinicalHematology:
		return "Clinical hematology"
	case PractitionerSpecialityClinicalImmunology:
		return "Clinical immunology"
	case PractitionerSpecialityClinicalMicrobiology:
		return "Clinical microbiology"
	case PractitionerSpecialityClinicalNeuroPhysiology:
		return "Clinical neuro-physiology"
	case PractitionerSpecialityClinicalOncology:
		return "Clinical oncology"
	case PractitionerSpecialityClinicalPharmacology:
		return "Clinical pharmacology"
	case PractitionerSpecialityClinicalPhysiology:
		return "Clinical physiology"
	case PractitionerSpecialityCommunityMedicine:
		return "Community medicine"
	case PractitionerSpecialityCriticalCareMedicine:
		return "Critical care medicine"
	case PractitionerSpecialityDentalMedicineSpecialties:
		return "Dental medicine specialties"
	case PractitionerSpecialityDentalGeneralDentalPractice:
		return "Dental-General dental practice"
	case PractitionerSpecialityDermatology:
		return "Dermatology"
	case PractitionerSpecialityDiabeticMedicine:
		return "Diabetic medicine"
	case PractitionerSpecialityDiveMedicine:
		return "Dive medicine"
	case PractitionerSpecialityEndocrinology:
		return "Endocrinology"
	case PractitionerSpecialityFamilyPractice:
		return "Family practice"
	case PractitionerSpecialityGastroenterology:
		return "Gastroenterology"
	case PractitionerSpecialityGeneralMedicalPractice:
		return "General medical practice"
	case PractitionerSpecialityGeneralMedicine:
		return "General medicine"
	case PractitionerSpecialityGeneralPathology:
		return "General pathology"
	case PractitionerSpecialityGeneralPractice:
		return "General practice"
	case PractitionerSpecialityGenitoUrinaryMedicine:
		return "Genito-urinary medicine"
	case PractitionerSpecialityGeriatricMedicine:
		return "Geriatric medicine"
	case PractitionerSpecialityGynecologicalOncology:
		return "Gynecological oncology"
	case PractitionerSpecialityGynecology:
		return "Gynecology"
	case PractitionerSpecialityHematopathology:
		return "Hematopathology"
	case PractitionerSpecialityHepatology:
		return "Hepatology"
	case PractitionerSpecialityHistopathology:
		return "Histopathology"
	case PractitionerSpecialityImmunopathology:
		return "Immunopathology"
	case PractitionerSpecialityInfectiousDiseases:
		return "Infectious diseases"
	case PractitionerSpecialityInternalMedicine:
		return "Internal medicine"
	case PractitionerSpecialityLearningDisability:
		return "Learning disability"
	case PractitionerSpecialityMedicalOncology:
		return "Medical oncology"
	case PractitionerSpecialityMedicalOphthalmology:
		return "Medical ophthalmology"
	case PractitionerSpecialityMilitaryMedicine:
		return "Military medicine"
	case PractitionerSpecialityNephrology:
		return "Nephrology"
	case PractitionerSpecialityNeurology:
		return "Neurology"
	case PractitionerSpecialityNeuropathology:
		return "Neuropathology"
	case PractitionerSpecialityNuclearMedicine:
		return "Nuclear medicine"
	case PractitionerSpecialityObstetrics:
		return "Obstetrics"
	case PractitionerSpecialityObstetricsAndGynecology:
		return "Obstetrics and gynecology"
	case PractitionerSpecialityOccupationalMedicine:
		return "Occupational medicine"
	case PractitionerSpecialityOphthalmicSurgery:
		return "Ophthalmic surgery"
	case PractitionerSpecialityOphthalmology:
		return "Ophthalmology"
	case PractitionerSpecialityOsteopathicManipulativeMedicine:
		return "Osteopathic manipulative medicine"
	case PractitionerSpecialityOtolaryngology:
		return "Otolaryngology"
	case PractitionerSpecialityPainManagement:
		return "Pain management"
	case PractitionerSpecialityPalliativeMedicine:
		return "Palliative medicine"
	case PractitionerSpecialityPediatricChildAndAdolescentPsychiatry:
		return "Pediatric (Child and adolescent) psychiatry"
	case PractitionerSpecialityPediatricCardiology:
		return "Pediatric cardiology"
	case PractitionerSpecialityPediatricDentistry:
		return "Pediatric dentistry"
	case PractitionerSpecialityPediatricEndocrinology:
		return "Pediatric endocrinology"
	case PractitionerSpecialityPediatricGastroenterology:
		return "Pediatric gastroenterology"
	case PractitionerSpecialityPediatricGenetics:
		return "Pediatric genetics"
	case PractitionerSpecialityPediatricHematology:
		return "Pediatric hematology"
	case PractitionerSpecialityPediatricImmunology:
		return "Pediatric immunology"
	case PractitionerSpecialityPediatricInfectiousDiseases:
		return "Pediatric infectious diseases"
	case PractitionerSpecialityPediatricNephrology:
		return "Pediatric nephrology"
	case PractitionerSpecialityPediatricOncology:
		return "Pediatric oncology"
	case PractitionerSpecialityPediatricOphthalmology:
		return "Pediatric ophthalmology"
	case PractitionerSpecialityPediatricPulmonology:
		return "Pediatric pulmonology"
	case PractitionerSpecialityPediatricRheumatology:
		return "Pediatric rheumatology"
	case PractitionerSpecialityPediatricSurgery:
		return "Pediatric surgery"
	case PractitionerSpecialityPediatricSurgeryBoneMarrowTransplantation:
		return "Pediatric surgery-bone marrow transplantation"
	case PractitionerSpecialityPreventiveMedicine:
		return "Preventive medicine"
	case PractitionerSpecialityPsychiatry:
		return "Psychiatry"
	case PractitionerSpecialityPsychotherapy:
		return "Psychotherapy"
	case PractitionerSpecialityPublicHealthMedicine:
		return "Public health medicine"
	case PractitionerSpecialityPulmonaryMedicine:
		return "Pulmonary medicine"
	case PractitionerSpecialityRadiationOncology:
		return "Radiation oncology"
	case PractitionerSpecialityRadiology:
		return "Radiology"
	case PractitionerSpecialityRadiologyInterventionalRadiology:
		return "Radiology-Interventional radiology"
	case PractitionerSpecialityRehabilitation:
		return "Rehabilitation"
	case PractitionerSpecialityRespiteCare:
		return "Respite care"
	case PractitionerSpecialityRheumatology:
		return "Rheumatology"
	case PractitionerSpecialitySleepStudies:
		return "Sleep studies"
	case PractitionerSpecialitySurgeryBoneAndMarrowTransplantation:
		return "Surgery-Bone and marrow transplantation"
	case PractitionerSpecialitySurgeryBreastSurgery:
		return "Surgery-Breast surgery"
	case PractitionerSpecialitySurgeryCardiacSurgery:
		return "Surgery-Cardiac surgery"
	case PractitionerSpecialitySurgeryCardiothoracicTransplantation:
		return "Surgery-Cardiothoracic transplantation"
	case PractitionerSpecialitySurgeryColorectalSurgery:
		return "Surgery-Colorectal surgery"
	case PractitionerSpecialitySurgeryDentalEndodontics:
		return "Surgery-Dental-Endodontics"
	case PractitionerSpecialitySurgeryDentalOralAndMaxillofacialSurgery:
		return "Surgery-Dental-Oral and maxillofacial surgery"
	case PractitionerSpecialitySurgeryDentalOralSurgery:
		return "Surgery-Dental-Oral surgery"
	case PractitionerSpecialitySurgeryDentalOrthodontics:
		return "Surgery-Dental-Orthodontics"
	case PractitionerSpecialitySurgeryDentalPeriodontalSurgery:
		return "Surgery-Dental-Periodontal surgery"
	case PractitionerSpecialitySurgeryDentalProstheticDentistryProsthodontics:
		return "Surgery-Dental-Prosthetic dentistry (Prosthodontics)"
	case PractitionerSpecialitySurgeryDentalSurgicalProsthodontics:
		return "Surgery-Dental-surgical-Prosthodontics"
	case PractitionerSpecialitySurgeryDentistryRestorativeDentistry:
		return "Surgery-Dentistry-Restorative dentistry"
	case PractitionerSpecialitySurgeryDentistrySurgical:
		return "Surgery-Dentistry--surgical"
	case PractitionerSpecialitySurgeryDentistrySurgicalOrthodontics:
		return "Surgery-Dentistry-surgical-Orthodontics"
	case PractitionerSpecialitySurgeryDermatologicSurgery:
		return "Surgery-Dermatologic surgery"
	case PractitionerSpecialitySurgeryEarNoseAndThroatSurgery:
		return "Surgery-Ear, nose and throat surgery"
	case PractitionerSpecialitySurgeryGeneral:
		return "Surgery-general"
	case PractitionerSpecialitySurgeryHepatobiliaryAndPancreaticSurgery:
		return "Surgery-Hepatobiliary and pancreatic surgery"
	case PractitionerSpecialitySurgeryNeurosurgery:
		return "Surgery-Neurosurgery"
	case PractitionerSpecialitySurgeryPlasticSurgery:
		return "Surgery-Plastic surgery"
	case PractitionerSpecialitySurgeryTransplantationSurgery:
		return "Surgery-Transplantation surgery"
	case PractitionerSpecialitySurgeryTraumaAndOrthopedics:
		return "Surgery-Trauma and orthopedics"
	case PractitionerSpecialitySurgeryVascular:
		return "Surgery-Vascular"
	case PractitionerSpecialitySurgicalOncology:
		return "Surgical oncology"
	case PractitionerSpecialitySurgicalAccidentEmergency:
		return "Surgical-Accident & emergency"
	case PractitionerSpecialityThoracicMedicine:
		return "Thoracic medicine"
	case PractitionerSpecialityToxicology:
		return "Toxicology"
	case PractitionerSpecialityTropicalMedicine:
		return "Tropical medicine"
	case PractitionerSpecialityUrologicalOncology:
		return "Urological oncology"
	case PractitionerSpecialityUrology:
		return "Urology"
	case PractitionerSpecialityOtherMedicalSpecialty:
		return "Medical specialty--OTHER--NOT LISTED"
	case PractitionerSpecialityOtherSurgicalSpecialty:
		return "Surgical specialty--OTHER-NOT LISTED"

	default:
		return "Unknown UDI Device Status"
	}
}

/*
Adult mental illness
Anesthetics
Audiological medicine
Blood banking and transfusion medicine
Burns care
Cardiology
Clinical cytogenetics and molecular genetics
Clinical genetics
Clinical hematology
Clinical immunology
Clinical microbiology
Clinical neuro-physiology
Clinical oncology
Clinical pharmacology
Clinical physiology
Community medicine
Critical care medicine
Dental medicine specialties
Dental-General dental practice
Dermatology
Diabetic medicine
Dive medicine
Endocrinology
Family practice
Gastroenterology
General medical practice
General medicine
General pathology
General practice
Genito-urinary medicine
Geriatric medicine
Gynecological oncology
Gynecology
Hematopathology
Hepatology
Histopathology
Immunopathology
Infectious diseases
Internal medicine
Learning disability
Medical oncology
Medical ophthalmology
Military medicine
Nephrology
Neurology
Neuropathology
Nuclear medicine
Obstetrics
Obstetrics and gynecology
Occupational medicine
Ophthalmic surgery
Ophthalmology
Osteopathic manipulative medicine
Otolaryngology
Pain management
Palliative medicine
Pediatric (Child and adolescent) psychiatry
Pediatric cardiology
Pediatric dentistry
Pediatric endocrinology
Pediatric gastroenterology
Pediatric genetics
Pediatric hematology
Pediatric immunology
Pediatric infectious diseases
Pediatric nephrology
Pediatric oncology
Pediatric ophthalmology
Pediatric pulmonology
Pediatric rheumatology
Pediatric surgery
Pediatric surgery-bone marrow transplantation
Preventive medicine
Psychiatry
Psychotherapy
Public health medicine
Pulmonary medicine
Radiation oncology
Radiology
Radiology-Interventional radiology
Rehabilitation
Respite care
Rheumatology
Sleep studies
Surgery-Bone and marrow transplantation
Surgery-Breast surgery
Surgery-Cardiac surgery
Surgery-Cardiothoracic transplantation
Surgery-Colorectal surgery
Surgery-Dental-Endodontics
Surgery-Dental-Oral and maxillofacial surgery
Surgery-Dental-Oral surgery
Surgery-Dental-Orthodontics
Surgery-Dental-Periodontal surgery
Surgery-Dental-Prosthetic dentistry (Prosthodontics)
Surgery-Dental-surgical-Prosthodontics
Surgery-Dentistry-Restorative dentistry
Surgery-Dentistry--surgical
Surgery-Dentistry-surgical-Orthodontics
Surgery-Dermatologic surgery
Surgery-Ear, nose and throat surgery
Surgery-general
Surgery-Hepatobiliary and pancreatic surgery
Surgery-Neurosurgery
Surgery-Plastic surgery
Surgery-Transplantation surgery
Surgery-Trauma and orthopedics
Surgery-Vascular
Surgical oncology
Surgical-Accident & emergency
Thoracic medicine
Toxicology
Tropical medicine
Urological oncology
Urology
Medical specialty--OTHER--NOT LISTED
Surgical specialty--OTHER-NOT LISTED

*/
