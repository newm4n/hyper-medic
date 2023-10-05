package code

import (
	"fmt"
	"strings"
)

type AnimalTissueType string

const (
	AnimalTissueTypeAllRelevantTissues                AnimalTissueType = "100000072091"
	AnimalTissueTypeFat                               AnimalTissueType = "100000072092"
	AnimalTissueTypeHoney                             AnimalTissueType = "100000072093"
	AnimalTissueTypeLiver                             AnimalTissueType = "100000072094"
	AnimalTissueTypeFreshMilk                         AnimalTissueType = "100000072095"
	AnimalTissueTypeMuscleAndSkinInNaturalProportions AnimalTissueType = "100000072096"
	AnimalTissueTypeEggs                              AnimalTissueType = "100000072104"
	AnimalTissueTypeSkinAndFat                        AnimalTissueType = "100000072105"
	AnimalTissueTypeKidney                            AnimalTissueType = "100000072106"
	AnimalTissueTypeMeatAndOffal                      AnimalTissueType = "100000072107"
	AnimalTissueTypeMuscle                            AnimalTissueType = "100000072108"
	AnimalTissueTypeUnspecified                       AnimalTissueType = "100000072109"
	AnimalTissueTypeAdiposeTissue                     AnimalTissueType = "100000111053"
	AnimalTissueTypeAdrenal                           AnimalTissueType = "100000111054"
	AnimalTissueTypeBloodVessels                      AnimalTissueType = "100000111055"
	AnimalTissueTypeBone                              AnimalTissueType = "100000111056"
	AnimalTissueTypeBoneMarrow                        AnimalTissueType = "100000111057"
	AnimalTissueTypeBrain                             AnimalTissueType = "100000111058"
	AnimalTissueTypeConnectiveTissue                  AnimalTissueType = "100000111059"
	AnimalTissueTypeCornea                            AnimalTissueType = "100000111060"
	AnimalTissueTypeDentalPulp                        AnimalTissueType = "100000111061"
	AnimalTissueTypeDuodenum                          AnimalTissueType = "100000111062"
	AnimalTissueTypeDuraMater                         AnimalTissueType = "100000111063"
	AnimalTissueTypeEggEmbryonated                    AnimalTissueType = "100000111064"
	AnimalTissueTypeEgg                               AnimalTissueType = "100000111065"
	AnimalTissueTypeEggWhite                          AnimalTissueType = "100000111066"
	AnimalTissueTypeEggYolk                           AnimalTissueType = "100000111067"
	AnimalTissueTypeEmbryos                           AnimalTissueType = "100000111068"
	AnimalTissueTypeEntericPlexuses                   AnimalTissueType = "100000111069"
	AnimalTissueTypeEsophagus                         AnimalTissueType = "100000111070"
	AnimalTissueTypeFeathers                          AnimalTissueType = "100000111071"
	AnimalTissueTypeFoetus                            AnimalTissueType = "100000111072"
	AnimalTissueTypeForeStomachRuminantsOnly          AnimalTissueType = "100000111073"
	AnimalTissueTypeGingivalTissue                    AnimalTissueType = "100000111074"
	AnimalTissueTypeHair                              AnimalTissueType = "100000111075"
	AnimalTissueTypeHeartPericardium                  AnimalTissueType = "100000111076"
	AnimalTissueTypeHide                              AnimalTissueType = "100000111077"
	AnimalTissueTypeHooves                            AnimalTissueType = "100000111078"
	AnimalTissueTypeIleum                             AnimalTissueType = "100000111079"
	AnimalTissueTypeJejunum                           AnimalTissueType = "100000111080"
	AnimalTissueTypeKidney2                           AnimalTissueType = "100000111081"
	AnimalTissueTypeLardLardOil                       AnimalTissueType = "100000111082"
	AnimalTissueTypeLargeIntestine                    AnimalTissueType = "100000111083"
	Liver                                             AnimalTissueType = "100000111084"
	AnimalTissueTypeLung                              AnimalTissueType = "100000111085"
	AnimalTissueTypeLymphNodes                        AnimalTissueType = "100000111086"
	AnimalTissueTypeMammaryGland                      AnimalTissueType = "100000111087"
	AnimalTissueTypeUdder                             AnimalTissueType = "100000111088"
	AnimalTissueTypeMammaryTumour                     AnimalTissueType = "100000111089"
	AnimalTissueTypeMeatExtract                       AnimalTissueType = "100000111090"
	AnimalTissueTypeNasopharyngeal                    AnimalTissueType = "100000111091"
	AnimalTissueTypeNictitatingMembrane               AnimalTissueType = "100000111092"
	AnimalTissueTypeNasalMucosa                       AnimalTissueType = "100000111093"
	AnimalTissueTypeOvary                             AnimalTissueType = "100000111094"
	AnimalTissueTypePancreas                          AnimalTissueType = "100000111095"
	AnimalTissueTypePeripheralNerves                  AnimalTissueType = "100000111096"
	AnimalTissueTypePituitaryGland                    AnimalTissueType = "100000111097"
	AnimalTissueTypePlacenta                          AnimalTissueType = "100000111098"
	AnimalTissueTypeProstate                          AnimalTissueType = "100000111099"
	AnimalTissueTypeEpididymis                        AnimalTissueType = "100000111100"
	AnimalTissueTypeSeminalVesicle                    AnimalTissueType = "100000111101"
	AnimalTissueTypeRennetCalf                        AnimalTissueType = "100000111102"
	AnimalTissueTypeRetina                            AnimalTissueType = "100000111103"
	AnimalTissueTypeOpticNerve                        AnimalTissueType = "100000111104"
	AnimalTissueTypeSalivaryGland                     AnimalTissueType = "100000111105"
	AnimalTissueTypeShank                             AnimalTissueType = "100000111106"
	AnimalTissueTypeSkeletalMuscle                    AnimalTissueType = "100000111107"
	AnimalTissueTypeSkin                              AnimalTissueType = "100000111108"
	AnimalTissueTypeSpinalGanglia                     AnimalTissueType = "100000111109"
	AnimalTissueTypeSpinalCord                        AnimalTissueType = "100000111110"
	AnimalTissueTypeSpleen                            AnimalTissueType = "100000111111"
	AnimalTissueTypeStomach                           AnimalTissueType = "100000111112"
	AnimalTissueTypeAbomasum                          AnimalTissueType = "100000111113"
	AnimalTissueTypeSubmaxillaryGlands                AnimalTissueType = "100000111114"
	AnimalTissueTypeTallow                            AnimalTissueType = "100000111115"
	AnimalTissueTypeTendon                            AnimalTissueType = "100000111116"
	AnimalTissueTypeTestis                            AnimalTissueType = "100000111117"
	AnimalTissueTypeThymus                            AnimalTissueType = "100000111118"
	AnimalTissueTypeThyroidGland                      AnimalTissueType = "100000111119"
	AnimalTissueTypeTongue                            AnimalTissueType = "100000111120"
	AnimalTissueTypeTonsil                            AnimalTissueType = "100000111121"
	AnimalTissueTypeTrachea                           AnimalTissueType = "100000111122"
	AnimalTissueTypeTrigeminalGanglia                 AnimalTissueType = "100000111123"
	AnimalTissueTypeTripe                             AnimalTissueType = "100000111124"
	AnimalTissueTypeUterusNonGravid                   AnimalTissueType = "100000111125"
	AnimalTissueTypeWool                              AnimalTissueType = "100000111126"
	AnimalTissueTypeAscitesFluid                      AnimalTissueType = "100000111127"
	AnimalTissueTypeBile                              AnimalTissueType = "100000111128"
	AnimalTissueTypeBlood1                            AnimalTissueType = "100000111129"
	AnimalTissueTypeBloodFoetal                       AnimalTissueType = "100000111130"
	AnimalTissueTypeColostrum                         AnimalTissueType = "100000111131"
	AnimalTissueTypeCordBlood                         AnimalTissueType = "100000111132"
	AnimalTissueTypeCsf                               AnimalTissueType = "100000111133"
	AnimalTissueTypeFaeces                            AnimalTissueType = "100000111134"
	AnimalTissueTypeMilk                              AnimalTissueType = "100000111135"
	AnimalTissueTypeNasalMucus                        AnimalTissueType = "100000111136"
	AnimalTissueTypePlacentaFluids                    AnimalTissueType = "100000111137"
	AnimalTissueTypePlasma                            AnimalTissueType = "100000111138"
	AnimalTissueTypeSaliva                            AnimalTissueType = "100000111139"
	AnimalTissueTypeSecretionFromBees                 AnimalTissueType = "100000111140"
	AnimalTissueTypeSemen                             AnimalTissueType = "100000111141"
	AnimalTissueTypeSerumCalf                         AnimalTissueType = "100000111142"
	AnimalTissueTypeSerumDonorAdultBovine             AnimalTissueType = "100000111143"
	AnimalTissueTypeSerumDonorCalf                    AnimalTissueType = "100000111144"
	AnimalTissueTypeSerumFoetalBovine                 AnimalTissueType = "100000111145"
	AnimalTissueTypeSerumNewbornCalf                  AnimalTissueType = "100000111146"
	AnimalTissueTypeSerumPlasmaDerivateAdultBovine    AnimalTissueType = "100000111147"
	AnimalTissueTypeSerumPlasmaAdultBovine            AnimalTissueType = "100000111148"
	AnimalTissueTypeSweat                             AnimalTissueType = "100000111149"
	AnimalTissueTypeTears                             AnimalTissueType = "100000111150"
	AnimalTissueTypeUrine                             AnimalTissueType = "100000111151"
	AnimalTissueTypeVenom                             AnimalTissueType = "100000111152"
	AnimalTissueTypeWhey                              AnimalTissueType = "100000111153"
	AnimalTissueTypeCasein                            AnimalTissueType = "100000111154"
	AnimalTissueTypeFermentationProducts              AnimalTissueType = "100000111155"
	AnimalTissueTypeGelatin                           AnimalTissueType = "100000111156"
	AnimalTissueTypeLactose                           AnimalTissueType = "100000111157"
	AnimalTissueTypeProtein                           AnimalTissueType = "100000111158"
	AnimalTissueTypeInsulin                           AnimalTissueType = "100000111159"
	AnimalTissueTypeCollagen                          AnimalTissueType = "100000111160"
	AnimalTissueTypeAnimalCharcoal                    AnimalTissueType = "100000111161"
	AnimalTissueTypePeptones                          AnimalTissueType = "100000111162"
	AnimalTissueTypeFattyAcids                        AnimalTissueType = "100000111163"
	AnimalTissueTypeGlycerol                          AnimalTissueType = "100000111164"
	AnimalTissueTypeNotApplicable                     AnimalTissueType = "100000125717"
	AnimalTissueTypeMeatAndOffalMilk                  AnimalTissueType = "100000136180"
	AnimalTissueTypeAgarBlood                         AnimalTissueType = "100000136181"
	AnimalTissueTypeCasaminoAcid                      AnimalTissueType = "100000136182"
	AnimalTissueTypeCaseinHydrolysate                 AnimalTissueType = "100000136183"
	AnimalTissueTypeCaseinPancreaticDigest            AnimalTissueType = "100000136184"
	AnimalTissueTypeCaseinPeptidesN3                  AnimalTissueType = "100000136185"
	AnimalTissueTypeCells                             AnimalTissueType = "100000136186"
	AnimalTissueTypeCellsBhk21                        AnimalTissueType = "100000136187"
	AnimalTissueTypeCellsCho                          AnimalTissueType = "100000136188"
	AnimalTissueTypeCellsCrfk                         AnimalTissueType = "100000136189"
	AnimalTissueTypeCellsEmbryoSpf                    AnimalTissueType = "100000136190"
	AnimalTissueTypeCellsIrc5                         AnimalTissueType = "100000136191"
	AnimalTissueTypeCellsKidney                       AnimalTissueType = "100000136192"
	AnimalTissueTypeCellsMdck                         AnimalTissueType = "100000136193"
	AnimalTissueTypeCellsRedBlood                     AnimalTissueType = "100000136194"
	AnimalTissueTypeCollagenHydrolysate               AnimalTissueType = "100000136195"
	AnimalTissueTypeCholesterol                       AnimalTissueType = "100000136196"
	AnimalTissueTypeEggSpfEmbryonated                 AnimalTissueType = "100000136197"
	AnimalTissueTypeEnzyme                            AnimalTissueType = "100000136198"
	AnimalTissueTypeEnzymePancreaticEnzymes           AnimalTissueType = "100000136199"
	AnimalTissueTypeEnzymePancreatin6Nf               AnimalTissueType = "100000136200"
	AnimalTissueTypeEnzymePepsin                      AnimalTissueType = "100000136201"
	AnimalTissueTypeEnzymePronase                     AnimalTissueType = "100000136202"
	AnimalTissueTypeEnzymeTrypsin                     AnimalTissueType = "100000136203"
	AnimalTissueTypeHeartDigest                       AnimalTissueType = "100000136204"
	AnimalTissueTypeHeartExtract                      AnimalTissueType = "100000136205"
	AnimalTissueTypeIntestinalMucosae                 AnimalTissueType = "100000136206"
	AnimalTissueTypeLactalbuminHydrolysate            AnimalTissueType = "100000136207"
	AnimalTissueTypeLiverDigest                       AnimalTissueType = "100000136208"
	AnimalTissueTypeLymphocytes                       AnimalTissueType = "100000136209"
	AnimalTissueTypeMeat                              AnimalTissueType = "100000136210"
	AnimalTissueTypeMeatEnzymicHydrolysate            AnimalTissueType = "100000136211"
	AnimalTissueTypeMediumCookedMeat                  AnimalTissueType = "100000136212"
	AnimalTissueTypeMediumF10199Medium                AnimalTissueType = "100000136213"
	AnimalTissueTypeMediumFmdCultureMedium            AnimalTissueType = "100000136214"
	AnimalTissueTypeMediumGlasgowMemCulture           AnimalTissueType = "100000136215"
	AnimalTissueTypeMediumLbAgarLennox                AnimalTissueType = "100000136216"
	AnimalTissueTypeMediumLbBrothLennox               AnimalTissueType = "100000136217"
	AnimalTissueTypeMediumModifiedThioglycolateMedium AnimalTissueType = "100000136218"
	AnimalTissueTypeMediumTrypticaseSoyBroth          AnimalTissueType = "100000136219"
	AnimalTissueTypeMediumTryptosePhosphateBroth      AnimalTissueType = "100000136220"
	AnimalTissueTypeMilkSkimmed                       AnimalTissueType = "100000136221"
	AnimalTissueTypePancreasExtract                   AnimalTissueType = "100000136222"
	AnimalTissueTypePeptonesCaseinHydrochloricPeptone AnimalTissueType = "100000136223"
	AnimalTissueTypePeptonesCaseinTrypticPeptone      AnimalTissueType = "100000136224"
	AnimalTissueTypePituitaryExtract                  AnimalTissueType = "100000136225"
	AnimalTissueTypeRennet                            AnimalTissueType = "100000136226"
	AnimalTissueTypeMediumNutrientBroth               AnimalTissueType = "100000136227"
	AnimalTissueTypeMediumNzAmine                     AnimalTissueType = "100000136228"
	AnimalTissueTypeMediumThioglycolateMedium         AnimalTissueType = "100000136229"
	AnimalTissueTypePeptonesProteosePeptone           AnimalTissueType = "100000136230"
	AnimalTissueTypeSerum                             AnimalTissueType = "100000136231"
	AnimalTissueTypeSerumAlbumin                      AnimalTissueType = "100000136232"
	AnimalTissueTypeSerumIronFortifiedCalf            AnimalTissueType = "100000136233"
	AnimalTissueTypeSkinConnectiveTissueAndBone       AnimalTissueType = "100000136234"
	AnimalTissueTypeSperm                             AnimalTissueType = "100000136235"
	AnimalTissueTypeTryptone                          AnimalTissueType = "100000136236"
	AnimalTissueTypeMeatExtractDesiccated             AnimalTissueType = "100000136237"
	AnimalTissueTypeStomachMucosa                     AnimalTissueType = "100000136247"
	AnimalTissueTypeTransferin                        AnimalTissueType = "100000136248"
	AnimalTissueTypeNonNeural                         AnimalTissueType = "100000136554"
	AnimalTissueTypeNotSpecified                      AnimalTissueType = "100000136555"
	AnimalTissueTypeOrganTissue                       AnimalTissueType = "100000136556"
	AnimalTissueTypeSkinAndFatInNaturalProportions    AnimalTissueType = "100000142485"
)

func AllAnimalTissueType() []AnimalTissueType {
	return []AnimalTissueType{
		AnimalTissueTypeAllRelevantTissues,
		AnimalTissueTypeFat,
		AnimalTissueTypeHoney,
		AnimalTissueTypeLiver,
		AnimalTissueTypeFreshMilk,
		AnimalTissueTypeMuscleAndSkinInNaturalProportions,
		AnimalTissueTypeEggs,
		AnimalTissueTypeSkinAndFat,
		AnimalTissueTypeKidney,
		AnimalTissueTypeMeatAndOffal,
		AnimalTissueTypeMuscle,
		AnimalTissueTypeUnspecified,
		AnimalTissueTypeAdiposeTissue,
		AnimalTissueTypeAdrenal,
		AnimalTissueTypeBloodVessels,
		AnimalTissueTypeBone,
		AnimalTissueTypeBoneMarrow,
		AnimalTissueTypeBrain,
		AnimalTissueTypeConnectiveTissue,
		AnimalTissueTypeCornea,
		AnimalTissueTypeDentalPulp,
		AnimalTissueTypeDuodenum,
		AnimalTissueTypeDuraMater,
		AnimalTissueTypeEggEmbryonated,
		AnimalTissueTypeEgg,
		AnimalTissueTypeEggWhite,
		AnimalTissueTypeEggYolk,
		AnimalTissueTypeEmbryos,
		AnimalTissueTypeEntericPlexuses,
		AnimalTissueTypeEsophagus,
		AnimalTissueTypeFeathers,
		AnimalTissueTypeFoetus,
		AnimalTissueTypeForeStomachRuminantsOnly,
		AnimalTissueTypeGingivalTissue,
		AnimalTissueTypeHair,
		AnimalTissueTypeHeartPericardium,
		AnimalTissueTypeHide,
		AnimalTissueTypeHooves,
		AnimalTissueTypeIleum,
		AnimalTissueTypeJejunum,
		AnimalTissueTypeLardLardOil,
		AnimalTissueTypeLargeIntestine,
		AnimalTissueTypeLung,
		AnimalTissueTypeLymphNodes,
		AnimalTissueTypeMammaryGland,
		AnimalTissueTypeUdder,
		AnimalTissueTypeMammaryTumour,
		AnimalTissueTypeMeatExtract,
		AnimalTissueTypeNasopharyngeal,
		AnimalTissueTypeNictitatingMembrane,
		AnimalTissueTypeNasalMucosa,
		AnimalTissueTypeOvary,
		AnimalTissueTypePancreas,
		AnimalTissueTypePeripheralNerves,
		AnimalTissueTypePituitaryGland,
		AnimalTissueTypePlacenta,
		AnimalTissueTypeProstate,
		AnimalTissueTypeEpididymis,
		AnimalTissueTypeSeminalVesicle,
		AnimalTissueTypeRennetCalf,
		AnimalTissueTypeRetina,
		AnimalTissueTypeOpticNerve,
		AnimalTissueTypeSalivaryGland,
		AnimalTissueTypeShank,
		AnimalTissueTypeSkeletalMuscle,
		AnimalTissueTypeSkin,
		AnimalTissueTypeSpinalGanglia,
		AnimalTissueTypeSpinalCord,
		AnimalTissueTypeSpleen,
		AnimalTissueTypeStomach,
		AnimalTissueTypeAbomasum,
		AnimalTissueTypeSubmaxillaryGlands,
		AnimalTissueTypeTallow,
		AnimalTissueTypeTendon,
		AnimalTissueTypeTestis,
		AnimalTissueTypeThymus,
		AnimalTissueTypeThyroidGland,
		AnimalTissueTypeTongue,
		AnimalTissueTypeTonsil,
		AnimalTissueTypeTrachea,
		AnimalTissueTypeTrigeminalGanglia,
		AnimalTissueTypeTripe,
		AnimalTissueTypeUterusNonGravid,
		AnimalTissueTypeWool,
		AnimalTissueTypeAscitesFluid,
		AnimalTissueTypeBile,
		AnimalTissueTypeBlood1,
		AnimalTissueTypeBloodFoetal,
		AnimalTissueTypeColostrum,
		AnimalTissueTypeCordBlood,
		AnimalTissueTypeCsf,
		AnimalTissueTypeFaeces,
		AnimalTissueTypeMilk,
		AnimalTissueTypeNasalMucus,
		AnimalTissueTypePlacentaFluids,
		AnimalTissueTypePlasma,
		AnimalTissueTypeSaliva,
		AnimalTissueTypeSecretionFromBees,
		AnimalTissueTypeSemen,
		AnimalTissueTypeSerumCalf,
		AnimalTissueTypeSerumDonorAdultBovine,
		AnimalTissueTypeSerumDonorCalf,
		AnimalTissueTypeSerumFoetalBovine,
		AnimalTissueTypeSerumNewbornCalf,
		AnimalTissueTypeSerumPlasmaDerivateAdultBovine,
		AnimalTissueTypeSerumPlasmaAdultBovine,
		AnimalTissueTypeSweat,
		AnimalTissueTypeTears,
		AnimalTissueTypeUrine,
		AnimalTissueTypeVenom,
		AnimalTissueTypeWhey,
		AnimalTissueTypeCasein,
		AnimalTissueTypeFermentationProducts,
		AnimalTissueTypeGelatin,
		AnimalTissueTypeLactose,
		AnimalTissueTypeProtein,
		AnimalTissueTypeInsulin,
		AnimalTissueTypeCollagen,
		AnimalTissueTypeAnimalCharcoal,
		AnimalTissueTypePeptones,
		AnimalTissueTypeFattyAcids,
		AnimalTissueTypeGlycerol,
		AnimalTissueTypeNotApplicable,
		AnimalTissueTypeMeatAndOffalMilk,
		AnimalTissueTypeAgarBlood,
		AnimalTissueTypeCasaminoAcid,
		AnimalTissueTypeCaseinHydrolysate,
		AnimalTissueTypeCaseinPancreaticDigest,
		AnimalTissueTypeCaseinPeptidesN3,
		AnimalTissueTypeCells,
		AnimalTissueTypeCellsBhk21,
		AnimalTissueTypeCellsCho,
		AnimalTissueTypeCellsCrfk,
		AnimalTissueTypeCellsEmbryoSpf,
		AnimalTissueTypeCellsIrc5,
		AnimalTissueTypeCellsKidney,
		AnimalTissueTypeCellsMdck,
		AnimalTissueTypeCellsRedBlood,
		AnimalTissueTypeCollagenHydrolysate,
		AnimalTissueTypeCholesterol,
		AnimalTissueTypeEggSpfEmbryonated,
		AnimalTissueTypeEnzyme,
		AnimalTissueTypeEnzymePancreaticEnzymes,
		AnimalTissueTypeEnzymePancreatin6Nf,
		AnimalTissueTypeEnzymePepsin,
		AnimalTissueTypeEnzymePronase,
		AnimalTissueTypeEnzymeTrypsin,
		AnimalTissueTypeHeartDigest,
		AnimalTissueTypeHeartExtract,
		AnimalTissueTypeIntestinalMucosae,
		AnimalTissueTypeLactalbuminHydrolysate,
		AnimalTissueTypeLiverDigest,
		AnimalTissueTypeLymphocytes,
		AnimalTissueTypeMeat,
		AnimalTissueTypeMeatEnzymicHydrolysate,
		AnimalTissueTypeMediumCookedMeat,
		AnimalTissueTypeMediumF10199Medium,
		AnimalTissueTypeMediumFmdCultureMedium,
		AnimalTissueTypeMediumGlasgowMemCulture,
		AnimalTissueTypeMediumLbAgarLennox,
		AnimalTissueTypeMediumLbBrothLennox,
		AnimalTissueTypeMediumModifiedThioglycolateMedium,
		AnimalTissueTypeMediumTrypticaseSoyBroth,
		AnimalTissueTypeMediumTryptosePhosphateBroth,
		AnimalTissueTypeMilkSkimmed,
		AnimalTissueTypePancreasExtract,
		AnimalTissueTypePeptonesCaseinHydrochloricPeptone,
		AnimalTissueTypePeptonesCaseinTrypticPeptone,
		AnimalTissueTypePituitaryExtract,
		AnimalTissueTypeRennet,
		AnimalTissueTypeMediumNutrientBroth,
		AnimalTissueTypeMediumNzAmine,
		AnimalTissueTypeMediumThioglycolateMedium,
		AnimalTissueTypePeptonesProteosePeptone,
		AnimalTissueTypeSerum,
		AnimalTissueTypeSerumAlbumin,
		AnimalTissueTypeSerumIronFortifiedCalf,
		AnimalTissueTypeSkinConnectiveTissueAndBone,
		AnimalTissueTypeSperm,
		AnimalTissueTypeTryptone,
		AnimalTissueTypeMeatExtractDesiccated,
		AnimalTissueTypeStomachMucosa,
		AnimalTissueTypeTransferin,
		AnimalTissueTypeNonNeural,
		AnimalTissueTypeNotSpecified,
		AnimalTissueTypeOrganTissue,
		AnimalTissueTypeSkinAndFatInNaturalProportions,
	}
}

func FindAnimalTissueType(filter string) []AnimalTissueType {
	ret := make([]AnimalTissueType, 0)
	for _, at := range AllAnimalTissueType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AnimalTissueType) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AnimalTissueType) String() string {
	switch cpt {
	case AnimalTissueTypeAllRelevantTissues:
		return "All relevant tissues"
	case AnimalTissueTypeFat:
		return "Fat"
	case AnimalTissueTypeHoney:
		return "Honey"
	case AnimalTissueTypeLiver:
		return "Liver"
	case AnimalTissueTypeFreshMilk:
		return "Fresh Milk"
	case AnimalTissueTypeMuscleAndSkinInNaturalProportions:
		return "Muscle and skin in natural proportions"
	case AnimalTissueTypeEggs:
		return "Eggs"
	case AnimalTissueTypeSkinAndFat:
		return "Skin and fat"
	case AnimalTissueTypeKidney:
		return "Kidney"
	case AnimalTissueTypeMeatAndOffal:
		return "Meat and offal"
	case AnimalTissueTypeMuscle:
		return "Muscle"
	case AnimalTissueTypeUnspecified:
		return "Unspecified"
	case AnimalTissueTypeAdiposeTissue:
		return "Adipose tissue"
	case AnimalTissueTypeAdrenal:
		return "Adrenal"
	case AnimalTissueTypeBloodVessels:
		return "Blood vessels"
	case AnimalTissueTypeBone:
		return "Bone"
	case AnimalTissueTypeBoneMarrow:
		return "Bone marrow"
	case AnimalTissueTypeBrain:
		return "Brain"
	case AnimalTissueTypeConnectiveTissue:
		return "Connective tissue"
	case AnimalTissueTypeCornea:
		return "Cornea"
	case AnimalTissueTypeDentalPulp:
		return "Dental pulp"
	case AnimalTissueTypeDuodenum:
		return "Duodenum"
	case AnimalTissueTypeDuraMater:
		return "Dura mater"
	case AnimalTissueTypeEggEmbryonated:
		return "Egg, embryonated"
	case AnimalTissueTypeEgg:
		return "Egg"
	case AnimalTissueTypeEggWhite:
		return "Egg white"
	case AnimalTissueTypeEggYolk:
		return "Egg yolk"
	case AnimalTissueTypeEmbryos:
		return "Embryos"
	case AnimalTissueTypeEntericPlexuses:
		return "Enteric plexuses"
	case AnimalTissueTypeEsophagus:
		return "Esophagus"
	case AnimalTissueTypeFeathers:
		return "Feathers"
	case AnimalTissueTypeFoetus:
		return "Foetus"
	case AnimalTissueTypeForeStomachRuminantsOnly:
		return "Fore-stomach (ruminants only)"
	case AnimalTissueTypeGingivalTissue:
		return "Gingival tissue"
	case AnimalTissueTypeHair:
		return "Hair"
	case AnimalTissueTypeHeartPericardium:
		return "Heart/pericardium"
	case AnimalTissueTypeHide:
		return "Hide"
	case AnimalTissueTypeHooves:
		return "Hooves"
	case AnimalTissueTypeIleum:
		return "Ileum"
	case AnimalTissueTypeJejunum:
		return "Jejunum"
	case AnimalTissueTypeLardLardOil:
		return "Lard/lard oil"
	case AnimalTissueTypeLargeIntestine:
		return "Large intestine"
	case AnimalTissueTypeLung:
		return "Lung"
	case AnimalTissueTypeLymphNodes:
		return "Lymph nodes"
	case AnimalTissueTypeMammaryGland:
		return "Mammary gland"
	case AnimalTissueTypeUdder:
		return "Udder"
	case AnimalTissueTypeMammaryTumour:
		return "Mammary tumour"
	case AnimalTissueTypeMeatExtract:
		return "Meat extract"
	case AnimalTissueTypeNasopharyngeal:
		return "Nasopharyngeal"
	case AnimalTissueTypeNictitatingMembrane:
		return "Nictitating membrane"
	case AnimalTissueTypeNasalMucosa:
		return "Nasal mucosa"
	case AnimalTissueTypeOvary:
		return "Ovary"
	case AnimalTissueTypePancreas:
		return "Pancreas"
	case AnimalTissueTypePeripheralNerves:
		return "Peripheral nerves"
	case AnimalTissueTypePituitaryGland:
		return "Pituitary gland"
	case AnimalTissueTypePlacenta:
		return "Placenta"
	case AnimalTissueTypeProstate:
		return "Prostate"
	case AnimalTissueTypeEpididymis:
		return "Epididymis"
	case AnimalTissueTypeSeminalVesicle:
		return "Seminal vesicle"
	case AnimalTissueTypeRennetCalf:
		return "Rennet, calf"
	case AnimalTissueTypeRetina:
		return "Retina"
	case AnimalTissueTypeOpticNerve:
		return "Optic nerve"
	case AnimalTissueTypeSalivaryGland:
		return "Salivary gland"
	case AnimalTissueTypeShank:
		return "Shank"
	case AnimalTissueTypeSkeletalMuscle:
		return "Skeletal muscle"
	case AnimalTissueTypeSkin:
		return "Skin"
	case AnimalTissueTypeSpinalGanglia:
		return "Spinal ganglia"
	case AnimalTissueTypeSpinalCord:
		return "Spinal cord"
	case AnimalTissueTypeSpleen:
		return "Spleen"
	case AnimalTissueTypeStomach:
		return "Stomach"
	case AnimalTissueTypeAbomasum:
		return "Abomasum"
	case AnimalTissueTypeSubmaxillaryGlands:
		return "Submaxillary glands"
	case AnimalTissueTypeTallow:
		return "Tallow"
	case AnimalTissueTypeTendon:
		return "Tendon"
	case AnimalTissueTypeTestis:
		return "Testis"
	case AnimalTissueTypeThymus:
		return "Thymus"
	case AnimalTissueTypeThyroidGland:
		return "Thyroid gland"
	case AnimalTissueTypeTongue:
		return "Tongue"
	case AnimalTissueTypeTonsil:
		return "Tonsil"
	case AnimalTissueTypeTrachea:
		return "Trachea"
	case AnimalTissueTypeTrigeminalGanglia:
		return "Trigeminal ganglia"
	case AnimalTissueTypeTripe:
		return "Tripe"
	case AnimalTissueTypeUterusNonGravid:
		return "Uterus (Non-gravid)"
	case AnimalTissueTypeWool:
		return "Wool"
	case AnimalTissueTypeAscitesFluid:
		return "Ascites fluid"
	case AnimalTissueTypeBile:
		return "Bile"
	case AnimalTissueTypeBlood1:
		return "Blood1"
	case AnimalTissueTypeBloodFoetal:
		return "Blood, foetal"
	case AnimalTissueTypeColostrum:
		return "Colostrum"
	case AnimalTissueTypeCordBlood:
		return "Cord blood"
	case AnimalTissueTypeCsf:
		return "CSF"
	case AnimalTissueTypeFaeces:
		return "Faeces"
	case AnimalTissueTypeMilk:
		return "Milk"
	case AnimalTissueTypeNasalMucus:
		return "Nasal mucus"
	case AnimalTissueTypePlacentaFluids:
		return "Placenta fluids"
	case AnimalTissueTypePlasma:
		return "Plasma"
	case AnimalTissueTypeSaliva:
		return "Saliva"
	case AnimalTissueTypeSecretionFromBees:
		return "Secretion from bees"
	case AnimalTissueTypeSemen:
		return "Semen"
	case AnimalTissueTypeSerumCalf:
		return "Serum, calf"
	case AnimalTissueTypeSerumDonorAdultBovine:
		return "Serum, donor adult bovine"
	case AnimalTissueTypeSerumDonorCalf:
		return "Serum, donor calf"
	case AnimalTissueTypeSerumFoetalBovine:
		return "Serum, foetal bovine"
	case AnimalTissueTypeSerumNewbornCalf:
		return "Serum, newborn calf"
	case AnimalTissueTypeSerumPlasmaDerivateAdultBovine:
		return "Serum/plasma derivate, adult bovine"
	case AnimalTissueTypeSerumPlasmaAdultBovine:
		return "Serum/plasma, adult bovine"
	case AnimalTissueTypeSweat:
		return "Sweat"
	case AnimalTissueTypeTears:
		return "Tears"
	case AnimalTissueTypeUrine:
		return "Urine"
	case AnimalTissueTypeVenom:
		return "Venom"
	case AnimalTissueTypeWhey:
		return "Whey"
	case AnimalTissueTypeCasein:
		return "Casein"
	case AnimalTissueTypeFermentationProducts:
		return "Fermentation products"
	case AnimalTissueTypeGelatin:
		return "Gelatin"
	case AnimalTissueTypeLactose:
		return "Lactose"
	case AnimalTissueTypeProtein:
		return "Protein"
	case AnimalTissueTypeInsulin:
		return "Insulin"
	case AnimalTissueTypeCollagen:
		return "Collagen"
	case AnimalTissueTypeAnimalCharcoal:
		return "Animal Charcoal"
	case AnimalTissueTypePeptones:
		return "Peptones"
	case AnimalTissueTypeFattyAcids:
		return "Fatty acids"
	case AnimalTissueTypeGlycerol:
		return "Glycerol"
	case AnimalTissueTypeNotApplicable:
		return "Not applicable"
	case AnimalTissueTypeMeatAndOffalMilk:
		return "Meat and offal, milk"
	case AnimalTissueTypeAgarBlood:
		return "Agar blood"
	case AnimalTissueTypeCasaminoAcid:
		return "Casamino acid"
	case AnimalTissueTypeCaseinHydrolysate:
		return "Casein, hydrolysate"
	case AnimalTissueTypeCaseinPancreaticDigest:
		return "Casein, pancreatic digest"
	case AnimalTissueTypeCaseinPeptidesN3:
		return "Casein, peptides N3"
	case AnimalTissueTypeCells:
		return "Cells"
	case AnimalTissueTypeCellsBhk21:
		return "Cells, BHK21"
	case AnimalTissueTypeCellsCho:
		return "Cells, CHO"
	case AnimalTissueTypeCellsCrfk:
		return "Cells, CRFK"
	case AnimalTissueTypeCellsEmbryoSpf:
		return "Cells, embryo SPF"
	case AnimalTissueTypeCellsIrc5:
		return "Cells, IRC5"
	case AnimalTissueTypeCellsKidney:
		return "Cells, kidney"
	case AnimalTissueTypeCellsMdck:
		return "Cells, MDCK"
	case AnimalTissueTypeCellsRedBlood:
		return "Cells, red blood"
	case AnimalTissueTypeCollagenHydrolysate:
		return "Collagen, hydrolysate"
	case AnimalTissueTypeCholesterol:
		return "Cholesterol"
	case AnimalTissueTypeEggSpfEmbryonated:
		return "Egg, SPF embryonated"
	case AnimalTissueTypeEnzyme:
		return "Enzyme"
	case AnimalTissueTypeEnzymePancreaticEnzymes:
		return "Enzyme, pancreatic enzymes"
	case AnimalTissueTypeEnzymePancreatin6Nf:
		return "Enzyme, pancreatin 6NF"
	case AnimalTissueTypeEnzymePepsin:
		return "Enzyme, pepsin"
	case AnimalTissueTypeEnzymePronase:
		return "Enzyme, pronase"
	case AnimalTissueTypeEnzymeTrypsin:
		return "Enzyme, trypsin"
	case AnimalTissueTypeHeartDigest:
		return "Heart, digest"
	case AnimalTissueTypeHeartExtract:
		return "Heart, extract"
	case AnimalTissueTypeIntestinalMucosae:
		return "Intestinal mucosae"
	case AnimalTissueTypeLactalbuminHydrolysate:
		return "Lactalbumin hydrolysate"
	case AnimalTissueTypeLiverDigest:
		return "Liver, digest"
	case AnimalTissueTypeLymphocytes:
		return "Lymphocytes"
	case AnimalTissueTypeMeat:
		return "Meat"
	case AnimalTissueTypeMeatEnzymicHydrolysate:
		return "Meat, enzymic hydrolysate"
	case AnimalTissueTypeMediumCookedMeat:
		return "Medium, cooked meat"
	case AnimalTissueTypeMediumF10199Medium:
		return "Medium, F10-199 medium"
	case AnimalTissueTypeMediumFmdCultureMedium:
		return "Medium, FMD culture medium"
	case AnimalTissueTypeMediumGlasgowMemCulture:
		return "Medium, Glasgow MEM culture"
	case AnimalTissueTypeMediumLbAgarLennox:
		return "Medium, LB Agar Lennox"
	case AnimalTissueTypeMediumLbBrothLennox:
		return "Medium, LB Broth Lennox"
	case AnimalTissueTypeMediumModifiedThioglycolateMedium:
		return "Medium, modified thioglycolate medium"
	case AnimalTissueTypeMediumTrypticaseSoyBroth:
		return "Medium, trypticase soy broth"
	case AnimalTissueTypeMediumTryptosePhosphateBroth:
		return "Medium, tryptose phosphate broth"
	case AnimalTissueTypeMilkSkimmed:
		return "Milk, skimmed"
	case AnimalTissueTypePancreasExtract:
		return "Pancreas, extract"
	case AnimalTissueTypePeptonesCaseinHydrochloricPeptone:
		return "Peptones, casein hydrochloric peptone"
	case AnimalTissueTypePeptonesCaseinTrypticPeptone:
		return "Peptones, casein tryptic peptone"
	case AnimalTissueTypePituitaryExtract:
		return "Pituitary extract"
	case AnimalTissueTypeRennet:
		return "Rennet"
	case AnimalTissueTypeMediumNutrientBroth:
		return "Medium, nutrient broth"
	case AnimalTissueTypeMediumNzAmine:
		return "Medium, NZ-Amine"
	case AnimalTissueTypeMediumThioglycolateMedium:
		return "Medium, thioglycolate medium"
	case AnimalTissueTypePeptonesProteosePeptone:
		return "Peptones, proteose peptone"
	case AnimalTissueTypeSerum:
		return "Serum"
	case AnimalTissueTypeSerumAlbumin:
		return "Serum, albumin"
	case AnimalTissueTypeSerumIronFortifiedCalf:
		return "Serum, Iron fortified calf"
	case AnimalTissueTypeSkinConnectiveTissueAndBone:
		return "Skin, connective tissue and bone"
	case AnimalTissueTypeSperm:
		return "Sperm"
	case AnimalTissueTypeTryptone:
		return "Tryptone"
	case AnimalTissueTypeMeatExtractDesiccated:
		return "Meat, extract desiccated"
	case AnimalTissueTypeStomachMucosa:
		return "Stomach mucosa"
	case AnimalTissueTypeTransferin:
		return "Transferin"
	case AnimalTissueTypeNonNeural:
		return "Non-neural"
	case AnimalTissueTypeNotSpecified:
		return "Not specified"
	case AnimalTissueTypeOrganTissue:
		return "Organ tissue"
	case AnimalTissueTypeSkinAndFatInNaturalProportions:
		return "Skin and fat in natural proportions"
	case AnimalTissueTypeKidney2:
		return "Kidney"
	default:
		return "Unknown AnimalTissueType"
	}
}

/*
All relevant tissues
Fat
Honey
Liver
Fresh Milk
Muscle and skin in natural proportions
Eggs
Skin and fat
Kidney
Meat and offal
Muscle
Unspecified
Adipose tissue
Adrenal
Blood vessels
Bone
Bone marrow
Brain
Connective tissue
Cornea
Dental pulp
Duodenum
Dura mater
Egg, embryonated
Egg
Egg white
Egg yolk
Embryos
Enteric plexuses
Esophagus
Feathers
Foetus
Fore-stomach (ruminants only)
Gingival tissue
Hair
Heart/pericardium
Hide
Hooves
Ileum
Jejunum
Kidney
Lard/lard oil
Large intestine
Liver
Lung
Lymph nodes
Mammary gland
Udder
Mammary tumour
Meat extract
Nasopharyngeal
Nictitating membrane
Nasal mucosa
Ovary
Pancreas
Peripheral nerves
Pituitary gland
Placenta
Prostate
Epididymis
Seminal vesicle
Rennet, calf
Retina
Optic nerve
Salivary gland
Shank
Skeletal muscle
Skin
Spinal ganglia
Spinal cord
Spleen
Stomach
Abomasum
Submaxillary glands
Tallow
Tendon
Testis
Thymus
Thyroid gland
Tongue
Tonsil
Trachea
Trigeminal ganglia
Tripe
Uterus (Non-gravid)
Wool
Ascites fluid
Bile
Blood1
Blood, foetal
Colostrum
Cord blood
CSF
Faeces
Milk
Nasal mucus
Placenta fluids
Plasma
Saliva
Secretion from bees
Semen
Serum, calf
Serum, donor adult bovine
Serum, donor calf
Serum, foetal bovine
Serum, newborn calf
Serum/plasma derivate, adult bovine
Serum/plasma, adult bovine
Sweat
Tears
Urine
Venom
Whey
Casein
Fermentation products
Gelatin
Lactose
Protein
Insulin
Collagen
Animal Charcoal
Peptones
Fatty acids
Glycerol
Not applicable
Meat and offal, milk
Agar blood
Casamino acid
Casein, hydrolysate
Casein, pancreatic digest
Casein, peptides N3
Cells
Cells, BHK21
Cells, CHO
Cells, CRFK
Cells, embryo SPF
Cells, IRC5
Cells, kidney
Cells, MDCK
Cells, red blood
Collagen, hydrolysate
Cholesterol
Egg, SPF embryonated
Enzyme
Enzyme, pancreatic enzymes
Enzyme, pancreatin 6NF
Enzyme, pepsin
Enzyme, pronase
Enzyme, trypsin
Heart, digest
Heart, extract
Intestinal mucosae
Lactalbumin hydrolysate
Liver, digest
Lymphocytes
Meat
Meat, enzymic hydrolysate
Medium, cooked meat
Medium, F10-199 medium
Medium, FMD culture medium
Medium, Glasgow MEM culture
Medium, LB Agar Lennox
Medium, LB Broth Lennox
Medium, modified thioglycolate medium
Medium, trypticase soy broth
Medium, tryptose phosphate broth
Milk, skimmed
Pancreas, extract
Peptones, casein hydrochloric peptone
Peptones, casein tryptic peptone
Pituitary extract
Rennet
Medium, nutrient broth
Medium, NZ-Amine
Medium, thioglycolate medium
Peptones, proteose peptone
Serum
Serum, albumin
Serum, Iron fortified calf
Skin, connective tissue and bone
Sperm
Tryptone
Meat, extract desiccated
Stomach mucosa
Transferin
Non-neural
Not specified
Organ tissue
Skin and fat in natural proportions

*/
