package code

import (
	"fmt"
	"strings"
)

type AppointmentRecurrenceType string

const (
	AppointmentRecurrenceTypeDaily   AppointmentRecurrenceType = "d"
	AppointmentRecurrenceTypeWeekly  AppointmentRecurrenceType = "wk"
	AppointmentRecurrenceTypeMonthly AppointmentRecurrenceType = "mo"
	AppointmentRecurrenceTypeYearly  AppointmentRecurrenceType = "a"
)

func AllAppointmentRecurrenceType() []AppointmentRecurrenceType {
	return []AppointmentRecurrenceType{
		AppointmentRecurrenceTypeDaily,
		AppointmentRecurrenceTypeWeekly,
		AppointmentRecurrenceTypeMonthly,
		AppointmentRecurrenceTypeYearly,
	}
}

func FindAppointmentRecurrenceType(filter string) []AppointmentRecurrenceType {
	ret := make([]AppointmentRecurrenceType, 0)
	for _, at := range AllAppointmentRecurrenceType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AppointmentRecurrenceType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AppointmentRecurrenceType) String() string {
	switch cpt {
	case AppointmentRecurrenceTypeDaily:
		return "Daily"
	case AppointmentRecurrenceTypeWeekly:
		return "Weekly"
	case AppointmentRecurrenceTypeMonthly:
		return "Monthly"
	case AppointmentRecurrenceTypeYearly:
		return "Yearly"
	default:
		return "Unknown Appointment Recurrence Type"
	}
}

/*
Daily
Weekly
Monthly
Yearly
*/
