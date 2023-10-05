package code

import (
	"fmt"
	"strings"
)

type HttpVerb string

const (
	HttpVerbGET    HttpVerb = "GET"
	HttpVerbHEAD   HttpVerb = "HEAD"
	HttpVerbPOST   HttpVerb = "POST"
	HttpVerbPUT    HttpVerb = "PUT"
	HttpVerbDELETE HttpVerb = "DELETE"
	HttpVerbPATCH  HttpVerb = "PATCH"
)

func AllHttpVerb() []HttpVerb {
	return []HttpVerb{
		HttpVerbGET,
		HttpVerbHEAD,
		HttpVerbPOST,
		HttpVerbPUT,
		HttpVerbDELETE,
		HttpVerbPATCH,
	}
}

func FindHttpVerb(filter string) []HttpVerb {
	ret := make([]HttpVerb, 0)
	for _, at := range AllHttpVerb() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au HttpVerb) ToString() {
	fmt.Println(au.String())
}
func (au HttpVerb) String() string {
	switch au {
	case HttpVerbGET:
		return "GET"
	case HttpVerbHEAD:
		return "HEAD"
	case HttpVerbPOST:
		return "POST"
	case HttpVerbPUT:
		return "PUT"
	case HttpVerbDELETE:
		return "DELETE"
	case HttpVerbPATCH:
		return "PATCH"
	default:
		return "Unknown Http Verb"
	}
}
