package modelspec

type IPractitionerRole interface {
	GetIndentifier() string
	IsActive() bool
	GetPeriod() IPeriod
	GetPractitioner() IPractitioner
	GetOrganization() IOrganization
	GetCodes() []IPractitionerRole
	GetSpecialities()
	GetLocations()
	GetHealthCareServices()
	GetContacts()
	GetCharacteristics()
	GetCommunications()
	GetAvailabilities()
}
