package code

import (
	"fmt"
	"strings"
)

type AdverseEventOutcome int

const (
	AdverseEventOutcome1 AdverseEventOutcome = iota
	AdverseEventOutcome2
	AdverseEventOutcome3
	AdverseEventOutcome4
	AdverseEventOutcome5
)

func AllAdverseEventOutcome() []AdverseEventOutcome {
	return []AdverseEventOutcome{
		AdverseEventOutcome1,
		AdverseEventOutcome2,
		AdverseEventOutcome3,
		AdverseEventOutcome4,
		AdverseEventOutcome5,
	}
}

func FindAdverseEventOutcome(filter string) []AdverseEventOutcome {
	ret := make([]AdverseEventOutcome, 0)
	for _, at := range AllAdverseEventOutcome() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventOutcome) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventOutcome) String() string {
	switch cpt {
	case AdverseEventOutcome1:
		return "Transient abnormality unnoticed by the patient (finding)"
	case AdverseEventOutcome2:
		return "Transient abnormality with full recovery (finding)"
	case AdverseEventOutcome3:
		return "Adverse incident resulting in potentially permanent but not disabling damage (finding)"
	case AdverseEventOutcome4:
		return "Adverse incident resulting in potentially permanent disabling damage (finding)"
	case AdverseEventOutcome5:
		return "Adverse incident resulting in death (finding)"
	default:
		return "Unknown Adverse Event Outcome"
	}
}

/*
Transient abnormality unnoticed by the patient (finding)
Transient abnormality with full recovery (finding)
Adverse incident resulting in potentially permanent but not disabling damage (finding)
Adverse incident resulting in potentially permanent disabling damage (finding)
Adverse incident resulting in death (finding)
*/
