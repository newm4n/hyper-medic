package code

import (
	"fmt"
	"strings"
)

type ArtifactAssessmentDisposition string

const (
	ArtifactAssessmentDispositionUnresolved        ArtifactAssessmentDisposition = "unresolved"
	ArtifactAssessmentDispositionNotPersuasive     ArtifactAssessmentDisposition = "not-persuasive"
	ArtifactAssessmentDispositionPersuasive        ArtifactAssessmentDisposition = "persuasive"
	ArtifactAssessmentDispositionPersuasiveWMod    ArtifactAssessmentDisposition = "persuasive-with-modification"
	ArtifactAssessmentDispositionNotPersuasiveWMod ArtifactAssessmentDisposition = "not-persuasive-with-modification"
)

func AllArtifactAssessmentDisposition() []ArtifactAssessmentDisposition {
	return []ArtifactAssessmentDisposition{
		ArtifactAssessmentDispositionUnresolved,
		ArtifactAssessmentDispositionNotPersuasive,
		ArtifactAssessmentDispositionPersuasive,
		ArtifactAssessmentDispositionPersuasiveWMod,
		ArtifactAssessmentDispositionNotPersuasiveWMod,
	}
}

func FindArtifactAssessmentDisposition(filter string) []ArtifactAssessmentDisposition {
	ret := make([]ArtifactAssessmentDisposition, 0)
	for _, at := range AllArtifactAssessmentDisposition() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ArtifactAssessmentDisposition) ToString() {
	fmt.Println(au.String())
}
func (au ArtifactAssessmentDisposition) String() string {
	switch au {
	case ArtifactAssessmentDispositionUnresolved:
		return "Unresolved"
	case ArtifactAssessmentDispositionNotPersuasive:
		return "Not Persuasive"
	case ArtifactAssessmentDispositionPersuasive:
		return "Persuasive"
	case ArtifactAssessmentDispositionPersuasiveWMod:
		return "Persuasive with Modification"
	case ArtifactAssessmentDispositionNotPersuasiveWMod:
		return "Not Persuasive with Modification"
	default:
		return "Unknown Artifact Assessment Disposition"
	}
}

/*
Unresolved
Not Persuasive
Persuasive
Persuasive with Modification
Not Persuasive with Modification
*/
