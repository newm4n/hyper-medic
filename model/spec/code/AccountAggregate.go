package code

import (
	"fmt"
	"strings"
)

type AccountAggregate string

const (
	AccountAggregatePatent    AccountAggregate = "patient"
	AccountAggregateInsurance AccountAggregate = "insurance"
	AccountAggregateTotal     AccountAggregate = "total"
)

func AllAccountAggregate() []AccountAggregate {
	return []AccountAggregate{
		AccountAggregatePatent,
		AccountAggregateInsurance,
		AccountAggregateTotal,
	}
}

func FindAccountAggregate(filter string) []AccountAggregate {
	ret := make([]AccountAggregate, 0)
	for _, at := range AllAccountAggregate() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AccountAggregate) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AccountAggregate) String() string {
	switch cpt {
	case AccountAggregatePatent:
		return "Patent"
	case AccountAggregateInsurance:
		return "Insurance"
	case AccountAggregateTotal:
		return "Total"
	default:
		return "Unknown Account Aggregate"
	}
}

/*
Patent
Insurance
Total
*/
