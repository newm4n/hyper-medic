package code

import (
	"fmt"
	"strings"
)

type AppointmentStatus string

const (
	AppointmentStatusProposed       AppointmentStatus = "Proposed"
	AppointmentStatusPending        AppointmentStatus = "Pending"
	AppointmentStatusBooked         AppointmentStatus = "Booked"
	AppointmentStatusArrived        AppointmentStatus = "Arrived"
	AppointmentStatusFulfilled      AppointmentStatus = "Fulfilled"
	AppointmentStatusCancelled      AppointmentStatus = "Cancelled"
	AppointmentStatusNoShow         AppointmentStatus = "No Show"
	AppointmentStatusEnteredInError AppointmentStatus = "Entered in error"
	AppointmentStatusCheckedIn      AppointmentStatus = "Checked In"
	AppointmentStatusWaitlisted     AppointmentStatus = "Waitlisted"
)

func AllAppointmentStatus() []AppointmentStatus {
	return []AppointmentStatus{
		AppointmentStatusProposed,
		AppointmentStatusPending,
		AppointmentStatusBooked,
		AppointmentStatusArrived,
		AppointmentStatusFulfilled,
		AppointmentStatusCancelled,
		AppointmentStatusNoShow,
		AppointmentStatusEnteredInError,
		AppointmentStatusCheckedIn,
		AppointmentStatusWaitlisted,
	}
}

func FindAppointmentStatus(filter string) []AppointmentStatus {
	ret := make([]AppointmentStatus, 0)
	for _, at := range AllActionConditionKind() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AppointmentStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AppointmentStatus) String() string {
	switch cpt {
	case AppointmentStatusProposed:
		return "Proposed"
	case AppointmentStatusPending:
		return "Pending"
	case AppointmentStatusBooked:
		return "Booked"
	case AppointmentStatusArrived:
		return "Arrived"
	case AppointmentStatusFulfilled:
		return "Fulfilled"
	case AppointmentStatusCancelled:
		return "Cancelled"
	case AppointmentStatusNoShow:
		return "No Show"
	case AppointmentStatusEnteredInError:
		return "Entered in error"
	case AppointmentStatusCheckedIn:
		return "Checked In"
	case AppointmentStatusWaitlisted:
		return "Waitlisted"
	default:
		return "Unknown Appointment Status"
	}
}
