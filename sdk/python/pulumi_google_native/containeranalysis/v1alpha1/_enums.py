# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'BuildSignatureKeyType',
    'CVSSAttackComplexity',
    'CVSSAttackVector',
    'CVSSAuthentication',
    'CVSSAvailabilityImpact',
    'CVSSConfidentialityImpact',
    'CVSSIntegrityImpact',
    'CVSSPrivilegesRequired',
    'CVSSScope',
    'CVSSUserInteraction',
    'CisBenchmarkSeverity',
    'DeploymentPlatform',
    'DiscoveredAnalysisStatus',
    'DiscoveredContinuousAnalysis',
    'DiscoveryAnalysisKind',
    'DistributionArchitecture',
    'ExternalRefCategory',
    'FileNoteFileType',
    'GoogleDevtoolsContaineranalysisV1alpha1AliasContextKind',
    'HashType',
    'LayerDirective',
    'PgpSignedAttestationContentType',
    'RelationshipNoteType',
    'VersionKind',
    'VulnerabilityDetailsEffectiveSeverity',
    'VulnerabilityTypeSeverity',
]


class BuildSignatureKeyType(str, Enum):
    """
    The type of the key, either stored in `public_key` or referenced in `key_id`
    """
    KEY_TYPE_UNSPECIFIED = "KEY_TYPE_UNSPECIFIED"
    """
    `KeyType` is not set.
    """
    PGP_ASCII_ARMORED = "PGP_ASCII_ARMORED"
    """
    `PGP ASCII Armored` public key.
    """
    PKIX_PEM = "PKIX_PEM"
    """
    `PKIX PEM` public key.
    """


class CVSSAttackComplexity(str, Enum):
    ATTACK_COMPLEXITY_UNSPECIFIED = "ATTACK_COMPLEXITY_UNSPECIFIED"
    ATTACK_COMPLEXITY_LOW = "ATTACK_COMPLEXITY_LOW"
    ATTACK_COMPLEXITY_HIGH = "ATTACK_COMPLEXITY_HIGH"


class CVSSAttackVector(str, Enum):
    """
    Base Metrics Represents the intrinsic characteristics of a vulnerability that are constant over time and across user environments.
    """
    ATTACK_VECTOR_UNSPECIFIED = "ATTACK_VECTOR_UNSPECIFIED"
    ATTACK_VECTOR_NETWORK = "ATTACK_VECTOR_NETWORK"
    ATTACK_VECTOR_ADJACENT = "ATTACK_VECTOR_ADJACENT"
    ATTACK_VECTOR_LOCAL = "ATTACK_VECTOR_LOCAL"
    ATTACK_VECTOR_PHYSICAL = "ATTACK_VECTOR_PHYSICAL"


class CVSSAuthentication(str, Enum):
    AUTHENTICATION_UNSPECIFIED = "AUTHENTICATION_UNSPECIFIED"
    AUTHENTICATION_MULTIPLE = "AUTHENTICATION_MULTIPLE"
    AUTHENTICATION_SINGLE = "AUTHENTICATION_SINGLE"
    AUTHENTICATION_NONE = "AUTHENTICATION_NONE"


class CVSSAvailabilityImpact(str, Enum):
    IMPACT_UNSPECIFIED = "IMPACT_UNSPECIFIED"
    IMPACT_HIGH = "IMPACT_HIGH"
    IMPACT_LOW = "IMPACT_LOW"
    IMPACT_NONE = "IMPACT_NONE"


class CVSSConfidentialityImpact(str, Enum):
    IMPACT_UNSPECIFIED = "IMPACT_UNSPECIFIED"
    IMPACT_HIGH = "IMPACT_HIGH"
    IMPACT_LOW = "IMPACT_LOW"
    IMPACT_NONE = "IMPACT_NONE"


class CVSSIntegrityImpact(str, Enum):
    IMPACT_UNSPECIFIED = "IMPACT_UNSPECIFIED"
    IMPACT_HIGH = "IMPACT_HIGH"
    IMPACT_LOW = "IMPACT_LOW"
    IMPACT_NONE = "IMPACT_NONE"


class CVSSPrivilegesRequired(str, Enum):
    PRIVILEGES_REQUIRED_UNSPECIFIED = "PRIVILEGES_REQUIRED_UNSPECIFIED"
    PRIVILEGES_REQUIRED_NONE = "PRIVILEGES_REQUIRED_NONE"
    PRIVILEGES_REQUIRED_LOW = "PRIVILEGES_REQUIRED_LOW"
    PRIVILEGES_REQUIRED_HIGH = "PRIVILEGES_REQUIRED_HIGH"


class CVSSScope(str, Enum):
    SCOPE_UNSPECIFIED = "SCOPE_UNSPECIFIED"
    SCOPE_UNCHANGED = "SCOPE_UNCHANGED"
    SCOPE_CHANGED = "SCOPE_CHANGED"


class CVSSUserInteraction(str, Enum):
    USER_INTERACTION_UNSPECIFIED = "USER_INTERACTION_UNSPECIFIED"
    USER_INTERACTION_NONE = "USER_INTERACTION_NONE"
    USER_INTERACTION_REQUIRED = "USER_INTERACTION_REQUIRED"


class CisBenchmarkSeverity(str, Enum):
    """
    The severity level of this CIS benchmark check.
    """
    SEVERITY_UNSPECIFIED = "SEVERITY_UNSPECIFIED"
    """
    Unknown Impact
    """
    MINIMAL = "MINIMAL"
    """
    Minimal Impact
    """
    LOW = "LOW"
    """
    Low Impact
    """
    MEDIUM = "MEDIUM"
    """
    Medium Impact
    """
    HIGH = "HIGH"
    """
    High Impact
    """
    CRITICAL = "CRITICAL"
    """
    Critical Impact
    """


class DeploymentPlatform(str, Enum):
    """
    Platform hosting this deployment.
    """
    PLATFORM_UNSPECIFIED = "PLATFORM_UNSPECIFIED"
    """
    Unknown
    """
    GKE = "GKE"
    """
    Google Container Engine
    """
    FLEX = "FLEX"
    """
    Google App Engine: Flexible Environment
    """
    CUSTOM = "CUSTOM"
    """
    Custom user-defined platform
    """


class DiscoveredAnalysisStatus(str, Enum):
    """
    The status of discovery for the resource.
    """
    ANALYSIS_STATUS_UNSPECIFIED = "ANALYSIS_STATUS_UNSPECIFIED"
    """
    Unknown
    """
    PENDING = "PENDING"
    """
    Resource is known but no action has been taken yet.
    """
    SCANNING = "SCANNING"
    """
    Resource is being analyzed.
    """
    FINISHED_SUCCESS = "FINISHED_SUCCESS"
    """
    Analysis has finished successfully.
    """
    FINISHED_FAILED = "FINISHED_FAILED"
    """
    Analysis has finished unsuccessfully, the analysis itself is in a bad state.
    """
    FINISHED_UNSUPPORTED = "FINISHED_UNSUPPORTED"
    """
    The resource is known not to be supported.
    """


class DiscoveredContinuousAnalysis(str, Enum):
    """
    Whether the resource is continuously analyzed.
    """
    CONTINUOUS_ANALYSIS_UNSPECIFIED = "CONTINUOUS_ANALYSIS_UNSPECIFIED"
    """
    Unknown
    """
    ACTIVE = "ACTIVE"
    """
    The resource is continuously analyzed.
    """
    INACTIVE = "INACTIVE"
    """
    The resource is ignored for continuous analysis.
    """


class DiscoveryAnalysisKind(str, Enum):
    """
    The kind of analysis that is handled by this discovery.
    """
    KIND_UNSPECIFIED = "KIND_UNSPECIFIED"
    """
    Unknown
    """
    PACKAGE_VULNERABILITY = "PACKAGE_VULNERABILITY"
    """
    The note and occurrence represent a package vulnerability.
    """
    BUILD_DETAILS = "BUILD_DETAILS"
    """
    The note and occurrence assert build provenance.
    """
    IMAGE_BASIS = "IMAGE_BASIS"
    """
    This represents an image basis relationship.
    """
    PACKAGE_MANAGER = "PACKAGE_MANAGER"
    """
    This represents a package installed via a package manager.
    """
    DEPLOYABLE = "DEPLOYABLE"
    """
    The note and occurrence track deployment events.
    """
    DISCOVERY = "DISCOVERY"
    """
    The note and occurrence track the initial discovery status of a resource.
    """
    ATTESTATION_AUTHORITY = "ATTESTATION_AUTHORITY"
    """
    This represents a logical "role" that can attest to artifacts.
    """
    UPGRADE = "UPGRADE"
    """
    This represents an available software upgrade.
    """
    COMPLIANCE = "COMPLIANCE"
    """
    This represents a compliance check that can be applied to a resource.
    """
    SBOM = "SBOM"
    """
    This represents a software bill of materials.
    """
    SPDX_PACKAGE = "SPDX_PACKAGE"
    """
    This represents an SPDX Package.
    """
    SPDX_FILE = "SPDX_FILE"
    """
    This represents an SPDX File.
    """
    SPDX_RELATIONSHIP = "SPDX_RELATIONSHIP"
    """
    This represents an SPDX Relationship.
    """
    DSSE_ATTESTATION = "DSSE_ATTESTATION"
    """
    This represents a DSSE attestation Note
    """


class DistributionArchitecture(str, Enum):
    """
    The CPU architecture for which packages in this distribution channel were built
    """
    ARCHITECTURE_UNSPECIFIED = "ARCHITECTURE_UNSPECIFIED"
    """
    Unknown architecture
    """
    X86 = "X86"
    """
    X86 architecture
    """
    X64 = "X64"
    """
    X64 architecture
    """


class ExternalRefCategory(str, Enum):
    """
    An External Reference allows a Package to reference an external source of additional information, metadata, enumerations, asset identifiers, or downloadable content believed to be relevant to the Package
    """
    CATEGORY_UNSPECIFIED = "CATEGORY_UNSPECIFIED"
    """
    Unspecified
    """
    SECURITY = "SECURITY"
    """
    Security (e.g. cpe22Type, cpe23Type)
    """
    PACKAGE_MANAGER = "PACKAGE_MANAGER"
    """
    Package Manager (e.g. maven-central, npm, nuget, bower, purl)
    """
    PERSISTENT_ID = "PERSISTENT_ID"
    """
    Persistent-Id (e.g. swh)
    """
    OTHER = "OTHER"
    """
    Other
    """


class FileNoteFileType(str, Enum):
    """
    This field provides information about the type of file identified
    """
    FILE_TYPE_UNSPECIFIED = "FILE_TYPE_UNSPECIFIED"
    """
    Unspecified
    """
    SOURCE = "SOURCE"
    """
    The file is human readable source code (.c, .html, etc.)
    """
    BINARY = "BINARY"
    """
    The file is a compiled object, target image or binary executable (.o, .a, etc.)
    """
    ARCHIVE = "ARCHIVE"
    """
    The file represents an archive (.tar, .jar, etc.)
    """
    APPLICATION = "APPLICATION"
    """
    The file is associated with a specific application type (MIME type of application/*)
    """
    AUDIO = "AUDIO"
    """
    The file is associated with an audio file (MIME type of audio/* , e.g. .mp3)
    """
    IMAGE = "IMAGE"
    """
    The file is associated with an picture image file (MIME type of image/*, e.g., .jpg, .gif)
    """
    TEXT = "TEXT"
    """
    The file is human readable text file (MIME type of text/*)
    """
    VIDEO = "VIDEO"
    """
    The file is associated with a video file type (MIME type of video/*)
    """
    DOCUMENTATION = "DOCUMENTATION"
    """
    The file serves as documentation
    """
    SPDX = "SPDX"
    """
    The file is an SPDX document
    """
    OTHER = "OTHER"
    """
    The file doesn't fit into the above categories (generated artifacts, data files, etc.)
    """


class GoogleDevtoolsContaineranalysisV1alpha1AliasContextKind(str, Enum):
    """
    The alias kind.
    """
    KIND_UNSPECIFIED = "KIND_UNSPECIFIED"
    """
    Unknown.
    """
    FIXED = "FIXED"
    """
    Git tag.
    """
    MOVABLE = "MOVABLE"
    """
    Git branch.
    """
    OTHER = "OTHER"
    """
    Used to specify non-standard aliases. For example, if a Git repo has a ref named "refs/foo/bar".
    """


class HashType(str, Enum):
    """
    The type of hash that was performed.
    """
    NONE = "NONE"
    """
    No hash requested.
    """
    SHA256 = "SHA256"
    """
    A sha256 hash.
    """


class LayerDirective(str, Enum):
    """
    The recovered Dockerfile directive used to construct this layer.
    """
    DIRECTIVE_UNSPECIFIED = "DIRECTIVE_UNSPECIFIED"
    """
    Default value for unsupported/missing directive
    """
    MAINTAINER = "MAINTAINER"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    RUN = "RUN"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    CMD = "CMD"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    LABEL = "LABEL"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    EXPOSE = "EXPOSE"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    ENV = "ENV"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    ADD = "ADD"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    COPY = "COPY"
    """
    https://docs.docker.com/reference/builder/#copy
    """
    ENTRYPOINT = "ENTRYPOINT"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    VOLUME = "VOLUME"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    USER = "USER"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    WORKDIR = "WORKDIR"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    ARG = "ARG"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    ONBUILD = "ONBUILD"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    STOPSIGNAL = "STOPSIGNAL"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    HEALTHCHECK = "HEALTHCHECK"
    """
    https://docs.docker.com/engine/reference/builder/
    """
    SHELL = "SHELL"
    """
    https://docs.docker.com/engine/reference/builder/
    """


class PgpSignedAttestationContentType(str, Enum):
    """
    Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
    """
    CONTENT_TYPE_UNSPECIFIED = "CONTENT_TYPE_UNSPECIFIED"
    """
    `ContentType` is not set.
    """
    SIMPLE_SIGNING_JSON = "SIMPLE_SIGNING_JSON"
    """
    Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted from `signature` is a JSON blob conforming to the linked schema.
    """


class RelationshipNoteType(str, Enum):
    """
    The type of relationship between the source and target SPDX elements
    """
    RELATIONSHIP_TYPE_UNSPECIFIED = "RELATIONSHIP_TYPE_UNSPECIFIED"
    """
    Unspecified
    """
    DESCRIBES = "DESCRIBES"
    """
    Is to be used when SPDXRef-DOCUMENT describes SPDXRef-A
    """
    DESCRIBED_BY = "DESCRIBED_BY"
    """
    Is to be used when SPDXRef-A is described by SPDXREF-Document
    """
    CONTAINS = "CONTAINS"
    """
    Is to be used when SPDXRef-A contains SPDXRef-B
    """
    CONTAINED_BY = "CONTAINED_BY"
    """
    Is to be used when SPDXRef-A is contained by SPDXRef-B
    """
    DEPENDS_ON = "DEPENDS_ON"
    """
    Is to be used when SPDXRef-A depends on SPDXRef-B
    """
    DEPENDENCY_OF = "DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is dependency of SPDXRef-B
    """
    DEPENDENCY_MANIFEST_OF = "DEPENDENCY_MANIFEST_OF"
    """
    Is to be used when SPDXRef-A is a manifest file that lists a set of dependencies for SPDXRef-B
    """
    BUILD_DEPENDENCY_OF = "BUILD_DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is a build dependency of SPDXRef-B
    """
    DEV_DEPENDENCY_OF = "DEV_DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is a development dependency of SPDXRef-B
    """
    OPTIONAL_DEPENDENCY_OF = "OPTIONAL_DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is an optional dependency of SPDXRef-B
    """
    PROVIDED_DEPENDENCY_OF = "PROVIDED_DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is a to be provided dependency of SPDXRef-B
    """
    TEST_DEPENDENCY_OF = "TEST_DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is a test dependency of SPDXRef-B
    """
    RUNTIME_DEPENDENCY_OF = "RUNTIME_DEPENDENCY_OF"
    """
    Is to be used when SPDXRef-A is a dependency required for the execution of SPDXRef-B
    """
    EXAMPLE_OF = "EXAMPLE_OF"
    """
    Is to be used when SPDXRef-A is an example of SPDXRef-B
    """
    GENERATES = "GENERATES"
    """
    Is to be used when SPDXRef-A generates SPDXRef-B
    """
    GENERATED_FROM = "GENERATED_FROM"
    """
    Is to be used when SPDXRef-A was generated from SPDXRef-B
    """
    ANCESTOR_OF = "ANCESTOR_OF"
    """
    Is to be used when SPDXRef-A is an ancestor (same lineage but pre-dates) SPDXRef-B
    """
    DESCENDANT_OF = "DESCENDANT_OF"
    """
    Is to be used when SPDXRef-A is a descendant of (same lineage but postdates) SPDXRef-B
    """
    VARIANT_OF = "VARIANT_OF"
    """
    Is to be used when SPDXRef-A is a variant of (same lineage but not clear which came first) SPDXRef-B
    """
    DISTRIBUTION_ARTIFACT = "DISTRIBUTION_ARTIFACT"
    """
    Is to be used when distributing SPDXRef-A requires that SPDXRef-B also be distributed
    """
    PATCH_FOR = "PATCH_FOR"
    """
    Is to be used when SPDXRef-A is a patch file for (to be applied to) SPDXRef-B
    """
    PATCH_APPLIED = "PATCH_APPLIED"
    """
    Is to be used when SPDXRef-A is a patch file that has been applied to SPDXRef-B
    """
    COPY_OF = "COPY_OF"
    """
    Is to be used when SPDXRef-A is an exact copy of SPDXRef-B
    """
    FILE_ADDED = "FILE_ADDED"
    """
    Is to be used when SPDXRef-A is a file that was added to SPDXRef-B
    """
    FILE_DELETED = "FILE_DELETED"
    """
    Is to be used when SPDXRef-A is a file that was deleted from SPDXRef-B
    """
    FILE_MODIFIED = "FILE_MODIFIED"
    """
    Is to be used when SPDXRef-A is a file that was modified from SPDXRef-B
    """
    EXPANDED_FROM_ARCHIVE = "EXPANDED_FROM_ARCHIVE"
    """
    Is to be used when SPDXRef-A is expanded from the archive SPDXRef-B
    """
    DYNAMIC_LINK = "DYNAMIC_LINK"
    """
    Is to be used when SPDXRef-A dynamically links to SPDXRef-B
    """
    STATIC_LINK = "STATIC_LINK"
    """
    Is to be used when SPDXRef-A statically links to SPDXRef-B
    """
    DATA_FILE_OF = "DATA_FILE_OF"
    """
    Is to be used when SPDXRef-A is a data file used in SPDXRef-B
    """
    TEST_CASE_OF = "TEST_CASE_OF"
    """
    Is to be used when SPDXRef-A is a test case used in testing SPDXRef-B
    """
    BUILD_TOOL_OF = "BUILD_TOOL_OF"
    """
    Is to be used when SPDXRef-A is used to build SPDXRef-B
    """
    DEV_TOOL_OF = "DEV_TOOL_OF"
    """
    Is to be used when SPDXRef-A is used as a development tool for SPDXRef-B
    """
    TEST_OF = "TEST_OF"
    """
    Is to be used when SPDXRef-A is used for testing SPDXRef-B
    """
    TEST_TOOL_OF = "TEST_TOOL_OF"
    """
    Is to be used when SPDXRef-A is used as a test tool for SPDXRef-B
    """
    DOCUMENTATION_OF = "DOCUMENTATION_OF"
    """
    Is to be used when SPDXRef-A provides documentation of SPDXRef-B
    """
    OPTIONAL_COMPONENT_OF = "OPTIONAL_COMPONENT_OF"
    """
    Is to be used when SPDXRef-A is an optional component of SPDXRef-B
    """
    METAFILE_OF = "METAFILE_OF"
    """
    Is to be used when SPDXRef-A is a metafile of SPDXRef-B
    """
    PACKAGE_OF = "PACKAGE_OF"
    """
    Is to be used when SPDXRef-A is used as a package as part of SPDXRef-B
    """
    AMENDS = "AMENDS"
    """
    Is to be used when (current) SPDXRef-DOCUMENT amends the SPDX information in SPDXRef-B
    """
    PREREQUISITE_FOR = "PREREQUISITE_FOR"
    """
    Is to be used when SPDXRef-A is a prerequisite for SPDXRef-B
    """
    HAS_PREREQUISITE = "HAS_PREREQUISITE"
    """
    Is to be used when SPDXRef-A has as a prerequisite SPDXRef-B
    """
    OTHER = "OTHER"
    """
    Is to be used for a relationship which has not been defined in the formal SPDX specification. A description of the relationship should be included in the Relationship comments field
    """


class VersionKind(str, Enum):
    """
    Distinguish between sentinel MIN/MAX versions and normal versions. If kind is not NORMAL, then the other fields are ignored.
    """
    NORMAL = "NORMAL"
    """
    A standard package version, defined by the other fields.
    """
    MINIMUM = "MINIMUM"
    """
    A special version representing negative infinity, other fields are ignored.
    """
    MAXIMUM = "MAXIMUM"
    """
    A special version representing positive infinity, other fields are ignored.
    """


class VulnerabilityDetailsEffectiveSeverity(str, Enum):
    """
    The distro assigned severity for this vulnerability when that is available and note provider assigned severity when distro has not yet assigned a severity for this vulnerability. When there are multiple package issues for this vulnerability, they can have different effective severities because some might come from the distro and some might come from installed language packs (e.g. Maven JARs or Go binaries). For this reason, it is advised to use the effective severity on the PackageIssue level, as this field may eventually be deprecated. In the case where multiple PackageIssues have different effective severities, the one set here will be the highest severity of any of the PackageIssues.
    """
    SEVERITY_UNSPECIFIED = "SEVERITY_UNSPECIFIED"
    """
    Unknown Impact
    """
    MINIMAL = "MINIMAL"
    """
    Minimal Impact
    """
    LOW = "LOW"
    """
    Low Impact
    """
    MEDIUM = "MEDIUM"
    """
    Medium Impact
    """
    HIGH = "HIGH"
    """
    High Impact
    """
    CRITICAL = "CRITICAL"
    """
    Critical Impact
    """


class VulnerabilityTypeSeverity(str, Enum):
    """
    Note provider assigned impact of the vulnerability
    """
    SEVERITY_UNSPECIFIED = "SEVERITY_UNSPECIFIED"
    """
    Unknown Impact
    """
    MINIMAL = "MINIMAL"
    """
    Minimal Impact
    """
    LOW = "LOW"
    """
    Low Impact
    """
    MEDIUM = "MEDIUM"
    """
    Medium Impact
    """
    HIGH = "HIGH"
    """
    High Impact
    """
    CRITICAL = "CRITICAL"
    """
    Critical Impact
    """
