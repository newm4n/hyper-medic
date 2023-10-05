package code

import (
	"fmt"
	"strings"
)

type IdentifierType string

const (
	IdentifierTypeDL   IdentifierType = "DL"
	IdentifierTypePPN  IdentifierType = "PPN"
	IdentifierTypeBRN  IdentifierType = "BRN"
	IdentifierTypeMR   IdentifierType = "MR"
	IdentifierTypeMCN  IdentifierType = "MCN"
	IdentifierTypeEN   IdentifierType = "EN"
	IdentifierTypeTAX  IdentifierType = "TAX"
	IdentifierTypeNIIP IdentifierType = "NIIP"
	IdentifierTypePRN  IdentifierType = "PRN"
	IdentifierTypeMD   IdentifierType = "MD"
	IdentifierTypeDR   IdentifierType = "DR"
	IdentifierTypeACSN IdentifierType = "ACSN"
	IdentifierTypeUDI  IdentifierType = "UDI"
	IdentifierTypeSNO  IdentifierType = "SNO"
	IdentifierTypeSB   IdentifierType = "SB"
	IdentifierTypePLAC IdentifierType = "PLAC"
	IdentifierTypeFILL IdentifierType = "FILL"
	IdentifierTypeJHN  IdentifierType = "JHN"
)

func AllIdentifierType() []IdentifierType {
	return []IdentifierType{
		IdentifierTypeDL,
		IdentifierTypePPN,
		IdentifierTypeBRN,
		IdentifierTypeMR,
		IdentifierTypeMCN,
		IdentifierTypeEN,
		IdentifierTypeTAX,
		IdentifierTypeNIIP,
		IdentifierTypePRN,
		IdentifierTypeMD,
		IdentifierTypeDR,
		IdentifierTypeACSN,
		IdentifierTypeUDI,
		IdentifierTypeSNO,
		IdentifierTypeSB,
		IdentifierTypePLAC,
		IdentifierTypeFILL,
		IdentifierTypeJHN,
	}
}

func FindIdentifierType(filter string) []IdentifierType {
	ret := make([]IdentifierType, 0)
	for _, at := range AllIdentifierType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au IdentifierType) ToString() {
	fmt.Println(au.String())
}
func (au IdentifierType) String() string {
	switch au {
	case IdentifierTypeDL:
		return "Driver's license number"
	case IdentifierTypePPN:
		return "Passport number"
	case IdentifierTypeBRN:
		return "Breed Registry Number"
	case IdentifierTypeMR:
		return "Medical record number"
	case IdentifierTypeMCN:
		return "Microchip Number"
	case IdentifierTypeEN:
		return "Employer number"
	case IdentifierTypeTAX:
		return "Tax ID number"
	case IdentifierTypeNIIP:
		return "National Insurance Payor Identifier (Payor)"
	case IdentifierTypePRN:
		return "Provider number"
	case IdentifierTypeMD:
		return "Medical License number"
	case IdentifierTypeDR:
		return "Donor Registration Number"
	case IdentifierTypeACSN:
		return "Accession ID"
	case IdentifierTypeUDI:
		return "Universal Device Identifier"
	case IdentifierTypeSNO:
		return "Serial Number"
	case IdentifierTypeSB:
		return "Social Beneficiary Identifier"
	case IdentifierTypePLAC:
		return "Placer Identifier"
	case IdentifierTypeFILL:
		return "Filler Identifier"
	case IdentifierTypeJHN:
		return "Jurisdictional health number"

	default:
		return "Unknown Identifier Type"
	}
}
