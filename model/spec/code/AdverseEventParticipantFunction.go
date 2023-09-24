package code

type AdverseEventParticipantFunction int

const (
	AdverseEventParticipantFunctionInformant AdverseEventParticipantFunction = iota
	AdverseEventParticipantFunctionParticipation
	AdverseEventParticipantFunctionWitness
	AdverseEventParticipantFunctionAuthorOriginator
)

/**
informant
Participation
witness
author (originator)
*/
