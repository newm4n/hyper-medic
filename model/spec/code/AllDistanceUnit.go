package code

import (
	"fmt"
	"strings"
)

type AllDistanceUnit int

const (
	AllDistanceUnitMilliMeter AllDistanceUnit = iota
	AllDistanceUnitCentiMeter
	AllDistanceUnitDesiMeter
	AllDistanceUnitMeter
	AllDistanceUnitKiloMeter
	AllDistanceUnitNanoMeter
)

func AllAllDistanceUnit() []AllDistanceUnit {
	return []AllDistanceUnit{
		AllDistanceUnitMilliMeter,
		AllDistanceUnitCentiMeter,
		AllDistanceUnitDesiMeter,
		AllDistanceUnitMeter,
		AllDistanceUnitKiloMeter,
		AllDistanceUnitNanoMeter,
	}
}

func FindAllDistanceUnit(filter string) []AllDistanceUnit {
	ret := make([]AllDistanceUnit, 0)
	for _, at := range AllAllDistanceUnit() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllDistanceUnit) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllDistanceUnit) String() string {
	switch cpt {
	case AllDistanceUnitMilliMeter:
		return "mm"
	case AllDistanceUnitCentiMeter:
		return "cm"
	case AllDistanceUnitDesiMeter:
		return "dm"
	case AllDistanceUnitMeter:
		return "m"
	case AllDistanceUnitKiloMeter:
		return "km"
	case AllDistanceUnitNanoMeter:
		return "nm"
	default:
		return "Unknown Allergy Intolerance Clinical"
	}
}

/*
m
dm
hm
km
*/
