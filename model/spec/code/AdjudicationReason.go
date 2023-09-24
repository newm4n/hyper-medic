package code

type AdjudicationReason int

const (
	AdjudicationReasonNotCovered AdjudicationReason = iota
	AdjudicationReasonPlanLimitReached
)

/*
Not covered
Plan Limit Reached
*/
