package code

import (
	"fmt"
	"strings"
)

type CareTeamStatus int

const (
	CareTeamStatusProposed CareTeamStatus = iota
	CareTeamStatusActive
	CareTeamStatusSuspended
	CareTeamStatusInactive
	CareTeamStatusEnteredInError
)

func AllCareTeamStatus() []CareTeamStatus {
	return []CareTeamStatus{
		CareTeamStatusProposed,
		CareTeamStatusActive,
		CareTeamStatusSuspended,
		CareTeamStatusInactive,
		CareTeamStatusEnteredInError,
	}
}

func FindCareTeamStatus(filter string) []CareTeamStatus {
	ret := make([]CareTeamStatus, 0)
	for _, at := range AllCareTeamStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt CareTeamStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt CareTeamStatus) String() string {
	switch cpt {
	case CareTeamStatusProposed:
		return "proposed"
	case CareTeamStatusActive:
		return "active"
	case CareTeamStatusSuspended:
		return "suspended"
	case CareTeamStatusInactive:
		return "inactive"
	case CareTeamStatusEnteredInError:
		return "entered-in-error"
	default:
		return "Unknown Contact Point Type"
	}
}

/*
	proposed | active | suspended | inactive | entered-in-error
*/
