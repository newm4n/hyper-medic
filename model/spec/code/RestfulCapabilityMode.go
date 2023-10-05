package code

import (
	"fmt"
	"strings"
)

type RestfulCapabilityMode string

const (
	RestfulCapabilityModeClient RestfulCapabilityMode = "client"
	RestfulCapabilityModeServer RestfulCapabilityMode = "server"
)

func AllRestfulCapabilityMode() []RestfulCapabilityMode {
	return []RestfulCapabilityMode{
		RestfulCapabilityModeClient,
		RestfulCapabilityModeServer,
	}
}

func FindRestfulCapabilityMode(filter string) []RestfulCapabilityMode {
	ret := make([]RestfulCapabilityMode, 0)
	for _, at := range AllRestfulCapabilityMode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au RestfulCapabilityMode) ToString() {
	fmt.Println(au.String())
}
func (au RestfulCapabilityMode) String() string {
	switch au {
	case RestfulCapabilityModeClient:
		return "Client"
	case RestfulCapabilityModeServer:
		return "Server"
	default:
		return "Unknown Restful Capability Mode"
	}
}
