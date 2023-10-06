package code

import (
	"fmt"
	"strings"
)

type BenefitType string

const (
	BenefitTypeBenefit            BenefitType = "benefit"
	BenefitTypeDeductible         BenefitType = "deductible"
	BenefitTypeVisit              BenefitType = "visit"
	BenefitTypeRoom               BenefitType = "room"
	BenefitTypeCopay              BenefitType = "copay"
	BenefitTypeCopayPercent       BenefitType = "copay-percent"
	BenefitTypeCopayMaximum       BenefitType = "copay-maximum"
	BenefitTypeVisionExam         BenefitType = "vision-exam"
	BenefitTypeVisionGlasses      BenefitType = "vision-glasses"
	BenefitTypeVisionContacts     BenefitType = "vision-contacts"
	BenefitTypeMedicalPrimarycare BenefitType = "medical-primarycare"
	BenefitTypePharmacyDispense   BenefitType = "pharmacy-dispense"
)

func AllBenefitType() []BenefitType {
	return []BenefitType{
		BenefitTypeBenefit,
		BenefitTypeDeductible,
		BenefitTypeVisit,
		BenefitTypeRoom,
		BenefitTypeCopay,
		BenefitTypeCopayPercent,
		BenefitTypeCopayMaximum,
		BenefitTypeVisionExam,
		BenefitTypeVisionGlasses,
		BenefitTypeVisionContacts,
		BenefitTypeMedicalPrimarycare,
		BenefitTypePharmacyDispense,
	}
}

func FindBenefitType(filter string) []BenefitType {
	ret := make([]BenefitType, 0)
	for _, at := range AllBenefitType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt BenefitType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt BenefitType) String() string {
	switch cpt {
	case BenefitTypeBenefit:
		return "Benefit"
	case BenefitTypeDeductible:
		return "Deductible"
	case BenefitTypeVisit:
		return "Visit"
	case BenefitTypeRoom:
		return "Room"
	case BenefitTypeCopay:
		return "Copayment per service"
	case BenefitTypeCopayPercent:
		return "Copayment Percent per service"
	case BenefitTypeCopayMaximum:
		return "Copayment maximum per service"
	case BenefitTypeVisionExam:
		return "Vision Exam"
	case BenefitTypeVisionGlasses:
		return "Vision Glasses"
	case BenefitTypeVisionContacts:
		return "Vision Contacts Coverage"
	case BenefitTypeMedicalPrimarycare:
		return "Medical Primary Health Coverage"
	case BenefitTypePharmacyDispense:
		return "Pharmacy Dispense Coverage"

	default:
		return "Unknown Benefit Type"
	}
}
