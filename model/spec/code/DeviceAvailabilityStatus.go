package code

import (
	"fmt"
	"strings"
)

type DeviceAvailabilityStatus int

const (
	DeviceAvailabilityStatusLost DeviceAvailabilityStatus = iota
	DeviceAvailabilityStatusDamaged
	DeviceAvailabilityStatusDestroyed
	DeviceAvailabilityStatusAvailable
)

func AllDeviceAvailabilityStatus() []DeviceAvailabilityStatus {
	return []DeviceAvailabilityStatus{
		DeviceAvailabilityStatusLost,
		DeviceAvailabilityStatusDamaged,
		DeviceAvailabilityStatusDestroyed,
		DeviceAvailabilityStatusAvailable,
	}
}

func FindDeviceAvailabilityStatus(filter string) []DeviceAvailabilityStatus {
	ret := make([]DeviceAvailabilityStatus, 0)
	for _, at := range AllDeviceAvailabilityStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt DeviceAvailabilityStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt DeviceAvailabilityStatus) String() string {
	switch cpt {
	case DeviceAvailabilityStatusLost:
		return "Lost"
	case DeviceAvailabilityStatusDamaged:
		return "Damaged"
	case DeviceAvailabilityStatusDestroyed:
		return "Destroyed"
	case DeviceAvailabilityStatusAvailable:
		return "Available"
	default:
		return "Unknown Device Name Type"
	}
}

/*
Lost
Damaged
Destroyed
Available
*/
