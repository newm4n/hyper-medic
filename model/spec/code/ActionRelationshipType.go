package code

type ActionRelationshipType int

const (
	ActionRelationshipTypeBefore ActionRelationshipType = iota
	ActionRelationshipTypeBeforeStart
	ActionRelationshipTypeBeforeEnd
	ActionRelationshipTypeConcurrent
	ActionRelationshipTypeConcurrentWithStart
	ActionRelationshipTypeConcurrentWithEnd
	ActionRelationshipTypeAfter
	ActionRelationshipTypeAfterStart
	ActionRelationshipTypeAfterEnd
)

/*
Before
Before Start
Before End
Concurrent
Concurrent with start
Concurrent with end
After
After Start
After End
*/
