package code

import (
	"fmt"
	"strings"
)

type ArtifactURLClassifier string

const (
	ArtifactURLClassifierAbstract           ArtifactURLClassifier = "abstract"
	ArtifactURLClassifierFullText           ArtifactURLClassifier = "full-text"
	ArtifactURLClassifierSupplement         ArtifactURLClassifier = "supplement"
	ArtifactURLClassifierWebpage            ArtifactURLClassifier = "webpage"
	ArtifactURLClassifierFileDirectory      ArtifactURLClassifier = "file-directory"
	ArtifactURLClassifierCodeRepository     ArtifactURLClassifier = "code-repository"
	ArtifactURLClassifierRestricted         ArtifactURLClassifier = "restricted"
	ArtifactURLClassifierCompressedFile     ArtifactURLClassifier = "compressed-file"
	ArtifactURLClassifierDOIBased           ArtifactURLClassifier = "doi-based"
	ArtifactURLClassifierPDF                ArtifactURLClassifier = "pdf"
	ArtifactURLClassifierJSON               ArtifactURLClassifier = "json"
	ArtifactURLClassifierXML                ArtifactURLClassifier = "xml"
	ArtifactURLClassifierVersionSpecific    ArtifactURLClassifier = "version-specific"
	ArtifactURLClassifierComputableResource ArtifactURLClassifier = "computable-resource"
	ArtifactURLClassifierNotSpecified       ArtifactURLClassifier = "not-specified"
)

func AllArtifactURLClassifier() []ArtifactURLClassifier {
	return []ArtifactURLClassifier{
		ArtifactURLClassifierAbstract,
		ArtifactURLClassifierFullText,
		ArtifactURLClassifierSupplement,
		ArtifactURLClassifierWebpage,
		ArtifactURLClassifierFileDirectory,
		ArtifactURLClassifierCodeRepository,
		ArtifactURLClassifierRestricted,
		ArtifactURLClassifierCompressedFile,
		ArtifactURLClassifierDOIBased,
		ArtifactURLClassifierPDF,
		ArtifactURLClassifierJSON,
		ArtifactURLClassifierXML,
		ArtifactURLClassifierVersionSpecific,
		ArtifactURLClassifierComputableResource,
		ArtifactURLClassifierNotSpecified,
	}
}

func FindArtifactURLClassifier(filter string) []ArtifactURLClassifier {
	ret := make([]ArtifactURLClassifier, 0)
	for _, at := range AllArtifactURLClassifier() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ArtifactURLClassifier) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ArtifactURLClassifier) String() string {
	switch cpt {
	case ArtifactURLClassifierAbstract:
		return "Abstract"
	case ArtifactURLClassifierFullText:
		return "Full-Text"
	case ArtifactURLClassifierSupplement:
		return "Supplement"
	case ArtifactURLClassifierWebpage:
		return "Webpage"
	case ArtifactURLClassifierFileDirectory:
		return "File directory"
	case ArtifactURLClassifierCodeRepository:
		return "Code repository"
	case ArtifactURLClassifierRestricted:
		return "Restricted"
	case ArtifactURLClassifierCompressedFile:
		return "Compressed file"
	case ArtifactURLClassifierDOIBased:
		return "DOI Based"
	case ArtifactURLClassifierPDF:
		return "PDF"
	case ArtifactURLClassifierJSON:
		return "JSON"
	case ArtifactURLClassifierXML:
		return "XML"
	case ArtifactURLClassifierVersionSpecific:
		return "Version Specific"
	case ArtifactURLClassifierComputableResource:
		return "Computable resource"
	case ArtifactURLClassifierNotSpecified:
		return "Not Specified"
	default:
		return "Unknown Artifact URL Classifier"
	}
}
