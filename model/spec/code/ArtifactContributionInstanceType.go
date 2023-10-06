package code

import (
	"fmt"
	"strings"
)

type ArtifactContributionInstanceType string

const (
	ArtifactContributionInstanceTypeReviewed ArtifactContributionInstanceType = "reviewed"
	ArtifactContributionInstanceTypeApproved ArtifactContributionInstanceType = "approved"
	ArtifactContributionInstanceTypeEdited   ArtifactContributionInstanceType = "edited"
)

func AllArtifactContributionInstanceType() []ArtifactContributionInstanceType {
	return []ArtifactContributionInstanceType{
		ArtifactContributionInstanceTypeReviewed,
		ArtifactContributionInstanceTypeApproved,
		ArtifactContributionInstanceTypeEdited,
	}
}

func FindArtifactContributionInstanceType(filter string) []ArtifactContributionInstanceType {
	ret := make([]ArtifactContributionInstanceType, 0)
	for _, at := range AllArtifactContributionInstanceType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ArtifactContributionInstanceType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ArtifactContributionInstanceType) String() string {
	switch cpt {
	case ArtifactContributionInstanceTypeReviewed:
		return "Reviewed"
	case ArtifactContributionInstanceTypeApproved:
		return "Approved"
	case ArtifactContributionInstanceTypeEdited:
		return "Edited"
	default:
		return "Unknown Artifact Contribution Instance Type"
	}
}
