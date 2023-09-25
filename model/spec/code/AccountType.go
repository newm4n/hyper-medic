package code

import (
	"fmt"
	"strings"
)

type AccountType int

const (
	AccountTypeActAccountCode AccountType = iota
	AccountTypeAccountReceivable
	AccountTypeCash
	AccountTypeCreditCard
	AccountTypeAmericanExpress
	AccountTypeDinersClub
	AccountTypeDiscoverCard
	AccountTypeMasterCard
	AccountTypeVisa
	AccountTypePatientBillingAccount
)

func AllAccountType() []AccountType {
	return []AccountType{
		AccountTypeActAccountCode,
		AccountTypeAccountReceivable,
		AccountTypeCash,
		AccountTypeCreditCard,
		AccountTypeAmericanExpress,
		AccountTypeDinersClub,
		AccountTypeDiscoverCard,
		AccountTypeMasterCard,
		AccountTypeVisa,
		AccountTypePatientBillingAccount,
	}
}

func FindAccountType(filter string) []AccountType {
	ret := make([]AccountType, 0)
	for _, at := range AllAccountType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AccountType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AccountType) String() string {
	switch cpt {
	case AccountTypeActAccountCode:
		return "ActAccountCode"
	case AccountTypeAccountReceivable:
		return "account Receivable"
	case AccountTypeCash:
		return "Cash"
	case AccountTypeCreditCard:
		return "Credit Card"
	case AccountTypeAmericanExpress:
		return "American Express"
	case AccountTypeDinersClub:
		return "Diner's Club"
	case AccountTypeDiscoverCard:
		return "Discover Card"
	case AccountTypeMasterCard:
		return "Master Card"
	case AccountTypeVisa:
		return "Visa"
	case AccountTypePatientBillingAccount:
		return "Patient Billing Account"
	default:
		return "Unknown Action Grouping Behavior"
	}
}

/*
ActAccountCode
account Receivable
Cash
Credit Card
American Express
Diner's Club
Discover Card
Master Card
Visa
Patient Billing Account

*/
