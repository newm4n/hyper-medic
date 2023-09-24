package code

import (
	"fmt"
	"strings"
)

type UDIEntryType int

const (
	UDIEntryTypeBarcode UDIEntryType = iota
	UDIEntryTypeRFID
	UDIEntryTypeManual
	UDIEntryTypeCard
	UDIEntryTypeSelfReported
	UDIEntryTypeElectronicTransmission
	UDIEntryTypeUnknown
)

func AllUDIEntryType() []UDIEntryType {
	return []UDIEntryType{
		UDIEntryTypeBarcode,
		UDIEntryTypeRFID,
		UDIEntryTypeManual,
		UDIEntryTypeCard,
		UDIEntryTypeSelfReported,
		UDIEntryTypeElectronicTransmission,
		UDIEntryTypeUnknown,
	}
}

func FindUDIEntryType(filter string) []UDIEntryType {
	ret := make([]UDIEntryType, 0)
	for _, at := range AllUDIEntryType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt UDIEntryType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt UDIEntryType) String() string {
	switch cpt {
	case UDIEntryTypeBarcode:
		return "Barcode"
	case UDIEntryTypeRFID:
		return "RFID"
	case UDIEntryTypeManual:
		return "Manual"
	case UDIEntryTypeCard:
		return "Card"
	case UDIEntryTypeSelfReported:
		return "Self Reported"
	case UDIEntryTypeElectronicTransmission:
		return "Electronic Transmission"
	case UDIEntryTypeUnknown:
		return "Unknown"
	default:
		return "Unknown Contact Point Type"
	}
}

/*
Barcode
RFID
Manual
Card
Self Reported
Electronic Transmission
Unknown
*/
