package code

type AdverseEventOutcome int

const (
	AdverseEventOutcome1 AdverseEventOutcome = iota
	AdverseEventOutcome2
	AdverseEventOutcome3
	AdverseEventOutcome4
	AdverseEventOutcome5
)

/*
Transient abnormality unnoticed by the patient (finding)
Transient abnormality with full recovery (finding)
Adverse incident resulting in potentially permanent but not disabling damage (finding)
Adverse incident resulting in potentially permanent disabling damage (finding)
Adverse incident resulting in death (finding)
*/
