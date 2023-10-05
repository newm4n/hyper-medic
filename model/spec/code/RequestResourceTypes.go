package code

import (
	"fmt"
	"strings"
)

type RequestResourceTypes string

const (
	RequestResourceTypesAppointment                RequestResourceTypes = "Appointment"
	RequestResourceTypesAppointmentResponse        RequestResourceTypes = "AppointmentResponse"
	RequestResourceTypesCarePlan                   RequestResourceTypes = "CarePlan"
	RequestResourceTypesClaim                      RequestResourceTypes = "Claim"
	RequestResourceTypesCommunicationRequest       RequestResourceTypes = "CommunicationRequest"
	RequestResourceTypesCoverageEligibilityRequest RequestResourceTypes = "CoverageEligibilityRequest"
	RequestResourceTypesDeviceRequest              RequestResourceTypes = "DeviceRequest"
	RequestResourceTypesEnrollmentRequest          RequestResourceTypes = "EnrollmentRequest"
	RequestResourceTypesImmunizationRecommendation RequestResourceTypes = "ImmunizationRecommendation"
	RequestResourceTypesMedicationRequest          RequestResourceTypes = "MedicationRequest"
	RequestResourceTypesNutritionOrder             RequestResourceTypes = "NutritionOrder"
	RequestResourceTypesRequestOrchestration       RequestResourceTypes = "RequestOrchestration"
	RequestResourceTypesServiceRequest             RequestResourceTypes = "ServiceRequest"
	RequestResourceTypesSupplyRequest              RequestResourceTypes = "SupplyRequest"
	RequestResourceTypesTask                       RequestResourceTypes = "Task"
	RequestResourceTypesTransport                  RequestResourceTypes = "Transport"
	RequestResourceTypesVisionPrescription         RequestResourceTypes = "VisionPrescription"
)

func AllRequestResourceTypes() []RequestResourceTypes {
	return []RequestResourceTypes{
		RequestResourceTypesAppointment,
		RequestResourceTypesAppointmentResponse,
		RequestResourceTypesCarePlan,
		RequestResourceTypesClaim,
		RequestResourceTypesCommunicationRequest,
		RequestResourceTypesCoverageEligibilityRequest,
		RequestResourceTypesDeviceRequest,
		RequestResourceTypesEnrollmentRequest,
		RequestResourceTypesImmunizationRecommendation,
		RequestResourceTypesMedicationRequest,
		RequestResourceTypesNutritionOrder,
		RequestResourceTypesRequestOrchestration,
		RequestResourceTypesServiceRequest,
		RequestResourceTypesSupplyRequest,
		RequestResourceTypesTask,
		RequestResourceTypesTransport,
		RequestResourceTypesVisionPrescription,
	}
}

func FindRequestResourceTypes(filter string) []RequestResourceTypes {
	ret := make([]RequestResourceTypes, 0)
	for _, at := range AllRequestResourceTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au RequestResourceTypes) ToString() {
	fmt.Println(au.String())
}
func (au RequestResourceTypes) String() string {
	switch au {
	case RequestResourceTypesAppointment:
		return "Appointment"
	case RequestResourceTypesAppointmentResponse:
		return "AppointmentResponse"
	case RequestResourceTypesCarePlan:
		return "CarePlan"
	case RequestResourceTypesClaim:
		return "Claim"
	case RequestResourceTypesCommunicationRequest:
		return "CommunicationRequest"
	case RequestResourceTypesCoverageEligibilityRequest:
		return "CoverageEligibilityRequest"
	case RequestResourceTypesDeviceRequest:
		return "DeviceRequest"
	case RequestResourceTypesEnrollmentRequest:
		return "EnrollmentRequest"
	case RequestResourceTypesImmunizationRecommendation:
		return "ImmunizationRecommendation"
	case RequestResourceTypesMedicationRequest:
		return "MedicationRequest"
	case RequestResourceTypesNutritionOrder:
		return "NutritionOrder"
	case RequestResourceTypesRequestOrchestration:
		return "RequestOrchestration"
	case RequestResourceTypesServiceRequest:
		return "ServiceRequest"
	case RequestResourceTypesSupplyRequest:
		return "SupplyRequest"
	case RequestResourceTypesTask:
		return "Task"
	case RequestResourceTypesTransport:
		return "Transport"
	case RequestResourceTypesVisionPrescription:
		return "VisionPrescription"
	default:
		return "Unknown Constraint Severity"
	}
}
