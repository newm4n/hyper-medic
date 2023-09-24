package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IContactPoint interface {
	GetSystem() code.ContactPointType
	GetValue() string
	GetUse() code.ContactPointUse
	GetRank() uint
	GetPeriod() IPeriod
}
