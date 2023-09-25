package code

import (
	"fmt"
	"strings"
)

type AccountBillingStatus int

const (
	AccountBillingStatusOpen AccountBillingStatus = iota
	AccountBillingStatusCareCompleteNotbilled
	AccountBillingStatusBilling
	AccountBillingStatusClosedBaddebt
	AccountBillingStatusClosedVoided
	AccountBillingStatusClosedCompleted
	AccountBillingStatusClosedCombined
)

func AllAccountBillingStatus() []AccountBillingStatus {
	return []AccountBillingStatus{
		AccountBillingStatusOpen,
		AccountBillingStatusCareCompleteNotbilled,
		AccountBillingStatusBilling,
		AccountBillingStatusClosedBaddebt,
		AccountBillingStatusClosedVoided,
		AccountBillingStatusClosedCompleted,
		AccountBillingStatusClosedCombined,
	}
}

func FindAccountBillingStatus(filter string) []AccountBillingStatus {
	ret := make([]AccountBillingStatus, 0)
	for _, at := range AllAccountBillingStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AccountBillingStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AccountBillingStatus) String() string {
	switch cpt {
	case AccountBillingStatusOpen:
		return "open"
	case AccountBillingStatusCareCompleteNotbilled:
		return "carecomplete-notbilled"
	case AccountBillingStatusBilling:
		return "billing"
	case AccountBillingStatusClosedBaddebt:
		return "closed-baddebt"
	case AccountBillingStatusClosedVoided:
		return "closed-voided"
	case AccountBillingStatusClosedCompleted:
		return "closed-completed"
	case AccountBillingStatusClosedCombined:
		return "closed-combined"
	default:
		return "Unknown Account Relationship"
	}
}

/*
open
 carecomplete-notbilled
billing
closed-baddebt
closed-voided
 closed-completed
 closed-combined
*/
