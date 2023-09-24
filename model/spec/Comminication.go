package modelspec

type ICommunication interface {
	GetLanguage() string
	IsPreferred() bool
}
