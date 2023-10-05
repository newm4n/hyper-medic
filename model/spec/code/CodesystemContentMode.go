package code

import (
	"fmt"
	"strings"
)

type CodesystemContentMode string

const (
	CodesystemContentModeNotPresent CodesystemContentMode = "not-present"
	CodesystemContentModeExample    CodesystemContentMode = "example"
	CodesystemContentModeFragment   CodesystemContentMode = "fragment"
	CodesystemContentModeComplete   CodesystemContentMode = "complete"
	CodesystemContentModeSupplement CodesystemContentMode = "supplement"
)

func AllCodesystemContentMode() []CodesystemContentMode {
	return []CodesystemContentMode{
		CodesystemContentModeNotPresent,
		CodesystemContentModeExample,
		CodesystemContentModeFragment,
		CodesystemContentModeComplete,
		CodesystemContentModeSupplement,
	}
}

func FindCodesystemContentMode(filter string) []CodesystemContentMode {
	ret := make([]CodesystemContentMode, 0)
	for _, at := range AllCodesystemContentMode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au CodesystemContentMode) ToString() {
	fmt.Println(au.String())
}
func (au CodesystemContentMode) String() string {
	switch au {
	case CodesystemContentModeNotPresent:
		return "Not Present"
	case CodesystemContentModeExample:
		return "Example"
	case CodesystemContentModeFragment:
		return "Fragment"
	case CodesystemContentModeComplete:
		return "Complete"
	case CodesystemContentModeSupplement:
		return "Supplement"
	default:
		return "Unknown Codesystem Content Mode"
	}
}
