package code

type ActionConditionKind int

const (
	ActionConditionKindApplicability ActionConditionKind = iota
	ActionConditionKindStart
	ActionConditionKindStop
)

/*
Applicability
Start
Stop
*/
