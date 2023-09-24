package code

type ActionParticipantFunction int

const (
	ActionParticipantFunctionPerformer ActionParticipantFunction = iota
	ActionParticipantFunctionAuthor
	ActionParticipantFunctionReviewer
	ActionParticipantFunctionWitness
)

/*
Performer
Author
Reviewer
Witness
*/
