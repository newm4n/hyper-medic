package code

import (
	"fmt"
	"strings"
)

type CodesystemHierarchyMeaning string

const (
	CodesystemHierarchyMeaningGroupedBy      CodesystemHierarchyMeaning = "grouped-by"
	CodesystemHierarchyMeaningIsA            CodesystemHierarchyMeaning = "is-a"
	CodesystemHierarchyMeaningPartOf         CodesystemHierarchyMeaning = "part-of"
	CodesystemHierarchyMeaningClassifiedWith CodesystemHierarchyMeaning = "classified-with"
)

func AllCodesystemHierarchyMeaning() []CodesystemHierarchyMeaning {
	return []CodesystemHierarchyMeaning{
		CodesystemHierarchyMeaningGroupedBy,
		CodesystemHierarchyMeaningIsA,
		CodesystemHierarchyMeaningPartOf,
		CodesystemHierarchyMeaningClassifiedWith,
	}
}

func FindCodesystemHierarchyMeaning(filter string) []CodesystemHierarchyMeaning {
	ret := make([]CodesystemHierarchyMeaning, 0)
	for _, at := range AllCodesystemHierarchyMeaning() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au CodesystemHierarchyMeaning) ToString() {
	fmt.Println(au.String())
}
func (au CodesystemHierarchyMeaning) String() string {
	switch au {
	case CodesystemHierarchyMeaningGroupedBy:
		return "Grouped By"
	case CodesystemHierarchyMeaningIsA:
		return "Is-A"
	case CodesystemHierarchyMeaningPartOf:
		return "Part Of"
	case CodesystemHierarchyMeaningClassifiedWith:
		return "Classified With"
	default:
		return "Unknown Address Use"
	}
}
