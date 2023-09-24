package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IAddress interface {
	GetUse() code.AddressUse
	GetType() code.AddressType
	GetText() string
	GetLines() []string
	GetCity() string
	GetDistrict() string
	GetState() string
	GetPostalCode() string
	GetCountry() string
	GetPeriod() IPeriod
}
