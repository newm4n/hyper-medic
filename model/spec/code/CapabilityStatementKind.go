package code

import (
	"fmt"
	"strings"
)

type CapabilityStatementKind string

const (
	CapabilityStatementKindInstance     CapabilityStatementKind = "instance"
	CapabilityStatementKindCapability   CapabilityStatementKind = "capability"
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"
)

func AllCapabilityStatementKind() []CapabilityStatementKind {
	return []CapabilityStatementKind{
		CapabilityStatementKindInstance,
		CapabilityStatementKindCapability,
		CapabilityStatementKindRequirements,
	}
}

func FindCapabilityStatementKind(filter string) []CapabilityStatementKind {
	ret := make([]CapabilityStatementKind, 0)
	for _, at := range AllCapabilityStatementKind() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au CapabilityStatementKind) ToString() {
	fmt.Println(au.String())
}
func (au CapabilityStatementKind) String() string {
	switch au {
	case CapabilityStatementKindInstance:
		return "Instance"
	case CapabilityStatementKindCapability:
		return "Capability"
	case CapabilityStatementKindRequirements:
		return "Requirements"
	default:
		return "Unknown Capability Statement Kind"
	}
}
