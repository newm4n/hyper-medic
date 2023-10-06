package code

import (
	"fmt"
	"strings"
)

type ArtifactAssessmentInformationType string

const (
	ArtifactAssessmentInformationTypeComment       ArtifactAssessmentInformationType = "comment"
	ArtifactAssessmentInformationTypeClassifier    ArtifactAssessmentInformationType = "classifier"
	ArtifactAssessmentInformationTypeRating        ArtifactAssessmentInformationType = "rating"
	ArtifactAssessmentInformationTypeContainer     ArtifactAssessmentInformationType = "container"
	ArtifactAssessmentInformationTypeResponse      ArtifactAssessmentInformationType = "response"
	ArtifactAssessmentInformationTypeChangeRequest ArtifactAssessmentInformationType = "change-request"
)

func AllArtifactAssessmentInformationType() []ArtifactAssessmentInformationType {
	return []ArtifactAssessmentInformationType{
		ArtifactAssessmentInformationTypeComment,
		ArtifactAssessmentInformationTypeClassifier,
		ArtifactAssessmentInformationTypeRating,
		ArtifactAssessmentInformationTypeContainer,
		ArtifactAssessmentInformationTypeResponse,
		ArtifactAssessmentInformationTypeChangeRequest,
	}
}

func FindArtifactAssessmentInformationType(filter string) []ArtifactAssessmentInformationType {
	ret := make([]ArtifactAssessmentInformationType, 0)
	for _, at := range AllArtifactAssessmentInformationType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ArtifactAssessmentInformationType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ArtifactAssessmentInformationType) String() string {
	switch cpt {
	case ArtifactAssessmentInformationTypeComment:
		return "CareTeam"
	case ArtifactAssessmentInformationTypeClassifier:
		return "Device"
	case ArtifactAssessmentInformationTypeRating:
		return "Group"
	case ArtifactAssessmentInformationTypeContainer:
		return "HealthcareService"
	case ArtifactAssessmentInformationTypeResponse:
		return "Location"
	case ArtifactAssessmentInformationTypeChangeRequest:
		return "Organization"

	default:
		return "Unknown Action Participant Type"
	}
}

/*
Comment
Classifier
Rating
Container
Response
Change Request
*/
