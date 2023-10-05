package code

import "strings"

type AddressType string

const (
	AddressTypePostal   AddressType = "postal"
	AddressTypePhysical AddressType = "physical"
	AddressTypeBoth     AddressType = "both"
)

func AllAddressTypes() []AddressType {
	return []AddressType{
		AddressTypePostal,
		AddressTypePhysical,
		AddressTypeBoth,
	}
}

func FindAddressTypes(filter string) []AddressType {
	ret := make([]AddressType, 0)
	for _, at := range AllAddressTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (at AddressType) String() string {
	switch at {
	case AddressTypePostal:
		return "Postal"
	case AddressTypePhysical:
		return "Physical"
	case AddressTypeBoth:
		return "Both"
	default:
		return "Unknown Address Type"
	}
}
