package code

import (
	"fmt"
	"strings"
)

type DeviceNameType int

const (
	DeviceNameTypeRegistered DeviceNameType = iota
	DeviceNameTypeUser
	DeviceNameTypePatient
)

func AllDeviceNameType() []DeviceNameType {
	return []DeviceNameType{
		DeviceNameTypeRegistered,
		DeviceNameTypeUser,
		DeviceNameTypePatient,
	}
}

func FindDeviceNameType(filter string) []DeviceNameType {
	ret := make([]DeviceNameType, 0)
	for _, at := range AllDeviceNameType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt DeviceNameType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt DeviceNameType) String() string {
	switch cpt {
	case DeviceNameTypeRegistered:
		return "Registered name"
	case DeviceNameTypeUser:
		return "User Friendly name"
	case DeviceNameTypePatient:
		return "Patient Reported name"
	default:
		return "Unknown Device Name Type"
	}
}

/*
Registered name
User Friendly name
Patient Reported name
*/
