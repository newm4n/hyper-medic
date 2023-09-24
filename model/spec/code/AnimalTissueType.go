package code

import (
	"fmt"
	"strings"
)

type AnimalTissueType int

const (
	AnimalTissueTypeAllRelevantTissues AnimalTissueType = iota
	AnimalTissueTypeFat
	AnimalTissueTypeHoney
	AnimalTissueTypeLiver
	AnimalTissueTypeFreshMilk
	AnimalTissueTypeMuscleAndSkinInNaturalProportions
	AnimalTissueTypeEggs
	AnimalTissueTypeSkinAndFat
	AnimalTissueTypeKidney
	AnimalTissueTypeMeatAndOffal
	AnimalTissueTypeMuscle
	AnimalTissueTypeUnspecified
	AnimalTissueTypeAdiposeTissue
	AnimalTissueTypeAdrenal
	AnimalTissueTypeBloodVessels
	AnimalTissueTypeBone
	AnimalTissueTypeBoneMarrow
	AnimalTissueTypeBrain
	AnimalTissueTypeConnectiveTissue
	AnimalTissueTypeCornea
	AnimalTissueTypeDentalPulp
	AnimalTissueTypeDuodenum
	AnimalTissueTypeDuraMater
	AnimalTissueTypeEggEmbryonated
	AnimalTissueTypeEgg
	AnimalTissueTypeEggWhite
	AnimalTissueTypeEggYolk
	AnimalTissueTypeEmbryos
	AnimalTissueTypeEntericPlexuses
	AnimalTissueTypeEsophagus
	AnimalTissueTypeFeathers
	AnimalTissueTypeFoetus
	AnimalTissueTypeForeStomachRuminantsOnly
	AnimalTissueTypeGingivalTissue
	AnimalTissueTypeHair
	AnimalTissueTypeHeartPericardium
	AnimalTissueTypeHide
	AnimalTissueTypeHooves
	AnimalTissueTypeIleum
	AnimalTissueTypeJejunum
	AnimalTissueTypeLardLardOil
	AnimalTissueTypeLargeIntestine
	AnimalTissueTypeLung
	AnimalTissueTypeLymphNodes
	AnimalTissueTypeMammaryGland
	AnimalTissueTypeUdder
	AnimalTissueTypeMammaryTumour
	AnimalTissueTypeMeatExtract
	AnimalTissueTypeNasopharyngeal
	AnimalTissueTypeNictitatingMembrane
	AnimalTissueTypeNasalMucosa
	AnimalTissueTypeOvary
	AnimalTissueTypePancreas
	AnimalTissueTypePeripheralNerves
	AnimalTissueTypePituitaryGland
	AnimalTissueTypePlacenta
	AnimalTissueTypeProstate
	AnimalTissueTypeEpididymis
	AnimalTissueTypeSeminalVesicle
	AnimalTissueTypeRennetCalf
	AnimalTissueTypeRetina
	AnimalTissueTypeOpticNerve
	AnimalTissueTypeSalivaryGland
	AnimalTissueTypeShank
	AnimalTissueTypeSkeletalMuscle
	AnimalTissueTypeSkin
	AnimalTissueTypeSpinalGanglia
	AnimalTissueTypeSpinalCord
	AnimalTissueTypeSpleen
	AnimalTissueTypeStomach
	AnimalTissueTypeAbomasum
	AnimalTissueTypeSubmaxillaryGlands
	AnimalTissueTypeTallow
	AnimalTissueTypeTendon
	AnimalTissueTypeTestis
	AnimalTissueTypeThymus
	AnimalTissueTypeThyroidGland
	AnimalTissueTypeTongue
	AnimalTissueTypeTonsil
	AnimalTissueTypeTrachea
	AnimalTissueTypeTrigeminalGanglia
	AnimalTissueTypeTripe
	AnimalTissueTypeUterusNonGravid
	AnimalTissueTypeWool
	AnimalTissueTypeAscitesFluid
	AnimalTissueTypeBile
	AnimalTissueTypeBlood1
	AnimalTissueTypeBloodFoetal
	AnimalTissueTypeColostrum
	AnimalTissueTypeCordBlood
	AnimalTissueTypeCsf
	AnimalTissueTypeFaeces
	AnimalTissueTypeMilk
	AnimalTissueTypeNasalMucus
	AnimalTissueTypePlacentaFluids
	AnimalTissueTypePlasma
	AnimalTissueTypeSaliva
	AnimalTissueTypeSecretionFromBees
	AnimalTissueTypeSemen
	AnimalTissueTypeSerumCalf
	AnimalTissueTypeSerumDonorAdultBovine
	AnimalTissueTypeSerumDonorCalf
	AnimalTissueTypeSerumFoetalBovine
	AnimalTissueTypeSerumNewbornCalf
	AnimalTissueTypeSerumPlasmaDerivateAdultBovine
	AnimalTissueTypeSerumPlasmaAdultBovine
	AnimalTissueTypeSweat
	AnimalTissueTypeTears
	AnimalTissueTypeUrine
	AnimalTissueTypeVenom
	AnimalTissueTypeWhey
	AnimalTissueTypeCasein
	AnimalTissueTypeFermentationProducts
	AnimalTissueTypeGelatin
	AnimalTissueTypeLactose
	AnimalTissueTypeProtein
	AnimalTissueTypeInsulin
	AnimalTissueTypeCollagen
	AnimalTissueTypeAnimalCharcoal
	AnimalTissueTypePeptones
	AnimalTissueTypeFattyAcids
	AnimalTissueTypeGlycerol
	AnimalTissueTypeNotApplicable
	AnimalTissueTypeMeatAndOffalMilk
	AnimalTissueTypeAgarBlood
	AnimalTissueTypeCasaminoAcid
	AnimalTissueTypeCaseinHydrolysate
	AnimalTissueTypeCaseinPancreaticDigest
	AnimalTissueTypeCaseinPeptidesN3
	AnimalTissueTypeCells
	AnimalTissueTypeCellsBhk21
	AnimalTissueTypeCellsCho
	AnimalTissueTypeCellsCrfk
	AnimalTissueTypeCellsEmbryoSpf
	AnimalTissueTypeCellsIrc5
	AnimalTissueTypeCellsKidney
	AnimalTissueTypeCellsMdck
	AnimalTissueTypeCellsRedBlood
	AnimalTissueTypeCollagenHydrolysate
	AnimalTissueTypeCholesterol
	AnimalTissueTypeEggSpfEmbryonated
	AnimalTissueTypeEnzyme
	AnimalTissueTypeEnzymePancreaticEnzymes
	AnimalTissueTypeEnzymePancreatin6Nf
	AnimalTissueTypeEnzymePepsin
	AnimalTissueTypeEnzymePronase
	AnimalTissueTypeEnzymeTrypsin
	AnimalTissueTypeHeartDigest
	AnimalTissueTypeHeartExtract
	AnimalTissueTypeIntestinalMucosae
	AnimalTissueTypeLactalbuminHydrolysate
	AnimalTissueTypeLiverDigest
	AnimalTissueTypeLymphocytes
	AnimalTissueTypeMeat
	AnimalTissueTypeMeatEnzymicHydrolysate
	AnimalTissueTypeMediumCookedMeat
	AnimalTissueTypeMediumF10199Medium
	AnimalTissueTypeMediumFmdCultureMedium
	AnimalTissueTypeMediumGlasgowMemCulture
	AnimalTissueTypeMediumLbAgarLennox
	AnimalTissueTypeMediumLbBrothLennox
	AnimalTissueTypeMediumModifiedThioglycolateMedium
	AnimalTissueTypeMediumTrypticaseSoyBroth
	AnimalTissueTypeMediumTryptosePhosphateBroth
	AnimalTissueTypeMilkSkimmed
	AnimalTissueTypePancreasExtract
	AnimalTissueTypePeptonesCaseinHydrochloricPeptone
	AnimalTissueTypePeptonesCaseinTrypticPeptone
	AnimalTissueTypePituitaryExtract
	AnimalTissueTypeRennet
	AnimalTissueTypeMediumNutrientBroth
	AnimalTissueTypeMediumNzAmine
	AnimalTissueTypeMediumThioglycolateMedium
	AnimalTissueTypePeptonesProteosePeptone
	AnimalTissueTypeSerum
	AnimalTissueTypeSerumAlbumin
	AnimalTissueTypeSerumIronFortifiedCalf
	AnimalTissueTypeSkinConnectiveTissueAndBone
	AnimalTissueTypeSperm
	AnimalTissueTypeTryptone
	AnimalTissueTypeMeatExtractDesiccated
	AnimalTissueTypeStomachMucosa
	AnimalTissueTypeTransferin
	AnimalTissueTypeNonNeural
	AnimalTissueTypeNotSpecified
	AnimalTissueTypeOrganTissue
	AnimalTissueTypeSkinAndFatInNaturalProportions
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
