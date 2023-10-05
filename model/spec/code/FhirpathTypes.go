package code

import (
	"fmt"
	"strings"
)

type FhirpathTypes string

const (
	FhirpathTypesString   FhirpathTypes = "http://hl7.org/fhirpath/System.String"
	FhirpathTypesBoolean  FhirpathTypes = "http://hl7.org/fhirpath/System.Boolean"
	FhirpathTypesDate     FhirpathTypes = "http://hl7.org/fhirpath/System.Date"
	FhirpathTypesDateTime FhirpathTypes = "http://hl7.org/fhirpath/System.DateTime"
	FhirpathTypesDecimal  FhirpathTypes = "http://hl7.org/fhirpath/System.Decimal"
	FhirpathTypesInteger  FhirpathTypes = "http://hl7.org/fhirpath/System.Integer"
	FhirpathTypesTime     FhirpathTypes = "http://hl7.org/fhirpath/System.Time"
)

func AllFhirpathTypes() []FhirpathTypes {
	return []FhirpathTypes{
		FhirpathTypesString,
		FhirpathTypesBoolean,
		FhirpathTypesDate,
		FhirpathTypesDateTime,
		FhirpathTypesDecimal,
		FhirpathTypesInteger,
		FhirpathTypesTime,
	}
}

func FindFhirpathTypes(filter string) []FhirpathTypes {
	ret := make([]FhirpathTypes, 0)
	for _, at := range AllFhirpathTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au FhirpathTypes) ToString() {
	fmt.Println(au.String())
}
func (au FhirpathTypes) String() string {
	switch au {
	case FhirpathTypesString:
		return "String"
	case FhirpathTypesBoolean:
		return "Boolean"
	case FhirpathTypesDate:
		return "Date"
	case FhirpathTypesDateTime:
		return "DateTime"
	case FhirpathTypesDecimal:
		return "Decimal"
	case FhirpathTypesInteger:
		return "Integer"
	case FhirpathTypesTime:
		return "Time"
	default:
		return "Unknown Constraint Severity"
	}
}
