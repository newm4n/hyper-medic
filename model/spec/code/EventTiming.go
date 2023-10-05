package code

import (
	"fmt"
	"strings"
)

type EventTiming string

const (
	EventTimingMorning        EventTiming = "MORN"
	EventTimingEarlyMorning   EventTiming = "MORN.early"
	EventTimingLateMorning    EventTiming = "MORN.late"
	EventTimingNoon           EventTiming = "NOON"
	EventTimingAfternoon      EventTiming = "AFT"
	EventTimingEarlyAfternoon EventTiming = "AFT.early"
	EventTimingLateAfternoon  EventTiming = "AFT.late"
	EventTimingEvening        EventTiming = "EVE"
	EventTimingEarlyEvening   EventTiming = "EVE.early"
	EventTimingLateEvening    EventTiming = "EVE.late"
	EventTimingNight          EventTiming = "NIGHT"
	EventTimingAfterSleep     EventTiming = "PHS"
	EventTimingImmediate      EventTiming = "IMD"
	EventTimingHs             EventTiming = "HS"
	EventTimingWake           EventTiming = "WAKE"
	EventTimingC              EventTiming = "C"
	EventTimingCm             EventTiming = "CM"
	EventTimingCd             EventTiming = "CD"
	EventTimingCv             EventTiming = "CV"
	EventTimingAc             EventTiming = "AC"
	EventTimingAcm            EventTiming = "ACM"
	EventTimingAcd            EventTiming = "ACD"
	EventTimingAcv            EventTiming = "ACV"
	EventTimingPc             EventTiming = "PC"
	EventTimingPcm            EventTiming = "PCM"
	EventTimingPcd            EventTiming = "PCD"
	EventTimingPcv            EventTiming = "PCV"
)

func AllEventTiming() []EventTiming {
	return []EventTiming{
		EventTimingMorning,
		EventTimingEarlyMorning,
		EventTimingLateMorning,
		EventTimingNoon,
		EventTimingAfternoon,
		EventTimingEarlyAfternoon,
		EventTimingLateAfternoon,
		EventTimingEvening,
		EventTimingEarlyEvening,
		EventTimingLateEvening,
		EventTimingNight,
		EventTimingAfterSleep,
		EventTimingImmediate,
		EventTimingHs,
		EventTimingWake,
		EventTimingC,
		EventTimingCm,
		EventTimingCd,
		EventTimingCv,
		EventTimingAc,
		EventTimingAcm,
		EventTimingAcd,
		EventTimingAcv,
		EventTimingPc,
		EventTimingPcm,
		EventTimingPcd,
		EventTimingPcv,
	}
}

func FindEventTiming(filter string) []EventTiming {
	ret := make([]EventTiming, 0)
	for _, at := range AllEventTiming() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au EventTiming) ToString() {
	fmt.Println(au.String())
}
func (au EventTiming) String() string {
	switch au {
	case EventTimingMorning:
		return "Morning"
	case EventTimingEarlyMorning:
		return "Early Morning"
	case EventTimingLateMorning:
		return "Late Morning"
	case EventTimingNoon:
		return "Noon"
	case EventTimingAfternoon:
		return "Afternoon"
	case EventTimingEarlyAfternoon:
		return "Early Afternoon"
	case EventTimingLateAfternoon:
		return "Late Afternoon"
	case EventTimingEvening:
		return "Evening"
	case EventTimingEarlyEvening:
		return "Early Evening"
	case EventTimingLateEvening:
		return "Late Evening"
	case EventTimingNight:
		return "Night"
	case EventTimingAfterSleep:
		return "After Sleep"
	case EventTimingImmediate:
		return "Immediate"
	case EventTimingHs:
		return "HS"
	case EventTimingWake:
		return "WAKE"
	case EventTimingC:
		return "C"
	case EventTimingCm:
		return "CM"
	case EventTimingCd:
		return "CD"
	case EventTimingCv:
		return "CV"
	case EventTimingAc:
		return "AC"
	case EventTimingAcm:
		return "ACM"
	case EventTimingAcd:
		return "ACD"
	case EventTimingAcv:
		return "ACV"
	case EventTimingPc:
		return "PC"
	case EventTimingPcm:
		return "PCM"
	case EventTimingPcd:
		return "PCD"
	case EventTimingPcv:
		return "PCV"
	default:
		return "Unknown Constraint Severity"
	}
}
