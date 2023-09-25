package code

import (
	"fmt"
	"strings"
)

type AdministrableDoseForm int

const (
	AdministrableDoseFormOralSuspension AdministrableDoseForm = iota
	AdministrableDoseFormOralGel
	AdministrableDoseFormPowderForOralSolution
	AdministrableDoseFormGranulesForOralSolution
	AdministrableDoseFormLyophilisateForSuspension
	AdministrableDoseFormPowderForSyrup
	AdministrableDoseFormSolubleTablet
	AdministrableDoseFormHerbalTea
	AdministrableDoseFormInstantHerbalTea
	AdministrableDoseFormGranules
	AdministrableDoseFormGastroResistantGranules
	AdministrableDoseFormModifiedReleaseGranules
	AdministrableDoseFormCapsuleHard
	AdministrableDoseFormGastroResistantCapsuleHard
	AdministrableDoseFormChewableCapsuleSoft
	AdministrableDoseFormProlongedReleaseCapsuleSoft
	AdministrableDoseFormModifiedReleaseCapsuleSoft
	AdministrableDoseFormCoatedTablet
	AdministrableDoseFormOralDropsSolution
	AdministrableDoseFormOralDropsSuspension
	AdministrableDoseFormOralDropsEmulsion
	AdministrableDoseFormOralLiquid
	AdministrableDoseFormOralSolution
	AdministrableDoseFormOralEmulsion
	AdministrableDoseFormOralPaste
	AdministrableDoseFormPowderForOralSuspension
	AdministrableDoseFormGranulesForOralSuspension
	AdministrableDoseFormSyrup
	AdministrableDoseFormGranulesForSyrup
	AdministrableDoseFormDispersibleTablet
	AdministrableDoseFormOralPowder
	AdministrableDoseFormEffervescentPowder
	AdministrableDoseFormEffervescentGranules
	AdministrableDoseFormProlongedReleaseGranules
	AdministrableDoseFormCachet
	AdministrableDoseFormCapsuleSoft
	AdministrableDoseFormGastroResistantCapsuleSoft
	AdministrableDoseFormProlongedReleaseCapsuleHard
	AdministrableDoseFormModifiedReleaseCapsuleHard
	AdministrableDoseFormTablet
	AdministrableDoseFormFilmCoatedTablet
	AdministrableDoseFormOrodispersibleTablet
	AdministrableDoseFormGastroResistantTablet
	AdministrableDoseFormModifiedReleaseTablet
	AdministrableDoseFormMedicatedChewingGum
	AdministrableDoseFormPillules
	AdministrableDoseFormPulsatileReleaseIntraruminalDevice
	AdministrableDoseFormPremixForMedicatedFeedingStuff
	AdministrableDoseFormGargle
	AdministrableDoseFormGarglePowderForSolution
	AdministrableDoseFormOromucosalSuspension
	AdministrableDoseFormOromucosalSpray
	AdministrableDoseFormMouthwash
	AdministrableDoseFormGingivalSolution
	AdministrableDoseFormOromucosalPaste
	AdministrableDoseFormGingivalGel
	AdministrableDoseFormEffervescentTablet
	AdministrableDoseFormOralLyophilisate
	AdministrableDoseFormProlongedReleaseTablet
	AdministrableDoseFormChewableTablet
	AdministrableDoseFormOralGum
	AdministrableDoseFormContinuousReleaseIntraruminalDevice
	AdministrableDoseFormLickBlock
	AdministrableDoseFormMedicatedPellets
	AdministrableDoseFormConcentrateForGargle
	AdministrableDoseFormGargleTabletForSolution
	AdministrableDoseFormOromucosalSolution
	AdministrableDoseFormOromucosalDrops
	AdministrableDoseFormSublingualSpray
	AdministrableDoseFormMouthwashTabletForSolution
	AdministrableDoseFormOromucosalGel
	AdministrableDoseFormOromucosalCream
	AdministrableDoseFormGingivalPaste
	AdministrableDoseFormSublingualTablet
	AdministrableDoseFormBuccalTablet
	AdministrableDoseFormCompressedLozenge
	AdministrableDoseFormOromucosalCapsule
	AdministrableDoseFormMucoAdhesiveBuccalTablet
	AdministrableDoseFormLozenge
	AdministrableDoseFormPastille
	AdministrableDoseFormDentalGel
	AdministrableDoseFormDentalInsert
	AdministrableDoseFormDentalPowder
	AdministrableDoseFormDentalSuspension
	AdministrableDoseFormToothpaste
	AdministrableDoseFormPeriodontalGel
	AdministrableDoseFormBathAdditive
	AdministrableDoseFormCream
	AdministrableDoseFormOintment
	AdministrableDoseFormMedicatedPlaster
	AdministrableDoseFormShampoo
	AdministrableDoseFormCutaneousSpraySuspension
	AdministrableDoseFormCutaneousLiquid
	AdministrableDoseFormConcentrateForCutaneousSolution
	AdministrableDoseFormCutaneousEmulsion
	AdministrableDoseFormCutaneousPatch
	AdministrableDoseFormPeriodontalPowder
	AdministrableDoseFormDentalStick
	AdministrableDoseFormDentalSolution
	AdministrableDoseFormDentalEmulsion
	AdministrableDoseFormPeriodontalInsert
	AdministrableDoseFormGel
	AdministrableDoseFormCutaneousPaste
	AdministrableDoseFormCutaneousFoam
	AdministrableDoseFormCutaneousSpraySolution
	AdministrableDoseFormCutaneousSprayPowder
	AdministrableDoseFormCutaneousSolution
	AdministrableDoseFormCutaneousSuspension
	AdministrableDoseFormCutaneousPowder
	AdministrableDoseFormSolutionForIontophoresis
	AdministrableDoseFormCollodion
	AdministrableDoseFormPoultice
	AdministrableDoseFormCutaneousSponge
	AdministrableDoseFormCollar
	AdministrableDoseFormEarTag
	AdministrableDoseFormDipSuspension
	AdministrableDoseFormTransdermalPatch
	AdministrableDoseFormMedicatedNailLacquer
	AdministrableDoseFormCutaneousStick
	AdministrableDoseFormImpregnatedDressing
	AdministrableDoseFormMedicatedPendant
	AdministrableDoseFormDipSolution
	AdministrableDoseFormDipEmulsion
	AdministrableDoseFormConcentrateForDipSuspension
	AdministrableDoseFormPowderForDipSolution
	AdministrableDoseFormPowderForSuspensionForFishTreatment
	AdministrableDoseFormPourOnSuspension
	AdministrableDoseFormSpotOnSolution
	AdministrableDoseFormSpotOnEmulsion
	AdministrableDoseFormTeatDipSuspension
	AdministrableDoseFormTeatSpraySolution
	AdministrableDoseFormSolutionForSkinPrickTest
	AdministrableDoseFormPlasterForProvocationTest
	AdministrableDoseFormEyeGel
	AdministrableDoseFormEyeDropsSolution
	AdministrableDoseFormEyeDropsSuspension
	AdministrableDoseFormConcentrateForDipSolution
	AdministrableDoseFormConcentrateForDipEmulsion
	AdministrableDoseFormConcentrateForSolutionForFishTreatment
	AdministrableDoseFormPourOnSolution
	AdministrableDoseFormPourOnEmulsion
	AdministrableDoseFormSpotOnSuspension
	AdministrableDoseFormTeatDipSolution
	AdministrableDoseFormTeatDipEmulsion
	AdministrableDoseFormTransdermalSystem
	AdministrableDoseFormSolutionForSkinScratchTest
	AdministrableDoseFormEyeCream
	AdministrableDoseFormEyeOintment
	AdministrableDoseFormEyeDropsEmulsion
	AdministrableDoseFormEyeDropsSolventForReconstitution
	AdministrableDoseFormEyeLotion
	AdministrableDoseFormOphthalmicInsert
	AdministrableDoseFormEarCream
	AdministrableDoseFormEarOintment
	AdministrableDoseFormEarDropsSuspension
	AdministrableDoseFormEyeDropsProlongedRelease
	AdministrableDoseFormEyeLotionSolventForReconstitution
	AdministrableDoseFormOphthalmicStrip
	AdministrableDoseFormEarGel
	AdministrableDoseFormEarDropsSolution
	AdministrableDoseFormEarDropsEmulsion
	AdministrableDoseFormEarPowder
	AdministrableDoseFormEarSpraySuspension
	AdministrableDoseFormEarWashSolution
	AdministrableDoseFormEarTampon
	AdministrableDoseFormNasalCream
	AdministrableDoseFormNasalGel
	AdministrableDoseFormNasalDropsSolution
	AdministrableDoseFormNasalDropsEmulsion
	AdministrableDoseFormNasalSpraySolution
	AdministrableDoseFormNasalSprayEmulsion
	AdministrableDoseFormNasalStick
	AdministrableDoseFormVaginalGel
	AdministrableDoseFormVaginalFoam
	AdministrableDoseFormEarSpraySolution
	AdministrableDoseFormEarSprayEmulsion
	AdministrableDoseFormEarWashEmulsion
	AdministrableDoseFormEarStick
	AdministrableDoseFormNasalOintment
	AdministrableDoseFormNasalDropsSuspension
	AdministrableDoseFormNasalPowder
	AdministrableDoseFormNasalSpraySuspension
	AdministrableDoseFormNasalWash
	AdministrableDoseFormVaginalCream
	AdministrableDoseFormVaginalOintment
	AdministrableDoseFormVaginalSolution
	AdministrableDoseFormVaginalEmulsion
	AdministrableDoseFormPessary
	AdministrableDoseFormVaginalCapsuleSoft
	AdministrableDoseFormEffervescentVaginalTablet
	AdministrableDoseFormVaginalDeliverySystem
	AdministrableDoseFormRectalCream
	AdministrableDoseFormRectalFoam
	AdministrableDoseFormVaginalSuspension
	AdministrableDoseFormTabletForVaginalSolution
	AdministrableDoseFormVaginalCapsuleHard
	AdministrableDoseFormVaginalTablet
	AdministrableDoseFormMedicatedVaginalTampon
	AdministrableDoseFormVaginalSponge
	AdministrableDoseFormRectalGel
	AdministrableDoseFormSolutionForInjection
)

func AllAdministrableDoseForm() []AdministrableDoseForm {
	return []AdministrableDoseForm{
		AdministrableDoseFormOralSuspension,
		AdministrableDoseFormOralGel,
		AdministrableDoseFormPowderForOralSolution,
		AdministrableDoseFormGranulesForOralSolution,
		AdministrableDoseFormLyophilisateForSuspension,
		AdministrableDoseFormPowderForSyrup,
		AdministrableDoseFormSolubleTablet,
		AdministrableDoseFormHerbalTea,
		AdministrableDoseFormInstantHerbalTea,
		AdministrableDoseFormGranules,
		AdministrableDoseFormGastroResistantGranules,
		AdministrableDoseFormModifiedReleaseGranules,
		AdministrableDoseFormCapsuleHard,
		AdministrableDoseFormGastroResistantCapsuleHard,
		AdministrableDoseFormChewableCapsuleSoft,
		AdministrableDoseFormProlongedReleaseCapsuleSoft,
		AdministrableDoseFormModifiedReleaseCapsuleSoft,
		AdministrableDoseFormCoatedTablet,
		AdministrableDoseFormOralDropsSolution,
		AdministrableDoseFormOralDropsSuspension,
		AdministrableDoseFormOralDropsEmulsion,
		AdministrableDoseFormOralLiquid,
		AdministrableDoseFormOralSolution,
		AdministrableDoseFormOralEmulsion,
		AdministrableDoseFormOralPaste,
		AdministrableDoseFormPowderForOralSuspension,
		AdministrableDoseFormGranulesForOralSuspension,
		AdministrableDoseFormSyrup,
		AdministrableDoseFormGranulesForSyrup,
		AdministrableDoseFormDispersibleTablet,
		AdministrableDoseFormOralPowder,
		AdministrableDoseFormEffervescentPowder,
		AdministrableDoseFormEffervescentGranules,
		AdministrableDoseFormProlongedReleaseGranules,
		AdministrableDoseFormCachet,
		AdministrableDoseFormCapsuleSoft,
		AdministrableDoseFormGastroResistantCapsuleSoft,
		AdministrableDoseFormProlongedReleaseCapsuleHard,
		AdministrableDoseFormModifiedReleaseCapsuleHard,
		AdministrableDoseFormTablet,
		AdministrableDoseFormFilmCoatedTablet,
		AdministrableDoseFormOrodispersibleTablet,
		AdministrableDoseFormGastroResistantTablet,
		AdministrableDoseFormModifiedReleaseTablet,
		AdministrableDoseFormMedicatedChewingGum,
		AdministrableDoseFormPillules,
		AdministrableDoseFormPulsatileReleaseIntraruminalDevice,
		AdministrableDoseFormPremixForMedicatedFeedingStuff,
		AdministrableDoseFormGargle,
		AdministrableDoseFormGarglePowderForSolution,
		AdministrableDoseFormOromucosalSuspension,
		AdministrableDoseFormOromucosalSpray,
		AdministrableDoseFormMouthwash,
		AdministrableDoseFormGingivalSolution,
		AdministrableDoseFormOromucosalPaste,
		AdministrableDoseFormGingivalGel,
		AdministrableDoseFormEffervescentTablet,
		AdministrableDoseFormOralLyophilisate,
		AdministrableDoseFormProlongedReleaseTablet,
		AdministrableDoseFormChewableTablet,
		AdministrableDoseFormOralGum,
		AdministrableDoseFormContinuousReleaseIntraruminalDevice,
		AdministrableDoseFormLickBlock,
		AdministrableDoseFormMedicatedPellets,
		AdministrableDoseFormConcentrateForGargle,
		AdministrableDoseFormGargleTabletForSolution,
		AdministrableDoseFormOromucosalSolution,
		AdministrableDoseFormOromucosalDrops,
		AdministrableDoseFormSublingualSpray,
		AdministrableDoseFormMouthwashTabletForSolution,
		AdministrableDoseFormOromucosalGel,
		AdministrableDoseFormOromucosalCream,
		AdministrableDoseFormGingivalPaste,
		AdministrableDoseFormSublingualTablet,
		AdministrableDoseFormBuccalTablet,
		AdministrableDoseFormCompressedLozenge,
		AdministrableDoseFormOromucosalCapsule,
		AdministrableDoseFormMucoAdhesiveBuccalTablet,
		AdministrableDoseFormLozenge,
		AdministrableDoseFormPastille,
		AdministrableDoseFormDentalGel,
		AdministrableDoseFormDentalInsert,
		AdministrableDoseFormDentalPowder,
		AdministrableDoseFormDentalSuspension,
		AdministrableDoseFormToothpaste,
		AdministrableDoseFormPeriodontalGel,
		AdministrableDoseFormBathAdditive,
		AdministrableDoseFormCream,
		AdministrableDoseFormOintment,
		AdministrableDoseFormMedicatedPlaster,
		AdministrableDoseFormShampoo,
		AdministrableDoseFormCutaneousSpraySuspension,
		AdministrableDoseFormCutaneousLiquid,
		AdministrableDoseFormConcentrateForCutaneousSolution,
		AdministrableDoseFormCutaneousEmulsion,
		AdministrableDoseFormCutaneousPatch,
		AdministrableDoseFormPeriodontalPowder,
		AdministrableDoseFormDentalStick,
		AdministrableDoseFormDentalSolution,
		AdministrableDoseFormDentalEmulsion,
		AdministrableDoseFormPeriodontalInsert,
		AdministrableDoseFormGel,
		AdministrableDoseFormCutaneousPaste,
		AdministrableDoseFormCutaneousFoam,
		AdministrableDoseFormCutaneousSpraySolution,
		AdministrableDoseFormCutaneousSprayPowder,
		AdministrableDoseFormCutaneousSolution,
		AdministrableDoseFormCutaneousSuspension,
		AdministrableDoseFormCutaneousPowder,
		AdministrableDoseFormSolutionForIontophoresis,
		AdministrableDoseFormCollodion,
		AdministrableDoseFormPoultice,
		AdministrableDoseFormCutaneousSponge,
		AdministrableDoseFormCollar,
		AdministrableDoseFormEarTag,
		AdministrableDoseFormDipSuspension,
		AdministrableDoseFormTransdermalPatch,
		AdministrableDoseFormMedicatedNailLacquer,
		AdministrableDoseFormCutaneousStick,
		AdministrableDoseFormImpregnatedDressing,
		AdministrableDoseFormMedicatedPendant,
		AdministrableDoseFormDipSolution,
		AdministrableDoseFormDipEmulsion,
		AdministrableDoseFormConcentrateForDipSuspension,
		AdministrableDoseFormPowderForDipSolution,
		AdministrableDoseFormPowderForSuspensionForFishTreatment,
		AdministrableDoseFormPourOnSuspension,
		AdministrableDoseFormSpotOnSolution,
		AdministrableDoseFormSpotOnEmulsion,
		AdministrableDoseFormTeatDipSuspension,
		AdministrableDoseFormTeatSpraySolution,
		AdministrableDoseFormSolutionForSkinPrickTest,
		AdministrableDoseFormPlasterForProvocationTest,
		AdministrableDoseFormEyeGel,
		AdministrableDoseFormEyeDropsSolution,
		AdministrableDoseFormEyeDropsSuspension,
		AdministrableDoseFormConcentrateForDipSolution,
		AdministrableDoseFormConcentrateForDipEmulsion,
		AdministrableDoseFormConcentrateForSolutionForFishTreatment,
		AdministrableDoseFormPourOnSolution,
		AdministrableDoseFormPourOnEmulsion,
		AdministrableDoseFormSpotOnSuspension,
		AdministrableDoseFormTeatDipSolution,
		AdministrableDoseFormTeatDipEmulsion,
		AdministrableDoseFormTransdermalSystem,
		AdministrableDoseFormSolutionForSkinScratchTest,
		AdministrableDoseFormEyeCream,
		AdministrableDoseFormEyeOintment,
		AdministrableDoseFormEyeDropsEmulsion,
		AdministrableDoseFormEyeDropsSolventForReconstitution,
		AdministrableDoseFormEyeLotion,
		AdministrableDoseFormOphthalmicInsert,
		AdministrableDoseFormEarCream,
		AdministrableDoseFormEarOintment,
		AdministrableDoseFormEarDropsSuspension,
		AdministrableDoseFormEyeDropsProlongedRelease,
		AdministrableDoseFormEyeLotionSolventForReconstitution,
		AdministrableDoseFormOphthalmicStrip,
		AdministrableDoseFormEarGel,
		AdministrableDoseFormEarDropsSolution,
		AdministrableDoseFormEarDropsEmulsion,
		AdministrableDoseFormEarPowder,
		AdministrableDoseFormEarSpraySuspension,
		AdministrableDoseFormEarWashSolution,
		AdministrableDoseFormEarTampon,
		AdministrableDoseFormNasalCream,
		AdministrableDoseFormNasalGel,
		AdministrableDoseFormNasalDropsSolution,
		AdministrableDoseFormNasalDropsEmulsion,
		AdministrableDoseFormNasalSpraySolution,
		AdministrableDoseFormNasalSprayEmulsion,
		AdministrableDoseFormNasalStick,
		AdministrableDoseFormVaginalGel,
		AdministrableDoseFormVaginalFoam,
		AdministrableDoseFormEarSpraySolution,
		AdministrableDoseFormEarSprayEmulsion,
		AdministrableDoseFormEarWashEmulsion,
		AdministrableDoseFormEarStick,
		AdministrableDoseFormNasalOintment,
		AdministrableDoseFormNasalDropsSuspension,
		AdministrableDoseFormNasalPowder,
		AdministrableDoseFormNasalSpraySuspension,
		AdministrableDoseFormNasalWash,
		AdministrableDoseFormVaginalCream,
		AdministrableDoseFormVaginalOintment,
		AdministrableDoseFormVaginalSolution,
		AdministrableDoseFormVaginalEmulsion,
		AdministrableDoseFormPessary,
		AdministrableDoseFormVaginalCapsuleSoft,
		AdministrableDoseFormEffervescentVaginalTablet,
		AdministrableDoseFormVaginalDeliverySystem,
		AdministrableDoseFormRectalCream,
		AdministrableDoseFormRectalFoam,
		AdministrableDoseFormVaginalSuspension,
		AdministrableDoseFormTabletForVaginalSolution,
		AdministrableDoseFormVaginalCapsuleHard,
		AdministrableDoseFormVaginalTablet,
		AdministrableDoseFormMedicatedVaginalTampon,
		AdministrableDoseFormVaginalSponge,
		AdministrableDoseFormRectalGel,
		AdministrableDoseFormSolutionForInjection,
	}
}

func FindAdministrableDoseForm(filter string) []AdministrableDoseForm {
	ret := make([]AdministrableDoseForm, 0)
	for _, at := range AllAdministrableDoseForm() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (cpt AdministrableDoseForm) ToString() {
	fmt.Println(cpt.String())
}

func (cpt AdministrableDoseForm) String() string {
	switch cpt {
	case AdministrableDoseFormOralSuspension:
		return "Oral suspension"
	case AdministrableDoseFormOralGel:
		return "Oral gel"
	case AdministrableDoseFormPowderForOralSolution:
		return "Powder for oral solution"
	case AdministrableDoseFormGranulesForOralSolution:
		return "Granules for oral solution"
	case AdministrableDoseFormLyophilisateForSuspension:
		return "Lyophilisate for suspension"
	case AdministrableDoseFormPowderForSyrup:
		return "Powder for syrup"
	case AdministrableDoseFormSolubleTablet:
		return "Soluble tablet"
	case AdministrableDoseFormHerbalTea:
		return "Herbal tea"
	case AdministrableDoseFormInstantHerbalTea:
		return "Instant herbal tea"
	case AdministrableDoseFormGranules:
		return "Granules"
	case AdministrableDoseFormGastroResistantGranules:
		return "Gastro-resistant granules"
	case AdministrableDoseFormModifiedReleaseGranules:
		return "Modified-release granules"
	case AdministrableDoseFormCapsuleHard:
		return "Capsule, hard"
	case AdministrableDoseFormGastroResistantCapsuleHard:
		return "Gastro-resistant capsule, hard"
	case AdministrableDoseFormChewableCapsuleSoft:
		return "Chewable capsule, soft"
	case AdministrableDoseFormProlongedReleaseCapsuleSoft:
		return "Prolonged-release capsule, soft"
	case AdministrableDoseFormModifiedReleaseCapsuleSoft:
		return "Modified-release capsule, soft"
	case AdministrableDoseFormCoatedTablet:
		return "Coated tablet"
	case AdministrableDoseFormOralDropsSolution:
		return "Oral drops, solution"
	case AdministrableDoseFormOralDropsSuspension:
		return "Oral drops, suspension"
	case AdministrableDoseFormOralDropsEmulsion:
		return "Oral drops, emulsion"
	case AdministrableDoseFormOralLiquid:
		return "Oral liquid"
	case AdministrableDoseFormOralSolution:
		return "Oral solution"
	case AdministrableDoseFormOralEmulsion:
		return "Oral emulsion"
	case AdministrableDoseFormOralPaste:
		return "Oral paste"
	case AdministrableDoseFormPowderForOralSuspension:
		return "Powder for oral suspension"
	case AdministrableDoseFormGranulesForOralSuspension:
		return "Granules for oral suspension"
	case AdministrableDoseFormSyrup:
		return "Syrup"
	case AdministrableDoseFormGranulesForSyrup:
		return "Granules for syrup"
	case AdministrableDoseFormDispersibleTablet:
		return "Dispersible tablet"
	case AdministrableDoseFormOralPowder:
		return "Oral powder"
	case AdministrableDoseFormEffervescentPowder:
		return "Effervescent powder"
	case AdministrableDoseFormEffervescentGranules:
		return "Effervescent granules"
	case AdministrableDoseFormProlongedReleaseGranules:
		return "Prolonged-release granules"
	case AdministrableDoseFormCachet:
		return "Cachet"
	case AdministrableDoseFormCapsuleSoft:
		return "Capsule, soft"
	case AdministrableDoseFormGastroResistantCapsuleSoft:
		return "Gastro-resistant capsule, soft"
	case AdministrableDoseFormProlongedReleaseCapsuleHard:
		return "Prolonged-release capsule, hard"
	case AdministrableDoseFormModifiedReleaseCapsuleHard:
		return "Modified-release capsule, hard"
	case AdministrableDoseFormTablet:
		return "Tablet"
	case AdministrableDoseFormFilmCoatedTablet:
		return "Film-coated tablet"
	case AdministrableDoseFormOrodispersibleTablet:
		return "Orodispersible tablet"
	case AdministrableDoseFormGastroResistantTablet:
		return "Gastro-resistant tablet"
	case AdministrableDoseFormModifiedReleaseTablet:
		return "Modified-release tablet"
	case AdministrableDoseFormMedicatedChewingGum:
		return "Medicated chewing-gum"
	case AdministrableDoseFormPillules:
		return "Pillules"
	case AdministrableDoseFormPulsatileReleaseIntraruminalDevice:
		return "Pulsatile-release intraruminal device"
	case AdministrableDoseFormPremixForMedicatedFeedingStuff:
		return "Premix for medicated feeding stuff"
	case AdministrableDoseFormGargle:
		return "Gargle"
	case AdministrableDoseFormGarglePowderForSolution:
		return "Gargle, powder for solution"
	case AdministrableDoseFormOromucosalSuspension:
		return "Oromucosal suspension"
	case AdministrableDoseFormOromucosalSpray:
		return "Oromucosal spray"
	case AdministrableDoseFormMouthwash:
		return "Mouthwash"
	case AdministrableDoseFormGingivalSolution:
		return "Gingival solution"
	case AdministrableDoseFormOromucosalPaste:
		return "Oromucosal paste"
	case AdministrableDoseFormGingivalGel:
		return "Gingival gel"
	case AdministrableDoseFormEffervescentTablet:
		return "Effervescent tablet"
	case AdministrableDoseFormOralLyophilisate:
		return "Oral lyophilisate"
	case AdministrableDoseFormProlongedReleaseTablet:
		return "Prolonged-release tablet"
	case AdministrableDoseFormChewableTablet:
		return "Chewable tablet"
	case AdministrableDoseFormOralGum:
		return "Oral gum"
	case AdministrableDoseFormContinuousReleaseIntraruminalDevice:
		return "Continuous-release intraruminal device"
	case AdministrableDoseFormLickBlock:
		return "Lick block"
	case AdministrableDoseFormMedicatedPellets:
		return "Medicated pellets"
	case AdministrableDoseFormConcentrateForGargle:
		return "Concentrate for gargle"
	case AdministrableDoseFormGargleTabletForSolution:
		return "Gargle, tablet for solution"
	case AdministrableDoseFormOromucosalSolution:
		return "Oromucosal solution"
	case AdministrableDoseFormOromucosalDrops:
		return "Oromucosal drops"
	case AdministrableDoseFormSublingualSpray:
		return "Sublingual spray"
	case AdministrableDoseFormMouthwashTabletForSolution:
		return "Mouthwash, tablet for solution"
	case AdministrableDoseFormOromucosalGel:
		return "Oromucosal gel"
	case AdministrableDoseFormOromucosalCream:
		return "Oromucosal cream"
	case AdministrableDoseFormGingivalPaste:
		return "Gingival paste"
	case AdministrableDoseFormSublingualTablet:
		return "Sublingual tablet"
	case AdministrableDoseFormBuccalTablet:
		return "Buccal tablet"
	case AdministrableDoseFormCompressedLozenge:
		return "Compressed lozenge"
	case AdministrableDoseFormOromucosalCapsule:
		return "Oromucosal capsule"
	case AdministrableDoseFormMucoAdhesiveBuccalTablet:
		return "Muco-adhesive buccal tablet"
	case AdministrableDoseFormLozenge:
		return "Lozenge"
	case AdministrableDoseFormPastille:
		return "Pastille"
	case AdministrableDoseFormDentalGel:
		return "Dental gel"
	case AdministrableDoseFormDentalInsert:
		return "Dental insert"
	case AdministrableDoseFormDentalPowder:
		return "Dental powder"
	case AdministrableDoseFormDentalSuspension:
		return "Dental suspension"
	case AdministrableDoseFormToothpaste:
		return "Toothpaste"
	case AdministrableDoseFormPeriodontalGel:
		return "Periodontal gel"
	case AdministrableDoseFormBathAdditive:
		return "Bath additive"
	case AdministrableDoseFormCream:
		return "Cream"
	case AdministrableDoseFormOintment:
		return "Ointment"
	case AdministrableDoseFormMedicatedPlaster:
		return "Medicated plaster"
	case AdministrableDoseFormShampoo:
		return "Shampoo"
	case AdministrableDoseFormCutaneousSpraySuspension:
		return "Cutaneous spray, suspension"
	case AdministrableDoseFormCutaneousLiquid:
		return "Cutaneous liquid"
	case AdministrableDoseFormConcentrateForCutaneousSolution:
		return "Concentrate for cutaneous solution"
	case AdministrableDoseFormCutaneousEmulsion:
		return "Cutaneous emulsion"
	case AdministrableDoseFormCutaneousPatch:
		return "Cutaneous patch"
	case AdministrableDoseFormPeriodontalPowder:
		return "Periodontal powder"
	case AdministrableDoseFormDentalStick:
		return "Dental stick"
	case AdministrableDoseFormDentalSolution:
		return "Dental solution"
	case AdministrableDoseFormDentalEmulsion:
		return "Dental emulsion"
	case AdministrableDoseFormPeriodontalInsert:
		return "Periodontal insert"
	case AdministrableDoseFormGel:
		return "Gel"
	case AdministrableDoseFormCutaneousPaste:
		return "Cutaneous paste"
	case AdministrableDoseFormCutaneousFoam:
		return "Cutaneous foam"
	case AdministrableDoseFormCutaneousSpraySolution:
		return "Cutaneous spray, solution"
	case AdministrableDoseFormCutaneousSprayPowder:
		return "Cutaneous spray, powder"
	case AdministrableDoseFormCutaneousSolution:
		return "Cutaneous solution"
	case AdministrableDoseFormCutaneousSuspension:
		return "Cutaneous suspension"
	case AdministrableDoseFormCutaneousPowder:
		return "Cutaneous powder"
	case AdministrableDoseFormSolutionForIontophoresis:
		return "Solution for iontophoresis"
	case AdministrableDoseFormCollodion:
		return "Collodion"
	case AdministrableDoseFormPoultice:
		return "Poultice"
	case AdministrableDoseFormCutaneousSponge:
		return "Cutaneous sponge"
	case AdministrableDoseFormCollar:
		return "Collar"
	case AdministrableDoseFormEarTag:
		return "Ear tag"
	case AdministrableDoseFormDipSuspension:
		return "Dip suspension"
	case AdministrableDoseFormTransdermalPatch:
		return "Transdermal patch"
	case AdministrableDoseFormMedicatedNailLacquer:
		return "Medicated nail lacquer"
	case AdministrableDoseFormCutaneousStick:
		return "Cutaneous stick"
	case AdministrableDoseFormImpregnatedDressing:
		return "Impregnated dressing"
	case AdministrableDoseFormMedicatedPendant:
		return "Medicated pendant"
	case AdministrableDoseFormDipSolution:
		return "Dip solution"
	case AdministrableDoseFormDipEmulsion:
		return "Dip emulsion"
	case AdministrableDoseFormConcentrateForDipSuspension:
		return "Concentrate for dip suspension"
	case AdministrableDoseFormPowderForDipSolution:
		return "Powder for dip solution"
	case AdministrableDoseFormPowderForSuspensionForFishTreatment:
		return "Powder for suspension for fish treatment"
	case AdministrableDoseFormPourOnSuspension:
		return "Pour-on suspension"
	case AdministrableDoseFormSpotOnSolution:
		return "Spot-on solution"
	case AdministrableDoseFormSpotOnEmulsion:
		return "Spot-on emulsion"
	case AdministrableDoseFormTeatDipSuspension:
		return "Teat dip suspension"
	case AdministrableDoseFormTeatSpraySolution:
		return "Teat spray solution"
	case AdministrableDoseFormSolutionForSkinPrickTest:
		return "Solution for skin-prick test"
	case AdministrableDoseFormPlasterForProvocationTest:
		return "Plaster for provocation test"
	case AdministrableDoseFormEyeGel:
		return "Eye gel"
	case AdministrableDoseFormEyeDropsSolution:
		return "Eye drops, solution"
	case AdministrableDoseFormEyeDropsSuspension:
		return "Eye drops, suspension"
	case AdministrableDoseFormConcentrateForDipSolution:
		return "Concentrate for dip solution"
	case AdministrableDoseFormConcentrateForDipEmulsion:
		return "Concentrate for dip emulsion"
	case AdministrableDoseFormConcentrateForSolutionForFishTreatment:
		return "Concentrate for solution for fish treatment"
	case AdministrableDoseFormPourOnSolution:
		return "Pour-on solution"
	case AdministrableDoseFormPourOnEmulsion:
		return "Pour-on emulsion"
	case AdministrableDoseFormSpotOnSuspension:
		return "Spot-on suspension"
	case AdministrableDoseFormTeatDipSolution:
		return "Teat dip solution"
	case AdministrableDoseFormTeatDipEmulsion:
		return "Teat dip emulsion"
	case AdministrableDoseFormTransdermalSystem:
		return "Transdermal system"
	case AdministrableDoseFormSolutionForSkinScratchTest:
		return "Solution for skin-scratch test"
	case AdministrableDoseFormEyeCream:
		return "Eye cream"
	case AdministrableDoseFormEyeOintment:
		return "Eye ointment"
	case AdministrableDoseFormEyeDropsEmulsion:
		return "Eye drops, emulsion"
	case AdministrableDoseFormEyeDropsSolventForReconstitution:
		return "Eye drops, solvent for reconstitution"
	case AdministrableDoseFormEyeLotion:
		return "Eye lotion"
	case AdministrableDoseFormOphthalmicInsert:
		return "Ophthalmic insert"
	case AdministrableDoseFormEarCream:
		return "Ear cream"
	case AdministrableDoseFormEarOintment:
		return "Ear ointment"
	case AdministrableDoseFormEarDropsSuspension:
		return "Ear drops, suspension"
	case AdministrableDoseFormEyeDropsProlongedRelease:
		return "Eye drops, prolonged-release"
	case AdministrableDoseFormEyeLotionSolventForReconstitution:
		return "Eye lotion, solvent for reconstitution"
	case AdministrableDoseFormOphthalmicStrip:
		return "Ophthalmic strip"
	case AdministrableDoseFormEarGel:
		return "Ear gel"
	case AdministrableDoseFormEarDropsSolution:
		return "Ear drops, solution"
	case AdministrableDoseFormEarDropsEmulsion:
		return "Ear drops, emulsion"
	case AdministrableDoseFormEarPowder:
		return "Ear powder"
	case AdministrableDoseFormEarSpraySuspension:
		return "Ear spray, suspension"
	case AdministrableDoseFormEarWashSolution:
		return "Ear wash, solution"
	case AdministrableDoseFormEarTampon:
		return "Ear tampon"
	case AdministrableDoseFormNasalCream:
		return "Nasal cream"
	case AdministrableDoseFormNasalGel:
		return "Nasal gel"
	case AdministrableDoseFormNasalDropsSolution:
		return "Nasal drops, solution"
	case AdministrableDoseFormNasalDropsEmulsion:
		return "Nasal drops, emulsion"
	case AdministrableDoseFormNasalSpraySolution:
		return "Nasal spray, solution"
	case AdministrableDoseFormNasalSprayEmulsion:
		return "Nasal spray, emulsion"
	case AdministrableDoseFormNasalStick:
		return "Nasal stick"
	case AdministrableDoseFormVaginalGel:
		return "Vaginal gel"
	case AdministrableDoseFormVaginalFoam:
		return "Vaginal foam"
	case AdministrableDoseFormEarSpraySolution:
		return "Ear spray, solution"
	case AdministrableDoseFormEarSprayEmulsion:
		return "Ear spray, emulsion"
	case AdministrableDoseFormEarWashEmulsion:
		return "Ear wash, emulsion"
	case AdministrableDoseFormEarStick:
		return "Ear stick"
	case AdministrableDoseFormNasalOintment:
		return "Nasal ointment"
	case AdministrableDoseFormNasalDropsSuspension:
		return "Nasal drops, suspension"
	case AdministrableDoseFormNasalPowder:
		return "Nasal powder"
	case AdministrableDoseFormNasalSpraySuspension:
		return "Nasal spray, suspension"
	case AdministrableDoseFormNasalWash:
		return "Nasal wash"
	case AdministrableDoseFormVaginalCream:
		return "Vaginal cream"
	case AdministrableDoseFormVaginalOintment:
		return "Vaginal ointment"
	case AdministrableDoseFormVaginalSolution:
		return "Vaginal solution"
	case AdministrableDoseFormVaginalEmulsion:
		return "Vaginal emulsion"
	case AdministrableDoseFormPessary:
		return "Pessary"
	case AdministrableDoseFormVaginalCapsuleSoft:
		return "Vaginal capsule, soft"
	case AdministrableDoseFormEffervescentVaginalTablet:
		return "Effervescent vaginal tablet"
	case AdministrableDoseFormVaginalDeliverySystem:
		return "Vaginal delivery system"
	case AdministrableDoseFormRectalCream:
		return "Rectal cream"
	case AdministrableDoseFormRectalFoam:
		return "Rectal foam"
	case AdministrableDoseFormVaginalSuspension:
		return "Vaginal suspension"
	case AdministrableDoseFormTabletForVaginalSolution:
		return "Tablet for vaginal solution"
	case AdministrableDoseFormVaginalCapsuleHard:
		return "Vaginal capsule, hard"
	case AdministrableDoseFormVaginalTablet:
		return "Vaginal tablet"
	case AdministrableDoseFormMedicatedVaginalTampon:
		return "Medicated vaginal tampon"
	case AdministrableDoseFormVaginalSponge:
		return "Vaginal sponge"
	case AdministrableDoseFormRectalGel:
		return "Rectal gel"
	case AdministrableDoseFormSolutionForInjection:
		return "Solution for injection"
	default:
		return "Unknown Administration Subpotent Reason"
	}
}

/*
Oral suspension
Oral gel
Powder for oral solution
Granules for oral solution
Lyophilisate for suspension
Powder for syrup
Soluble tablet
Herbal tea
Instant herbal tea
Granules
Gastro-resistant granules
Modified-release granules
Capsule, hard
Gastro-resistant capsule, hard
Chewable capsule, soft
Prolonged-release capsule, soft
Modified-release capsule, soft
Coated tablet
Oral drops, solution
Oral drops, suspension
Oral drops, emulsion
Oral liquid
Oral solution
Oral emulsion
Oral paste
Powder for oral suspension
Granules for oral suspension
Syrup
Granules for syrup
Dispersible tablet
Oral powder
Effervescent powder
Effervescent granules
Prolonged-release granules
Cachet
Capsule, soft
Gastro-resistant capsule, soft
Prolonged-release capsule, hard
Modified-release capsule, hard
Tablet
Film-coated tablet
Orodispersible tablet
Gastro-resistant tablet
Modified-release tablet
Medicated chewing-gum
Pillules
Pulsatile-release intraruminal device
Premix for medicated feeding stuff
Gargle
Gargle, powder for solution
Oromucosal suspension
Oromucosal spray
Mouthwash
Gingival solution
Oromucosal paste
Gingival gel
Effervescent tablet
Oral lyophilisate
Prolonged-release tablet
Chewable tablet
Oral gum
Continuous-release intraruminal device
Lick block
Medicated pellets
Concentrate for gargle
Gargle, tablet for solution
Oromucosal solution
Oromucosal drops
Sublingual spray
Mouthwash, tablet for solution
Oromucosal gel
Oromucosal cream
Gingival paste
Sublingual tablet
Buccal tablet
Compressed lozenge
Oromucosal capsule
Muco-adhesive buccal tablet
Lozenge
Pastille
Dental gel
Dental insert
Dental powder
Dental suspension
Toothpaste
Periodontal gel
Bath additive
Cream
Ointment
Medicated plaster
Shampoo
Cutaneous spray, suspension
Cutaneous liquid
Concentrate for cutaneous solution
Cutaneous emulsion
Cutaneous patch
Periodontal powder
Dental stick
Dental solution
Dental emulsion
Periodontal insert
Gel
Cutaneous paste
Cutaneous foam
Cutaneous spray, solution
Cutaneous spray, powder
Cutaneous solution
Cutaneous suspension
Cutaneous powder
Solution for iontophoresis
Collodion
Poultice
Cutaneous sponge
Collar
Ear tag
Dip suspension
Transdermal patch
Medicated nail lacquer
Cutaneous stick
Impregnated dressing
Medicated pendant
Dip solution
Dip emulsion
Concentrate for dip suspension
Powder for dip solution
Powder for suspension for fish treatment
Pour-on suspension
Spot-on solution
Spot-on emulsion
Teat dip suspension
Teat spray solution
Solution for skin-prick test
Plaster for provocation test
Eye gel
Eye drops, solution
Eye drops, suspension
Concentrate for dip solution
Concentrate for dip emulsion
Concentrate for solution for fish treatment
Pour-on solution
Pour-on emulsion
Spot-on suspension
Teat dip solution
Teat dip emulsion
Transdermal system
Solution for skin-scratch test
Eye cream
Eye ointment
Eye drops, emulsion
Eye drops, solvent for reconstitution
Eye lotion
Ophthalmic insert
Ear cream
Ear ointment
Ear drops, suspension
Eye drops, prolonged-release
Eye lotion, solvent for reconstitution
Ophthalmic strip
Ear gel
Ear drops, solution
Ear drops, emulsion
Ear powder
Ear spray, suspension
Ear wash, solution
Ear tampon
Nasal cream
Nasal gel
Nasal drops, solution
Nasal drops, emulsion
Nasal spray, solution
Nasal spray, emulsion
Nasal stick
Vaginal gel
Vaginal foam
Ear spray, solution
Ear spray, emulsion
Ear wash, emulsion
Ear stick
Nasal ointment
Nasal drops, suspension
Nasal powder
Nasal spray, suspension
Nasal wash
Vaginal cream
Vaginal ointment
Vaginal solution
Vaginal emulsion
Pessary
Vaginal capsule, soft
Effervescent vaginal tablet
Vaginal delivery system
Rectal cream
Rectal foam
Vaginal suspension
Tablet for vaginal solution
Vaginal capsule, hard
Vaginal tablet
Medicated vaginal tampon
Vaginal sponge
Rectal gel
Solution for injection

*/
