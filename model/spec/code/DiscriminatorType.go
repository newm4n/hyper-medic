package code

import (
	"fmt"
	"strings"
)

type DiscriminatorType string

const (
	DiscriminatorTypeValue    DiscriminatorType = "value"
	DiscriminatorTypeExists   DiscriminatorType = "exists"
	DiscriminatorTypePattern  DiscriminatorType = "pattern"
	DiscriminatorTypeType     DiscriminatorType = "type"
	DiscriminatorTypeProfile  DiscriminatorType = "profile"
	DiscriminatorTypePosition DiscriminatorType = "position"
)

func AllDiscriminatorType() []DiscriminatorType {
	return []DiscriminatorType{
		DiscriminatorTypeValue,
		DiscriminatorTypeExists,
		DiscriminatorTypePattern,
		DiscriminatorTypeType,
		DiscriminatorTypeProfile,
		DiscriminatorTypePosition,
	}
}

func FindDiscriminatorType(filter string) []DiscriminatorType {
	ret := make([]DiscriminatorType, 0)
	for _, at := range AllDiscriminatorType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DiscriminatorType) ToString() {
	fmt.Println(au.String())
}
func (au DiscriminatorType) String() string {
	switch au {
	case DiscriminatorTypeValue:
		return "Value"
	case DiscriminatorTypeExists:
		return "Exists"
	case DiscriminatorTypePattern:
		return "Pattern"
	case DiscriminatorTypeType:
		return "Type"
	case DiscriminatorTypeProfile:
		return "Profile"
	case DiscriminatorTypePosition:
		return "Position)"
	default:
		return "Unknown Discriminator Type"
	}
}
