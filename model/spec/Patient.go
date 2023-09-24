package modelspec

import (
	"github.com/newm4n/hyper-medic/model/spec/code"
	"github.com/newm4n/hyper-medic/model/spec/types"
)

type IPatient interface {
	IDomainResource
	GetIdentifier() string
	IsActive() bool
	GetNames() []IHumanName
	GetTelecoms() []IContactPoint
	GetGender() code.Gender
	GetBirthDate() types.Date
	GetDeceasedStatus() types.DeceaseStatus
	GetAddresses() []IAddress
	GetMaritalStatus() code.MaritalStatus
	GetMultipleBirthStatus() types.MultipleBirthStatus
	GetPhotos() []IAttachment
	GetContacts() []IContact
	GetCommunications() []ICommunication
	GetPractitioners() []IGeneralPratitioner
	GetManagingOrganization() []IOrganization
}
