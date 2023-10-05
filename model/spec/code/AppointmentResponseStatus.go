package code

import (
	"fmt"
	"strings"
)

type AppointmentResponseStatus string

const (
	AppointmentResponseStatusAccepted       AppointmentResponseStatus = "accepted"
	AppointmentResponseStatusDeclined       AppointmentResponseStatus = "declined"
	AppointmentResponseStatusTentative      AppointmentResponseStatus = "tentative"
	AppointmentResponseStatusNeedsAction    AppointmentResponseStatus = "needs-action"
	AppointmentResponseStatusEnteredInError AppointmentResponseStatus = "entered-in-error"
)

func AllAppointmentResponseStatus() []AppointmentResponseStatus {
	return []AppointmentResponseStatus{
		AppointmentResponseStatusAccepted,
		AppointmentResponseStatusDeclined,
		AppointmentResponseStatusTentative,
		AppointmentResponseStatusNeedsAction,
		AppointmentResponseStatusEnteredInError,
	}
}

func FindAppointmentResponseStatus(filter string) []AppointmentResponseStatus {
	ret := make([]AppointmentResponseStatus, 0)
	for _, at := range AllAppointmentResponseStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AppointmentResponseStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AppointmentResponseStatus) String() string {
	switch cpt {
	case AppointmentResponseStatusAccepted:
		return "Accepted"
	case AppointmentResponseStatusDeclined:
		return "Declined"
	case AppointmentResponseStatusTentative:
		return "Tentative"
	case AppointmentResponseStatusNeedsAction:
		return "Needs Action"
	case AppointmentResponseStatusEnteredInError:
		return "Entered in error"
	default:
		return "Unknown Appointment Recurrence Type"
	}
}

/**
Accepted
Declined
Tentative
Needs Action
Entered in error
*/
