package code

import (
	"fmt"
	"strings"
)

type PropertyRepresentation string

const (
	PropertyRepresentationXMLAttribute  PropertyRepresentation = "xmlAttr"
	PropertyRepresentationXMLText       PropertyRepresentation = "xmlText"
	PropertyRepresentationTypeAttribute PropertyRepresentation = "typeAttr"
	PropertyRepresentationCDATextFormat PropertyRepresentation = "cdaText"
	PropertyRepresentationXHTML         PropertyRepresentation = "xhtml"
)

func AllPropertyRepresentation() []PropertyRepresentation {
	return []PropertyRepresentation{
		PropertyRepresentationXMLAttribute,
		PropertyRepresentationXMLText,
		PropertyRepresentationTypeAttribute,
		PropertyRepresentationCDATextFormat,
		PropertyRepresentationXHTML,
	}
}

func FindPropertyRepresentation(filter string) []PropertyRepresentation {
	ret := make([]PropertyRepresentation, 0)
	for _, at := range AllPropertyRepresentation() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au PropertyRepresentation) ToString() {
	fmt.Println(au.String())
}
func (au PropertyRepresentation) String() string {
	switch au {
	case PropertyRepresentationXMLAttribute:
		return "XML Attribute"
	case PropertyRepresentationXMLText:
		return "XML Text"
	case PropertyRepresentationTypeAttribute:
		return "Type Attribute"
	case PropertyRepresentationCDATextFormat:
		return "CDA Text Format"
	case PropertyRepresentationXHTML:
		return "XHTML"
	default:
		return "Unknown Property Representation"
	}
}
