package code

import (
	"fmt"
	"strings"
)

type AdministrationSubpotentReason int

const (
	AdministrationSubpotentReasonPartialDose AdministrationSubpotentReason = iota
	AdministrationSubpotentReasonVomited
	AdministrationSubpotentReasonColdChainBreak
	AdministrationSubpotentReasonManufacturerRecall
	AdministrationSubpotentReasonAdverseStorage
	AdministrationSubpotentReasonExpiredProduct
)

func AllAdministrationSubpotentReason() []AdministrationSubpotentReason {
	return []AdministrationSubpotentReason{
		AdministrationSubpotentReasonPartialDose,
		AdministrationSubpotentReasonVomited,
		AdministrationSubpotentReasonColdChainBreak,
		AdministrationSubpotentReasonManufacturerRecall,
		AdministrationSubpotentReasonAdverseStorage,
		AdministrationSubpotentReasonExpiredProduct,
	}
}

func FindAdministrationSubpotentReason(filter string) []AdministrationSubpotentReason {
	ret := make([]AdministrationSubpotentReason, 0)
	for _, at := range AllAdministrationSubpotentReason() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdministrationSubpotentReason) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdministrationSubpotentReason) String() string {
	switch cpt {
	case AdministrationSubpotentReasonPartialDose:
		return "Partial Dose"
	case AdministrationSubpotentReasonVomited:
		return "Vomited"
	case AdministrationSubpotentReasonColdChainBreak:
		return "Cold Chain Break"
	case AdministrationSubpotentReasonManufacturerRecall:
		return "Manufacturer Recall"
	case AdministrationSubpotentReasonAdverseStorage:
		return "Manufacturer Recall"
	case AdministrationSubpotentReasonExpiredProduct:
		return "Expired Product"
	default:
		return "Unknown Administration Subpotent Reason"
	}
}

/*
Partial Dose
Vomited
Cold Chain Break
Manufacturer Recall
Adverse Storage
Expired Product
*/
