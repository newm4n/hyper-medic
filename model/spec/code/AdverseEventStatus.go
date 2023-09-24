package code

type AdverseEventStatus int

const (
	AdverseEventStatusInProgress AdverseEventStatus = iota
	AdverseEventStatusCompleted
	AdverseEventStatusEnteredInError
	AdverseEventStatusUnknown
)

/**
In Progress
Completed
Entered in Error
Unknown
*/
