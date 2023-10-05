package code

import (
	"fmt"
	"strings"
)

type ConditionalDeleteStatus string

const (
	ConditionalDeleteStatusNotSupported             ConditionalDeleteStatus = "not-supported"
	ConditionalDeleteStatusSingleDeletesSupported   ConditionalDeleteStatus = "single"
	ConditionalDeleteStatusMultipleDeletesSupported ConditionalDeleteStatus = "multiple"
)

func AllConditionalDeleteStatus() []ConditionalDeleteStatus {
	return []ConditionalDeleteStatus{
		ConditionalDeleteStatusNotSupported,
		ConditionalDeleteStatusSingleDeletesSupported,
		ConditionalDeleteStatusMultipleDeletesSupported,
	}
}

func FindConditionalDeleteStatus(filter string) []ConditionalDeleteStatus {
	ret := make([]ConditionalDeleteStatus, 0)
	for _, at := range AllConditionalDeleteStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ConditionalDeleteStatus) ToString() {
	fmt.Println(au.String())
}
func (au ConditionalDeleteStatus) String() string {
	switch au {
	case ConditionalDeleteStatusNotSupported:
		return "Not Supported"
	case ConditionalDeleteStatusSingleDeletesSupported:
		return "Single Deletes Supported"
	case ConditionalDeleteStatusMultipleDeletesSupported:
		return "Multiple Deletes Supported"
	default:
		return "Unknown Conditional Delete Status"
	}
}
