package modelspec

import (
	"github.com/newm4n/hyper-medic/model/spec/code"
	"time"
)

type IDevice interface {
	IDomainResource
	GetIdentifier() string
	GetDisplayName() string
	GetDefinitionURL() string
	GetUDICarrier() []IUDICarrier
	GetStatus() code.UDIDeviceStatus
	GetAvailabilityStatus() code.DeviceAvailabilityStatus
	GetBiologicalSourceEvent() string
	GetManufacturer() string
	GetManufactureDate() time.Time
	GetExpirationDate() time.Time
	GetIotNumber() string
	GetSerialNumber() string
	GetDeviceNames() []IDeviceName
	GetModelNumber() string
	GetPartNumber() string
	GetCategory() code.DeviceCategory
	GetType()
}
