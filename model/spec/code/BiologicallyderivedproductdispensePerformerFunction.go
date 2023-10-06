package code

import (
	"fmt"
	"strings"
)

type BiologicallyderivedproductdispensePerformerFunction string

const (
	BiologicallyderivedproductdispensePerformerFunctionGroupType              BiologicallyderivedproductdispensePerformerFunction = "group-and-type"
	BiologicallyderivedproductdispensePerformerFunctionAntibodyScreen         BiologicallyderivedproductdispensePerformerFunction = "antibody-screen"
	BiologicallyderivedproductdispensePerformerFunctionAntibodyIdentification BiologicallyderivedproductdispensePerformerFunction = "antibody-identification"
	BiologicallyderivedproductdispensePerformerFunctionCrossmatch             BiologicallyderivedproductdispensePerformerFunction = "crossmatch"
	BiologicallyderivedproductdispensePerformerFunctionRelease                BiologicallyderivedproductdispensePerformerFunction = "release"
	BiologicallyderivedproductdispensePerformerFunctionTransport              BiologicallyderivedproductdispensePerformerFunction = "transport"
	BiologicallyderivedproductdispensePerformerFunctionReceipt                BiologicallyderivedproductdispensePerformerFunction = "receipt"
)

func AllBiologicallyderivedproductdispensePerformerFunction() []BiologicallyderivedproductdispensePerformerFunction {
	return []BiologicallyderivedproductdispensePerformerFunction{
		BiologicallyderivedproductdispensePerformerFunctionGroupType,
		BiologicallyderivedproductdispensePerformerFunctionAntibodyScreen,
		BiologicallyderivedproductdispensePerformerFunctionAntibodyIdentification,
		BiologicallyderivedproductdispensePerformerFunctionCrossmatch,
		BiologicallyderivedproductdispensePerformerFunctionRelease,
		BiologicallyderivedproductdispensePerformerFunctionTransport,
		BiologicallyderivedproductdispensePerformerFunctionReceipt,
	}
}

func FindBiologicallyderivedproductdispensePerformerFunction(filter string) []BiologicallyderivedproductdispensePerformerFunction {
	ret := make([]BiologicallyderivedproductdispensePerformerFunction, 0)
	for _, at := range AllBiologicallyderivedproductdispensePerformerFunction() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BiologicallyderivedproductdispensePerformerFunction) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BiologicallyderivedproductdispensePerformerFunction) String() string {
	switch cpt {
	case BiologicallyderivedproductdispensePerformerFunctionGroupType:
		return "Group and Type"
	case BiologicallyderivedproductdispensePerformerFunctionAntibodyScreen:
		return "Antibody Screen"
	case BiologicallyderivedproductdispensePerformerFunctionAntibodyIdentification:
		return "Antibody Identification"
	case BiologicallyderivedproductdispensePerformerFunctionCrossmatch:
		return "Crossmatch"
	case BiologicallyderivedproductdispensePerformerFunctionRelease:
		return "Release"
	case BiologicallyderivedproductdispensePerformerFunctionTransport:
		return "Transport"
	case BiologicallyderivedproductdispensePerformerFunctionReceipt:
		return "Receipt"
	default:
		return "Unknown Biologically derived product dispense Performer Function"
	}
}
