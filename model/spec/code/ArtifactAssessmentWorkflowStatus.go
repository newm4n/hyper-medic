package code

import (
	"fmt"
	"strings"
)

// https://www.hl7.org/fhir/valueset-artifactassessment-workflow-status.html

type ArtifactAssessmentWorkflowStatus string

const (
	ArtifactAssessmentWorkflowStatusSubmitted  ArtifactAssessmentWorkflowStatus = "submitted"
	ArtifactAssessmentWorkflowStatusTriaged    ArtifactAssessmentWorkflowStatus = "triaged"
	ArtifactAssessmentWorkflowStatusWaiting    ArtifactAssessmentWorkflowStatus = "waiting-for-input"
	ArtifactAssessmentWorkflowStatusResolvedNC ArtifactAssessmentWorkflowStatus = "resolved-no-change"
	ArtifactAssessmentWorkflowStatusResolvedCR ArtifactAssessmentWorkflowStatus = "resolved-change-required"
	ArtifactAssessmentWorkflowStatusDeferred   ArtifactAssessmentWorkflowStatus = "deferred"
	ArtifactAssessmentWorkflowStatusDuplicate  ArtifactAssessmentWorkflowStatus = "duplicate"
	ArtifactAssessmentWorkflowStatusApplied    ArtifactAssessmentWorkflowStatus = "applied"
	ArtifactAssessmentWorkflowStatusPublished  ArtifactAssessmentWorkflowStatus = "published"
	ArtifactAssessmentWorkflowStatusError      ArtifactAssessmentWorkflowStatus = "entered-in-error"
)

func AllArtifactAssessmentWorkflowStatus() []ArtifactAssessmentWorkflowStatus {
	return []ArtifactAssessmentWorkflowStatus{
		ArtifactAssessmentWorkflowStatusSubmitted,
		ArtifactAssessmentWorkflowStatusTriaged,
		ArtifactAssessmentWorkflowStatusWaiting,
		ArtifactAssessmentWorkflowStatusResolvedNC,
		ArtifactAssessmentWorkflowStatusResolvedCR,
		ArtifactAssessmentWorkflowStatusDeferred,
		ArtifactAssessmentWorkflowStatusDuplicate,
		ArtifactAssessmentWorkflowStatusApplied,
		ArtifactAssessmentWorkflowStatusPublished,
		ArtifactAssessmentWorkflowStatusError,
	}
}

func FindArtifactAssessmentWorkflowStatus(filter string) []ArtifactAssessmentWorkflowStatus {
	ret := make([]ArtifactAssessmentWorkflowStatus, 0)
	for _, at := range AllArtifactAssessmentWorkflowStatus() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ArtifactAssessmentWorkflowStatus) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ArtifactAssessmentWorkflowStatus) String() string {
	switch cpt {
	case ArtifactAssessmentWorkflowStatusSubmitted:
		return "Submitted"
	case ArtifactAssessmentWorkflowStatusTriaged:
		return "Triaged"
	case ArtifactAssessmentWorkflowStatusWaiting:
		return "Waiting for Input"
	case ArtifactAssessmentWorkflowStatusResolvedNC:
		return "Resolved - No Change"
	case ArtifactAssessmentWorkflowStatusResolvedCR:
		return "Resolved - Change Required"
	case ArtifactAssessmentWorkflowStatusDeferred:
		return "Deferred"
	case ArtifactAssessmentWorkflowStatusDuplicate:
		return "Duplicate"
	case ArtifactAssessmentWorkflowStatusApplied:
		return "Applied"
	case ArtifactAssessmentWorkflowStatusPublished:
		return "Published"
	case ArtifactAssessmentWorkflowStatusError:
		return "Entered in Error"
	default:
		return "Unknown Artifact Assessment Workflow Status"
	}
}
