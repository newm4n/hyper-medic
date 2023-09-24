package code

import (
	"fmt"
	"strings"
)

type PractitionerQualification int

const (
	PractitionerQualificationAdvancedPracticeNurse PractitionerQualification = iota
	PractitionerQualificationAssociateOfAppliedScience
	PractitionerQualificationAssociateOfArts
	PractitionerQualificationAssociateOfBusinessAdministration
	PractitionerQualificationAssociateOfEngineering
	PractitionerQualificationAssociateOfScience
	PractitionerQualificationBachelorOfArts
	PractitionerQualificationBachelorOfBusinessAdministration
	PractitionerQualificationBachelorOfEngineering
	PractitionerQualificationBachelorOfFineArts
	PractitionerQualificationBachelorOfNursing
	PractitionerQualificationBachelorOfScience
	PractitionerQualificationBachelorOfScienceLaw
	PractitionerQualificationBachelorOnScienceNursing
	PractitionerQualificationBachelorOfTheology
	PractitionerQualificationCertificate
	PractitionerQualificationCertifiedAdultNursePractitioner
	PractitionerQualificationCertifiedMedicalAssistant
	PractitionerQualificationCertifiedNursePractitioner
	PractitionerQualificationCertifiedNurseMidwife
	PractitionerQualificationCertifiedRegisteredNurse
	PractitionerQualificationCertifiedNurseSpecialist
	PractitionerQualificationCertifiedPediatricNursePractitioner
	PractitionerQualificationCertifiedTumorRegistrar
	PractitionerQualificationDiploma
	PractitionerQualificationDoctorOfBusinessAdministration
	PractitionerQualificationDoctorOfEducation
	PractitionerQualificationDoctorOfPharmacy
	PractitionerQualificationDoctorOfEngineering
	PractitionerQualificationDoctorOfPhilosophy
	PractitionerQualificationDoctorOfScience
	PractitionerQualificationDoctorOfMedicine
	PractitionerQualificationDoctorOfOsteopathy
	PractitionerQualificationEmergencyMedicalTechnician
	PractitionerQualificationEmergencyMedicalTechnicianParamedic
	PractitionerQualificationFamilyPracticeNurse
	PractitionerQualificationHighSchoolGraduate
	PractitionerQualificationJurisDoctor
	PractitionerQualificationMasterOfArts
	PractitionerQualificationMasterOfBusinessAdministration
	PractitionerQualificationMasterOfCivilEngineering
	PractitionerQualificationMasterOfDivinity
	PractitionerQualificationMasterOfEducation
	PractitionerQualificationMasterOfElectricalEngineering
	PractitionerQualificationMasterOfEngineering
	PractitionerQualificationMasterOfFineArts
	PractitionerQualificationMasterOfMechanicalEngineering
	PractitionerQualificationMasterOfScience
	PractitionerQualificationMasterOfScienceLaw
	PractitionerQualificationMasterOfScienceNursing
	PractitionerQualificationMasterOfTheology
	PractitionerQualificationMedicalAssistant
	PractitionerQualificationMedicalTechnician
	PractitionerQualificationNonGraduate
	PractitionerQualificationNursePractitioner
	PractitionerQualificationPhysicianAssistant
	PractitionerQualificationRegisteredMedicalAssistant
	PractitionerQualificationRegisteredNurse
	PractitionerQualificationRegisteredPharmacist
	PractitionerQualificationSecretarialCertificate
	PractitionerQualificationTradeSchoolGraduate
)

func AllPractitionerQualifications() []PractitionerQualification {
	return []PractitionerQualification{
		PractitionerQualificationAdvancedPracticeNurse,
		PractitionerQualificationAssociateOfAppliedScience,
		PractitionerQualificationAssociateOfArts,
		PractitionerQualificationAssociateOfBusinessAdministration,
		PractitionerQualificationAssociateOfEngineering,
		PractitionerQualificationAssociateOfScience,
		PractitionerQualificationBachelorOfArts,
		PractitionerQualificationBachelorOfBusinessAdministration,
		PractitionerQualificationBachelorOfEngineering,
		PractitionerQualificationBachelorOfFineArts,
		PractitionerQualificationBachelorOfNursing,
		PractitionerQualificationBachelorOfScience,
		PractitionerQualificationBachelorOfScienceLaw,
		PractitionerQualificationBachelorOnScienceNursing,
		PractitionerQualificationBachelorOfTheology,
		PractitionerQualificationCertificate,
		PractitionerQualificationCertifiedAdultNursePractitioner,
		PractitionerQualificationCertifiedMedicalAssistant,
		PractitionerQualificationCertifiedNursePractitioner,
		PractitionerQualificationCertifiedNurseMidwife,
		PractitionerQualificationCertifiedRegisteredNurse,
		PractitionerQualificationCertifiedNurseSpecialist,
		PractitionerQualificationCertifiedPediatricNursePractitioner,
		PractitionerQualificationCertifiedTumorRegistrar,
		PractitionerQualificationDiploma,
		PractitionerQualificationDoctorOfBusinessAdministration,
		PractitionerQualificationDoctorOfEducation,
		PractitionerQualificationDoctorOfPharmacy,
		PractitionerQualificationDoctorOfEngineering,
		PractitionerQualificationDoctorOfPhilosophy,
		PractitionerQualificationDoctorOfScience,
		PractitionerQualificationDoctorOfMedicine,
		PractitionerQualificationDoctorOfOsteopathy,
		PractitionerQualificationEmergencyMedicalTechnician,
		PractitionerQualificationEmergencyMedicalTechnicianParamedic,
		PractitionerQualificationFamilyPracticeNurse,
		PractitionerQualificationHighSchoolGraduate,
		PractitionerQualificationJurisDoctor,
		PractitionerQualificationMasterOfArts,
		PractitionerQualificationMasterOfBusinessAdministration,
		PractitionerQualificationMasterOfCivilEngineering,
		PractitionerQualificationMasterOfDivinity,
		PractitionerQualificationMasterOfEducation,
		PractitionerQualificationMasterOfElectricalEngineering,
		PractitionerQualificationMasterOfEngineering,
		PractitionerQualificationMasterOfFineArts,
		PractitionerQualificationMasterOfMechanicalEngineering,
		PractitionerQualificationMasterOfScience,
		PractitionerQualificationMasterOfScienceLaw,
		PractitionerQualificationMasterOfScienceNursing,
		PractitionerQualificationMasterOfTheology,
		PractitionerQualificationMedicalAssistant,
		PractitionerQualificationMedicalTechnician,
		PractitionerQualificationNonGraduate,
		PractitionerQualificationNursePractitioner,
		PractitionerQualificationPhysicianAssistant,
		PractitionerQualificationRegisteredMedicalAssistant,
		PractitionerQualificationRegisteredNurse,
		PractitionerQualificationRegisteredPharmacist,
		PractitionerQualificationSecretarialCertificate,
		PractitionerQualificationTradeSchoolGraduate,
	}
}

func FindPractitionerQualification(filter string) []PractitionerQualification {
	ret := make([]PractitionerQualification, 0)
	for _, at := range AllPractitionerQualifications() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (pq PractitionerQualification) ToString() {
	fmt.Println(pq.String())
}
func (pq PractitionerQualification) String() string {
	switch pq {
	case PractitionerQualificationAdvancedPracticeNurse:
		return "Advanced Practice Nurse"
	case PractitionerQualificationAssociateOfAppliedScience:
		return "Associate of Applied Science"
	case PractitionerQualificationAssociateOfArts:
		return "Associate of Arts"
	case PractitionerQualificationAssociateOfBusinessAdministration:
		return "Associate of Business Administration"
	case PractitionerQualificationAssociateOfEngineering:
		return "Associate of Engineering"
	case PractitionerQualificationAssociateOfScience:
		return "Associate of Science"
	case PractitionerQualificationBachelorOfArts:
		return "Bachelor of Arts"
	case PractitionerQualificationBachelorOfBusinessAdministration:
		return "Bachelor of Business Administration"
	case PractitionerQualificationBachelorOfEngineering:
		return "Bachelor or Engineering"
	case PractitionerQualificationBachelorOfFineArts:
		return "Bachelor of Fine Arts"
	case PractitionerQualificationBachelorOfNursing:
		return "Bachelor of Nursing"
	case PractitionerQualificationBachelorOfScience:
		return "Bachelor of Science"
	case PractitionerQualificationBachelorOfScienceLaw:
		return "Bachelor of Science - Law"
	case PractitionerQualificationBachelorOnScienceNursing:
		return "Bachelor on Science - Nursing"
	case PractitionerQualificationBachelorOfTheology:
		return "Bachelor of Theology"
	case PractitionerQualificationCertificate:
		return "Certificate"
	case PractitionerQualificationCertifiedAdultNursePractitioner:
		return "Certified Adult Nurse Practitioner"
	case PractitionerQualificationCertifiedMedicalAssistant:
		return "Certified Medical Assistant"
	case PractitionerQualificationCertifiedNursePractitioner:
		return "Certified Nurse Practitioner"
	case PractitionerQualificationCertifiedNurseMidwife:
		return "Certified Nurse Midwife"
	case PractitionerQualificationCertifiedRegisteredNurse:
		return "Certified Registered Nurse"
	case PractitionerQualificationCertifiedNurseSpecialist:
		return "Certified Nurse Specialist"
	case PractitionerQualificationCertifiedPediatricNursePractitioner:
		return "Certified Pediatric Nurse Practitioner"
	case PractitionerQualificationCertifiedTumorRegistrar:
		return "Certified Tumor Registrar"
	case PractitionerQualificationDiploma:
		return "Diploma"
	case PractitionerQualificationDoctorOfBusinessAdministration:
		return "Doctor of Business Administration"
	case PractitionerQualificationDoctorOfEducation:
		return "Doctor of Education"
	case PractitionerQualificationDoctorOfPharmacy:
		return "Doctor of Pharmacy"
	case PractitionerQualificationDoctorOfEngineering:
		return "Doctor of Engineering"
	case PractitionerQualificationDoctorOfPhilosophy:
		return "Doctor of Philosophy"
	case PractitionerQualificationDoctorOfScience:
		return "Doctor of Science"
	case PractitionerQualificationDoctorOfMedicine:
		return "Doctor of Medicine"
	case PractitionerQualificationDoctorOfOsteopathy:
		return "Doctor of Osteopathy"
	case PractitionerQualificationEmergencyMedicalTechnician:
		return "Emergency Medical Technician"
	case PractitionerQualificationEmergencyMedicalTechnicianParamedic:
		return "Emergency Medical Technician - Paramedic"
	case PractitionerQualificationFamilyPracticeNurse:
		return "Family Practice Nurse Practitioner"
	case PractitionerQualificationHighSchoolGraduate:
		return "High School Graduate"
	case PractitionerQualificationJurisDoctor:
		return "Juris Doctor"
	case PractitionerQualificationMasterOfArts:
		return "Master of Arts"
	case PractitionerQualificationMasterOfBusinessAdministration:
		return "Master of Business Administration"
	case PractitionerQualificationMasterOfCivilEngineering:
		return "Master of Civil Engineering"
	case PractitionerQualificationMasterOfDivinity:
		return "Master of Divinity"
	case PractitionerQualificationMasterOfEducation:
		return "Master of Education"
	case PractitionerQualificationMasterOfElectricalEngineering:
		return "Master of Electrical Engineering"
	case PractitionerQualificationMasterOfEngineering:
		return "Master of Engineering"
	case PractitionerQualificationMasterOfFineArts:
		return "Master of Fine Arts"
	case PractitionerQualificationMasterOfMechanicalEngineering:
		return "Master of Mechanical Engineering"
	case PractitionerQualificationMasterOfScience:
		return "Master of Science"
	case PractitionerQualificationMasterOfScienceLaw:
		return "Master of Science - Law"
	case PractitionerQualificationMasterOfScienceNursing:
		return "Master of Science - Nursing"
	case PractitionerQualificationMasterOfTheology:
		return "Master of Theology"
	case PractitionerQualificationMedicalAssistant:
		return "Medical Assistant"
	case PractitionerQualificationMedicalTechnician:
		return "Medical Technician"
	case PractitionerQualificationNonGraduate:
		return "Non-Graduate"
	case PractitionerQualificationNursePractitioner:
		return "Nurse Practitioner"
	case PractitionerQualificationPhysicianAssistant:
		return "Physician Assistant"
	case PractitionerQualificationRegisteredMedicalAssistant:
		return "Registered Medical Assistant"
	case PractitionerQualificationRegisteredNurse:
		return "Registered Nurse"
	case PractitionerQualificationRegisteredPharmacist:
		return "Registered Pharmacist"
	case PractitionerQualificationSecretarialCertificate:
		return "Secretarial Certificate"
	case PractitionerQualificationTradeSchoolGraduate:
		return "Trade School Graduate"
	default:
		return "Unknown Qualification"
	}
}
