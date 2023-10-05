package code

import (
	"fmt"
	"strings"
)

type OperationKind string

const (
	OperationKindOperation OperationKind = "operation"
	OperationKindQuery     OperationKind = "query"
)

func AllOperationKind() []OperationKind {
	return []OperationKind{
		OperationKindOperation,
		OperationKindQuery,
	}
}

func FindOperationKind(filter string) []OperationKind {
	ret := make([]OperationKind, 0)
	for _, at := range AllOperationKind() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au OperationKind) ToString() {
	fmt.Println(au.String())
}
func (au OperationKind) String() string {
	switch au {
	case OperationKindOperation:
		return "Operation"
	case OperationKindQuery:
		return "Query"
	default:
		return "Unknown Operation Kind"
	}
}
