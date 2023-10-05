package code

import (
	"fmt"
	"strings"
)

type DurationUnits string

const (
	DurationUnitsMilliseconds DurationUnits = "ms"
	DurationUnitsSeconds      DurationUnits = "s"
	DurationUnitsMinutes      DurationUnits = "min"
	DurationUnitsHours        DurationUnits = "h"
	DurationUnitsDays         DurationUnits = "d"
	DurationUnitsWeeks        DurationUnits = "wk"
	DurationUnitsMonths       DurationUnits = "mo"
	DurationUnitsYears        DurationUnits = "a"
)

func AllDurationUnits() []DurationUnits {
	return []DurationUnits{
		DurationUnitsMilliseconds,
		DurationUnitsSeconds,
		DurationUnitsMinutes,
		DurationUnitsHours,
		DurationUnitsDays,
		DurationUnitsWeeks,
		DurationUnitsMonths,
		DurationUnitsYears,
	}
}

func FindDurationUnits(filter string) []DurationUnits {
	ret := make([]DurationUnits, 0)
	for _, at := range AllDurationUnits() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DurationUnits) ToString() {
	fmt.Println(au.String())
}
func (au DurationUnits) String() string {
	switch au {
	case DurationUnitsMilliseconds:
		return "milliseconds"
	case DurationUnitsSeconds:
		return "seconds"
	case DurationUnitsMinutes:
		return "minutes"
	case DurationUnitsHours:
		return "hours"
	case DurationUnitsDays:
		return "days"
	case DurationUnitsWeeks:
		return "weeks"
	case DurationUnitsMonths:
		return "months"
	case DurationUnitsYears:
		return "years"
	default:
		return "Unknown Constraint Severity"
	}
}
