package code

import (
	"fmt"
	"strings"
)

type BindingStrength string

const (
	BindingStrengthRequired   BindingStrength = "required"
	BindingStrengthExtensible BindingStrength = "extensible"
	BindingStrengthPreferred  BindingStrength = "preferred"
	BindingStrengthExample    BindingStrength = "example"
)

func AllBindingStrength() []BindingStrength {
	return []BindingStrength{
		BindingStrengthRequired,
		BindingStrengthExtensible,
		BindingStrengthPreferred,
		BindingStrengthExample,
	}
}

func FindBindingStrength(filter string) []BindingStrength {
	ret := make([]BindingStrength, 0)
	for _, at := range AllBindingStrength() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BindingStrength) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BindingStrength) String() string {
	switch cpt {
	case BindingStrengthRequired:
		return "Required"
	case BindingStrengthExtensible:
		return "Extensible"
	case BindingStrengthPreferred:
		return "Preferred"
	case BindingStrengthExample:
		return "Example"
	default:
		return "Unknown Binding Strength"
	}
}
