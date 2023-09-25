package code

import (
	"fmt"
	"strings"
)

type AgeUnit int

const (
	AgeUnitMilliSeconds AgeUnit = iota
	AgeUnitSeconds
	AgeUnitMinutes
	AgeUnitHours
	AgeUnitDays
	AgeUnitWeeks
	AgeUnitMonths
	AgeUnitYears
)

func AllAgeUnit() []AgeUnit {
	return []AgeUnit{
		AgeUnitMilliSeconds,
		AgeUnitSeconds,
		AgeUnitMinutes,
		AgeUnitHours,
		AgeUnitDays,
		AgeUnitWeeks,
		AgeUnitMonths,
		AgeUnitYears,
	}
}

func FindAgeUnit(filter string) []AgeUnit {
	ret := make([]AgeUnit, 0)
	for _, at := range AllAgeUnit() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AgeUnit) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AgeUnit) String() string {
	switch cpt {
	case AgeUnitMilliSeconds:
		return "Milli Second"
	case AgeUnitSeconds:
		return "Second"
	case AgeUnitMinutes:
		return "Minute"
	case AgeUnitHours:
		return "Hour"
	case AgeUnitDays:
		return "Day"
	case AgeUnitWeeks:
		return "Week"
	case AgeUnitMonths:
		return "Month"
	case AgeUnitYears:
		return "Year"
	default:
		return "Unknown Age Unit"
	}
}

/*
Minutes
Hours
Days
Weeks
Months
Years
*/
