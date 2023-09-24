package code

import (
	"fmt"
	"strings"
)

type NameUse int

const (
	NameUseUsual NameUse = iota
	NameUseOfficial
	NameUseTemp
	NameUseNickname
	NameUseAnonymous
	NameUseOld
	NameUseMaiden
)

func AllNameUses() []NameUse {
	return []NameUse{
		NameUseUsual,
		NameUseOfficial,
		NameUseTemp,
		NameUseNickname,
		NameUseAnonymous,
		NameUseOld,
		NameUseMaiden,
	}
}

func FindNameUse(filter string) []NameUse {
	ret := make([]NameUse, 0)
	for _, at := range AllNameUses() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (nu NameUse) ToString() {
	fmt.Println(nu.String())
}
func (nu NameUse) String() string {
	switch nu {
	case NameUseUsual:
		return "Usual"
	case NameUseOfficial:
		return "Official"
	case NameUseTemp:
		return "Temporary"
	case NameUseNickname:
		return "Nickname"
	case NameUseAnonymous:
		return "Anonymous"
	case NameUseOld:
		return "Old"
	case NameUseMaiden:
		return "Maiden"
	default:
		return "Unknown Name Usage"
	}
}
