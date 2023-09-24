package code

type AllergyIntoleranceCriticality int

const (
	AllergyIntoleranceCriticalityLowRisk AllergyIntoleranceCriticality = iota
	AllergyIntoleranceCriticalityHighRisk
	AllergyIntoleranceCriticalityUnableToAssessRisk
)

/*
Low Risk
High Risk
Unable to Assess Risk
*/
