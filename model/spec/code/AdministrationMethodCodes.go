package code

import (
	"fmt"
	"strings"
)

type AdministrationMethodCode string

const (
	AdministrationMethodCodeDoseFormAdministrationMethod AdministrationMethodCode = "736665006"
	AdministrationMethodCodeAdminister                   AdministrationMethodCode = "738990001"
	AdministrationMethodCodeApply                        AdministrationMethodCode = "738991002"
	AdministrationMethodCodeChew                         AdministrationMethodCode = "738992009"
	AdministrationMethodCodeInsert                       AdministrationMethodCode = "738993004"
	AdministrationMethodCodeInstill                      AdministrationMethodCode = "738994005"
	AdministrationMethodCodeSwallow                      AdministrationMethodCode = "738995006"
	AdministrationMethodCodeSpray                        AdministrationMethodCode = "738996007"
	AdministrationMethodCodeInhale                       AdministrationMethodCode = "740666001"
	AdministrationMethodCodeInject                       AdministrationMethodCode = "740685003"
	AdministrationMethodCodeSuck                         AdministrationMethodCode = "764498003"
	AdministrationMethodCodeInfuse                       AdministrationMethodCode = "764794000"
	AdministrationMethodCodeRinse                        AdministrationMethodCode = "782155003"
	AdministrationMethodCodeGargle                       AdministrationMethodCode = "782168006"
	AdministrationMethodCodeRinseOrWash                  AdministrationMethodCode = "785900008"
	AdministrationMethodCodeOrodisperse                  AdministrationMethodCode = "823034001"
	AdministrationMethodCodeImplant                      AdministrationMethodCode = "827107003"
	AdministrationMethodCodeInsufflate                   AdministrationMethodCode = "1010690008"
	AdministrationMethodCodeDialysis                     AdministrationMethodCode = "1231460007"
	AdministrationMethodCodeBathe                        AdministrationMethodCode = "58841000052102"
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
