package code

import (
	"fmt"
	"strings"
)

type ExtensionContextType string

const (
	ExtensionContextTypeFHIRPath     ExtensionContextType = "fhirpath"
	ExtensionContextTypeElementID    ExtensionContextType = "element"
	ExtensionContextTypeExtensionURL ExtensionContextType = "extension"
)

func AllExtensionContextType() []ExtensionContextType {
	return []ExtensionContextType{
		ExtensionContextTypeFHIRPath,
		ExtensionContextTypeElementID,
		ExtensionContextTypeExtensionURL,
	}
}

func FindExtensionContextType(filter string) []ExtensionContextType {
	ret := make([]ExtensionContextType, 0)
	for _, at := range AllExtensionContextType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ExtensionContextType) ToString() {
	fmt.Println(au.String())
}
func (au ExtensionContextType) String() string {
	switch au {
	case ExtensionContextTypeFHIRPath:
		return "FHIRPath"
	case ExtensionContextTypeElementID:
		return "Element ID"
	case ExtensionContextTypeExtensionURL:
		return "Extension URL"
	default:
		return "Unknown Constraint Severity"
	}
}
