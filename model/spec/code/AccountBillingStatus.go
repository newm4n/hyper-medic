package code

type AccountBillingStatus int

const (
	AccountBillingStatusOpen AccountBillingStatus = iota
	AccountBillingStatusCareCompleteNotbilled
	AccountBillingStatusBilling
	AccountBillingStatusClosedBaddebt
	AccountBillingStatusClosedVoided
	AccountBillingStatusClosedCompleted
	AccountBillingStatusClosedCombined
)

/*
open
 carecomplete-notbilled
billing
closed-baddebt
closed-voided
 closed-completed
 closed-combined
*/
