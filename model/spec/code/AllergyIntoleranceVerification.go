package code

type AllergyIntoleranceVerification int

const (
	AllergyIntoleranceVerificationUnconfirmed AllergyIntoleranceVerification = iota
	AllergyIntoleranceVerificationPresumed
	AllergyIntoleranceVerificationConfirmed
	AllergyIntoleranceVerificationRefuted
	AllergyIntoleranceVerificationEnteredInError
)

/*
Unconfirmed
Presumed
Confirmed
Refuted
Entered in Error
*/
