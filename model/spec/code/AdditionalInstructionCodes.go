package code

import (
	"fmt"
	"strings"
)

type AdditionalInstructionCodes string

const (
	AdditionalInstructionCodes00 AdditionalInstructionCodes = "419492006"
	AdditionalInstructionCodes01 AdditionalInstructionCodes = "311501008"
	AdditionalInstructionCodes02 AdditionalInstructionCodes = "311504000"
	AdditionalInstructionCodes03 AdditionalInstructionCodes = "417929005"
	AdditionalInstructionCodes04 AdditionalInstructionCodes = "417980006"
	AdditionalInstructionCodes05 AdditionalInstructionCodes = "417995008"
	AdditionalInstructionCodes06 AdditionalInstructionCodes = "418071006"
	AdditionalInstructionCodes07 AdditionalInstructionCodes = "418194009"
	AdditionalInstructionCodes08 AdditionalInstructionCodes = "418281004"
	AdditionalInstructionCodes09 AdditionalInstructionCodes = "418443006"
	AdditionalInstructionCodes10 AdditionalInstructionCodes = "418521000"
	AdditionalInstructionCodes11 AdditionalInstructionCodes = "418577003"
	AdditionalInstructionCodes12 AdditionalInstructionCodes = "418637003"
	AdditionalInstructionCodes13 AdditionalInstructionCodes = "418639000"
	AdditionalInstructionCodes14 AdditionalInstructionCodes = "418693002"
	AdditionalInstructionCodes15 AdditionalInstructionCodes = "418849000"
	AdditionalInstructionCodes16 AdditionalInstructionCodes = "418850000"
	AdditionalInstructionCodes17 AdditionalInstructionCodes = "418914006"
	AdditionalInstructionCodes18 AdditionalInstructionCodes = "418954008"
	AdditionalInstructionCodes19 AdditionalInstructionCodes = "418991002"
	AdditionalInstructionCodes20 AdditionalInstructionCodes = "419111009"
	AdditionalInstructionCodes21 AdditionalInstructionCodes = "419115000"
	AdditionalInstructionCodes22 AdditionalInstructionCodes = "419303009"
	AdditionalInstructionCodes23 AdditionalInstructionCodes = "419437002"
	AdditionalInstructionCodes24 AdditionalInstructionCodes = "419439004"
	AdditionalInstructionCodes25 AdditionalInstructionCodes = "419444006"
	AdditionalInstructionCodes26 AdditionalInstructionCodes = "419473009"
	AdditionalInstructionCodes27 AdditionalInstructionCodes = "419529008"
	AdditionalInstructionCodes28 AdditionalInstructionCodes = "419822006"
	AdditionalInstructionCodes29 AdditionalInstructionCodes = "419974005"
	AdditionalInstructionCodes30 AdditionalInstructionCodes = "420003005"
	AdditionalInstructionCodes31 AdditionalInstructionCodes = "420082003"
	AdditionalInstructionCodes32 AdditionalInstructionCodes = "420110001"
	AdditionalInstructionCodes33 AdditionalInstructionCodes = "420162004"
	AdditionalInstructionCodes34 AdditionalInstructionCodes = "420652005"
	AdditionalInstructionCodes35 AdditionalInstructionCodes = "421484000"
	AdditionalInstructionCodes36 AdditionalInstructionCodes = "421769005"
	AdditionalInstructionCodes37 AdditionalInstructionCodes = "421984009"
	AdditionalInstructionCodes38 AdditionalInstructionCodes = "422327006"
	AdditionalInstructionCodes39 AdditionalInstructionCodes = "428579001"
	AdditionalInstructionCodes40 AdditionalInstructionCodes = "717154004"
)

func AllAdditionalInstructionCodes() []AdditionalInstructionCodes {
	return []AdditionalInstructionCodes{
		AdditionalInstructionCodes00,
		AdditionalInstructionCodes01,
		AdditionalInstructionCodes02,
		AdditionalInstructionCodes03,
		AdditionalInstructionCodes04,
		AdditionalInstructionCodes05,
		AdditionalInstructionCodes06,
		AdditionalInstructionCodes07,
		AdditionalInstructionCodes08,
		AdditionalInstructionCodes09,
		AdditionalInstructionCodes10,
		AdditionalInstructionCodes11,
		AdditionalInstructionCodes12,
		AdditionalInstructionCodes13,
		AdditionalInstructionCodes14,
		AdditionalInstructionCodes15,
		AdditionalInstructionCodes16,
		AdditionalInstructionCodes17,
		AdditionalInstructionCodes18,
		AdditionalInstructionCodes19,
		AdditionalInstructionCodes20,
		AdditionalInstructionCodes21,
		AdditionalInstructionCodes22,
		AdditionalInstructionCodes23,
		AdditionalInstructionCodes24,
		AdditionalInstructionCodes25,
		AdditionalInstructionCodes26,
		AdditionalInstructionCodes27,
		AdditionalInstructionCodes28,
		AdditionalInstructionCodes29,
		AdditionalInstructionCodes30,
		AdditionalInstructionCodes31,
		AdditionalInstructionCodes32,
		AdditionalInstructionCodes33,
		AdditionalInstructionCodes34,
		AdditionalInstructionCodes35,
		AdditionalInstructionCodes36,
		AdditionalInstructionCodes37,
		AdditionalInstructionCodes38,
		AdditionalInstructionCodes39,
		AdditionalInstructionCodes40,
	}
}

func FindAdditionalInstructionCodes(filter string) []AdditionalInstructionCodes {
	ret := make([]AdditionalInstructionCodes, 0)
	for _, at := range AllAdditionalInstructionCodes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdditionalInstructionCodes) ToString() {
	fmt.Println(cpt.String())
}

func (pq AdditionalInstructionCodes) String() string {
	switch pq {
	case AdditionalInstructionCodes00:
		return "Additional dosage instructions (qualifier value)"
	case AdditionalInstructionCodes01:
		return "Half to one hour before food"
	case AdditionalInstructionCodes02:
		return "With or after food"
	case AdditionalInstructionCodes03:
		return "Times (qualifier value)"
	case AdditionalInstructionCodes04:
		return "Contains aspirin (qualifier value)"
	case AdditionalInstructionCodes05:
		return "Dissolve or mix with water before taking (qualifier value)"
	case AdditionalInstructionCodes06:
		return "Warning. Causes drowsiness which may continue the next day. If affected do not drive or operate machinery. Avoid alcoholic drink (qualifier value)"
	case AdditionalInstructionCodes07:
		return "Contains an aspirin-like medicine (qualifier value)"
	case AdditionalInstructionCodes08:
		return "Do not take anything containing aspirin while taking this medicine (qualifier value)"
	case AdditionalInstructionCodes09:
		return "Do not take more than . . . in 24 hours or . . . in any one week (qualifier value)"
	case AdditionalInstructionCodes10:
		return "Avoid exposure of skin to direct sunlight or sun lamps (qualifier value)"
	case AdditionalInstructionCodes11:
		return "Take at regular intervals. Complete the prescribed course unless otherwise directed (qualifier value)"
	case AdditionalInstructionCodes12:
		return "Do not take with any other paracetamol products (qualifier value)"
	case AdditionalInstructionCodes13:
		return "Warning. May cause drowsiness (qualifier value)"
	case AdditionalInstructionCodes14:
		return "Swallowed whole, not chewed (qualifier value)"
	case AdditionalInstructionCodes15:
		return "Warning. Follow the printed instructions you have been given with this medicine (qualifier value)"
	case AdditionalInstructionCodes16:
		return "Contains aspirin and paracetamol. Do not take with any other paracetamol products (qualifier value)"
	case AdditionalInstructionCodes17:
		return "Warning. May cause drowsiness. If affected do not drive or operate machinery. Avoid alcoholic drink (qualifier value)"
	case AdditionalInstructionCodes18:
		return "Warning. May cause drowsiness. If affected do not drive or operate machinery (qualifier value)"
	case AdditionalInstructionCodes19:
		return "Sucked or chewed (qualifier value)"
	case AdditionalInstructionCodes20:
		return "Allow to dissolve under the tongue. Do not transfer from this container. Keep tightly closed. Discard eight weeks after opening (qualifier value)"
	case AdditionalInstructionCodes21:
		return "Do not take milk, indigestion remedies, or medicines containing iron or zinc at the same time of day as this medicine (qualifier value)"
	case AdditionalInstructionCodes22:
		return "With plenty of water (qualifier value)"
	case AdditionalInstructionCodes23:
		return "Do not take more than 2 at any one time. Do not take more than 8 in 24 hours (qualifier value)"
	case AdditionalInstructionCodes24:
		return "Caution flammable: keep away from fire or flames (qualifier value)"
	case AdditionalInstructionCodes25:
		return "Do not stop taking this medicine except on your doctor's advice (qualifier value)"
	case AdditionalInstructionCodes26:
		return "Each (qualifier value)"
	case AdditionalInstructionCodes27:
		return "Dissolved under the tongue (qualifier value)"
	case AdditionalInstructionCodes28:
		return "Warning. Avoid alcoholic drink (qualifier value)"
	case AdditionalInstructionCodes29:
		return "This medicine may color the urine"
	case AdditionalInstructionCodes30:
		return "Do not take more than . . . in 24 hours (qualifier value)"
	case AdditionalInstructionCodes31:
		return "Do not take indigestion remedies or medicines containing iron or zinc at the same time of day as this medicine (qualifier value)"
	case AdditionalInstructionCodes32:
		return "Do not take indigestion remedies at the same time of day as this medicine (qualifier value)"
	case AdditionalInstructionCodes33:
		return "To be spread thinly (qualifier value)"
	case AdditionalInstructionCodes34:
		return "Until gone - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes35:
		return "Then discontinue - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes36:
		return "Follow directions (qualifier value)"
	case AdditionalInstructionCodes37:
		return "Until finished - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes38:
		return "Then stop - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes39:
		return "Use with caution (qualifier value)"
	case AdditionalInstructionCodes40:
		return "Take on an empty stomach (qualifier value)"
	default:
		return "Unknown Additional Instruction Codes"
	}
}

/*
Additional dosage instructions (qualifier value)
Half to one hour before food
With or after food
Times (qualifier value)
Contains aspirin (qualifier value)
Dissolve or mix with water before taking (qualifier value)
Warning. Causes drowsiness which may continue the next day. If affected do not drive or operate machinery. Avoid alcoholic drink (qualifier value)
Contains an aspirin-like medicine (qualifier value)
Do not take anything containing aspirin while taking this medicine (qualifier value)
Do not take more than . . . in 24 hours or . . . in any one week (qualifier value)
Avoid exposure of skin to direct sunlight or sun lamps (qualifier value)
Take at regular intervals. Complete the prescribed course unless otherwise directed (qualifier value)
Do not take with any other paracetamol products (qualifier value)
Warning. May cause drowsiness (qualifier value)
Swallowed whole, not chewed (qualifier value)
Warning. Follow the printed instructions you have been given with this medicine (qualifier value)
Contains aspirin and paracetamol. Do not take with any other paracetamol products (qualifier value)
Warning. May cause drowsiness. If affected do not drive or operate machinery. Avoid alcoholic drink (qualifier value)
Warning. May cause drowsiness. If affected do not drive or operate machinery (qualifier value)
Sucked or chewed (qualifier value)
Allow to dissolve under the tongue. Do not transfer from this container. Keep tightly closed. Discard eight weeks after opening (qualifier value)
Do not take milk, indigestion remedies, or medicines containing iron or zinc at the same time of day as this medicine (qualifier value)
With plenty of water (qualifier value)
Do not take more than 2 at any one time. Do not take more than 8 in 24 hours (qualifier value)
Caution flammable: keep away from fire or flames (qualifier value)
Do not stop taking this medicine except on your doctor's advice (qualifier value)
Each (qualifier value)
Dissolved under the tongue (qualifier value)
Warning. Avoid alcoholic drink (qualifier value)
This medicine may color the urine
Do not take more than . . . in 24 hours (qualifier value)
Do not take indigestion remedies or medicines containing iron or zinc at the same time of day as this medicine (qualifier value)
Do not take indigestion remedies at the same time of day as this medicine (qualifier value)
To be spread thinly (qualifier value)
Until gone - dosing instruction fragment (qualifier value)
Then discontinue - dosing instruction fragment (qualifier value)
Follow directions (qualifier value)
Until finished - dosing instruction fragment (qualifier value)
Then stop - dosing instruction fragment (qualifier value)
Use with caution (qualifier value)
Take on an empty stomach (qualifier value)
*/
