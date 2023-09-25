package code

import (
	"fmt"
	"strings"
)

type AdditionalInstructionCodes int

const (
	AdditionalInstructionCodes00 AdditionalInstructionCodes = iota
	AdditionalInstructionCodes01
	AdditionalInstructionCodes02
	AdditionalInstructionCodes03
	AdditionalInstructionCodes04
	AdditionalInstructionCodes05
	AdditionalInstructionCodes06
	AdditionalInstructionCodes07
	AdditionalInstructionCodes08
	AdditionalInstructionCodes09
	AdditionalInstructionCodes10
	AdditionalInstructionCodes11
	AdditionalInstructionCodes12
	AdditionalInstructionCodes13
	AdditionalInstructionCodes14
	AdditionalInstructionCodes15
	AdditionalInstructionCodes16
	AdditionalInstructionCodes17
	AdditionalInstructionCodes18
	AdditionalInstructionCodes19
	AdditionalInstructionCodes20
	AdditionalInstructionCodes21
	AdditionalInstructionCodes22
	AdditionalInstructionCodes23
	AdditionalInstructionCodes24
	AdditionalInstructionCodes25
	AdditionalInstructionCodes26
	AdditionalInstructionCodes27
	AdditionalInstructionCodes28
	AdditionalInstructionCodes29
	AdditionalInstructionCodes30
	AdditionalInstructionCodes31
	AdditionalInstructionCodes32
	AdditionalInstructionCodes33
	AdditionalInstructionCodes34
	AdditionalInstructionCodes35
	AdditionalInstructionCodes36
	AdditionalInstructionCodes37
	AdditionalInstructionCodes38
	AdditionalInstructionCodes39
	AdditionalInstructionCodes40
	AdditionalInstructionCodes41
	AdditionalInstructionCodes42
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
		AdditionalInstructionCodes41,
		AdditionalInstructionCodes42,
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
	case AdditionalInstructionCodes01:
		return "Additional dosage instructions (qualifier value)"
	case AdditionalInstructionCodes02:
		return "Half to one hour before food"
	case AdditionalInstructionCodes04:
		return "With or after food"
	case AdditionalInstructionCodes05:
		return "Times (qualifier value)"
	case AdditionalInstructionCodes06:
		return "Contains aspirin (qualifier value)"
	case AdditionalInstructionCodes07:
		return "Dissolve or mix with water before taking (qualifier value)"
	case AdditionalInstructionCodes08:
		return "Warning. Causes drowsiness which may continue the next day. If affected do not drive or operate machinery. Avoid alcoholic drink (qualifier value)"
	case AdditionalInstructionCodes09:
		return "Contains an aspirin-like medicine (qualifier value)"
	case AdditionalInstructionCodes10:
		return "Do not take anything containing aspirin while taking this medicine (qualifier value)"
	case AdditionalInstructionCodes11:
		return "Do not take more than . . . in 24 hours or . . . in any one week (qualifier value)"
	case AdditionalInstructionCodes12:
		return "Avoid exposure of skin to direct sunlight or sun lamps (qualifier value)"
	case AdditionalInstructionCodes13:
		return "Take at regular intervals. Complete the prescribed course unless otherwise directed (qualifier value)"
	case AdditionalInstructionCodes14:
		return "Do not take with any other paracetamol products (qualifier value)"
	case AdditionalInstructionCodes15:
		return "Warning. May cause drowsiness (qualifier value)"
	case AdditionalInstructionCodes16:
		return "Swallowed whole, not chewed (qualifier value)"
	case AdditionalInstructionCodes17:
		return "Warning. Follow the printed instructions you have been given with this medicine (qualifier value)"
	case AdditionalInstructionCodes18:
		return "Contains aspirin and paracetamol. Do not take with any other paracetamol products (qualifier value)"
	case AdditionalInstructionCodes19:
		return "Warning. May cause drowsiness. If affected do not drive or operate machinery. Avoid alcoholic drink (qualifier value)"
	case AdditionalInstructionCodes20:
		return "Warning. May cause drowsiness. If affected do not drive or operate machinery (qualifier value)"
	case AdditionalInstructionCodes21:
		return "Sucked or chewed (qualifier value)"
	case AdditionalInstructionCodes22:
		return "Allow to dissolve under the tongue. Do not transfer from this container. Keep tightly closed. Discard eight weeks after opening (qualifier value)"
	case AdditionalInstructionCodes23:
		return "Do not take milk, indigestion remedies, or medicines containing iron or zinc at the same time of day as this medicine (qualifier value)"
	case AdditionalInstructionCodes24:
		return "With plenty of water (qualifier value)"
	case AdditionalInstructionCodes25:
		return "Do not take more than 2 at any one time. Do not take more than 8 in 24 hours (qualifier value)"
	case AdditionalInstructionCodes26:
		return "Caution flammable: keep away from fire or flames (qualifier value)"
	case AdditionalInstructionCodes27:
		return "Do not stop taking this medicine except on your doctor's advice (qualifier value)"
	case AdditionalInstructionCodes28:
		return "Each (qualifier value)"
	case AdditionalInstructionCodes29:
		return "Dissolved under the tongue (qualifier value)"
	case AdditionalInstructionCodes30:
		return "Warning. Avoid alcoholic drink (qualifier value)"
	case AdditionalInstructionCodes31:
		return "This medicine may color the urine"
	case AdditionalInstructionCodes32:
		return "Do not take more than . . . in 24 hours (qualifier value)"
	case AdditionalInstructionCodes33:
		return "Do not take indigestion remedies or medicines containing iron or zinc at the same time of day as this medicine (qualifier value)"
	case AdditionalInstructionCodes34:
		return "Do not take indigestion remedies at the same time of day as this medicine (qualifier value)"
	case AdditionalInstructionCodes35:
		return "To be spread thinly (qualifier value)"
	case AdditionalInstructionCodes36:
		return "Until gone - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes37:
		return "Then discontinue - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes38:
		return "Follow directions (qualifier value)"
	case AdditionalInstructionCodes39:
		return "Until finished - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes40:
		return "Then stop - dosing instruction fragment (qualifier value)"
	case AdditionalInstructionCodes41:
		return "Use with caution (qualifier value)"
	case AdditionalInstructionCodes42:
		return "Take on an empty stomach (qualifier value)"
	default:
		return "Unknown Qualification"
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
