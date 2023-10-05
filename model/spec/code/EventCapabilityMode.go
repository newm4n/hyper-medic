package code

import (
	"fmt"
	"strings"
)

type EventCapabilityMode string

const (
	EventCapabilityModeSender   EventCapabilityMode = "sender"
	EventCapabilityModeReceiver EventCapabilityMode = "receiver"
)

func AllEventCapabilityMode() []EventCapabilityMode {
	return []EventCapabilityMode{
		EventCapabilityModeSender,
		EventCapabilityModeSender,
	}
}

func FindEventCapabilityMode(filter string) []EventCapabilityMode {
	ret := make([]EventCapabilityMode, 0)
	for _, at := range AllEventCapabilityMode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au EventCapabilityMode) ToString() {
	fmt.Println(au.String())
}
func (au EventCapabilityMode) String() string {
	switch au {
	case EventCapabilityModeSender:
		return "Sender"
	case EventCapabilityModeReceiver:
		return "Receiver"
	default:
		return "Unknown Constraint Severity"
	}
}
