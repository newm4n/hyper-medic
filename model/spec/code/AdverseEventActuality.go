package code

import (
	"fmt"
	"strings"
)

type AdverseEventActuality int

const (
	AdverseEventActualityActual AdverseEventActuality = iota
	AdverseEventActualityPotential
)

func AllAdverseEventActuality() []AdverseEventActuality {
	return []AdverseEventActuality{
		AdverseEventActualityActual,
		AdverseEventActualityPotential,
	}
}

func FindAdverseEventActuality(filter string) []AdverseEventActuality {
	ret := make([]AdverseEventActuality, 0)
	for _, at := range AllAdverseEventActuality() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventActuality) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventActuality) String() string {
	switch cpt {
	case AdverseEventActualityActual:
		return "Actual"
	case AdverseEventActualityPotential:
		return "Potential"
	default:
		return "Unknown Adverse Event Actuality"
	}
}

/*
Actual
Potential
*/
