package code

import (
	"fmt"
	"strings"
)

type ResourceAggregationMode string

const (
	ResourceAggregationModeContained  ResourceAggregationMode = "contained"
	ResourceAggregationModeReferenced ResourceAggregationMode = "referenced"
	ResourceAggregationModeBundled    ResourceAggregationMode = "bundled"
)

func AllResourceAggregationMode() []ResourceAggregationMode {
	return []ResourceAggregationMode{
		ResourceAggregationModeContained,
		ResourceAggregationModeReferenced,
		ResourceAggregationModeBundled,
	}
}

func FindResourceAggregationMode(filter string) []ResourceAggregationMode {
	ret := make([]ResourceAggregationMode, 0)
	for _, at := range AllResourceAggregationMode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ResourceAggregationMode) ToString() {
	fmt.Println(au.String())
}
func (au ResourceAggregationMode) String() string {
	switch au {
	case ResourceAggregationModeContained:
		return "Contained"
	case ResourceAggregationModeReferenced:
		return "Referenced"
	case ResourceAggregationModeBundled:
		return "Bundled"
	default:
		return "Unknown Resource Aggregation Mode"
	}
}
