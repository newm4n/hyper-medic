package code

import (
	"fmt"
	"strings"
)

type AllTimeUnit int

const (
	AllTimeUnitMicroSecond AllTimeUnit = iota
	AllTimeUnitMillisecond
	AllTimeUnitSecond
	AllTimeUnitMinute
	AllTimeUnitHour
	AllTimeUnitDay
	AllTimeUnitWeek
	AllTimeUnitMonth
	AllTimeUnitYear
)

func AllAllTimeUnit() []AllTimeUnit {
	return []AllTimeUnit{
		AllTimeUnitMicroSecond,
		AllTimeUnitMillisecond,
		AllTimeUnitSecond,
		AllTimeUnitMinute,
		AllTimeUnitHour,
		AllTimeUnitDay,
		AllTimeUnitWeek,
		AllTimeUnitMonth,
		AllTimeUnitYear,
	}
}

func FindAllTimeUnit(filter string) []AllTimeUnit {
	ret := make([]AllTimeUnit, 0)
	for _, at := range AllAllTimeUnit() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllTimeUnit) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllTimeUnit) String() string {
	switch cpt {
	case AllTimeUnitMicroSecond:
		return "Micro Second"
	case AllTimeUnitMillisecond:
		return "Millisecond"
	case AllTimeUnitSecond:
		return "Second"
	case AllTimeUnitMinute:
		return "Minute"
	case AllTimeUnitHour:
		return "Hour"
	case AllTimeUnitDay:
		return "Day"
	case AllTimeUnitWeek:
		return "Week"
	case AllTimeUnitMonth:
		return "Month"
	case AllTimeUnitYear:
		return "Year"
	default:
		return "Unknown Appointment Recurrence Type"
	}
}

/*
Micro Second
Millisecond
Second
Minute
Hour
Day
Week
Month
Year
*/
