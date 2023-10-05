package code

import (
	"fmt"
	"strings"
)

type AllergyIntoleranceVerification string

const (
	AllergyIntoleranceVerificationUnconfirmed    AllergyIntoleranceVerification = "unconfirmed"
	AllergyIntoleranceVerificationPresumed       AllergyIntoleranceVerification = "presumed"
	AllergyIntoleranceVerificationConfirmed      AllergyIntoleranceVerification = "confirmed"
	AllergyIntoleranceVerificationRefuted        AllergyIntoleranceVerification = "refuted"
	AllergyIntoleranceVerificationEnteredInError AllergyIntoleranceVerification = "entered-in-error"
)

func AllAllergyIntoleranceVerification() []AllergyIntoleranceVerification {
	return []AllergyIntoleranceVerification{
		AllergyIntoleranceVerificationUnconfirmed,
		AllergyIntoleranceVerificationPresumed,
		AllergyIntoleranceVerificationConfirmed,
		AllergyIntoleranceVerificationRefuted,
		AllergyIntoleranceVerificationEnteredInError,
	}
}

func FindAllergyIntoleranceVerification(filter string) []AllergyIntoleranceVerification {
	ret := make([]AllergyIntoleranceVerification, 0)
	for _, at := range AllAllergyIntoleranceVerification() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AllergyIntoleranceVerification) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AllergyIntoleranceVerification) String() string {
	switch cpt {
	case AllergyIntoleranceVerificationUnconfirmed:
		return "Unconfirmed"
	case AllergyIntoleranceVerificationPresumed:
		return "Presumed"
	case AllergyIntoleranceVerificationConfirmed:
		return "Confirmed"
	case AllergyIntoleranceVerificationRefuted:
		return "Refuted"
	case AllergyIntoleranceVerificationEnteredInError:
		return "Entered in Error"
	default:
		return "Unknown Appointment Recurrence Type"
	}
}

/*
Unconfirmed
Presumed
Confirmed
Refuted
Entered in Error
*/
