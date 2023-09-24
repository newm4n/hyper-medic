package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IContact interface {
	GetRelationShip() code.ContactRelationship
}
