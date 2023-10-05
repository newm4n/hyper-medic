package code

import (
	"fmt"
	"strings"
)

/* TODO change type to string */
type AttributeEstimateType int

const (
	AttributeEstimateTypeCochransQstatistic AttributeEstimateType = iota
	AttributeEstimateTypeConfidenceInterval
	AttributeEstimateTypeCredibleInterval
	AttributeEstimateTypeIsquared
	AttributeEstimateTypeInterquartileRange
	AttributeEstimateTypePValue
	AttributeEstimateTypeRange
	AttributeEstimateTypeStandardDeviation
	AttributeEstimateTypeStandardErrorOfTheMean
	AttributeEstimateTypeTauSquared
	AttributeEstimateTypeVariance
)

func AllAttributeEstimateType() []AttributeEstimateType {
	return []AttributeEstimateType{
		AttributeEstimateTypeCochransQstatistic,
		AttributeEstimateTypeConfidenceInterval,
		AttributeEstimateTypeCredibleInterval,
		AttributeEstimateTypeIsquared,
		AttributeEstimateTypeInterquartileRange,
		AttributeEstimateTypePValue,
		AttributeEstimateTypeRange,
		AttributeEstimateTypeStandardDeviation,
		AttributeEstimateTypeStandardErrorOfTheMean,
		AttributeEstimateTypeTauSquared,
		AttributeEstimateTypeVariance,
	}
}

func FindAttributeEstimateType(filter string) []AttributeEstimateType {
	ret := make([]AttributeEstimateType, 0)
	for _, at := range AllAttributeEstimateType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au AttributeEstimateType) ToString() {
	fmt.Println(au.String())
}
func (au AttributeEstimateType) String() string {
	switch au {
	case AttributeEstimateTypeCochransQstatistic:
		return "Cochran's Q statistic"
	case AttributeEstimateTypeConfidenceInterval:
		return "Confidence interval"
	case AttributeEstimateTypeCredibleInterval:
		return "Credible interval"
	case AttributeEstimateTypeIsquared:
		return "I-squared"
	case AttributeEstimateTypeInterquartileRange:
		return "Interquartile range"
	case AttributeEstimateTypePValue:
		return "P-value"
	case AttributeEstimateTypeRange:
		return "Range"
	case AttributeEstimateTypeStandardDeviation:
		return "Standard deviation"
	case AttributeEstimateTypeStandardErrorOfTheMean:
		return "Standard error of the mean"
	case AttributeEstimateTypeTauSquared:
		return "Tau squared"
	case AttributeEstimateTypeVariance:
		return "Variance"
	default:
		return "Unknown Address Use"
	}
}

/*
Cochran's Q statistic
Confidence interval
Credible interval
I-squared
Interquartile range
P-value
Range
Standard deviation
Standard error of the mean
Tau squared
Variance

*/
