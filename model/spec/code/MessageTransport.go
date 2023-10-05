package code

import (
	"fmt"
	"strings"
)

type MessageTransport string

const (
	MessageTransportHTTP MessageTransport = "http"
	MessageTransportFTP  MessageTransport = "ftp"
	MessageTransportMLLP MessageTransport = "mllp"
)

func AllMessageTransport() []MessageTransport {
	return []MessageTransport{
		MessageTransportHTTP,
		MessageTransportFTP,
		MessageTransportMLLP,
	}
}

func FindMessageTransport(filter string) []MessageTransport {
	ret := make([]MessageTransport, 0)
	for _, at := range AllMessageTransport() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au MessageTransport) ToString() {
	fmt.Println(au.String())
}
func (au MessageTransport) String() string {
	switch au {
	case MessageTransportHTTP:
		return "HTTP"
	case MessageTransportFTP:
		return "FTP"
	case MessageTransportMLLP:
		return "MLLP"
	default:
		return "Unknown Constraint Severity"
	}
}
