package code

import (
	"fmt"
	"strings"
)

type ContactPointSystem string

const (
	ContactPointSystemPhone ContactPointSystem = "phone"
	ContactPointSystemFax   ContactPointSystem = "fax"
	ContactPointSystemEmail ContactPointSystem = "email"
	ContactPointSystemPager ContactPointSystem = "pager"
	ContactPointSystemURL   ContactPointSystem = "url"
	ContactPointSystemSMS   ContactPointSystem = "sms"
	ContactPointSystemOther ContactPointSystem = "other"
)

func AllContactPointSystem() []ContactPointSystem {
	return []ContactPointSystem{
		ContactPointSystemPhone,
		ContactPointSystemFax,
		ContactPointSystemEmail,
		ContactPointSystemPager,
		ContactPointSystemURL,
		ContactPointSystemSMS,
		ContactPointSystemOther,
	}
}

func FindContactPointSystem(filter string) []ContactPointSystem {
	ret := make([]ContactPointSystem, 0)
	for _, at := range AllContactPointSystem() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ContactPointSystem) ToString() {
	fmt.Println(au.String())
}
func (au ContactPointSystem) String() string {
	switch au {
	case ContactPointSystemPhone:
		return "Phone"
	case ContactPointSystemFax:
		return "Fax"
	case ContactPointSystemEmail:
		return "Email"
	case ContactPointSystemPager:
		return "Pager"
	case ContactPointSystemURL:
		return "URL"
	case ContactPointSystemSMS:
		return "SMS"
	case ContactPointSystemOther:
		return "Other"
	default:
		return "Unknown Contact Point System"
	}
}
