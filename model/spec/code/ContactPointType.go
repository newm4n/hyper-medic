package code

import (
	"fmt"
	"strings"
)

type ContactPointType int

const (
	ContactPointTypePhone ContactPointType = iota
	ContactPointTypeFax
	ContactPointTypeEmail
	ContactPointTypePager
	ContactPointTypeURL
	ContactPointTypeSMS
	ContactPointTypeOther
)

func AllContactPointType() []ContactPointType {
	return []ContactPointType{
		ContactPointTypePhone,
		ContactPointTypeFax,
		ContactPointTypeEmail,
		ContactPointTypePager,
		ContactPointTypeURL,
		ContactPointTypeSMS,
		ContactPointTypeOther,
	}
}

func FindContactPointType(filter string) []ContactPointType {
	ret := make([]ContactPointType, 0)
	for _, at := range AllContactPointType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ContactPointType) ToString() {
	fmt.Println(cpt.String())
}
func (cpt ContactPointType) String() string {
	switch cpt {
	case ContactPointTypePhone:
		return "Billing Contact"
	case ContactPointTypeFax:
		return "Fax"
	case ContactPointTypeEmail:
		return "Email"
	case ContactPointTypePager:
		return "Pager"
	case ContactPointTypeURL:
		return "URL"
	case ContactPointTypeSMS:
		return "SMS"
	case ContactPointTypeOther:
		return "Other"
	default:
		return "Unknown Contact Point Type"
	}
}
