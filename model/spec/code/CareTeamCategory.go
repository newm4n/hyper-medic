package code

import (
	"fmt"
	"strings"
)

type CareTeamCategory int

const (
	CareTeamCategoryEventFocused CareTeamCategory = iota
	CareTeamCategoryEncounterFocused
	CareTeamCategoryEpisodeOfCareFocused
	CareTeamCategoryConditionFocused
	CareTeamCategoryLongitudinalCareCoordinationFocused
	CareTeamCategoryHomeCommunityBasedServicesFocused
	CareTeamCategoryClinicalResearchFocused
	CareTeamCategoryPublicHealthFocused
)

func AllCareTeamCategory() []CareTeamCategory {
	return []CareTeamCategory{
		CareTeamCategoryEventFocused,
		CareTeamCategoryEncounterFocused,
		CareTeamCategoryEpisodeOfCareFocused,
		CareTeamCategoryConditionFocused,
		CareTeamCategoryLongitudinalCareCoordinationFocused,
		CareTeamCategoryHomeCommunityBasedServicesFocused,
		CareTeamCategoryClinicalResearchFocused,
		CareTeamCategoryPublicHealthFocused,
	}
}

func FindCareTeamCategory(filter string) []CareTeamCategory {
	ret := make([]CareTeamCategory, 0)
	for _, at := range AllCareTeamCategory() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt CareTeamCategory) ToString() {
	fmt.Println(cpt.String())
}

func (cpt CareTeamCategory) String() string {
	switch cpt {
	case CareTeamCategoryEventFocused:
		return "\nEvent-focused care team"
	case CareTeamCategoryEncounterFocused:
		return "Encounter-focused care team"
	case CareTeamCategoryEpisodeOfCareFocused:
		return "Episode of care-focused care team"
	case CareTeamCategoryConditionFocused:
		return "Condition-focused care team"
	case CareTeamCategoryLongitudinalCareCoordinationFocused:
		return "Longitudinal care-coordination focused care team"
	case CareTeamCategoryHomeCommunityBasedServicesFocused:
		return "Home & Community Based Services (HCBS)-focused care team"
	case CareTeamCategoryClinicalResearchFocused:
		return "Clinical research-focused care team"
	case CareTeamCategoryPublicHealthFocused:
		return "Public health-focused care team"
	default:
		return "Unknown CareTeamCategory"
	}
}

/**
Event-focused care team
	Encounter-focused care team
Episode of care-focused care team
Condition-focused care team
	Longitudinal care-coordination focused care team
Home & Community Based Services (HCBS)-focused care team
Clinical research-focused care team
Public health-focused care team
*/
