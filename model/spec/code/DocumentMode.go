package code

import (
	"fmt"
	"strings"
)

type DocumentMode string

const (
	DocumentModeProducer DocumentMode = "producer"
	DocumentModeConsumer DocumentMode = "consumer"
)

func AllDocumentMode() []DocumentMode {
	return []DocumentMode{
		DocumentModeProducer,
		DocumentModeConsumer,
	}
}

func FindDocumentMode(filter string) []DocumentMode {
	ret := make([]DocumentMode, 0)
	for _, at := range AllDocumentMode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DocumentMode) ToString() {
	fmt.Println(au.String())
}
func (au DocumentMode) String() string {
	switch au {
	case DocumentModeProducer:
		return "Producer"
	case DocumentModeConsumer:
		return "Consumer"
	default:
		return "Unknown Document Mode"
	}
}
