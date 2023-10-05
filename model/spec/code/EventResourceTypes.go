package code

import (
	"fmt"
	"strings"
)

type EventResourceTypes string

const (
	EventResourceTypesSender   EventResourceTypes = "sender"
	EventResourceTypesReceiver EventResourceTypes = "receiver"
)

func AllEventResourceTypes() []EventResourceTypes {
	return []EventResourceTypes{
		EventResourceTypesSender,
		EventResourceTypesReceiver,
	}
}

func FindEventResourceTypes(filter string) []EventResourceTypes {
	ret := make([]EventResourceTypes, 0)
	for _, at := range AllEventResourceTypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au EventResourceTypes) ToString() {
	fmt.Println(au.String())
}
func (au EventResourceTypes) String() string {
	switch au {
	case EventResourceTypesSender:
		return "Sender"
	case EventResourceTypesReceiver:
		return "Receiver"
	default:
		return "Unknown Event Resource Types"
	}
}
