package code

import (
	"fmt"
	"strings"
)

type AccountBalanceTerm int

const (
	AccountBalanceTermCurrent AccountBalanceTerm = iota
	AccountBalanceTerm30
	AccountBalanceTerm60
	AccountBalanceTerm90
	AccountBalanceTerm120
)

func AllAccountBalanceTerm() []AccountBalanceTerm {
	return []AccountBalanceTerm{
		AccountBalanceTermCurrent,
		AccountBalanceTerm30,
		AccountBalanceTerm60,
		AccountBalanceTerm90,
		AccountBalanceTerm120,
	}
}

func FindAccountBalanceTerm(filter string) []AccountBalanceTerm {
	ret := make([]AccountBalanceTerm, 0)
	for _, at := range AllAccountBalanceTerm() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AccountBalanceTerm) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AccountBalanceTerm) String() string {
	switch cpt {
	case AccountBalanceTermCurrent:
		return "current"
	case AccountBalanceTerm30:
		return "30"
	case AccountBalanceTerm60:
		return "60"
	case AccountBalanceTerm90:
		return "90"
	case AccountBalanceTerm120:
		return "120"
	default:
		return "Unknown Account Balance Term"
	}
}

/*
current
30
60
90
120
*/
