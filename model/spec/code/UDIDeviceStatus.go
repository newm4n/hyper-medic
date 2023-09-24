package code

import (
	"fmt"
	"strings"
)

type UDIDeviceStatus int

const (
	UDIDeviceStatusActive UDIDeviceStatus = iota
	UDIDeviceStatusInactive
	UDIDeviceStatusEnteredInError
)

func AllUDIDeviceStatus() []UDIDeviceStatus {
	return []UDIDeviceStatus{
		UDIDeviceStatusActive,
		UDIDeviceStatusInactive,
		UDIDeviceStatusEnteredInError,
	}
}

func FindUDIDeviceStatus(filter string) []UDIDeviceStatus {
	ret := make([]UDIDeviceStatus, 0)
	for _, at := range AllUDIDeviceStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt UDIDeviceStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt UDIDeviceStatus) String() string {
	switch cpt {
	case UDIDeviceStatusActive:
		return "Active"
	case UDIDeviceStatusInactive:
		return "Inactive"
	case UDIDeviceStatusEnteredInError:
		return "Entered in Error"
	default:
		return "Unknown UDI Device Status"
	}
}

/*
Active
Inactive
Entered in Error
*/
