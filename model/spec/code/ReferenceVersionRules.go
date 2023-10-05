package code

import (
	"fmt"
	"strings"
)

type ReferenceVersionRules string

const (
	ReferenceVersionRulesEither      ReferenceVersionRules = "either"
	ReferenceVersionRulesIndependent ReferenceVersionRules = "independent"
	ReferenceVersionRulesSpecific    ReferenceVersionRules = "specific"
)

func AllReferenceVersionRules() []ReferenceVersionRules {
	return []ReferenceVersionRules{
		ReferenceVersionRulesEither,
		ReferenceVersionRulesIndependent,
		ReferenceVersionRulesSpecific,
	}
}

func FindReferenceVersionRules(filter string) []ReferenceVersionRules {
	ret := make([]ReferenceVersionRules, 0)
	for _, at := range AllReferenceVersionRules() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ReferenceVersionRules) ToString() {
	fmt.Println(au.String())
}
func (au ReferenceVersionRules) String() string {
	switch au {
	case ReferenceVersionRulesEither:
		return "Either Specific or independent"
	case ReferenceVersionRulesIndependent:
		return "Version independent"
	case ReferenceVersionRulesSpecific:
		return "Version Specific"
	default:
		return "Unknown Reference Version Rules"
	}
}
