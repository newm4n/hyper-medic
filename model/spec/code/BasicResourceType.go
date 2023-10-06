package code

import (
	"fmt"
	"strings"
)

type BasicResourceType string

const (
	BasicResourceTypeConsent       BasicResourceType = "consent"
	BasicResourceTypeReferral      BasicResourceType = "referral"
	BasicResourceTypeAdvevent      BasicResourceType = "advevent"
	BasicResourceTypeAptmtreq      BasicResourceType = "aptmtreq"
	BasicResourceTypeTransfer      BasicResourceType = "transfer"
	BasicResourceTypeDiet          BasicResourceType = "diet"
	BasicResourceTypeAdminact      BasicResourceType = "adminact"
	BasicResourceTypeExposure      BasicResourceType = "exposure"
	BasicResourceTypeInvestigation BasicResourceType = "investigation"
	BasicResourceTypeAccount       BasicResourceType = "account"
	BasicResourceTypeInvoice       BasicResourceType = "invoice"
	BasicResourceTypeAdjudicat     BasicResourceType = "adjudicat"
	BasicResourceTypePredetreq     BasicResourceType = "predetreq"
	BasicResourceTypePredetermine  BasicResourceType = "predetermine"
	BasicResourceTypeStudy         BasicResourceType = "study"
	BasicResourceTypeProtocol      BasicResourceType = "protocol"
)

func AllBasicResourceType() []BasicResourceType {
	return []BasicResourceType{
		BasicResourceTypeConsent,
		BasicResourceTypeReferral,
		BasicResourceTypeAdvevent,
		BasicResourceTypeAptmtreq,
		BasicResourceTypeTransfer,
		BasicResourceTypeDiet,
		BasicResourceTypeAdminact,
		BasicResourceTypeExposure,
		BasicResourceTypeInvestigation,
		BasicResourceTypeAccount,
		BasicResourceTypeInvoice,
		BasicResourceTypeAdjudicat,
		BasicResourceTypePredetreq,
		BasicResourceTypePredetermine,
		BasicResourceTypeStudy,
		BasicResourceTypeProtocol,
	}
}

func FindBasicResourceType(filter string) []BasicResourceType {
	ret := make([]BasicResourceType, 0)
	for _, at := range AllBasicResourceType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BasicResourceType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BasicResourceType) String() string {
	switch cpt {
	case BasicResourceTypeConsent:
		return "Consent"
	case BasicResourceTypeReferral:
		return "Referral"
	case BasicResourceTypeAdvevent:
		return "Adverse Event"
	case BasicResourceTypeAptmtreq:
		return "Appointment Request"
	case BasicResourceTypeTransfer:
		return "Transfer"
	case BasicResourceTypeDiet:
		return "Diet"
	case BasicResourceTypeAdminact:
		return "Administrative Activity"
	case BasicResourceTypeExposure:
		return "Exposure"
	case BasicResourceTypeInvestigation:
		return "Investigation"
	case BasicResourceTypeAccount:
		return "Account"
	case BasicResourceTypeInvoice:
		return "Invoice"
	case BasicResourceTypeAdjudicat:
		return "Invoice Adjudication"
	case BasicResourceTypePredetreq:
		return "Pre-determination Request"
	case BasicResourceTypePredetermine:
		return "Predetermination"
	case BasicResourceTypeStudy:
		return "Study"
	case BasicResourceTypeProtocol:
		return "Protocol"

	default:
		return "Unknown Action Grouping Behavior"
	}
}
