package code

type ActionSelectionBehavior int

const (
	ActionSelectionBehaviorAny ActionSelectionBehavior = iota
	ActionSelectionBehaviorAll
	ActionSelectionBehaviorAllOrNone
	ActionSelectionBehaviorExactlyOne
	ActionSelectionBehaviorAtMostOne
	ActionSelectionBehaviorOneOrMore
)

/*
Any
All
All Or None
Exactly One
At Most One
One Or More
*/
