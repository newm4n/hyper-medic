package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IDeviceName interface {
	GetValue() string
	GetType() code.DeviceNameType
	IsDisplay() bool
}
