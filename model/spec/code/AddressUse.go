package code

import (
	"fmt"
	"strings"
)

type AddressUse string

const (
	AddressUseHome    AddressUse = "home"
	AddressUseWork    AddressUse = "work"
	AddressUseTemp    AddressUse = "temp"
	AddressUseOld     AddressUse = "old"
	AddressUseBilling AddressUse = "billing"
)

func AllAddressUse() []AddressUse {
	return []AddressUse{
		AddressUseHome,
		AddressUseWork,
		AddressUseTemp,
		AddressUseOld,
		AddressUseBilling,
	}
}

func FindAddressUse(filter string) []AddressUse {
	ret := make([]AddressUse, 0)
	for _, at := range AllAddressUse() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au AddressUse) ToString() {
	fmt.Println(au.String())
}
func (au AddressUse) String() string {
	switch au {
	case AddressUseHome:
		return "Home"
	case AddressUseWork:
		return "Work"
	case AddressUseTemp:
		return "Temporary"
	case AddressUseOld:
		return "Old"
	case AddressUseBilling:
		return "Billing"
	default:
		return "Unknown Address Use"
	}
}
