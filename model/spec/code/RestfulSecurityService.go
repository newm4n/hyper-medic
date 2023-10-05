package code

import (
	"fmt"
	"strings"
)

type RestfulSecurityService string

const (
	RestfulSecurityServiceOAuth        RestfulSecurityService = "OAuth"
	RestfulSecurityServiceSMARTOnFHIR  RestfulSecurityService = "SMART-on-FHIR"
	RestfulSecurityServiceNTLM         RestfulSecurityService = "NTLM"
	RestfulSecurityServiceBasic        RestfulSecurityService = "Basic"
	RestfulSecurityServiceKerberos     RestfulSecurityService = "Kerberos"
	RestfulSecurityServiceCertificates RestfulSecurityService = "Certificates"
)

func AllRestfulSecurityService() []RestfulSecurityService {
	return []RestfulSecurityService{
		RestfulSecurityServiceOAuth,
		RestfulSecurityServiceSMARTOnFHIR,
		RestfulSecurityServiceNTLM,
		RestfulSecurityServiceBasic,
		RestfulSecurityServiceKerberos,
		RestfulSecurityServiceCertificates,
	}
}

func FindRestfulSecurityService(filter string) []RestfulSecurityService {
	ret := make([]RestfulSecurityService, 0)
	for _, at := range AllRestfulSecurityService() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au RestfulSecurityService) ToString() {
	fmt.Println(au.String())
}
func (au RestfulSecurityService) String() string {
	switch au {
	case RestfulSecurityServiceOAuth:
		return "OAuth"
	case RestfulSecurityServiceSMARTOnFHIR:
		return "SMART-on-FHIR"
	case RestfulSecurityServiceNTLM:
		return "NTLM"
	case RestfulSecurityServiceBasic:
		return "Basic"
	case RestfulSecurityServiceKerberos:
		return "Kerberos"
	case RestfulSecurityServiceCertificates:
		return "Certificates"
	default:
		return "Unknown Restful Security Service"
	}
}
