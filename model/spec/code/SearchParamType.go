package code

import (
	"fmt"
	"strings"
)

type SearchParamType string

const (
	SearchParamTypeNumber    SearchParamType = "number"
	SearchParamTypeDate      SearchParamType = "date"
	SearchParamTypeString    SearchParamType = "string"
	SearchParamTypeToken     SearchParamType = "token"
	SearchParamTypeReference SearchParamType = "reference"
	SearchParamTypeComposite SearchParamType = "composite"
	SearchParamTypeQuantity  SearchParamType = "quantity"
	SearchParamTypeURI       SearchParamType = "uri"
	SearchParamTypeSpecial   SearchParamType = "special"
)

func AllSearchParamType() []SearchParamType {
	return []SearchParamType{
		SearchParamTypeNumber,
		SearchParamTypeDate,
		SearchParamTypeString,
		SearchParamTypeToken,
		SearchParamTypeReference,
		SearchParamTypeComposite,
		SearchParamTypeQuantity,
		SearchParamTypeURI,
		SearchParamTypeSpecial,
	}
}

func FindSearchParamType(filter string) []SearchParamType {
	ret := make([]SearchParamType, 0)
	for _, at := range AllSearchParamType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au SearchParamType) ToString() {
	fmt.Println(au.String())
}
func (au SearchParamType) String() string {
	switch au {
	case SearchParamTypeNumber:
		return "Number"
	case SearchParamTypeDate:
		return "Date/DateTime"
	case SearchParamTypeString:
		return "String"
	case SearchParamTypeToken:
		return "Token"
	case SearchParamTypeReference:
		return "Reference"
	case SearchParamTypeComposite:
		return "Composite"
	case SearchParamTypeQuantity:
		return "Quantity"
	case SearchParamTypeURI:
		return "URI"
	case SearchParamTypeSpecial:
		return "Special"
	default:
		return "Unknown Search Param Type"
	}
}
