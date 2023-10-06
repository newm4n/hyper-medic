package code

import (
	"fmt"
	"strings"
)

type BiologicallyderivedproductdispenseMatchStatus string

const (
	BiologicallyderivedproductdispenseMatchStatusCrossmatched      BiologicallyderivedproductdispenseMatchStatus = "crossmatched"
	BiologicallyderivedproductdispenseMatchStatusSelected          BiologicallyderivedproductdispenseMatchStatus = "selected"
	BiologicallyderivedproductdispenseMatchStatusUnmatched         BiologicallyderivedproductdispenseMatchStatus = "unmatched"
	BiologicallyderivedproductdispenseMatchStatusLeastincompatible BiologicallyderivedproductdispenseMatchStatus = "least-incompatible"
)

func AllBiologicallyderivedproductdispenseMatchStatus() []BiologicallyderivedproductdispenseMatchStatus {
	return []BiologicallyderivedproductdispenseMatchStatus{
		BiologicallyderivedproductdispenseMatchStatusCrossmatched,
		BiologicallyderivedproductdispenseMatchStatusSelected,
		BiologicallyderivedproductdispenseMatchStatusUnmatched,
		BiologicallyderivedproductdispenseMatchStatusLeastincompatible,
	}
}

func FindBiologicallyderivedproductdispenseMatchStatus(filter string) []BiologicallyderivedproductdispenseMatchStatus {
	ret := make([]BiologicallyderivedproductdispenseMatchStatus, 0)
	for _, at := range AllBiologicallyderivedproductdispenseMatchStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BiologicallyderivedproductdispenseMatchStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BiologicallyderivedproductdispenseMatchStatus) String() string {
	switch cpt {
	case BiologicallyderivedproductdispenseMatchStatusCrossmatched:
		return "Crossmatched"
	case BiologicallyderivedproductdispenseMatchStatusSelected:
		return "Selected"
	case BiologicallyderivedproductdispenseMatchStatusUnmatched:
		return "Unmatched"
	case BiologicallyderivedproductdispenseMatchStatusLeastincompatible:
		return "Least incompatible"
	default:
		return "Unknown Biologically derived product dispense Match Status"
	}
}
