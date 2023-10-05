package code

import (
	"fmt"
	"strings"
)

type AdministrableDoseForm string

const (
	AdministrableDoseFormOralSuspension                         AdministrableDoseForm = "100000073362"
	AdministrableDoseFormOralGel                                AdministrableDoseForm = "100000073363"
	AdministrableDoseFormPowderForOralSolution                  AdministrableDoseForm = "100000073364"
	AdministrableDoseFormGranulesForOralSolution                AdministrableDoseForm = "100000073365"
	AdministrableDoseFormLyophilisateForSuspension              AdministrableDoseForm = "100000073367"
	AdministrableDoseFormPowderForSyrup                         AdministrableDoseForm = "100000073368"
	AdministrableDoseFormSolubleTablet                          AdministrableDoseForm = "100000073369"
	AdministrableDoseFormHerbalTea                              AdministrableDoseForm = "100000073370"
	AdministrableDoseFormInstantHerbalTea                       AdministrableDoseForm = "100000073371"
	AdministrableDoseFormGranules                               AdministrableDoseForm = "100000073372"
	AdministrableDoseFormGastroResistantGranules                AdministrableDoseForm = "100000073373"
	AdministrableDoseFormModifiedReleaseGranules                AdministrableDoseForm = "100000073374"
	AdministrableDoseFormCapsuleHard                            AdministrableDoseForm = "100000073375"
	AdministrableDoseFormGastroResistantCapsuleHard             AdministrableDoseForm = "100000073376"
	AdministrableDoseFormChewableCapsuleSoft                    AdministrableDoseForm = "100000073377"
	AdministrableDoseFormProlongedReleaseCapsuleSoft            AdministrableDoseForm = "100000073378"
	AdministrableDoseFormModifiedReleaseCapsuleSoft             AdministrableDoseForm = "100000073379"
	AdministrableDoseFormCoatedTablet                           AdministrableDoseForm = "100000073380"
	AdministrableDoseFormOralDropsSolution                      AdministrableDoseForm = "100000073642"
	AdministrableDoseFormOralDropsSuspension                    AdministrableDoseForm = "100000073643"
	AdministrableDoseFormOralDropsEmulsion                      AdministrableDoseForm = "100000073644"
	AdministrableDoseFormOralLiquid                             AdministrableDoseForm = "100000073645"
	AdministrableDoseFormOralSolution                           AdministrableDoseForm = "100000073646"
	AdministrableDoseFormOralEmulsion                           AdministrableDoseForm = "100000073647"
	AdministrableDoseFormOralPaste                              AdministrableDoseForm = "100000073648"
	AdministrableDoseFormPowderForOralSuspension                AdministrableDoseForm = "100000073649"
	AdministrableDoseFormGranulesForOralSuspension              AdministrableDoseForm = "100000073650"
	AdministrableDoseFormSyrup                                  AdministrableDoseForm = "100000073652"
	AdministrableDoseFormGranulesForSyrup                       AdministrableDoseForm = "100000073653"
	AdministrableDoseFormDispersibleTablet                      AdministrableDoseForm = "100000073654"
	AdministrableDoseFormOralPowder                             AdministrableDoseForm = "100000073655"
	AdministrableDoseFormEffervescentPowder                     AdministrableDoseForm = "100000073656"
	AdministrableDoseFormEffervescentGranules                   AdministrableDoseForm = "100000073657"
	AdministrableDoseFormProlongedReleaseGranules               AdministrableDoseForm = "100000073658"
	AdministrableDoseFormCachet                                 AdministrableDoseForm = "100000073659"
	AdministrableDoseFormCapsuleSoft                            AdministrableDoseForm = "100000073660"
	AdministrableDoseFormGastroResistantCapsuleSoft             AdministrableDoseForm = "100000073661"
	AdministrableDoseFormProlongedReleaseCapsuleHard            AdministrableDoseForm = "100000073662"
	AdministrableDoseFormModifiedReleaseCapsuleHard             AdministrableDoseForm = "100000073663"
	AdministrableDoseFormTablet                                 AdministrableDoseForm = "100000073664"
	AdministrableDoseFormFilmCoatedTablet                       AdministrableDoseForm = "100000073665"
	AdministrableDoseFormOrodispersibleTablet                   AdministrableDoseForm = "100000073666"
	AdministrableDoseFormGastroResistantTablet                  AdministrableDoseForm = "100000073667"
	AdministrableDoseFormModifiedReleaseTablet                  AdministrableDoseForm = "100000073668"
	AdministrableDoseFormMedicatedChewingGum                    AdministrableDoseForm = "100000073669"
	AdministrableDoseFormPillules                               AdministrableDoseForm = "100000073670"
	AdministrableDoseFormPulsatileReleaseIntraruminalDevice     AdministrableDoseForm = "100000073671"
	AdministrableDoseFormPremixForMedicatedFeedingStuff         AdministrableDoseForm = "100000073672"
	AdministrableDoseFormGargle                                 AdministrableDoseForm = "100000073673"
	AdministrableDoseFormGarglePowderForSolution                AdministrableDoseForm = "100000073674"
	AdministrableDoseFormOromucosalSuspension                   AdministrableDoseForm = "100000073675"
	AdministrableDoseFormOromucosalSpray                        AdministrableDoseForm = "100000073676"
	AdministrableDoseFormMouthwash                              AdministrableDoseForm = "100000073677"
	AdministrableDoseFormGingivalSolution                       AdministrableDoseForm = "100000073678"
	AdministrableDoseFormOromucosalPaste                        AdministrableDoseForm = "100000073679"
	AdministrableDoseFormGingivalGel                            AdministrableDoseForm = "100000073680"
	AdministrableDoseFormEffervescentTablet                     AdministrableDoseForm = "100000073681"
	AdministrableDoseFormOralLyophilisate                       AdministrableDoseForm = "100000073682"
	AdministrableDoseFormProlongedReleaseTablet                 AdministrableDoseForm = "100000073683"
	AdministrableDoseFormChewableTablet                         AdministrableDoseForm = "100000073684"
	AdministrableDoseFormOralGum                                AdministrableDoseForm = "100000073685"
	AdministrableDoseFormContinuousReleaseIntraruminalDevice    AdministrableDoseForm = "100000073686"
	AdministrableDoseFormLickBlock                              AdministrableDoseForm = "100000073687"
	AdministrableDoseFormMedicatedPellets                       AdministrableDoseForm = "100000073688"
	AdministrableDoseFormConcentrateForGargle                   AdministrableDoseForm = "100000073689"
	AdministrableDoseFormGargleTabletForSolution                AdministrableDoseForm = "100000073690"
	AdministrableDoseFormOromucosalSolution                     AdministrableDoseForm = "100000073691"
	AdministrableDoseFormOromucosalDrops                        AdministrableDoseForm = "100000073692"
	AdministrableDoseFormSublingualSpray                        AdministrableDoseForm = "100000073693"
	AdministrableDoseFormMouthwashTabletForSolution             AdministrableDoseForm = "100000073694"
	AdministrableDoseFormOromucosalGel                          AdministrableDoseForm = "100000073695"
	AdministrableDoseFormOromucosalCream                        AdministrableDoseForm = "100000073696"
	AdministrableDoseFormGingivalPaste                          AdministrableDoseForm = "100000073697"
	AdministrableDoseFormSublingualTablet                       AdministrableDoseForm = "100000073698"
	AdministrableDoseFormBuccalTablet                           AdministrableDoseForm = "100000073699"
	AdministrableDoseFormCompressedLozenge                      AdministrableDoseForm = "100000073700"
	AdministrableDoseFormOromucosalCapsule                      AdministrableDoseForm = "100000073701"
	AdministrableDoseFormMucoAdhesiveBuccalTablet               AdministrableDoseForm = "100000073702"
	AdministrableDoseFormLozenge                                AdministrableDoseForm = "100000073703"
	AdministrableDoseFormPastille                               AdministrableDoseForm = "100000073704"
	AdministrableDoseFormDentalGel                              AdministrableDoseForm = "100000073705"
	AdministrableDoseFormDentalInsert                           AdministrableDoseForm = "100000073706"
	AdministrableDoseFormDentalPowder                           AdministrableDoseForm = "100000073707"
	AdministrableDoseFormDentalSuspension                       AdministrableDoseForm = "100000073708"
	AdministrableDoseFormToothpaste                             AdministrableDoseForm = "100000073709"
	AdministrableDoseFormPeriodontalGel                         AdministrableDoseForm = "100000073710"
	AdministrableDoseFormBathAdditive                           AdministrableDoseForm = "100000073711"
	AdministrableDoseFormCream                                  AdministrableDoseForm = "100000073712"
	AdministrableDoseFormOintment                               AdministrableDoseForm = "100000073713"
	AdministrableDoseFormMedicatedPlaster                       AdministrableDoseForm = "100000073714"
	AdministrableDoseFormShampoo                                AdministrableDoseForm = "100000073715"
	AdministrableDoseFormCutaneousSpraySuspension               AdministrableDoseForm = "100000073716"
	AdministrableDoseFormCutaneousLiquid                        AdministrableDoseForm = "100000073717"
	AdministrableDoseFormConcentrateForCutaneousSolution        AdministrableDoseForm = "100000073718"
	AdministrableDoseFormCutaneousEmulsion                      AdministrableDoseForm = "100000073719"
	AdministrableDoseFormCutaneousPatch                         AdministrableDoseForm = "100000073720"
	AdministrableDoseFormPeriodontalPowder                      AdministrableDoseForm = "100000073721"
	AdministrableDoseFormDentalStick                            AdministrableDoseForm = "100000073722"
	AdministrableDoseFormDentalSolution                         AdministrableDoseForm = "100000073723"
	AdministrableDoseFormDentalEmulsion                         AdministrableDoseForm = "100000073724"
	AdministrableDoseFormPeriodontalInsert                      AdministrableDoseForm = "100000073725"
	AdministrableDoseFormGel                                    AdministrableDoseForm = "100000073726"
	AdministrableDoseFormCutaneousPaste                         AdministrableDoseForm = "100000073727"
	AdministrableDoseFormCutaneousFoam                          AdministrableDoseForm = "100000073728"
	AdministrableDoseFormCutaneousSpraySolution                 AdministrableDoseForm = "100000073729"
	AdministrableDoseFormCutaneousSprayPowder                   AdministrableDoseForm = "100000073730"
	AdministrableDoseFormCutaneousSolution                      AdministrableDoseForm = "100000073731"
	AdministrableDoseFormCutaneousSuspension                    AdministrableDoseForm = "100000073732"
	AdministrableDoseFormCutaneousPowder                        AdministrableDoseForm = "100000073733"
	AdministrableDoseFormSolutionForIontophoresis               AdministrableDoseForm = "100000073734"
	AdministrableDoseFormCollodion                              AdministrableDoseForm = "100000073735"
	AdministrableDoseFormPoultice                               AdministrableDoseForm = "100000073736"
	AdministrableDoseFormCutaneousSponge                        AdministrableDoseForm = "100000073737"
	AdministrableDoseFormCollar                                 AdministrableDoseForm = "100000073738"
	AdministrableDoseFormEarTag                                 AdministrableDoseForm = "100000073739"
	AdministrableDoseFormDipSuspension                          AdministrableDoseForm = "100000073740"
	AdministrableDoseFormTransdermalPatch                       AdministrableDoseForm = "100000073741"
	AdministrableDoseFormMedicatedNailLacquer                   AdministrableDoseForm = "100000073742"
	AdministrableDoseFormCutaneousStick                         AdministrableDoseForm = "100000073743"
	AdministrableDoseFormImpregnatedDressing                    AdministrableDoseForm = "100000073744"
	AdministrableDoseFormMedicatedPendant                       AdministrableDoseForm = "100000073745"
	AdministrableDoseFormDipSolution                            AdministrableDoseForm = "100000073746"
	AdministrableDoseFormDipEmulsion                            AdministrableDoseForm = "100000073747"
	AdministrableDoseFormConcentrateForDipSuspension            AdministrableDoseForm = "100000073748"
	AdministrableDoseFormPowderForDipSolution                   AdministrableDoseForm = "100000073749"
	AdministrableDoseFormPowderForSuspensionForFishTreatment    AdministrableDoseForm = "100000073750"
	AdministrableDoseFormPourOnSuspension                       AdministrableDoseForm = "100000073751"
	AdministrableDoseFormSpotOnSolution                         AdministrableDoseForm = "100000073752"
	AdministrableDoseFormSpotOnEmulsion                         AdministrableDoseForm = "100000073753"
	AdministrableDoseFormTeatDipSuspension                      AdministrableDoseForm = "100000073754"
	AdministrableDoseFormTeatSpraySolution                      AdministrableDoseForm = "100000073755"
	AdministrableDoseFormSolutionForSkinPrickTest               AdministrableDoseForm = "100000073756"
	AdministrableDoseFormPlasterForProvocationTest              AdministrableDoseForm = "100000073757"
	AdministrableDoseFormEyeGel                                 AdministrableDoseForm = "100000073758"
	AdministrableDoseFormEyeDropsSolution                       AdministrableDoseForm = "100000073759"
	AdministrableDoseFormEyeDropsSuspension                     AdministrableDoseForm = "100000073760"
	AdministrableDoseFormConcentrateForDipSolution              AdministrableDoseForm = "100000073761"
	AdministrableDoseFormConcentrateForDipEmulsion              AdministrableDoseForm = "100000073762"
	AdministrableDoseFormConcentrateForSolutionForFishTreatment AdministrableDoseForm = "100000073763"
	AdministrableDoseFormPourOnSolution                         AdministrableDoseForm = "100000073764"
	AdministrableDoseFormPourOnEmulsion                         AdministrableDoseForm = "100000073765"
	AdministrableDoseFormSpotOnSuspension                       AdministrableDoseForm = "100000073766"
	AdministrableDoseFormTeatDipSolution                        AdministrableDoseForm = "100000073767"
	AdministrableDoseFormTeatDipEmulsion                        AdministrableDoseForm = "100000073768"
	AdministrableDoseFormTransdermalSystem                      AdministrableDoseForm = "100000073769"
	AdministrableDoseFormSolutionForSkinScratchTest             AdministrableDoseForm = "100000073770"
	AdministrableDoseFormEyeCream                               AdministrableDoseForm = "100000073771"
	AdministrableDoseFormEyeOintment                            AdministrableDoseForm = "100000073772"
	AdministrableDoseFormEyeDropsEmulsion                       AdministrableDoseForm = "100000073773"
	AdministrableDoseFormEyeDropsSolventForReconstitution       AdministrableDoseForm = "100000073775"
	AdministrableDoseFormEyeLotion                              AdministrableDoseForm = "100000073776"
	AdministrableDoseFormOphthalmicInsert                       AdministrableDoseForm = "100000073777"
	AdministrableDoseFormEarCream                               AdministrableDoseForm = "100000073778"
	AdministrableDoseFormEarOintment                            AdministrableDoseForm = "100000073779"
	AdministrableDoseFormEarDropsSuspension                     AdministrableDoseForm = "100000073780"
	AdministrableDoseFormEyeDropsProlongedRelease               AdministrableDoseForm = "100000073782"
	AdministrableDoseFormEyeLotionSolventForReconstitution      AdministrableDoseForm = "100000073783"
	AdministrableDoseFormOphthalmicStrip                        AdministrableDoseForm = "100000073784"
	AdministrableDoseFormEarGel                                 AdministrableDoseForm = "100000073785"
	AdministrableDoseFormEarDropsSolution                       AdministrableDoseForm = "100000073786"
	AdministrableDoseFormEarDropsEmulsion                       AdministrableDoseForm = "100000073787"
	AdministrableDoseFormEarPowder                              AdministrableDoseForm = "100000073788"
	AdministrableDoseFormEarSpraySuspension                     AdministrableDoseForm = "100000073789"
	AdministrableDoseFormEarWashSolution                        AdministrableDoseForm = "100000073790"
	AdministrableDoseFormEarTampon                              AdministrableDoseForm = "100000073791"
	AdministrableDoseFormNasalCream                             AdministrableDoseForm = "100000073792"
	AdministrableDoseFormNasalGel                               AdministrableDoseForm = "100000073793"
	AdministrableDoseFormNasalDropsSolution                     AdministrableDoseForm = "100000073794"
	AdministrableDoseFormNasalDropsEmulsion                     AdministrableDoseForm = "100000073795"
	AdministrableDoseFormNasalSpraySolution                     AdministrableDoseForm = "100000073796"
	AdministrableDoseFormNasalSprayEmulsion                     AdministrableDoseForm = "100000073797"
	AdministrableDoseFormNasalStick                             AdministrableDoseForm = "100000073798"
	AdministrableDoseFormVaginalGel                             AdministrableDoseForm = "100000073799"
	AdministrableDoseFormVaginalFoam                            AdministrableDoseForm = "100000073800"
	AdministrableDoseFormEarSpraySolution                       AdministrableDoseForm = "100000073802"
	AdministrableDoseFormEarSprayEmulsion                       AdministrableDoseForm = "100000073803"
	AdministrableDoseFormEarWashEmulsion                        AdministrableDoseForm = "100000073804"
	AdministrableDoseFormEarStick                               AdministrableDoseForm = "100000073805"
	AdministrableDoseFormNasalOintment                          AdministrableDoseForm = "100000073806"
	AdministrableDoseFormNasalDropsSuspension                   AdministrableDoseForm = "100000073807"
	AdministrableDoseFormNasalPowder                            AdministrableDoseForm = "100000073808"
	AdministrableDoseFormNasalSpraySuspension                   AdministrableDoseForm = "100000073809"
	AdministrableDoseFormNasalWash                              AdministrableDoseForm = "100000073810"
	AdministrableDoseFormVaginalCream                           AdministrableDoseForm = "100000073811"
	AdministrableDoseFormVaginalOintment                        AdministrableDoseForm = "100000073812"
	AdministrableDoseFormVaginalSolution                        AdministrableDoseForm = "100000073813"
	AdministrableDoseFormVaginalEmulsion                        AdministrableDoseForm = "100000073814"
	AdministrableDoseFormPessary                                AdministrableDoseForm = "100000073815"
	AdministrableDoseFormVaginalCapsuleSoft                     AdministrableDoseForm = "100000073816"
	AdministrableDoseFormEffervescentVaginalTablet              AdministrableDoseForm = "100000073817"
	AdministrableDoseFormVaginalDeliverySystem                  AdministrableDoseForm = "100000073818"
	AdministrableDoseFormRectalCream                            AdministrableDoseForm = "100000073819"
	AdministrableDoseFormRectalFoam                             AdministrableDoseForm = "100000073820"
	AdministrableDoseFormVaginalSuspension                      AdministrableDoseForm = "100000073821"
	AdministrableDoseFormTabletForVaginalSolution               AdministrableDoseForm = "100000073822"
	AdministrableDoseFormVaginalCapsuleHard                     AdministrableDoseForm = "100000073823"
	AdministrableDoseFormVaginalTablet                          AdministrableDoseForm = "100000073824"
	AdministrableDoseFormMedicatedVaginalTampon                 AdministrableDoseForm = "100000073825"
	AdministrableDoseFormVaginalSponge                          AdministrableDoseForm = "100000073826"
	AdministrableDoseFormRectalGel                              AdministrableDoseForm = "100000073827"
	AdministrableDoseFormSolutionForInjection                   AdministrableDoseForm = "100000073863"
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
