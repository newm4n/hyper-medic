package code

import (
	"fmt"
	"strings"
)

type AdverseEventSeriousness string

const (
	AdverseEventSeriousnessNonserious AdverseEventSeriousness = "non-serious"
	AdverseEventSeriousnessSerious    AdverseEventSeriousness = "serious"
)

func AllAdverseEventSeriousness() []AdverseEventSeriousness {
	return []AdverseEventSeriousness{
		AdverseEventSeriousnessNonserious,
		AdverseEventSeriousnessSerious,
	}
}

func FindAdverseEventSeriousness(filter string) []AdverseEventSeriousness {
	ret := make([]AdverseEventSeriousness, 0)
	for _, at := range AllAdverseEventSeriousness() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventSeriousness) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventSeriousness) String() string {
	switch cpt {
	case AdverseEventSeriousnessNonserious:
		return "Non-serious"
	case AdverseEventSeriousnessSerious:
		return "Serious"
	default:
		return "Unknown Adverse Event Seriousness"
	}
}

/*
Non-serious
Serious
*/
