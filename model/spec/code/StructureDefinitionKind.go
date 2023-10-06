package code

import (
	"fmt"
	"strings"
)

type StructureDefinitionKind string

const (
	StructureDefinitionKindPrimitive StructureDefinitionKind = "primitive-type"
	StructureDefinitionKindComplex   StructureDefinitionKind = "complex-type"
	StructureDefinitionKindResource  StructureDefinitionKind = "resource"
	StructureDefinitionKindLogical   StructureDefinitionKind = "logical"
)

func AllStructureDefinitionKind() []StructureDefinitionKind {
	return []StructureDefinitionKind{
		StructureDefinitionKindPrimitive,
		StructureDefinitionKindComplex,
		StructureDefinitionKindResource,
		StructureDefinitionKindLogical,
	}
}

func FindStructureDefinitionKind(filter string) []StructureDefinitionKind {
	ret := make([]StructureDefinitionKind, 0)
	for _, at := range AllStructureDefinitionKind() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au StructureDefinitionKind) ToString() {
	fmt.Println(au.String())
}
func (au StructureDefinitionKind) String() string {
	switch au {
	case StructureDefinitionKindPrimitive:
		return "Primitive Data Type"
	case StructureDefinitionKindComplex:
		return "Complex Data Type"
	case StructureDefinitionKindResource:
		return "Resource"
	case StructureDefinitionKindLogical:
		return "Logical"
	default:
		return "Unknown Structure Definition Kind"
	}
}
