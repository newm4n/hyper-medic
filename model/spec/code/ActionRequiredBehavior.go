package code

type ActionRequiredBehavior int

const (
	ActionRequiredBehaviorMust ActionRequiredBehavior = iota
	ActionRequiredBehaviorCould
	ActionRequiredBehaviorMustUnlessDocumented
)

/*
Must
Could
Must Unless Documented
*/
