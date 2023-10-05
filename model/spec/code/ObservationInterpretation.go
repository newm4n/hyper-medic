package code

import (
	"fmt"
	"strings"
)

type ObservationInterpretation string

const (
	ObservationInterpretationGeneticObservationInterpretation        ObservationInterpretation = "GeneticObservationInterpretation"
	ObservationInterpretationCAR                                     ObservationInterpretation = "CAR"
	ObservationInterpretationCarrier                                 ObservationInterpretation = "Carrier"
	ObservationInterpretationObservationInterpretationChange         ObservationInterpretation = "ObservationInterpretationChange"
	ObservationInterpretationB                                       ObservationInterpretation = "B"
	ObservationInterpretationD                                       ObservationInterpretation = "D"
	ObservationInterpretationU                                       ObservationInterpretation = "U"
	ObservationInterpretationW                                       ObservationInterpretation = "W"
	ObservationInterpretationObservationInterpretationExceptions     ObservationInterpretation = "ObservationInterpretationExceptions"
	ObservationInterpretationSymbLT                                  ObservationInterpretation = "<"
	ObservationInterpretationSymbGT                                  ObservationInterpretation = ">"
	ObservationInterpretationAC                                      ObservationInterpretation = "AC"
	ObservationInterpretationIE                                      ObservationInterpretation = "IE"
	ObservationInterpretationQCF                                     ObservationInterpretation = "QCF"
	ObservationInterpretationTOX                                     ObservationInterpretation = "TOX"
	ObservationInterpretationObservationInterpretationNormality      ObservationInterpretation = "ObservationInterpretationNormality"
	ObservationInterpretationA                                       ObservationInterpretation = "A"
	ObservationInterpretationAA                                      ObservationInterpretation = "AA"
	ObservationInterpretationHH                                      ObservationInterpretation = "HH"
	ObservationInterpretationLL                                      ObservationInterpretation = "LL"
	ObservationInterpretationH                                       ObservationInterpretation = "H"
	ObservationInterpretationHSymbGT                                 ObservationInterpretation = "H>"
	ObservationInterpretationHU                                      ObservationInterpretation = "HU"
	ObservationInterpretationL                                       ObservationInterpretation = "L"
	ObservationInterpretationLSymbLT                                 ObservationInterpretation = "L<"
	ObservationInterpretationLU                                      ObservationInterpretation = "LU"
	ObservationInterpretationN                                       ObservationInterpretation = "N"
	ObservationInterpretationObservationInterpretationSusceptibility ObservationInterpretation = "ObservationInterpretationSusceptibility"
	ObservationInterpretationI                                       ObservationInterpretation = "I"
	ObservationInterpretationMS                                      ObservationInterpretation = "MS"
	ObservationInterpretationNCL                                     ObservationInterpretation = "NCL"
	ObservationInterpretationNS                                      ObservationInterpretation = "NS"
	ObservationInterpretationR                                       ObservationInterpretation = "R"
	ObservationInterpretationSYNdR                                   ObservationInterpretation = "SYN-R"
	ObservationInterpretationS                                       ObservationInterpretation = "S"
	ObservationInterpretationSDD                                     ObservationInterpretation = "SDD"
	ObservationInterpretationSYNdS                                   ObservationInterpretation = "SYN-S"
	ObservationInterpretationVS                                      ObservationInterpretation = "VS"
	ObservationInterpretationEX                                      ObservationInterpretation = "EX"
	ObservationInterpretationHX                                      ObservationInterpretation = "HX"
	ObservationInterpretationLX                                      ObservationInterpretation = "LX"
	ObservationInterpretationHM                                      ObservationInterpretation = "HM"
	ObservationInterpretationObservationInterpretationDetection      ObservationInterpretation = "ObservationInterpretationDetection"
	ObservationInterpretationIND                                     ObservationInterpretation = "IND"
	ObservationInterpretationE                                       ObservationInterpretation = "E"
	ObservationInterpretationNEG                                     ObservationInterpretation = "NEG"
	ObservationInterpretationND                                      ObservationInterpretation = "ND"
	ObservationInterpretationPOS                                     ObservationInterpretation = "POS"
	ObservationInterpretationDET                                     ObservationInterpretation = "DET"
	ObservationInterpretationObservationInterpretationExpectation    ObservationInterpretation = "ObservationInterpretationExpectation"
	ObservationInterpretationEXP                                     ObservationInterpretation = "EXP"
	ObservationInterpretationUNE                                     ObservationInterpretation = "UNE"
	ObservationInterpretationOBX                                     ObservationInterpretation = "OBX"
	ObservationInterpretationReactivityObservationInterpretation     ObservationInterpretation = "ReactivityObservationInterpretation"
	ObservationInterpretationNR                                      ObservationInterpretation = "NR"
	ObservationInterpretationRR                                      ObservationInterpretation = "RR"
	ObservationInterpretationWR                                      ObservationInterpretation = "WR"
)

func AllObservationInterpretation() []ObservationInterpretation {
	return []ObservationInterpretation{
		ObservationInterpretationGeneticObservationInterpretation,
		ObservationInterpretationCAR,
		ObservationInterpretationCarrier,
		ObservationInterpretationObservationInterpretationChange,
		ObservationInterpretationB,
		ObservationInterpretationD,
		ObservationInterpretationU,
		ObservationInterpretationW,
		ObservationInterpretationObservationInterpretationExceptions,
		ObservationInterpretationSymbLT,
		ObservationInterpretationSymbGT,
		ObservationInterpretationAC,
		ObservationInterpretationIE,
		ObservationInterpretationQCF,
		ObservationInterpretationTOX,
		ObservationInterpretationObservationInterpretationNormality,
		ObservationInterpretationA,
		ObservationInterpretationAA,
		ObservationInterpretationHH,
		ObservationInterpretationLL,
		ObservationInterpretationH,
		ObservationInterpretationHSymbGT,
		ObservationInterpretationHU,
		ObservationInterpretationL,
		ObservationInterpretationLSymbLT,
		ObservationInterpretationLU,
		ObservationInterpretationN,
		ObservationInterpretationObservationInterpretationSusceptibility,
		ObservationInterpretationI,
		ObservationInterpretationMS,
		ObservationInterpretationNCL,
		ObservationInterpretationNS,
		ObservationInterpretationR,
		ObservationInterpretationSYNdR,
		ObservationInterpretationS,
		ObservationInterpretationSDD,
		ObservationInterpretationSYNdS,
		ObservationInterpretationVS,
		ObservationInterpretationEX,
		ObservationInterpretationHX,
		ObservationInterpretationLX,
		ObservationInterpretationHM,
		ObservationInterpretationObservationInterpretationDetection,
		ObservationInterpretationIND,
		ObservationInterpretationE,
		ObservationInterpretationNEG,
		ObservationInterpretationND,
		ObservationInterpretationPOS,
		ObservationInterpretationDET,
		ObservationInterpretationObservationInterpretationExpectation,
		ObservationInterpretationEXP,
		ObservationInterpretationUNE,
		ObservationInterpretationOBX,
		ObservationInterpretationReactivityObservationInterpretation,
		ObservationInterpretationNR,
		ObservationInterpretationRR,
		ObservationInterpretationWR,
	}
}

func FindObservationInterpretation(filter string) []ObservationInterpretation {
	ret := make([]ObservationInterpretation, 0)
	for _, at := range AllObservationInterpretation() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au ObservationInterpretation) ToString() {
	fmt.Println(au.String())
}
func (au ObservationInterpretation) String() string {
	switch au {
	case ObservationInterpretationGeneticObservationInterpretation:
		return "GeneticObservationInterpretation"
	case ObservationInterpretationCAR:
		return "Carrier"
	case ObservationInterpretationCarrier:
		return "Carrier"
	case ObservationInterpretationObservationInterpretationChange:
		return "ObservationInterpretationChange"
	case ObservationInterpretationB:
		return "Better"
	case ObservationInterpretationD:
		return "Significant change down"
	case ObservationInterpretationU:
		return "Significant change up"
	case ObservationInterpretationW:
		return "Worse"
	case ObservationInterpretationObservationInterpretationExceptions:
		return "ObservationInterpretationExceptions"
	case ObservationInterpretationSymbLT:
		return "Off scale low"
	case ObservationInterpretationSymbGT:
		return "Off scale high"
	case ObservationInterpretationAC:
		return "Anti-complementary substances present"
	case ObservationInterpretationIE:
		return "Insufficient evidence"
	case ObservationInterpretationQCF:
		return "Quality control failure"
	case ObservationInterpretationTOX:
		return "Cytotoxic substance present"
	case ObservationInterpretationObservationInterpretationNormality:
		return "ObservationInterpretationNormality"
	case ObservationInterpretationA:
		return "Abnormal"
	case ObservationInterpretationAA:
		return "Critical abnormal"
	case ObservationInterpretationHH:
		return "Critical high"
	case ObservationInterpretationLL:
		return "Critical low"
	case ObservationInterpretationH:
		return "High"
	case ObservationInterpretationHSymbGT:
		return "Significantly high"
	case ObservationInterpretationHU:
		return "Significantly high"
	case ObservationInterpretationL:
		return "Low"
	case ObservationInterpretationLSymbLT:
		return "Significantly low"
	case ObservationInterpretationLU:
		return "Significantly low"
	case ObservationInterpretationN:
		return "Normal"
	case ObservationInterpretationObservationInterpretationSusceptibility:
		return "ObservationInterpretationSusceptibility"
	case ObservationInterpretationI:
		return "Intermediate"
	case ObservationInterpretationMS:
		return "moderately susceptible"
	case ObservationInterpretationNCL:
		return "No CLSI defined breakpoint"
	case ObservationInterpretationNS:
		return "Non-susceptible"
	case ObservationInterpretationR:
		return "Resistant"
	case ObservationInterpretationSYNdR:
		return "Synergy - resistant"
	case ObservationInterpretationS:
		return "Susceptible"
	case ObservationInterpretationSDD:
		return "Susceptible-dose dependent"
	case ObservationInterpretationSYNdS:
		return "Synergy - susceptible"
	case ObservationInterpretationVS:
		return "very susceptible"
	case ObservationInterpretationEX:
		return "outside threshold"
	case ObservationInterpretationHX:
		return "above high threshold"
	case ObservationInterpretationLX:
		return "below low threshold"
	case ObservationInterpretationHM:
		return "Hold for Medical Review"
	case ObservationInterpretationObservationInterpretationDetection:
		return "ObservationInterpretationDetection"
	case ObservationInterpretationIND:
		return "Indeterminate"
	case ObservationInterpretationE:
		return "Equivocal"
	case ObservationInterpretationNEG:
		return "Negative"
	case ObservationInterpretationND:
		return "Not detected"
	case ObservationInterpretationPOS:
		return "Positive"
	case ObservationInterpretationDET:
		return "Detected"
	case ObservationInterpretationObservationInterpretationExpectation:
		return "ObservationInterpretationExpectation"
	case ObservationInterpretationEXP:
		return "Expected"
	case ObservationInterpretationUNE:
		return "Unexpected"
	case ObservationInterpretationOBX:
		return "Interpretation qualifiers in separate OBX segments"
	case ObservationInterpretationReactivityObservationInterpretation:
		return "ReactivityObservationInterpretation"
	case ObservationInterpretationNR:
		return "Non-reactive"
	case ObservationInterpretationRR:
		return "Reactive"
	case ObservationInterpretationWR:
		return "Weakly reactive"

	default:
		return "Unknown Constraint Severity"
	}
}
