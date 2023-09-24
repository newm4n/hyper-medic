package code

type Adjudication int

const (
	AdjudicationSubmittedAmount Adjudication = iota
	AdjudicationCoPay
	AdjudicationEligibleAmount
	AdjudicationDeductible
	AdjudicationUnallocatedDeductible
	AdjudicationEligiblePercentage
	Adjudicationax
	AdjudicationBenefitAmount
)

/*
Submitted Amount
CoPay
Eligible Amount
Deductible
Unallocated Deductible
Eligible %
Tax
Benefit Amount
*/
