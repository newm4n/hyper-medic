package code

import (
	"fmt"
	"strings"
)

/* TODO change type to string */
type ArtifactContributionType int

const (
	ArtifactContributionTypeConceptualization ArtifactContributionType = iota
	ArtifactContributionTypeDataCuration
	ArtifactContributionTypeFormalAnalysis
	ArtifactContributionTypeFundingAcquisition
	ArtifactContributionTypeInvestigation
	ArtifactContributionTypeMethodology
	ArtifactContributionTypeProjectAdministration
	ArtifactContributionTypeResources
	ArtifactContributionTypeSoftware
	ArtifactContributionTypeSupervision
	ArtifactContributionTypeValidation
	ArtifactContributionTypeVisualization
	ArtifactContributionTypeWritingOriginalDraft
	ArtifactContributionTypeWritingReviewEditing
)

func AllArtifactContributionType() []ArtifactContributionType {
	return []ArtifactContributionType{
		ArtifactContributionTypeConceptualization,
		ArtifactContributionTypeDataCuration,
		ArtifactContributionTypeFormalAnalysis,
		ArtifactContributionTypeFundingAcquisition,
		ArtifactContributionTypeInvestigation,
		ArtifactContributionTypeMethodology,
		ArtifactContributionTypeProjectAdministration,
		ArtifactContributionTypeResources,
		ArtifactContributionTypeSoftware,
		ArtifactContributionTypeSupervision,
		ArtifactContributionTypeValidation,
		ArtifactContributionTypeVisualization,
		ArtifactContributionTypeWritingOriginalDraft,
		ArtifactContributionTypeWritingReviewEditing,
	}
}

func FindArtifactContributionType(filter string) []ArtifactContributionType {
	ret := make([]ArtifactContributionType, 0)
	for _, at := range AllArtifactContributionType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ArtifactContributionType) ToString() {
	fmt.Println(au.String())
}
func (au ArtifactContributionType) String() string {
	switch au {
	case ArtifactContributionTypeConceptualization:
		return "Conceptualization"
	case ArtifactContributionTypeDataCuration:
		return "Data curation"
	case ArtifactContributionTypeFormalAnalysis:
		return "Formal analysis"
	case ArtifactContributionTypeFundingAcquisition:
		return "Funding acquisition"
	case ArtifactContributionTypeInvestigation:
		return "Investigation"
	case ArtifactContributionTypeMethodology:
		return "Methodology"
	case ArtifactContributionTypeProjectAdministration:
		return "Project administration"
	case ArtifactContributionTypeResources:
		return "Resources"
	case ArtifactContributionTypeSoftware:
		return "Software"
	case ArtifactContributionTypeSupervision:
		return "Supervision"
	case ArtifactContributionTypeValidation:
		return "Validation"
	case ArtifactContributionTypeVisualization:
		return "Visualization"
	case ArtifactContributionTypeWritingOriginalDraft:
		return "Writing - original draft"
	case ArtifactContributionTypeWritingReviewEditing:
		return "Writing - review & editing"
	default:
		return "Unknown Artifact Contribution Type"
	}
}

/*
Conceptualization
Data curation
Formal analysis
Funding acquisition
Investigation
Methodology
Project administration
Resources
Software
Supervision
Validation
Visualization
Writing - original draft
Writing - review & editing

*/
