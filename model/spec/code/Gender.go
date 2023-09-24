package code

import (
	"fmt"
	"strings"
)

type Gender int

const (
	GenderMale Gender = iota
	GenderFemale
	GenderOther
	GenderUnknown
)

func AllGenders() []Gender {
	return []Gender{
		GenderMale,
		GenderFemale,
		GenderOther,
		GenderUnknown,
	}
}

func FindGender(filter string) []Gender {
	ret := make([]Gender, 0)
	for _, at := range AllGenders() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (g Gender) ToString() {
	fmt.Println(g.String())
}
func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "Male"
	case GenderFemale:
		return "Female"
	case GenderOther:
		return "Other"
	case GenderUnknown:
		return "Unknown"
	default:
		return "Unknown Gender"
	}
}
