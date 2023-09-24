package code

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

/*
Wrong Patient
Procedure Mishap
Medication Mishap
Device
Unsafe Physical Environment
Hospital Acquired Infection
Wrong Body Site
*/
