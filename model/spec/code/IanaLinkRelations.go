package code

import (
	"fmt"
	"strings"
)

type IanaLinkRelations string

const (
	IanaLinkRelationsAbout                  IanaLinkRelations = "about"
	IanaLinkRelationsAcl                    IanaLinkRelations = "acl"
	IanaLinkRelationsAlternate              IanaLinkRelations = "alternate"
	IanaLinkRelationsAmphtml                IanaLinkRelations = "amphtml"
	IanaLinkRelationsAppendix               IanaLinkRelations = "appendix"
	IanaLinkRelationsAppleTouchIcon         IanaLinkRelations = "apple-touch-icon"
	IanaLinkRelationsAppleTouchStartupImage IanaLinkRelations = "apple-touch-startup-image"
	IanaLinkRelationsArchives               IanaLinkRelations = "archives"
	IanaLinkRelationsAuthor                 IanaLinkRelations = "author"
	IanaLinkRelationsBlockedBy              IanaLinkRelations = "blocked-by"
	IanaLinkRelationsBookmark               IanaLinkRelations = "bookmark"
	IanaLinkRelationsCanonical              IanaLinkRelations = "canonical"
	IanaLinkRelationsChapter                IanaLinkRelations = "chapter"
	IanaLinkRelationsCiteAs                 IanaLinkRelations = "cite-as"
	IanaLinkRelationsCollection             IanaLinkRelations = "collection"
	IanaLinkRelationsContents               IanaLinkRelations = "contents"
	IanaLinkRelationsConvertedfrom          IanaLinkRelations = "convertedFrom"
	IanaLinkRelationsCopyright              IanaLinkRelations = "copyright"
	IanaLinkRelationsCreateForm             IanaLinkRelations = "create-form"
	IanaLinkRelationsCurrent                IanaLinkRelations = "current"
	IanaLinkRelationsDescribedby            IanaLinkRelations = "describedby"
	IanaLinkRelationsDescribes              IanaLinkRelations = "describes"
	IanaLinkRelationsDisclosure             IanaLinkRelations = "disclosure"
	IanaLinkRelationsDnsPrefetch            IanaLinkRelations = "dns-prefetch"
	IanaLinkRelationsDuplicate              IanaLinkRelations = "duplicate"
	IanaLinkRelationsEdit                   IanaLinkRelations = "edit"
	IanaLinkRelationsEditForm               IanaLinkRelations = "edit-form"
	IanaLinkRelationsEditMedia              IanaLinkRelations = "edit-media"
	IanaLinkRelationsEnclosure              IanaLinkRelations = "enclosure"
	IanaLinkRelationsExternal               IanaLinkRelations = "external"
	IanaLinkRelationsFirst                  IanaLinkRelations = "first"
	IanaLinkRelationsGlossary               IanaLinkRelations = "glossary"
	IanaLinkRelationsHelp                   IanaLinkRelations = "help"
	IanaLinkRelationsHosts                  IanaLinkRelations = "hosts"
	IanaLinkRelationsHub                    IanaLinkRelations = "hub"
	IanaLinkRelationsIcon                   IanaLinkRelations = "icon"
	IanaLinkRelationsIndex                  IanaLinkRelations = "index"
	IanaLinkRelationsIntervalafter          IanaLinkRelations = "intervalAfter"
	IanaLinkRelationsIntervalbefore         IanaLinkRelations = "intervalBefore"
	IanaLinkRelationsIntervalcontains       IanaLinkRelations = "intervalContains"
	IanaLinkRelationsIntervaldisjoint       IanaLinkRelations = "intervalDisjoint"
	IanaLinkRelationsIntervalduring         IanaLinkRelations = "intervalDuring"
	IanaLinkRelationsIntervalequals         IanaLinkRelations = "intervalEquals"
	IanaLinkRelationsIntervalfinishedby     IanaLinkRelations = "intervalFinishedBy"
	IanaLinkRelationsIntervalfinishes       IanaLinkRelations = "intervalFinishes"
	IanaLinkRelationsIntervalin             IanaLinkRelations = "intervalIn"
	IanaLinkRelationsIntervalmeets          IanaLinkRelations = "intervalMeets"
	IanaLinkRelationsIntervalmetby          IanaLinkRelations = "intervalMetBy"
	IanaLinkRelationsIntervaloverlappedby   IanaLinkRelations = "intervalOverlappedBy"
	IanaLinkRelationsIntervaloverlaps       IanaLinkRelations = "intervalOverlaps"
	IanaLinkRelationsIntervalstartedby      IanaLinkRelations = "intervalStartedBy"
	IanaLinkRelationsIntervalstarts         IanaLinkRelations = "intervalStarts"
	IanaLinkRelationsItem                   IanaLinkRelations = "item"
	IanaLinkRelationsLast                   IanaLinkRelations = "last"
	IanaLinkRelationsLatestVersion          IanaLinkRelations = "latest-version"
	IanaLinkRelationsLicense                IanaLinkRelations = "license"
	IanaLinkRelationsLinkset                IanaLinkRelations = "linkset"
	IanaLinkRelationsLrdd                   IanaLinkRelations = "lrdd"
	IanaLinkRelationsManifest               IanaLinkRelations = "manifest"
	IanaLinkRelationsMaskIcon               IanaLinkRelations = "mask-icon"
	IanaLinkRelationsMediaFeed              IanaLinkRelations = "media-feed"
	IanaLinkRelationsMemento                IanaLinkRelations = "memento"
	IanaLinkRelationsMicropub               IanaLinkRelations = "micropub"
	IanaLinkRelationsModulepreload          IanaLinkRelations = "modulepreload"
	IanaLinkRelationsMonitor                IanaLinkRelations = "monitor"
	IanaLinkRelationsMonitorGroup           IanaLinkRelations = "monitor-group"
	IanaLinkRelationsNext                   IanaLinkRelations = "next"
	IanaLinkRelationsNextArchive            IanaLinkRelations = "next-archive"
	IanaLinkRelationsNofollow               IanaLinkRelations = "nofollow"
	IanaLinkRelationsNoopener               IanaLinkRelations = "noopener"
	IanaLinkRelationsNoreferrer             IanaLinkRelations = "noreferrer"
	IanaLinkRelationsOpener                 IanaLinkRelations = "opener"
	IanaLinkRelationsOpenid2LocalId         IanaLinkRelations = "openid2.local_id"
	IanaLinkRelationsOpenid2Provider        IanaLinkRelations = "openid2.provider"
	IanaLinkRelationsOriginal               IanaLinkRelations = "original"
	IanaLinkRelationsP3Pv1                  IanaLinkRelations = "P3Pv1"
	IanaLinkRelationsPayment                IanaLinkRelations = "payment"
	IanaLinkRelationsPingback               IanaLinkRelations = "pingback"
	IanaLinkRelationsPreconnect             IanaLinkRelations = "preconnect"
	IanaLinkRelationsPredecessorVersion     IanaLinkRelations = "predecessor-version"
	IanaLinkRelationsPrefetch               IanaLinkRelations = "prefetch"
	IanaLinkRelationsPreload                IanaLinkRelations = "preload"
	IanaLinkRelationsPrerender              IanaLinkRelations = "prerender"
	IanaLinkRelationsPrev                   IanaLinkRelations = "prev"
	IanaLinkRelationsPreview                IanaLinkRelations = "preview"
	IanaLinkRelationsPrevious               IanaLinkRelations = "previous"
	IanaLinkRelationsPrevArchive            IanaLinkRelations = "prev-archive"
	IanaLinkRelationsPrivacyPolicy          IanaLinkRelations = "privacy-policy"
	IanaLinkRelationsProfile                IanaLinkRelations = "profile"
	IanaLinkRelationsPublication            IanaLinkRelations = "publication"
	IanaLinkRelationsRelated                IanaLinkRelations = "related"
	IanaLinkRelationsRestconf               IanaLinkRelations = "restconf"
	IanaLinkRelationsReplies                IanaLinkRelations = "replies"
	IanaLinkRelationsRuleinput              IanaLinkRelations = "ruleinput"
	IanaLinkRelationsSearch                 IanaLinkRelations = "search"
	IanaLinkRelationsSection                IanaLinkRelations = "section"
	IanaLinkRelationsSelf                   IanaLinkRelations = "self"
	IanaLinkRelationsService                IanaLinkRelations = "service"
	IanaLinkRelationsServiceDesc            IanaLinkRelations = "service-desc"
	IanaLinkRelationsServiceDoc             IanaLinkRelations = "service-doc"
	IanaLinkRelationsServiceMeta            IanaLinkRelations = "service-meta"
	IanaLinkRelationsSponsored              IanaLinkRelations = "sponsored"
	IanaLinkRelationsStart                  IanaLinkRelations = "start"
	IanaLinkRelationsStatus                 IanaLinkRelations = "status"
	IanaLinkRelationsStylesheet             IanaLinkRelations = "stylesheet"
	IanaLinkRelationsSubsection             IanaLinkRelations = "subsection"
	IanaLinkRelationsSuccessorVersion       IanaLinkRelations = "successor-version"
	IanaLinkRelationsSunset                 IanaLinkRelations = "sunset"
	IanaLinkRelationsTag                    IanaLinkRelations = "tag"
	IanaLinkRelationsTermsOfService         IanaLinkRelations = "terms-of-service"
	IanaLinkRelationsTimegate               IanaLinkRelations = "timegate"
	IanaLinkRelationsTimemap                IanaLinkRelations = "timemap"
	IanaLinkRelationsType                   IanaLinkRelations = "type"
	IanaLinkRelationsUgc                    IanaLinkRelations = "ugc"
	IanaLinkRelationsUp                     IanaLinkRelations = "up"
	IanaLinkRelationsVersionHistory         IanaLinkRelations = "version-history"
	IanaLinkRelationsVia                    IanaLinkRelations = "via"
	IanaLinkRelationsWebmention             IanaLinkRelations = "webmention"
	IanaLinkRelationsWorkingCopy            IanaLinkRelations = "working-copy"
	IanaLinkRelationsWorkingCopyOf          IanaLinkRelations = "working-copy-of"
)

func AllIanaLinkRelations() []IanaLinkRelations {
	return []IanaLinkRelations{
		IanaLinkRelationsAbout,
		IanaLinkRelationsAcl,
		IanaLinkRelationsAlternate,
		IanaLinkRelationsAmphtml,
		IanaLinkRelationsAppendix,
		IanaLinkRelationsAppleTouchIcon,
		IanaLinkRelationsAppleTouchStartupImage,
		IanaLinkRelationsArchives,
		IanaLinkRelationsAuthor,
		IanaLinkRelationsBlockedBy,
		IanaLinkRelationsBookmark,
		IanaLinkRelationsCanonical,
		IanaLinkRelationsChapter,
		IanaLinkRelationsCiteAs,
		IanaLinkRelationsCollection,
		IanaLinkRelationsContents,
		IanaLinkRelationsConvertedfrom,
		IanaLinkRelationsCopyright,
		IanaLinkRelationsCreateForm,
		IanaLinkRelationsCurrent,
		IanaLinkRelationsDescribedby,
		IanaLinkRelationsDescribes,
		IanaLinkRelationsDisclosure,
		IanaLinkRelationsDnsPrefetch,
		IanaLinkRelationsDuplicate,
		IanaLinkRelationsEdit,
		IanaLinkRelationsEditForm,
		IanaLinkRelationsEditMedia,
		IanaLinkRelationsEnclosure,
		IanaLinkRelationsExternal,
		IanaLinkRelationsFirst,
		IanaLinkRelationsGlossary,
		IanaLinkRelationsHelp,
		IanaLinkRelationsHosts,
		IanaLinkRelationsHub,
		IanaLinkRelationsIcon,
		IanaLinkRelationsIndex,
		IanaLinkRelationsIntervalafter,
		IanaLinkRelationsIntervalbefore,
		IanaLinkRelationsIntervalcontains,
		IanaLinkRelationsIntervaldisjoint,
		IanaLinkRelationsIntervalduring,
		IanaLinkRelationsIntervalequals,
		IanaLinkRelationsIntervalfinishedby,
		IanaLinkRelationsIntervalfinishes,
		IanaLinkRelationsIntervalin,
		IanaLinkRelationsIntervalmeets,
		IanaLinkRelationsIntervalmetby,
		IanaLinkRelationsIntervaloverlappedby,
		IanaLinkRelationsIntervaloverlaps,
		IanaLinkRelationsIntervalstartedby,
		IanaLinkRelationsIntervalstarts,
		IanaLinkRelationsItem,
		IanaLinkRelationsLast,
		IanaLinkRelationsLatestVersion,
		IanaLinkRelationsLicense,
		IanaLinkRelationsLinkset,
		IanaLinkRelationsLrdd,
		IanaLinkRelationsManifest,
		IanaLinkRelationsMaskIcon,
		IanaLinkRelationsMediaFeed,
		IanaLinkRelationsMemento,
		IanaLinkRelationsMicropub,
		IanaLinkRelationsModulepreload,
		IanaLinkRelationsMonitor,
		IanaLinkRelationsMonitorGroup,
		IanaLinkRelationsNext,
		IanaLinkRelationsNextArchive,
		IanaLinkRelationsNofollow,
		IanaLinkRelationsNoopener,
		IanaLinkRelationsNoreferrer,
		IanaLinkRelationsOpener,
		IanaLinkRelationsOpenid2LocalId,
		IanaLinkRelationsOpenid2Provider,
		IanaLinkRelationsOriginal,
		IanaLinkRelationsP3Pv1,
		IanaLinkRelationsPayment,
		IanaLinkRelationsPingback,
		IanaLinkRelationsPreconnect,
		IanaLinkRelationsPredecessorVersion,
		IanaLinkRelationsPrefetch,
		IanaLinkRelationsPreload,
		IanaLinkRelationsPrerender,
		IanaLinkRelationsPrev,
		IanaLinkRelationsPreview,
		IanaLinkRelationsPrevious,
		IanaLinkRelationsPrevArchive,
		IanaLinkRelationsPrivacyPolicy,
		IanaLinkRelationsProfile,
		IanaLinkRelationsPublication,
		IanaLinkRelationsRelated,
		IanaLinkRelationsRestconf,
		IanaLinkRelationsReplies,
		IanaLinkRelationsRuleinput,
		IanaLinkRelationsSearch,
		IanaLinkRelationsSection,
		IanaLinkRelationsSelf,
		IanaLinkRelationsService,
		IanaLinkRelationsServiceDesc,
		IanaLinkRelationsServiceDoc,
		IanaLinkRelationsServiceMeta,
		IanaLinkRelationsSponsored,
		IanaLinkRelationsStart,
		IanaLinkRelationsStatus,
		IanaLinkRelationsStylesheet,
		IanaLinkRelationsSubsection,
		IanaLinkRelationsSuccessorVersion,
		IanaLinkRelationsSunset,
		IanaLinkRelationsTag,
		IanaLinkRelationsTermsOfService,
		IanaLinkRelationsTimegate,
		IanaLinkRelationsTimemap,
		IanaLinkRelationsType,
		IanaLinkRelationsUgc,
		IanaLinkRelationsUp,
		IanaLinkRelationsVersionHistory,
		IanaLinkRelationsVia,
		IanaLinkRelationsWebmention,
		IanaLinkRelationsWorkingCopy,
		IanaLinkRelationsWorkingCopyOf,
	}
}

func FindIanaLinkRelations(filter string) []IanaLinkRelations {
	ret := make([]IanaLinkRelations, 0)
	for _, at := range AllIanaLinkRelations() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au IanaLinkRelations) ToString() {
	fmt.Println(au.String())
}
func (au IanaLinkRelations) String() string {
	switch au {
	case IanaLinkRelationsAbout:
		return "Refers to a resource that is the subject of the link's context."
	case IanaLinkRelationsAcl:
		return "Asserts that the link target provides an access control description for the link context."
	case IanaLinkRelationsAlternate:
		return "Refers to a substitute for this context"
	case IanaLinkRelationsAmphtml:
		return "Used to reference alternative content that uses the AMP profile of the HTML format."
	case IanaLinkRelationsAppendix:
		return "Refers to an appendix."
	case IanaLinkRelationsAppleTouchIcon:
		return "Refers to an icon for the context. Synonym for icon."
	case IanaLinkRelationsAppleTouchStartupImage:
		return "Refers to a launch screen for the context."
	case IanaLinkRelationsArchives:
		return "Refers to a collection of records, documents, or other materials of historical interest."
	case IanaLinkRelationsAuthor:
		return "Refers to the context's author."
	case IanaLinkRelationsBlockedBy:
		return "Identifies the entity that blocks access to a resource following receipt of a legal demand."
	case IanaLinkRelationsBookmark:
		return "Gives a permanent link to use for bookmarking purposes."
	case IanaLinkRelationsCanonical:
		return "Designates the preferred version of a resource (the IRI and its contents)."
	case IanaLinkRelationsChapter:
		return "Refers to a chapter in a collection of resources."
	case IanaLinkRelationsCiteAs:
		return "Indicates that the link target is preferred over the link context for the purpose of permanent citation."
	case IanaLinkRelationsCollection:
		return "The target IRI points to a resource which represents the collection resource for the context IRI."
	case IanaLinkRelationsContents:
		return "Refers to a table of contents."
	case IanaLinkRelationsConvertedfrom:
		return "The document linked to was later converted to the document that contains this link relation. For example, an RFC can have a link to the Internet-Draft that became the RFC; in that case, the link relation would be \"convertedFrom\"."
	case IanaLinkRelationsCopyright:
		return "Refers to a copyright statement that applies to the link's context."
	case IanaLinkRelationsCreateForm:
		return "The target IRI points to a resource where a submission form can be obtained."
	case IanaLinkRelationsCurrent:
		return "Refers to a resource containing the most recent item(s) in a collection of resources."
	case IanaLinkRelationsDescribedby:
		return "Refers to a resource providing information about the link's context."
	case IanaLinkRelationsDescribes:
		return "The relationship A 'describes' B asserts that resource A provides a description of resource B. There are no constraints on the format or representation of either A or B, neither are there any further constraints on either resource."
	case IanaLinkRelationsDisclosure:
		return "Refers to a list of patent disclosures made with respect to material for which 'disclosure' relation is specified."
	case IanaLinkRelationsDnsPrefetch:
		return "Used to indicate an origin that will be used to fetch required resources for the link context, and that the user agent ought to resolve as early as possible."
	case IanaLinkRelationsDuplicate:
		return "Refers to a resource whose available representations are byte-for-byte identical with the corresponding representations of the context IRI."
	case IanaLinkRelationsEdit:
		return "Refers to a resource that can be used to edit the link's context."
	case IanaLinkRelationsEditForm:
		return "The target IRI points to a resource where a submission form for editing associated resource can be obtained."
	case IanaLinkRelationsEditMedia:
		return "Refers to a resource that can be used to edit media associated with the link's context."
	case IanaLinkRelationsEnclosure:
		return "Identifies a related resource that is potentially large and might require special handling."
	case IanaLinkRelationsExternal:
		return "Refers to a resource that is not part of the same site as the current context."
	case IanaLinkRelationsFirst:
		return "An IRI that refers to the furthest preceding resource in a series of resources."
	case IanaLinkRelationsGlossary:
		return "Refers to a glossary of terms."
	case IanaLinkRelationsHelp:
		return "Refers to context-sensitive help."
	case IanaLinkRelationsHosts:
		return "Refers to a resource hosted by the server indicated by the link context."
	case IanaLinkRelationsHub:
		return "Refers to a hub that enables registration for notification of updates to the context."
	case IanaLinkRelationsIcon:
		return "Refers to an icon representing the link's context."
	case IanaLinkRelationsIndex:
		return "Refers to an index."
	case IanaLinkRelationsIntervalafter:
		return "refers to a resource associated with a time interval that ends before the beginning of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalbefore:
		return "refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalcontains:
		return "refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervaldisjoint:
		return "refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource, or ends before the beginning of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalduring:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalequals:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalfinishedby:
		return "refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalfinishes:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalin:
		return "refers to a resource associated with a time interval that begins before or is coincident with the beginning of the time interval associated with the context resource, and ends after or is coincident with the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalmeets:
		return "refers to a resource associated with a time interval whose beginning coincides with the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalmetby:
		return "refers to a resource associated with a time interval whose end coincides with the beginning of the time interval associated with the context resource"
	case IanaLinkRelationsIntervaloverlappedby:
		return "refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the beginning of the time interval associated with the context resource"
	case IanaLinkRelationsIntervaloverlaps:
		return "refers to a resource associated with a time interval that begins before the end of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalstartedby:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource"
	case IanaLinkRelationsIntervalstarts:
		return "refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource"
	case IanaLinkRelationsItem:
		return "The target IRI points to a resource that is a member of the collection represented by the context IRI."
	case IanaLinkRelationsLast:
		return "An IRI that refers to the furthest following resource in a series of resources."
	case IanaLinkRelationsLatestVersion:
		return "Points to a resource containing the latest (e.g., current) version of the context."
	case IanaLinkRelationsLicense:
		return "Refers to a license associated with this context."
	case IanaLinkRelationsLinkset:
		return "The link target of a link with the \"linkset\" relation type provides a set of links, including links in which the link context of the link participates."
	case IanaLinkRelationsLrdd:
		return "Refers to further information about the link's context, expressed as a LRDD (\"Link-based Resource Descriptor Document\") resource. See for information about processing this relation type in host-meta documents. When used elsewhere, it refers to additional links and other metadata. Multiple instances indicate additional LRDD resources. LRDD resources MUST have an \"application/xrd+xml\" representation, and MAY have others."
	case IanaLinkRelationsManifest:
		return "Links to a manifest file for the context."
	case IanaLinkRelationsMaskIcon:
		return "Refers to a mask that can be applied to the icon for the context."
	case IanaLinkRelationsMediaFeed:
		return "Refers to a feed of personalised media recommendations relevant to the link context."
	case IanaLinkRelationsMemento:
		return "The Target IRI points to a Memento, a fixed resource that will not change state anymore."
	case IanaLinkRelationsMicropub:
		return "Links to the context's Micropub endpoint."
	case IanaLinkRelationsModulepreload:
		return "Refers to a module that the user agent is to preemptively fetch and store for use in the current context."
	case IanaLinkRelationsMonitor:
		return "Refers to a resource that can be used to monitor changes in an HTTP resource."
	case IanaLinkRelationsMonitorGroup:
		return "Refers to a resource that can be used to monitor changes in a specified group of HTTP resources."
	case IanaLinkRelationsNext:
		return "Indicates that the link's context is a part of a series, and that the next in the series is the link target."
	case IanaLinkRelationsNextArchive:
		return "Refers to the immediately following archive resource."
	case IanaLinkRelationsNofollow:
		return "Indicates that the context’s original author or publisher does not endorse the link target."
	case IanaLinkRelationsNoopener:
		return "Indicates that any newly created top-level browsing context which results from following the link will not be an auxiliary browsing context."
	case IanaLinkRelationsNoreferrer:
		return "Indicates that no referrer information is to be leaked when following the link."
	case IanaLinkRelationsOpener:
		return "Indicates that any newly created top-level browsing context which results from following the link will be an auxiliary browsing context."
	case IanaLinkRelationsOpenid2LocalId:
		return "Refers to an OpenID Authentication server on which the context relies for an assertion that the end user controls an Identifier."
	case IanaLinkRelationsOpenid2Provider:
		return "Refers to a resource which accepts OpenID Authentication protocol messages for the context."
	case IanaLinkRelationsOriginal:
		return "The Target IRI points to an Original Resource."
	case IanaLinkRelationsP3Pv1:
		return "Refers to a P3P privacy policy for the context."
	case IanaLinkRelationsPayment:
		return "Indicates a resource where payment is accepted."
	case IanaLinkRelationsPingback:
		return "Gives the address of the pingback resource for the link context."
	case IanaLinkRelationsPreconnect:
		return "Used to indicate an origin that will be used to fetch required resources for the link context. Initiating an early connection, which includes the DNS lookup, TCP handshake, and optional TLS negotiation, allows the user agent to mask the high latency costs of establishing a connection."
	case IanaLinkRelationsPredecessorVersion:
		return "Points to a resource containing the predecessor version in the version history."
	case IanaLinkRelationsPrefetch:
		return "The prefetch link relation type is used to identify a resource that might be required by the next navigation from the link context, and that the user agent ought to fetch, such that the user agent can deliver a faster response once the resource is requested in the future."
	case IanaLinkRelationsPreload:
		return "Refers to a resource that should be loaded early in the processing of the link's context, without blocking rendering."
	case IanaLinkRelationsPrerender:
		return "Used to identify a resource that might be required by the next navigation from the link context, and that the user agent ought to fetch and execute, such that the user agent can deliver a faster response once the resource is requested in the future."
	case IanaLinkRelationsPrev:
		return "Indicates that the link's context is a part of a series, and that the previous in the series is the link target."
	case IanaLinkRelationsPreview:
		return "Refers to a resource that provides a preview of the link's context."
	case IanaLinkRelationsPrevious:
		return "Refers to the previous resource in an ordered series of resources. Synonym for \"prev\"."
	case IanaLinkRelationsPrevArchive:
		return "Refers to the immediately preceding archive resource."
	case IanaLinkRelationsPrivacyPolicy:
		return "Refers to a privacy policy associated with the link's context."
	case IanaLinkRelationsProfile:
		return "Identifying that a resource representation conforms to a certain profile, without affecting the non-profile semantics of the resource representation."
	case IanaLinkRelationsPublication:
		return "Links to a publication manifest. A manifest represents structured information about a publication, such as informative metadata, a list of resources, and a default reading order."
	case IanaLinkRelationsRelated:
		return "Identifies a related resource."
	case IanaLinkRelationsRestconf:
		return "Identifies the root of RESTCONF API as configured on this HTTP server. The \"restconf\" relation defines the root of the API defined in RFC8040. Subsequent revisions of RESTCONF will use alternate relation values to support protocol versioning."
	case IanaLinkRelationsReplies:
		return "Identifies a resource that is a reply to the context of the link."
	case IanaLinkRelationsRuleinput:
		return "The resource identified by the link target provides an input value to an instance of a rule, where the resource which represents the rule instance is identified by the link context."
	case IanaLinkRelationsSearch:
		return "Refers to a resource that can be used to search through the link's context and related resources."
	case IanaLinkRelationsSection:
		return "Refers to a section in a collection of resources."
	case IanaLinkRelationsSelf:
		return "Conveys an identifier for the link's context."
	case IanaLinkRelationsService:
		return "Indicates a URI that can be used to retrieve a service document."
	case IanaLinkRelationsServiceDesc:
		return "Identifies service description for the context that is primarily intended for consumption by machines."
	case IanaLinkRelationsServiceDoc:
		return "Identifies service documentation for the context that is primarily intended for human consumption."
	case IanaLinkRelationsServiceMeta:
		return "Identifies general metadata for the context that is primarily intended for consumption by machines."
	case IanaLinkRelationsSponsored:
		return "Refers to a resource that is within a context that is sponsored (such as advertising or another compensation agreement)."
	case IanaLinkRelationsStart:
		return "Refers to the first resource in a collection of resources."
	case IanaLinkRelationsStatus:
		return "Identifies a resource that represents the context's status."
	case IanaLinkRelationsStylesheet:
		return "Refers to a stylesheet."
	case IanaLinkRelationsSubsection:
		return "Refers to a resource serving as a subsection in a collection of resources."
	case IanaLinkRelationsSuccessorVersion:
		return "Points to a resource containing the successor version in the version history."
	case IanaLinkRelationsSunset:
		return "Identifies a resource that provides information about the context's retirement policy."
	case IanaLinkRelationsTag:
		return "Gives a tag (identified by the given address) that applies to the current document."
	case IanaLinkRelationsTermsOfService:
		return "Refers to the terms of service associated with the link's context."
	case IanaLinkRelationsTimegate:
		return "The Target IRI points to a TimeGate for an Original Resource."
	case IanaLinkRelationsTimemap:
		return "The Target IRI points to a TimeMap for an Original Resource."
	case IanaLinkRelationsType:
		return "Refers to a resource identifying the abstract semantic type of which the link's context is considered to be an instance."
	case IanaLinkRelationsUgc:
		return "Refers to a resource that is within a context that is User Generated Content."
	case IanaLinkRelationsUp:
		return "Refers to a parent document in a hierarchy of documents."
	case IanaLinkRelationsVersionHistory:
		return "Points to a resource containing the version history for the context."
	case IanaLinkRelationsVia:
		return "Identifies a resource that is the source of the information in the link's context."
	case IanaLinkRelationsWebmention:
		return "Identifies a target URI that supports the Webmention protocol. This allows clients that mention a resource in some form of publishing process to contact that endpoint and inform it that this resource has been mentioned."
	case IanaLinkRelationsWorkingCopy:
		return "Points to a working copy for this resource."
	case IanaLinkRelationsWorkingCopyOf:
		return "Points to the versioned resource from which this working copy was obtained."
	default:
		return "Unknown Constraint Severity"
	}
}
