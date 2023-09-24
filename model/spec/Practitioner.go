package modelspec

import (
	"github.com/newm4n/hyper-medic/model/spec/code"
	"github.com/newm4n/hyper-medic/model/spec/types"
)

type IPractitioner interface {
	IDomainResource
	GetIdentifier() string
	IsActive() bool
	GetNames() []IHumanName
	GetTelecoms() []IContactPoint
	GetGender() code.Gender
	GetBirthDate() types.Date
	GetDeceasedStatus() types.DeceaseStatus
	GetAddresses() []IAddress
	GetPhotos() []IAttachment
	GetCommunications() []ICommunication
}
