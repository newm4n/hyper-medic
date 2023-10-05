package code

import (
	"fmt"
	"strings"
)

type SearchEntryMode string

const (
	SearchEntryModeMatch   SearchEntryMode = "match"
	SearchEntryModeInclude SearchEntryMode = "include"
	SearchEntryModeOutcome SearchEntryMode = "outcome"
)

func AllSearchEntryMode() []SearchEntryMode {
	return []SearchEntryMode{
		SearchEntryModeMatch,
		SearchEntryModeInclude,
		SearchEntryModeOutcome,
	}
}

func FindSearchEntryMode(filter string) []SearchEntryMode {
	ret := make([]SearchEntryMode, 0)
	for _, at := range AllSearchEntryMode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au SearchEntryMode) ToString() {
	fmt.Println(au.String())
}
func (au SearchEntryMode) String() string {
	switch au {
	case SearchEntryModeMatch:
		return "Match"
	case SearchEntryModeInclude:
		return "Include"
	case SearchEntryModeOutcome:
		return "Outcome"
	default:
		return "Unknown Search Entry Mode"
	}
}
