package code

import (
	"fmt"
	"strings"
)

type AdverseEventType int

const (
	AdverseEventType00 AdverseEventType = iota
	AdverseEventType01
	AdverseEventType02
	AdverseEventType03
	AdverseEventType04
	AdverseEventType05
	AdverseEventType06
	AdverseEventType07
	AdverseEventType08
	AdverseEventType09
	AdverseEventType10
	AdverseEventType11
	AdverseEventType12
	AdverseEventType13
	AdverseEventType14
	AdverseEventType15
	AdverseEventType16
	AdverseEventType17
	AdverseEventType18
	AdverseEventType19
	AdverseEventType20
	AdverseEventType21
	AdverseEventType22
	AdverseEventType23
	AdverseEventType24
	AdverseEventType25
	AdverseEventType26
	AdverseEventType27
	AdverseEventType28
	AdverseEventType29
	AdverseEventType30
	AdverseEventType31
	AdverseEventType32
	AdverseEventType33
	AdverseEventType34
	AdverseEventType35
	AdverseEventType36
	AdverseEventType37
	AdverseEventType38
	AdverseEventType39
	AdverseEventType40
	AdverseEventType41
	AdverseEventType42
	AdverseEventType43
	AdverseEventType44
	AdverseEventType45
	AdverseEventType46
)

func AllAdverseEventType() []AdverseEventType {
	return []AdverseEventType{
		AdverseEventType00,
		AdverseEventType01,
		AdverseEventType02,
		AdverseEventType03,
		AdverseEventType04,
		AdverseEventType05,
		AdverseEventType06,
		AdverseEventType07,
		AdverseEventType08,
		AdverseEventType09,
		AdverseEventType10,
		AdverseEventType11,
		AdverseEventType12,
		AdverseEventType13,
		AdverseEventType14,
		AdverseEventType15,
		AdverseEventType16,
		AdverseEventType17,
		AdverseEventType18,
		AdverseEventType19,
		AdverseEventType20,
		AdverseEventType21,
		AdverseEventType22,
		AdverseEventType23,
		AdverseEventType24,
		AdverseEventType25,
		AdverseEventType26,
		AdverseEventType27,
		AdverseEventType28,
		AdverseEventType29,
		AdverseEventType30,
		AdverseEventType31,
		AdverseEventType32,
		AdverseEventType33,
		AdverseEventType34,
		AdverseEventType35,
		AdverseEventType36,
		AdverseEventType37,
		AdverseEventType38,
		AdverseEventType39,
		AdverseEventType40,
		AdverseEventType41,
		AdverseEventType42,
		AdverseEventType43,
		AdverseEventType44,
		AdverseEventType45,
		AdverseEventType46,
	}
}

func FindAdverseEventType(filter string) []AdverseEventType {
	ret := make([]AdverseEventType, 0)
	for _, at := range AllAdverseEventType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdverseEventType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdverseEventType) String() string {
	switch cpt {
	case AdverseEventType00:
		return "Serious reportable event (event)"
	case AdverseEventType01:
		return "Postoperative death"
	case AdverseEventType02:
		return "Death during anesthetic induction"
	case AdverseEventType03:
		return "Anesthetic death"
	case AdverseEventType04:
		return "Intraoperative death"
	case AdverseEventType05:
		return "Serious reportable event associated with surgery (event)"
	case AdverseEventType06:
		return "Surgery performed on the wrong body part (event)"
	case AdverseEventType07:
		return "Surgery performed on the wrong patient (event)"
	case AdverseEventType08:
		return "Wrong surgical procedure performed on a patient (event)"
	case AdverseEventType09:
		return "Retention of foreign object in a patient after surgery or other procedure (event)"
	case AdverseEventType10:
		return "Intraoperative or immediately post-operative death in an ASA Class I patient"
	case AdverseEventType11:
		return "Serious reportable event associated with product or device (event)"
	case AdverseEventType12:
		return "Patient death or serious disability associated with the use of contaminated drugs, devices, or biologics provided by the healthcare facility (event)"
	case AdverseEventType13:
		return "Patient death or serious disability associated with the use or function of a device in patient care, in which the device is used or functions other than as intended (event)"
	case AdverseEventType14:
		return "Patient death or serious disability associated with intravascular air embolism that occurs while being cared for in a healthcare facility (event)"
	case AdverseEventType15:
		return "Serious reportable event associated with patient protection (event)"
	case AdverseEventType16:
		return "Infant discharged to the wrong person (event)"
	case AdverseEventType17:
		return "Patient death or serious disability associated with patient elopement (disappearance) for more than four hours (event)"
	case AdverseEventType18:
		return "Patient suicide, or attempted suicide resulting in serious disability, while being cared for in a healthcare facility (event)"
	case AdverseEventType19:
		return "Serious reportable event associated with care management (event)"
	case AdverseEventType20:
		return "Patient death or serious disability associated with a medication error (event)"
	case AdverseEventType21:
		return "Patient death or serious disability associated with a hemolytic reaction due to the administration of ABO-incompatible blood or blood products (event)"
	case AdverseEventType22:
		return "Maternal death or serious disability associated with labor or delivery in a low-risk pregnancy while being cared for in a healthcare facility (event)"
	case AdverseEventType23:
		return "Patient death or serious disability associated with hypoglycemia, the onset of which occurs while the patient is being cared for in a healthcare facility (event)"
	case AdverseEventType24:
		return "Death or serious disability (kernicterus) associated with failure to identify and treat hyperbilirubinemia in neonates (event)"
	case AdverseEventType25:
		return "Stage 3 or 4 pressure ulcers acquired after admission to a healthcare facility (event)"
	case AdverseEventType26:
		return "Patient death or serious disability due to spinal manipulative therapy (event)"
	case AdverseEventType27:
		return "Serious reportable event associated with environment (event)"
	case AdverseEventType28:
		return "Patient death or serious disability associated with an electric shock while being cared for in a healthcare facility (event)"
	case AdverseEventType29:
		return "Any incident in which a line designated for oxygen or other gas to be delivered to a patient contains the wrong gas or is contaminated by toxic substances (event)"
	case AdverseEventType30:
		return "Patient death or serious disability associated with a burn incurred from any source while being cared for in a healthcare facility (event)"
	case AdverseEventType31:
		return "Patient death associated with a fall while being cared for in a healthcare facility (event)"
	case AdverseEventType32:
		return "Patient death or serious disability associated with the use of restraints or bedrails while being cared for in a healthcare facility (event)"
	case AdverseEventType33:
		return "Serious reportable event associated with criminal activity (event)"
	case AdverseEventType34:
		return "Any instance of care ordered by or provided by someone impersonating a physician, nurse, pharmacist, or other licensed healthcare provider (event)"
	case AdverseEventType35:
		return "Abduction of a patient of any age (event)"
	case AdverseEventType36:
		return "Sexual assault on a patient within or on the grounds of the healthcare facility (event)"
	case AdverseEventType37:
		return "Death or significant injury of a patient or staff member resulting from a physical assault (i.e., battery) that occurs within or on the grounds of the healthcare facility (event)"
	case AdverseEventType38:
		return "Staff injury or damage (event)"
	case AdverseEventType39:
		return "Perioperative death"
	case AdverseEventType40:
		return "Critical incident (event)"
	case AdverseEventType41:
		return "Preoperative anesthetic death"
	case AdverseEventType42:
		return "Artificial insemination with the wrong donor sperm or donor egg (event)"
	case AdverseEventType43:
		return "Anesthetic procedure performed on the wrong body part (event)"
	case AdverseEventType44:
		return "Anesthetic procedure performed on the wrong patient (event)"
	case AdverseEventType45:
		return "Transmission of infectious agent via medicinal product"
	case AdverseEventType46:
		return "Fall"
	default:
		return "Unknown Adverse Event Type"
	}
}

/*
Serious reportable event (event)
Postoperative death
Death during anesthetic induction
Anesthetic death
Intraoperative death
Serious reportable event associated with surgery (event)
Surgery performed on the wrong body part (event)
Surgery performed on the wrong patient (event)
Wrong surgical procedure performed on a patient (event)
Retention of foreign object in a patient after surgery or other procedure (event)
Intraoperative or immediately post-operative death in an ASA Class I patient
Serious reportable event associated with product or device (event)
Patient death or serious disability associated with the use of contaminated drugs, devices, or biologics provided by the healthcare facility (event)
Patient death or serious disability associated with the use or function of a device in patient care, in which the device is used or functions other than as intended (event)
Patient death or serious disability associated with intravascular air embolism that occurs while being cared for in a healthcare facility (event)
Serious reportable event associated with patient protection (event)
Infant discharged to the wrong person (event)
Patient death or serious disability associated with patient elopement (disappearance) for more than four hours (event)
Patient suicide, or attempted suicide resulting in serious disability, while being cared for in a healthcare facility (event)
Serious reportable event associated with care management (event)
Patient death or serious disability associated with a medication error (event)
Patient death or serious disability associated with a hemolytic reaction due to the administration of ABO-incompatible blood or blood products (event)
Maternal death or serious disability associated with labor or delivery in a low-risk pregnancy while being cared for in a healthcare facility (event)
Patient death or serious disability associated with hypoglycemia, the onset of which occurs while the patient is being cared for in a healthcare facility (event)
Death or serious disability (kernicterus) associated with failure to identify and treat hyperbilirubinemia in neonates (event)
Stage 3 or 4 pressure ulcers acquired after admission to a healthcare facility (event)
Patient death or serious disability due to spinal manipulative therapy (event)
Serious reportable event associated with environment (event)
Patient death or serious disability associated with an electric shock while being cared for in a healthcare facility (event)
Any incident in which a line designated for oxygen or other gas to be delivered to a patient contains the wrong gas or is contaminated by toxic substances (event)
Patient death or serious disability associated with a burn incurred from any source while being cared for in a healthcare facility (event)
Patient death associated with a fall while being cared for in a healthcare facility (event)
Patient death or serious disability associated with the use of restraints or bedrails while being cared for in a healthcare facility (event)
Serious reportable event associated with criminal activity (event)
Any instance of care ordered by or provided by someone impersonating a physician, nurse, pharmacist, or other licensed healthcare provider (event)
Abduction of a patient of any age (event)
Sexual assault on a patient within or on the grounds of the healthcare facility (event)
Death or significant injury of a patient or staff member resulting from a physical assault (i.e., battery) that occurs within or on the grounds of the healthcare facility (event)
Staff injury or damage (event)
Perioperative death
Critical incident (event)
Preoperative anesthetic death
Artificial insemination with the wrong donor sperm or donor egg (event)
Anesthetic procedure performed on the wrong body part (event)
Anesthetic procedure performed on the wrong patient (event)
Transmission of infectious agent via medicinal product
Fall

*/
