package code

import (
	"fmt"
	"strings"
)

type AccountStatus string

const (
	AccountStatusActive         AccountStatus = "active"
	AccountStatusInactive       AccountStatus = "inactive"
	AccountStatusEnteredInError AccountStatus = "entered-in-error"
	AccountStatusUnHold         AccountStatus = "on-hold"
	AccountStatusUnknown        AccountStatus = "unknown"
)

func AllAccountStatus() []AccountStatus {
	return []AccountStatus{
		AccountStatusActive,
		AccountStatusInactive,
		AccountStatusEnteredInError,
		AccountStatusUnHold,
		AccountStatusUnknown,
	}
}

func FindAccountStatus(filter string) []AccountStatus {
	ret := make([]AccountStatus, 0)
	for _, at := range AllAccountStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AccountStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AccountStatus) String() string {
	switch cpt {
	case AccountStatusActive:
		return "Active"
	case AccountStatusInactive:
		return "Inactive"
	case AccountStatusEnteredInError:
		return "Entered In Error"
	case AccountStatusUnHold:
		return "UnHold"
	case AccountStatusUnknown:
		return "Unknown"
	default:
		return "Unknown Account Status"
	}
}

/*
Active
Inactive
EnteredInError
UnHold
Unknown
*/
