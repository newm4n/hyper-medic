package code

import (
	"fmt"
	"strings"
)

type Language int

const (
	AB Language = iota
	AA
	AF
	AK
	SQ
	AM
	AR
	AN
	HY
	AS
	AV
	AE
	AY
	AZ
	BM
	BA
	EU
	BE
	BN
	BH
	BI
	BS
	BR
	BG
	MY
	CA
	CH
	CE
	NY
	ZH
	ZH_HANS
	ZH_HANT
	CV
	KW
	CO
	CR
	HR
	CS
	DA
	DV
	NL
	DZ
	EN
	EO
	ET
	EE
	FO
	FJ
	FI
	FR
	FF
	GL
	GD
	GV
	KA
	DE
	EL
	KL
	GN
	GU
	HT
	HA
	HE
	HZ
	HI
	HO
	HU
	IS
	IO
	IG
	ID
	IA
	IE
	IU
	IK
	GA
	IT
	JA
	JV
	KN
	KR
	KS
	KK
	KM
	KI
	RW
	RN
	KY
	KV
	KG
	KO
	KU
	KJ
	LO
	LA
	LV
	LI
	LN
	LT
	LU
	LG
	LB
	MK
	MG
	MS
	ML
	MT
	MI
	MR
	MH
	MO
	MN
	NA
	NV
	NG
	ND
	NE
	NO
	NB
	NN
	II
	OC
	OJ
	CU
	OR
	OM
	OS
	PI
	PS
	FA
	PL
	PT
	PA
	QU
	RM
	RO
	RU
	SE
	SM
	SG
	SA
	SR
	SH
	ST
	TN
	SN
	SD
	SI
	SS
	SK
	SL
	SO
	NR
	ES
	SU
	SW
	SV
	TL
	TY
	TG
	TA
	TT
	TE
	TH
	BO
	TI
	TO
	TS
	TR
	TK
	TW
	UG
	UK
	UR
	UZ
	VE
	VI
	VO
	WA
	CY
	WO
	FY
	XH
	YI
	YO
	ZA
	ZU
)

func AllLanguage() []Language {
	return []Language{
		AB,
		AA,
		AF,
		AK,
		SQ,
		AM,
		AR,
		AN,
		HY,
		AS,
		AV,
		AE,
		AY,
		AZ,
		BM,
		BA,
		EU,
		BE,
		BN,
		BH,
		BI,
		BS,
		BR,
		BG,
		MY,
		CA,
		CH,
		CE,
		NY,
		ZH,
		ZH_HANS,
		ZH_HANT,
		CV,
		KW,
		CO,
		CR,
		HR,
		CS,
		DA,
		DV,
		NL,
		DZ,
		EN,
		EO,
		ET,
		EE,
		FO,
		FJ,
		FI,
		FR,
		FF,
		GL,
		GD,
		GV,
		KA,
		DE,
		EL,
		KL,
		GN,
		GU,
		HT,
		HA,
		HE,
		HZ,
		HI,
		HO,
		HU,
		IS,
		IO,
		IG,
		ID,
		IA,
		IE,
		IU,
		IK,
		GA,
		IT,
		JA,
		JV,
		KN,
		KR,
		KS,
		KK,
		KM,
		KI,
		RW,
		RN,
		KY,
		KV,
		KG,
		KO,
		KU,
		KJ,
		LO,
		LA,
		LV,
		LI,
		LN,
		LT,
		LU,
		LG,
		LB,
		MK,
		MG,
		MS,
		ML,
		MT,
		MI,
		MR,
		MH,
		MO,
		MN,
		NA,
		NV,
		NG,
		ND,
		NE,
		NO,
		NB,
		NN,
		II,
		OC,
		OJ,
		CU,
		OR,
		OM,
		OS,
		PI,
		PS,
		FA,
		PL,
		PT,
		PA,
		QU,
		RM,
		RO,
		RU,
		SE,
		SM,
		SG,
		SA,
		SR,
		SH,
		ST,
		TN,
		SN,
		SD,
		SI,
		SS,
		SK,
		SL,
		SO,
		NR,
		ES,
		SU,
		SW,
		SV,
		TL,
		TY,
		TG,
		TA,
		TT,
		TE,
		TH,
		BO,
		TI,
		TO,
		TS,
		TR,
		TK,
		TW,
		UG,
		UK,
		UR,
		UZ,
		VE,
		VI,
		VO,
		WA,
		CY,
		WO,
		FY,
		XH,
		YI,
		YO,
		ZA,
		ZU,
	}
}

func FindLanguage(filter string) []Language {
	ret := make([]Language, 0)
	for _, at := range AllLanguage() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt Language) ToString() {
	fmt.Println(cpt.String())
}

func (cpt Language) String() string {
	switch cpt {
	case AB:
		return "Abkhazian"
	case AA:
		return "Afar"
	case AF:
		return "Afrikaans"
	case AK:
		return "Akan"
	case SQ:
		return "Albanian"
	case AM:
		return "Amharic"
	case AR:
		return "Arabic"
	case AN:
		return "Aragonese"
	case HY:
		return "Armenian"
	case AS:
		return "Assamese"
	case AV:
		return "Avaric"
	case AE:
		return "Avestan"
	case AY:
		return "Aymara"
	case AZ:
		return "Azerbaijani"
	case BM:
		return "Bambara"
	case BA:
		return "Bashkir"
	case EU:
		return "Basque"
	case BE:
		return "Belarusian"
	case BN:
		return "Bengali (Bangla)"
	case BH:
		return "Bihari"
	case BI:
		return "Bislama"
	case BS:
		return "Bosnian"
	case BR:
		return "Breton"
	case BG:
		return "Bulgarian"
	case MY:
		return "Burmese"
	case CA:
		return "Catalan"
	case CH:
		return "Chamorro"
	case CE:
		return "Chechen"
	case NY:
		return "Chichewa, Chewa, Nyanja"
	case ZH:
		return "Chinese"
	case ZH_HANS:
		return "Chinese (Simplified)"
	case ZH_HANT:
		return "Chinese (Traditional)"
	case CV:
		return "Chuvash"
	case KW:
		return "Cornish"
	case CO:
		return "Corsican"
	case CR:
		return "Cree"
	case HR:
		return "Croatian"
	case CS:
		return "Czech"
	case DA:
		return "Danish"
	case DV:
		return "Divehi, Dhivehi, Maldivian"
	case NL:
		return "Dutch"
	case DZ:
		return "Dzongkha"
	case EN:
		return "English"
	case EO:
		return "Esperanto"
	case ET:
		return "Estonian"
	case EE:
		return "Ewe"
	case FO:
		return "Faroese"
	case FJ:
		return "Fijian"
	case FI:
		return "Finnish"
	case FR:
		return "French"
	case FF:
		return "Fula, Fulah, Pulaar, Pular"
	case GL:
		return "Galician"
	case GD:
		return "Gaelic (Scottish)"
	case GV:
		return "Gaelic (Manx)"
	case KA:
		return "Georgian"
	case DE:
		return "German"
	case EL:
		return "Greek"
	case KL:
		return "Greenlandic"
	case GN:
		return "Guarani"
	case GU:
		return "Gujarati"
	case HT:
		return "Haitian Creole"
	case HA:
		return "Hausa"
	case HE:
		return "Hebrew"
	case HZ:
		return "Herero"
	case HI:
		return "Hindi"
	case HO:
		return "Hiri Motu"
	case HU:
		return "Hungarian"
	case IS:
		return "Icelandic"
	case IO:
		return "Ido"
	case IG:
		return "Igbo"
	case ID:
		return "Indonesian"
	case IA:
		return "Interlingua"
	case IE:
		return "Interlingue"
	case IU:
		return "Inuktitut"
	case IK:
		return "Inupiak"
	case GA:
		return "Irish"
	case IT:
		return "Italian"
	case JA:
		return "Japanese"
	case JV:
		return "Javanese"
	case KN:
		return "Kannada"
	case KR:
		return "Kanuri"
	case KS:
		return "Kashmiri"
	case KK:
		return "Kazakh"
	case KM:
		return "Khmer"
	case KI:
		return "Kikuyu"
	case RW:
		return "Kinyarwanda (Rwanda)"
	case RN:
		return "Kirundi"
	case KY:
		return "Kyrgyz"
	case KV:
		return "Komi"
	case KG:
		return "Kongo"
	case KO:
		return "Korean"
	case KU:
		return "Kurdish"
	case KJ:
		return "Kwanyama"
	case LO:
		return "Lao"
	case LA:
		return "Latin"
	case LV:
		return "Latvian (Lettish)"
	case LI:
		return "Limburgish ( Limburger)"
	case LN:
		return "Lingala"
	case LT:
		return "Lithuanian"
	case LU:
		return "Luga-Katanga"
	case LG:
		return "Luganda, Ganda"
	case LB:
		return "Luxembourgish"
	case MK:
		return "Macedonian"
	case MG:
		return "Malagasy"
	case MS:
		return "Malay"
	case ML:
		return "Malayalam"
	case MT:
		return "Maltese"
	case MI:
		return "Maori"
	case MR:
		return "Marathi"
	case MH:
		return "Marshallese"
	case MO:
		return "Moldavian"
	case MN:
		return "Mongolian"
	case NA:
		return "Nauru"
	case NV:
		return "Navajo"
	case NG:
		return "Ndonga"
	case ND:
		return "Northern Ndebele"
	case NE:
		return "Nepali"
	case NO:
		return "Norwegian"
	case NB:
		return "Norwegian bokmål"
	case NN:
		return "Norwegian nynorsk"
	case II:
		return "Nuosu"
	case OC:
		return "Occitan"
	case OJ:
		return "Ojibwe"
	case CU:
		return "Old Church Slavonic, Old Bulgarian"
	case OR:
		return "Oriya"
	case OM:
		return "Oromo (Afaan Oromo)"
	case OS:
		return "Ossetian"
	case PI:
		return "Pāli"
	case PS:
		return "Pashto, Pushto"
	case FA:
		return "Persian (Farsi)"
	case PL:
		return "Polish"
	case PT:
		return "Portuguese"
	case PA:
		return "Punjabi (Eastern)"
	case QU:
		return "Quechua"
	case RM:
		return "Romansh"
	case RO:
		return "Romanian"
	case RU:
		return "Russian"
	case SE:
		return "Sami"
	case SM:
		return "Samoan"
	case SG:
		return "Sango"
	case SA:
		return "Sanskrit"
	case SR:
		return "Serbian"
	case SH:
		return "Serbo-Croatian"
	case ST:
		return "Sesotho"
	case TN:
		return "Setswana"
	case SN:
		return "Shona"
	case SD:
		return "Sindhi"
	case SI:
		return "Sinhalese"
	case SS:
		return "Siswati"
	case SK:
		return "Slovak"
	case SL:
		return "Slovenian"
	case SO:
		return "Somali"
	case NR:
		return "Southern Ndebele"
	case ES:
		return "Spanish"
	case SU:
		return "Sundanese"
	case SW:
		return "Swahili (Kiswahili)"
	case SV:
		return "Swedish"
	case TL:
		return "Tagalog"
	case TY:
		return "Tahitian"
	case TG:
		return "Tajik"
	case TA:
		return "Tamil"
	case TT:
		return "Tatar"
	case TE:
		return "Telugu"
	case TH:
		return "Thai"
	case BO:
		return "Tibetan"
	case TI:
		return "Tigrinya"
	case TO:
		return "Tonga"
	case TS:
		return "Tsonga"
	case TR:
		return "Turkish"
	case TK:
		return "Turkmen"
	case TW:
		return "Twi"
	case UG:
		return "Uyghur"
	case UK:
		return "Ukrainian"
	case UR:
		return "Urdu"
	case UZ:
		return "Uzbek"
	case VE:
		return "Venda"
	case VI:
		return "Vietnamese"
	case VO:
		return "Volapük"
	case WA:
		return "Wallon"
	case CY:
		return "Welsh"
	case WO:
		return "Wolof"
	case FY:
		return "Western Frisian"
	case XH:
		return "Xhosa"
	case YI:
		return "Yiddish"
	case YO:
		return "Yoruba"
	case ZA:
		return "Zhuang, Chuang"
	case ZU:
		return "Zulu"

	default:
		return "Unknown Language"
	}
}
