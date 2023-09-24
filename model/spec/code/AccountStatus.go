package code

type AccountStatus int

const (
	AccountStatusActive AccountStatus = iota
	AccountStatusInactive
	AccountStatusEnteredInError
	AccountStatusUnHold
	AccountStatusUnknown
)

/*
Active
Inactive
EnteredInError
UnHold
Unknown
*/
