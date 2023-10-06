package code

import (
	"fmt"
	"strings"
)

type AttributeEstimateType string

const (
	AttributeEstimateTypeCochransQstatistic     AttributeEstimateType = "0000419"
	AttributeEstimateTypeConfidenceInterval     AttributeEstimateType = "C53324"
	AttributeEstimateTypeCredibleInterval       AttributeEstimateType = "0000455"
	AttributeEstimateTypeIsquared               AttributeEstimateType = "0000420"
	AttributeEstimateTypeInterquartileRange     AttributeEstimateType = "C53245"
	AttributeEstimateTypePValue                 AttributeEstimateType = "C44185"
	AttributeEstimateTypeRange                  AttributeEstimateType = "C38013"
	AttributeEstimateTypeStandardDeviation      AttributeEstimateType = "C53322"
	AttributeEstimateTypeStandardErrorOfTheMean AttributeEstimateType = "0000037 "
	AttributeEstimateTypeTauSquared             AttributeEstimateType = "0000421"
	AttributeEstimateTypeVariance               AttributeEstimateType = "C48918"
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
