package code

import (
	"fmt"
	"strings"
)

type DataAbsentReason string

const (
	DataAbsentReasonUnknown          DataAbsentReason = "unknown"
	DataAbsentReasonAskedUnknown     DataAbsentReason = "asked-unknown"
	DataAbsentReasonTempUnknown      DataAbsentReason = "temp-unknown"
	DataAbsentReasonNotAsked         DataAbsentReason = "not-asked"
	DataAbsentReasonAskedDeclined    DataAbsentReason = "asked-declined"
	DataAbsentReasonMasked           DataAbsentReason = "masked"
	DataAbsentReasonNotApplicable    DataAbsentReason = "not-applicable"
	DataAbsentReasonUnsupported      DataAbsentReason = "unsupported"
	DataAbsentReasonAsText           DataAbsentReason = "as-text"
	DataAbsentReasonError            DataAbsentReason = "error"
	DataAbsentReasonNotANumber       DataAbsentReason = "not-a-number"
	DataAbsentReasonNegativeInfinity DataAbsentReason = "negative-infinity"
	DataAbsentReasonPositiveInfinity DataAbsentReason = "positive-infinity"
	DataAbsentReasonNotPerformed     DataAbsentReason = "not-performed"
	DataAbsentReasonNotPermitted     DataAbsentReason = "not-permitted"
)

func AllDataAbsentReason() []DataAbsentReason {
	return []DataAbsentReason{
		DataAbsentReasonUnknown,
		DataAbsentReasonAskedUnknown,
		DataAbsentReasonTempUnknown,
		DataAbsentReasonNotAsked,
		DataAbsentReasonAskedDeclined,
		DataAbsentReasonMasked,
		DataAbsentReasonNotApplicable,
		DataAbsentReasonUnsupported,
		DataAbsentReasonAsText,
		DataAbsentReasonError,
		DataAbsentReasonNotANumber,
		DataAbsentReasonNegativeInfinity,
		DataAbsentReasonPositiveInfinity,
		DataAbsentReasonNotPerformed,
		DataAbsentReasonNotPermitted,
	}
}

func FindDataAbsentReason(filter string) []DataAbsentReason {
	ret := make([]DataAbsentReason, 0)
	for _, at := range AllDataAbsentReason() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au DataAbsentReason) ToString() {
	fmt.Println(au.String())
}
func (au DataAbsentReason) String() string {
	switch au {
	case DataAbsentReasonUnknown:
		return "Unknown"
	case DataAbsentReasonAskedUnknown:
		return "Asked But Unknown"
	case DataAbsentReasonTempUnknown:
		return "Temporarily Unknown"
	case DataAbsentReasonNotAsked:
		return "Not Asked"
	case DataAbsentReasonAskedDeclined:
		return "Asked But Declined"
	case DataAbsentReasonMasked:
		return "Masked"
	case DataAbsentReasonNotApplicable:
		return "Not Applicable"
	case DataAbsentReasonUnsupported:
		return "Unsupported"
	case DataAbsentReasonAsText:
		return "As Text"
	case DataAbsentReasonError:
		return "Error"
	case DataAbsentReasonNotANumber:
		return "Not a Number (NaN)"
	case DataAbsentReasonNegativeInfinity:
		return "Negative Infinity (NINF)"
	case DataAbsentReasonPositiveInfinity:
		return "Positive Infinity (PINF)"
	case DataAbsentReasonNotPerformed:
		return "Not Performed"
	case DataAbsentReasonNotPermitted:
		return "Not Permitted"
	default:
		return "Unknown Data Absent Reason"
	}
}
