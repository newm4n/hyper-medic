package code

import (
	"fmt"
	"strings"
)

type ConceptmapAttributeType string

const (
	ConceptmapAttributeTypeCode     ConceptmapAttributeType = "code"
	ConceptmapAttributeTypeCoding   ConceptmapAttributeType = "Coding"
	ConceptmapAttributeTypeString   ConceptmapAttributeType = "string"
	ConceptmapAttributeTypeBoolean  ConceptmapAttributeType = "boolean"
	ConceptmapAttributeTypeQuantity ConceptmapAttributeType = "Quantity"
)

func AllConceptmapAttributeType() []ConceptmapAttributeType {
	return []ConceptmapAttributeType{
		ConceptmapAttributeTypeCode,
		ConceptmapAttributeTypeCoding,
		ConceptmapAttributeTypeString,
		ConceptmapAttributeTypeBoolean,
		ConceptmapAttributeTypeQuantity,
	}
}

func FindConceptmapAttributeType(filter string) []ConceptmapAttributeType {
	ret := make([]ConceptmapAttributeType, 0)
	for _, at := range AllConceptmapAttributeType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ConceptmapAttributeType) ToString() {
	fmt.Println(au.String())
}
func (au ConceptmapAttributeType) String() string {
	switch au {
	case ConceptmapAttributeTypeCode:
		return "code (internal reference)"
	case ConceptmapAttributeTypeCoding:
		return "Coding (external reference)"
	case ConceptmapAttributeTypeString:
		return "string"
	case ConceptmapAttributeTypeBoolean:
		return "boolean"
	case ConceptmapAttributeTypeQuantity:
		return "Quantity"
	default:
		return "Unknown Conceptmap Attribute Type"
	}
}
