package code

import (
	"fmt"
	"strings"
)

type AccountRelationship string

const (
	AccountRelationshipParent    AccountRelationship = "parent"
	AccountRelationshipGuarantor AccountRelationship = "guarantor"
)

func AllAccountRelationship() []AccountRelationship {
	return []AccountRelationship{
		AccountRelationshipParent,
		AccountRelationshipGuarantor,
	}
}

func FindAccountRelationship(filter string) []AccountRelationship {
	ret := make([]AccountRelationship, 0)
	for _, at := range AllAccountRelationship() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AccountRelationship) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AccountRelationship) String() string {
	switch cpt {
	case AccountRelationshipParent:
		return "Parent"
	case AccountRelationshipGuarantor:
		return "Guarantor"
	default:
		return "Unknown Account Relationship"
	}
}

/*
Parent
Guarantor
*/
