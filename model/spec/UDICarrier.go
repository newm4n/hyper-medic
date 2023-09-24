package modelspec

import "github.com/newm4n/hyper-medic/model/spec/code"

type IUDICarrier interface {
	GetIdentifier() string
	GetIssuerURI()
	GetJurisdictionURI()
	GetCarrierAIDC() []byte
	GetCarrierHRF() string
	GetEntryType() code.UDIEntryType
}
