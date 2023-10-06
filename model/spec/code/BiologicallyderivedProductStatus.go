package code

import (
	"fmt"
	"strings"
)

type BiologicallyderivedProductStatus string

const (
	BiologicallyderivedProductStatusAvailable   BiologicallyderivedProductStatus = "available"
	BiologicallyderivedProductStatusUnavailable BiologicallyderivedProductStatus = "unavailable"
)

func AllBiologicallyderivedProductStatus() []BiologicallyderivedProductStatus {
	return []BiologicallyderivedProductStatus{
		BiologicallyderivedProductStatusAvailable,
		BiologicallyderivedProductStatusUnavailable,
	}
}

func FindBiologicallyderivedProductStatus(filter string) []BiologicallyderivedProductStatus {
	ret := make([]BiologicallyderivedProductStatus, 0)
	for _, at := range AllBiologicallyderivedProductStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BiologicallyderivedProductStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BiologicallyderivedProductStatus) String() string {
	switch cpt {
	case BiologicallyderivedProductStatusAvailable:
		return "Available"
	case BiologicallyderivedProductStatusUnavailable:
		return "Unavailable"
	default:
		return "Unknown Biologically derived Product Status"
	}
}
