package code

import (
	"fmt"
	"strings"
)

type ArtifactContributionType string

const (
	ArtifactContributionTypeConceptualization     ArtifactContributionType = "conceptualization"
	ArtifactContributionTypeDataCuration          ArtifactContributionType = "data-curation"
	ArtifactContributionTypeFormalAnalysis        ArtifactContributionType = "formal-analysis"
	ArtifactContributionTypeFundingAcquisition    ArtifactContributionType = "funding-acquisition"
	ArtifactContributionTypeInvestigation         ArtifactContributionType = "investigation"
	ArtifactContributionTypeMethodology           ArtifactContributionType = "methodology"
	ArtifactContributionTypeProjectAdministration ArtifactContributionType = "project-administration"
	ArtifactContributionTypeResources             ArtifactContributionType = "resources"
	ArtifactContributionTypeSoftware              ArtifactContributionType = "software"
	ArtifactContributionTypeSupervision           ArtifactContributionType = "supervision"
	ArtifactContributionTypeValidation            ArtifactContributionType = "validation"
	ArtifactContributionTypeVisualization         ArtifactContributionType = "visualization"
	ArtifactContributionTypeWritingOriginalDraft  ArtifactContributionType = "writing-original-draft"
	ArtifactContributionTypeWritingReviewEditing  ArtifactContributionType = "writing-review-editing"
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
