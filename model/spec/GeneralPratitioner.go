package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

const (
	Organization GeneralParctitionerType = iota
	Practitioner
	PractitionerRole
)

type GeneralParctitionerType int

type IGeneralPratitioner interface {
	GetGeneralParctitionerType() GeneralParctitionerType
	GetPractitioner() IPractitioner
	GetOrganzation() IOrganization
	GetPractitionerRole() code.ParticipantRole
}
