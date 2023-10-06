package code

import (
	"fmt"
	"strings"
)

type BiologicallyderivedProductcodes string

const (
	BiologicallyderivedProductcodese0398 BiologicallyderivedProductcodes = "e0398"
	BiologicallyderivedProductcodess1128 BiologicallyderivedProductcodes = "s1128"
	BiologicallyderivedProductcodess1194 BiologicallyderivedProductcodes = "s1194"
	BiologicallyderivedProductcodess1195 BiologicallyderivedProductcodes = "s1195"
	BiologicallyderivedProductcodess1310 BiologicallyderivedProductcodes = "s1310"
	BiologicallyderivedProductcodess1398 BiologicallyderivedProductcodes = "s1398"
	BiologicallyderivedProductcodess2598 BiologicallyderivedProductcodes = "s2598"
	BiologicallyderivedProductcodese4377 BiologicallyderivedProductcodes = "e4377"
	BiologicallyderivedProductcodest1396 BiologicallyderivedProductcodes = "t1396"
)

func AllBiologicallyderivedProductcodes() []BiologicallyderivedProductcodes {
	return []BiologicallyderivedProductcodes{
		BiologicallyderivedProductcodese0398,
		BiologicallyderivedProductcodess1128,
		BiologicallyderivedProductcodess1194,
		BiologicallyderivedProductcodess1195,
		BiologicallyderivedProductcodess1310,
		BiologicallyderivedProductcodess1398,
		BiologicallyderivedProductcodess2598,
		BiologicallyderivedProductcodese4377,
		BiologicallyderivedProductcodest1396,
	}
}

func FindBiologicallyderivedProductcodes(filter string) []BiologicallyderivedProductcodes {
	ret := make([]BiologicallyderivedProductcodes, 0)
	for _, at := range AllBiologicallyderivedProductcodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BiologicallyderivedProductcodes) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BiologicallyderivedProductcodes) String() string {
	switch cpt {
	case BiologicallyderivedProductcodese0398:
		return "RED BLOOD CELLS|CPD>AS5/450mL/refg|Irr|ResLeu"
	case BiologicallyderivedProductcodess1128:
		return "HPC, APHERESIS/Citrate/XX/refg/Mobilized"
	case BiologicallyderivedProductcodess1194:
		return "HPC, APHERESIS|NS/XX/<=-120C|10% DMSO|Cryopreserved|Mobilized"
	case BiologicallyderivedProductcodess1195:
		return "productcodes\tHPC, APHERESIS|NS/XX/<=-120C|5% DMSO|Cryopreserved|Mobilized"
	case BiologicallyderivedProductcodess1310:
		return "HPC, APHERESIS|None/XX/refg|3rd Party Comp:Yes|Other Additives:Yes|Mobilized|CD34 enriched"
	case BiologicallyderivedProductcodess1398:
		return "\tHPC, MARROW|NS/XX/rt|Plasma reduced"
	case BiologicallyderivedProductcodess2598:
		return "HPC, MARROW|NS/XX/<=-150C|10% DMSO|3rd Party Comp:Yes|Cryopreserved|RBC reduced"
	case BiologicallyderivedProductcodese4377:
		return "\tApheresis RED BLOOD CELLS|ACD-A/XX/refg|Irradiated|1st container"
	case BiologicallyderivedProductcodest1396:
		return "BONE, FEMUR|Frozen|Right|Radiation sterilization"
	default:
		return "Unknown Biologically derived Product codes"
	}
}
