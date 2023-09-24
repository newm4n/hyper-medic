package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IOrganization interface {
	GetIdentifier() string
	IsActive() bool
	GetTypes() []*code.OrganizationType
	GetName() string
	GetAliases() []string
	GetDescription() string
	GetContacts() IExtendedContactDetail
}
