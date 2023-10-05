package code

import (
	"fmt"
	"strings"
)

type Currencies string

const (
	CurrenciesAED Currencies = "AED"
	CurrenciesAFN Currencies = "AFN"
	CurrenciesALL Currencies = "ALL"
	CurrenciesAMD Currencies = "AMD"
	CurrenciesANG Currencies = "ANG"
	CurrenciesAOA Currencies = "AOA"
	CurrenciesARS Currencies = "ARS"
	CurrenciesAUD Currencies = "AUD"
	CurrenciesAWG Currencies = "AWG"
	CurrenciesAZN Currencies = "AZN"
	CurrenciesBAM Currencies = "BAM"
	CurrenciesBBD Currencies = "BBD"
	CurrenciesBDT Currencies = "BDT"
	CurrenciesBGN Currencies = "BGN"
	CurrenciesBHD Currencies = "BHD"
	CurrenciesBIF Currencies = "BIF"
	CurrenciesBMD Currencies = "BMD"
	CurrenciesBND Currencies = "BND"
	CurrenciesBOB Currencies = "BOB"
	CurrenciesBOV Currencies = "BOV"
	CurrenciesBRL Currencies = "BRL"
	CurrenciesBSD Currencies = "BSD"
	CurrenciesBTN Currencies = "BTN"
	CurrenciesBWP Currencies = "BWP"
	CurrenciesBYN Currencies = "BYN"
	CurrenciesBZD Currencies = "BZD"
	CurrenciesCAD Currencies = "CAD"
	CurrenciesCDF Currencies = "CDF"
	CurrenciesCHE Currencies = "CHE"
	CurrenciesCHF Currencies = "CHF"
	CurrenciesCHW Currencies = "CHW"
	CurrenciesCLF Currencies = "CLF"
	CurrenciesCLP Currencies = "CLP"
	CurrenciesCNY Currencies = "CNY"
	CurrenciesCOP Currencies = "COP"
	CurrenciesCOU Currencies = "COU"
	CurrenciesCRC Currencies = "CRC"
	CurrenciesCUC Currencies = "CUC"
	CurrenciesCUP Currencies = "CUP"
	CurrenciesCVE Currencies = "CVE"
	CurrenciesCZK Currencies = "CZK"
	CurrenciesDJF Currencies = "DJF"
	CurrenciesDKK Currencies = "DKK"
	CurrenciesDOP Currencies = "DOP"
	CurrenciesDZD Currencies = "DZD"
	CurrenciesEGP Currencies = "EGP"
	CurrenciesERN Currencies = "ERN"
	CurrenciesETB Currencies = "ETB"
	CurrenciesEUR Currencies = "EUR"
	CurrenciesFJD Currencies = "FJD"
	CurrenciesFKP Currencies = "FKP"
	CurrenciesGBP Currencies = "GBP"
	CurrenciesGEL Currencies = "GEL"
	CurrenciesGGP Currencies = "GGP"
	CurrenciesGHS Currencies = "GHS"
	CurrenciesGIP Currencies = "GIP"
	CurrenciesGMD Currencies = "GMD"
	CurrenciesGNF Currencies = "GNF"
	CurrenciesGTQ Currencies = "GTQ"
	CurrenciesGYD Currencies = "GYD"
	CurrenciesHKD Currencies = "HKD"
	CurrenciesHNL Currencies = "HNL"
	CurrenciesHRK Currencies = "HRK"
	CurrenciesHTG Currencies = "HTG"
	CurrenciesHUF Currencies = "HUF"
	CurrenciesIDR Currencies = "IDR"
	CurrenciesILS Currencies = "ILS"
	CurrenciesIMP Currencies = "IMP"
	CurrenciesINR Currencies = "INR"
	CurrenciesIQD Currencies = "IQD"
	CurrenciesIRR Currencies = "IRR"
	CurrenciesISK Currencies = "ISK"
	CurrenciesJEP Currencies = "JEP"
	CurrenciesJMD Currencies = "JMD"
	CurrenciesJOD Currencies = "JOD"
	CurrenciesJPY Currencies = "JPY"
	CurrenciesKES Currencies = "KES"
	CurrenciesKGS Currencies = "KGS"
	CurrenciesKHR Currencies = "KHR"
	CurrenciesKMF Currencies = "KMF"
	CurrenciesKPW Currencies = "KPW"
	CurrenciesKRW Currencies = "KRW"
	CurrenciesKWD Currencies = "KWD"
	CurrenciesKYD Currencies = "KYD"
	CurrenciesKZT Currencies = "KZT"
	CurrenciesLAK Currencies = "LAK"
	CurrenciesLBP Currencies = "LBP"
	CurrenciesLKR Currencies = "LKR"
	CurrenciesLRD Currencies = "LRD"
	CurrenciesLSL Currencies = "LSL"
	CurrenciesLYD Currencies = "LYD"
	CurrenciesMAD Currencies = "MAD"
	CurrenciesMDL Currencies = "MDL"
	CurrenciesMGA Currencies = "MGA"
	CurrenciesMKD Currencies = "MKD"
	CurrenciesMMK Currencies = "MMK"
	CurrenciesMNT Currencies = "MNT"
	CurrenciesMOP Currencies = "MOP"
	CurrenciesMRU Currencies = "MRU"
	CurrenciesMUR Currencies = "MUR"
	CurrenciesMVR Currencies = "MVR"
	CurrenciesMWK Currencies = "MWK"
	CurrenciesMXN Currencies = "MXN"
	CurrenciesMXV Currencies = "MXV"
	CurrenciesMYR Currencies = "MYR"
	CurrenciesMZN Currencies = "MZN"
	CurrenciesNAD Currencies = "NAD"
	CurrenciesNGN Currencies = "NGN"
	CurrenciesNIO Currencies = "NIO"
	CurrenciesNOK Currencies = "NOK"
	CurrenciesNPR Currencies = "NPR"
	CurrenciesNZD Currencies = "NZD"
	CurrenciesOMR Currencies = "OMR"
	CurrenciesPAB Currencies = "PAB"
	CurrenciesPEN Currencies = "PEN"
	CurrenciesPGK Currencies = "PGK"
	CurrenciesPHP Currencies = "PHP"
	CurrenciesPKR Currencies = "PKR"
	CurrenciesPLN Currencies = "PLN"
	CurrenciesPYG Currencies = "PYG"
	CurrenciesQAR Currencies = "QAR"
	CurrenciesRON Currencies = "RON"
	CurrenciesRSD Currencies = "RSD"
	CurrenciesRUB Currencies = "RUB"
	CurrenciesRWF Currencies = "RWF"
	CurrenciesSAR Currencies = "SAR"
	CurrenciesSBD Currencies = "SBD"
	CurrenciesSCR Currencies = "SCR"
	CurrenciesSDG Currencies = "SDG"
	CurrenciesSEK Currencies = "SEK"
	CurrenciesSGD Currencies = "SGD"
	CurrenciesSHP Currencies = "SHP"
	CurrenciesSLL Currencies = "SLL"
	CurrenciesSOS Currencies = "SOS"
	CurrenciesSRD Currencies = "SRD"
	CurrenciesSSP Currencies = "SSP"
	CurrenciesSTN Currencies = "STN"
	CurrenciesSVC Currencies = "SVC"
	CurrenciesSYP Currencies = "SYP"
	CurrenciesSZL Currencies = "SZL"
	CurrenciesTHB Currencies = "THB"
	CurrenciesTJS Currencies = "TJS"
	CurrenciesTMT Currencies = "TMT"
	CurrenciesTND Currencies = "TND"
	CurrenciesTOP Currencies = "TOP"
	CurrenciesTRY Currencies = "TRY"
	CurrenciesTTD Currencies = "TTD"
	CurrenciesTVD Currencies = "TVD"
	CurrenciesTWD Currencies = "TWD"
	CurrenciesTZS Currencies = "TZS"
	CurrenciesUAH Currencies = "UAH"
	CurrenciesUGX Currencies = "UGX"
	CurrenciesUSD Currencies = "USD"
	CurrenciesUSN Currencies = "USN"
	CurrenciesUYI Currencies = "UYI"
	CurrenciesUYU Currencies = "UYU"
	CurrenciesUZS Currencies = "UZS"
	CurrenciesVEF Currencies = "VEF"
	CurrenciesVND Currencies = "VND"
	CurrenciesVUV Currencies = "VUV"
	CurrenciesWST Currencies = "WST"
	CurrenciesXAF Currencies = "XAF"
	CurrenciesXAG Currencies = "XAG"
	CurrenciesXAU Currencies = "XAU"
	CurrenciesXBA Currencies = "XBA"
	CurrenciesXBB Currencies = "XBB"
	CurrenciesXBC Currencies = "XBC"
	CurrenciesXBD Currencies = "XBD"
	CurrenciesXCD Currencies = "XCD"
	CurrenciesXDR Currencies = "XDR"
	CurrenciesXOF Currencies = "XOF"
	CurrenciesXPD Currencies = "XPD"
	CurrenciesXPF Currencies = "XPF"
	CurrenciesXPT Currencies = "XPT"
	CurrenciesXSU Currencies = "XSU"
	CurrenciesXTS Currencies = "XTS"
	CurrenciesXUA Currencies = "XUA"
	CurrenciesXXX Currencies = "XXX"
	CurrenciesYER Currencies = "YER"
	CurrenciesZAR Currencies = "ZAR"
	CurrenciesZMW Currencies = "ZMW"
	CurrenciesZWL Currencies = "ZWL"
)

func AllCurrencies() []Currencies {
	return []Currencies{
		CurrenciesAED,
		CurrenciesAFN,
		CurrenciesALL,
		CurrenciesAMD,
		CurrenciesANG,
		CurrenciesAOA,
		CurrenciesARS,
		CurrenciesAUD,
		CurrenciesAWG,
		CurrenciesAZN,
		CurrenciesBAM,
		CurrenciesBBD,
		CurrenciesBDT,
		CurrenciesBGN,
		CurrenciesBHD,
		CurrenciesBIF,
		CurrenciesBMD,
		CurrenciesBND,
		CurrenciesBOB,
		CurrenciesBOV,
		CurrenciesBRL,
		CurrenciesBSD,
		CurrenciesBTN,
		CurrenciesBWP,
		CurrenciesBYN,
		CurrenciesBZD,
		CurrenciesCAD,
		CurrenciesCDF,
		CurrenciesCHE,
		CurrenciesCHF,
		CurrenciesCHW,
		CurrenciesCLF,
		CurrenciesCLP,
		CurrenciesCNY,
		CurrenciesCOP,
		CurrenciesCOU,
		CurrenciesCRC,
		CurrenciesCUC,
		CurrenciesCUP,
		CurrenciesCVE,
		CurrenciesCZK,
		CurrenciesDJF,
		CurrenciesDKK,
		CurrenciesDOP,
		CurrenciesDZD,
		CurrenciesEGP,
		CurrenciesERN,
		CurrenciesETB,
		CurrenciesEUR,
		CurrenciesFJD,
		CurrenciesFKP,
		CurrenciesGBP,
		CurrenciesGEL,
		CurrenciesGGP,
		CurrenciesGHS,
		CurrenciesGIP,
		CurrenciesGMD,
		CurrenciesGNF,
		CurrenciesGTQ,
		CurrenciesGYD,
		CurrenciesHKD,
		CurrenciesHNL,
		CurrenciesHRK,
		CurrenciesHTG,
		CurrenciesHUF,
		CurrenciesIDR,
		CurrenciesILS,
		CurrenciesIMP,
		CurrenciesINR,
		CurrenciesIQD,
		CurrenciesIRR,
		CurrenciesISK,
		CurrenciesJEP,
		CurrenciesJMD,
		CurrenciesJOD,
		CurrenciesJPY,
		CurrenciesKES,
		CurrenciesKGS,
		CurrenciesKHR,
		CurrenciesKMF,
		CurrenciesKPW,
		CurrenciesKRW,
		CurrenciesKWD,
		CurrenciesKYD,
		CurrenciesKZT,
		CurrenciesLAK,
		CurrenciesLBP,
		CurrenciesLKR,
		CurrenciesLRD,
		CurrenciesLSL,
		CurrenciesLYD,
		CurrenciesMAD,
		CurrenciesMDL,
		CurrenciesMGA,
		CurrenciesMKD,
		CurrenciesMMK,
		CurrenciesMNT,
		CurrenciesMOP,
		CurrenciesMRU,
		CurrenciesMUR,
		CurrenciesMVR,
		CurrenciesMWK,
		CurrenciesMXN,
		CurrenciesMXV,
		CurrenciesMYR,
		CurrenciesMZN,
		CurrenciesNAD,
		CurrenciesNGN,
		CurrenciesNIO,
		CurrenciesNOK,
		CurrenciesNPR,
		CurrenciesNZD,
		CurrenciesOMR,
		CurrenciesPAB,
		CurrenciesPEN,
		CurrenciesPGK,
		CurrenciesPHP,
		CurrenciesPKR,
		CurrenciesPLN,
		CurrenciesPYG,
		CurrenciesQAR,
		CurrenciesRON,
		CurrenciesRSD,
		CurrenciesRUB,
		CurrenciesRWF,
		CurrenciesSAR,
		CurrenciesSBD,
		CurrenciesSCR,
		CurrenciesSDG,
		CurrenciesSEK,
		CurrenciesSGD,
		CurrenciesSHP,
		CurrenciesSLL,
		CurrenciesSOS,
		CurrenciesSRD,
		CurrenciesSSP,
		CurrenciesSTN,
		CurrenciesSVC,
		CurrenciesSYP,
		CurrenciesSZL,
		CurrenciesTHB,
		CurrenciesTJS,
		CurrenciesTMT,
		CurrenciesTND,
		CurrenciesTOP,
		CurrenciesTRY,
		CurrenciesTTD,
		CurrenciesTVD,
		CurrenciesTWD,
		CurrenciesTZS,
		CurrenciesUAH,
		CurrenciesUGX,
		CurrenciesUSD,
		CurrenciesUSN,
		CurrenciesUYI,
		CurrenciesUYU,
		CurrenciesUZS,
		CurrenciesVEF,
		CurrenciesVND,
		CurrenciesVUV,
		CurrenciesWST,
		CurrenciesXAF,
		CurrenciesXAG,
		CurrenciesXAU,
		CurrenciesXBA,
		CurrenciesXBB,
		CurrenciesXBC,
		CurrenciesXBD,
		CurrenciesXCD,
		CurrenciesXDR,
		CurrenciesXOF,
		CurrenciesXPD,
		CurrenciesXPF,
		CurrenciesXPT,
		CurrenciesXSU,
		CurrenciesXTS,
		CurrenciesXUA,
		CurrenciesXXX,
		CurrenciesYER,
		CurrenciesZAR,
		CurrenciesZMW,
		CurrenciesZWL,
	}
}

func FindCurrencies(filter string) []Currencies {
	ret := make([]Currencies, 0)
	for _, at := range AllCurrencies() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au Currencies) ToString() {
	fmt.Println(au.String())
}
func (au Currencies) String() string {
	switch au {
	case CurrenciesAED:
		return "United Arab Emirates dirham"
	case CurrenciesAFN:
		return "Afghan afghani"
	case CurrenciesALL:
		return "Albanian lek"
	case CurrenciesAMD:
		return "Armenian dram"
	case CurrenciesANG:
		return "Netherlands Antillean guilder"
	case CurrenciesAOA:
		return "Angolan kwanza"
	case CurrenciesARS:
		return "Argentine peso"
	case CurrenciesAUD:
		return "Australian dollar"
	case CurrenciesAWG:
		return "Aruban florin"
	case CurrenciesAZN:
		return "Azerbaijani manat"
	case CurrenciesBAM:
		return "Bosnia and Herzegovina convertible mark"
	case CurrenciesBBD:
		return "Barbados dollar"
	case CurrenciesBDT:
		return "Bangladeshi taka"
	case CurrenciesBGN:
		return "Bulgarian lev"
	case CurrenciesBHD:
		return "Bahraini dinar"
	case CurrenciesBIF:
		return "Burundian franc"
	case CurrenciesBMD:
		return "Bermudian dollar"
	case CurrenciesBND:
		return "Brunei dollar"
	case CurrenciesBOB:
		return "Boliviano"
	case CurrenciesBOV:
		return "Bolivian Mvdol (funds code)"
	case CurrenciesBRL:
		return "Brazilian real"
	case CurrenciesBSD:
		return "Bahamian dollar"
	case CurrenciesBTN:
		return "Bhutanese ngultrum"
	case CurrenciesBWP:
		return "Botswana pula"
	case CurrenciesBYN:
		return "Belarusian ruble"
	case CurrenciesBZD:
		return "Belize dollar"
	case CurrenciesCAD:
		return "Canadian dollar"
	case CurrenciesCDF:
		return "Congolese franc"
	case CurrenciesCHE:
		return "WIR Euro (complementary currency)"
	case CurrenciesCHF:
		return "Swiss franc"
	case CurrenciesCHW:
		return "WIR Franc (complementary currency)"
	case CurrenciesCLF:
		return "Unidad de Fomento (funds code)"
	case CurrenciesCLP:
		return "Chilean peso"
	case CurrenciesCNY:
		return "Renminbi (Chinese) yuan[8]"
	case CurrenciesCOP:
		return "Colombian peso"
	case CurrenciesCOU:
		return "Unidad de Valor Real (UVR) (funds code)[9]"
	case CurrenciesCRC:
		return "Costa Rican colon"
	case CurrenciesCUC:
		return "Cuban convertible peso"
	case CurrenciesCUP:
		return "Cuban peso"
	case CurrenciesCVE:
		return "Cape Verde escudo"
	case CurrenciesCZK:
		return "Czech koruna"
	case CurrenciesDJF:
		return "Djiboutian franc"
	case CurrenciesDKK:
		return "Danish krone"
	case CurrenciesDOP:
		return "Dominican peso"
	case CurrenciesDZD:
		return "Algerian dinar"
	case CurrenciesEGP:
		return "Egyptian pound"
	case CurrenciesERN:
		return "Eritrean nakfa"
	case CurrenciesETB:
		return "Ethiopian birr"
	case CurrenciesEUR:
		return "Euro"
	case CurrenciesFJD:
		return "Fiji dollar"
	case CurrenciesFKP:
		return "Falkland Islands pound"
	case CurrenciesGBP:
		return "Pound sterling"
	case CurrenciesGEL:
		return "Georgian lari"
	case CurrenciesGGP:
		return "Guernsey Pound"
	case CurrenciesGHS:
		return "Ghanaian cedi"
	case CurrenciesGIP:
		return "Gibraltar pound"
	case CurrenciesGMD:
		return "Gambian dalasi"
	case CurrenciesGNF:
		return "Guinean franc"
	case CurrenciesGTQ:
		return "Guatemalan quetzal"
	case CurrenciesGYD:
		return "Guyanese dollar"
	case CurrenciesHKD:
		return "Hong Kong dollar"
	case CurrenciesHNL:
		return "Honduran lempira"
	case CurrenciesHRK:
		return "Croatian kuna"
	case CurrenciesHTG:
		return "Haitian gourde"
	case CurrenciesHUF:
		return "Hungarian forint"
	case CurrenciesIDR:
		return "Indonesian rupiah"
	case CurrenciesILS:
		return "Israeli new shekel"
	case CurrenciesIMP:
		return "Isle of Man Pound"
	case CurrenciesINR:
		return "Indian rupee"
	case CurrenciesIQD:
		return "Iraqi dinar"
	case CurrenciesIRR:
		return "Iranian rial"
	case CurrenciesISK:
		return "Icelandic króna"
	case CurrenciesJEP:
		return "Jersey Pound"
	case CurrenciesJMD:
		return "Jamaican dollar"
	case CurrenciesJOD:
		return "Jordanian dinar"
	case CurrenciesJPY:
		return "Japanese yen"
	case CurrenciesKES:
		return "Kenyan shilling"
	case CurrenciesKGS:
		return "Kyrgyzstani som"
	case CurrenciesKHR:
		return "Cambodian riel"
	case CurrenciesKMF:
		return "Comoro franc"
	case CurrenciesKPW:
		return "North Korean won"
	case CurrenciesKRW:
		return "South Korean won"
	case CurrenciesKWD:
		return "Kuwaiti dinar"
	case CurrenciesKYD:
		return "Cayman Islands dollar"
	case CurrenciesKZT:
		return "Kazakhstani tenge"
	case CurrenciesLAK:
		return "Lao kip"
	case CurrenciesLBP:
		return "Lebanese pound"
	case CurrenciesLKR:
		return "Sri Lankan rupee"
	case CurrenciesLRD:
		return "Liberian dollar"
	case CurrenciesLSL:
		return "Lesotho loti"
	case CurrenciesLYD:
		return "Libyan dinar"
	case CurrenciesMAD:
		return "Moroccan dirham"
	case CurrenciesMDL:
		return "Moldovan leu"
	case CurrenciesMGA:
		return "Malagasy ariary"
	case CurrenciesMKD:
		return "Macedonian denar"
	case CurrenciesMMK:
		return "Myanmar kyat"
	case CurrenciesMNT:
		return "Mongolian tögrög"
	case CurrenciesMOP:
		return "Macanese pataca"
	case CurrenciesMRU:
		return "Mauritanian ouguiya"
	case CurrenciesMUR:
		return "Mauritian rupee"
	case CurrenciesMVR:
		return "Maldivian rufiyaa"
	case CurrenciesMWK:
		return "Malawian kwacha"
	case CurrenciesMXN:
		return "Mexican peso"
	case CurrenciesMXV:
		return "Mexican Unidad de Inversion (UDI) (funds code)"
	case CurrenciesMYR:
		return "Malaysian ringgit"
	case CurrenciesMZN:
		return "Mozambican metical"
	case CurrenciesNAD:
		return "Namibian dollar"
	case CurrenciesNGN:
		return "Nigerian naira"
	case CurrenciesNIO:
		return "Nicaraguan córdoba"
	case CurrenciesNOK:
		return "Norwegian krone"
	case CurrenciesNPR:
		return "Nepalese rupee"
	case CurrenciesNZD:
		return "New Zealand dollar"
	case CurrenciesOMR:
		return "Omani rial"
	case CurrenciesPAB:
		return "Panamanian balboa"
	case CurrenciesPEN:
		return "Peruvian Sol"
	case CurrenciesPGK:
		return "Papua New Guinean kina"
	case CurrenciesPHP:
		return "Philippine piso[13]"
	case CurrenciesPKR:
		return "Pakistani rupee"
	case CurrenciesPLN:
		return "Polish złoty"
	case CurrenciesPYG:
		return "Paraguayan guaraní"
	case CurrenciesQAR:
		return "Qatari riyal"
	case CurrenciesRON:
		return "Romanian leu"
	case CurrenciesRSD:
		return "Serbian dinar"
	case CurrenciesRUB:
		return "Russian ruble"
	case CurrenciesRWF:
		return "Rwandan franc"
	case CurrenciesSAR:
		return "Saudi riyal"
	case CurrenciesSBD:
		return "Solomon Islands dollar"
	case CurrenciesSCR:
		return "Seychelles rupee"
	case CurrenciesSDG:
		return "Sudanese pound"
	case CurrenciesSEK:
		return "Swedish krona/kronor"
	case CurrenciesSGD:
		return "Singapore dollar"
	case CurrenciesSHP:
		return "Saint Helena pound"
	case CurrenciesSLL:
		return "Sierra Leonean leone"
	case CurrenciesSOS:
		return "Somali shilling"
	case CurrenciesSRD:
		return "Surinamese dollar"
	case CurrenciesSSP:
		return "South Sudanese pound"
	case CurrenciesSTN:
		return "São Tomé and Príncipe dobra"
	case CurrenciesSVC:
		return "Salvadoran colón"
	case CurrenciesSYP:
		return "Syrian pound"
	case CurrenciesSZL:
		return "Swazi lilangeni"
	case CurrenciesTHB:
		return "Thai baht"
	case CurrenciesTJS:
		return "Tajikistani somoni"
	case CurrenciesTMT:
		return "Turkmenistan manat"
	case CurrenciesTND:
		return "Tunisian dinar"
	case CurrenciesTOP:
		return "Tongan paʻanga"
	case CurrenciesTRY:
		return "Turkish lira"
	case CurrenciesTTD:
		return "Trinidad and Tobago dollar"
	case CurrenciesTVD:
		return "Tuvalu Dollar"
	case CurrenciesTWD:
		return "New Taiwan dollar"
	case CurrenciesTZS:
		return "Tanzanian shilling"
	case CurrenciesUAH:
		return "Ukrainian hryvnia"
	case CurrenciesUGX:
		return "Ugandan shilling"
	case CurrenciesUSD:
		return "United States dollar"
	case CurrenciesUSN:
		return "United States dollar (next day) (funds code)"
	case CurrenciesUYI:
		return "Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)"
	case CurrenciesUYU:
		return "Uruguayan peso"
	case CurrenciesUZS:
		return "Uzbekistan som"
	case CurrenciesVEF:
		return "Venezuelan bolívar"
	case CurrenciesVND:
		return "Vietnamese đồng"
	case CurrenciesVUV:
		return "Vanuatu vatu"
	case CurrenciesWST:
		return "Samoan tala"
	case CurrenciesXAF:
		return "CFA franc BEAC"
	case CurrenciesXAG:
		return "Silver (one troy ounce)"
	case CurrenciesXAU:
		return "Gold (one troy ounce)"
	case CurrenciesXBA:
		return "European Composite Unit (EURCO) (bond market unit)"
	case CurrenciesXBB:
		return "European Monetary Unit (E.M.U.-6) (bond market unit)"
	case CurrenciesXBC:
		return "European Unit of Account 9 (E.U.A.-9) (bond market unit)"
	case CurrenciesXBD:
		return "European Unit of Account 17 (E.U.A.-17) (bond market unit)"
	case CurrenciesXCD:
		return "East Caribbean dollar"
	case CurrenciesXDR:
		return "Special drawing rights"
	case CurrenciesXOF:
		return "CFA franc BCEAO"
	case CurrenciesXPD:
		return "Palladium (one troy ounce)"
	case CurrenciesXPF:
		return "CFP franc (franc Pacifique)"
	case CurrenciesXPT:
		return "Platinum (one troy ounce)"
	case CurrenciesXSU:
		return "SUCRE"
	case CurrenciesXTS:
		return "Code reserved for testing purposes"
	case CurrenciesXUA:
		return "ADB Unit of Account"
	case CurrenciesXXX:
		return "No currency"
	case CurrenciesYER:
		return "Yemeni rial"
	case CurrenciesZAR:
		return "South African rand"
	case CurrenciesZMW:
		return "Zambian kwacha"
	case CurrenciesZWL:
		return "Zimbabwean dollar A/10"

	default:
		return "Unknown Constraint Severity"
	}
}
