package code

import (
	"fmt"
	"strings"
)

type SecurityLabels string

const (
	SecurityLabelsL                                           SecurityLabels = "L"
	SecurityLabelsM                                           SecurityLabels = "M"
	SecurityLabelsN                                           SecurityLabels = "N"
	SecurityLabelsR                                           SecurityLabels = "R"
	SecurityLabelsU                                           SecurityLabels = "U"
	SecurityLabelsV                                           SecurityLabels = "V"
	SecurityLabels_InformationSensitivityPolicy               SecurityLabels = "_InformationSensitivityPolicy"
	SecurityLabels_ActInformationSensitivityPolicy            SecurityLabels = "_ActInformationSensitivityPolicy"
	SecurityLabelsETH                                         SecurityLabels = "ETH"
	SecurityLabelsGDIS                                        SecurityLabels = "GDIS"
	SecurityLabelsHIV                                         SecurityLabels = "HIV"
	SecurityLabelsMST                                         SecurityLabels = "MST"
	SecurityLabelsPREGNANT                                    SecurityLabels = "PREGNANT"
	SecurityLabelsSCA                                         SecurityLabels = "SCA"
	SecurityLabelsSDV                                         SecurityLabels = "SDV"
	SecurityLabelsSEX                                         SecurityLabels = "SEX"
	SecurityLabelsSPI                                         SecurityLabels = "SPI"
	SecurityLabelsBH                                          SecurityLabels = "BH"
	SecurityLabelsCOGN                                        SecurityLabels = "COGN"
	SecurityLabelsDVD                                         SecurityLabels = "DVD"
	SecurityLabelsEMOTDIS                                     SecurityLabels = "EMOTDIS"
	SecurityLabelsMH                                          SecurityLabels = "MH"
	SecurityLabelsPSY                                         SecurityLabels = "PSY"
	SecurityLabelsPSYTHPN                                     SecurityLabels = "PSYTHPN"
	SecurityLabelsSUD                                         SecurityLabels = "SUD"
	SecurityLabelsETHUD                                       SecurityLabels = "ETHUD"
	SecurityLabelsOPIOIDUD                                    SecurityLabels = "OPIOIDUD"
	SecurityLabelsSTD                                         SecurityLabels = "STD"
	SecurityLabelsTBOO                                        SecurityLabels = "TBOO"
	SecurityLabelsVIO                                         SecurityLabels = "VIO"
	SecurityLabelsIDS                                         SecurityLabels = "IDS"
	SecurityLabelsSICKLE                                      SecurityLabels = "SICKLE"
	SecurityLabels_EntitySensitivityPolicyType                SecurityLabels = "_EntitySensitivityPolicyType"
	SecurityLabelsDEMO                                        SecurityLabels = "DEMO"
	SecurityLabelsDOB                                         SecurityLabels = "DOB"
	SecurityLabelsGENDER                                      SecurityLabels = "GENDER"
	SecurityLabelsLIVARG                                      SecurityLabels = "LIVARG"
	SecurityLabelsMARST                                       SecurityLabels = "MARST"
	SecurityLabelsPATLOC                                      SecurityLabels = "PATLOC"
	SecurityLabelsRACE                                        SecurityLabels = "RACE"
	SecurityLabelsREL                                         SecurityLabels = "REL"
	SecurityLabels_RoleInformationSensitivityPolicy           SecurityLabels = "_RoleInformationSensitivityPolicy"
	SecurityLabelsB                                           SecurityLabels = "B"
	SecurityLabelsEMPL                                        SecurityLabels = "EMPL"
	SecurityLabelsLOCIS                                       SecurityLabels = "LOCIS"
	SecurityLabelsSSP                                         SecurityLabels = "SSP"
	SecurityLabelsADOL                                        SecurityLabels = "ADOL"
	SecurityLabelsCEL                                         SecurityLabels = "CEL"
	SecurityLabelsVIP                                         SecurityLabels = "VIP"
	SecurityLabelsDIA                                         SecurityLabels = "DIA"
	SecurityLabelsDRGIS                                       SecurityLabels = "DRGIS"
	SecurityLabelsEMP                                         SecurityLabels = "EMP"
	SecurityLabelsPDS                                         SecurityLabels = "PDS"
	SecurityLabelsPHY                                         SecurityLabels = "PHY"
	SecurityLabelsPRS                                         SecurityLabels = "PRS"
	SecurityLabelsCOMPT                                       SecurityLabels = "COMPT"
	SecurityLabelsACOCOMPT                                    SecurityLabels = "ACOCOMPT"
	SecurityLabelsCDSSCOMPT                                   SecurityLabels = "CDSSCOMPT"
	SecurityLabelsCTCOMPT                                     SecurityLabels = "CTCOMPT"
	SecurityLabelsFMCOMPT                                     SecurityLabels = "FMCOMPT"
	SecurityLabelsHRCOMPT                                     SecurityLabels = "HRCOMPT"
	SecurityLabelsLRCOMPT                                     SecurityLabels = "LRCOMPT"
	SecurityLabelsPACOMPT                                     SecurityLabels = "PACOMPT"
	SecurityLabelsRESCOMPT                                    SecurityLabels = "RESCOMPT"
	SecurityLabelsRMGTCOMPT                                   SecurityLabels = "RMGTCOMPT"
	SecurityLabels_SECALTINTOBV                               SecurityLabels = "_SECALTINTOBV"
	SecurityLabelsABSTRED                                     SecurityLabels = "ABSTRED"
	SecurityLabelsAGGRED                                      SecurityLabels = "AGGRED"
	SecurityLabelsANONYED                                     SecurityLabels = "ANONYED"
	SecurityLabelsMAPPED                                      SecurityLabels = "MAPPED"
	SecurityLabelsMASKED                                      SecurityLabels = "MASKED"
	SecurityLabelsPSEUDED                                     SecurityLabels = "PSEUDED"
	SecurityLabelsREDACTED                                    SecurityLabels = "REDACTED"
	SecurityLabelsSUBSETTED                                   SecurityLabels = "SUBSETTED"
	SecurityLabelsSYNTAC                                      SecurityLabels = "SYNTAC"
	SecurityLabelsTRSLT                                       SecurityLabels = "TRSLT"
	SecurityLabelsVERSIONED                                   SecurityLabels = "VERSIONED"
	SecurityLabels_SECDATINTOBV                               SecurityLabels = "_SECDATINTOBV"
	SecurityLabelsCRYTOHASH                                   SecurityLabels = "CRYTOHASH"
	SecurityLabelsDIGSIG                                      SecurityLabels = "DIGSIG"
	SecurityLabels_SECINTCONOBV                               SecurityLabels = "_SECINTCONOBV"
	SecurityLabelsHRELIABLE                                   SecurityLabels = "HRELIABLE"
	SecurityLabelsRELIABLE                                    SecurityLabels = "RELIABLE"
	SecurityLabelsUNCERTREL                                   SecurityLabels = "UNCERTREL"
	SecurityLabelsUNRELIABLE                                  SecurityLabels = "UNRELIABLE"
	SecurityLabels_SECINTPRVOBV                               SecurityLabels = "_SECINTPRVOBV"
	SecurityLabels_SECINTPRVABOBV                             SecurityLabels = "_SECINTPRVABOBV"
	SecurityLabelsCLINAST                                     SecurityLabels = "CLINAST"
	SecurityLabelsDEVAST                                      SecurityLabels = "DEVAST"
	SecurityLabelsHCPAST                                      SecurityLabels = "HCPAST"
	SecurityLabelsPACQAST                                     SecurityLabels = "PACQAST"
	SecurityLabelsPATAST                                      SecurityLabels = "PATAST"
	SecurityLabelsPAYAST                                      SecurityLabels = "PAYAST"
	SecurityLabelsPROAST                                      SecurityLabels = "PROAST"
	SecurityLabelsSDMAST                                      SecurityLabels = "SDMAST"
	SecurityLabels_SECINTPRVRBOBV                             SecurityLabels = "_SECINTPRVRBOBV"
	SecurityLabelsCLINRPT                                     SecurityLabels = "CLINRPT"
	SecurityLabelsDEVRPT                                      SecurityLabels = "DEVRPT"
	SecurityLabelsHCPRPT                                      SecurityLabels = "HCPRPT"
	SecurityLabelsPACQRPT                                     SecurityLabels = "PACQRPT"
	SecurityLabelsPATRPT                                      SecurityLabels = "PATRPT"
	SecurityLabelsPAYRPT                                      SecurityLabels = "PAYRPT"
	SecurityLabelsPRORPT                                      SecurityLabels = "PRORPT"
	SecurityLabelsSDMRPT                                      SecurityLabels = "SDMRPT"
	SecurityLabels_SECINTSTOBV                                SecurityLabels = "_SECINTSTOBV"
	SecurityLabelsSecurityPolicy                              SecurityLabels = "SecurityPolicy"
	SecurityLabelsAUTHPOL                                     SecurityLabels = "AUTHPOL"
	SecurityLabelsACCESSCONSCHEME                             SecurityLabels = "ACCESSCONSCHEME"
	SecurityLabelsDELEPOL                                     SecurityLabels = "DELEPOL"
	SecurityLabelsObligationPolicy                            SecurityLabels = "ObligationPolicy"
	SecurityLabelsANONY                                       SecurityLabels = "ANONY"
	SecurityLabelsAOD                                         SecurityLabels = "AOD"
	SecurityLabelsAUDIT                                       SecurityLabels = "AUDIT"
	SecurityLabelsAUDTR                                       SecurityLabels = "AUDTR"
	SecurityLabelsCPLYPOL                                     SecurityLabels = "CPLYPOL"
	SecurityLabelsCPLYCC                                      SecurityLabels = "CPLYCC"
	SecurityLabelsCPLYCD                                      SecurityLabels = "CPLYCD"
	SecurityLabelsCPLYCUI                                     SecurityLabels = "CPLYCUI"
	SecurityLabelsCPLYJPP                                     SecurityLabels = "CPLYJPP"
	SecurityLabelsCPLYJSP                                     SecurityLabels = "CPLYJSP"
	SecurityLabelsCPLYOPP                                     SecurityLabels = "CPLYOPP"
	SecurityLabelsCPLYOSP                                     SecurityLabels = "CPLYOSP"
	SecurityLabelsDECLASSIFYLABEL                             SecurityLabels = "DECLASSIFYLABEL"
	SecurityLabelsDEID                                        SecurityLabels = "DEID"
	SecurityLabelsDELAU                                       SecurityLabels = "DELAU"
	SecurityLabelsDOWNGRDLABEL                                SecurityLabels = "DOWNGRDLABEL"
	SecurityLabelsDRIVLABEL                                   SecurityLabels = "DRIVLABEL"
	SecurityLabelsENCRYPT                                     SecurityLabels = "ENCRYPT"
	SecurityLabelsENCRYPTR                                    SecurityLabels = "ENCRYPTR"
	SecurityLabelsENCRYPTT                                    SecurityLabels = "ENCRYPTT"
	SecurityLabelsENCRYPTU                                    SecurityLabels = "ENCRYPTU"
	SecurityLabelsHUAPRV                                      SecurityLabels = "HUAPRV"
	SecurityLabelsLABEL                                       SecurityLabels = "LABEL"
	SecurityLabelsMASK                                        SecurityLabels = "MASK"
	SecurityLabelsMINEC                                       SecurityLabels = "MINEC"
	SecurityLabelsPERSISTLABEL                                SecurityLabels = "PERSISTLABEL"
	SecurityLabelsPRIVMARK                                    SecurityLabels = "PRIVMARK"
	SecurityLabelsCUIMark                                     SecurityLabels = "CUIMark"
	SecurityLabelsPSEUD                                       SecurityLabels = "PSEUD"
	SecurityLabelsREDACT                                      SecurityLabels = "REDACT"
	SecurityLabelsUPGRDLABEL                                  SecurityLabels = "UPGRDLABEL"
	SecurityLabelsPROCESSINLINELABEL                          SecurityLabels = "PROCESSINLINELABEL"
	SecurityLabelsPrivacyMark                                 SecurityLabels = "PrivacyMark"
	SecurityLabelsControlledUnclassifiedInformation           SecurityLabels = "ControlledUnclassifiedInformation"
	SecurityLabelsCONTROLLED                                  SecurityLabels = "CONTROLLED"
	SecurityLabelsCUI                                         SecurityLabels = "CUI"
	SecurityLabelsCUIHLTH                                     SecurityLabels = "CUIHLTH"
	SecurityLabelsCUIHLTHP                                    SecurityLabels = "CUIHLTHP"
	SecurityLabelsCUIP                                        SecurityLabels = "CUIP"
	SecurityLabelsCUIPRVCY                                    SecurityLabels = "CUIPRVCY"
	SecurityLabelsCUIPRVCYP                                   SecurityLabels = "CUIPRVCYP"
	SecurityLabelsCUISPHLTH                                   SecurityLabels = "CUISP-HLTH"
	SecurityLabelsCUISPHLTHP                                  SecurityLabels = "CUISP-HLTHP"
	SecurityLabelsCUISPPRVCY                                  SecurityLabels = "CUISP-PRVCY"
	SecurityLabelsCUISPPRVCYP                                 SecurityLabels = "CUISP-PRVCYP"
	SecurityLabelsUUI                                         SecurityLabels = "UUI"
	SecurityLabelsSecurityLabelMark                           SecurityLabels = "SecurityLabelMark"
	SecurityLabelsConfidentialMark                            SecurityLabels = "ConfidentialMark"
	SecurityLabelsCOPYMark                                    SecurityLabels = "COPYMark"
	SecurityLabelsDeliverToAddresseeOnlyMark                  SecurityLabels = "DeliverToAddresseeOnlyMark"
	SecurityLabelsRedisclosureProhibitionMark                 SecurityLabels = "RedisclosureProhibitionMark"
	SecurityLabelsRestrictedConfidentialityMark               SecurityLabels = "RestrictedConfidentialityMark"
	SecurityLabelsDRAFTMark                                   SecurityLabels = "DRAFTMark"
	SecurityLabelsRefrainPolicy                               SecurityLabels = "RefrainPolicy"
	SecurityLabelsNOAUTH                                      SecurityLabels = "NOAUTH"
	SecurityLabelsNOCOLLECT                                   SecurityLabels = "NOCOLLECT"
	SecurityLabelsNODSCLCD                                    SecurityLabels = "NODSCLCD"
	SecurityLabelsNODSCLCDS                                   SecurityLabels = "NODSCLCDS"
	SecurityLabelsNOINTEGRATE                                 SecurityLabels = "NOINTEGRATE"
	SecurityLabelsNOLIST                                      SecurityLabels = "NOLIST"
	SecurityLabelsNOMOU                                       SecurityLabels = "NOMOU"
	SecurityLabelsNOORGPOL                                    SecurityLabels = "NOORGPOL"
	SecurityLabelsNOPAT                                       SecurityLabels = "NOPAT"
	SecurityLabelsNOPERSISTP                                  SecurityLabels = "NOPERSISTP"
	SecurityLabelsNORDSCLCD                                   SecurityLabels = "NORDSCLCD"
	SecurityLabelsNORDSLCD                                    SecurityLabels = "NORDSLCD"
	SecurityLabelsNORDSCLCDS                                  SecurityLabels = "NORDSCLCDS"
	SecurityLabelsNORDSCLW                                    SecurityLabels = "NORDSCLW"
	SecurityLabelsNORELINK                                    SecurityLabels = "NORELINK"
	SecurityLabelsNOREUSE                                     SecurityLabels = "NOREUSE"
	SecurityLabelsNOVIP                                       SecurityLabels = "NOVIP"
	SecurityLabelsORCON                                       SecurityLabels = "ORCON"
	SecurityLabelsHMARKT                                      SecurityLabels = "HMARKT"
	SecurityLabelsHOPERAT                                     SecurityLabels = "HOPERAT"
	SecurityLabelsCAREMGT                                     SecurityLabels = "CAREMGT"
	SecurityLabelsDONAT                                       SecurityLabels = "DONAT"
	SecurityLabelsFRAUD                                       SecurityLabels = "FRAUD"
	SecurityLabelsGOV                                         SecurityLabels = "GOV"
	SecurityLabelsHACCRED                                     SecurityLabels = "HACCRED"
	SecurityLabelsHCOMPL                                      SecurityLabels = "HCOMPL"
	SecurityLabelsHDECD                                       SecurityLabels = "HDECD"
	SecurityLabelsHDIRECT                                     SecurityLabels = "HDIRECT"
	SecurityLabelsHDM                                         SecurityLabels = "HDM"
	SecurityLabelsHLEGAL                                      SecurityLabels = "HLEGAL"
	SecurityLabelsHOUTCOMS                                    SecurityLabels = "HOUTCOMS"
	SecurityLabelsHPRGRP                                      SecurityLabels = "HPRGRP"
	SecurityLabelsHQUALIMP                                    SecurityLabels = "HQUALIMP"
	SecurityLabelsHSYSADMIN                                   SecurityLabels = "HSYSADMIN"
	SecurityLabelsLABELING                                    SecurityLabels = "LABELING"
	SecurityLabelsMETAMGT                                     SecurityLabels = "METAMGT"
	SecurityLabelsMEMADMIN                                    SecurityLabels = "MEMADMIN"
	SecurityLabelsMILCDM                                      SecurityLabels = "MILCDM"
	SecurityLabelsPATADMIN                                    SecurityLabels = "PATADMIN"
	SecurityLabelsPATSFTY                                     SecurityLabels = "PATSFTY"
	SecurityLabelsPERFMSR                                     SecurityLabels = "PERFMSR"
	SecurityLabelsRECORDMGT                                   SecurityLabels = "RECORDMGT"
	SecurityLabelsSYSDEV                                      SecurityLabels = "SYSDEV"
	SecurityLabelsHTEST                                       SecurityLabels = "HTEST"
	SecurityLabelsTRAIN                                       SecurityLabels = "TRAIN"
	SecurityLabelsHPAYMT                                      SecurityLabels = "HPAYMT"
	SecurityLabelsCLMATTCH                                    SecurityLabels = "CLMATTCH"
	SecurityLabelsCOVAUTH                                     SecurityLabels = "COVAUTH"
	SecurityLabelsCOVERAGE                                    SecurityLabels = "COVERAGE"
	SecurityLabelsELIGDTRM                                    SecurityLabels = "ELIGDTRM"
	SecurityLabelsELIGVER                                     SecurityLabels = "ELIGVER"
	SecurityLabelsENROLLM                                     SecurityLabels = "ENROLLM"
	SecurityLabelsMILDCRG                                     SecurityLabels = "MILDCRG"
	SecurityLabelsREMITADV                                    SecurityLabels = "REMITADV"
	SecurityLabelsHRESCH                                      SecurityLabels = "HRESCH"
	SecurityLabelsBIORCH                                      SecurityLabels = "BIORCH"
	SecurityLabelsCLINTRCH                                    SecurityLabels = "CLINTRCH"
	SecurityLabelsCLINTRCHNPC                                 SecurityLabels = "CLINTRCHNPC"
	SecurityLabelsCLINTRCHPC                                  SecurityLabels = "CLINTRCHPC"
	SecurityLabelsPRECLINTRCH                                 SecurityLabels = "PRECLINTRCH"
	SecurityLabelsDSRCH                                       SecurityLabels = "DSRCH"
	SecurityLabelsPOARCH                                      SecurityLabels = "POARCH"
	SecurityLabelsTRANSRCH                                    SecurityLabels = "TRANSRCH"
	SecurityLabelsPATRQT                                      SecurityLabels = "PATRQT"
	SecurityLabelsFAMRQT                                      SecurityLabels = "FAMRQT"
	SecurityLabelsPWATRNY                                     SecurityLabels = "PWATRNY"
	SecurityLabelsSUPNWK                                      SecurityLabels = "SUPNWK"
	SecurityLabelsPUBHLTH                                     SecurityLabels = "PUBHLTH"
	SecurityLabelsDISASTER                                    SecurityLabels = "DISASTER"
	SecurityLabelsTHREAT                                      SecurityLabels = "THREAT"
	SecurityLabelsTREAT                                       SecurityLabels = "TREAT"
	SecurityLabelsCLINTRL                                     SecurityLabels = "CLINTRL"
	SecurityLabelsCOC                                         SecurityLabels = "COC"
	SecurityLabelsETREAT                                      SecurityLabels = "ETREAT"
	SecurityLabelsBTG                                         SecurityLabels = "BTG"
	SecurityLabelsERTREAT                                     SecurityLabels = "ERTREAT"
	SecurityLabelsPOPHLTH                                     SecurityLabels = "POPHLTH"
	SecurityLabels_ActCoverageAssessmentObservationValue      SecurityLabels = "_ActCoverageAssessmentObservationValue"
	SecurityLabels_ActFinancialStatusObservationValue         SecurityLabels = "_ActFinancialStatusObservationValue"
	SecurityLabelsASSET                                       SecurityLabels = "ASSET"
	SecurityLabelsANNUITY                                     SecurityLabels = "ANNUITY"
	SecurityLabelsPROP                                        SecurityLabels = "PROP"
	SecurityLabelsRETACCT                                     SecurityLabels = "RETACCT"
	SecurityLabelsTRUST                                       SecurityLabels = "TRUST"
	SecurityLabelsINCOME                                      SecurityLabels = "INCOME"
	SecurityLabelsCHILD                                       SecurityLabels = "CHILD"
	SecurityLabelsDISABL                                      SecurityLabels = "DISABL"
	SecurityLabelsINVEST                                      SecurityLabels = "INVEST"
	SecurityLabelsPAY                                         SecurityLabels = "PAY"
	SecurityLabelsRETIRE                                      SecurityLabels = "RETIRE"
	SecurityLabelsSPOUSAL                                     SecurityLabels = "SPOUSAL"
	SecurityLabelsSUPPLE                                      SecurityLabels = "SUPPLE"
	SecurityLabelsTAX                                         SecurityLabels = "TAX"
	SecurityLabelsLIVEXP                                      SecurityLabels = "LIVEXP"
	SecurityLabelsCLOTH                                       SecurityLabels = "CLOTH"
	SecurityLabelsFOOD                                        SecurityLabels = "FOOD"
	SecurityLabelsHEALTH                                      SecurityLabels = "HEALTH"
	SecurityLabelsHOUSE                                       SecurityLabels = "HOUSE"
	SecurityLabelsLEGAL                                       SecurityLabels = "LEGAL"
	SecurityLabelsMORTG                                       SecurityLabels = "MORTG"
	SecurityLabelsRENT                                        SecurityLabels = "RENT"
	SecurityLabelsSUNDRY                                      SecurityLabels = "SUNDRY"
	SecurityLabelsTRANS                                       SecurityLabels = "TRANS"
	SecurityLabelsUTIL                                        SecurityLabels = "UTIL"
	SecurityLabelsELSTAT                                      SecurityLabels = "ELSTAT"
	SecurityLabelsADOPT                                       SecurityLabels = "ADOPT"
	SecurityLabelsBTHCERT                                     SecurityLabels = "BTHCERT"
	SecurityLabelsCCOC                                        SecurityLabels = "CCOC"
	SecurityLabelsDRLIC                                       SecurityLabels = "DRLIC"
	SecurityLabelsFOSTER                                      SecurityLabels = "FOSTER"
	SecurityLabelsMEMBER                                      SecurityLabels = "MEMBER"
	SecurityLabelsMIL                                         SecurityLabels = "MIL"
	SecurityLabelsMRGCERT                                     SecurityLabels = "MRGCERT"
	SecurityLabelsPASSPORT                                    SecurityLabels = "PASSPORT"
	SecurityLabelsSTUDENRL                                    SecurityLabels = "STUDENRL"
	SecurityLabelsHLSTAT                                      SecurityLabels = "HLSTAT"
	SecurityLabelsDISABLE                                     SecurityLabels = "DISABLE"
	SecurityLabelsDRUG                                        SecurityLabels = "DRUG"
	SecurityLabelsIVDRG                                       SecurityLabels = "IVDRG"
	SecurityLabelsPGNT                                        SecurityLabels = "PGNT"
	SecurityLabelsLIVDEP                                      SecurityLabels = "LIVDEP"
	SecurityLabelsRELDEP                                      SecurityLabels = "RELDEP"
	SecurityLabelsSPSDEP                                      SecurityLabels = "SPSDEP"
	SecurityLabelsURELDEP                                     SecurityLabels = "URELDEP"
	SecurityLabelsLIVSIT                                      SecurityLabels = "LIVSIT"
	SecurityLabelsALONE                                       SecurityLabels = "ALONE"
	SecurityLabelsDEPCHD                                      SecurityLabels = "DEPCHD"
	SecurityLabelsDEPSPS                                      SecurityLabels = "DEPSPS"
	SecurityLabelsDEPYGCHD                                    SecurityLabels = "DEPYGCHD"
	SecurityLabelsFAM                                         SecurityLabels = "FAM"
	SecurityLabelsRELAT                                       SecurityLabels = "RELAT"
	SecurityLabelsSPS                                         SecurityLabels = "SPS"
	SecurityLabelsUNREL                                       SecurityLabels = "UNREL"
	SecurityLabelsSOECSTAT                                    SecurityLabels = "SOECSTAT"
	SecurityLabelsABUSE                                       SecurityLabels = "ABUSE"
	SecurityLabelsHMLESS                                      SecurityLabels = "HMLESS"
	SecurityLabelsILGIM                                       SecurityLabels = "ILGIM"
	SecurityLabelsINCAR                                       SecurityLabels = "INCAR"
	SecurityLabelsPROB                                        SecurityLabels = "PROB"
	SecurityLabelsREFUG                                       SecurityLabels = "REFUG"
	SecurityLabelsUNEMPL                                      SecurityLabels = "UNEMPL"
	SecurityLabels_AllergyTestValue                           SecurityLabels = "_AllergyTestValue"
	SecurityLabelsA0                                          SecurityLabels = "A0"
	SecurityLabelsA1                                          SecurityLabels = "A1"
	SecurityLabelsA2                                          SecurityLabels = "A2"
	SecurityLabelsA3                                          SecurityLabels = "A3"
	SecurityLabelsA4                                          SecurityLabels = "A4"
	SecurityLabels_CompositeMeasureScoring                    SecurityLabels = "_CompositeMeasureScoring"
	SecurityLabelsALLORNONESCR                                SecurityLabels = "ALLORNONESCR"
	SecurityLabelsLINEARSCR                                   SecurityLabels = "LINEARSCR"
	SecurityLabelsOPPORSCR                                    SecurityLabels = "OPPORSCR"
	SecurityLabelsWEIGHTSCR                                   SecurityLabels = "WEIGHTSCR"
	SecurityLabels_CoverageLimitObservationValue              SecurityLabels = "_CoverageLimitObservationValue"
	SecurityLabels_CoverageLevelObservationValue              SecurityLabels = "_CoverageLevelObservationValue"
	SecurityLabelsADC                                         SecurityLabels = "ADC"
	SecurityLabelsCHD                                         SecurityLabels = "CHD"
	SecurityLabelsDEP                                         SecurityLabels = "DEP"
	SecurityLabelsDP                                          SecurityLabels = "DP"
	SecurityLabelsECH                                         SecurityLabels = "ECH"
	SecurityLabelsFLY                                         SecurityLabels = "FLY"
	SecurityLabelsIND                                         SecurityLabels = "IND"
	SecurityLabels_CoverageItemLimitObservationValue          SecurityLabels = "_CoverageItemLimitObservationValue"
	SecurityLabels_CoverageLocationLimitObservationValue      SecurityLabels = "_CoverageLocationLimitObservationValue"
	SecurityLabels_CriticalityObservationValue                SecurityLabels = "_CriticalityObservationValue"
	SecurityLabelsCRITH                                       SecurityLabels = "CRITH"
	SecurityLabelsCRITL                                       SecurityLabels = "CRITL"
	SecurityLabelsCRITU                                       SecurityLabels = "CRITU"
	SecurityLabels_EmploymentStatus                           SecurityLabels = "_EmploymentStatus"
	SecurityLabelsEmployed                                    SecurityLabels = "Employed"
	SecurityLabelsNotInLaborForce                             SecurityLabels = "NotInLaborForce"
	SecurityLabelsUnemployed                                  SecurityLabels = "Unemployed"
	SecurityLabels_GeneticObservationValue                    SecurityLabels = "_GeneticObservationValue"
	SecurityLabelsHomozygote                                  SecurityLabels = "Homozygote"
	SecurityLabels_MeasurementImprovementNotation             SecurityLabels = "_MeasurementImprovementNotation"
	SecurityLabelsDecrIsImp                                   SecurityLabels = "DecrIsImp"
	SecurityLabelsIncrIsImp                                   SecurityLabels = "IncrIsImp"
	SecurityLabels_ObservationMeasureScoring                  SecurityLabels = "_ObservationMeasureScoring"
	SecurityLabelsCOHORT                                      SecurityLabels = "COHORT"
	SecurityLabelsCONTVAR                                     SecurityLabels = "CONTVAR"
	SecurityLabelsPROPOR                                      SecurityLabels = "PROPOR"
	SecurityLabelsRATIO                                       SecurityLabels = "RATIO"
	SecurityLabels_ObservationMeasureType                     SecurityLabels = "_ObservationMeasureType"
	SecurityLabelsCOMPOSITE                                   SecurityLabels = "COMPOSITE"
	SecurityLabelsEFFICIENCY                                  SecurityLabels = "EFFICIENCY"
	SecurityLabelsEXPERIENCE                                  SecurityLabels = "EXPERIENCE"
	SecurityLabelsOUTCOME                                     SecurityLabels = "OUTCOME"
	SecurityLabelsINTERMOM                                    SecurityLabels = "INTERM-OM"
	SecurityLabelsPROPM                                       SecurityLabels = "PRO-PM"
	SecurityLabelsPROCESS                                     SecurityLabels = "PROCESS"
	SecurityLabelsAPPROPRIATE                                 SecurityLabels = "APPROPRIATE"
	SecurityLabelsRESOURCE                                    SecurityLabels = "RESOURCE"
	SecurityLabelsSTRUCTURE                                   SecurityLabels = "STRUCTURE"
	SecurityLabels_ObservationPopulationInclusion             SecurityLabels = "_ObservationPopulationInclusion"
	SecurityLabelsDENEX                                       SecurityLabels = "DENEX"
	SecurityLabelsDENEXCEP                                    SecurityLabels = "DENEXCEP"
	SecurityLabelsDENOM                                       SecurityLabels = "DENOM"
	SecurityLabelsIP                                          SecurityLabels = "IP"
	SecurityLabelsIPP                                         SecurityLabels = "IPP"
	SecurityLabelsMSRPOPL                                     SecurityLabels = "MSRPOPL"
	SecurityLabelsNUMER                                       SecurityLabels = "NUMER"
	SecurityLabelsNUMEX                                       SecurityLabels = "NUMEX"
	SecurityLabels_PartialCompletionScale                     SecurityLabels = "_PartialCompletionScale"
	SecurityLabelsG                                           SecurityLabels = "G"
	SecurityLabelsLE                                          SecurityLabels = "LE"
	SecurityLabelsME                                          SecurityLabels = "ME"
	SecurityLabelsMI                                          SecurityLabels = "MI"
	SecurityLabelsS                                           SecurityLabels = "S"
	SecurityLabels_SecurityObservationValue                   SecurityLabels = "_SecurityObservationValue"
	SecurityLabels_SECCATOBV                                  SecurityLabels = "_SECCATOBV"
	SecurityLabels_SECCLASSOBV                                SecurityLabels = "_SECCLASSOBV"
	SecurityLabels_SECCONOBV                                  SecurityLabels = "_SECCONOBV"
	SecurityLabels_SECINTOBV                                  SecurityLabels = "_SECINTOBV"
	SecurityLabelsSECTRSTOBV                                  SecurityLabels = "SECTRSTOBV"
	SecurityLabelsTRSTACCRDOBV                                SecurityLabels = "TRSTACCRDOBV"
	SecurityLabelsTRSTAGREOBV                                 SecurityLabels = "TRSTAGREOBV"
	SecurityLabelsTRSTCERTOBV                                 SecurityLabels = "TRSTCERTOBV"
	SecurityLabelsTRSTFWKOBV                                  SecurityLabels = "TRSTFWKOBV"
	SecurityLabelsTRSTLOAOBV                                  SecurityLabels = "TRSTLOAOBV"
	SecurityLabelsLOAAN                                       SecurityLabels = "LOAAN"
	SecurityLabelsLOAAN1                                      SecurityLabels = "LOAAN1"
	SecurityLabelsLOAAN2                                      SecurityLabels = "LOAAN2"
	SecurityLabelsLOAAN3                                      SecurityLabels = "LOAAN3"
	SecurityLabelsLOAAN4                                      SecurityLabels = "LOAAN4"
	SecurityLabelsLOAAP                                       SecurityLabels = "LOAAP"
	SecurityLabelsLOAAP1                                      SecurityLabels = "LOAAP1"
	SecurityLabelsLOAAP2                                      SecurityLabels = "LOAAP2"
	SecurityLabelsLOAAP3                                      SecurityLabels = "LOAAP3"
	SecurityLabelsLOAAP4                                      SecurityLabels = "LOAAP4"
	SecurityLabelsLOAAS                                       SecurityLabels = "LOAAS"
	SecurityLabelsLOAAS1                                      SecurityLabels = "LOAAS1"
	SecurityLabelsLOAAS2                                      SecurityLabels = "LOAAS2"
	SecurityLabelsLOAAS3                                      SecurityLabels = "LOAAS3"
	SecurityLabelsLOAAS4                                      SecurityLabels = "LOAAS4"
	SecurityLabelsLOACM                                       SecurityLabels = "LOACM"
	SecurityLabelsLOACM1                                      SecurityLabels = "LOACM1"
	SecurityLabelsLOACM2                                      SecurityLabels = "LOACM2"
	SecurityLabelsLOACM3                                      SecurityLabels = "LOACM3"
	SecurityLabelsLOACM4                                      SecurityLabels = "LOACM4"
	SecurityLabelsLOAID                                       SecurityLabels = "LOAID"
	SecurityLabelsLOAID1                                      SecurityLabels = "LOAID1"
	SecurityLabelsLOAID2                                      SecurityLabels = "LOAID2"
	SecurityLabelsLOAID3                                      SecurityLabels = "LOAID3"
	SecurityLabelsLOAID4                                      SecurityLabels = "LOAID4"
	SecurityLabelsLOANR                                       SecurityLabels = "LOANR"
	SecurityLabelsLOANR1                                      SecurityLabels = "LOANR1"
	SecurityLabelsLOANR2                                      SecurityLabels = "LOANR2"
	SecurityLabelsLOANR3                                      SecurityLabels = "LOANR3"
	SecurityLabelsLOANR4                                      SecurityLabels = "LOANR4"
	SecurityLabelsLOARA                                       SecurityLabels = "LOARA"
	SecurityLabelsLOARA1                                      SecurityLabels = "LOARA1"
	SecurityLabelsLOARA2                                      SecurityLabels = "LOARA2"
	SecurityLabelsLOARA3                                      SecurityLabels = "LOARA3"
	SecurityLabelsLOARA4                                      SecurityLabels = "LOARA4"
	SecurityLabelsLOATK                                       SecurityLabels = "LOATK"
	SecurityLabelsLOATK1                                      SecurityLabels = "LOATK1"
	SecurityLabelsLOATK2                                      SecurityLabels = "LOATK2"
	SecurityLabelsLOATK3                                      SecurityLabels = "LOATK3"
	SecurityLabelsLOATK4                                      SecurityLabels = "LOATK4"
	SecurityLabelsTRSTMECOBV                                  SecurityLabels = "TRSTMECOBV"
	SecurityLabels_SeverityObservation                        SecurityLabels = "_SeverityObservation"
	SecurityLabelsH                                           SecurityLabels = "H"
	SecurityLabels_SubjectBodyPosition                        SecurityLabels = "_SubjectBodyPosition"
	SecurityLabelsLLD                                         SecurityLabels = "LLD"
	SecurityLabelsPRN                                         SecurityLabels = "PRN"
	SecurityLabelsRLD                                         SecurityLabels = "RLD"
	SecurityLabelsSFWL                                        SecurityLabels = "SFWL"
	SecurityLabelsSIT                                         SecurityLabels = "SIT"
	SecurityLabelsSTN                                         SecurityLabels = "STN"
	SecurityLabelsSUP                                         SecurityLabels = "SUP"
	SecurityLabelsRTRD                                        SecurityLabels = "RTRD"
	SecurityLabelsTRD                                         SecurityLabels = "TRD"
	SecurityLabels_VerificationOutcomeValue                   SecurityLabels = "_VerificationOutcomeValue"
	SecurityLabelsACT                                         SecurityLabels = "ACT"
	SecurityLabelsACTPEND                                     SecurityLabels = "ACTPEND"
	SecurityLabelsELG                                         SecurityLabels = "ELG"
	SecurityLabelsINACT                                       SecurityLabels = "INACT"
	SecurityLabelsINPNDINV                                    SecurityLabels = "INPNDINV"
	SecurityLabelsINPNDUPD                                    SecurityLabels = "INPNDUPD"
	SecurityLabelsNELG                                        SecurityLabels = "NELG"
	SecurityLabels_WorkSchedule                               SecurityLabels = "_WorkSchedule"
	SecurityLabelsDS                                          SecurityLabels = "DS"
	SecurityLabelsEMS                                         SecurityLabels = "EMS"
	SecurityLabelsES                                          SecurityLabels = "ES"
	SecurityLabelsNS                                          SecurityLabels = "NS"
	SecurityLabelsRSWN                                        SecurityLabels = "RSWN"
	SecurityLabelsRSWON                                       SecurityLabels = "RSWON"
	SecurityLabelsSS                                          SecurityLabels = "SS"
	SecurityLabelsVLS                                         SecurityLabels = "VLS"
	SecurityLabelsVS                                          SecurityLabels = "VS"
	SecurityLabels_AnnotationValue                            SecurityLabels = "_AnnotationValue"
	SecurityLabels_ECGAnnotationValue                         SecurityLabels = "_ECGAnnotationValue"
	SecurityLabels_CommonClinicalObservationValue             SecurityLabels = "_CommonClinicalObservationValue"
	SecurityLabels_CommonClinicalObservationAssertionValue    SecurityLabels = "_CommonClinicalObservationAssertionValue"
	SecurityLabels_CommonClinicalObservationResultValue       SecurityLabels = "_CommonClinicalObservationResultValue"
	SecurityLabels_CoverageChemicalDependencyValue            SecurityLabels = "_CoverageChemicalDependencyValue"
	SecurityLabels_IndividualCaseSafetyReportValueDomains     SecurityLabels = "_IndividualCaseSafetyReportValueDomains"
	SecurityLabels_CaseSeriousnessCriteria                    SecurityLabels = "_CaseSeriousnessCriteria"
	SecurityLabels_DeviceManufacturerEvaluationInterpretation SecurityLabels = "_DeviceManufacturerEvaluationInterpretation"
	SecurityLabels_DeviceManufacturerEvaluationMethod         SecurityLabels = "_DeviceManufacturerEvaluationMethod"
	SecurityLabels_DeviceManufacturerEvaluationResult         SecurityLabels = "_DeviceManufacturerEvaluationResult"
	SecurityLabels_PertinentReactionRelatedness               SecurityLabels = "_PertinentReactionRelatedness"
	SecurityLabels_ProductCharacterization                    SecurityLabels = "_ProductCharacterization"
	SecurityLabels_ReactionActionTaken                        SecurityLabels = "_ReactionActionTaken"
	SecurityLabels_SubjectReaction                            SecurityLabels = "_SubjectReaction"
	SecurityLabels_SubjectReactionEmphasis                    SecurityLabels = "_SubjectReactionEmphasis"
	SecurityLabels_SubjectReactionOutcome                     SecurityLabels = "_SubjectReactionOutcome"
	SecurityLabels_InjuryObservationValue                     SecurityLabels = "_InjuryObservationValue"
	SecurityLabels_IntoleranceValue                           SecurityLabels = "_IntoleranceValue"
	SecurityLabels_IssueTriggerObservationValue               SecurityLabels = "_IssueTriggerObservationValue"
	SecurityLabels_OtherIndicationValue                       SecurityLabels = "_OtherIndicationValue"
	SecurityLabels_IndicationValue                            SecurityLabels = "_IndicationValue"
	SecurityLabels_DiagnosisValue                             SecurityLabels = "_DiagnosisValue"
	SecurityLabels_SymptomValue                               SecurityLabels = "_SymptomValue"
	SecurityLabels_ActUSPrivacyLaw                            SecurityLabels = "_ActUSPrivacyLaw"
	SecurityLabels42CFRPart2                                  SecurityLabels = "42CFRPart2"
	SecurityLabelsCommonRule                                  SecurityLabels = "CommonRule"
	SecurityLabelsHIPAANOPP                                   SecurityLabels = "HIPAANOPP"
	SecurityLabelsHIPAAPsyNotes                               SecurityLabels = "HIPAAPsyNotes"
	SecurityLabelsHIPAASelfPay                                SecurityLabels = "HIPAASelfPay"
	SecurityLabelsTitle38Section7332                          SecurityLabels = "Title38Section7332"
	SecurityLabelsTitle38Part1                                SecurityLabels = "Title38Part1"
)

func AllSecurityLabels() []SecurityLabels {
	return []SecurityLabels{
		SecurityLabelsL,
		SecurityLabelsM,
		SecurityLabelsN,
		SecurityLabelsR,
		SecurityLabelsU,
		SecurityLabelsV,
		SecurityLabels_InformationSensitivityPolicy,
		SecurityLabels_ActInformationSensitivityPolicy,
		SecurityLabelsETH,
		SecurityLabelsGDIS,
		SecurityLabelsHIV,
		SecurityLabelsMST,
		SecurityLabelsPREGNANT,
		SecurityLabelsSCA,
		SecurityLabelsSDV,
		SecurityLabelsSEX,
		SecurityLabelsSPI,
		SecurityLabelsBH,
		SecurityLabelsCOGN,
		SecurityLabelsDVD,
		SecurityLabelsEMOTDIS,
		SecurityLabelsMH,
		SecurityLabelsPSY,
		SecurityLabelsPSYTHPN,
		SecurityLabelsSUD,
		SecurityLabelsETHUD,
		SecurityLabelsOPIOIDUD,
		SecurityLabelsSTD,
		SecurityLabelsTBOO,
		SecurityLabelsVIO,
		SecurityLabelsIDS,
		SecurityLabelsSICKLE,
		SecurityLabels_EntitySensitivityPolicyType,
		SecurityLabelsDEMO,
		SecurityLabelsDOB,
		SecurityLabelsGENDER,
		SecurityLabelsLIVARG,
		SecurityLabelsMARST,
		SecurityLabelsPATLOC,
		SecurityLabelsRACE,
		SecurityLabelsREL,
		SecurityLabels_RoleInformationSensitivityPolicy,
		SecurityLabelsB,
		SecurityLabelsEMPL,
		SecurityLabelsLOCIS,
		SecurityLabelsSSP,
		SecurityLabelsADOL,
		SecurityLabelsCEL,
		SecurityLabelsVIP,
		SecurityLabelsDIA,
		SecurityLabelsDRGIS,
		SecurityLabelsEMP,
		SecurityLabelsPDS,
		SecurityLabelsPHY,
		SecurityLabelsPRS,
		SecurityLabelsCOMPT,
		SecurityLabelsACOCOMPT,
		SecurityLabelsCDSSCOMPT,
		SecurityLabelsCTCOMPT,
		SecurityLabelsFMCOMPT,
		SecurityLabelsHRCOMPT,
		SecurityLabelsLRCOMPT,
		SecurityLabelsPACOMPT,
		SecurityLabelsRESCOMPT,
		SecurityLabelsRMGTCOMPT,
		SecurityLabels_SECALTINTOBV,
		SecurityLabelsABSTRED,
		SecurityLabelsAGGRED,
		SecurityLabelsANONYED,
		SecurityLabelsMAPPED,
		SecurityLabelsMASKED,
		SecurityLabelsPSEUDED,
		SecurityLabelsREDACTED,
		SecurityLabelsSUBSETTED,
		SecurityLabelsSYNTAC,
		SecurityLabelsTRSLT,
		SecurityLabelsVERSIONED,
		SecurityLabels_SECDATINTOBV,
		SecurityLabelsCRYTOHASH,
		SecurityLabelsDIGSIG,
		SecurityLabels_SECINTCONOBV,
		SecurityLabelsHRELIABLE,
		SecurityLabelsRELIABLE,
		SecurityLabelsUNCERTREL,
		SecurityLabelsUNRELIABLE,
		SecurityLabels_SECINTPRVOBV,
		SecurityLabels_SECINTPRVABOBV,
		SecurityLabelsCLINAST,
		SecurityLabelsDEVAST,
		SecurityLabelsHCPAST,
		SecurityLabelsPACQAST,
		SecurityLabelsPATAST,
		SecurityLabelsPAYAST,
		SecurityLabelsPROAST,
		SecurityLabelsSDMAST,
		SecurityLabels_SECINTPRVRBOBV,
		SecurityLabelsCLINRPT,
		SecurityLabelsDEVRPT,
		SecurityLabelsHCPRPT,
		SecurityLabelsPACQRPT,
		SecurityLabelsPATRPT,
		SecurityLabelsPAYRPT,
		SecurityLabelsPRORPT,
		SecurityLabelsSDMRPT,
		SecurityLabels_SECINTSTOBV,
		SecurityLabelsSecurityPolicy,
		SecurityLabelsAUTHPOL,
		SecurityLabelsACCESSCONSCHEME,
		SecurityLabelsDELEPOL,
		SecurityLabelsObligationPolicy,
		SecurityLabelsANONY,
		SecurityLabelsAOD,
		SecurityLabelsAUDIT,
		SecurityLabelsAUDTR,
		SecurityLabelsCPLYPOL,
		SecurityLabelsCPLYCC,
		SecurityLabelsCPLYCD,
		SecurityLabelsCPLYCUI,
		SecurityLabelsCPLYJPP,
		SecurityLabelsCPLYJSP,
		SecurityLabelsCPLYOPP,
		SecurityLabelsCPLYOSP,
		SecurityLabelsDECLASSIFYLABEL,
		SecurityLabelsDEID,
		SecurityLabelsDELAU,
		SecurityLabelsDOWNGRDLABEL,
		SecurityLabelsDRIVLABEL,
		SecurityLabelsENCRYPT,
		SecurityLabelsENCRYPTR,
		SecurityLabelsENCRYPTT,
		SecurityLabelsENCRYPTU,
		SecurityLabelsHUAPRV,
		SecurityLabelsLABEL,
		SecurityLabelsMASK,
		SecurityLabelsMINEC,
		SecurityLabelsPERSISTLABEL,
		SecurityLabelsPRIVMARK,
		SecurityLabelsCUIMark,
		SecurityLabelsPSEUD,
		SecurityLabelsREDACT,
		SecurityLabelsUPGRDLABEL,
		SecurityLabelsPROCESSINLINELABEL,
		SecurityLabelsPrivacyMark,
		SecurityLabelsControlledUnclassifiedInformation,
		SecurityLabelsCONTROLLED,
		SecurityLabelsCUI,
		SecurityLabelsCUIHLTH,
		SecurityLabelsCUIHLTHP,
		SecurityLabelsCUIP,
		SecurityLabelsCUIPRVCY,
		SecurityLabelsCUIPRVCYP,
		SecurityLabelsCUISPHLTH,
		SecurityLabelsCUISPHLTHP,
		SecurityLabelsCUISPPRVCY,
		SecurityLabelsCUISPPRVCYP,
		SecurityLabelsUUI,
		SecurityLabelsSecurityLabelMark,
		SecurityLabelsConfidentialMark,
		SecurityLabelsCOPYMark,
		SecurityLabelsDeliverToAddresseeOnlyMark,
		SecurityLabelsRedisclosureProhibitionMark,
		SecurityLabelsRestrictedConfidentialityMark,
		SecurityLabelsDRAFTMark,
		SecurityLabelsRefrainPolicy,
		SecurityLabelsNOAUTH,
		SecurityLabelsNOCOLLECT,
		SecurityLabelsNODSCLCD,
		SecurityLabelsNODSCLCDS,
		SecurityLabelsNOINTEGRATE,
		SecurityLabelsNOLIST,
		SecurityLabelsNOMOU,
		SecurityLabelsNOORGPOL,
		SecurityLabelsNOPAT,
		SecurityLabelsNOPERSISTP,
		SecurityLabelsNORDSCLCD,
		SecurityLabelsNORDSLCD,
		SecurityLabelsNORDSCLCDS,
		SecurityLabelsNORDSCLW,
		SecurityLabelsNORELINK,
		SecurityLabelsNOREUSE,
		SecurityLabelsNOVIP,
		SecurityLabelsORCON,
		SecurityLabelsHMARKT,
		SecurityLabelsHOPERAT,
		SecurityLabelsCAREMGT,
		SecurityLabelsDONAT,
		SecurityLabelsFRAUD,
		SecurityLabelsGOV,
		SecurityLabelsHACCRED,
		SecurityLabelsHCOMPL,
		SecurityLabelsHDECD,
		SecurityLabelsHDIRECT,
		SecurityLabelsHDM,
		SecurityLabelsHLEGAL,
		SecurityLabelsHOUTCOMS,
		SecurityLabelsHPRGRP,
		SecurityLabelsHQUALIMP,
		SecurityLabelsHSYSADMIN,
		SecurityLabelsLABELING,
		SecurityLabelsMETAMGT,
		SecurityLabelsMEMADMIN,
		SecurityLabelsMILCDM,
		SecurityLabelsPATADMIN,
		SecurityLabelsPATSFTY,
		SecurityLabelsPERFMSR,
		SecurityLabelsRECORDMGT,
		SecurityLabelsSYSDEV,
		SecurityLabelsHTEST,
		SecurityLabelsTRAIN,
		SecurityLabelsHPAYMT,
		SecurityLabelsCLMATTCH,
		SecurityLabelsCOVAUTH,
		SecurityLabelsCOVERAGE,
		SecurityLabelsELIGDTRM,
		SecurityLabelsELIGVER,
		SecurityLabelsENROLLM,
		SecurityLabelsMILDCRG,
		SecurityLabelsREMITADV,
		SecurityLabelsHRESCH,
		SecurityLabelsBIORCH,
		SecurityLabelsCLINTRCH,
		SecurityLabelsCLINTRCHNPC,
		SecurityLabelsCLINTRCHPC,
		SecurityLabelsPRECLINTRCH,
		SecurityLabelsDSRCH,
		SecurityLabelsPOARCH,
		SecurityLabelsTRANSRCH,
		SecurityLabelsPATRQT,
		SecurityLabelsFAMRQT,
		SecurityLabelsPWATRNY,
		SecurityLabelsSUPNWK,
		SecurityLabelsPUBHLTH,
		SecurityLabelsDISASTER,
		SecurityLabelsTHREAT,
		SecurityLabelsTREAT,
		SecurityLabelsCLINTRL,
		SecurityLabelsCOC,
		SecurityLabelsETREAT,
		SecurityLabelsBTG,
		SecurityLabelsERTREAT,
		SecurityLabelsPOPHLTH,
		SecurityLabels_ActCoverageAssessmentObservationValue,
		SecurityLabels_ActFinancialStatusObservationValue,
		SecurityLabelsASSET,
		SecurityLabelsANNUITY,
		SecurityLabelsPROP,
		SecurityLabelsRETACCT,
		SecurityLabelsTRUST,
		SecurityLabelsINCOME,
		SecurityLabelsCHILD,
		SecurityLabelsDISABL,
		SecurityLabelsINVEST,
		SecurityLabelsPAY,
		SecurityLabelsRETIRE,
		SecurityLabelsSPOUSAL,
		SecurityLabelsSUPPLE,
		SecurityLabelsTAX,
		SecurityLabelsLIVEXP,
		SecurityLabelsCLOTH,
		SecurityLabelsFOOD,
		SecurityLabelsHEALTH,
		SecurityLabelsHOUSE,
		SecurityLabelsLEGAL,
		SecurityLabelsMORTG,
		SecurityLabelsRENT,
		SecurityLabelsSUNDRY,
		SecurityLabelsTRANS,
		SecurityLabelsUTIL,
		SecurityLabelsELSTAT,
		SecurityLabelsADOPT,
		SecurityLabelsBTHCERT,
		SecurityLabelsCCOC,
		SecurityLabelsDRLIC,
		SecurityLabelsFOSTER,
		SecurityLabelsMEMBER,
		SecurityLabelsMIL,
		SecurityLabelsMRGCERT,
		SecurityLabelsPASSPORT,
		SecurityLabelsSTUDENRL,
		SecurityLabelsHLSTAT,
		SecurityLabelsDISABLE,
		SecurityLabelsDRUG,
		SecurityLabelsIVDRG,
		SecurityLabelsPGNT,
		SecurityLabelsLIVDEP,
		SecurityLabelsRELDEP,
		SecurityLabelsSPSDEP,
		SecurityLabelsURELDEP,
		SecurityLabelsLIVSIT,
		SecurityLabelsALONE,
		SecurityLabelsDEPCHD,
		SecurityLabelsDEPSPS,
		SecurityLabelsDEPYGCHD,
		SecurityLabelsFAM,
		SecurityLabelsRELAT,
		SecurityLabelsSPS,
		SecurityLabelsUNREL,
		SecurityLabelsSOECSTAT,
		SecurityLabelsABUSE,
		SecurityLabelsHMLESS,
		SecurityLabelsILGIM,
		SecurityLabelsINCAR,
		SecurityLabelsPROB,
		SecurityLabelsREFUG,
		SecurityLabelsUNEMPL,
		SecurityLabels_AllergyTestValue,
		SecurityLabelsA0,
		SecurityLabelsA1,
		SecurityLabelsA2,
		SecurityLabelsA3,
		SecurityLabelsA4,
		SecurityLabels_CompositeMeasureScoring,
		SecurityLabelsALLORNONESCR,
		SecurityLabelsLINEARSCR,
		SecurityLabelsOPPORSCR,
		SecurityLabelsWEIGHTSCR,
		SecurityLabels_CoverageLimitObservationValue,
		SecurityLabels_CoverageLevelObservationValue,
		SecurityLabelsADC,
		SecurityLabelsCHD,
		SecurityLabelsDEP,
		SecurityLabelsDP,
		SecurityLabelsECH,
		SecurityLabelsFLY,
		SecurityLabelsIND,
		SecurityLabels_CoverageItemLimitObservationValue,
		SecurityLabels_CoverageLocationLimitObservationValue,
		SecurityLabels_CriticalityObservationValue,
		SecurityLabelsCRITH,
		SecurityLabelsCRITL,
		SecurityLabelsCRITU,
		SecurityLabels_EmploymentStatus,
		SecurityLabelsEmployed,
		SecurityLabelsNotInLaborForce,
		SecurityLabelsUnemployed,
		SecurityLabels_GeneticObservationValue,
		SecurityLabelsHomozygote,
		SecurityLabels_MeasurementImprovementNotation,
		SecurityLabelsDecrIsImp,
		SecurityLabelsIncrIsImp,
		SecurityLabels_ObservationMeasureScoring,
		SecurityLabelsCOHORT,
		SecurityLabelsCONTVAR,
		SecurityLabelsPROPOR,
		SecurityLabelsRATIO,
		SecurityLabels_ObservationMeasureType,
		SecurityLabelsCOMPOSITE,
		SecurityLabelsEFFICIENCY,
		SecurityLabelsEXPERIENCE,
		SecurityLabelsOUTCOME,
		SecurityLabelsINTERMOM,
		SecurityLabelsPROPM,
		SecurityLabelsPROCESS,
		SecurityLabelsAPPROPRIATE,
		SecurityLabelsRESOURCE,
		SecurityLabelsSTRUCTURE,
		SecurityLabels_ObservationPopulationInclusion,
		SecurityLabelsDENEX,
		SecurityLabelsDENEXCEP,
		SecurityLabelsDENOM,
		SecurityLabelsIP,
		SecurityLabelsIPP,
		SecurityLabelsMSRPOPL,
		SecurityLabelsNUMER,
		SecurityLabelsNUMEX,
		SecurityLabels_PartialCompletionScale,
		SecurityLabelsG,
		SecurityLabelsLE,
		SecurityLabelsME,
		SecurityLabelsMI,
		SecurityLabelsS,
		SecurityLabels_SecurityObservationValue,
		SecurityLabels_SECCATOBV,
		SecurityLabels_SECCLASSOBV,
		SecurityLabels_SECCONOBV,
		SecurityLabels_SECINTOBV,
		SecurityLabelsSECTRSTOBV,
		SecurityLabelsTRSTACCRDOBV,
		SecurityLabelsTRSTAGREOBV,
		SecurityLabelsTRSTCERTOBV,
		SecurityLabelsTRSTFWKOBV,
		SecurityLabelsTRSTLOAOBV,
		SecurityLabelsLOAAN,
		SecurityLabelsLOAAN1,
		SecurityLabelsLOAAN2,
		SecurityLabelsLOAAN3,
		SecurityLabelsLOAAN4,
		SecurityLabelsLOAAP,
		SecurityLabelsLOAAP1,
		SecurityLabelsLOAAP2,
		SecurityLabelsLOAAP3,
		SecurityLabelsLOAAP4,
		SecurityLabelsLOAAS,
		SecurityLabelsLOAAS1,
		SecurityLabelsLOAAS2,
		SecurityLabelsLOAAS3,
		SecurityLabelsLOAAS4,
		SecurityLabelsLOACM,
		SecurityLabelsLOACM1,
		SecurityLabelsLOACM2,
		SecurityLabelsLOACM3,
		SecurityLabelsLOACM4,
		SecurityLabelsLOAID,
		SecurityLabelsLOAID1,
		SecurityLabelsLOAID2,
		SecurityLabelsLOAID3,
		SecurityLabelsLOAID4,
		SecurityLabelsLOANR,
		SecurityLabelsLOANR1,
		SecurityLabelsLOANR2,
		SecurityLabelsLOANR3,
		SecurityLabelsLOANR4,
		SecurityLabelsLOARA,
		SecurityLabelsLOARA1,
		SecurityLabelsLOARA2,
		SecurityLabelsLOARA3,
		SecurityLabelsLOARA4,
		SecurityLabelsLOATK,
		SecurityLabelsLOATK1,
		SecurityLabelsLOATK2,
		SecurityLabelsLOATK3,
		SecurityLabelsLOATK4,
		SecurityLabelsTRSTMECOBV,
		SecurityLabels_SeverityObservation,
		SecurityLabelsH,
		SecurityLabels_SubjectBodyPosition,
		SecurityLabelsLLD,
		SecurityLabelsPRN,
		SecurityLabelsRLD,
		SecurityLabelsSFWL,
		SecurityLabelsSIT,
		SecurityLabelsSTN,
		SecurityLabelsSUP,
		SecurityLabelsRTRD,
		SecurityLabelsTRD,
		SecurityLabels_VerificationOutcomeValue,
		SecurityLabelsACT,
		SecurityLabelsACTPEND,
		SecurityLabelsELG,
		SecurityLabelsINACT,
		SecurityLabelsINPNDINV,
		SecurityLabelsINPNDUPD,
		SecurityLabelsNELG,
		SecurityLabels_WorkSchedule,
		SecurityLabelsDS,
		SecurityLabelsEMS,
		SecurityLabelsES,
		SecurityLabelsNS,
		SecurityLabelsRSWN,
		SecurityLabelsRSWON,
		SecurityLabelsSS,
		SecurityLabelsVLS,
		SecurityLabelsVS,
		SecurityLabels_AnnotationValue,
		SecurityLabels_ECGAnnotationValue,
		SecurityLabels_CommonClinicalObservationValue,
		SecurityLabels_CommonClinicalObservationAssertionValue,
		SecurityLabels_CommonClinicalObservationResultValue,
		SecurityLabels_CoverageChemicalDependencyValue,
		SecurityLabels_IndividualCaseSafetyReportValueDomains,
		SecurityLabels_CaseSeriousnessCriteria,
		SecurityLabels_DeviceManufacturerEvaluationInterpretation,
		SecurityLabels_DeviceManufacturerEvaluationMethod,
		SecurityLabels_DeviceManufacturerEvaluationResult,
		SecurityLabels_PertinentReactionRelatedness,
		SecurityLabels_ProductCharacterization,
		SecurityLabels_ReactionActionTaken,
		SecurityLabels_SubjectReaction,
		SecurityLabels_SubjectReactionEmphasis,
		SecurityLabels_SubjectReactionOutcome,
		SecurityLabels_InjuryObservationValue,
		SecurityLabels_IntoleranceValue,
		SecurityLabels_IssueTriggerObservationValue,
		SecurityLabels_OtherIndicationValue,
		SecurityLabels_IndicationValue,
		SecurityLabels_DiagnosisValue,
		SecurityLabels_SymptomValue,
		SecurityLabels_ActUSPrivacyLaw,
		SecurityLabels42CFRPart2,
		SecurityLabelsCommonRule,
		SecurityLabelsHIPAANOPP,
		SecurityLabelsHIPAAPsyNotes,
		SecurityLabelsHIPAASelfPay,
		SecurityLabelsTitle38Section7332,
		SecurityLabelsTitle38Part1,
	}
}

func FindSecurityLabels(filter string) []SecurityLabels {
	ret := make([]SecurityLabels, 0)
	for _, at := range AllSecurityLabels() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au SecurityLabels) ToString() {
	fmt.Println(au.String())
}
func (au SecurityLabels) String() string {
	switch au {
	case SecurityLabelsL:
		return "low"
	case SecurityLabelsM:
		return "moderate"
	case SecurityLabelsN:
		return "normal"
	case SecurityLabelsR:
		return "restricted"
	case SecurityLabelsU:
		return "unrestricted"
	case SecurityLabelsV:
		return "very restricted"
	case SecurityLabels_InformationSensitivityPolicy:
		return "InformationSensitivityPolicy"
	case SecurityLabels_ActInformationSensitivityPolicy:
		return "ActInformationSensitivityPolicy"
	case SecurityLabelsETH:
		return "substance abuse information sensitivity"
	case SecurityLabelsGDIS:
		return "genetic disease information sensitivity"
	case SecurityLabelsHIV:
		return "HIV/AIDS information sensitivity"
	case SecurityLabelsMST:
		return "military sexual trauma information sensitivity"
	case SecurityLabelsPREGNANT:
		return "pregnancy information sensitivity"
	case SecurityLabelsSCA:
		return "sickle cell anemia information sensitivity"
	case SecurityLabelsSDV:
		return "sexual assault, abuse, or domestic violence information sensitivity"
	case SecurityLabelsSEX:
		return "sexuality and reproductive health information sensitivity"
	case SecurityLabelsSPI:
		return "specially protected information sensitivity"
	case SecurityLabelsBH:
		return "behavioral health information sensitivity"
	case SecurityLabelsCOGN:
		return "cognitive disability information sensitivity"
	case SecurityLabelsDVD:
		return "developmental disability information sensitivity"
	case SecurityLabelsEMOTDIS:
		return "emotional disturbance information sensitivity"
	case SecurityLabelsMH:
		return "mental health information sensitivity"
	case SecurityLabelsPSY:
		return "psychiatry disorder information sensitivity"
	case SecurityLabelsPSYTHPN:
		return "psychotherapy note information sensitivity"
	case SecurityLabelsSUD:
		return "substance use disorder information sensitivity"
	case SecurityLabelsETHUD:
		return "alcohol use disorder information sensitivity"
	case SecurityLabelsOPIOIDUD:
		return "opioid use disorder information sensitivity"
	case SecurityLabelsSTD:
		return "sexually transmitted disease information sensitivity"
	case SecurityLabelsTBOO:
		return "taboo"
	case SecurityLabelsVIO:
		return "violence information sensitivity"
	case SecurityLabelsIDS:
		return "Identifier Sensitivity"
	case SecurityLabelsSICKLE:
		return "sickle cell"
	case SecurityLabels_EntitySensitivityPolicyType:
		return "EntityInformationSensitivityPolicy"
	case SecurityLabelsDEMO:
		return "all demographic information sensitivity"
	case SecurityLabelsDOB:
		return "date of birth information sensitivity"
	case SecurityLabelsGENDER:
		return "gender and sexual orientation information sensitivity"
	case SecurityLabelsLIVARG:
		return "living arrangement information sensitivity"
	case SecurityLabelsMARST:
		return "marital status information sensitivity"
	case SecurityLabelsPATLOC:
		return "patient location"
	case SecurityLabelsRACE:
		return "race information sensitivity"
	case SecurityLabelsREL:
		return "religion information sensitivity"
	case SecurityLabels_RoleInformationSensitivityPolicy:
		return "RoleInformationSensitivityPolicy"
	case SecurityLabelsB:
		return "business information sensitivity"
	case SecurityLabelsEMPL:
		return "employer information sensitivity"
	case SecurityLabelsLOCIS:
		return "location information sensitivity"
	case SecurityLabelsSSP:
		return "sensitive service provider information sensitivity"
	case SecurityLabelsADOL:
		return "adolescent information sensitivity"
	case SecurityLabelsCEL:
		return "celebrity information sensitivity"
	case SecurityLabelsVIP:
		return "celebrity information sensitivity"
	case SecurityLabelsDIA:
		return "diagnosis information sensitivity"
	case SecurityLabelsDRGIS:
		return "drug information sensitivity"
	case SecurityLabelsEMP:
		return "employee information sensitivity"
	case SecurityLabelsPDS:
		return "patient default information sensitivity"
	case SecurityLabelsPHY:
		return "physician requested information sensitivity"
	case SecurityLabelsPRS:
		return "patient requested information sensitivity"
	case SecurityLabelsCOMPT:
		return "compartment"
	case SecurityLabelsACOCOMPT:
		return "accountable care organization compartment"
	case SecurityLabelsCDSSCOMPT:
		return "CDS system compartment"
	case SecurityLabelsCTCOMPT:
		return "care team compartment"
	case SecurityLabelsFMCOMPT:
		return "financial management compartment"
	case SecurityLabelsHRCOMPT:
		return "human resource compartment"
	case SecurityLabelsLRCOMPT:
		return "legitimate relationship compartment"
	case SecurityLabelsPACOMPT:
		return "patient administration compartment"
	case SecurityLabelsRESCOMPT:
		return "research project compartment"
	case SecurityLabelsRMGTCOMPT:
		return "records management compartment"
	case SecurityLabels_SECALTINTOBV:
		return "alteration integrity"
	case SecurityLabelsABSTRED:
		return "abstracted"
	case SecurityLabelsAGGRED:
		return "aggregated"
	case SecurityLabelsANONYED:
		return "anonymized"
	case SecurityLabelsMAPPED:
		return "mapped"
	case SecurityLabelsMASKED:
		return "masked"
	case SecurityLabelsPSEUDED:
		return "pseudonymized"
	case SecurityLabelsREDACTED:
		return "redacted"
	case SecurityLabelsSUBSETTED:
		return "subsetted"
	case SecurityLabelsSYNTAC:
		return "syntactic transform"
	case SecurityLabelsTRSLT:
		return "translated"
	case SecurityLabelsVERSIONED:
		return "versioned"
	case SecurityLabels_SECDATINTOBV:
		return "data integrity"
	case SecurityLabelsCRYTOHASH:
		return "cryptographic hash function"
	case SecurityLabelsDIGSIG:
		return "digital signature"
	case SecurityLabels_SECINTCONOBV:
		return "integrity confidence"
	case SecurityLabelsHRELIABLE:
		return "highly reliable"
	case SecurityLabelsRELIABLE:
		return "reliable"
	case SecurityLabelsUNCERTREL:
		return "uncertain reliability"
	case SecurityLabelsUNRELIABLE:
		return "unreliable"
	case SecurityLabels_SECINTPRVOBV:
		return "provenance"
	case SecurityLabels_SECINTPRVABOBV:
		return "provenance asserted by"
	case SecurityLabelsCLINAST:
		return "clinician asserted"
	case SecurityLabelsDEVAST:
		return "device asserted"
	case SecurityLabelsHCPAST:
		return "healthcare professional asserted"
	case SecurityLabelsPACQAST:
		return "patient acquaintance asserted"
	case SecurityLabelsPATAST:
		return "patient asserted"
	case SecurityLabelsPAYAST:
		return "payer asserted"
	case SecurityLabelsPROAST:
		return "professional asserted"
	case SecurityLabelsSDMAST:
		return "substitute decision maker asserted"
	case SecurityLabels_SECINTPRVRBOBV:
		return "provenance reported by"
	case SecurityLabelsCLINRPT:
		return "clinician reported"
	case SecurityLabelsDEVRPT:
		return "device reported"
	case SecurityLabelsHCPRPT:
		return "healthcare professional reported"
	case SecurityLabelsPACQRPT:
		return "patient acquaintance reported"
	case SecurityLabelsPATRPT:
		return "patient reported"
	case SecurityLabelsPAYRPT:
		return "payer reported"
	case SecurityLabelsPRORPT:
		return "professional reported"
	case SecurityLabelsSDMRPT:
		return "substitute decision maker reported"
	case SecurityLabels_SECINTSTOBV:
		return "integrity status"
	case SecurityLabelsSecurityPolicy:
		return "security policy"
	case SecurityLabelsAUTHPOL:
		return "authorization policy"
	case SecurityLabelsACCESSCONSCHEME:
		return "access control scheme"
	case SecurityLabelsDELEPOL:
		return "delegation policy"
	case SecurityLabelsObligationPolicy:
		return "obligation policy"
	case SecurityLabelsANONY:
		return "anonymize"
	case SecurityLabelsAOD:
		return "accounting of disclosure"
	case SecurityLabelsAUDIT:
		return "audit"
	case SecurityLabelsAUDTR:
		return "audit trail"
	case SecurityLabelsCPLYPOL:
		return "comply with policy"
	case SecurityLabelsCPLYCC:
		return "comply with confidentiality code"
	case SecurityLabelsCPLYCD:
		return "comply with consent directive"
	case SecurityLabelsCPLYCUI:
		return "comply with controlled unclassified information policy"
	case SecurityLabelsCPLYJPP:
		return "comply with jurisdictional privacy policy"
	case SecurityLabelsCPLYJSP:
		return "comply with jurisdictional security policy"
	case SecurityLabelsCPLYOPP:
		return "comply with organizational privacy policy"
	case SecurityLabelsCPLYOSP:
		return "comply with organizational security policy"
	case SecurityLabelsDECLASSIFYLABEL:
		return "declassify security label"
	case SecurityLabelsDEID:
		return "deidentify"
	case SecurityLabelsDELAU:
		return "delete after use"
	case SecurityLabelsDOWNGRDLABEL:
		return "downgrade security label"
	case SecurityLabelsDRIVLABEL:
		return "derive security label"
	case SecurityLabelsENCRYPT:
		return "encrypt"
	case SecurityLabelsENCRYPTR:
		return "encrypt at rest"
	case SecurityLabelsENCRYPTT:
		return "encrypt in transit"
	case SecurityLabelsENCRYPTU:
		return "encrypt in use"
	case SecurityLabelsHUAPRV:
		return "human approval"
	case SecurityLabelsLABEL:
		return "assign security label"
	case SecurityLabelsMASK:
		return "mask"
	case SecurityLabelsMINEC:
		return "minimum necessary"
	case SecurityLabelsPERSISTLABEL:
		return "persist security label"
	case SecurityLabelsPRIVMARK:
		return "privacy mark"
	case SecurityLabelsCUIMark:
		return "CUI Mark"
	case SecurityLabelsPSEUD:
		return "pseudonymize"
	case SecurityLabelsREDACT:
		return "redact"
	case SecurityLabelsUPGRDLABEL:
		return "upgrade security label"
	case SecurityLabelsPROCESSINLINELABEL:
		return "process inline security label"
	case SecurityLabelsPrivacyMark:
		return "privacy mark"
	case SecurityLabelsControlledUnclassifiedInformation:
		return "ControlledUnclassifiedInformation"
	case SecurityLabelsCONTROLLED:
		return "CONTROLLED"
	case SecurityLabelsCUI:
		return "CUI"
	case SecurityLabelsCUIHLTH:
		return "CUI//HLTH"
	case SecurityLabelsCUIHLTHP:
		return "(CUI//HLTH)"
	case SecurityLabelsCUIP:
		return "(CUI)"
	case SecurityLabelsCUIPRVCY:
		return "CUI//PRVCY"
	case SecurityLabelsCUIPRVCYP:
		return "(CUI//PRVCY)"
	case SecurityLabelsCUISPHLTH:
		return "CUI//SP-HLTH"
	case SecurityLabelsCUISPHLTHP:
		return "(CUI//SP-HLTH)"
	case SecurityLabelsCUISPPRVCY:
		return "CUI//SP-PRVCY"
	case SecurityLabelsCUISPPRVCYP:
		return "(CUI//SP-PRVCY)"
	case SecurityLabelsUUI:
		return "(U)"
	case SecurityLabelsSecurityLabelMark:
		return "Security Label Mark"
	case SecurityLabelsConfidentialMark:
		return "confidential mark"
	case SecurityLabelsCOPYMark:
		return "copy of original mark"
	case SecurityLabelsDeliverToAddresseeOnlyMark:
		return "deliver only to addressee mark"
	case SecurityLabelsRedisclosureProhibitionMark:
		return "prohibition against redisclosure mark"
	case SecurityLabelsRestrictedConfidentialityMark:
		return "restricted confidentiality mark"
	case SecurityLabelsDRAFTMark:
		return "Draft Mark"
	case SecurityLabelsRefrainPolicy:
		return "refrain policy"
	case SecurityLabelsNOAUTH:
		return "no disclosure without subject authorization"
	case SecurityLabelsNOCOLLECT:
		return "no collection"
	case SecurityLabelsNODSCLCD:
		return "no disclosure without consent directive"
	case SecurityLabelsNODSCLCDS:
		return "no disclosure without information subject's consent directive"
	case SecurityLabelsNOINTEGRATE:
		return "no integration"
	case SecurityLabelsNOLIST:
		return "no unlisted entity disclosure"
	case SecurityLabelsNOMOU:
		return "no disclosure without MOU"
	case SecurityLabelsNOORGPOL:
		return "no disclosure without organizational authorization"
	case SecurityLabelsNOPAT:
		return "no disclosure to patient, family or caregivers without attending provider's authorization"
	case SecurityLabelsNOPERSISTP:
		return "no collection beyond purpose of use"
	case SecurityLabelsNORDSCLCD:
		return "no redisclosure without consent directive"
	case SecurityLabelsNORDSLCD:
		return "no redisclosure without consent directive"
	case SecurityLabelsNORDSCLCDS:
		return "no redisclosure without information subject's consent directive"
	case SecurityLabelsNORDSCLW:
		return "no disclosure without jurisdictional authorization"
	case SecurityLabelsNORELINK:
		return "no relinking"
	case SecurityLabelsNOREUSE:
		return "no reuse beyond purpose of use"
	case SecurityLabelsNOVIP:
		return "no unauthorized VIP disclosure"
	case SecurityLabelsORCON:
		return "no disclosure without originator authorization"
	case SecurityLabelsHMARKT:
		return "healthcare marketing"
	case SecurityLabelsHOPERAT:
		return "healthcare operations"
	case SecurityLabelsCAREMGT:
		return "care management"
	case SecurityLabelsDONAT:
		return "donation"
	case SecurityLabelsFRAUD:
		return "fraud"
	case SecurityLabelsGOV:
		return "government"
	case SecurityLabelsHACCRED:
		return "health accreditation"
	case SecurityLabelsHCOMPL:
		return "health compliance"
	case SecurityLabelsHDECD:
		return "decedent"
	case SecurityLabelsHDIRECT:
		return "directory"
	case SecurityLabelsHDM:
		return "healthcare delivery management"
	case SecurityLabelsHLEGAL:
		return "legal"
	case SecurityLabelsHOUTCOMS:
		return "health outcome measure"
	case SecurityLabelsHPRGRP:
		return "health program reporting"
	case SecurityLabelsHQUALIMP:
		return "health quality improvement"
	case SecurityLabelsHSYSADMIN:
		return "health system administration"
	case SecurityLabelsLABELING:
		return "labeling"
	case SecurityLabelsMETAMGT:
		return "metadata management"
	case SecurityLabelsMEMADMIN:
		return "member administration"
	case SecurityLabelsMILCDM:
		return "military command"
	case SecurityLabelsPATADMIN:
		return "patient administration"
	case SecurityLabelsPATSFTY:
		return "patient safety"
	case SecurityLabelsPERFMSR:
		return "performance measure"
	case SecurityLabelsRECORDMGT:
		return "records management"
	case SecurityLabelsSYSDEV:
		return "system development"
	case SecurityLabelsHTEST:
		return "test health data"
	case SecurityLabelsTRAIN:
		return "training"
	case SecurityLabelsHPAYMT:
		return "healthcare payment"
	case SecurityLabelsCLMATTCH:
		return "claim attachment"
	case SecurityLabelsCOVAUTH:
		return "coverage authorization"
	case SecurityLabelsCOVERAGE:
		return "coverage under policy or program"
	case SecurityLabelsELIGDTRM:
		return "eligibility determination"
	case SecurityLabelsELIGVER:
		return "eligibility verification"
	case SecurityLabelsENROLLM:
		return "enrollment"
	case SecurityLabelsMILDCRG:
		return "military discharge"
	case SecurityLabelsREMITADV:
		return "remittance advice"
	case SecurityLabelsHRESCH:
		return "healthcare research"
	case SecurityLabelsBIORCH:
		return "biomedical research"
	case SecurityLabelsCLINTRCH:
		return "clinical trial research"
	case SecurityLabelsCLINTRCHNPC:
		return "clinical trial research without patient care"
	case SecurityLabelsCLINTRCHPC:
		return "clinical trial research with patient care"
	case SecurityLabelsPRECLINTRCH:
		return "preclinical trial research"
	case SecurityLabelsDSRCH:
		return "disease specific healthcare research"
	case SecurityLabelsPOARCH:
		return "population origins or ancestry healthcare research"
	case SecurityLabelsTRANSRCH:
		return "translational healthcare research"
	case SecurityLabelsPATRQT:
		return "patient requested"
	case SecurityLabelsFAMRQT:
		return "family requested"
	case SecurityLabelsPWATRNY:
		return "power of attorney"
	case SecurityLabelsSUPNWK:
		return "support network"
	case SecurityLabelsPUBHLTH:
		return "public health"
	case SecurityLabelsDISASTER:
		return "disaster"
	case SecurityLabelsTHREAT:
		return "threat"
	case SecurityLabelsTREAT:
		return "treatment"
	case SecurityLabelsCLINTRL:
		return "clinical trial"
	case SecurityLabelsCOC:
		return "coordination of care"
	case SecurityLabelsETREAT:
		return "Emergency Treatment"
	case SecurityLabelsBTG:
		return "break the glass"
	case SecurityLabelsERTREAT:
		return "emergency room treatment"
	case SecurityLabelsPOPHLTH:
		return "population health"
	case SecurityLabels_ActCoverageAssessmentObservationValue:
		return "ActCoverageAssessmentObservationValue"
	case SecurityLabels_ActFinancialStatusObservationValue:
		return "ActFinancialStatusObservationValue"
	case SecurityLabelsASSET:
		return "asset"
	case SecurityLabelsANNUITY:
		return "annuity"
	case SecurityLabelsPROP:
		return "real property"
	case SecurityLabelsRETACCT:
		return "retirement investment account"
	case SecurityLabelsTRUST:
		return "trust"
	case SecurityLabelsINCOME:
		return "income"
	case SecurityLabelsCHILD:
		return "child support"
	case SecurityLabelsDISABL:
		return "disability pay"
	case SecurityLabelsINVEST:
		return "investment income"
	case SecurityLabelsPAY:
		return "paid employment"
	case SecurityLabelsRETIRE:
		return "retirement pay"
	case SecurityLabelsSPOUSAL:
		return "spousal or partner support"
	case SecurityLabelsSUPPLE:
		return "income supplement"
	case SecurityLabelsTAX:
		return "tax obligation"
	case SecurityLabelsLIVEXP:
		return "living expense"
	case SecurityLabelsCLOTH:
		return "clothing expense"
	case SecurityLabelsFOOD:
		return "food expense"
	case SecurityLabelsHEALTH:
		return "health expense"
	case SecurityLabelsHOUSE:
		return "household expense"
	case SecurityLabelsLEGAL:
		return "legal expense"
	case SecurityLabelsMORTG:
		return "mortgage"
	case SecurityLabelsRENT:
		return "rent"
	case SecurityLabelsSUNDRY:
		return "sundry expense"
	case SecurityLabelsTRANS:
		return "transportation expense"
	case SecurityLabelsUTIL:
		return "utility expense"
	case SecurityLabelsELSTAT:
		return "eligibility indicator"
	case SecurityLabelsADOPT:
		return "adoption document"
	case SecurityLabelsBTHCERT:
		return "birth certificate"
	case SecurityLabelsCCOC:
		return "creditable coverage document"
	case SecurityLabelsDRLIC:
		return "driver license"
	case SecurityLabelsFOSTER:
		return "foster child document"
	case SecurityLabelsMEMBER:
		return "program or policy member"
	case SecurityLabelsMIL:
		return "military identification"
	case SecurityLabelsMRGCERT:
		return "marriage certificate"
	case SecurityLabelsPASSPORT:
		return "passport"
	case SecurityLabelsSTUDENRL:
		return "student enrollment"
	case SecurityLabelsHLSTAT:
		return "health status"
	case SecurityLabelsDISABLE:
		return "disabled"
	case SecurityLabelsDRUG:
		return "drug use"
	case SecurityLabelsIVDRG:
		return "IV drug use"
	case SecurityLabelsPGNT:
		return "pregnant"
	case SecurityLabelsLIVDEP:
		return "living dependency"
	case SecurityLabelsRELDEP:
		return "relative dependent"
	case SecurityLabelsSPSDEP:
		return "spouse dependent"
	case SecurityLabelsURELDEP:
		return "unrelated person dependent"
	case SecurityLabelsLIVSIT:
		return "living situation"
	case SecurityLabelsALONE:
		return "alone"
	case SecurityLabelsDEPCHD:
		return "dependent children"
	case SecurityLabelsDEPSPS:
		return "dependent spouse"
	case SecurityLabelsDEPYGCHD:
		return "dependent young children"
	case SecurityLabelsFAM:
		return "live with family"
	case SecurityLabelsRELAT:
		return "relative"
	case SecurityLabelsSPS:
		return "spouse only"
	case SecurityLabelsUNREL:
		return "unrelated person"
	case SecurityLabelsSOECSTAT:
		return "socio economic status"
	case SecurityLabelsABUSE:
		return "abuse victim"
	case SecurityLabelsHMLESS:
		return "homeless"
	case SecurityLabelsILGIM:
		return "illegal immigrant"
	case SecurityLabelsINCAR:
		return "incarcerated"
	case SecurityLabelsPROB:
		return "probation"
	case SecurityLabelsREFUG:
		return "refugee"
	case SecurityLabelsUNEMPL:
		return "unemployed"
	case SecurityLabels_AllergyTestValue:
		return "AllergyTestValue"
	case SecurityLabelsA0:
		return "no reaction"
	case SecurityLabelsA1:
		return "minimal reaction"
	case SecurityLabelsA2:
		return "mild reaction"
	case SecurityLabelsA3:
		return "moderate reaction"
	case SecurityLabelsA4:
		return "severe reaction"
	case SecurityLabels_CompositeMeasureScoring:
		return "CompositeMeasureScoring"
	case SecurityLabelsALLORNONESCR:
		return "All-or-nothing Scoring"
	case SecurityLabelsLINEARSCR:
		return "Linear Scoring"
	case SecurityLabelsOPPORSCR:
		return "Opportunity Scoring"
	case SecurityLabelsWEIGHTSCR:
		return "Weighted Scoring"
	case SecurityLabels_CoverageLimitObservationValue:
		return "CoverageLimitObservationValue"
	case SecurityLabels_CoverageLevelObservationValue:
		return "CoverageLevelObservationValue"
	case SecurityLabelsADC:
		return "adult child"
	case SecurityLabelsCHD:
		return "child"
	case SecurityLabelsDEP:
		return "dependent"
	case SecurityLabelsDP:
		return "domestic partner"
	case SecurityLabelsECH:
		return "employee"
	case SecurityLabelsFLY:
		return "family coverage"
	case SecurityLabelsIND:
		return "individual"
	case SecurityLabels_CoverageItemLimitObservationValue:
		return "CoverageItemLimitObservationValue"
	case SecurityLabels_CoverageLocationLimitObservationValue:
		return "CoverageLocationLimitObservationValue"
	case SecurityLabels_CriticalityObservationValue:
		return "CriticalityObservationValue"
	case SecurityLabelsCRITH:
		return "high criticality"
	case SecurityLabelsCRITL:
		return "low criticality"
	case SecurityLabelsCRITU:
		return "unable to assess criticality"
	case SecurityLabels_EmploymentStatus:
		return "_EmploymentStatus"
	case SecurityLabelsEmployed:
		return "Employed"
	case SecurityLabelsNotInLaborForce:
		return "Not In Labor Force"
	case SecurityLabelsUnemployed:
		return "Unemployed"
	case SecurityLabels_GeneticObservationValue:
		return "GeneticObservationValue"
	case SecurityLabelsHomozygote:
		return "HOMO"
	case SecurityLabels_MeasurementImprovementNotation:
		return "Measurement Improvement Notation"
	case SecurityLabelsDecrIsImp:
		return "Decreased score indicates improvement"
	case SecurityLabelsIncrIsImp:
		return "Increased score indicates improvement"
	case SecurityLabels_ObservationMeasureScoring:
		return "ObservationMeasureScoring"
	case SecurityLabelsCOHORT:
		return "cohort measure scoring"
	case SecurityLabelsCONTVAR:
		return "continuous variable measure scoring"
	case SecurityLabelsPROPOR:
		return "proportion measure scoring"
	case SecurityLabelsRATIO:
		return "ratio measure scoring"
	case SecurityLabels_ObservationMeasureType:
		return "ObservationMeasureType"
	case SecurityLabelsCOMPOSITE:
		return "composite measure type"
	case SecurityLabelsEFFICIENCY:
		return "efficiency measure type"
	case SecurityLabelsEXPERIENCE:
		return "experience measure type"
	case SecurityLabelsOUTCOME:
		return "outcome measure type"
	case SecurityLabelsINTERMOM:
		return "intermediate clinical outcome measure"
	case SecurityLabelsPROPM:
		return "patient reported outcome performance measure"
	case SecurityLabelsPROCESS:
		return "process measure type"
	case SecurityLabelsAPPROPRIATE:
		return "appropriate use process measure"
	case SecurityLabelsRESOURCE:
		return "resource use measure type"
	case SecurityLabelsSTRUCTURE:
		return "structure measure type"
	case SecurityLabels_ObservationPopulationInclusion:
		return "ObservationPopulationInclusion"
	case SecurityLabelsDENEX:
		return "denominator exclusions"
	case SecurityLabelsDENEXCEP:
		return "denominator exceptions"
	case SecurityLabelsDENOM:
		return "denominator"
	case SecurityLabelsIP:
		return "initial population"
	case SecurityLabelsIPP:
		return "initial patient population"
	case SecurityLabelsMSRPOPL:
		return "measure population"
	case SecurityLabelsNUMER:
		return "numerator"
	case SecurityLabelsNUMEX:
		return "numerator exclusions"
	case SecurityLabels_PartialCompletionScale:
		return "PartialCompletionScale"
	case SecurityLabelsG:
		return "Great extent"
	case SecurityLabelsLE:
		return "Large extent"
	case SecurityLabelsME:
		return "Medium extent"
	case SecurityLabelsMI:
		return "Minimal extent"
	case SecurityLabelsS:
		return "Some extent"
	case SecurityLabels_SecurityObservationValue:
		return "SecurityObservationValue"
	case SecurityLabels_SECCATOBV:
		return "security category"
	case SecurityLabels_SECCLASSOBV:
		return "security classification"
	case SecurityLabels_SECCONOBV:
		return "security control"
	case SecurityLabels_SECINTOBV:
		return "security integrity"
	case SecurityLabelsSECTRSTOBV:
		return "security trust observation"
	case SecurityLabelsTRSTACCRDOBV:
		return "trust accreditation observation"
	case SecurityLabelsTRSTAGREOBV:
		return "trust agreement observation"
	case SecurityLabelsTRSTCERTOBV:
		return "trust certificate observation"
	case SecurityLabelsTRSTFWKOBV:
		return "none supplied 5"
	case SecurityLabelsTRSTLOAOBV:
		return "trust assurance observation"
	case SecurityLabelsLOAAN:
		return "authentication level of assurance value"
	case SecurityLabelsLOAAN1:
		return "low authentication level of assurance"
	case SecurityLabelsLOAAN2:
		return "basic authentication level of assurance"
	case SecurityLabelsLOAAN3:
		return "medium authentication level of assurance"
	case SecurityLabelsLOAAN4:
		return "high authentication level of assurance"
	case SecurityLabelsLOAAP:
		return "authentication process level of assurance value"
	case SecurityLabelsLOAAP1:
		return "low authentication process level of assurance"
	case SecurityLabelsLOAAP2:
		return "basic authentication process level of assurance"
	case SecurityLabelsLOAAP3:
		return "medium authentication process level of assurance"
	case SecurityLabelsLOAAP4:
		return "high authentication process level of assurance"
	case SecurityLabelsLOAAS:
		return "assertion level of assurance value"
	case SecurityLabelsLOAAS1:
		return "low assertion level of assurance"
	case SecurityLabelsLOAAS2:
		return "basic assertion level of assurance"
	case SecurityLabelsLOAAS3:
		return "medium assertion level of assurance"
	case SecurityLabelsLOAAS4:
		return "high assertion level of assurance"
	case SecurityLabelsLOACM:
		return "token and credential management level of assurance value)"
	case SecurityLabelsLOACM1:
		return "low token and credential management level of assurance"
	case SecurityLabelsLOACM2:
		return "basic token and credential management level of assurance"
	case SecurityLabelsLOACM3:
		return "medium token and credential management level of assurance"
	case SecurityLabelsLOACM4:
		return "high token and credential management level of assurance"
	case SecurityLabelsLOAID:
		return "identity proofing level of assurance"
	case SecurityLabelsLOAID1:
		return "low identity proofing level of assurance"
	case SecurityLabelsLOAID2:
		return "basic identity proofing level of assurance"
	case SecurityLabelsLOAID3:
		return "medium identity proofing level of assurance"
	case SecurityLabelsLOAID4:
		return "high identity proofing level of assurance"
	case SecurityLabelsLOANR:
		return "non-repudiation level of assurance value"
	case SecurityLabelsLOANR1:
		return "low non-repudiation level of assurance"
	case SecurityLabelsLOANR2:
		return "basic non-repudiation level of assurance"
	case SecurityLabelsLOANR3:
		return "medium non-repudiation level of assurance"
	case SecurityLabelsLOANR4:
		return "high non-repudiation level of assurance"
	case SecurityLabelsLOARA:
		return "remote access level of assurance value"
	case SecurityLabelsLOARA1:
		return "low remote access level of assurance"
	case SecurityLabelsLOARA2:
		return "basic remote access level of assurance"
	case SecurityLabelsLOARA3:
		return "medium remote access level of assurance"
	case SecurityLabelsLOARA4:
		return "high remote access level of assurance"
	case SecurityLabelsLOATK:
		return "token level of assurance value"
	case SecurityLabelsLOATK1:
		return "low token level of assurance"
	case SecurityLabelsLOATK2:
		return "basic token level of assurance"
	case SecurityLabelsLOATK3:
		return "medium token level of assurance"
	case SecurityLabelsLOATK4:
		return "high token level of assurance"
	case SecurityLabelsTRSTMECOBV:
		return "none supplied 6"
	case SecurityLabels_SeverityObservation:
		return "SeverityObservation"
	case SecurityLabelsH:
		return "High"
	case SecurityLabels_SubjectBodyPosition:
		return "_SubjectBodyPosition"
	case SecurityLabelsLLD:
		return "left lateral decubitus"
	case SecurityLabelsPRN:
		return "prone"
	case SecurityLabelsRLD:
		return "right lateral decubitus"
	case SecurityLabelsSFWL:
		return "Semi-Fowler's"
	case SecurityLabelsSIT:
		return "sitting"
	case SecurityLabelsSTN:
		return "standing"
	case SecurityLabelsSUP:
		return "supine"
	case SecurityLabelsRTRD:
		return "reverse trendelenburg"
	case SecurityLabelsTRD:
		return "trendelenburg"
	case SecurityLabels_VerificationOutcomeValue:
		return "verification outcome"
	case SecurityLabelsACT:
		return "active coverage"
	case SecurityLabelsACTPEND:
		return "active - pending investigation"
	case SecurityLabelsELG:
		return "eligible"
	case SecurityLabelsINACT:
		return "inactive"
	case SecurityLabelsINPNDINV:
		return "inactive - pending investigation"
	case SecurityLabelsINPNDUPD:
		return "inactive - pending eligibility update"
	case SecurityLabelsNELG:
		return "not eligible"
	case SecurityLabels_WorkSchedule:
		return "_WorkSchedule"
	case SecurityLabelsDS:
		return "daytime shift"
	case SecurityLabelsEMS:
		return "early morning shift"
	case SecurityLabelsES:
		return "evening shift"
	case SecurityLabelsNS:
		return "night shift"
	case SecurityLabelsRSWN:
		return "rotating shift with nights"
	case SecurityLabelsRSWON:
		return "rotating shift without nights"
	case SecurityLabelsSS:
		return "split shift"
	case SecurityLabelsVLS:
		return "very long shift"
	case SecurityLabelsVS:
		return "variable shift"
	case SecurityLabels_AnnotationValue:
		return "AnnotationValue"
	case SecurityLabels_ECGAnnotationValue:
		return "ECGAnnotationValue"
	case SecurityLabels_CommonClinicalObservationValue:
		return "common clinical observation"
	case SecurityLabels_CommonClinicalObservationAssertionValue:
		return "CommonClinicalObservationAssertionValue"
	case SecurityLabels_CommonClinicalObservationResultValue:
		return "CommonClinicalObservationResultValue"
	case SecurityLabels_CoverageChemicalDependencyValue:
		return "CoverageChemicalDependencyValue"
	case SecurityLabels_IndividualCaseSafetyReportValueDomains:
		return "Individual Case Safety Report Value Domains"
	case SecurityLabels_CaseSeriousnessCriteria:
		return "CaseSeriousnessCriteria"
	case SecurityLabels_DeviceManufacturerEvaluationInterpretation:
		return "DeviceManufacturerEvaluationInterpretation"
	case SecurityLabels_DeviceManufacturerEvaluationMethod:
		return "DeviceManufacturerEvaluationMethod"
	case SecurityLabels_DeviceManufacturerEvaluationResult:
		return "DeviceManufacturerEvaluationResult"
	case SecurityLabels_PertinentReactionRelatedness:
		return "Pertinent Reaction Relatedness"
	case SecurityLabels_ProductCharacterization:
		return "Product Characterization"
	case SecurityLabels_ReactionActionTaken:
		return "ReactionActionTaken"
	case SecurityLabels_SubjectReaction:
		return "Subject Reaction"
	case SecurityLabels_SubjectReactionEmphasis:
		return "SubjectReactionEmphasis"
	case SecurityLabels_SubjectReactionOutcome:
		return "SubjectReactionOutcome"
	case SecurityLabels_InjuryObservationValue:
		return "InjuryObservationValue"
	case SecurityLabels_IntoleranceValue:
		return "IntoleranceValue"
	case SecurityLabels_IssueTriggerObservationValue:
		return "IssueTriggerObservationValue"
	case SecurityLabels_OtherIndicationValue:
		return "OtherIndicationValue"
	case SecurityLabels_IndicationValue:
		return "IndicationValue"
	case SecurityLabels_DiagnosisValue:
		return "DiagnosisValue"
	case SecurityLabels_SymptomValue:
		return "SymptomValue"
	case SecurityLabels_ActUSPrivacyLaw:
		return "ActUSPrivacyLaw"
	case SecurityLabels42CFRPart2:
		return "42 CFR Part2"
	case SecurityLabelsCommonRule:
		return "Common Rule"
	case SecurityLabelsHIPAANOPP:
		return "HIPAA notice of privacy practices"
	case SecurityLabelsHIPAAPsyNotes:
		return "HIPAA psychotherapy notes"
	case SecurityLabelsHIPAASelfPay:
		return "HIPAA self-pay"
	case SecurityLabelsTitle38Section7332:
		return "Title 38 Section 7332"
	case SecurityLabelsTitle38Part1:
		return "Title 38 Section 7332"

	default:
		return "Unknown Security Labels"
	}
}
