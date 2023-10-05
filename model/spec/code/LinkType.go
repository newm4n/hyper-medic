package code

import (
	"fmt"
	"strings"
)

type LinkType string

const (
	LinkTypeReplacedBy LinkType = "replaced"
	LinkTypeReplaces   LinkType = "replaces"
	LinkTypeRefer      LinkType = "refer"
	LinkTypeSeeAlso    LinkType = "seealso"
)

func AllLinkType() []LinkType {
	return []LinkType{
		LinkTypeReplacedBy,
		LinkTypeReplaces,
		LinkTypeRefer,
		LinkTypeSeeAlso,
	}
}

func FindLinkType(filter string) []LinkType {
	ret := make([]LinkType, 0)
	for _, at := range AllLinkType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au LinkType) ToString() {
	fmt.Println(au.String())
}
func (au LinkType) String() string {
	switch au {
	case LinkTypeReplacedBy:
		return "Replaced-by"
	case LinkTypeReplaces:
		return "Replaces"
	case LinkTypeRefer:
		return "Refer"
	case LinkTypeSeeAlso:
		return "See also"
	default:
		return "Unknown Link Type"
	}
}
