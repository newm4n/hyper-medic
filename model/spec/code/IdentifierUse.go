package code

import (
	"fmt"
	"strings"
)

type IdentifierUse string

const (
	IdentifierUseUsual     IdentifierUse = "usual"
	IdentifierUseOfficial  IdentifierUse = "official"
	IdentifierUseTemp      IdentifierUse = "temp"
	IdentifierUseSecondary IdentifierUse = "secondary"
	IdentifierUseOld       IdentifierUse = "old"
)

func AllIdentifierUse() []IdentifierUse {
	return []IdentifierUse{
		IdentifierUseUsual,
		IdentifierUseOfficial,
		IdentifierUseTemp,
		IdentifierUseSecondary,
		IdentifierUseOld,
	}
}

func FindIdentifierUse(filter string) []IdentifierUse {
	ret := make([]IdentifierUse, 0)
	for _, at := range AllIdentifierUse() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au IdentifierUse) ToString() {
	fmt.Println(au.String())
}
func (au IdentifierUse) String() string {
	switch au {
	case IdentifierUseUsual:
		return "Usual"
	case IdentifierUseOfficial:
		return "Official"
	case IdentifierUseTemp:
		return "Temp"
	case IdentifierUseSecondary:
		return "Secondary"
	case IdentifierUseOld:
		return "Old"
	default:
		return "Unknown Identifier Use"
	}
}
