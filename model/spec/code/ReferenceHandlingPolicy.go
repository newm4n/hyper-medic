package code

import (
	"fmt"
	"strings"
)

type ReferenceHandlingPolicy string

const (
	ReferenceHandlingPolicyLiteral           ReferenceHandlingPolicy = "literal"
	ReferenceHandlingPolicyLogical           ReferenceHandlingPolicy = "logical"
	ReferenceHandlingPolicyResolves          ReferenceHandlingPolicy = "resolves"
	ReferenceHandlingPolicyIntegrityEnforced ReferenceHandlingPolicy = "enforced"
	ReferenceHandlingPolicyLocalOnly         ReferenceHandlingPolicy = "local"
)

func AllReferenceHandlingPolicy() []ReferenceHandlingPolicy {
	return []ReferenceHandlingPolicy{
		ReferenceHandlingPolicyLiteral,
		ReferenceHandlingPolicyLogical,
		ReferenceHandlingPolicyResolves,
		ReferenceHandlingPolicyIntegrityEnforced,
		ReferenceHandlingPolicyLocalOnly,
	}
}

func FindReferenceHandlingPolicy(filter string) []ReferenceHandlingPolicy {
	ret := make([]ReferenceHandlingPolicy, 0)
	for _, at := range AllReferenceHandlingPolicy() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ReferenceHandlingPolicy) ToString() {
	fmt.Println(au.String())
}
func (au ReferenceHandlingPolicy) String() string {
	switch au {
	case ReferenceHandlingPolicyLiteral:
		return "Literal References"
	case ReferenceHandlingPolicyLogical:
		return "Logical References"
	case ReferenceHandlingPolicyResolves:
		return "Resolves References"
	case ReferenceHandlingPolicyIntegrityEnforced:
		return "Reference Integrity Enforced"
	case ReferenceHandlingPolicyLocalOnly:
		return "Local References Only"
	default:
		return "Unknown Constraint Severity"
	}
}
