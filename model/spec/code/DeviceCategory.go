package code

import (
	"fmt"
	"strings"
)

type DeviceCategory int

const (
	DeviceCategoryActive DeviceCategory = iota
	DeviceCategorycommunicating
	DeviceCategoryDurableMedicalEquipment
	DeviceCategoryMaintenance
	DeviceCategoryImplantable
	DeviceCategoryInVitro
	DeviceCategoryPointOfCare
	DeviceCategorySingleUse
	DeviceCategoryReusable
	DeviceCategorySoftware
)

func AllDeviceCategory() []DeviceCategory {
	return []DeviceCategory{
		DeviceCategoryActive,
		DeviceCategorycommunicating,
		DeviceCategoryDurableMedicalEquipment,
		DeviceCategoryMaintenance,
		DeviceCategoryImplantable,
		DeviceCategoryInVitro,
		DeviceCategoryPointOfCare,
		DeviceCategorySingleUse,
		DeviceCategoryReusable,
		DeviceCategorySoftware,
	}
}

func FindDeviceCategory(filter string) []DeviceCategory {
	ret := make([]DeviceCategory, 0)
	for _, at := range AllDeviceCategory() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt DeviceCategory) ToString() {
	fmt.Println(cpt.String())
}

func (cpt DeviceCategory) String() string {
	switch cpt {
	case DeviceCategoryActive:
		return "Active"
	case DeviceCategorycommunicating:
		return "communicating"
	case DeviceCategoryDurableMedicalEquipment:
		return "Durable Medical Equipment"
	case DeviceCategoryMaintenance:
		return "Maintenance"
	case DeviceCategoryImplantable:
		return "Implantable"
	case DeviceCategoryInVitro:
		return "In vitro"
	case DeviceCategoryPointOfCare:
		return "Point of Care"
	case DeviceCategorySingleUse:
		return "Single Use"
	case DeviceCategoryReusable:
		return "Reusable"
	case DeviceCategorySoftware:
		return "Software"
	default:
		return "Unknown Device Name Type"
	}
}

/*
Active
communicating
Durable Medical Equipment
Maintenance
Implantable
In vitro
Point of Care
Single Use
Reusable
Software
*/
