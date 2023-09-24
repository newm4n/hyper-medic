package modelspec

import (
	"github.com/newm4n/hyper-medic/model/spec/code"
	"github.com/newm4n/hyper-medic/model/spec/types"
)

type IRelatedPerson interface {
	GetIndentifier() string
	IsActive() bool
	GetPatient() IPatient
	GetRelationshio() code.RelatedPersonRelationShipType
	GetName() IHumanName
	GetTelecom() []IContactPoint
	GetGender() code.Gender
	GetBirthDate() types.Date
	GetPhoto() []IAttachment
	GetPeriod() IPeriod
	GetCommunication() []ICommunication
}
