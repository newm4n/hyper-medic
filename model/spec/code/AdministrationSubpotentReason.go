package code

type AdministrationSubpotentReason int

const (
	AdministrationSubpotentReasonPartialDose AdministrationSubpotentReason = iota
	AdministrationSubpotentReasonVomited
	AdministrationSubpotentReasonColdChainBreak
	AdministrationSubpotentReasonManufacturerRecall
	AdministrationSubpotentReasonAdverseStorage
	AdministrationSubpotentReasonExpiredProduct
)

/*
Partial Dose
Vomited
Cold Chain Break
Manufacturer Recall
Adverse Storage
Expired Product
*/
