package code

import (
	"fmt"
	"strings"
)

type DesignationUse string

const (
	DesignationUseFullySpecifiedName DesignationUse = "900000000000003001"
	DesignationUseSynonym            DesignationUse = "900000000000013009 "
)

func AllDesignationUse() []DesignationUse {
	return []DesignationUse{
		DesignationUseFullySpecifiedName,
		DesignationUseSynonym,
	}
}

func FindDesignationUse(filter string) []DesignationUse {
	ret := make([]DesignationUse, 0)
	for _, at := range AllDesignationUse() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DesignationUse) ToString() {
	fmt.Println(au.String())
}
func (au DesignationUse) String() string {
	switch au {
	case DesignationUseFullySpecifiedName:
		return "Fully specified name"
	case DesignationUseSynonym:
		return "Synonym (core metadata concept)"
	default:
		return "Unknown Designation Use"
	}
}
