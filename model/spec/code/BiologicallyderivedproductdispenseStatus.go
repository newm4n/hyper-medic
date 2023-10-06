package code

import (
	"fmt"
	"strings"
)

type BiologicallyderivedproductdispenseStatus string

const (
	BiologicallyderivedproductdispenseStatusPreparation    BiologicallyderivedproductdispenseStatus = "preparation"
	BiologicallyderivedproductdispenseStatusInProgress     BiologicallyderivedproductdispenseStatus = "in-progress"
	BiologicallyderivedproductdispenseStatusAllocated      BiologicallyderivedproductdispenseStatus = "allocated"
	BiologicallyderivedproductdispenseStatusIssued         BiologicallyderivedproductdispenseStatus = "issued"
	BiologicallyderivedproductdispenseStatusUnfulfilled    BiologicallyderivedproductdispenseStatus = "unfulfilled"
	BiologicallyderivedproductdispenseStatusReturned       BiologicallyderivedproductdispenseStatus = "returned"
	BiologicallyderivedproductdispenseStatusEnteredInError BiologicallyderivedproductdispenseStatus = "entered-in-error"
	BiologicallyderivedproductdispenseStatusUnknown        BiologicallyderivedproductdispenseStatus = "unknown"
)

func AllBiologicallyderivedproductdispenseStatus() []BiologicallyderivedproductdispenseStatus {
	return []BiologicallyderivedproductdispenseStatus{
		BiologicallyderivedproductdispenseStatusPreparation,
		BiologicallyderivedproductdispenseStatusInProgress,
		BiologicallyderivedproductdispenseStatusAllocated,
		BiologicallyderivedproductdispenseStatusIssued,
		BiologicallyderivedproductdispenseStatusUnfulfilled,
		BiologicallyderivedproductdispenseStatusReturned,
		BiologicallyderivedproductdispenseStatusEnteredInError,
		BiologicallyderivedproductdispenseStatusUnknown,
	}
}

func FindBiologicallyderivedproductdispenseStatus(filter string) []BiologicallyderivedproductdispenseStatus {
	ret := make([]BiologicallyderivedproductdispenseStatus, 0)
	for _, at := range AllBiologicallyderivedproductdispenseStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BiologicallyderivedproductdispenseStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BiologicallyderivedproductdispenseStatus) String() string {
	switch cpt {
	case BiologicallyderivedproductdispenseStatusPreparation:
		return "Preparation"
	case BiologicallyderivedproductdispenseStatusInProgress:
		return "In Progress"
	case BiologicallyderivedproductdispenseStatusAllocated:
		return "Allocated"
	case BiologicallyderivedproductdispenseStatusIssued:
		return "Issued"
	case BiologicallyderivedproductdispenseStatusUnfulfilled:
		return "Unfulfilled"
	case BiologicallyderivedproductdispenseStatusReturned:
		return "Returned"
	case BiologicallyderivedproductdispenseStatusEnteredInError:
		return "Entered in Error"
	case BiologicallyderivedproductdispenseStatusUnknown:
		return "Unknown"
	default:
		return "Unknown Biologically derived product dispense Status"
	}
}
