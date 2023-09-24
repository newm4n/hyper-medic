package code

type AdverseEventCausalityMethod int

const (
	AdverseEventCausalityMethodProbabilityScale AdverseEventCausalityMethod = iota
	AdverseEventCausalityMethodBayesian
	AdverseEventCausalityMethodChecklist
)

/*
Probability Scale
Bayesian
Checklist
*/
