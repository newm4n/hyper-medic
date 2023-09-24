package modelspec

type IExtendedContactDetail interface {
	GetPurpose() string
	GetName() []IHumanName
	GetContactPoint() []IContactPoint
	GetAddress() IAddress
	GetPeriod() IPeriod
}
