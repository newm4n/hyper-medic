package code

import (
	"fmt"
	"strings"
)

type QuantityComparator string

const (
	QuantityComparatorLT  QuantityComparator = "<"
	QuantityComparatorLTE QuantityComparator = "<="
	QuantityComparatorGTE QuantityComparator = ">="
	QuantityComparatorGT  QuantityComparator = ">"
	QuantityComparatorE   QuantityComparator = "="
)

func AllQuantityComparator() []QuantityComparator {
	return []QuantityComparator{
		QuantityComparatorLT,
		QuantityComparatorLTE,
		QuantityComparatorGTE,
		QuantityComparatorGT,
		QuantityComparatorE,
	}
}

func FindQuantityComparator(filter string) []QuantityComparator {
	ret := make([]QuantityComparator, 0)
	for _, at := range AllQuantityComparator() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au QuantityComparator) ToString() {
	fmt.Println(au.String())
}
func (au QuantityComparator) String() string {
	switch au {
	case QuantityComparatorLT:
		return "Less than"
	case QuantityComparatorLTE:
		return "Less or Equal to"
	case QuantityComparatorGTE:
		return "Greater or Equal to"
	case QuantityComparatorGT:
		return "Greater than"
	case QuantityComparatorE:
		return "Sufficient to achieve this total quantity"
	default:
		return "Unknown Quantity Comparator"
	}
}
