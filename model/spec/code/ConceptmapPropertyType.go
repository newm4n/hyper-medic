package code

import (
	"fmt"
	"strings"
)

type ConceptmapPropertyType string

const (
	ConceptmapPropertyTypeCode     ConceptmapPropertyType = "code"
	ConceptmapPropertyTypeCoding   ConceptmapPropertyType = "Coding"
	ConceptmapPropertyTypeString   ConceptmapPropertyType = "string"
	ConceptmapPropertyTypeInteger  ConceptmapPropertyType = "integer"
	ConceptmapPropertyTypeBoolean  ConceptmapPropertyType = "boolean"
	ConceptmapPropertyTypeDateTime ConceptmapPropertyType = "dateTime"
	ConceptmapPropertyTypeDecimal  ConceptmapPropertyType = "decimal"
)

func AllConceptmapPropertyType() []ConceptmapPropertyType {
	return []ConceptmapPropertyType{
		ConceptmapPropertyTypeCode,
		ConceptmapPropertyTypeCoding,
		ConceptmapPropertyTypeString,
		ConceptmapPropertyTypeInteger,
		ConceptmapPropertyTypeBoolean,
		ConceptmapPropertyTypeDateTime,
		ConceptmapPropertyTypeDecimal,
	}
}

func FindConceptmapPropertyType(filter string) []ConceptmapPropertyType {
	ret := make([]ConceptmapPropertyType, 0)
	for _, at := range AllConceptmapPropertyType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ConceptmapPropertyType) ToString() {
	fmt.Println(au.String())
}
func (au ConceptmapPropertyType) String() string {
	switch au {
	case ConceptmapPropertyTypeCode:
		return "code (internal reference)"
	case ConceptmapPropertyTypeCoding:
		return "Coding (external reference)"
	case ConceptmapPropertyTypeString:
		return "string"
	case ConceptmapPropertyTypeInteger:
		return "integer"
	case ConceptmapPropertyTypeBoolean:
		return "boolean"
	case ConceptmapPropertyTypeDateTime:
		return "dateTime"
	case ConceptmapPropertyTypeDecimal:
		return "decimal"
	default:
		return "Unknown Concept Property Type"
	}
}
