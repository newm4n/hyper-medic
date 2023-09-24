package code

type ActionType int

const (
	ActionTypeCreate ActionType = iota
	ActionTypeUpdate
	ActionTypeRemove
	ActionTypeFireEvent
)

/*
Create
Update
Remove
Fire Event
*/
