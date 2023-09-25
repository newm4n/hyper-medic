package code

import (
	"fmt"
	"strings"
)

type AdministrationMethodCode int

const (
	AdministrationMethodCodeDoseFormAdministrationMethod AdministrationMethodCode = iota
	AdministrationMethodCodeAdminister
	AdministrationMethodCodeApply
	AdministrationMethodCodeChew
	AdministrationMethodCodeInsert
	AdministrationMethodCodeInstill
	AdministrationMethodCodeSwallow
	AdministrationMethodCodeSpray
	AdministrationMethodCodeInhale
	AdministrationMethodCodeInject
	AdministrationMethodCodeSuck
	AdministrationMethodCodeInfuse
	AdministrationMethodCodeRinse
	AdministrationMethodCodeGargle
	AdministrationMethodCodeRinseOrWash
	AdministrationMethodCodeOrodisperse
	AdministrationMethodCodeImplant
	AdministrationMethodCodeInsufflate
	AdministrationMethodCodeDialysis
	AdministrationMethodCodeBathe
)

func AllAdministrationMethodCode() []AdministrationMethodCode {
	return []AdministrationMethodCode{
		AdministrationMethodCodeDoseFormAdministrationMethod,
		AdministrationMethodCodeAdminister,
		AdministrationMethodCodeApply,
		AdministrationMethodCodeChew,
		AdministrationMethodCodeInsert,
		AdministrationMethodCodeInstill,
		AdministrationMethodCodeSwallow,
		AdministrationMethodCodeSpray,
		AdministrationMethodCodeInhale,
		AdministrationMethodCodeInject,
		AdministrationMethodCodeSuck,
		AdministrationMethodCodeInfuse,
		AdministrationMethodCodeRinse,
		AdministrationMethodCodeGargle,
		AdministrationMethodCodeRinseOrWash,
		AdministrationMethodCodeOrodisperse,
		AdministrationMethodCodeImplant,
		AdministrationMethodCodeInsufflate,
		AdministrationMethodCodeDialysis,
		AdministrationMethodCodeBathe,
	}
}

func FindAdministrationMethodCode(filter string) []AdministrationMethodCode {
	ret := make([]AdministrationMethodCode, 0)
	for _, at := range AllAdministrationMethodCode() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdministrationMethodCode) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdministrationMethodCode) String() string {
	switch cpt {
	case AdministrationMethodCodeDoseFormAdministrationMethod:
		return "Dose form administration method"
	case AdministrationMethodCodeAdminister:
		return "Administer"
	case AdministrationMethodCodeApply:
		return "Apply"
	case AdministrationMethodCodeChew:
		return "Chew"
	case AdministrationMethodCodeInsert:
		return "Insert"
	case AdministrationMethodCodeInstill:
		return "Instill"
	case AdministrationMethodCodeSwallow:
		return "Swallow"
	case AdministrationMethodCodeSpray:
		return "Spray"
	case AdministrationMethodCodeInhale:
		return "Inhale"
	case AdministrationMethodCodeInject:
		return "Inject"
	case AdministrationMethodCodeSuck:
		return "Suck"
	case AdministrationMethodCodeInfuse:
		return "Infuse"
	case AdministrationMethodCodeRinse:
		return "Rinse"
	case AdministrationMethodCodeGargle:
		return "Gargle"
	case AdministrationMethodCodeRinseOrWash:
		return "Rinse or wash"
	case AdministrationMethodCodeOrodisperse:
		return "Orodisperse"
	case AdministrationMethodCodeImplant:
		return "Implant"
	case AdministrationMethodCodeInsufflate:
		return "Insufflate"
	case AdministrationMethodCodeDialysis:
		return "Dialysis"
	case AdministrationMethodCodeBathe:
		return "Bathe"
	default:
		return "Unknown Adverse Event Actuality"
	}
}

/**
Dose form administration method
Administer
Apply
Chew
Insert
Instill
Swallow
Spray
Inhale
Inject
Suck
Infuse
Rinse
Gargle
Rinse or wash
Orodisperse
Implant
Insufflate
Dialysis
Bathe

*/
