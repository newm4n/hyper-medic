package code

import (
	"fmt"
	"strings"
)

type Languages string

const (
	LanguagesAR   Languages = "ar"
	LanguagesBG   Languages = "bg"
	LanguagesBGBG Languages = "bg-BG"
	LanguagesBN   Languages = "bn"
	LanguagesCS   Languages = "cs"
	LanguagesCSCZ Languages = "cs-CZ"
	LanguagesBS   Languages = "bs"
	LanguagesBSBA Languages = "bs-BA"
	LanguagesDA   Languages = "da"
	LanguagesDADK Languages = "da-DK"
	LanguagesDE   Languages = "de"
	LanguagesDEAT Languages = "de-AT"
	LanguagesDECH Languages = "de-CH"
	LanguagesDEDE Languages = "de-DE"
	LanguagesEL   Languages = "el"
	LanguagesELGR Languages = "el-GR"
	LanguagesEN   Languages = "en"
	LanguagesENAU Languages = "en-AU"
	LanguagesENCA Languages = "en-CA"
	LanguagesENGB Languages = "en-GB"
	LanguagesENIN Languages = "en-IN"
	LanguagesENNZ Languages = "en-NZ"
	LanguagesENSG Languages = "en-SG"
	LanguagesENUS Languages = "en-US"
	LanguagesES   Languages = "es"
	LanguagesESAR Languages = "es-AR"
	LanguagesESES Languages = "es-ES"
	LanguagesESUY Languages = "es-UY"
	LanguagesET   Languages = "et"
	LanguagesETEE Languages = "et-EE"
	LanguagesFI   Languages = "fi"
	LanguagesFR   Languages = "fr"
	LanguagesFRBE Languages = "fr-BE"
	LanguagesFRCH Languages = "fr-CH"
	LanguagesFRFR Languages = "fr-FR"
	LanguagesFIFI Languages = "fi-FI"
	LanguagesFRCA Languages = "fr-CA"
	LanguagesFY   Languages = "fy"
	LanguagesFYNL Languages = "fy-NL"
	LanguagesHI   Languages = "hi"
	LanguagesHR   Languages = "hr"
	LanguagesHRHR Languages = "hr-HR"
	LanguagesIS   Languages = "is"
	LanguagesISIS Languages = "is-IS"
	LanguagesIT   Languages = "it"
	LanguagesITCH Languages = "it-CH"
	LanguagesITIT Languages = "it-IT"
	LanguagesJA   Languages = "ja"
	LanguagesKO   Languages = "ko"
	LanguagesLT   Languages = "lt"
	LanguagesLTLT Languages = "lt-LT"
	LanguagesLV   Languages = "lv"
	LanguagesLVLV Languages = "lv-LV"
	LanguagesNL   Languages = "nl"
	LanguagesNLBE Languages = "nl-BE"
	LanguagesNLNL Languages = "nl-NL"
	LanguagesNO   Languages = "no"
	LanguagesNONO Languages = "no-NO"
	LanguagesPA   Languages = "pa"
	LanguagesPL   Languages = "pl"
	LanguagesPLPL Languages = "pl-PL"
	LanguagesPT   Languages = "pt"
	LanguagesPTPT Languages = "pt-PT"
	LanguagesPTBR Languages = "pt-BR"
	LanguagesRO   Languages = "ro"
	LanguagesRORO Languages = "ro-RO"
	LanguagesRU   Languages = "ru"
	LanguagesRURU Languages = "ru-RU"
	LanguagesSK   Languages = "sk"
	LanguagesSKSK Languages = "sk-SK"
	LanguagesSL   Languages = "sl"
	LanguagesSLSI Languages = "sl-SI"
	LanguagesSR   Languages = "sr"
	LanguagesSRRS Languages = "sr-RS"
	LanguagesSV   Languages = "sv"
	LanguagesSVSE Languages = "sv-SE"
	LanguagesTE   Languages = "te"
	LanguagesZH   Languages = "zh"
	LanguagesZHCN Languages = "zh-CN"
	LanguagesZHHK Languages = "zh-HK"
	LanguagesZHSG Languages = "zh-SG"
	LanguagesZHTW Languages = "zh-TW"
)

func AllLanguages() []Languages {
	return []Languages{
		LanguagesAR,
		LanguagesBG,
		LanguagesBGBG,
		LanguagesBN,
		LanguagesCS,
		LanguagesCSCZ,
		LanguagesBS,
		LanguagesBSBA,
		LanguagesDA,
		LanguagesDADK,
		LanguagesDE,
		LanguagesDEAT,
		LanguagesDECH,
		LanguagesDEDE,
		LanguagesEL,
		LanguagesELGR,
		LanguagesEN,
		LanguagesENAU,
		LanguagesENCA,
		LanguagesENGB,
		LanguagesENIN,
		LanguagesENNZ,
		LanguagesENSG,
		LanguagesENUS,
		LanguagesES,
		LanguagesESAR,
		LanguagesESES,
		LanguagesESUY,
		LanguagesET,
		LanguagesETEE,
		LanguagesFI,
		LanguagesFR,
		LanguagesFRBE,
		LanguagesFRCH,
		LanguagesFRFR,
		LanguagesFIFI,
		LanguagesFRCA,
		LanguagesFY,
		LanguagesFYNL,
		LanguagesHI,
		LanguagesHR,
		LanguagesHRHR,
		LanguagesIS,
		LanguagesISIS,
		LanguagesIT,
		LanguagesITCH,
		LanguagesITIT,
		LanguagesJA,
		LanguagesKO,
		LanguagesLT,
		LanguagesLTLT,
		LanguagesLV,
		LanguagesLVLV,
		LanguagesNL,
		LanguagesNLBE,
		LanguagesNLNL,
		LanguagesNO,
		LanguagesNONO,
		LanguagesPA,
		LanguagesPL,
		LanguagesPLPL,
		LanguagesPT,
		LanguagesPTPT,
		LanguagesPTBR,
		LanguagesRO,
		LanguagesRORO,
		LanguagesRU,
		LanguagesRURU,
		LanguagesSK,
		LanguagesSKSK,
		LanguagesSL,
		LanguagesSLSI,
		LanguagesSR,
		LanguagesSRRS,
		LanguagesSV,
		LanguagesSVSE,
		LanguagesTE,
		LanguagesZH,
		LanguagesZHCN,
		LanguagesZHHK,
		LanguagesZHSG,
		LanguagesZHTW,
	}
}

func FindLanguages(filter string) []Languages {
	ret := make([]Languages, 0)
	for _, at := range AllLanguages() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au Languages) ToString() {
	fmt.Println(au.String())
}
func (au Languages) String() string {
	switch au {
	case LanguagesAR:
		return "Arabisk"
	case LanguagesBG:
		return "Bulgarian"
	case LanguagesBGBG:
		return "Bulgarian (Bulgaria)"
	case LanguagesBN:
		return "Bengali"
	case LanguagesCS:
		return "Czech"
	case LanguagesCSCZ:
		return "Czech (Czechia)"
	case LanguagesBS:
		return "Bosnian"
	case LanguagesBSBA:
		return "Bosnian (Bosnia and Herzegovina))"
	case LanguagesDA:
		return "Danish"
	case LanguagesDADK:
		return "Danish (Denmark)"
	case LanguagesDE:
		return "German"
	case LanguagesDEAT:
		return "German (Austria)"
	case LanguagesDECH:
		return "German (Switzerland)"
	case LanguagesDEDE:
		return "German (Germany)"
	case LanguagesEL:
		return "Greek"
	case LanguagesELGR:
		return "Greek (Greece)"
	case LanguagesEN:
		return "English"
	case LanguagesENAU:
		return "English (Australia)"
	case LanguagesENCA:
		return "English (Canada)"
	case LanguagesENGB:
		return "English (Great Britain)"
	case LanguagesENIN:
		return "English (India)"
	case LanguagesENNZ:
		return "English (New Zealand)"
	case LanguagesENSG:
		return "English (Singapore)"
	case LanguagesENUS:
		return "English (United States)"
	case LanguagesES:
		return "Spanish"
	case LanguagesESAR:
		return "Spanish (Argentina)"
	case LanguagesESES:
		return "Spanish (Spain)"
	case LanguagesESUY:
		return "Spanish (Uruguay)"
	case LanguagesET:
		return "Estonian"
	case LanguagesETEE:
		return "Estonian (Estonia)"
	case LanguagesFI:
		return "Finnish"
	case LanguagesFR:
		return "French"
	case LanguagesFRBE:
		return "French (Belgium)"
	case LanguagesFRCH:
		return "French (Switzerland)"
	case LanguagesFRFR:
		return "French (France)"
	case LanguagesFIFI:
		return "Finnish (Finland)"
	case LanguagesFRCA:
		return "French (Canada)"
	case LanguagesFY:
		return "Frisian"
	case LanguagesFYNL:
		return "Frisian (Netherlands)"
	case LanguagesHI:
		return "Hindi"
	case LanguagesHR:
		return "Croatian"
	case LanguagesHRHR:
		return "Croatian (Croatia)"
	case LanguagesIS:
		return "Icelandic"
	case LanguagesISIS:
		return "Icelandic (Iceland)"
	case LanguagesIT:
		return "Italian"
	case LanguagesITCH:
		return "Italian (Switzerland)"
	case LanguagesITIT:
		return "Italian (Italy)"
	case LanguagesJA:
		return "Japanese"
	case LanguagesKO:
		return "Korean"
	case LanguagesLT:
		return "Lithuanian"
	case LanguagesLTLT:
		return "Lithuanian (Lithuania)"
	case LanguagesLV:
		return "Latvian"
	case LanguagesLVLV:
		return "Latvian (Latvia)"
	case LanguagesNL:
		return "Dutch"
	case LanguagesNLBE:
		return "Dutch (Belgium)"
	case LanguagesNLNL:
		return "Dutch (Netherlands)"
	case LanguagesNO:
		return "Norwegian"
	case LanguagesNONO:
		return "Norwegian (Norway)"
	case LanguagesPA:
		return "Punjabi"
	case LanguagesPL:
		return "Polskie"
	case LanguagesPLPL:
		return "Polish (Poland)"
	case LanguagesPT:
		return "Portuguese"
	case LanguagesPTPT:
		return "Portuguese (Portugal)"
	case LanguagesPTBR:
		return "Portuguese (Brazil)"
	case LanguagesRO:
		return "Romanian"
	case LanguagesRORO:
		return "Romanian (Romania)"
	case LanguagesRU:
		return "Russian"
	case LanguagesRURU:
		return "Russian (Russia)"
	case LanguagesSK:
		return "Slovakian"
	case LanguagesSKSK:
		return "Slovakian (Slovakia)"
	case LanguagesSL:
		return "Slovenian"
	case LanguagesSLSI:
		return "Slovenian (Slovenia)"
	case LanguagesSR:
		return "Serbian"
	case LanguagesSRRS:
		return "Serbian (Serbia)"
	case LanguagesSV:
		return "Swedish"
	case LanguagesSVSE:
		return "Swedish (Sweden)"
	case LanguagesTE:
		return "Telugu"
	case LanguagesZH:
		return "Chinese"
	case LanguagesZHCN:
		return "Chinese (China)"
	case LanguagesZHHK:
		return "Chinese (Hong Kong)"
	case LanguagesZHSG:
		return "Chinese (Singapore)"
	case LanguagesZHTW:
		return "Chinese (Taiwan)"

	default:
		return "Unknown Constraint Severity"
	}
}
