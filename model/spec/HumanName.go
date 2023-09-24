package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IHumanName interface {
	GetUse() code.NameUse
	GetText() string
	GetFamily() string
	GetGiven() []string
	GetPrefix() []string
	GetSuffix() []string
	GetPeriod() IPeriod
}
