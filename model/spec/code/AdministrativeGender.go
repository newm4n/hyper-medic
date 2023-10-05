package code

import (
	"fmt"
	"strings"
)

type AdministrativeGender string

const (
	AdministrativeGenderMale    AdministrativeGender = "male"
	AdministrativeGenderFemale  AdministrativeGender = "female"
	AdministrativeGenderOther   AdministrativeGender = "other"
	AdministrativeGenderUnknown AdministrativeGender = "unknown"
)

func AllAdministrativeGender() []AdministrativeGender {
	return []AdministrativeGender{
		AdministrativeGenderMale,
		AdministrativeGenderFemale,
		AdministrativeGenderOther,
		AdministrativeGenderUnknown,
	}
}

func FindAdministrativeGender(filter string) []AdministrativeGender {
	ret := make([]AdministrativeGender, 0)
	for _, at := range AllAdministrativeGender() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdministrativeGender) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdministrativeGender) String() string {
	switch cpt {
	case AdministrativeGenderMale:
		return "Male"
	case AdministrativeGenderFemale:
		return "Female"
	case AdministrativeGenderOther:
		return "Other"
	case AdministrativeGenderUnknown:
		return "Unknown"
	default:
		return "Unknown Adverse Event Actuality"
	}
}

/*
Male
Female
Other
Unknown
*/
