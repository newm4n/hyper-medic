package code

import (
	"fmt"
	"strings"
)

type AdverseEventCategory int

const (
	AdverseEventCategoryWrongPatient AdverseEventCategory = iota
	AdverseEventCategoryProcedureMishap
	AdverseEventCategoryMedicationMishap
	AdverseEventCategoryDevice
	AdverseEventCategoryUnsafePhysicalEnvironment
	AdverseEventCategoryHospitalAcquiredInfection
	AdverseEventCategoryWrongBodySite
)

func AllAdverseEventCategory() []AdverseEventCategory {
	return []AdverseEventCategory{
		AdverseEventCategoryWrongPatient,
		AdverseEventCategoryProcedureMishap,
		AdverseEventCategoryMedicationMishap,
		AdverseEventCategoryDevice,
		AdverseEventCategoryUnsafePhysicalEnvironment,
		AdverseEventCategoryHospitalAcquiredInfection,
		AdverseEventCategoryWrongBodySite,
	}
}

func FindAdverseEventCategory(filter string) []AdverseEventCategory {
	ret := make([]AdverseEventCategory, 0)
	for _, at := range AllAdverseEventCategory() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventCategory) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventCategory) String() string {
	switch cpt {
	case AdverseEventCategoryWrongPatient:
		return "Wrong Patient"
	case AdverseEventCategoryProcedureMishap:
		return "Procedure Mishap"
	case AdverseEventCategoryMedicationMishap:
		return "Medication Mishap"
	case AdverseEventCategoryDevice:
		return "Device"
	case AdverseEventCategoryUnsafePhysicalEnvironment:
		return "Unsafe Physical Environment"
	case AdverseEventCategoryHospitalAcquiredInfection:
		return "Hospital Acquired Infection"
	case AdverseEventCategoryWrongBodySite:
		return "Wrong Body Site"
	default:
		return "Unknown Adverse Event Category"
	}
}

/*
Wrong Patient
Procedure Mishap
Medication Mishap
Device
Unsafe Physical Environment
Hospital Acquired Infection
Wrong Body Site
*/
