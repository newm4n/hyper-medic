package code

import (
	"fmt"
	"strings"
)

type BenefitNetwork string

const (
	BenefitNetworkIn  BenefitNetwork = "in"
	BenefitNetworkOut BenefitNetwork = "out"
)

func AllBenefitNetwork() []BenefitNetwork {
	return []BenefitNetwork{
		BenefitNetworkIn,
		BenefitNetworkOut,
	}
}

func FindBenefitNetwork(filter string) []BenefitNetwork {
	ret := make([]BenefitNetwork, 0)
	for _, at := range AllBenefitNetwork() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BenefitNetwork) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BenefitNetwork) String() string {
	switch cpt {
	case BenefitNetworkIn:
		return "In Network"
	case BenefitNetworkOut:
		return "Out of Network"
	default:
		return "Unknown Benefit Network"
	}
}
