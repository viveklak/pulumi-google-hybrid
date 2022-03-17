// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CloudRunConfigLoadBalancerType = {
    /**
     * Load balancer type for Cloud Run is unspecified.
     */
    LoadBalancerTypeUnspecified: "LOAD_BALANCER_TYPE_UNSPECIFIED",
    /**
     * Install external load balancer for Cloud Run.
     */
    LoadBalancerTypeExternal: "LOAD_BALANCER_TYPE_EXTERNAL",
    /**
     * Install internal load balancer for Cloud Run.
     */
    LoadBalancerTypeInternal: "LOAD_BALANCER_TYPE_INTERNAL",
} as const;

/**
 * Which load balancer type is installed for Cloud Run.
 */
export type CloudRunConfigLoadBalancerType = (typeof CloudRunConfigLoadBalancerType)[keyof typeof CloudRunConfigLoadBalancerType];

export const ClusterAutoscalingAutoscalingProfile = {
    /**
     * No change to autoscaling configuration.
     */
    ProfileUnspecified: "PROFILE_UNSPECIFIED",
    /**
     * Prioritize optimizing utilization of resources.
     */
    OptimizeUtilization: "OPTIMIZE_UTILIZATION",
    /**
     * Use default (balanced) autoscaling configuration.
     */
    Balanced: "BALANCED",
} as const;

/**
 * Defines autoscaling behaviour.
 */
export type ClusterAutoscalingAutoscalingProfile = (typeof ClusterAutoscalingAutoscalingProfile)[keyof typeof ClusterAutoscalingAutoscalingProfile];

export const ClusterTelemetryType = {
    /**
     * Not set.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * Monitoring integration is disabled.
     */
    Disabled: "DISABLED",
    /**
     * Monitoring integration is enabled.
     */
    Enabled: "ENABLED",
    /**
     * Only system components are monitored and logged.
     */
    SystemOnly: "SYSTEM_ONLY",
} as const;

/**
 * Type of the integration.
 */
export type ClusterTelemetryType = (typeof ClusterTelemetryType)[keyof typeof ClusterTelemetryType];

export const ClusterUpdateDesiredDatapathProvider = {
    /**
     * Default value.
     */
    DatapathProviderUnspecified: "DATAPATH_PROVIDER_UNSPECIFIED",
    /**
     * Use the IPTables implementation based on kube-proxy.
     */
    LegacyDatapath: "LEGACY_DATAPATH",
    /**
     * Use the eBPF based GKE Dataplane V2 with additional features. See the [GKE Dataplane V2 documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/dataplane-v2) for more.
     */
    AdvancedDatapath: "ADVANCED_DATAPATH",
} as const;

/**
 * The desired datapath provider for the cluster.
 */
export type ClusterUpdateDesiredDatapathProvider = (typeof ClusterUpdateDesiredDatapathProvider)[keyof typeof ClusterUpdateDesiredDatapathProvider];

export const ClusterUpdateDesiredPrivateIpv6GoogleAccess = {
    /**
     * Default value. Same as DISABLED
     */
    PrivateIpv6GoogleAccessUnspecified: "PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED",
    /**
     * No private access to or from Google Services
     */
    PrivateIpv6GoogleAccessDisabled: "PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED",
    /**
     * Enables private IPv6 access to Google Services from GKE
     */
    PrivateIpv6GoogleAccessToGoogle: "PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE",
    /**
     * Enables priate IPv6 access to and from Google Services
     */
    PrivateIpv6GoogleAccessBidirectional: "PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL",
} as const;

/**
 * The desired state of IPv6 connectivity to Google Services.
 */
export type ClusterUpdateDesiredPrivateIpv6GoogleAccess = (typeof ClusterUpdateDesiredPrivateIpv6GoogleAccess)[keyof typeof ClusterUpdateDesiredPrivateIpv6GoogleAccess];

export const DNSConfigClusterDns = {
    /**
     * Default value
     */
    ProviderUnspecified: "PROVIDER_UNSPECIFIED",
    /**
     * Use GKE default DNS provider(kube-dns) for DNS resolution.
     */
    PlatformDefault: "PLATFORM_DEFAULT",
    /**
     * Use CloudDNS for DNS resolution.
     */
    CloudDns: "CLOUD_DNS",
} as const;

/**
 * cluster_dns indicates which in-cluster DNS provider should be used.
 */
export type DNSConfigClusterDns = (typeof DNSConfigClusterDns)[keyof typeof DNSConfigClusterDns];

export const DNSConfigClusterDnsScope = {
    /**
     * Default value, will be inferred as cluster scope.
     */
    DnsScopeUnspecified: "DNS_SCOPE_UNSPECIFIED",
    /**
     * DNS records are accessible from within the cluster.
     */
    ClusterScope: "CLUSTER_SCOPE",
    /**
     * DNS records are accessible from within the VPC.
     */
    VpcScope: "VPC_SCOPE",
} as const;

/**
 * cluster_dns_scope indicates the scope of access to cluster DNS records.
 */
export type DNSConfigClusterDnsScope = (typeof DNSConfigClusterDnsScope)[keyof typeof DNSConfigClusterDnsScope];

export const DatabaseEncryptionState = {
    /**
     * Should never be set
     */
    Unknown: "UNKNOWN",
    /**
     * Secrets in etcd are encrypted.
     */
    Encrypted: "ENCRYPTED",
    /**
     * Secrets in etcd are stored in plain text (at etcd level) - this is unrelated to Compute Engine level full disk encryption.
     */
    Decrypted: "DECRYPTED",
} as const;

/**
 * Denotes the state of etcd encryption.
 */
export type DatabaseEncryptionState = (typeof DatabaseEncryptionState)[keyof typeof DatabaseEncryptionState];

export const FilterEventTypeItem = {
    /**
     * Not set, will be ignored.
     */
    EventTypeUnspecified: "EVENT_TYPE_UNSPECIFIED",
    /**
     * Corresponds with UpgradeAvailableEvent.
     */
    UpgradeAvailableEvent: "UPGRADE_AVAILABLE_EVENT",
    /**
     * Corresponds with UpgradeEvent.
     */
    UpgradeEvent: "UPGRADE_EVENT",
    /**
     * Corresponds with SecurityBulletinEvent.
     */
    SecurityBulletinEvent: "SECURITY_BULLETIN_EVENT",
} as const;

export type FilterEventTypeItem = (typeof FilterEventTypeItem)[keyof typeof FilterEventTypeItem];

export const IstioConfigAuth = {
    /**
     * auth not enabled
     */
    AuthNone: "AUTH_NONE",
    /**
     * auth mutual TLS enabled
     */
    AuthMutualTls: "AUTH_MUTUAL_TLS",
} as const;

/**
 * The specified Istio auth mode, either none, or mutual TLS.
 */
export type IstioConfigAuth = (typeof IstioConfigAuth)[keyof typeof IstioConfigAuth];

export const LoggingComponentConfigEnableComponentsItem = {
    /**
     * Default value. This shouldn't be used.
     */
    ComponentUnspecified: "COMPONENT_UNSPECIFIED",
    /**
     * system components
     */
    SystemComponents: "SYSTEM_COMPONENTS",
    /**
     * workloads
     */
    Workloads: "WORKLOADS",
} as const;

export type LoggingComponentConfigEnableComponentsItem = (typeof LoggingComponentConfigEnableComponentsItem)[keyof typeof LoggingComponentConfigEnableComponentsItem];

export const MaintenanceExclusionOptionsScope = {
    /**
     * NO_UPGRADES excludes all upgrades, including patch upgrades and minor upgrades across control planes and nodes. This is the default exclusion behavior.
     */
    NoUpgrades: "NO_UPGRADES",
    /**
     * NO_MINOR_UPGRADES excludes all minor upgrades for the cluster, only patches are allowed.
     */
    NoMinorUpgrades: "NO_MINOR_UPGRADES",
    /**
     * NO_MINOR_OR_NODE_UPGRADES excludes all minor upgrades for the cluster, and also exclude all node pool upgrades. Only control plane patches are allowed.
     */
    NoMinorOrNodeUpgrades: "NO_MINOR_OR_NODE_UPGRADES",
} as const;

/**
 * Scope specifies the upgrade scope which upgrades are blocked by the exclusion.
 */
export type MaintenanceExclusionOptionsScope = (typeof MaintenanceExclusionOptionsScope)[keyof typeof MaintenanceExclusionOptionsScope];

export const MonitoringComponentConfigEnableComponentsItem = {
    /**
     * Default value. This shouldn't be used.
     */
    ComponentUnspecified: "COMPONENT_UNSPECIFIED",
    /**
     * system components
     */
    SystemComponents: "SYSTEM_COMPONENTS",
    /**
     * workloads
     */
    Workloads: "WORKLOADS",
} as const;

export type MonitoringComponentConfigEnableComponentsItem = (typeof MonitoringComponentConfigEnableComponentsItem)[keyof typeof MonitoringComponentConfigEnableComponentsItem];

export const NetworkConfigDatapathProvider = {
    /**
     * Default value.
     */
    DatapathProviderUnspecified: "DATAPATH_PROVIDER_UNSPECIFIED",
    /**
     * Use the IPTables implementation based on kube-proxy.
     */
    LegacyDatapath: "LEGACY_DATAPATH",
    /**
     * Use the eBPF based GKE Dataplane V2 with additional features. See the [GKE Dataplane V2 documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/dataplane-v2) for more.
     */
    AdvancedDatapath: "ADVANCED_DATAPATH",
} as const;

/**
 * The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation.
 */
export type NetworkConfigDatapathProvider = (typeof NetworkConfigDatapathProvider)[keyof typeof NetworkConfigDatapathProvider];

export const NetworkConfigPrivateIpv6GoogleAccess = {
    /**
     * Default value. Same as DISABLED
     */
    PrivateIpv6GoogleAccessUnspecified: "PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED",
    /**
     * No private access to or from Google Services
     */
    PrivateIpv6GoogleAccessDisabled: "PRIVATE_IPV6_GOOGLE_ACCESS_DISABLED",
    /**
     * Enables private IPv6 access to Google Services from GKE
     */
    PrivateIpv6GoogleAccessToGoogle: "PRIVATE_IPV6_GOOGLE_ACCESS_TO_GOOGLE",
    /**
     * Enables priate IPv6 access to and from Google Services
     */
    PrivateIpv6GoogleAccessBidirectional: "PRIVATE_IPV6_GOOGLE_ACCESS_BIDIRECTIONAL",
} as const;

/**
 * The desired state of IPv6 connectivity to Google Services. By default, no private IPv6 access to or from Google Services (all access will be via IPv4)
 */
export type NetworkConfigPrivateIpv6GoogleAccess = (typeof NetworkConfigPrivateIpv6GoogleAccess)[keyof typeof NetworkConfigPrivateIpv6GoogleAccess];

export const NetworkPolicyProvider = {
    /**
     * Not set
     */
    ProviderUnspecified: "PROVIDER_UNSPECIFIED",
    /**
     * Tigera (Calico Felix).
     */
    Calico: "CALICO",
} as const;

/**
 * The selected network policy provider.
 */
export type NetworkPolicyProvider = (typeof NetworkPolicyProvider)[keyof typeof NetworkPolicyProvider];

export const NodeTaintEffect = {
    /**
     * Not set
     */
    EffectUnspecified: "EFFECT_UNSPECIFIED",
    /**
     * NoSchedule
     */
    NoSchedule: "NO_SCHEDULE",
    /**
     * PreferNoSchedule
     */
    PreferNoSchedule: "PREFER_NO_SCHEDULE",
    /**
     * NoExecute
     */
    NoExecute: "NO_EXECUTE",
} as const;

/**
 * Effect for taint.
 */
export type NodeTaintEffect = (typeof NodeTaintEffect)[keyof typeof NodeTaintEffect];

export const PlacementPolicyType = {
    /**
     * TYPE_UNSPECIFIED specifies no requirements on nodes placement.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * COMPACT specifies node placement in the same availability domain to ensure low communication latency.
     */
    Compact: "COMPACT",
} as const;

/**
 * The type of placement.
 */
export type PlacementPolicyType = (typeof PlacementPolicyType)[keyof typeof PlacementPolicyType];

export const ReleaseChannelChannel = {
    /**
     * No channel specified.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * RAPID channel is offered on an early access basis for customers who want to test new releases. WARNING: Versions available in the RAPID Channel may be subject to unresolved issues with no known workaround and are not subject to any SLAs.
     */
    Rapid: "RAPID",
    /**
     * Clusters subscribed to REGULAR receive versions that are considered GA quality. REGULAR is intended for production users who want to take advantage of new features.
     */
    Regular: "REGULAR",
    /**
     * Clusters subscribed to STABLE receive versions that are known to be stable and reliable in production.
     */
    Stable: "STABLE",
} as const;

/**
 * channel specifies which release channel the cluster is subscribed to.
 */
export type ReleaseChannelChannel = (typeof ReleaseChannelChannel)[keyof typeof ReleaseChannelChannel];

export const ReservationAffinityConsumeReservationType = {
    /**
     * Default value. This should not be used.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * Do not consume from any reserved capacity.
     */
    NoReservation: "NO_RESERVATION",
    /**
     * Consume any reservation available.
     */
    AnyReservation: "ANY_RESERVATION",
    /**
     * Must consume from a specific reservation. Must specify key value fields for specifying the reservations.
     */
    SpecificReservation: "SPECIFIC_RESERVATION",
} as const;

/**
 * Corresponds to the type of reservation consumption.
 */
export type ReservationAffinityConsumeReservationType = (typeof ReservationAffinityConsumeReservationType)[keyof typeof ReservationAffinityConsumeReservationType];

export const SandboxConfigType = {
    /**
     * Default value. This should not be used.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * Run sandbox using gvisor.
     */
    Gvisor: "GVISOR",
} as const;

/**
 * Type of the sandbox to use for the node.
 */
export type SandboxConfigType = (typeof SandboxConfigType)[keyof typeof SandboxConfigType];

export const StatusConditionCanonicalCode = {
    /**
     * Not an error; returned on success HTTP Mapping: 200 OK
     */
    Ok: "OK",
    /**
     * The operation was cancelled, typically by the caller. HTTP Mapping: 499 Client Closed Request
     */
    Cancelled: "CANCELLED",
    /**
     * Unknown error. For example, this error may be returned when a `Status` value received from another address space belongs to an error space that is not known in this address space. Also errors raised by APIs that do not return enough error information may be converted to this error. HTTP Mapping: 500 Internal Server Error
     */
    Unknown: "UNKNOWN",
    /**
     * The client specified an invalid argument. Note that this differs from `FAILED_PRECONDITION`. `INVALID_ARGUMENT` indicates arguments that are problematic regardless of the state of the system (e.g., a malformed file name). HTTP Mapping: 400 Bad Request
     */
    InvalidArgument: "INVALID_ARGUMENT",
    /**
     * The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully. For example, a successful response from a server could have been delayed long enough for the deadline to expire. HTTP Mapping: 504 Gateway Timeout
     */
    DeadlineExceeded: "DEADLINE_EXCEEDED",
    /**
     * Some requested entity (e.g., file or directory) was not found. Note to server developers: if a request is denied for an entire class of users, such as gradual feature rollout or undocumented allowlist, `NOT_FOUND` may be used. If a request is denied for some users within a class of users, such as user-based access control, `PERMISSION_DENIED` must be used. HTTP Mapping: 404 Not Found
     */
    NotFound: "NOT_FOUND",
    /**
     * The entity that a client attempted to create (e.g., file or directory) already exists. HTTP Mapping: 409 Conflict
     */
    AlreadyExists: "ALREADY_EXISTS",
    /**
     * The caller does not have permission to execute the specified operation. `PERMISSION_DENIED` must not be used for rejections caused by exhausting some resource (use `RESOURCE_EXHAUSTED` instead for those errors). `PERMISSION_DENIED` must not be used if the caller can not be identified (use `UNAUTHENTICATED` instead for those errors). This error code does not imply the request is valid or the requested entity exists or satisfies other pre-conditions. HTTP Mapping: 403 Forbidden
     */
    PermissionDenied: "PERMISSION_DENIED",
    /**
     * The request does not have valid authentication credentials for the operation. HTTP Mapping: 401 Unauthorized
     */
    Unauthenticated: "UNAUTHENTICATED",
    /**
     * Some resource has been exhausted, perhaps a per-user quota, or perhaps the entire file system is out of space. HTTP Mapping: 429 Too Many Requests
     */
    ResourceExhausted: "RESOURCE_EXHAUSTED",
    /**
     * The operation was rejected because the system is not in a state required for the operation's execution. For example, the directory to be deleted is non-empty, an rmdir operation is applied to a non-directory, etc. Service implementors can use the following guidelines to decide between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`: (a) Use `UNAVAILABLE` if the client can retry just the failing call. (b) Use `ABORTED` if the client should retry at a higher level. For example, when a client-specified test-and-set fails, indicating the client should restart a read-modify-write sequence. (c) Use `FAILED_PRECONDITION` if the client should not retry until the system state has been explicitly fixed. For example, if an "rmdir" fails because the directory is non-empty, `FAILED_PRECONDITION` should be returned since the client should not retry unless the files are deleted from the directory. HTTP Mapping: 400 Bad Request
     */
    FailedPrecondition: "FAILED_PRECONDITION",
    /**
     * The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort. See the guidelines above for deciding between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`. HTTP Mapping: 409 Conflict
     */
    Aborted: "ABORTED",
    /**
     * The operation was attempted past the valid range. E.g., seeking or reading past end-of-file. Unlike `INVALID_ARGUMENT`, this error indicates a problem that may be fixed if the system state changes. For example, a 32-bit file system will generate `INVALID_ARGUMENT` if asked to read at an offset that is not in the range [0,2^32-1], but it will generate `OUT_OF_RANGE` if asked to read from an offset past the current file size. There is a fair bit of overlap between `FAILED_PRECONDITION` and `OUT_OF_RANGE`. We recommend using `OUT_OF_RANGE` (the more specific error) when it applies so that callers who are iterating through a space can easily look for an `OUT_OF_RANGE` error to detect when they are done. HTTP Mapping: 400 Bad Request
     */
    OutOfRange: "OUT_OF_RANGE",
    /**
     * The operation is not implemented or is not supported/enabled in this service. HTTP Mapping: 501 Not Implemented
     */
    Unimplemented: "UNIMPLEMENTED",
    /**
     * Internal errors. This means that some invariants expected by the underlying system have been broken. This error code is reserved for serious errors. HTTP Mapping: 500 Internal Server Error
     */
    Internal: "INTERNAL",
    /**
     * The service is currently unavailable. This is most likely a transient condition, which can be corrected by retrying with a backoff. Note that it is not always safe to retry non-idempotent operations. See the guidelines above for deciding between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`. HTTP Mapping: 503 Service Unavailable
     */
    Unavailable: "UNAVAILABLE",
    /**
     * Unrecoverable data loss or corruption. HTTP Mapping: 500 Internal Server Error
     */
    DataLoss: "DATA_LOSS",
} as const;

/**
 * Canonical code of the condition.
 */
export type StatusConditionCanonicalCode = (typeof StatusConditionCanonicalCode)[keyof typeof StatusConditionCanonicalCode];

export const StatusConditionCode = {
    /**
     * UNKNOWN indicates a generic condition.
     */
    Unknown: "UNKNOWN",
    /**
     * GCE_STOCKOUT indicates that Google Compute Engine resources are temporarily unavailable.
     */
    GceStockout: "GCE_STOCKOUT",
    /**
     * GKE_SERVICE_ACCOUNT_DELETED indicates that the user deleted their robot service account.
     */
    GkeServiceAccountDeleted: "GKE_SERVICE_ACCOUNT_DELETED",
    /**
     * Google Compute Engine quota was exceeded.
     */
    GceQuotaExceeded: "GCE_QUOTA_EXCEEDED",
    /**
     * Cluster state was manually changed by an SRE due to a system logic error.
     */
    SetByOperator: "SET_BY_OPERATOR",
    /**
     * Unable to perform an encrypt operation against the CloudKMS key used for etcd level encryption.
     */
    CloudKmsKeyError: "CLOUD_KMS_KEY_ERROR",
    /**
     * Cluster CA is expiring soon. More codes TBA
     */
    CaExpiring: "CA_EXPIRING",
} as const;

/**
 * Machine-friendly representation of the condition Deprecated. Use canonical_code instead.
 */
export type StatusConditionCode = (typeof StatusConditionCode)[keyof typeof StatusConditionCode];

export const WorkloadMetadataConfigMode = {
    /**
     * Not set.
     */
    ModeUnspecified: "MODE_UNSPECIFIED",
    /**
     * Expose all Compute Engine metadata to pods.
     */
    GceMetadata: "GCE_METADATA",
    /**
     * Run the GKE Metadata Server on this node. The GKE Metadata Server exposes a metadata API to workloads that is compatible with the V1 Compute Metadata APIs exposed by the Compute Engine and App Engine Metadata Servers. This feature can only be enabled if Workload Identity is enabled at the cluster level.
     */
    GkeMetadata: "GKE_METADATA",
} as const;

/**
 * Mode is the configuration for how to expose metadata to workloads running on the node pool.
 */
export type WorkloadMetadataConfigMode = (typeof WorkloadMetadataConfigMode)[keyof typeof WorkloadMetadataConfigMode];

export const WorkloadMetadataConfigNodeMetadata = {
    /**
     * Not set.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * Prevent workloads not in hostNetwork from accessing certain VM metadata, specifically kube-env, which contains Kubelet credentials, and the instance identity token. Metadata concealment is a temporary security solution available while the bootstrapping process for cluster nodes is being redesigned with significant security improvements. This feature is scheduled to be deprecated in the future and later removed.
     */
    Secure: "SECURE",
    /**
     * Expose all VM metadata to pods.
     */
    Expose: "EXPOSE",
    /**
     * Run the GKE Metadata Server on this node. The GKE Metadata Server exposes a metadata API to workloads that is compatible with the V1 Compute Metadata APIs exposed by the Compute Engine and App Engine Metadata Servers. This feature can only be enabled if Workload Identity is enabled at the cluster level.
     */
    GkeMetadataServer: "GKE_METADATA_SERVER",
} as const;

/**
 * NodeMetadata is the configuration for how to expose metadata to the workloads running on the node.
 */
export type WorkloadMetadataConfigNodeMetadata = (typeof WorkloadMetadataConfigNodeMetadata)[keyof typeof WorkloadMetadataConfigNodeMetadata];
