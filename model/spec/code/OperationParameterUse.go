package code

import (
	"fmt"
	"strings"
)

type OperationParameterUse string

const (
	OperationParameterUseIn  OperationParameterUse = "in"
	OperationParameterUseOut OperationParameterUse = "out"
)

func AllOperationParameterUse() []OperationParameterUse {
	return []OperationParameterUse{
		OperationParameterUseIn,
		OperationParameterUseOut,
	}
}

func FindOperationParameterUse(filter string) []OperationParameterUse {
	ret := make([]OperationParameterUse, 0)
	for _, at := range AllOperationParameterUse() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au OperationParameterUse) ToString() {
	fmt.Println(au.String())
}
func (au OperationParameterUse) String() string {
	switch au {
	case OperationParameterUseIn:
		return "In"
	case OperationParameterUseOut:
		return "Out"
	default:
		return "Unknown Constraint Severity"
	}
}
