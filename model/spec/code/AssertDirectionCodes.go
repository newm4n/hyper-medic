package code

import (
	"fmt"
	"strings"
)

type AssertDirectionCodes string

const (
	AssertDirectionCodesResponse AssertDirectionCodes = "response"
	AssertDirectionCodesRequest  AssertDirectionCodes = "request"
)

func AllAssertDirectionCodes() []AssertDirectionCodes {
	return []AssertDirectionCodes{
		AssertDirectionCodesResponse,
		AssertDirectionCodesRequest,
	}
}

func FindAssertDirectionCodes(filter string) []AssertDirectionCodes {
	ret := make([]AssertDirectionCodes, 0)
	for _, at := range AllAssertDirectionCodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au AssertDirectionCodes) ToString() {
	fmt.Println(au.String())
}
func (au AssertDirectionCodes) String() string {
	switch au {
	case AssertDirectionCodesResponse:
		return "response"
	case AssertDirectionCodesRequest:
		return "request"
	default:
		return "Unknown Address Use"
	}
}

/*
response
request
*/
