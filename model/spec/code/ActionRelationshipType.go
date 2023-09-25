package code

import (
	"fmt"
	"strings"
)

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

func AllActionRelationshipType() []ActionRelationshipType {
	return []ActionRelationshipType{
		ActionRelationshipTypeBefore,
		ActionRelationshipTypeBeforeStart,
		ActionRelationshipTypeBeforeEnd,
		ActionRelationshipTypeConcurrent,
		ActionRelationshipTypeConcurrentWithStart,
		ActionRelationshipTypeConcurrentWithEnd,
		ActionRelationshipTypeAfter,
		ActionRelationshipTypeAfterStart,
		ActionRelationshipTypeAfterEnd,
	}
}

func FindActionRelationshipType(filter string) []ActionRelationshipType {
	ret := make([]ActionRelationshipType, 0)
	for _, at := range AllActionType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt ActionRelationshipType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt ActionRelationshipType) String() string {
	switch cpt {
	case ActionRelationshipTypeBefore:
		return "Before"
	case ActionRelationshipTypeBeforeStart:
		return "Before Start"
	case ActionRelationshipTypeBeforeEnd:
		return "Before End"
	case ActionRelationshipTypeConcurrent:
		return "Concurrent"
	case ActionRelationshipTypeConcurrentWithStart:
		return "Concurrent with start"
	case ActionRelationshipTypeConcurrentWithEnd:
		return "Concurrent with end"
	case ActionRelationshipTypeAfter:
		return "After"
	case ActionRelationshipTypeAfterStart:
		return "After Start"
	case ActionRelationshipTypeAfterEnd:
		return "After End"
	default:
		return "Unknown Action Relationship Type"
	}
}

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
