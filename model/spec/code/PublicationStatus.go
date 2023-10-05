package code

import (
	"fmt"
	"strings"
)

type PublicationStatus string

const (
	PublicationStatusDraft   PublicationStatus = "draft"
	PublicationStatusActive  PublicationStatus = "active"
	PublicationStatusRetired PublicationStatus = "retired"
	PublicationStatusUnknown PublicationStatus = "unknown"
)

func AllPublicationStatus() []PublicationStatus {
	return []PublicationStatus{
		PublicationStatusDraft,
		PublicationStatusActive,
		PublicationStatusRetired,
		PublicationStatusUnknown,
	}
}

func FindPublicationStatus(filter string) []PublicationStatus {
	ret := make([]PublicationStatus, 0)
	for _, at := range AllPublicationStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au PublicationStatus) ToString() {
	fmt.Println(au.String())
}
func (au PublicationStatus) String() string {
	switch au {
	case PublicationStatusDraft:
		return "Draft"
	case PublicationStatusActive:
		return "Active"
	case PublicationStatusRetired:
		return "Retired"
	case PublicationStatusUnknown:
		return "Unknown"
	default:
		return "Unknown Publication Status"
	}
}
