package code

import (
	"fmt"
	"strings"
)

type OperationParameterScope string

const (
	OperationParameterScopeInstance OperationParameterScope = "instance"
	OperationParameterScopeType     OperationParameterScope = "type"
	OperationParameterScopeSystem   OperationParameterScope = "system"
)

func AllOperationParameterScope() []OperationParameterScope {
	return []OperationParameterScope{
		OperationParameterScopeInstance,
		OperationParameterScopeType,
		OperationParameterScopeSystem,
	}
}

func FindOperationParameterScope(filter string) []OperationParameterScope {
	ret := make([]OperationParameterScope, 0)
	for _, at := range AllOperationParameterScope() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au OperationParameterScope) ToString() {
	fmt.Println(au.String())
}
func (au OperationParameterScope) String() string {
	switch au {
	case OperationParameterScopeInstance:
		return "Instance"
	case OperationParameterScopeType:
		return "Type"
	case OperationParameterScopeSystem:
		return "System"
	default:
		return "Unknown Operation Parameter Scope"
	}
}
