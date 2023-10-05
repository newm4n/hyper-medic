package code

import (
	"fmt"
	"strings"
)

type Mimetypes string

const (
	MimetypesX3D         Mimetypes = "application/vnd.hzn-3d-crossword"
	Mimetypes3GP         Mimetypes = "video/3gpp"
	Mimetypes3G2         Mimetypes = "video/3gpp2"
	MimetypesMSEQ        Mimetypes = "application/vnd.mseq"
	MimetypesPWN         Mimetypes = "application/vnd.3m.post-it-notes"
	MimetypesPLB         Mimetypes = "application/vnd.3gpp.pic-bw-large"
	MimetypesPSB         Mimetypes = "application/vnd.3gpp.pic-bw-small"
	MimetypesPVB         Mimetypes = "application/vnd.3gpp.pic-bw-var"
	MimetypesTCAP        Mimetypes = "application/vnd.3gpp2.tcap"
	Mimetypes7Z          Mimetypes = "application/x-7z-compressed"
	MimetypesABW         Mimetypes = "application/x-abiword"
	MimetypesACE         Mimetypes = "application/x-ace-compressed"
	MimetypesACC         Mimetypes = "application/vnd.americandynamics.acc"
	MimetypesACU         Mimetypes = "application/vnd.acucobol"
	MimetypesATC         Mimetypes = "application/vnd.acucorp"
	MimetypesADP         Mimetypes = "audio/adpcm"
	MimetypesAAB         Mimetypes = "application/x-authorware-bin"
	MimetypesAAM         Mimetypes = "application/x-authorware-map"
	MimetypesAAS         Mimetypes = "application/x-authorware-seg"
	MimetypesAIR         Mimetypes = "application/vnd.adobe.air-application-installer-package+zip"
	MimetypesSWF         Mimetypes = "application/x-shockwave-flash"
	MimetypesFXP         Mimetypes = "application/vnd.adobe.fxp"
	MimetypesPDF         Mimetypes = "application/pdf"
	MimetypesPPD         Mimetypes = "application/vnd.cups-ppd"
	MimetypesDIR         Mimetypes = "application/x-director"
	MimetypesXDP         Mimetypes = "application/vnd.adobe.xdp+xml"
	MimetypesXFDF        Mimetypes = "application/vnd.adobe.xfdf"
	MimetypesAAC         Mimetypes = "audio/x-aac"
	MimetypesAHEAD       Mimetypes = "application/vnd.ahead.space"
	MimetypesAZF         Mimetypes = "application/vnd.airzip.filesecure.azf"
	MimetypesAZS         Mimetypes = "application/vnd.airzip.filesecure.azs"
	MimetypesAZW         Mimetypes = "application/vnd.amazon.ebook"
	MimetypesAMI         Mimetypes = "application/vnd.amiga.ami"
	MimetypesANI         Mimetypes = "application/andrew-inset"
	MimetypesAPK         Mimetypes = "application/vnd.android.package-archive"
	MimetypesCII         Mimetypes = "application/vnd.anser-web-certificate-issue-initiation"
	MimetypesFTI         Mimetypes = "application/vnd.anser-web-funds-transfer-initiation"
	MimetypesATX         Mimetypes = "application/vnd.antix.game-component"
	MimetypesDMG         Mimetypes = "application/x-apple-diskimage"
	MimetypesMPKG        Mimetypes = "application/vnd.apple.installer+xml"
	MimetypesAW          Mimetypes = "application/applixware"
	MimetypesLES         Mimetypes = "application/vnd.hhe.lesson-player"
	MimetypesARC         Mimetypes = "application/x-freearc"
	MimetypesSWI         Mimetypes = "application/vnd.aristanetworks.swi"
	MimetypesS           Mimetypes = "text/x-asm"
	MimetypesATOMCAT     Mimetypes = "application/atomcat+xml"
	MimetypesATOMSVC     Mimetypes = "application/atomsvc+xml"
	MimetypesATOMXML     Mimetypes = "application/atom+xml"
	MimetypesAC          Mimetypes = "application/pkix-attr-cert"
	MimetypesAIF         Mimetypes = "audio/x-aiff"
	MimetypesAVI         Mimetypes = "video/x-msvideo"
	MimetypesAEP         Mimetypes = "application/vnd.audiograph"
	MimetypesDXF         Mimetypes = "image/vnd.dxf"
	MimetypesDWF         Mimetypes = "model/vnd.dwf"
	MimetypesAVIF        Mimetypes = "image/avif"
	MimetypesPAR         Mimetypes = "text/plain-bas"
	MimetypesBCPIO       Mimetypes = "application/x-bcpio"
	MimetypesBIN         Mimetypes = "application/octet-stream"
	MimetypesBMP         Mimetypes = "image/bmp"
	MimetypesTORRENT     Mimetypes = "application/x-bittorrent"
	MimetypesCOD         Mimetypes = "application/vnd.rim.cod"
	MimetypesMPM         Mimetypes = "application/vnd.blueice.multipass"
	MimetypesBMI         Mimetypes = "application/vnd.bmi"
	MimetypesSH          Mimetypes = "application/x-sh"
	MimetypesBTIF        Mimetypes = "image/prs.btif"
	MimetypesREP         Mimetypes = "application/vnd.businessobjects"
	MimetypesBZ          Mimetypes = "application/x-bzip"
	MimetypesBZ2         Mimetypes = "application/x-bzip2"
	MimetypesCSH         Mimetypes = "application/x-csh"
	MimetypesC           Mimetypes = "text/x-c"
	MimetypesCDXML       Mimetypes = "application/vnd.chemdraw+xml"
	MimetypesCSS         Mimetypes = "text/css"
	MimetypesCDA         Mimetypes = "application/x-cdf"
	MimetypesCDX         Mimetypes = "chemical/x-cdx"
	MimetypesCML         Mimetypes = "chemical/x-cml"
	MimetypesCSML        Mimetypes = "chemical/x-csml"
	MimetypesCDBCMSG     Mimetypes = "application/vnd.contact.cmsg"
	MimetypesCLA         Mimetypes = "application/vnd.claymore"
	MimetypesC4G         Mimetypes = "application/vnd.clonk.c4group"
	MimetypesSUB         Mimetypes = "image/vnd.dvb.subtitle"
	MimetypesCDMIA       Mimetypes = "application/cdmi-capability"
	MimetypesCDMIC       Mimetypes = "application/cdmi-container"
	MimetypesCDMID       Mimetypes = "application/cdmi-domain"
	MimetypesCDMIO       Mimetypes = "application/cdmi-object"
	MimetypesCDMIQ       Mimetypes = "application/cdmi-queue"
	MimetypesC11AMC      Mimetypes = "application/vnd.cluetrust.cartomobile-config"
	MimetypesC11AMZ      Mimetypes = "application/vnd.cluetrust.cartomobile-config-pkg"
	MimetypesRAS         Mimetypes = "image/x-cmu-raster"
	MimetypesDAE         Mimetypes = "model/vnd.collada+xml"
	MimetypesCSV         Mimetypes = "text/csv"
	MimetypesCPT         Mimetypes = "application/mac-compactpro"
	MimetypesWMLC        Mimetypes = "application/vnd.wap.wmlc"
	MimetypesCGM         Mimetypes = "image/cgm"
	MimetypesICE         Mimetypes = "x-conference/x-cooltalk"
	MimetypesCMX         Mimetypes = "image/x-cmx"
	MimetypesXAR         Mimetypes = "application/vnd.xara"
	MimetypesCMC         Mimetypes = "application/vnd.cosmocaller"
	MimetypesCPIO        Mimetypes = "application/x-cpio"
	MimetypesCLKX        Mimetypes = "application/vnd.crick.clicker"
	MimetypesCLKK        Mimetypes = "application/vnd.crick.clicker.keyboard"
	MimetypesCLKP        Mimetypes = "application/vnd.crick.clicker.palette"
	MimetypesCLKT        Mimetypes = "application/vnd.crick.clicker.template"
	MimetypesCLKW        Mimetypes = "application/vnd.crick.clicker.wordbank"
	MimetypesWBS         Mimetypes = "application/vnd.criticaltools.wbs+xml"
	MimetypesCRYPTONOTE  Mimetypes = "application/vnd.rig.cryptonote"
	MimetypesCIF         Mimetypes = "chemical/x-cif"
	MimetypesCMDF        Mimetypes = "chemical/x-cmdf"
	MimetypesCU          Mimetypes = "application/cu-seeme"
	MimetypesCWW         Mimetypes = "application/prs.cww"
	MimetypesCURL        Mimetypes = "text/vnd.curl"
	MimetypesDCURL       Mimetypes = "text/vnd.curl.dcurl"
	MimetypesMCURL       Mimetypes = "text/vnd.curl.mcurl"
	MimetypesSCURL       Mimetypes = "text/vnd.curl.scurl"
	MimetypesCAR         Mimetypes = "application/vnd.curl.car"
	MimetypesPCURL       Mimetypes = "application/vnd.curl.pcurl"
	MimetypesCMP         Mimetypes = "application/vnd.yellowriver-custom-menu"
	MimetypesDSSC        Mimetypes = "application/dssc+der"
	MimetypesXDSSC       Mimetypes = "application/dssc+xml"
	MimetypesDEB         Mimetypes = "application/x-debian-package"
	MimetypesUVA         Mimetypes = "audio/vnd.dece.audio"
	MimetypesUVI         Mimetypes = "image/vnd.dece.graphic"
	MimetypesUVH         Mimetypes = "video/vnd.dece.hd"
	MimetypesUVM         Mimetypes = "video/vnd.dece.mobile"
	MimetypesUVU         Mimetypes = "video/vnd.uvvu.mp4"
	MimetypesUVP         Mimetypes = "video/vnd.dece.pd"
	MimetypesUVS         Mimetypes = "video/vnd.dece.sd"
	MimetypesUVV         Mimetypes = "video/vnd.dece.video"
	MimetypesDVI         Mimetypes = "application/x-dvi"
	MimetypesSEED        Mimetypes = "application/vnd.fdsn.seed"
	MimetypesDTB         Mimetypes = "application/x-dtbook+xml"
	MimetypesRES         Mimetypes = "application/x-dtbresource+xml"
	MimetypesAIT         Mimetypes = "application/vnd.dvb.ait"
	MimetypesSVC         Mimetypes = "application/vnd.dvb.service"
	MimetypesEOL         Mimetypes = "audio/vnd.digital-winds"
	MimetypesDJVU        Mimetypes = "image/vnd.djvu"
	MimetypesDTD         Mimetypes = "application/xml-dtd"
	MimetypesMLP         Mimetypes = "application/vnd.dolby.mlp"
	MimetypesWAD         Mimetypes = "application/x-doom"
	MimetypesDPG         Mimetypes = "application/vnd.dpgraph"
	MimetypesDRA         Mimetypes = "audio/vnd.dra"
	MimetypesDFAC        Mimetypes = "application/vnd.dreamfactory"
	MimetypesDTS         Mimetypes = "audio/vnd.dts"
	MimetypesDTSHD       Mimetypes = "audio/vnd.dts.hd"
	MimetypesDWG         Mimetypes = "image/vnd.dwg"
	MimetypesGEO         Mimetypes = "application/vnd.dynageo"
	MimetypesES          Mimetypes = "application/ecmascript"
	MimetypesMAG         Mimetypes = "application/vnd.ecowin.chart"
	MimetypesMMR         Mimetypes = "image/vnd.fujixerox.edmics-mmr"
	MimetypesRLC         Mimetypes = "image/vnd.fujixerox.edmics-rlc"
	MimetypesEXI         Mimetypes = "application/exi"
	MimetypesMGZ         Mimetypes = "application/vnd.proteus.magazine"
	MimetypesEPUB        Mimetypes = "application/epub+zip"
	MimetypesEML         Mimetypes = "message/rfc822"
	MimetypesNML         Mimetypes = "application/vnd.enliven"
	MimetypesXPR         Mimetypes = "application/vnd.is-xpr"
	MimetypesXIF         Mimetypes = "image/vnd.xiff"
	MimetypesXFDL        Mimetypes = "application/vnd.xfdl"
	MimetypesEMMA        Mimetypes = "application/emma+xml"
	MimetypesEZ2         Mimetypes = "application/vnd.ezpix-album"
	MimetypesEZ3         Mimetypes = "application/vnd.ezpix-package"
	MimetypesFST         Mimetypes = "image/vnd.fst"
	MimetypesFVT         Mimetypes = "video/vnd.fvt"
	MimetypesFBS         Mimetypes = "image/vnd.fastbidsheet"
	MimetypesFE_LAUNCH   Mimetypes = "application/vnd.denovo.fcselayout-link"
	MimetypesF4V         Mimetypes = "video/x-f4v"
	MimetypesFLV         Mimetypes = "video/x-flv"
	MimetypesFPX         Mimetypes = "image/vnd.fpx"
	MimetypesNPX         Mimetypes = "image/vnd.net-fpx"
	MimetypesFLX         Mimetypes = "text/vnd.fmi.flexstor"
	MimetypesFLI         Mimetypes = "video/x-fli"
	MimetypesFTC         Mimetypes = "application/vnd.fluxtime.clip"
	MimetypesFDF         Mimetypes = "application/vnd.fdf"
	MimetypesF           Mimetypes = "text/x-fortran"
	MimetypesMIF         Mimetypes = "application/vnd.mif"
	MimetypesFM          Mimetypes = "application/vnd.framemaker"
	MimetypesFH          Mimetypes = "image/x-freehand"
	MimetypesFSC         Mimetypes = "application/vnd.fsc.weblaunch"
	MimetypesFNC         Mimetypes = "application/vnd.frogans.fnc"
	MimetypesLTF         Mimetypes = "application/vnd.frogans.ltf"
	MimetypesDDD         Mimetypes = "application/vnd.fujixerox.ddd"
	MimetypesXDW         Mimetypes = "application/vnd.fujixerox.docuworks"
	MimetypesXBD         Mimetypes = "application/vnd.fujixerox.docuworks.binder"
	MimetypesOAS         Mimetypes = "application/vnd.fujitsu.oasys"
	MimetypesOA2         Mimetypes = "application/vnd.fujitsu.oasys2"
	MimetypesOA3         Mimetypes = "application/vnd.fujitsu.oasys3"
	MimetypesFG5         Mimetypes = "application/vnd.fujitsu.oasysgp"
	MimetypesBH2         Mimetypes = "application/vnd.fujitsu.oasysprs"
	MimetypesSPL         Mimetypes = "application/x-futuresplash"
	MimetypesFZS         Mimetypes = "application/vnd.fuzzysheet"
	MimetypesG3          Mimetypes = "image/g3fax"
	MimetypesGMX         Mimetypes = "application/vnd.gmx"
	MimetypesGTW         Mimetypes = "model/vnd.gtw"
	MimetypesTXD         Mimetypes = "application/vnd.genomatix.tuxedo"
	MimetypesGGB         Mimetypes = "application/vnd.geogebra.file"
	MimetypesGGT         Mimetypes = "application/vnd.geogebra.tool"
	MimetypesGDL         Mimetypes = "model/vnd.gdl"
	MimetypesGEX         Mimetypes = "application/vnd.geometry-explorer"
	MimetypesGXT         Mimetypes = "application/vnd.geonext"
	MimetypesG2W         Mimetypes = "application/vnd.geoplan"
	MimetypesG3W         Mimetypes = "application/vnd.geospace"
	MimetypesGSF         Mimetypes = "application/x-font-ghostscript"
	MimetypesBDF         Mimetypes = "application/x-font-bdf"
	MimetypesGTAR        Mimetypes = "application/x-gtar"
	MimetypesTEXINFO     Mimetypes = "application/x-texinfo"
	MimetypesGNUMERIC    Mimetypes = "application/x-gnumeric"
	MimetypesKML         Mimetypes = "application/vnd.google-earth.kml+xml"
	MimetypesKMZ         Mimetypes = "application/vnd.google-earth.kmz"
	MimetypesGPX         Mimetypes = "application/gpx+xml"
	MimetypesGQF         Mimetypes = "application/vnd.grafeq"
	MimetypesGIF         Mimetypes = "image/gif"
	MimetypesGV          Mimetypes = "text/vnd.graphviz"
	MimetypesGAC         Mimetypes = "application/vnd.groove-account"
	MimetypesGHF         Mimetypes = "application/vnd.groove-help"
	MimetypesGIM         Mimetypes = "application/vnd.groove-identity-message"
	MimetypesGRV         Mimetypes = "application/vnd.groove-injector"
	MimetypesGTM         Mimetypes = "application/vnd.groove-tool-message"
	MimetypesTPL         Mimetypes = "application/vnd.groove-tool-template"
	MimetypesVCG         Mimetypes = "application/vnd.groove-vcard"
	MimetypesGZ          Mimetypes = "application/gzip"
	MimetypesH261        Mimetypes = "video/h261"
	MimetypesH263        Mimetypes = "video/h263"
	MimetypesH264        Mimetypes = "video/h264"
	MimetypesHPID        Mimetypes = "application/vnd.hp-hpid"
	MimetypesHPS         Mimetypes = "application/vnd.hp-hps"
	MimetypesHDF         Mimetypes = "application/x-hdf"
	MimetypesRIP         Mimetypes = "audio/vnd.rip"
	MimetypesHBCI        Mimetypes = "application/vnd.hbci"
	MimetypesJLT         Mimetypes = "application/vnd.hp-jlyt"
	MimetypesPCL         Mimetypes = "application/vnd.hp-pcl"
	MimetypesHPGL        Mimetypes = "application/vnd.hp-hpgl"
	MimetypesHVS         Mimetypes = "application/vnd.yamaha.hv-script"
	MimetypesHVD         Mimetypes = "application/vnd.yamaha.hv-dic"
	MimetypesHVP         Mimetypes = "application/vnd.yamaha.hv-voice"
	MimetypesSFHDSTX     Mimetypes = "application/vnd.hydrostatix.sof-data"
	MimetypesSTK         Mimetypes = "application/hyperstudio"
	MimetypesHAL         Mimetypes = "application/vnd.hal+xml"
	MimetypesHTML        Mimetypes = "text/html"
	MimetypesIRM         Mimetypes = "application/vnd.ibm.rights-management"
	MimetypesSC          Mimetypes = "application/vnd.ibm.secure-container"
	MimetypesICS         Mimetypes = "text/calendar"
	MimetypesICC         Mimetypes = "application/vnd.iccprofile"
	MimetypesICO         Mimetypes = "image/x-icon"
	MimetypesIGL         Mimetypes = "application/vnd.igloader"
	MimetypesIEF         Mimetypes = "image/ief"
	MimetypesIVP         Mimetypes = "application/vnd.immervision-ivp"
	MimetypesIVU         Mimetypes = "application/vnd.immervision-ivu"
	MimetypesRIF         Mimetypes = "application/reginfo+xml"
	Mimetypes3DML        Mimetypes = "text/vnd.in3d.3dml"
	MimetypesSPOT        Mimetypes = "text/vnd.in3d.spot"
	MimetypesIGS         Mimetypes = "model/iges"
	MimetypesI2G         Mimetypes = "application/vnd.intergeo"
	MimetypesCDY         Mimetypes = "application/vnd.cinderella"
	MimetypesXPW         Mimetypes = "application/vnd.intercon.formnet"
	MimetypesFCS         Mimetypes = "application/vnd.isac.fcs"
	MimetypesIPFIX       Mimetypes = "application/ipfix"
	MimetypesCER         Mimetypes = "application/pkix-cert"
	MimetypesPKI         Mimetypes = "application/pkixcmp"
	MimetypesCRL         Mimetypes = "application/pkix-crl"
	MimetypesPKIPATH     Mimetypes = "application/pkix-pkipath"
	MimetypesIGM         Mimetypes = "application/vnd.insors.igm"
	MimetypesRCPROFILE   Mimetypes = "application/vnd.ipunplugged.rcprofile"
	MimetypesIRP         Mimetypes = "application/vnd.irepository.package+xml"
	MimetypesJAD         Mimetypes = "text/vnd.sun.j2me.app-descriptor"
	MimetypesJAR         Mimetypes = "application/java-archive"
	MimetypesCLASS       Mimetypes = "application/java-vm"
	MimetypesJNLP        Mimetypes = "application/x-java-jnlp-file"
	MimetypesSER         Mimetypes = "application/java-serialized-object"
	MimetypesJAVA        Mimetypes = "text/x-java-source,java"
	MimetypesJS          Mimetypes = "application/javascript"
	MimetypesTXTJS       Mimetypes = "text/javascript"
	MimetypesMJS         Mimetypes = "text/javascript"
	MimetypesJSON        Mimetypes = "application/json"
	MimetypesJODA        Mimetypes = "application/vnd.joost.joda-archive"
	MimetypesJPM         Mimetypes = "video/jpm"
	MimetypesIMGJPEG     Mimetypes = "image/jpeg"
	MimetypesCITRIXJPEG  Mimetypes = "image/x-citrix-jpeg"
	MimetypesPJPEG       Mimetypes = "image/pjpeg"
	MimetypesJPGV        Mimetypes = "video/jpeg"
	MimetypesJSONLD      Mimetypes = "application/ld+json"
	MimetypesKTZ         Mimetypes = "application/vnd.kahootz"
	MimetypesMMD         Mimetypes = "application/vnd.chipnuts.karaoke-mmd"
	MimetypesKARBON      Mimetypes = "application/vnd.kde.karbon"
	MimetypesCHRT        Mimetypes = "application/vnd.kde.kchart"
	MimetypesKFO         Mimetypes = "application/vnd.kde.kformula"
	MimetypesFLW         Mimetypes = "application/vnd.kde.kivio"
	MimetypesKON         Mimetypes = "application/vnd.kde.kontour"
	MimetypesKPR         Mimetypes = "application/vnd.kde.kpresenter"
	MimetypesKSP         Mimetypes = "application/vnd.kde.kspread"
	MimetypesKWD         Mimetypes = "application/vnd.kde.kword"
	MimetypesHTKE        Mimetypes = "application/vnd.kenameaapp"
	MimetypesKIA         Mimetypes = "application/vnd.kidspiration"
	MimetypesKNE         Mimetypes = "application/vnd.kinar"
	MimetypesSSE         Mimetypes = "application/vnd.kodak-descriptor"
	MimetypesLASXML      Mimetypes = "application/vnd.las.las+xml"
	MimetypesLATEX       Mimetypes = "application/x-latex"
	MimetypesLBD         Mimetypes = "application/vnd.llamagraphics.life-balance.desktop"
	MimetypesLBE         Mimetypes = "application/vnd.llamagraphics.life-balance.exchange+xml"
	MimetypesJAM         Mimetypes = "application/vnd.jam"
	Mimetypes123         Mimetypes = "application/vnd.lotus-1-2-3"
	MimetypesAPR         Mimetypes = "application/vnd.lotus-approach"
	MimetypesPRE         Mimetypes = "application/vnd.lotus-freelance"
	MimetypesNSF         Mimetypes = "application/vnd.lotus-notes"
	MimetypesORG         Mimetypes = "application/vnd.lotus-organizer"
	MimetypesSCM         Mimetypes = "application/vnd.lotus-screencam"
	MimetypesLWP         Mimetypes = "application/vnd.lotus-wordpro"
	MimetypesLVP         Mimetypes = "audio/vnd.lucent.voice"
	MimetypesM3U         Mimetypes = "audio/x-mpegurl"
	MimetypesM4V         Mimetypes = "video/x-m4v"
	MimetypesHQX         Mimetypes = "application/mac-binhex40"
	MimetypesPORTPKG     Mimetypes = "application/vnd.macports.portpkg"
	MimetypesMGP         Mimetypes = "application/vnd.osgeo.mapguide.package"
	MimetypesMRC         Mimetypes = "application/marc"
	MimetypesMRCX        Mimetypes = "application/marcxml+xml"
	MimetypesMXF         Mimetypes = "application/mxf"
	MimetypesNBP         Mimetypes = "application/vnd.wolfram.player"
	MimetypesMA          Mimetypes = "application/mathematica"
	MimetypesMATHML      Mimetypes = "application/mathml+xml"
	MimetypesMBOX        Mimetypes = "application/mbox"
	MimetypesMC1         Mimetypes = "application/vnd.medcalcdata"
	MimetypesMSCML       Mimetypes = "application/mediaservercontrol+xml"
	MimetypesCDKEY       Mimetypes = "application/vnd.mediastation.cdkey"
	MimetypesMWF         Mimetypes = "application/vnd.mfer"
	MimetypesMFM         Mimetypes = "application/vnd.mfmp"
	MimetypesMSH         Mimetypes = "model/mesh"
	MimetypesMADS        Mimetypes = "application/mads+xml"
	MimetypesMETS        Mimetypes = "application/mets+xml"
	MimetypesMODS        Mimetypes = "application/mods+xml"
	MimetypesMETA4       Mimetypes = "application/metalink4+xml"
	MimetypesMCD         Mimetypes = "application/vnd.mcd"
	MimetypesFLO         Mimetypes = "application/vnd.micrografx.flo"
	MimetypesIGX         Mimetypes = "application/vnd.micrografx.igx"
	MimetypesES3         Mimetypes = "application/vnd.eszigno3+xml"
	MimetypesMDB         Mimetypes = "application/x-msaccess"
	MimetypesASF         Mimetypes = "video/x-ms-asf"
	MimetypesEXE         Mimetypes = "application/x-msdownload"
	MimetypesCIL         Mimetypes = "application/vnd.ms-artgalry"
	MimetypesCAB         Mimetypes = "application/vnd.ms-cab-compressed"
	MimetypesIMS         Mimetypes = "application/vnd.ms-ims"
	MimetypesAPPLICATION Mimetypes = "application/x-ms-application"
	MimetypesCLP         Mimetypes = "application/x-msclip"
	MimetypesMDI         Mimetypes = "image/vnd.ms-modi"
	MimetypesEOT         Mimetypes = "application/vnd.ms-fontobject"
	MimetypesXLS         Mimetypes = "application/vnd.ms-excel"
	MimetypesXLAM        Mimetypes = "application/vnd.ms-excel.addin.macroenabled.12"
	MimetypesXLSB        Mimetypes = "application/vnd.ms-excel.sheet.binary.macroenabled.12"
	MimetypesXLTM        Mimetypes = "application/vnd.ms-excel.template.macroenabled.12"
	MimetypesXLSM        Mimetypes = "application/vnd.ms-excel.sheet.macroenabled.12"
	MimetypesCHM         Mimetypes = "application/vnd.ms-htmlhelp"
	MimetypesCRD         Mimetypes = "application/x-mscardfile"
	MimetypesLRM         Mimetypes = "application/vnd.ms-lrm"
	MimetypesMVB         Mimetypes = "application/x-msmediaview"
	MimetypesMNY         Mimetypes = "application/x-msmoney"
	MimetypesPPTX        Mimetypes = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MimetypesSLDX        Mimetypes = "application/vnd.openxmlformats-officedocument.presentationml.slide"
	MimetypesPPSX        Mimetypes = "application/vnd.openxmlformats-officedocument.presentationml.slideshow"
	MimetypesPOTX        Mimetypes = "application/vnd.openxmlformats-officedocument.presentationml.template"
	MimetypesXLSX        Mimetypes = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MimetypesXLTX        Mimetypes = "application/vnd.openxmlformats-officedocument.spreadsheetml.template"
	MimetypesDOCX        Mimetypes = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MimetypesDOTX        Mimetypes = "application/vnd.openxmlformats-officedocument.wordprocessingml.template"
	MimetypesOBD         Mimetypes = "application/x-msbinder"
	MimetypesTHMX        Mimetypes = "application/vnd.ms-officetheme"
	MimetypesONETOC      Mimetypes = "application/onenote"
	MimetypesPYA         Mimetypes = "audio/vnd.ms-playready.media.pya"
	MimetypesPYV         Mimetypes = "video/vnd.ms-playready.media.pyv"
	MimetypesPPT         Mimetypes = "application/vnd.ms-powerpoint"
	MimetypesPPAM        Mimetypes = "application/vnd.ms-powerpoint.addin.macroenabled.12"
	MimetypesSLDM        Mimetypes = "application/vnd.ms-powerpoint.slide.macroenabled.12"
	MimetypesPPTM        Mimetypes = "application/vnd.ms-powerpoint.presentation.macroenabled.12"
	MimetypesPPSM        Mimetypes = "application/vnd.ms-powerpoint.slideshow.macroenabled.12"
	MimetypesPOTM        Mimetypes = "application/vnd.ms-powerpoint.template.macroenabled.12"
	MimetypesMPP         Mimetypes = "application/vnd.ms-project"
	MimetypesPUB         Mimetypes = "application/x-mspublisher"
	MimetypesSCD         Mimetypes = "application/x-msschedule"
	MimetypesXAP         Mimetypes = "application/x-silverlight-app"
	MimetypesSTL         Mimetypes = "application/vnd.ms-pki.stl"
	MimetypesCAT         Mimetypes = "application/vnd.ms-pki.seccat"
	MimetypesVSD         Mimetypes = "application/vnd.visio"
	MimetypesVSDX        Mimetypes = "application/vnd.visio2013"
	MimetypesWM          Mimetypes = "video/x-ms-wm"
	MimetypesWMA         Mimetypes = "audio/x-ms-wma"
	MimetypesWAX         Mimetypes = "audio/x-ms-wax"
	MimetypesWMX         Mimetypes = "video/x-ms-wmx"
	MimetypesWMD         Mimetypes = "application/x-ms-wmd"
	MimetypesWPL         Mimetypes = "application/vnd.ms-wpl"
	MimetypesWMZ         Mimetypes = "application/x-ms-wmz"
	MimetypesWMV         Mimetypes = "video/x-ms-wmv"
	MimetypesWVX         Mimetypes = "video/x-ms-wvx"
	MimetypesWMF         Mimetypes = "application/x-msmetafile"
	MimetypesTRM         Mimetypes = "application/x-msterminal"
	MimetypesDOC         Mimetypes = "application/msword"
	MimetypesDOCM        Mimetypes = "application/vnd.ms-word.document.macroenabled.12"
	MimetypesDOTM        Mimetypes = "application/vnd.ms-word.template.macroenabled.12"
	MimetypesWRI         Mimetypes = "application/x-mswrite"
	MimetypesWPS         Mimetypes = "application/vnd.ms-works"
	MimetypesXBAP        Mimetypes = "application/x-ms-xbap"
	MimetypesXPS         Mimetypes = "application/vnd.ms-xpsdocument"
	MimetypesMIDI        Mimetypes = "audio/midi"
	MimetypesMID         Mimetypes = "audio/midi"
	MimetypesMPY         Mimetypes = "application/vnd.ibm.minipay"
	MimetypesAFP         Mimetypes = "application/vnd.ibm.modcap"
	MimetypesRMS         Mimetypes = "application/vnd.jcp.javame.midlet-rms"
	MimetypesTMO         Mimetypes = "application/vnd.tmobile-livetv"
	MimetypesPRC         Mimetypes = "application/x-mobipocket-ebook"
	MimetypesMBK         Mimetypes = "application/vnd.mobius.mbk"
	MimetypesDIS         Mimetypes = "application/vnd.mobius.dis"
	MimetypesPLC         Mimetypes = "application/vnd.mobius.plc"
	MimetypesMQY         Mimetypes = "application/vnd.mobius.mqy"
	MimetypesMSL         Mimetypes = "application/vnd.mobius.msl"
	MimetypesTXF         Mimetypes = "application/vnd.mobius.txf"
	MimetypesDAF         Mimetypes = "application/vnd.mobius.daf"
	MimetypesFLY         Mimetypes = "text/vnd.fly"
	MimetypesMPC         Mimetypes = "application/vnd.mophun.certificate"
	MimetypesMPN         Mimetypes = "application/vnd.mophun.application"
	MimetypesMJ2         Mimetypes = "video/mj2"
	MimetypesMPGA        Mimetypes = "audio/mpeg"
	MimetypesTS          Mimetypes = "video/mp2t"
	MimetypesMXU         Mimetypes = "video/vnd.mpegurl"
	MimetypesMPEG        Mimetypes = "video/mpeg"
	MimetypesM21         Mimetypes = "application/mp21"
	MimetypesMP4A        Mimetypes = "audio/mp4"
	MimetypesVDMP4       Mimetypes = "video/mp4"
	MimetypesMP4         Mimetypes = "application/mp4"
	MimetypesM3U8        Mimetypes = "application/vnd.apple.mpegurl"
	MimetypesMUS         Mimetypes = "application/vnd.musician"
	MimetypesMSTY        Mimetypes = "application/vnd.muvee.style"
	MimetypesMXML        Mimetypes = "application/xv+xml"
	MimetypesNGDAT       Mimetypes = "application/vnd.nokia.n-gage.data"
	MimetypesNGAGE       Mimetypes = "application/vnd.nokia.n-gage.symbian.install"
	MimetypesNCX         Mimetypes = "application/x-dtbncx+xml"
	MimetypesNC          Mimetypes = "application/x-netcdf"
	MimetypesNLU         Mimetypes = "application/vnd.neurolanguage.nlu"
	MimetypesDNA         Mimetypes = "application/vnd.dna"
	MimetypesNND         Mimetypes = "application/vnd.noblenet-directory"
	MimetypesNNS         Mimetypes = "application/vnd.noblenet-sealer"
	MimetypesNNW         Mimetypes = "application/vnd.noblenet-web"
	MimetypesRPST        Mimetypes = "application/vnd.nokia.radio-preset"
	MimetypesRPSS        Mimetypes = "application/vnd.nokia.radio-presets"
	MimetypesN3          Mimetypes = "text/n3"
	MimetypesEDM         Mimetypes = "application/vnd.novadigm.edm"
	MimetypesEDX         Mimetypes = "application/vnd.novadigm.edx"
	MimetypesEXT         Mimetypes = "application/vnd.novadigm.ext"
	MimetypesGPH         Mimetypes = "application/vnd.flographit"
	MimetypesECELP4800   Mimetypes = "audio/vnd.nuera.ecelp4800"
	MimetypesECELP7470   Mimetypes = "audio/vnd.nuera.ecelp7470"
	MimetypesECELP9600   Mimetypes = "audio/vnd.nuera.ecelp9600"
	MimetypesODA         Mimetypes = "application/oda"
	MimetypesOGX         Mimetypes = "application/ogg"
	MimetypesOGA         Mimetypes = "audio/ogg"
	MimetypesOGV         Mimetypes = "video/ogg"
	MimetypesDD2         Mimetypes = "application/vnd.oma.dd2+xml"
	MimetypesOTH         Mimetypes = "application/vnd.oasis.opendocument.text-web"
	MimetypesOPF         Mimetypes = "application/oebps-package+xml"
	MimetypesQBO         Mimetypes = "application/vnd.intu.qbo"
	MimetypesOXT         Mimetypes = "application/vnd.openofficeorg.extension"
	MimetypesOSF         Mimetypes = "application/vnd.yamaha.openscoreformat"
	MimetypesWEBA        Mimetypes = "audio/webm"
	MimetypesWEBM        Mimetypes = "video/webm"
	MimetypesODC         Mimetypes = "application/vnd.oasis.opendocument.chart"
	MimetypesOTC         Mimetypes = "application/vnd.oasis.opendocument.chart-template"
	MimetypesODB         Mimetypes = "application/vnd.oasis.opendocument.database"
	MimetypesODF         Mimetypes = "application/vnd.oasis.opendocument.formula"
	MimetypesODFT        Mimetypes = "application/vnd.oasis.opendocument.formula-template"
	MimetypesODG         Mimetypes = "application/vnd.oasis.opendocument.graphics"
	MimetypesOTG         Mimetypes = "application/vnd.oasis.opendocument.graphics-template"
	MimetypesODI         Mimetypes = "application/vnd.oasis.opendocument.image"
	MimetypesOTI         Mimetypes = "application/vnd.oasis.opendocument.image-template"
	MimetypesODP         Mimetypes = "application/vnd.oasis.opendocument.presentation"
	MimetypesOTP         Mimetypes = "application/vnd.oasis.opendocument.presentation-template"
	MimetypesODS         Mimetypes = "application/vnd.oasis.opendocument.spreadsheet"
	MimetypesOTS         Mimetypes = "application/vnd.oasis.opendocument.spreadsheet-template"
	MimetypesODT         Mimetypes = "application/vnd.oasis.opendocument.text"
	MimetypesODM         Mimetypes = "application/vnd.oasis.opendocument.text-master"
	MimetypesOTT         Mimetypes = "application/vnd.oasis.opendocument.text-template"
	MimetypesKTX         Mimetypes = "image/ktx"
	MimetypesSXC         Mimetypes = "application/vnd.sun.xml.calc"
	MimetypesSTC         Mimetypes = "application/vnd.sun.xml.calc.template"
	MimetypesSXD         Mimetypes = "application/vnd.sun.xml.draw"
	MimetypesSTD         Mimetypes = "application/vnd.sun.xml.draw.template"
	MimetypesSXI         Mimetypes = "application/vnd.sun.xml.impress"
	MimetypesSTI         Mimetypes = "application/vnd.sun.xml.impress.template"
	MimetypesSXM         Mimetypes = "application/vnd.sun.xml.math"
	MimetypesSXW         Mimetypes = "application/vnd.sun.xml.writer"
	MimetypesSXG         Mimetypes = "application/vnd.sun.xml.writer.global"
	MimetypesSTW         Mimetypes = "application/vnd.sun.xml.writer.template"
	MimetypesOTF         Mimetypes = "application/x-font-otf"
	MimetypesOPUS        Mimetypes = "audio/opus"
	MimetypesOSFPVG      Mimetypes = "application/vnd.yamaha.openscoreformat.osfpvg+xml"
	MimetypesDP          Mimetypes = "application/vnd.osgi.dp"
	MimetypesPDB         Mimetypes = "application/vnd.palm"
	MimetypesP           Mimetypes = "text/x-pascal"
	MimetypesPAW         Mimetypes = "application/vnd.pawaafile"
	MimetypesPCLXL       Mimetypes = "application/vnd.hp-pclxl"
	MimetypesEFIF        Mimetypes = "application/vnd.picsel"
	MimetypesPCX         Mimetypes = "image/x-pcx"
	MimetypesPSD         Mimetypes = "image/vnd.adobe.photoshop"
	MimetypesPRF         Mimetypes = "application/pics-rules"
	MimetypesPIC         Mimetypes = "image/x-pict"
	MimetypesCHAT        Mimetypes = "application/x-chat"
	MimetypesP10         Mimetypes = "application/pkcs10"
	MimetypesP12         Mimetypes = "application/x-pkcs12"
	MimetypesP7M         Mimetypes = "application/pkcs7-mime"
	MimetypesP7S         Mimetypes = "application/pkcs7-signature"
	MimetypesP7R         Mimetypes = "application/x-pkcs7-certreqresp"
	MimetypesP7B         Mimetypes = "application/x-pkcs7-certificates"
	MimetypesP8          Mimetypes = "application/pkcs8"
	MimetypesPLF         Mimetypes = "application/vnd.pocketlearn"
	MimetypesPNM         Mimetypes = "image/x-portable-anymap"
	MimetypesPBM         Mimetypes = "image/x-portable-bitmap"
	MimetypesPCF         Mimetypes = "application/x-font-pcf"
	MimetypesPFR         Mimetypes = "application/font-tdpfr"
	MimetypesPGN         Mimetypes = "application/x-chess-pgn"
	MimetypesPGM         Mimetypes = "image/x-portable-graymap"
	MimetypesIMGPNG      Mimetypes = "image/png"
	MimetypesPNG         Mimetypes = "image/x-citrix-png"
	MimetypesXPNG        Mimetypes = "image/x-png"
	MimetypesPPM         Mimetypes = "image/x-portable-pixmap"
	MimetypesPSKCXML     Mimetypes = "application/pskc+xml"
	MimetypesPML         Mimetypes = "application/vnd.ctc-posml"
	MimetypesAI          Mimetypes = "application/postscript"
	MimetypesPFA         Mimetypes = "application/x-font-type1"
	MimetypesPBD         Mimetypes = "application/vnd.powerbuilder6"
	MimetypesCRYPGP      Mimetypes = "application/pgp-encrypted"
	MimetypesSIGPGP      Mimetypes = "application/pgp-signature"
	MimetypesBOX         Mimetypes = "application/vnd.previewsystems.box"
	MimetypesPTID        Mimetypes = "application/vnd.pvi.ptid1"
	MimetypesPLS         Mimetypes = "application/pls+xml"
	MimetypesSTR         Mimetypes = "application/vnd.pg.format"
	MimetypesEI6         Mimetypes = "application/vnd.pg.osasli"
	MimetypesDSC         Mimetypes = "text/prs.lines.tag"
	MimetypesPSF         Mimetypes = "application/x-font-linux-psf"
	MimetypesQPS         Mimetypes = "application/vnd.publishare-delta-tree"
	MimetypesWG          Mimetypes = "application/vnd.pmi.widget"
	MimetypesQXD         Mimetypes = "application/vnd.quark.quarkxpress"
	MimetypesESF         Mimetypes = "application/vnd.epson.esf"
	MimetypesMSF         Mimetypes = "application/vnd.epson.msf"
	MimetypesSSF         Mimetypes = "application/vnd.epson.ssf"
	MimetypesQAM         Mimetypes = "application/vnd.epson.quickanime"
	MimetypesQFX         Mimetypes = "application/vnd.intu.qfx"
	MimetypesQT          Mimetypes = "video/quicktime"
	MimetypesRAR         Mimetypes = "application/x-rar-compressed"
	MimetypesRAM         Mimetypes = "audio/x-pn-realaudio"
	MimetypesRMP         Mimetypes = "audio/x-pn-realaudio-plugin"
	MimetypesRSD         Mimetypes = "application/rsd+xml"
	MimetypesRM          Mimetypes = "application/vnd.rn-realmedia"
	MimetypesBED         Mimetypes = "application/vnd.realvnc.bed"
	MimetypesMXL         Mimetypes = "application/vnd.recordare.musicxml"
	MimetypesMUSICXML    Mimetypes = "application/vnd.recordare.musicxml+xml"
	MimetypesRNC         Mimetypes = "application/relax-ng-compact-syntax"
	MimetypesRDZ         Mimetypes = "application/vnd.data-vision.rdz"
	MimetypesRDF         Mimetypes = "application/rdf+xml"
	MimetypesRP9         Mimetypes = "application/vnd.cloanto.rp9"
	MimetypesJISP        Mimetypes = "application/vnd.jisp"
	MimetypesRTF         Mimetypes = "application/rtf"
	MimetypesRTX         Mimetypes = "text/richtext"
	MimetypesLINK66      Mimetypes = "application/vnd.route66.link66+xml"
	MimetypesRSSXML      Mimetypes = "application/rss+xml"
	MimetypesSHF         Mimetypes = "application/shf+xml"
	MimetypesST          Mimetypes = "application/vnd.sailingtracker.track"
	MimetypesSVG         Mimetypes = "image/svg+xml"
	MimetypesSUS         Mimetypes = "application/vnd.sus-calendar"
	MimetypesSRU         Mimetypes = "application/sru+xml"
	MimetypesSETPAY      Mimetypes = "application/set-payment-initiation"
	MimetypesSETREG      Mimetypes = "application/set-registration-initiation"
	MimetypesSEMA        Mimetypes = "application/vnd.sema"
	MimetypesSEMD        Mimetypes = "application/vnd.semd"
	MimetypesSEMF        Mimetypes = "application/vnd.semf"
	MimetypesSEE         Mimetypes = "application/vnd.seemail"
	MimetypesSNF         Mimetypes = "application/x-font-snf"
	MimetypesSPQ         Mimetypes = "application/scvp-vp-request"
	MimetypesSPP         Mimetypes = "application/scvp-vp-response"
	MimetypesSCQ         Mimetypes = "application/scvp-cv-request"
	MimetypesSCS         Mimetypes = "application/scvp-cv-response"
	MimetypesSDP         Mimetypes = "application/sdp"
	MimetypesETX         Mimetypes = "text/x-setext"
	MimetypesMOVIE       Mimetypes = "video/x-sgi-movie"
	MimetypesIFM         Mimetypes = "application/vnd.shana.informed.formdata"
	MimetypesITP         Mimetypes = "application/vnd.shana.informed.formtemplate"
	MimetypesIIF         Mimetypes = "application/vnd.shana.informed.interchange"
	MimetypesIPK         Mimetypes = "application/vnd.shana.informed.package"
	MimetypesTFI         Mimetypes = "application/thraud+xml"
	MimetypesSHAR        Mimetypes = "application/x-shar"
	MimetypesRGB         Mimetypes = "image/x-rgb"
	MimetypesSLT         Mimetypes = "application/vnd.epson.salt"
	MimetypesASO         Mimetypes = "application/vnd.accpac.simply.aso"
	MimetypesIMP         Mimetypes = "application/vnd.accpac.simply.imp"
	MimetypesTWD         Mimetypes = "application/vnd.simtech-mindmapper"
	MimetypesCSP         Mimetypes = "application/vnd.commonspace"
	MimetypesSAF         Mimetypes = "application/vnd.yamaha.smaf-audio"
	MimetypesMMF         Mimetypes = "application/vnd.smaf"
	MimetypesSPF         Mimetypes = "application/vnd.yamaha.smaf-phrase"
	MimetypesTEACHER     Mimetypes = "application/vnd.smart.teacher"
	MimetypesSVD         Mimetypes = "application/vnd.svd"
	MimetypesRQ          Mimetypes = "application/sparql-query"
	MimetypesSRX         Mimetypes = "application/sparql-results+xml"
	MimetypesGRAM        Mimetypes = "application/srgs"
	MimetypesGRXML       Mimetypes = "application/srgs+xml"
	MimetypesSSML        Mimetypes = "application/ssml+xml"
	MimetypesSKP         Mimetypes = "application/vnd.koan"
	MimetypesSGML        Mimetypes = "text/sgml"
	MimetypesSDC         Mimetypes = "application/vnd.stardivision.calc"
	MimetypesSDA         Mimetypes = "application/vnd.stardivision.draw"
	MimetypesSDD         Mimetypes = "application/vnd.stardivision.impress"
	MimetypesSMF         Mimetypes = "application/vnd.stardivision.math"
	MimetypesSDW         Mimetypes = "application/vnd.stardivision.writer"
	MimetypesSGL         Mimetypes = "application/vnd.stardivision.writer-global"
	MimetypesSM          Mimetypes = "application/vnd.stepmania.stepchart"
	MimetypesSIT         Mimetypes = "application/x-stuffit"
	MimetypesSITX        Mimetypes = "application/x-stuffitx"
	MimetypesSDKM        Mimetypes = "application/vnd.solent.sdkm+xml"
	MimetypesXO          Mimetypes = "application/vnd.olpc-sugar"
	MimetypesAU          Mimetypes = "audio/basic"
	MimetypesWQD         Mimetypes = "application/vnd.wqd"
	MimetypesSIS         Mimetypes = "application/vnd.symbian.install"
	MimetypesSMI         Mimetypes = "application/smil+xml"
	MimetypesXSM         Mimetypes = "application/vnd.syncml+xml"
	MimetypesBDM         Mimetypes = "application/vnd.syncml.dm+wbxml"
	MimetypesXDM         Mimetypes = "application/vnd.syncml.dm+xml"
	MimetypesSV4CPIO     Mimetypes = "application/x-sv4cpio"
	MimetypesSV4CRC      Mimetypes = "application/x-sv4crc"
	MimetypesSBML        Mimetypes = "application/sbml+xml"
	MimetypesTSV         Mimetypes = "text/tab-separated-values"
	MimetypesTIFF        Mimetypes = "image/tiff"
	MimetypesTAO         Mimetypes = "application/vnd.tao.intent-module-archive"
	MimetypesTAR         Mimetypes = "application/x-tar"
	MimetypesTCL         Mimetypes = "application/x-tcl"
	MimetypesTEX         Mimetypes = "application/x-tex"
	MimetypesTFM         Mimetypes = "application/x-tex-tfm"
	MimetypesTEI         Mimetypes = "application/tei+xml"
	MimetypesTXT         Mimetypes = "text/plain"
	MimetypesDXP         Mimetypes = "application/vnd.spotfire.dxp"
	MimetypesSFS         Mimetypes = "application/vnd.spotfire.sfs"
	MimetypesTSD         Mimetypes = "application/timestamped-data"
	MimetypesTPT         Mimetypes = "application/vnd.trid.tpt"
	MimetypesMXS         Mimetypes = "application/vnd.triscape.mxs"
	MimetypesT           Mimetypes = "text/troff"
	MimetypesTRA         Mimetypes = "application/vnd.trueapp"
	MimetypesTTF         Mimetypes = "application/x-font-ttf"
	MimetypesTTL         Mimetypes = "text/turtle"
	MimetypesUMJ         Mimetypes = "application/vnd.umajin"
	MimetypesUOML        Mimetypes = "application/vnd.uoml+xml"
	MimetypesUNITYWEB    Mimetypes = "application/vnd.unity"
	MimetypesUFD         Mimetypes = "application/vnd.ufdl"
	MimetypesURI         Mimetypes = "text/uri-list"
	MimetypesUTZ         Mimetypes = "application/vnd.uiq.theme"
	MimetypesUSTAR       Mimetypes = "application/x-ustar"
	MimetypesUU          Mimetypes = "text/x-uuencode"
	MimetypesVCS         Mimetypes = "text/x-vcalendar"
	MimetypesVCF         Mimetypes = "text/x-vcard"
	MimetypesVCD         Mimetypes = "application/x-cdlink"
	MimetypesVSF         Mimetypes = "application/vnd.vsf"
	MimetypesWRL         Mimetypes = "model/vrml"
	MimetypesVCX         Mimetypes = "application/vnd.vcx"
	MimetypesMTS         Mimetypes = "model/vnd.mts"
	MimetypesVTU         Mimetypes = "model/vnd.vtu"
	MimetypesVIS         Mimetypes = "application/vnd.visionary"
	MimetypesVIV         Mimetypes = "video/vnd.vivo"
	MimetypesCCXML       Mimetypes = "application/ccxml+xml,"
	MimetypesVXML        Mimetypes = "application/voicexml+xml"
	MimetypesSRC         Mimetypes = "application/x-wais-source"
	MimetypesWBXML       Mimetypes = "application/vnd.wap.wbxml"
	MimetypesWBMP        Mimetypes = "image/vnd.wap.wbmp"
	MimetypesWAV         Mimetypes = "audio/x-wav"
	MimetypesDAVMOUNT    Mimetypes = "application/davmount+xml"
	MimetypesWOFF        Mimetypes = "application/x-font-woff"
	MimetypesWSPOLICY    Mimetypes = "application/wspolicy+xml"
	MimetypesWEBP        Mimetypes = "image/webp"
	MimetypesWTB         Mimetypes = "application/vnd.webturbo"
	MimetypesWGT         Mimetypes = "application/widget"
	MimetypesHLP         Mimetypes = "application/winhlp"
	MimetypesWML         Mimetypes = "text/vnd.wap.wml"
	MimetypesWMLS        Mimetypes = "text/vnd.wap.wmlscript"
	MimetypesWMLSC       Mimetypes = "application/vnd.wap.wmlscriptc"
	MimetypesWPD         Mimetypes = "application/vnd.wordperfect"
	MimetypesSTF         Mimetypes = "application/vnd.wt.stf"
	MimetypesWSDL        Mimetypes = "application/wsdl+xml"
	MimetypesXBM         Mimetypes = "image/x-xbitmap"
	MimetypesXPM         Mimetypes = "image/x-xpixmap"
	MimetypesXWD         Mimetypes = "image/x-xwindowdump"
	MimetypesDER         Mimetypes = "application/x-x509-ca-cert"
	MimetypesFIG         Mimetypes = "application/x-xfig"
	MimetypesXHTML       Mimetypes = "application/xhtml+xml"
	MimetypesXML         Mimetypes = "application/xml"
	MimetypesXDF         Mimetypes = "application/xcap-diff+xml"
	MimetypesXENC        Mimetypes = "application/xenc+xml"
	MimetypesXER         Mimetypes = "application/patch-ops-error+xml"
	MimetypesRL          Mimetypes = "application/resource-lists+xml"
	MimetypesRS          Mimetypes = "application/rls-services+xml"
	MimetypesRLD         Mimetypes = "application/resource-lists-diff+xml"
	MimetypesXSLT        Mimetypes = "application/xslt+xml"
	MimetypesXOP         Mimetypes = "application/xop+xml"
	MimetypesXPI         Mimetypes = "application/x-xpinstall"
	MimetypesXSPF        Mimetypes = "application/xspf+xml"
	MimetypesXUL         Mimetypes = "application/vnd.mozilla.xul+xml"
	MimetypesXYZ         Mimetypes = "chemical/x-xyz"
	MimetypesYAML        Mimetypes = "text/yaml"
	MimetypesYANG        Mimetypes = "application/yang"
	MimetypesYIN         Mimetypes = "application/yin+xml"
	MimetypesZIR         Mimetypes = "application/vnd.zul"
	MimetypesZIP         Mimetypes = "application/zip"
	MimetypesZMM         Mimetypes = "application/vnd.handheld-entertainment+xml"
	MimetypesZAZ         Mimetypes = "application/vnd.zzazz.deck+xml"
)

func AllMimetypes() []Mimetypes {
	return []Mimetypes{
		MimetypesX3D,
		Mimetypes3GP,
		Mimetypes3G2,
		MimetypesMSEQ,
		MimetypesPWN,
		MimetypesPLB,
		MimetypesPSB,
		MimetypesPVB,
		MimetypesTCAP,
		Mimetypes7Z,
		MimetypesABW,
		MimetypesACE,
		MimetypesACC,
		MimetypesACU,
		MimetypesATC,
		MimetypesADP,
		MimetypesAAB,
		MimetypesAAM,
		MimetypesAAS,
		MimetypesAIR,
		MimetypesSWF,
		MimetypesFXP,
		MimetypesPDF,
		MimetypesPPD,
		MimetypesDIR,
		MimetypesXDP,
		MimetypesXFDF,
		MimetypesAAC,
		MimetypesAHEAD,
		MimetypesAZF,
		MimetypesAZS,
		MimetypesAZW,
		MimetypesAMI,
		MimetypesANI,
		MimetypesAPK,
		MimetypesCII,
		MimetypesFTI,
		MimetypesATX,
		MimetypesDMG,
		MimetypesMPKG,
		MimetypesAW,
		MimetypesLES,
		MimetypesARC,
		MimetypesSWI,
		MimetypesS,
		MimetypesATOMCAT,
		MimetypesATOMSVC,
		MimetypesATOMXML,
		MimetypesAC,
		MimetypesAIF,
		MimetypesAVI,
		MimetypesAEP,
		MimetypesDXF,
		MimetypesDWF,
		MimetypesAVIF,
		MimetypesPAR,
		MimetypesBCPIO,
		MimetypesBIN,
		MimetypesBMP,
		MimetypesTORRENT,
		MimetypesCOD,
		MimetypesMPM,
		MimetypesBMI,
		MimetypesSH,
		MimetypesBTIF,
		MimetypesREP,
		MimetypesBZ,
		MimetypesBZ2,
		MimetypesCSH,
		MimetypesC,
		MimetypesCDXML,
		MimetypesCSS,
		MimetypesCDA,
		MimetypesCDX,
		MimetypesCML,
		MimetypesCSML,
		MimetypesCDBCMSG,
		MimetypesCLA,
		MimetypesC4G,
		MimetypesSUB,
		MimetypesCDMIA,
		MimetypesCDMIC,
		MimetypesCDMID,
		MimetypesCDMIO,
		MimetypesCDMIQ,
		MimetypesC11AMC,
		MimetypesC11AMZ,
		MimetypesRAS,
		MimetypesDAE,
		MimetypesCSV,
		MimetypesCPT,
		MimetypesWMLC,
		MimetypesCGM,
		MimetypesICE,
		MimetypesCMX,
		MimetypesXAR,
		MimetypesCMC,
		MimetypesCPIO,
		MimetypesCLKX,
		MimetypesCLKK,
		MimetypesCLKP,
		MimetypesCLKT,
		MimetypesCLKW,
		MimetypesWBS,
		MimetypesCRYPTONOTE,
		MimetypesCIF,
		MimetypesCMDF,
		MimetypesCU,
		MimetypesCWW,
		MimetypesCURL,
		MimetypesDCURL,
		MimetypesMCURL,
		MimetypesSCURL,
		MimetypesCAR,
		MimetypesPCURL,
		MimetypesCMP,
		MimetypesDSSC,
		MimetypesXDSSC,
		MimetypesDEB,
		MimetypesUVA,
		MimetypesUVI,
		MimetypesUVH,
		MimetypesUVM,
		MimetypesUVU,
		MimetypesUVP,
		MimetypesUVS,
		MimetypesUVV,
		MimetypesDVI,
		MimetypesSEED,
		MimetypesDTB,
		MimetypesRES,
		MimetypesAIT,
		MimetypesSVC,
		MimetypesEOL,
		MimetypesDJVU,
		MimetypesDTD,
		MimetypesMLP,
		MimetypesWAD,
		MimetypesDPG,
		MimetypesDRA,
		MimetypesDFAC,
		MimetypesDTS,
		MimetypesDTSHD,
		MimetypesDWG,
		MimetypesGEO,
		MimetypesES,
		MimetypesMAG,
		MimetypesMMR,
		MimetypesRLC,
		MimetypesEXI,
		MimetypesMGZ,
		MimetypesEPUB,
		MimetypesEML,
		MimetypesNML,
		MimetypesXPR,
		MimetypesXIF,
		MimetypesXFDL,
		MimetypesEMMA,
		MimetypesEZ2,
		MimetypesEZ3,
		MimetypesFST,
		MimetypesFVT,
		MimetypesFBS,
		MimetypesFE_LAUNCH,
		MimetypesF4V,
		MimetypesFLV,
		MimetypesFPX,
		MimetypesNPX,
		MimetypesFLX,
		MimetypesFLI,
		MimetypesFTC,
		MimetypesFDF,
		MimetypesF,
		MimetypesMIF,
		MimetypesFM,
		MimetypesFH,
		MimetypesFSC,
		MimetypesFNC,
		MimetypesLTF,
		MimetypesDDD,
		MimetypesXDW,
		MimetypesXBD,
		MimetypesOAS,
		MimetypesOA2,
		MimetypesOA3,
		MimetypesFG5,
		MimetypesBH2,
		MimetypesSPL,
		MimetypesFZS,
		MimetypesG3,
		MimetypesGMX,
		MimetypesGTW,
		MimetypesTXD,
		MimetypesGGB,
		MimetypesGGT,
		MimetypesGDL,
		MimetypesGEX,
		MimetypesGXT,
		MimetypesG2W,
		MimetypesG3W,
		MimetypesGSF,
		MimetypesBDF,
		MimetypesGTAR,
		MimetypesTEXINFO,
		MimetypesGNUMERIC,
		MimetypesKML,
		MimetypesKMZ,
		MimetypesGPX,
		MimetypesGQF,
		MimetypesGIF,
		MimetypesGV,
		MimetypesGAC,
		MimetypesGHF,
		MimetypesGIM,
		MimetypesGRV,
		MimetypesGTM,
		MimetypesTPL,
		MimetypesVCG,
		MimetypesGZ,
		MimetypesH261,
		MimetypesH263,
		MimetypesH264,
		MimetypesHPID,
		MimetypesHPS,
		MimetypesHDF,
		MimetypesRIP,
		MimetypesHBCI,
		MimetypesJLT,
		MimetypesPCL,
		MimetypesHPGL,
		MimetypesHVS,
		MimetypesHVD,
		MimetypesHVP,
		MimetypesSFHDSTX,
		MimetypesSTK,
		MimetypesHAL,
		MimetypesHTML,
		MimetypesIRM,
		MimetypesSC,
		MimetypesICS,
		MimetypesICC,
		MimetypesICO,
		MimetypesIGL,
		MimetypesIEF,
		MimetypesIVP,
		MimetypesIVU,
		MimetypesRIF,
		Mimetypes3DML,
		MimetypesSPOT,
		MimetypesIGS,
		MimetypesI2G,
		MimetypesCDY,
		MimetypesXPW,
		MimetypesFCS,
		MimetypesIPFIX,
		MimetypesCER,
		MimetypesPKI,
		MimetypesCRL,
		MimetypesPKIPATH,
		MimetypesIGM,
		MimetypesRCPROFILE,
		MimetypesIRP,
		MimetypesJAD,
		MimetypesJAR,
		MimetypesCLASS,
		MimetypesJNLP,
		MimetypesSER,
		MimetypesJAVA,
		MimetypesJS,
		MimetypesTXTJS,
		MimetypesMJS,
		MimetypesJSON,
		MimetypesJODA,
		MimetypesJPM,
		MimetypesIMGJPEG,
		MimetypesCITRIXJPEG,
		MimetypesPJPEG,
		MimetypesJPGV,
		MimetypesJSONLD,
		MimetypesKTZ,
		MimetypesMMD,
		MimetypesKARBON,
		MimetypesCHRT,
		MimetypesKFO,
		MimetypesFLW,
		MimetypesKON,
		MimetypesKPR,
		MimetypesKSP,
		MimetypesKWD,
		MimetypesHTKE,
		MimetypesKIA,
		MimetypesKNE,
		MimetypesSSE,
		MimetypesLASXML,
		MimetypesLATEX,
		MimetypesLBD,
		MimetypesLBE,
		MimetypesJAM,
		Mimetypes123,
		MimetypesAPR,
		MimetypesPRE,
		MimetypesNSF,
		MimetypesORG,
		MimetypesSCM,
		MimetypesLWP,
		MimetypesLVP,
		MimetypesM3U,
		MimetypesM4V,
		MimetypesHQX,
		MimetypesPORTPKG,
		MimetypesMGP,
		MimetypesMRC,
		MimetypesMRCX,
		MimetypesMXF,
		MimetypesNBP,
		MimetypesMA,
		MimetypesMATHML,
		MimetypesMBOX,
		MimetypesMC1,
		MimetypesMSCML,
		MimetypesCDKEY,
		MimetypesMWF,
		MimetypesMFM,
		MimetypesMSH,
		MimetypesMADS,
		MimetypesMETS,
		MimetypesMODS,
		MimetypesMETA4,
		MimetypesMCD,
		MimetypesFLO,
		MimetypesIGX,
		MimetypesES3,
		MimetypesMDB,
		MimetypesASF,
		MimetypesEXE,
		MimetypesCIL,
		MimetypesCAB,
		MimetypesIMS,
		MimetypesAPPLICATION,
		MimetypesCLP,
		MimetypesMDI,
		MimetypesEOT,
		MimetypesXLS,
		MimetypesXLAM,
		MimetypesXLSB,
		MimetypesXLTM,
		MimetypesXLSM,
		MimetypesCHM,
		MimetypesCRD,
		MimetypesLRM,
		MimetypesMVB,
		MimetypesMNY,
		MimetypesPPTX,
		MimetypesSLDX,
		MimetypesPPSX,
		MimetypesPOTX,
		MimetypesXLSX,
		MimetypesXLTX,
		MimetypesDOCX,
		MimetypesDOTX,
		MimetypesOBD,
		MimetypesTHMX,
		MimetypesONETOC,
		MimetypesPYA,
		MimetypesPYV,
		MimetypesPPT,
		MimetypesPPAM,
		MimetypesSLDM,
		MimetypesPPTM,
		MimetypesPPSM,
		MimetypesPOTM,
		MimetypesMPP,
		MimetypesPUB,
		MimetypesSCD,
		MimetypesXAP,
		MimetypesSTL,
		MimetypesCAT,
		MimetypesVSD,
		MimetypesVSDX,
		MimetypesWM,
		MimetypesWMA,
		MimetypesWAX,
		MimetypesWMX,
		MimetypesWMD,
		MimetypesWPL,
		MimetypesWMZ,
		MimetypesWMV,
		MimetypesWVX,
		MimetypesWMF,
		MimetypesTRM,
		MimetypesDOC,
		MimetypesDOCM,
		MimetypesDOTM,
		MimetypesWRI,
		MimetypesWPS,
		MimetypesXBAP,
		MimetypesXPS,
		MimetypesMIDI,
		MimetypesMID,
		MimetypesMPY,
		MimetypesAFP,
		MimetypesRMS,
		MimetypesTMO,
		MimetypesPRC,
		MimetypesMBK,
		MimetypesDIS,
		MimetypesPLC,
		MimetypesMQY,
		MimetypesMSL,
		MimetypesTXF,
		MimetypesDAF,
		MimetypesFLY,
		MimetypesMPC,
		MimetypesMPN,
		MimetypesMJ2,
		MimetypesMPGA,
		MimetypesTS,
		MimetypesMXU,
		MimetypesMPEG,
		MimetypesM21,
		MimetypesMP4A,
		MimetypesVDMP4,
		MimetypesMP4,
		MimetypesM3U8,
		MimetypesMUS,
		MimetypesMSTY,
		MimetypesMXML,
		MimetypesNGDAT,
		MimetypesNGAGE,
		MimetypesNCX,
		MimetypesNC,
		MimetypesNLU,
		MimetypesDNA,
		MimetypesNND,
		MimetypesNNS,
		MimetypesNNW,
		MimetypesRPST,
		MimetypesRPSS,
		MimetypesN3,
		MimetypesEDM,
		MimetypesEDX,
		MimetypesEXT,
		MimetypesGPH,
		MimetypesECELP4800,
		MimetypesECELP7470,
		MimetypesECELP9600,
		MimetypesODA,
		MimetypesOGX,
		MimetypesOGA,
		MimetypesOGV,
		MimetypesDD2,
		MimetypesOTH,
		MimetypesOPF,
		MimetypesQBO,
		MimetypesOXT,
		MimetypesOSF,
		MimetypesWEBA,
		MimetypesWEBM,
		MimetypesODC,
		MimetypesOTC,
		MimetypesODB,
		MimetypesODF,
		MimetypesODFT,
		MimetypesODG,
		MimetypesOTG,
		MimetypesODI,
		MimetypesOTI,
		MimetypesODP,
		MimetypesOTP,
		MimetypesODS,
		MimetypesOTS,
		MimetypesODT,
		MimetypesODM,
		MimetypesOTT,
		MimetypesKTX,
		MimetypesSXC,
		MimetypesSTC,
		MimetypesSXD,
		MimetypesSTD,
		MimetypesSXI,
		MimetypesSTI,
		MimetypesSXM,
		MimetypesSXW,
		MimetypesSXG,
		MimetypesSTW,
		MimetypesOTF,
		MimetypesOPUS,
		MimetypesOSFPVG,
		MimetypesDP,
		MimetypesPDB,
		MimetypesP,
		MimetypesPAW,
		MimetypesPCLXL,
		MimetypesEFIF,
		MimetypesPCX,
		MimetypesPSD,
		MimetypesPRF,
		MimetypesPIC,
		MimetypesCHAT,
		MimetypesP10,
		MimetypesP12,
		MimetypesP7M,
		MimetypesP7S,
		MimetypesP7R,
		MimetypesP7B,
		MimetypesP8,
		MimetypesPLF,
		MimetypesPNM,
		MimetypesPBM,
		MimetypesPCF,
		MimetypesPFR,
		MimetypesPGN,
		MimetypesPGM,
		MimetypesIMGPNG,
		MimetypesPNG,
		MimetypesXPNG,
		MimetypesPPM,
		MimetypesPSKCXML,
		MimetypesPML,
		MimetypesAI,
		MimetypesPFA,
		MimetypesPBD,
		MimetypesCRYPGP,
		MimetypesSIGPGP,
		MimetypesBOX,
		MimetypesPTID,
		MimetypesPLS,
		MimetypesSTR,
		MimetypesEI6,
		MimetypesDSC,
		MimetypesPSF,
		MimetypesQPS,
		MimetypesWG,
		MimetypesQXD,
		MimetypesESF,
		MimetypesMSF,
		MimetypesSSF,
		MimetypesQAM,
		MimetypesQFX,
		MimetypesQT,
		MimetypesRAR,
		MimetypesRAM,
		MimetypesRMP,
		MimetypesRSD,
		MimetypesRM,
		MimetypesBED,
		MimetypesMXL,
		MimetypesMUSICXML,
		MimetypesRNC,
		MimetypesRDZ,
		MimetypesRDF,
		MimetypesRP9,
		MimetypesJISP,
		MimetypesRTF,
		MimetypesRTX,
		MimetypesLINK66,
		MimetypesRSSXML,
		MimetypesSHF,
		MimetypesST,
		MimetypesSVG,
		MimetypesSUS,
		MimetypesSRU,
		MimetypesSETPAY,
		MimetypesSETREG,
		MimetypesSEMA,
		MimetypesSEMD,
		MimetypesSEMF,
		MimetypesSEE,
		MimetypesSNF,
		MimetypesSPQ,
		MimetypesSPP,
		MimetypesSCQ,
		MimetypesSCS,
		MimetypesSDP,
		MimetypesETX,
		MimetypesMOVIE,
		MimetypesIFM,
		MimetypesITP,
		MimetypesIIF,
		MimetypesIPK,
		MimetypesTFI,
		MimetypesSHAR,
		MimetypesRGB,
		MimetypesSLT,
		MimetypesASO,
		MimetypesIMP,
		MimetypesTWD,
		MimetypesCSP,
		MimetypesSAF,
		MimetypesMMF,
		MimetypesSPF,
		MimetypesTEACHER,
		MimetypesSVD,
		MimetypesRQ,
		MimetypesSRX,
		MimetypesGRAM,
		MimetypesGRXML,
		MimetypesSSML,
		MimetypesSKP,
		MimetypesSGML,
		MimetypesSDC,
		MimetypesSDA,
		MimetypesSDD,
		MimetypesSMF,
		MimetypesSDW,
		MimetypesSGL,
		MimetypesSM,
		MimetypesSIT,
		MimetypesSITX,
		MimetypesSDKM,
		MimetypesXO,
		MimetypesAU,
		MimetypesWQD,
		MimetypesSIS,
		MimetypesSMI,
		MimetypesXSM,
		MimetypesBDM,
		MimetypesXDM,
		MimetypesSV4CPIO,
		MimetypesSV4CRC,
		MimetypesSBML,
		MimetypesTSV,
		MimetypesTIFF,
		MimetypesTAO,
		MimetypesTAR,
		MimetypesTCL,
		MimetypesTEX,
		MimetypesTFM,
		MimetypesTEI,
		MimetypesTXT,
		MimetypesDXP,
		MimetypesSFS,
		MimetypesTSD,
		MimetypesTPT,
		MimetypesMXS,
		MimetypesT,
		MimetypesTRA,
		MimetypesTTF,
		MimetypesTTL,
		MimetypesUMJ,
		MimetypesUOML,
		MimetypesUNITYWEB,
		MimetypesUFD,
		MimetypesURI,
		MimetypesUTZ,
		MimetypesUSTAR,
		MimetypesUU,
		MimetypesVCS,
		MimetypesVCF,
		MimetypesVCD,
		MimetypesVSF,
		MimetypesWRL,
		MimetypesVCX,
		MimetypesMTS,
		MimetypesVTU,
		MimetypesVIS,
		MimetypesVIV,
		MimetypesCCXML,
		MimetypesVXML,
		MimetypesSRC,
		MimetypesWBXML,
		MimetypesWBMP,
		MimetypesWAV,
		MimetypesDAVMOUNT,
		MimetypesWOFF,
		MimetypesWSPOLICY,
		MimetypesWEBP,
		MimetypesWTB,
		MimetypesWGT,
		MimetypesHLP,
		MimetypesWML,
		MimetypesWMLS,
		MimetypesWMLSC,
		MimetypesWPD,
		MimetypesSTF,
		MimetypesWSDL,
		MimetypesXBM,
		MimetypesXPM,
		MimetypesXWD,
		MimetypesDER,
		MimetypesFIG,
		MimetypesXHTML,
		MimetypesXML,
		MimetypesXDF,
		MimetypesXENC,
		MimetypesXER,
		MimetypesRL,
		MimetypesRS,
		MimetypesRLD,
		MimetypesXSLT,
		MimetypesXOP,
		MimetypesXPI,
		MimetypesXSPF,
		MimetypesXUL,
		MimetypesXYZ,
		MimetypesYAML,
		MimetypesYANG,
		MimetypesYIN,
		MimetypesZIR,
		MimetypesZIP,
		MimetypesZMM,
		MimetypesZAZ,
	}
}

func FindMimetypes(filter string) []Mimetypes {
	ret := make([]Mimetypes, 0)
	for _, at := range AllMimetypes() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au Mimetypes) ToString() {
	fmt.Println(au.String())
}
func (au Mimetypes) String() string {
	switch au {
	case MimetypesX3D:
		return "3D Crossword Plugin"
	case Mimetypes3GP:
		return "3GP"
	case Mimetypes3G2:
		return "3GP2"
	case MimetypesMSEQ:
		return "3GPP MSEQ File"
	case MimetypesPWN:
		return "3M Post It Notes"
	case MimetypesPLB:
		return "3rd Generation Partnership Project - Pic Large"
	case MimetypesPSB:
		return "3rd Generation Partnership Project - Pic Small"
	case MimetypesPVB:
		return "3rd Generation Partnership Project - Pic Var"
	case MimetypesTCAP:
		return "3rd Generation Partnership Project - Transaction Capabilities Application Part"
	case Mimetypes7Z:
		return "7-Zip"
	case MimetypesABW:
		return "AbiWord"
	case MimetypesACE:
		return "Ace Archive"
	case MimetypesACC:
		return "Active Content Compression"
	case MimetypesACU:
		return "ACU Cobol"
	case MimetypesATC:
		return "ACU Cobol"
	case MimetypesADP:
		return "Adaptive differential pulse-code modulation"
	case MimetypesAAB:
		return "Adobe (Macropedia) Authorware - Binary File"
	case MimetypesAAM:
		return "Adobe (Macropedia) Authorware - Map"
	case MimetypesAAS:
		return "Adobe (Macropedia) Authorware - Segment File"
	case MimetypesAIR:
		return "Adobe AIR Application"
	case MimetypesSWF:
		return "Adobe Flash"
	case MimetypesFXP:
		return "Adobe Flex Project"
	case MimetypesPDF:
		return "Adobe Portable Document Format"
	case MimetypesPPD:
		return "Adobe PostScript Printer Description File Format"
	case MimetypesDIR:
		return "Adobe Shockwave Player"
	case MimetypesXDP:
		return "Adobe XML Data Package"
	case MimetypesXFDF:
		return "Adobe XML Forms Data Format"
	case MimetypesAAC:
		return "Advanced Audio Coding (AAC)"
	case MimetypesAHEAD:
		return "Ahead AIR Application"
	case MimetypesAZF:
		return "AirZip FileSECURE"
	case MimetypesAZS:
		return "AirZip FileSECURE"
	case MimetypesAZW:
		return "Amazon Kindle eBook format"
	case MimetypesAMI:
		return "AmigaDE"
	case MimetypesANI:
		return "Andrew Toolkit"
	case MimetypesAPK:
		return "Android Package Archive"
	case MimetypesCII:
		return "ANSER-WEB Terminal Client - Certificate Issue"
	case MimetypesFTI:
		return "ANSER-WEB Terminal Client - Web Funds Transfer"
	case MimetypesATX:
		return "Antix Game Player"
	case MimetypesDMG:
		return "Apple Disk Image"
	case MimetypesMPKG:
		return "Apple Installer Package"
	case MimetypesAW:
		return "Applixware"
	case MimetypesLES:
		return "Archipelago Lesson Player"
	case MimetypesARC:
		return "Archive document - Multiple Fils Embedded"
	case MimetypesSWI:
		return "Arista Networks Software Image"
	case MimetypesS:
		return "Assembler Source File"
	case MimetypesATOMCAT:
		return "Atom Publishing Protocol"
	case MimetypesATOMSVC:
		return "Atom Publishing Protocol Service Document"
	case MimetypesATOMXML:
		return "Atom Syndication Format"
	case MimetypesAC:
		return "Attribute Certificate"
	case MimetypesAIF:
		return "Audio Interchange File Format"
	case MimetypesAVI:
		return "Audio Video Interleave (AVI)"
	case MimetypesAEP:
		return "Audiograph"
	case MimetypesDXF:
		return "AutoCAD DXF"
	case MimetypesDWF:
		return "Autodesk Design Web Format (DWF)"
	case MimetypesAVIF:
		return "AV1 Image File"
	case MimetypesPAR:
		return "BAS Partitur Format"
	case MimetypesBCPIO:
		return "Binary CPIO Archive"
	case MimetypesBIN:
		return "Binary Data"
	case MimetypesBMP:
		return "Bitmap Image File"
	case MimetypesTORRENT:
		return "BitTorrent"
	case MimetypesCOD:
		return "Blackberry COD File"
	case MimetypesMPM:
		return "Blueice Research Multipass"
	case MimetypesBMI:
		return "BMI Drawing Data Interchange"
	case MimetypesSH:
		return "Bourne Shell Script"
	case MimetypesBTIF:
		return "BTIF"
	case MimetypesREP:
		return "BusinessObjects"
	case MimetypesBZ:
		return "Bzip Archive"
	case MimetypesBZ2:
		return "Bzip2 Archive"
	case MimetypesCSH:
		return "C Shell Script"
	case MimetypesC:
		return "C Source File"
	case MimetypesCDXML:
		return "CambridgeSoft Chem Draw"
	case MimetypesCSS:
		return "Cascading Style Sheets (CSS)"
	case MimetypesCDA:
		return "CD Audio"
	case MimetypesCDX:
		return "ChemDraw eXchange file"
	case MimetypesCML:
		return "Chemical Markup Language"
	case MimetypesCSML:
		return "Chemical Style Markup Language"
	case MimetypesCDBCMSG:
		return "CIM Database"
	case MimetypesCLA:
		return "Claymore Data Files"
	case MimetypesC4G:
		return "Clonk Game"
	case MimetypesSUB:
		return "Close Captioning - Subtitle"
	case MimetypesCDMIA:
		return "Cloud Data Management Interface (CDMI) - Capability"
	case MimetypesCDMIC:
		return "Cloud Data Management Interface (CDMI) - Contaimer"
	case MimetypesCDMID:
		return "Cloud Data Management Interface (CDMI) - Domain"
	case MimetypesCDMIO:
		return "Cloud Data Management Interface (CDMI) - Object"
	case MimetypesCDMIQ:
		return "Cloud Data Management Interface (CDMI) - Queue"
	case MimetypesC11AMC:
		return "ClueTrust CartoMobile - Config"
	case MimetypesC11AMZ:
		return "ClueTrust CartoMobile - Config Package"
	case MimetypesRAS:
		return "CMU Image"
	case MimetypesDAE:
		return "COLLADA"
	case MimetypesCSV:
		return "Comma-Seperated Values"
	case MimetypesCPT:
		return "Compact Pro"
	case MimetypesWMLC:
		return "Compiled Wireless Markup Language (WMLC)"
	case MimetypesCGM:
		return "Computer Graphics Metafile"
	case MimetypesICE:
		return "CoolTalk"
	case MimetypesCMX:
		return "Corel Metafile Exchange (CMX)"
	case MimetypesXAR:
		return "CorelXARA"
	case MimetypesCMC:
		return "CosmoCaller"
	case MimetypesCPIO:
		return "CPIO Archive"
	case MimetypesCLKX:
		return "CrickSoftware - Clicker"
	case MimetypesCLKK:
		return "CrickSoftware - Clicker - Keyboard"
	case MimetypesCLKP:
		return "CrickSoftware - Clicker - Palette"
	case MimetypesCLKT:
		return "CrickSoftware - Clicker - Template"
	case MimetypesCLKW:
		return "CrickSoftware - Clicker - Wordbank"
	case MimetypesWBS:
		return "Critical Tools - PERT Chart EXPERT"
	case MimetypesCRYPTONOTE:
		return "CryptoNote"
	case MimetypesCIF:
		return "Crystallographic Interchange Format"
	case MimetypesCMDF:
		return "CrystalMaker Data Format"
	case MimetypesCU:
		return "CU-SeeMe"
	case MimetypesCWW:
		return "CU-Writer"
	case MimetypesCURL:
		return "Curl - Applet"
	case MimetypesDCURL:
		return "Curl - Detached Applet"
	case MimetypesMCURL:
		return "Curl - Manifest File"
	case MimetypesSCURL:
		return "Curl - Source Code"
	case MimetypesCAR:
		return "CURL Applet"
	case MimetypesPCURL:
		return "CURL Applet"
	case MimetypesCMP:
		return "CustomMenu"
	case MimetypesDSSC:
		return "Data Structure for the Security Suitability of Cryptographic Algorithms"
	case MimetypesXDSSC:
		return "Data Structure for the Security Suitability of Cryptographic Algorithms"
	case MimetypesDEB:
		return "Debian Package"
	case MimetypesUVA:
		return "DECE Audio"
	case MimetypesUVI:
		return "DECE Graphic"
	case MimetypesUVH:
		return "DECE High Definition Video"
	case MimetypesUVM:
		return "DECE Mobile Video"
	case MimetypesUVU:
		return "DECE MP4"
	case MimetypesUVP:
		return "DECE PD Video"
	case MimetypesUVS:
		return "DECE SD Video"
	case MimetypesUVV:
		return "DECE Video"
	case MimetypesDVI:
		return "Device Independent File Format (DVI)"
	case MimetypesSEED:
		return "Digital Siesmograph Networks - SEED Datafiles"
	case MimetypesDTB:
		return "Digital Talking Book"
	case MimetypesRES:
		return "Digital Talking Book - Resource File"
	case MimetypesAIT:
		return "Digital Video Broadcasting"
	case MimetypesSVC:
		return "Digital Video Broadcasting"
	case MimetypesEOL:
		return "Digital Winds Music"
	case MimetypesDJVU:
		return "DjVu"
	case MimetypesDTD:
		return "Document Type Definition"
	case MimetypesMLP:
		return "Dolby Meridian Lossless Packing"
	case MimetypesWAD:
		return "Doom Video Game"
	case MimetypesDPG:
		return "DPGraph"
	case MimetypesDRA:
		return "DRA Audio"
	case MimetypesDFAC:
		return "DreamFactory"
	case MimetypesDTS:
		return "DTS Audio"
	case MimetypesDTSHD:
		return "DTS High Definition Audio"
	case MimetypesDWG:
		return "DWG Drawing"
	case MimetypesGEO:
		return "DynaGeo"
	case MimetypesES:
		return "ECMAScript"
	case MimetypesMAG:
		return "EcoWin Chart"
	case MimetypesMMR:
		return "EDMICS 2000"
	case MimetypesRLC:
		return "EDMICS 2000"
	case MimetypesEXI:
		return "Efficient XML Interchange"
	case MimetypesMGZ:
		return "EFI Proteus"
	case MimetypesEPUB:
		return "Electronic Publication"
	case MimetypesEML:
		return "Email Message"
	case MimetypesNML:
		return "Enliven Viewer"
	case MimetypesXPR:
		return "Express by Infoseek"
	case MimetypesXIF:
		return "eXtended Image File Format (XIFF)"
	case MimetypesXFDL:
		return "Extensible Forms Description Language"
	case MimetypesEMMA:
		return "Extensible MultiModal Annotation"
	case MimetypesEZ2:
		return "EZPix Secure Photo Album"
	case MimetypesEZ3:
		return "EZPix Secure Photo Album"
	case MimetypesFST:
		return "FAST Search & Transfer ASA"
	case MimetypesFVT:
		return "FAST Search & Transfer ASA"
	case MimetypesFBS:
		return "FastBid Sheet"
	case MimetypesFE_LAUNCH:
		return "FCS Express Layout Link"
	case MimetypesF4V:
		return "Flash Video"
	case MimetypesFLV:
		return "Flash Video"
	case MimetypesFPX:
		return "FlashPix"
	case MimetypesNPX:
		return "FlashPix"
	case MimetypesFLX:
		return "FLEXSTOR"
	case MimetypesFLI:
		return "FLI/FLC Animation Format"
	case MimetypesFTC:
		return "FluxTime Clip"
	case MimetypesFDF:
		return "Forms Data Format"
	case MimetypesF:
		return "Fortran Source File"
	case MimetypesMIF:
		return "FrameMaker Interchange Format"
	case MimetypesFM:
		return "FrameMaker Normal Format"
	case MimetypesFH:
		return "FreeHand MX"
	case MimetypesFSC:
		return "Friendly Software Corporation"
	case MimetypesFNC:
		return "Frogans Player"
	case MimetypesLTF:
		return "Frogans Player"
	case MimetypesDDD:
		return "Fujitsu - Xerox 2D CAD Data"
	case MimetypesXDW:
		return "Fujitsu - Xerox DocuWorks"
	case MimetypesXBD:
		return "Fujitsu - Xerox DocuWorks Binder"
	case MimetypesOAS:
		return "Fujitsu Oasys"
	case MimetypesOA2:
		return "Fujitsu Oasys"
	case MimetypesOA3:
		return "Fujitsu Oasys"
	case MimetypesFG5:
		return "Fujitsu Oasys"
	case MimetypesBH2:
		return "Fujitsu Oasys"
	case MimetypesSPL:
		return "FutureSplash Animator"
	case MimetypesFZS:
		return "FuzzySheet"
	case MimetypesG3:
		return "G3 Fax Image"
	case MimetypesGMX:
		return "GameMaker ActiveX"
	case MimetypesGTW:
		return "Gen-Trix Studio"
	case MimetypesTXD:
		return "Genomatix Tuxedo Framework"
	case MimetypesGGB:
		return "GeoGebra"
	case MimetypesGGT:
		return "GeoGebra"
	case MimetypesGDL:
		return "Geometric Description Language (GDL)"
	case MimetypesGEX:
		return "GeoMetry Explorer"
	case MimetypesGXT:
		return "GEONExT and JSXGraph"
	case MimetypesG2W:
		return "GeoplanW"
	case MimetypesG3W:
		return "GeospacW"
	case MimetypesGSF:
		return "Ghostscript Font"
	case MimetypesBDF:
		return "Glyph Bitmap Distribution Format"
	case MimetypesGTAR:
		return "GNU Tar Files"
	case MimetypesTEXINFO:
		return "GNU Texinfo Document"
	case MimetypesGNUMERIC:
		return "Gnumeric"
	case MimetypesKML:
		return "Google Earth - KML"
	case MimetypesKMZ:
		return "Google Earth - Zipped KML"
	case MimetypesGPX:
		return "GPS eXchange Format"
	case MimetypesGQF:
		return "GrafEq"
	case MimetypesGIF:
		return "Graphics Interchange Format"
	case MimetypesGV:
		return "Graphviz"
	case MimetypesGAC:
		return "Groove - Account"
	case MimetypesGHF:
		return "Groove - Help"
	case MimetypesGIM:
		return "Groove - Identity Message"
	case MimetypesGRV:
		return "Groove - Injector"
	case MimetypesGTM:
		return "Groove - Tool Message"
	case MimetypesTPL:
		return "Groove - Tool Template"
	case MimetypesVCG:
		return "Groove - Vcard"
	case MimetypesGZ:
		return "GZip"
	case MimetypesH261:
		return "H.261"
	case MimetypesH263:
		return "H.263"
	case MimetypesH264:
		return "H.264"
	case MimetypesHPID:
		return "Hewlett Packard Instant Delivery"
	case MimetypesHPS:
		return "Hewlett-Packard's WebPrintSmart"
	case MimetypesHDF:
		return "Hierarchical Data Format"
	case MimetypesRIP:
		return "Hit'n'Mix"
	case MimetypesHBCI:
		return "Homebanking Computer Interface (HBCI)"
	case MimetypesJLT:
		return "HP Indigo Digital Press - Job Layout Languate"
	case MimetypesPCL:
		return "HP Printer Command Language"
	case MimetypesHPGL:
		return "HP-GL/2 and HP RTL"
	case MimetypesHVS:
		return "HV Script"
	case MimetypesHVD:
		return "HV Voice Dictionary"
	case MimetypesHVP:
		return "HV Voice Parameter"
	case MimetypesSFHDSTX:
		return "Hydrostatix Master Suite"
	case MimetypesSTK:
		return "Hyperstudio"
	case MimetypesHAL:
		return "Hypertext Application Language"
	case MimetypesHTML:
		return "HyperText Markup Language (HTML)"
	case MimetypesIRM:
		return "IBM DB2 Rights Manager"
	case MimetypesSC:
		return "IBM Electronic Media Management System - Secure Container"
	case MimetypesICS:
		return "iCalendar"
	case MimetypesICC:
		return "ICC profile"
	case MimetypesICO:
		return "Icon Image"
	case MimetypesIGL:
		return "igLoader"
	case MimetypesIEF:
		return "Image Exchange Format"
	case MimetypesIVP:
		return "ImmerVision PURE Players"
	case MimetypesIVU:
		return "ImmerVision PURE Players"
	case MimetypesRIF:
		return "IMS Networks"
	case Mimetypes3DML:
		return "In3D - 3DML"
	case MimetypesSPOT:
		return "In3D - 3DML"
	case MimetypesIGS:
		return "Initial Graphics Exchange Specification (IGES)"
	case MimetypesI2G:
		return "Interactive Geometry Software"
	case MimetypesCDY:
		return "Interactive Geometry Software Cinderella"
	case MimetypesXPW:
		return "Intercon FormNet"
	case MimetypesFCS:
		return "International Society for Advancement of Cytometry"
	case MimetypesIPFIX:
		return "Internet Protocol Flow Information Export"
	case MimetypesCER:
		return "Internet Public Key Infrastructure - Certificate"
	case MimetypesPKI:
		return "Internet Public Key Infrastructure - Certificate Management Protocole"
	case MimetypesCRL:
		return "Internet Public Key Infrastructure - Certificate Revocation Lists"
	case MimetypesPKIPATH:
		return "Internet Public Key Infrastructure - Certification Path"
	case MimetypesIGM:
		return "IOCOM Visimeet"
	case MimetypesRCPROFILE:
		return "IP Unplugged Roaming Client"
	case MimetypesIRP:
		return "iRepository / Lucidoc Editor"
	case MimetypesJAD:
		return "J2ME App Descriptor"
	case MimetypesJAR:
		return "Java Archive"
	case MimetypesCLASS:
		return "Java Bytecode File"
	case MimetypesJNLP:
		return "Java Network Launching Protocol"
	case MimetypesSER:
		return "Java Serialized Object"
	case MimetypesJAVA:
		return "Java Source File"
	case MimetypesJS:
		return "JavaScript"
	case MimetypesTXTJS:
		return "JavaScript Module"
	case MimetypesMJS:
		return "JavaScript Module"
	case MimetypesJSON:
		return "JavaScript Object Notation (JSON)"
	case MimetypesJODA:
		return "Joda Archive"
	case MimetypesJPM:
		return "JPEG 2000 Compound Image File Format"
	case MimetypesIMGJPEG:
		return "JPEG Image"
	case MimetypesCITRIXJPEG:
		return "JPEG Image (Citrix client)"
	case MimetypesPJPEG:
		return "JPEG Image (Progressive)"
	case MimetypesJPGV:
		return "JPGVideo"
	case MimetypesJSONLD:
		return "JSON - Linked Data"
	case MimetypesKTZ:
		return "Kahootz"
	case MimetypesMMD:
		return "Karaoke on Chipnuts Chipsets"
	case MimetypesKARBON:
		return "KDE KOffice Office Suite - Karbon"
	case MimetypesCHRT:
		return "KDE KOffice Office Suite - KChart"
	case MimetypesKFO:
		return "KDE KOffice Office Suite - Kformula"
	case MimetypesFLW:
		return "KDE KOffice Office Suite - Kivio"
	case MimetypesKON:
		return "KDE KOffice Office Suite - Kontour"
	case MimetypesKPR:
		return "KDE KOffice Office Suite - Kpresenter"
	case MimetypesKSP:
		return "KDE KOffice Office Suite - Kspread"
	case MimetypesKWD:
		return "KDE KOffice Office Suite - Kword"
	case MimetypesHTKE:
		return "Kenamea App"
	case MimetypesKIA:
		return "Kidspiration"
	case MimetypesKNE:
		return "Kinar Applications"
	case MimetypesSSE:
		return "Kodak Storyshare"
	case MimetypesLASXML:
		return "Laser App Enterprise"
	case MimetypesLATEX:
		return "LaTeX"
	case MimetypesLBD:
		return "Life Balance - Desktop Edition"
	case MimetypesLBE:
		return "Life Balance - Exchange Format"
	case MimetypesJAM:
		return "Lightspeed Audio Lab"
	case Mimetypes123:
		return "Lotus 1-2-3"
	case MimetypesAPR:
		return "Lotus Approach"
	case MimetypesPRE:
		return "Lotus Freelance"
	case MimetypesNSF:
		return "Lotus Notes"
	case MimetypesORG:
		return "Lotus Organizer"
	case MimetypesSCM:
		return "Lotus Screencam"
	case MimetypesLWP:
		return "Lotus Wordpro"
	case MimetypesLVP:
		return "Lucent Voice"
	case MimetypesM3U:
		return "M3U (Multimedia Playlist)"
	case MimetypesM4V:
		return "M4v"
	case MimetypesHQX:
		return "Macintosh BinHex 4.0"
	case MimetypesPORTPKG:
		return "MacPorts Port System"
	case MimetypesMGP:
		return "MapGuide DBXML"
	case MimetypesMRC:
		return "MARC Formats"
	case MimetypesMRCX:
		return "MARC21 XML Schema"
	case MimetypesMXF:
		return "Material Exchange Format"
	case MimetypesNBP:
		return "Mathematica Notebook Player"
	case MimetypesMA:
		return "Mathematica Notebooks"
	case MimetypesMATHML:
		return "Mathematical Markup Language"
	case MimetypesMBOX:
		return "Mbox database files"
	case MimetypesMC1:
		return "MedCalc"
	case MimetypesMSCML:
		return "Media Server Control Markup Language"
	case MimetypesCDKEY:
		return "MediaRemote"
	case MimetypesMWF:
		return "Medical Waveform Encoding Format"
	case MimetypesMFM:
		return "Melody Format for Mobile Platform"
	case MimetypesMSH:
		return "Mesh Data Type"
	case MimetypesMADS:
		return "Metadata Authority Description Schema"
	case MimetypesMETS:
		return "Metadata Encoding and Transmission Standard"
	case MimetypesMODS:
		return "Metadata Object Description Schema"
	case MimetypesMETA4:
		return "Metalink"
	case MimetypesMCD:
		return "Micro CADAM Helix D&D"
	case MimetypesFLO:
		return "Micrografx"
	case MimetypesIGX:
		return "Micrografx iGrafx Professional"
	case MimetypesES3:
		return "MICROSEC e-Szign "
	case MimetypesMDB:
		return "Microsoft Access"
	case MimetypesASF:
		return "Microsoft Advanced Systems Format (ASF)"
	case MimetypesEXE:
		return "Microsoft Application"
	case MimetypesCIL:
		return "Microsoft Artgalry"
	case MimetypesCAB:
		return "Microsoft Cabinet File"
	case MimetypesIMS:
		return "Microsoft Class Server"
	case MimetypesAPPLICATION:
		return "Microsoft ClickOnce"
	case MimetypesCLP:
		return "Microsoft Clipboard Clip"
	case MimetypesMDI:
		return "Microsoft Document Imaging Format"
	case MimetypesEOT:
		return "Microsoft Embedded OpenType"
	case MimetypesXLS:
		return "Microsoft Excel"
	case MimetypesXLAM:
		return "Microsoft Excel - Add-In File"
	case MimetypesXLSB:
		return "Microsoft Excel - Binary Workbook"
	case MimetypesXLTM:
		return "Microsoft Excel - Macro-Enabled Template File"
	case MimetypesXLSM:
		return "Microsoft Excel - Macro-Enabled Workbook"
	case MimetypesCHM:
		return "Microsoft Html Help File"
	case MimetypesCRD:
		return "Microsoft Information Card"
	case MimetypesLRM:
		return "Microsoft Learning Resource Module"
	case MimetypesMVB:
		return "Microsoft MediaView"
	case MimetypesMNY:
		return "Microsoft Money"
	case MimetypesPPTX:
		return "Microsoft Office - OOXML - Presentation"
	case MimetypesSLDX:
		return "Microsoft Office - OOXML - Presentation (Slide)"
	case MimetypesPPSX:
		return "Microsoft Office - OOXML - Presentation (Slideshow)"
	case MimetypesPOTX:
		return "Microsoft Office - OOXML - Presentation Template"
	case MimetypesXLSX:
		return "Microsoft Office - OOXML - Spreadsheet"
	case MimetypesXLTX:
		return "Microsoft Office - OOXML - Spreadsheet Template"
	case MimetypesDOCX:
		return "Microsoft Office - OOXML - Word Document"
	case MimetypesDOTX:
		return "Microsoft Office - OOXML - Word Document Template"
	case MimetypesOBD:
		return "Microsoft Office Binder"
	case MimetypesTHMX:
		return "Microsoft Office System Release Theme"
	case MimetypesONETOC:
		return "Microsoft OneNote"
	case MimetypesPYA:
		return "Microsoft PlayReady Ecosystem"
	case MimetypesPYV:
		return "Microsoft PlayReady Ecosystem Video"
	case MimetypesPPT:
		return "Microsoft PowerPoint"
	case MimetypesPPAM:
		return "Microsoft PowerPoint - Add-in file"
	case MimetypesSLDM:
		return "Microsoft PowerPoint - Macro-Enabled Open XML Slide"
	case MimetypesPPTM:
		return "Microsoft PowerPoint - Macro-Enabled Presentation File"
	case MimetypesPPSM:
		return "Microsoft PowerPoint - Macro-Enabled Slide Show File"
	case MimetypesPOTM:
		return "Microsoft PowerPoint - Macro-Enabled Template File"
	case MimetypesMPP:
		return "Microsoft Project"
	case MimetypesPUB:
		return "Microsoft Publisher"
	case MimetypesSCD:
		return "Microsoft Schedule+"
	case MimetypesXAP:
		return "Microsoft Silverlight"
	case MimetypesSTL:
		return "Microsoft Trust UI Provider - Certificate Trust Link"
	case MimetypesCAT:
		return "Microsoft Trust UI Provider - Security Catalog"
	case MimetypesVSD:
		return "Microsoft Visio"
	case MimetypesVSDX:
		return "Microsoft Visio 2013"
	case MimetypesWM:
		return "Microsoft Windows Media"
	case MimetypesWMA:
		return "Microsoft Windows Media Audio"
	case MimetypesWAX:
		return "Microsoft Windows Media Audio Redirector"
	case MimetypesWMX:
		return "Microsoft Windows Media Audio/Video Playlist"
	case MimetypesWMD:
		return "Microsoft Windows Media Player Download Package"
	case MimetypesWPL:
		return "Microsoft Windows Media Player Playlist"
	case MimetypesWMZ:
		return "Microsoft Windows Media Player Skin Package"
	case MimetypesWMV:
		return "Microsoft Windows Media Video"
	case MimetypesWVX:
		return "Microsoft Windows Media Video Playlist"
	case MimetypesWMF:
		return "Microsoft Windows Metafile"
	case MimetypesTRM:
		return "Microsoft Windows Terminal Services"
	case MimetypesDOC:
		return "Microsoft Word"
	case MimetypesDOCM:
		return "Microsoft Word - Macro-Enabled Document"
	case MimetypesDOTM:
		return "Microsoft Word - Macro-Enabled Template"
	case MimetypesWRI:
		return "Microsoft Wordpad"
	case MimetypesWPS:
		return "Microsoft Works"
	case MimetypesXBAP:
		return "Microsoft XAML Browser Application"
	case MimetypesXPS:
		return "Microsoft XML Paper Specification"
	case MimetypesMIDI:
		return "MIDI"
	case MimetypesMID:
		return "MIDI - Musical Instrument Digital Interface"
	case MimetypesMPY:
		return "MiniPay"
	case MimetypesAFP:
		return "MO:DCA-P"
	case MimetypesRMS:
		return "Mobile Information Device Profile"
	case MimetypesTMO:
		return "MobileTV"
	case MimetypesPRC:
		return "Mobipocket"
	case MimetypesMBK:
		return "Mobius Management Systems - Basket file"
	case MimetypesDIS:
		return "Mobius Management Systems - Distribution Database"
	case MimetypesPLC:
		return "Mobius Management Systems - Policy Definition Language File"
	case MimetypesMQY:
		return "Mobius Management Systems - Query File"
	case MimetypesMSL:
		return "Mobius Management Systems - Script Language"
	case MimetypesTXF:
		return "Mobius Management Systems - Topic Index File"
	case MimetypesDAF:
		return "Mobius Management Systems - UniversalArchive"
	case MimetypesFLY:
		return "mod_fly / fly.cgi"
	case MimetypesMPC:
		return "Mophun Certificate"
	case MimetypesMPN:
		return "Mophun VM"
	case MimetypesMJ2:
		return "Motion JPEG 2000"
	case MimetypesMPGA:
		return "MPEG Audio"
	case MimetypesTS:
		return "MPEG Transport Stream"
	case MimetypesMXU:
		return "MPEG Url"
	case MimetypesMPEG:
		return "MPEG Video"
	case MimetypesM21:
		return "MPEG-21"
	case MimetypesMP4A:
		return "MPEG-4 Audio"
	case MimetypesVDMP4:
		return "MPEG-4 Video"
	case MimetypesMP4:
		return "MPEG4"
	case MimetypesM3U8:
		return "Multimedia Playlist Unicode"
	case MimetypesMUS:
		return "MUsical Score Interpreted Code Invented for the ASCII designation of Notation"
	case MimetypesMSTY:
		return "Muvee Automatic Video Editing"
	case MimetypesMXML:
		return "MXML"
	case MimetypesNGDAT:
		return "N-Gage Game Data"
	case MimetypesNGAGE:
		return "N-Gage Game Installer"
	case MimetypesNCX:
		return "Navigation Control file for XML (for ePub)"
	case MimetypesNC:
		return "Network Common Data Form (NetCDF)"
	case MimetypesNLU:
		return "neuroLanguage"
	case MimetypesDNA:
		return "New Moon Liftoff/DNA"
	case MimetypesNND:
		return "NobleNet Directory"
	case MimetypesNNS:
		return "NobleNet Sealer"
	case MimetypesNNW:
		return "NobleNet Web"
	case MimetypesRPST:
		return "Nokia Radio Application - Preset"
	case MimetypesRPSS:
		return "Nokia Radio Application - Preset"
	case MimetypesN3:
		return "Notation3"
	case MimetypesEDM:
		return "Novadigm's RADIA and EDM products"
	case MimetypesEDX:
		return "Novadigm's RADIA and EDM products"
	case MimetypesEXT:
		return "Novadigm's RADIA and EDM products"
	case MimetypesGPH:
		return "NpGraphIt"
	case MimetypesECELP4800:
		return "Nuera ECELP 4800"
	case MimetypesECELP7470:
		return "Nuera ECELP 7470"
	case MimetypesECELP9600:
		return "Nuera ECELP 9600"
	case MimetypesODA:
		return "Office Document Architecture"
	case MimetypesOGX:
		return "Ogg"
	case MimetypesOGA:
		return "Ogg Audio"
	case MimetypesOGV:
		return "Ogg Video"
	case MimetypesDD2:
		return "OMA Download Agents"
	case MimetypesOTH:
		return "Open Document Text Web"
	case MimetypesOPF:
		return "Open eBook Publication Structure"
	case MimetypesQBO:
		return "Open Financial Exchange"
	case MimetypesOXT:
		return "Open Office Extension"
	case MimetypesOSF:
		return "Open Score Format"
	case MimetypesWEBA:
		return "Open Web Media Project - Audio"
	case MimetypesWEBM:
		return "Open Web Media Project - Video"
	case MimetypesODC:
		return "OpenDocument Chart"
	case MimetypesOTC:
		return "OpenDocument Chart Template"
	case MimetypesODB:
		return "OpenDocument Database"
	case MimetypesODF:
		return "OpenDocument Formula"
	case MimetypesODFT:
		return "OpenDocument Formula Template"
	case MimetypesODG:
		return "OpenDocument Graphics"
	case MimetypesOTG:
		return "OpenDocument Graphics Template"
	case MimetypesODI:
		return "OpenDocument Image"
	case MimetypesOTI:
		return "OpenDocument Image Template"
	case MimetypesODP:
		return "OpenDocument Presentation"
	case MimetypesOTP:
		return "OpenDocument Presentation Template"
	case MimetypesODS:
		return "OpenDocument Spreadsheet"
	case MimetypesOTS:
		return "OpenDocument Spreadsheet Template"
	case MimetypesODT:
		return "OpenDocument Text"
	case MimetypesODM:
		return "OpenDocument Text Master"
	case MimetypesOTT:
		return "OpenDocument Text Template"
	case MimetypesKTX:
		return "OpenGL Textures (KTX)"
	case MimetypesSXC:
		return "OpenOffice - Calc (Spreadsheet)"
	case MimetypesSTC:
		return "OpenOffice - Calc Template (Spreadsheet)"
	case MimetypesSXD:
		return "OpenOffice - Draw (Graphics)"
	case MimetypesSTD:
		return "OpenOffice - Draw Template (Graphics)"
	case MimetypesSXI:
		return "OpenOffice - Impress (Presentation)"
	case MimetypesSTI:
		return "OpenOffice - Impress Template (Presentation)"
	case MimetypesSXM:
		return "OpenOffice - Math (Formula)"
	case MimetypesSXW:
		return "OpenOffice - Writer (Text - HTML)"
	case MimetypesSXG:
		return "OpenOffice - Writer (Text - HTML)"
	case MimetypesSTW:
		return "OpenOffice - Writer Template (Text - HTML)"
	case MimetypesOTF:
		return "OpenType Font File"
	case MimetypesOPUS:
		return "Opus Audio"
	case MimetypesOSFPVG:
		return "OSFPVG"
	case MimetypesDP:
		return "OSGi Deployment Package"
	case MimetypesPDB:
		return "PalmOS Data"
	case MimetypesP:
		return "Pascal Source File"
	case MimetypesPAW:
		return "PawaaFILE"
	case MimetypesPCLXL:
		return "PCL 6 Enhanced (Formely PCL XL)"
	case MimetypesEFIF:
		return "Pcsel eFIF File"
	case MimetypesPCX:
		return "PCX Image"
	case MimetypesPSD:
		return "Photoshop Document"
	case MimetypesPRF:
		return "PICSRules"
	case MimetypesPIC:
		return "PICT Image"
	case MimetypesCHAT:
		return "pIRCh"
	case MimetypesP10:
		return "PKCS #10 - Certification Request Standard"
	case MimetypesP12:
		return "PKCS #12 - Personal Information Exchange Syntax Standard"
	case MimetypesP7M:
		return "PKCS #7 - Cryptographic Message Syntax Standard"
	case MimetypesP7S:
		return "PKCS #7 - Cryptographic Message Syntax Standard"
	case MimetypesP7R:
		return "PKCS #7 - Cryptographic Message Syntax Standard (Certificate Request Response)"
	case MimetypesP7B:
		return "PKCS #7 - Cryptographic Message Syntax Standard (Certificates)"
	case MimetypesP8:
		return "PKCS #8 - Private-Key Information Syntax Standard"
	case MimetypesPLF:
		return "PocketLearn Viewers"
	case MimetypesPNM:
		return "Portable Anymap Image"
	case MimetypesPBM:
		return "Portable Bitmap Format"
	case MimetypesPCF:
		return "Portable Compiled Format"
	case MimetypesPFR:
		return "Portable Font Resource"
	case MimetypesPGN:
		return "Portable Game Notation (Chess Games)"
	case MimetypesPGM:
		return "Portable Graymap Format"
	case MimetypesIMGPNG:
		return "Portable Network Graphics (PNG)"
	case MimetypesPNG:
		return "Portable Network Graphics (PNG) (Citrix client)"
	case MimetypesXPNG:
		return "Portable Network Graphics (PNG) (x-token)"
	case MimetypesPPM:
		return "Portable Pixmap Format"
	case MimetypesPSKCXML:
		return "Portable Symmetric Key Container"
	case MimetypesPML:
		return "PosML"
	case MimetypesAI:
		return "PostScript"
	case MimetypesPFA:
		return "PostScript Fonts"
	case MimetypesPBD:
		return "PowerBuilder"
	case MimetypesCRYPGP:
		return "Pretty Good Privacy"
	case MimetypesSIGPGP:
		return "Pretty Good Privacy - Signature"
	case MimetypesBOX:
		return "Preview Systems ZipLock/VBox"
	case MimetypesPTID:
		return "Princeton Video Image"
	case MimetypesPLS:
		return "Pronunciation Lexicon Specification"
	case MimetypesSTR:
		return "Proprietary P&G Standard Reporting System"
	case MimetypesEI6:
		return "Proprietary P&G Standard Reporting System"
	case MimetypesDSC:
		return "PRS Lines Tag"
	case MimetypesPSF:
		return "PSF Fonts"
	case MimetypesQPS:
		return "PubliShare Objects"
	case MimetypesWG:
		return "Qualcomm's Plaza Mobile Internet"
	case MimetypesQXD:
		return "QuarkXpress"
	case MimetypesESF:
		return "QUASS Stream Player"
	case MimetypesMSF:
		return "QUASS Stream Player"
	case MimetypesSSF:
		return "QUASS Stream Player"
	case MimetypesQAM:
		return "QuickAnime Player"
	case MimetypesQFX:
		return "Quicken"
	case MimetypesQT:
		return "Quicktime Video"
	case MimetypesRAR:
		return "RAR Archive"
	case MimetypesRAM:
		return "Real Audio Sound"
	case MimetypesRMP:
		return "Real Audio Sound"
	case MimetypesRSD:
		return "Really Simple Discovery"
	case MimetypesRM:
		return "RealMedia"
	case MimetypesBED:
		return "RealVNC"
	case MimetypesMXL:
		return "Recordare Applications"
	case MimetypesMUSICXML:
		return "Recordare Applications"
	case MimetypesRNC:
		return "Relax NG Compact Syntax"
	case MimetypesRDZ:
		return "RemoteDocs R-Viewer"
	case MimetypesRDF:
		return "Resource Description Framework"
	case MimetypesRP9:
		return "RetroPlatform Player"
	case MimetypesJISP:
		return "RhymBox"
	case MimetypesRTF:
		return "Rich Text Format"
	case MimetypesRTX:
		return "Rich Text Format (RTF)"
	case MimetypesLINK66:
		return "ROUTE 66 Location Based Services"
	case MimetypesRSSXML:
		return "RSS - Really Simple Syndication"
	case MimetypesSHF:
		return "S Hexdump Format"
	case MimetypesST:
		return "SailingTracker"
	case MimetypesSVG:
		return "Scalable Vector Graphics (SVG)"
	case MimetypesSUS:
		return "ScheduleUs"
	case MimetypesSRU:
		return "Search/Retrieve via URL Response Format"
	case MimetypesSETPAY:
		return "Secure Electronic Transaction - Payment"
	case MimetypesSETREG:
		return "Secure Electronic Transaction - Registration"
	case MimetypesSEMA:
		return "Secured eMail"
	case MimetypesSEMD:
		return "Secured eMail"
	case MimetypesSEMF:
		return "Secured eMail"
	case MimetypesSEE:
		return "SeeMail"
	case MimetypesSNF:
		return "Server Normal Format"
	case MimetypesSPQ:
		return "Server-Based Certificate Validation Protocol - Validation Policies - Request"
	case MimetypesSPP:
		return "Server-Based Certificate Validation Protocol - Validation Policies - Response"
	case MimetypesSCQ:
		return "Server-Based Certificate Validation Protocol - Validation Request"
	case MimetypesSCS:
		return "Server-Based Certificate Validation Protocol - Validation Response"
	case MimetypesSDP:
		return "Session Description Protocol"
	case MimetypesETX:
		return "Setext"
	case MimetypesMOVIE:
		return "SGI Movie"
	case MimetypesIFM:
		return "Shana Informed Filler"
	case MimetypesITP:
		return "Shana Informed Filler"
	case MimetypesIIF:
		return "Shana Informed Filler"
	case MimetypesIPK:
		return "Shana Informed Filler"
	case MimetypesTFI:
		return "Sharing Transaction Fraud Data"
	case MimetypesSHAR:
		return "Shell Archive"
	case MimetypesRGB:
		return "Silicon Graphics RGB Bitmap"
	case MimetypesSLT:
		return "SimpleAnimeLite Player"
	case MimetypesASO:
		return "Simply Accounting"
	case MimetypesIMP:
		return "Simply Accounting - Data Import"
	case MimetypesTWD:
		return "SimTech MindMapper"
	case MimetypesCSP:
		return "Sixth Floor Media - CommonSpace"
	case MimetypesSAF:
		return "SMAF Audio"
	case MimetypesMMF:
		return "SMAF File"
	case MimetypesSPF:
		return "SMAF Phrase"
	case MimetypesTEACHER:
		return "SMART Technologies Apps"
	case MimetypesSVD:
		return "SourceView Document"
	case MimetypesRQ:
		return "SPARQL - Query"
	case MimetypesSRX:
		return "SPARQL - Results"
	case MimetypesGRAM:
		return "Speech Recognition Grammar Specification"
	case MimetypesGRXML:
		return "Speech Recognition Grammar Specification - XML"
	case MimetypesSSML:
		return "Speech Synthesis Markup Language"
	case MimetypesSKP:
		return "SSEYO Koan Play File"
	case MimetypesSGML:
		return "Standard Generalized Markup Language (SGML)"
	case MimetypesSDC:
		return "StarOffice - Calc"
	case MimetypesSDA:
		return "StarOffice - Draw"
	case MimetypesSDD:
		return "StarOffice - Impress"
	case MimetypesSMF:
		return "StarOffice - Math"
	case MimetypesSDW:
		return "StarOffice - Writer"
	case MimetypesSGL:
		return "StarOffice - Writer (Global)"
	case MimetypesSM:
		return "StepMania"
	case MimetypesSIT:
		return "Stuffit Archive"
	case MimetypesSITX:
		return "Stuffit Archive"
	case MimetypesSDKM:
		return "SudokuMagic"
	case MimetypesXO:
		return "Sugar Linux Application Bundle"
	case MimetypesAU:
		return "Sun Audio - Au file format"
	case MimetypesWQD:
		return "SundaHus WQ"
	case MimetypesSIS:
		return "Symbian Install Package"
	case MimetypesSMI:
		return "Synchronized Multimedia Integration Language"
	case MimetypesXSM:
		return "SyncML"
	case MimetypesBDM:
		return "SyncML - Device Management"
	case MimetypesXDM:
		return "SyncML - Device Management"
	case MimetypesSV4CPIO:
		return "System V Release 4 CPIO Archive"
	case MimetypesSV4CRC:
		return "System V Release 4 CPIO Checksum Data"
	case MimetypesSBML:
		return "Systems Biology Markup Language"
	case MimetypesTSV:
		return "Tab Seperated Values"
	case MimetypesTIFF:
		return "Tagged Image File Format"
	case MimetypesTAO:
		return "Tao Intent"
	case MimetypesTAR:
		return "Tar File (Tape Archive)"
	case MimetypesTCL:
		return "Tcl Script"
	case MimetypesTEX:
		return "TeX"
	case MimetypesTFM:
		return "TeX Font Metric"
	case MimetypesTEI:
		return "Text Encoding and Interchange"
	case MimetypesTXT:
		return "Text File"
	case MimetypesDXP:
		return "TIBCO Spotfire"
	case MimetypesSFS:
		return "TIBCO Spotfire"
	case MimetypesTSD:
		return "Time Stamped Data Envelope"
	case MimetypesTPT:
		return "TRI Systems Config"
	case MimetypesMXS:
		return "Triscape Map Explorer"
	case MimetypesT:
		return "troff"
	case MimetypesTRA:
		return "True BASIC"
	case MimetypesTTF:
		return "TrueType Font"
	case MimetypesTTL:
		return "Turtle (Terse RDF Triple Language)"
	case MimetypesUMJ:
		return "UMAJIN"
	case MimetypesUOML:
		return "Unique Object Markup Language"
	case MimetypesUNITYWEB:
		return "Unity 3d"
	case MimetypesUFD:
		return "Universal Forms Description Language"
	case MimetypesURI:
		return "URI Resolution Services"
	case MimetypesUTZ:
		return "User Interface Quartz - Theme (Symbian)"
	case MimetypesUSTAR:
		return "Ustar (Uniform Standard Tape Archive)"
	case MimetypesUU:
		return "UUEncode"
	case MimetypesVCS:
		return "vCalendar"
	case MimetypesVCF:
		return "vCard"
	case MimetypesVCD:
		return "Video CD"
	case MimetypesVSF:
		return "Viewport+"
	case MimetypesWRL:
		return "Virtual Reality Modeling Language"
	case MimetypesVCX:
		return "VirtualCatalog"
	case MimetypesMTS:
		return "Virtue MTS"
	case MimetypesVTU:
		return "Virtue VTU"
	case MimetypesVIS:
		return "Visionary"
	case MimetypesVIV:
		return "Vivo"
	case MimetypesCCXML:
		return "Voice Browser Call Control"
	case MimetypesVXML:
		return "VoiceXML"
	case MimetypesSRC:
		return "WAIS Source"
	case MimetypesWBXML:
		return "WAP Binary XML (WBXML)"
	case MimetypesWBMP:
		return "WAP Bitamp (WBMP)"
	case MimetypesWAV:
		return "Waveform Audio File Format (WAV)"
	case MimetypesDAVMOUNT:
		return "Web Distributed Authoring and Versioning"
	case MimetypesWOFF:
		return "Web Open Font Format"
	case MimetypesWSPOLICY:
		return "Web Services Policy"
	case MimetypesWEBP:
		return "WebP Image"
	case MimetypesWTB:
		return "WebTurbo"
	case MimetypesWGT:
		return "Widget Packaging and XML Configuration"
	case MimetypesHLP:
		return "WinHelp"
	case MimetypesWML:
		return "Wireless Markup Language (WML)"
	case MimetypesWMLS:
		return "Wireless Markup Language Script (WMLScript)"
	case MimetypesWMLSC:
		return "WMLScript"
	case MimetypesWPD:
		return "Wordperfect"
	case MimetypesSTF:
		return "Worldtalk"
	case MimetypesWSDL:
		return "WSDL - Web Services Description Language"
	case MimetypesXBM:
		return "X BitMap"
	case MimetypesXPM:
		return "X PixMap"
	case MimetypesXWD:
		return "X Window Dump"
	case MimetypesDER:
		return "X.509 Certificate"
	case MimetypesFIG:
		return "Xfig"
	case MimetypesXHTML:
		return "XHTML - The Extensible HyperText Markup Language"
	case MimetypesXML:
		return "XML - Extensible Markup Language"
	case MimetypesXDF:
		return "XML Configuration Access Protocol - XCAP Diff"
	case MimetypesXENC:
		return "XML Encryption Syntax and Processing"
	case MimetypesXER:
		return "XML Patch Framework"
	case MimetypesRL:
		return "XML Resource Lists"
	case MimetypesRS:
		return "XML Resource Lists"
	case MimetypesRLD:
		return "XML Resource Lists Diff"
	case MimetypesXSLT:
		return "XML Transformations"
	case MimetypesXOP:
		return "XML-Binary Optimized Packaging"
	case MimetypesXPI:
		return "XPInstall - Mozilla"
	case MimetypesXSPF:
		return "XSPF - XML Shareable Playlist Format"
	case MimetypesXUL:
		return "XUL - XML User Interface Language"
	case MimetypesXYZ:
		return "XYZ File Format"
	case MimetypesYAML:
		return "YAML Ain't Markup Language / Yet Another Markup Language"
	case MimetypesYANG:
		return "YANG Data Modeling Language"
	case MimetypesYIN:
		return "YIN (YANG - XML)"
	case MimetypesZIR:
		return "Z.U.L. Geometry"
	case MimetypesZIP:
		return "Zip Archive"
	case MimetypesZMM:
		return "ZVUE Media Manager"
	case MimetypesZAZ:
		return "Zzazz Deck"
	default:
		return "Unknown Constraint Severity"
	}
}
