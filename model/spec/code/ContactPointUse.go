package code

import (
	"fmt"
	"strings"
)

type ContactPointUse int

const (
	ContactPointUseHome ContactPointUse = iota
	ContactPointUseWork
	ContactPointUseTemp
	ContactPointUseOld
	ContactPointUseMobile
)

func AllContactPointUses() []ContactPointUse {
	return []ContactPointUse{
		ContactPointUseHome,
		ContactPointUseWork,
		ContactPointUseTemp,
		ContactPointUseOld,
		ContactPointUseMobile,
	}
}

func FindContactPointUses(filter string) []ContactPointUse {
	ret := make([]ContactPointUse, 0)
	for _, at := range AllContactPointUses() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpu ContactPointUse) ToString() {
	fmt.Println(cpu.String())
}
func (cpu ContactPointUse) String() string {
	switch cpu {
	case ContactPointUseHome:
		return "Home"
	case ContactPointUseWork:
		return "Work"
	case ContactPointUseTemp:
		return "Temporary"
	case ContactPointUseOld:
		return "Old"
	case ContactPointUseMobile:
		return "Mobile"
	default:
		return "Unknown Contact Point Use"
	}
}
