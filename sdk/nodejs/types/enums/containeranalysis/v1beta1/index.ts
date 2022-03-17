// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AliasContextKind = {
    /**
     * Unknown.
     */
    KindUnspecified: "KIND_UNSPECIFIED",
    /**
     * Git tag.
     */
    Fixed: "FIXED",
    /**
     * Git branch.
     */
    Movable: "MOVABLE",
    /**
     * Used to specify non-standard aliases. For example, if a Git repo has a ref named "refs/foo/bar".
     */
    Other: "OTHER",
} as const;

/**
 * The alias kind.
 */
export type AliasContextKind = (typeof AliasContextKind)[keyof typeof AliasContextKind];

export const BuildSignatureKeyType = {
    /**
     * `KeyType` is not set.
     */
    KeyTypeUnspecified: "KEY_TYPE_UNSPECIFIED",
    /**
     * `PGP ASCII Armored` public key.
     */
    PgpAsciiArmored: "PGP_ASCII_ARMORED",
    /**
     * `PKIX PEM` public key.
     */
    PkixPem: "PKIX_PEM",
} as const;

/**
 * The type of the key, either stored in `public_key` or referenced in `key_id`.
 */
export type BuildSignatureKeyType = (typeof BuildSignatureKeyType)[keyof typeof BuildSignatureKeyType];

export const CVSSv3AttackComplexity = {
    AttackComplexityUnspecified: "ATTACK_COMPLEXITY_UNSPECIFIED",
    AttackComplexityLow: "ATTACK_COMPLEXITY_LOW",
    AttackComplexityHigh: "ATTACK_COMPLEXITY_HIGH",
} as const;

export type CVSSv3AttackComplexity = (typeof CVSSv3AttackComplexity)[keyof typeof CVSSv3AttackComplexity];

export const CVSSv3AttackVector = {
    AttackVectorUnspecified: "ATTACK_VECTOR_UNSPECIFIED",
    AttackVectorNetwork: "ATTACK_VECTOR_NETWORK",
    AttackVectorAdjacent: "ATTACK_VECTOR_ADJACENT",
    AttackVectorLocal: "ATTACK_VECTOR_LOCAL",
    AttackVectorPhysical: "ATTACK_VECTOR_PHYSICAL",
} as const;

/**
 * Base Metrics Represents the intrinsic characteristics of a vulnerability that are constant over time and across user environments.
 */
export type CVSSv3AttackVector = (typeof CVSSv3AttackVector)[keyof typeof CVSSv3AttackVector];

export const CVSSv3AvailabilityImpact = {
    ImpactUnspecified: "IMPACT_UNSPECIFIED",
    ImpactHigh: "IMPACT_HIGH",
    ImpactLow: "IMPACT_LOW",
    ImpactNone: "IMPACT_NONE",
} as const;

export type CVSSv3AvailabilityImpact = (typeof CVSSv3AvailabilityImpact)[keyof typeof CVSSv3AvailabilityImpact];

export const CVSSv3ConfidentialityImpact = {
    ImpactUnspecified: "IMPACT_UNSPECIFIED",
    ImpactHigh: "IMPACT_HIGH",
    ImpactLow: "IMPACT_LOW",
    ImpactNone: "IMPACT_NONE",
} as const;

export type CVSSv3ConfidentialityImpact = (typeof CVSSv3ConfidentialityImpact)[keyof typeof CVSSv3ConfidentialityImpact];

export const CVSSv3IntegrityImpact = {
    ImpactUnspecified: "IMPACT_UNSPECIFIED",
    ImpactHigh: "IMPACT_HIGH",
    ImpactLow: "IMPACT_LOW",
    ImpactNone: "IMPACT_NONE",
} as const;

export type CVSSv3IntegrityImpact = (typeof CVSSv3IntegrityImpact)[keyof typeof CVSSv3IntegrityImpact];

export const CVSSv3PrivilegesRequired = {
    PrivilegesRequiredUnspecified: "PRIVILEGES_REQUIRED_UNSPECIFIED",
    PrivilegesRequiredNone: "PRIVILEGES_REQUIRED_NONE",
    PrivilegesRequiredLow: "PRIVILEGES_REQUIRED_LOW",
    PrivilegesRequiredHigh: "PRIVILEGES_REQUIRED_HIGH",
} as const;

export type CVSSv3PrivilegesRequired = (typeof CVSSv3PrivilegesRequired)[keyof typeof CVSSv3PrivilegesRequired];

export const CVSSv3Scope = {
    ScopeUnspecified: "SCOPE_UNSPECIFIED",
    ScopeUnchanged: "SCOPE_UNCHANGED",
    ScopeChanged: "SCOPE_CHANGED",
} as const;

export type CVSSv3Scope = (typeof CVSSv3Scope)[keyof typeof CVSSv3Scope];

export const CVSSv3UserInteraction = {
    UserInteractionUnspecified: "USER_INTERACTION_UNSPECIFIED",
    UserInteractionNone: "USER_INTERACTION_NONE",
    UserInteractionRequired: "USER_INTERACTION_REQUIRED",
} as const;

export type CVSSv3UserInteraction = (typeof CVSSv3UserInteraction)[keyof typeof CVSSv3UserInteraction];

export const DeploymentPlatform = {
    /**
     * Unknown.
     */
    PlatformUnspecified: "PLATFORM_UNSPECIFIED",
    /**
     * Google Container Engine.
     */
    Gke: "GKE",
    /**
     * Google App Engine: Flexible Environment.
     */
    Flex: "FLEX",
    /**
     * Custom user-defined platform.
     */
    Custom: "CUSTOM",
} as const;

/**
 * Platform hosting this deployment.
 */
export type DeploymentPlatform = (typeof DeploymentPlatform)[keyof typeof DeploymentPlatform];

export const DiscoveredAnalysisStatus = {
    /**
     * Unknown.
     */
    AnalysisStatusUnspecified: "ANALYSIS_STATUS_UNSPECIFIED",
    /**
     * Resource is known but no action has been taken yet.
     */
    Pending: "PENDING",
    /**
     * Resource is being analyzed.
     */
    Scanning: "SCANNING",
    /**
     * Analysis has finished successfully.
     */
    FinishedSuccess: "FINISHED_SUCCESS",
    /**
     * Analysis has finished unsuccessfully, the analysis itself is in a bad state.
     */
    FinishedFailed: "FINISHED_FAILED",
    /**
     * The resource is known not to be supported
     */
    FinishedUnsupported: "FINISHED_UNSUPPORTED",
} as const;

/**
 * The status of discovery for the resource.
 */
export type DiscoveredAnalysisStatus = (typeof DiscoveredAnalysisStatus)[keyof typeof DiscoveredAnalysisStatus];

export const DiscoveredContinuousAnalysis = {
    /**
     * Unknown.
     */
    ContinuousAnalysisUnspecified: "CONTINUOUS_ANALYSIS_UNSPECIFIED",
    /**
     * The resource is continuously analyzed.
     */
    Active: "ACTIVE",
    /**
     * The resource is ignored for continuous analysis.
     */
    Inactive: "INACTIVE",
} as const;

/**
 * Whether the resource is continuously analyzed.
 */
export type DiscoveredContinuousAnalysis = (typeof DiscoveredContinuousAnalysis)[keyof typeof DiscoveredContinuousAnalysis];

export const DiscoveryAnalysisKind = {
    /**
     * Default value. This value is unused.
     */
    NoteKindUnspecified: "NOTE_KIND_UNSPECIFIED",
    /**
     * The note and occurrence represent a package vulnerability.
     */
    Vulnerability: "VULNERABILITY",
    /**
     * The note and occurrence assert build provenance.
     */
    Build: "BUILD",
    /**
     * This represents an image basis relationship.
     */
    Image: "IMAGE",
    /**
     * This represents a package installed via a package manager.
     */
    Package: "PACKAGE",
    /**
     * The note and occurrence track deployment events.
     */
    Deployment: "DEPLOYMENT",
    /**
     * The note and occurrence track the initial discovery status of a resource.
     */
    Discovery: "DISCOVERY",
    /**
     * This represents a logical "role" that can attest to artifacts.
     */
    Attestation: "ATTESTATION",
    /**
     * This represents an in-toto link.
     */
    Intoto: "INTOTO",
    /**
     * This represents a software bill of materials.
     */
    Sbom: "SBOM",
    /**
     * This represents an SPDX Package.
     */
    SpdxPackage: "SPDX_PACKAGE",
    /**
     * This represents an SPDX File.
     */
    SpdxFile: "SPDX_FILE",
    /**
     * This represents an SPDX Relationship.
     */
    SpdxRelationship: "SPDX_RELATIONSHIP",
} as const;

/**
 * Required. Immutable. The kind of analysis that is handled by this discovery.
 */
export type DiscoveryAnalysisKind = (typeof DiscoveryAnalysisKind)[keyof typeof DiscoveryAnalysisKind];

export const DistributionArchitecture = {
    /**
     * Unknown architecture.
     */
    ArchitectureUnspecified: "ARCHITECTURE_UNSPECIFIED",
    /**
     * X86 architecture.
     */
    X86: "X86",
    /**
     * X64 architecture.
     */
    X64: "X64",
} as const;

/**
 * The CPU architecture for which packages in this distribution channel were built.
 */
export type DistributionArchitecture = (typeof DistributionArchitecture)[keyof typeof DistributionArchitecture];

export const ExternalRefCategory = {
    /**
     * Unspecified
     */
    CategoryUnspecified: "CATEGORY_UNSPECIFIED",
    /**
     * Security (e.g. cpe22Type, cpe23Type)
     */
    Security: "SECURITY",
    /**
     * Package Manager (e.g. maven-central, npm, nuget, bower, purl)
     */
    PackageManager: "PACKAGE_MANAGER",
    /**
     * Persistent-Id (e.g. swh)
     */
    PersistentId: "PERSISTENT_ID",
    /**
     * Other
     */
    Other: "OTHER",
} as const;

/**
 * An External Reference allows a Package to reference an external source of additional information, metadata, enumerations, asset identifiers, or downloadable content believed to be relevant to the Package
 */
export type ExternalRefCategory = (typeof ExternalRefCategory)[keyof typeof ExternalRefCategory];

export const FileNoteFileType = {
    /**
     * Unspecified
     */
    FileTypeUnspecified: "FILE_TYPE_UNSPECIFIED",
    /**
     * The file is human readable source code (.c, .html, etc.)
     */
    Source: "SOURCE",
    /**
     * The file is a compiled object, target image or binary executable (.o, .a, etc.)
     */
    Binary: "BINARY",
    /**
     * The file represents an archive (.tar, .jar, etc.)
     */
    Archive: "ARCHIVE",
    /**
     * The file is associated with a specific application type (MIME type of application/*)
     */
    Application: "APPLICATION",
    /**
     * The file is associated with an audio file (MIME type of audio/* , e.g. .mp3)
     */
    Audio: "AUDIO",
    /**
     * The file is associated with an picture image file (MIME type of image/*, e.g., .jpg, .gif)
     */
    Image: "IMAGE",
    /**
     * The file is human readable text file (MIME type of text/*)
     */
    Text: "TEXT",
    /**
     * The file is associated with a video file type (MIME type of video/*)
     */
    Video: "VIDEO",
    /**
     * The file serves as documentation
     */
    Documentation: "DOCUMENTATION",
    /**
     * The file is an SPDX document
     */
    Spdx: "SPDX",
    /**
     * The file doesn't fit into the above categories (generated artifacts, data files, etc.)
     */
    Other: "OTHER",
} as const;

/**
 * This field provides information about the type of file identified
 */
export type FileNoteFileType = (typeof FileNoteFileType)[keyof typeof FileNoteFileType];

export const GenericSignedAttestationContentType = {
    /**
     * `ContentType` is not set.
     */
    ContentTypeUnspecified: "CONTENT_TYPE_UNSPECIFIED",
    /**
     * Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted in `plaintext` is a JSON blob conforming to the linked schema.
     */
    SimpleSigningJson: "SIMPLE_SIGNING_JSON",
} as const;

/**
 * Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
 */
export type GenericSignedAttestationContentType = (typeof GenericSignedAttestationContentType)[keyof typeof GenericSignedAttestationContentType];

export const GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity = {
    /**
     * Unknown.
     */
    SeverityUnspecified: "SEVERITY_UNSPECIFIED",
    /**
     * Minimal severity.
     */
    Minimal: "MINIMAL",
    /**
     * Low severity.
     */
    Low: "LOW",
    /**
     * Medium severity.
     */
    Medium: "MEDIUM",
    /**
     * High severity.
     */
    High: "HIGH",
    /**
     * Critical severity.
     */
    Critical: "CRITICAL",
} as const;

/**
 * The distro assigned severity for this vulnerability when it is available, and note provider assigned severity when distro has not yet assigned a severity for this vulnerability. When there are multiple PackageIssues for this vulnerability, they can have different effective severities because some might be provided by the distro while others are provided by the language ecosystem for a language pack. For this reason, it is advised to use the effective severity on the PackageIssue level. In the case where multiple PackageIssues have differing effective severities, this field should be the highest severity for any of the PackageIssues.
 */
export type GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity = (typeof GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity)[keyof typeof GrafeasV1beta1VulnerabilityDetailsEffectiveSeverity];

export const HashType = {
    /**
     * Unknown.
     */
    HashTypeUnspecified: "HASH_TYPE_UNSPECIFIED",
    /**
     * A SHA-256 hash.
     */
    Sha256: "SHA256",
} as const;

/**
 * Required. The type of hash that was performed.
 */
export type HashType = (typeof HashType)[keyof typeof HashType];

export const LayerDirective = {
    /**
     * Default value for unsupported/missing directive.
     */
    DirectiveUnspecified: "DIRECTIVE_UNSPECIFIED",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Maintainer: "MAINTAINER",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Run: "RUN",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Cmd: "CMD",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Label: "LABEL",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Expose: "EXPOSE",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Env: "ENV",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Add: "ADD",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Copy: "COPY",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Entrypoint: "ENTRYPOINT",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Volume: "VOLUME",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    User: "USER",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Workdir: "WORKDIR",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Arg: "ARG",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Onbuild: "ONBUILD",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Stopsignal: "STOPSIGNAL",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Healthcheck: "HEALTHCHECK",
    /**
     * https://docs.docker.com/engine/reference/builder/
     */
    Shell: "SHELL",
} as const;

/**
 * Required. The recovered Dockerfile directive used to construct this layer.
 */
export type LayerDirective = (typeof LayerDirective)[keyof typeof LayerDirective];

export const PgpSignedAttestationContentType = {
    /**
     * `ContentType` is not set.
     */
    ContentTypeUnspecified: "CONTENT_TYPE_UNSPECIFIED",
    /**
     * Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted from `signature` is a JSON blob conforming to the linked schema.
     */
    SimpleSigningJson: "SIMPLE_SIGNING_JSON",
} as const;

/**
 * Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
 */
export type PgpSignedAttestationContentType = (typeof PgpSignedAttestationContentType)[keyof typeof PgpSignedAttestationContentType];

export const RelationshipNoteType = {
    /**
     * Unspecified
     */
    RelationshipTypeUnspecified: "RELATIONSHIP_TYPE_UNSPECIFIED",
    /**
     * Is to be used when SPDXRef-DOCUMENT describes SPDXRef-A
     */
    Describes: "DESCRIBES",
    /**
     * Is to be used when SPDXRef-A is described by SPDXREF-Document
     */
    DescribedBy: "DESCRIBED_BY",
    /**
     * Is to be used when SPDXRef-A contains SPDXRef-B
     */
    Contains: "CONTAINS",
    /**
     * Is to be used when SPDXRef-A is contained by SPDXRef-B
     */
    ContainedBy: "CONTAINED_BY",
    /**
     * Is to be used when SPDXRef-A depends on SPDXRef-B
     */
    DependsOn: "DEPENDS_ON",
    /**
     * Is to be used when SPDXRef-A is dependency of SPDXRef-B
     */
    DependencyOf: "DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is a manifest file that lists a set of dependencies for SPDXRef-B
     */
    DependencyManifestOf: "DEPENDENCY_MANIFEST_OF",
    /**
     * Is to be used when SPDXRef-A is a build dependency of SPDXRef-B
     */
    BuildDependencyOf: "BUILD_DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is a development dependency of SPDXRef-B
     */
    DevDependencyOf: "DEV_DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is an optional dependency of SPDXRef-B
     */
    OptionalDependencyOf: "OPTIONAL_DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is a to be provided dependency of SPDXRef-B
     */
    ProvidedDependencyOf: "PROVIDED_DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is a test dependency of SPDXRef-B
     */
    TestDependencyOf: "TEST_DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is a dependency required for the execution of SPDXRef-B
     */
    RuntimeDependencyOf: "RUNTIME_DEPENDENCY_OF",
    /**
     * Is to be used when SPDXRef-A is an example of SPDXRef-B
     */
    ExampleOf: "EXAMPLE_OF",
    /**
     * Is to be used when SPDXRef-A generates SPDXRef-B
     */
    Generates: "GENERATES",
    /**
     * Is to be used when SPDXRef-A was generated from SPDXRef-B
     */
    GeneratedFrom: "GENERATED_FROM",
    /**
     * Is to be used when SPDXRef-A is an ancestor (same lineage but pre-dates) SPDXRef-B
     */
    AncestorOf: "ANCESTOR_OF",
    /**
     * Is to be used when SPDXRef-A is a descendant of (same lineage but postdates) SPDXRef-B
     */
    DescendantOf: "DESCENDANT_OF",
    /**
     * Is to be used when SPDXRef-A is a variant of (same lineage but not clear which came first) SPDXRef-B
     */
    VariantOf: "VARIANT_OF",
    /**
     * Is to be used when distributing SPDXRef-A requires that SPDXRef-B also be distributed
     */
    DistributionArtifact: "DISTRIBUTION_ARTIFACT",
    /**
     * Is to be used when SPDXRef-A is a patch file for (to be applied to) SPDXRef-B
     */
    PatchFor: "PATCH_FOR",
    /**
     * Is to be used when SPDXRef-A is a patch file that has been applied to SPDXRef-B
     */
    PatchApplied: "PATCH_APPLIED",
    /**
     * Is to be used when SPDXRef-A is an exact copy of SPDXRef-B
     */
    CopyOf: "COPY_OF",
    /**
     * Is to be used when SPDXRef-A is a file that was added to SPDXRef-B
     */
    FileAdded: "FILE_ADDED",
    /**
     * Is to be used when SPDXRef-A is a file that was deleted from SPDXRef-B
     */
    FileDeleted: "FILE_DELETED",
    /**
     * Is to be used when SPDXRef-A is a file that was modified from SPDXRef-B
     */
    FileModified: "FILE_MODIFIED",
    /**
     * Is to be used when SPDXRef-A is expanded from the archive SPDXRef-B
     */
    ExpandedFromArchive: "EXPANDED_FROM_ARCHIVE",
    /**
     * Is to be used when SPDXRef-A dynamically links to SPDXRef-B
     */
    DynamicLink: "DYNAMIC_LINK",
    /**
     * Is to be used when SPDXRef-A statically links to SPDXRef-B
     */
    StaticLink: "STATIC_LINK",
    /**
     * Is to be used when SPDXRef-A is a data file used in SPDXRef-B
     */
    DataFileOf: "DATA_FILE_OF",
    /**
     * Is to be used when SPDXRef-A is a test case used in testing SPDXRef-B
     */
    TestCaseOf: "TEST_CASE_OF",
    /**
     * Is to be used when SPDXRef-A is used to build SPDXRef-B
     */
    BuildToolOf: "BUILD_TOOL_OF",
    /**
     * Is to be used when SPDXRef-A is used as a development tool for SPDXRef-B
     */
    DevToolOf: "DEV_TOOL_OF",
    /**
     * Is to be used when SPDXRef-A is used for testing SPDXRef-B
     */
    TestOf: "TEST_OF",
    /**
     * Is to be used when SPDXRef-A is used as a test tool for SPDXRef-B
     */
    TestToolOf: "TEST_TOOL_OF",
    /**
     * Is to be used when SPDXRef-A provides documentation of SPDXRef-B
     */
    DocumentationOf: "DOCUMENTATION_OF",
    /**
     * Is to be used when SPDXRef-A is an optional component of SPDXRef-B
     */
    OptionalComponentOf: "OPTIONAL_COMPONENT_OF",
    /**
     * Is to be used when SPDXRef-A is a metafile of SPDXRef-B
     */
    MetafileOf: "METAFILE_OF",
    /**
     * Is to be used when SPDXRef-A is used as a package as part of SPDXRef-B
     */
    PackageOf: "PACKAGE_OF",
    /**
     * Is to be used when (current) SPDXRef-DOCUMENT amends the SPDX information in SPDXRef-B
     */
    Amends: "AMENDS",
    /**
     * Is to be used when SPDXRef-A is a prerequisite for SPDXRef-B
     */
    PrerequisiteFor: "PREREQUISITE_FOR",
    /**
     * Is to be used when SPDXRef-A has as a prerequisite SPDXRef-B
     */
    HasPrerequisite: "HAS_PREREQUISITE",
    /**
     * Is to be used for a relationship which has not been defined in the formal SPDX specification. A description of the relationship should be included in the Relationship comments field
     */
    Other: "OTHER",
} as const;

/**
 * The type of relationship between the source and target SPDX elements
 */
export type RelationshipNoteType = (typeof RelationshipNoteType)[keyof typeof RelationshipNoteType];

export const VersionKind = {
    /**
     * Unknown.
     */
    VersionKindUnspecified: "VERSION_KIND_UNSPECIFIED",
    /**
     * A standard package version.
     */
    Normal: "NORMAL",
    /**
     * A special version representing negative infinity.
     */
    Minimum: "MINIMUM",
    /**
     * A special version representing positive infinity.
     */
    Maximum: "MAXIMUM",
} as const;

/**
 * Required. Distinguishes between sentinel MIN/MAX versions and normal versions.
 */
export type VersionKind = (typeof VersionKind)[keyof typeof VersionKind];

export const VulnerabilitySeverity = {
    /**
     * Unknown.
     */
    SeverityUnspecified: "SEVERITY_UNSPECIFIED",
    /**
     * Minimal severity.
     */
    Minimal: "MINIMAL",
    /**
     * Low severity.
     */
    Low: "LOW",
    /**
     * Medium severity.
     */
    Medium: "MEDIUM",
    /**
     * High severity.
     */
    High: "HIGH",
    /**
     * Critical severity.
     */
    Critical: "CRITICAL",
} as const;

/**
 * Note provider assigned impact of the vulnerability.
 */
export type VulnerabilitySeverity = (typeof VulnerabilitySeverity)[keyof typeof VulnerabilitySeverity];
