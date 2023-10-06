package code

import (
	"fmt"
	"strings"
)

type BiologicallyderivedproductdispenseOriginRelationship string

const (
	BiologicallyderivedproductdispenseOriginRelationshipAutologous BiologicallyderivedproductdispenseOriginRelationship = "autologous"
	BiologicallyderivedproductdispenseOriginRelationshipRelated    BiologicallyderivedproductdispenseOriginRelationship = "related"
	BiologicallyderivedproductdispenseOriginRelationshipDirected   BiologicallyderivedproductdispenseOriginRelationship = "directed"
	BiologicallyderivedproductdispenseOriginRelationshipAllogeneic BiologicallyderivedproductdispenseOriginRelationship = "allogeneic"
	BiologicallyderivedproductdispenseOriginRelationshipXenogenic  BiologicallyderivedproductdispenseOriginRelationship = "xenogenic"
)

func AllBiologicallyderivedproductdispenseOriginRelationship() []BiologicallyderivedproductdispenseOriginRelationship {
	return []BiologicallyderivedproductdispenseOriginRelationship{
		BiologicallyderivedproductdispenseOriginRelationshipAutologous,
		BiologicallyderivedproductdispenseOriginRelationshipRelated,
		BiologicallyderivedproductdispenseOriginRelationshipDirected,
		BiologicallyderivedproductdispenseOriginRelationshipAllogeneic,
		BiologicallyderivedproductdispenseOriginRelationshipXenogenic,
	}
}

func FindBiologicallyderivedproductdispenseOriginRelationship(filter string) []BiologicallyderivedproductdispenseOriginRelationship {
	ret := make([]BiologicallyderivedproductdispenseOriginRelationship, 0)
	for _, at := range AllBiologicallyderivedproductdispenseOriginRelationship() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BiologicallyderivedproductdispenseOriginRelationship) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BiologicallyderivedproductdispenseOriginRelationship) String() string {
	switch cpt {
	case BiologicallyderivedproductdispenseOriginRelationshipAutologous:
		return "Autologous"
	case BiologicallyderivedproductdispenseOriginRelationshipRelated:
		return "Related"
	case BiologicallyderivedproductdispenseOriginRelationshipDirected:
		return "Directed"
	case BiologicallyderivedproductdispenseOriginRelationshipAllogeneic:
		return "Allogeneic"
	case BiologicallyderivedproductdispenseOriginRelationshipXenogenic:
		return "Xenogenic"
	default:
		return "Unknown Biologically derived product dispense Origin Relationship"
	}
}
