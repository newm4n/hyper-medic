package code

import (
	"fmt"
	"strings"
)

type DaysOfWeek string

const (
	DaysOfWeekMonday    DaysOfWeek = "mon"
	DaysOfWeekTuesday   DaysOfWeek = "tue"
	DaysOfWeekWednesday DaysOfWeek = "wed"
	DaysOfWeekThursday  DaysOfWeek = "thu"
	DaysOfWeekFriday    DaysOfWeek = "fri"
	DaysOfWeekSaturday  DaysOfWeek = "sat"
	DaysOfWeekSunday    DaysOfWeek = "sun"
)

func AllDaysOfWeek() []DaysOfWeek {
	return []DaysOfWeek{
		DaysOfWeekMonday,
		DaysOfWeekTuesday,
		DaysOfWeekWednesday,
		DaysOfWeekThursday,
		DaysOfWeekFriday,
		DaysOfWeekSaturday,
		DaysOfWeekSunday,
	}
}

func FindDaysOfWeek(filter string) []DaysOfWeek {
	ret := make([]DaysOfWeek, 0)
	for _, at := range AllDaysOfWeek() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DaysOfWeek) ToString() {
	fmt.Println(au.String())
}
func (au DaysOfWeek) String() string {
	switch au {
	case DaysOfWeekMonday:
		return "Monday"
	case DaysOfWeekTuesday:
		return "Tuesday"
	case DaysOfWeekWednesday:
		return "Wednesday"
	case DaysOfWeekThursday:
		return "Thursday"
	case DaysOfWeekFriday:
		return "Friday"
	case DaysOfWeekSaturday:
		return "Saturday"
	case DaysOfWeekSunday:
		return "Sunday"
	default:
		return "Unknown Days Of Week"
	}
}
