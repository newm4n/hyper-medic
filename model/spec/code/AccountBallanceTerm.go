package code

type AccountBalanceTerm int

const (
	AccountBalanceTermCurrent AccountBalanceTerm = iota
	AccountBalanceTerm30
	AccountBalanceTerm60
	AccountBalanceTerm90
	AccountBalanceTerm120
)

/*
current
30
60
90
120
*/
