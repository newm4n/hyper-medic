package code

import (
	"fmt"
	"strings"
)

type ConditionalReadStatus string

const (
	ConditionalReadStatusNotSupported    ConditionalReadStatus = "not-supported"
	ConditionalReadStatusIfModifiedSince ConditionalReadStatus = "modified-since"
	ConditionalReadStatusIfNoneMatch     ConditionalReadStatus = "not-match"
	ConditionalReadStatusFullSupport     ConditionalReadStatus = "full-support"
)

func AllConditionalReadStatus() []ConditionalReadStatus {
	return []ConditionalReadStatus{
		ConditionalReadStatusNotSupported,
		ConditionalReadStatusIfModifiedSince,
		ConditionalReadStatusIfNoneMatch,
		ConditionalReadStatusFullSupport,
	}
}

func FindConditionalReadStatus(filter string) []ConditionalReadStatus {
	ret := make([]ConditionalReadStatus, 0)
	for _, at := range AllConditionalReadStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ConditionalReadStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ConditionalReadStatus) String() string {
	switch cpt {
	case ConditionalReadStatusNotSupported:
		return "Not Supported"
	case ConditionalReadStatusIfModifiedSince:
		return "If-Modified-Since"
	case ConditionalReadStatusIfNoneMatch:
		return "If-None-Match"
	case ConditionalReadStatusFullSupport:
		return "Full Support"
	default:
		return "Unknown Conditional Read Status"
	}
}
