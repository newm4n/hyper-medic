package code

import (
	"fmt"
	"strings"
)

type MaritalStatus int

const (
	MaritalStatusAnnulled MaritalStatus = iota
	MaritalStatusDivorced
	MaritalStatusInterlocutory
	MaritalStatusSeparated
	MaritalStatusMarried
	MaritalStatusCommonLaw
	MaritalStatusPolygamus
	MaritalStatusDomesticPartner
	MaritalStatusUnMarried
	MaritalStatusNeverMarried
	MaritalStatusWidowed
	MaritalStatusUnknown
)

func AllMaritalStatuses() []MaritalStatus {
	return []MaritalStatus{
		MaritalStatusAnnulled,
		MaritalStatusDivorced,
		MaritalStatusInterlocutory,
		MaritalStatusSeparated,
		MaritalStatusMarried,
		MaritalStatusCommonLaw,
		MaritalStatusPolygamus,
		MaritalStatusDomesticPartner,
		MaritalStatusUnMarried,
		MaritalStatusNeverMarried,
		MaritalStatusWidowed,
		MaritalStatusUnknown,
	}
}

func FindMaritalStatus(filter string) []MaritalStatus {
	ret := make([]MaritalStatus, 0)
	for _, at := range AllMaritalStatuses() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (mt MaritalStatus) ToString() {
	fmt.Println(mt.String())
}

func (mt MaritalStatus) String() string {
	switch mt {
	case MaritalStatusAnnulled:
		return "Annulled"
	case MaritalStatusDivorced:
		return "Divorced"
	case MaritalStatusInterlocutory:
		return "Interlocutory"
	case MaritalStatusSeparated:
		return "Separated"
	case MaritalStatusMarried:
		return "Married"
	case MaritalStatusCommonLaw:
		return "Common Law"
	case MaritalStatusPolygamus:
		return "Polygamus"
	case MaritalStatusDomesticPartner:
		return "Partner"
	case MaritalStatusUnMarried:
		return "Married"
	case MaritalStatusNeverMarried:
		return "Never Married"
	case MaritalStatusWidowed:
		return "Widowed"
	case MaritalStatusUnknown:
		return "Unknown"
	default:
		return "Unknown Gender"
	}
}
