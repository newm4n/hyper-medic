package code

import (
	"fmt"
	"strings"
)

type ConceptPropertyType string

const (
	ConceptPropertyTypeCode     ConceptPropertyType = "code"
	ConceptPropertyTypeCoding   ConceptPropertyType = "Coding"
	ConceptPropertyTypeString   ConceptPropertyType = "string"
	ConceptPropertyTypeInteger  ConceptPropertyType = "integer"
	ConceptPropertyTypeBoolean  ConceptPropertyType = "boolean"
	ConceptPropertyTypeDateTime ConceptPropertyType = "dateTime"
	ConceptPropertyTypeDecimal  ConceptPropertyType = "decimal"
)

func AllConceptPropertyType() []ConceptPropertyType {
	return []ConceptPropertyType{
		ConceptPropertyTypeCode,
		ConceptPropertyTypeCoding,
		ConceptPropertyTypeString,
		ConceptPropertyTypeInteger,
		ConceptPropertyTypeBoolean,
		ConceptPropertyTypeDateTime,
		ConceptPropertyTypeDecimal,
	}
}

func FindConceptPropertyType(filter string) []ConceptPropertyType {
	ret := make([]ConceptPropertyType, 0)
	for _, at := range AllConceptPropertyType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ConceptPropertyType) ToString() {
	fmt.Println(au.String())
}
func (au ConceptPropertyType) String() string {
	switch au {
	case ConceptPropertyTypeCode:
		return "code (internal reference)"
	case ConceptPropertyTypeCoding:
		return "Coding (external reference)"
	case ConceptPropertyTypeString:
		return "string"
	case ConceptPropertyTypeInteger:
		return "integer"
	case ConceptPropertyTypeBoolean:
		return "boolean"
	case ConceptPropertyTypeDateTime:
		return "dateTime"
	case ConceptPropertyTypeDecimal:
		return "decimal"
	default:
		return "Unknown Concept Property Type"
	}
}
