package code

type AdverseEventSeriousness int

const (
	AdverseEventSeriousnessNonserious AdverseEventSeriousness = iota
	AdverseEventSeriousnessSerious
)

/*
Non-serious
Serious
*/
