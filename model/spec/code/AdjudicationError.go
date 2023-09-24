package code

type AdjudicationError int

const (
	AdjudicationErrorMissingIdentifier AdjudicationError = iota
	AdjudicationErrorMissingCreationDate
)

/**
Missing Identifier
Missing Creation Date
*/
