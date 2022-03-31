// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.CloudFunctions.V1
{
    /// <summary>
    /// The log type that this config enables.
    /// </summary>
    [EnumType]
    public readonly struct AuditLogConfigLogType : IEquatable<AuditLogConfigLogType>
    {
        private readonly string _value;

        private AuditLogConfigLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default case. Should never be this.
        /// </summary>
        public static AuditLogConfigLogType LogTypeUnspecified { get; } = new AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED");
        /// <summary>
        /// Admin reads. Example: CloudIAM getIamPolicy
        /// </summary>
        public static AuditLogConfigLogType AdminRead { get; } = new AuditLogConfigLogType("ADMIN_READ");
        /// <summary>
        /// Data writes. Example: CloudSQL Users create
        /// </summary>
        public static AuditLogConfigLogType DataWrite { get; } = new AuditLogConfigLogType("DATA_WRITE");
        /// <summary>
        /// Data reads. Example: CloudSQL Users list
        /// </summary>
        public static AuditLogConfigLogType DataRead { get; } = new AuditLogConfigLogType("DATA_READ");

        public static bool operator ==(AuditLogConfigLogType left, AuditLogConfigLogType right) => left.Equals(right);
        public static bool operator !=(AuditLogConfigLogType left, AuditLogConfigLogType right) => !left.Equals(right);

        public static explicit operator string(AuditLogConfigLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuditLogConfigLogType other && Equals(other);
        public bool Equals(AuditLogConfigLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Docker Registry to use for this deployment. If `docker_repository` field is specified, this field will be automatically set as `ARTIFACT_REGISTRY`. If unspecified, it currently defaults to `CONTAINER_REGISTRY`. This field may be overridden by the backend for eligible deployments.
    /// </summary>
    [EnumType]
    public readonly struct FunctionDockerRegistry : IEquatable<FunctionDockerRegistry>
    {
        private readonly string _value;

        private FunctionDockerRegistry(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified.
        /// </summary>
        public static FunctionDockerRegistry DockerRegistryUnspecified { get; } = new FunctionDockerRegistry("DOCKER_REGISTRY_UNSPECIFIED");
        /// <summary>
        /// Docker images will be stored in multi-regional Container Registry repositories named `gcf`.
        /// </summary>
        public static FunctionDockerRegistry ContainerRegistry { get; } = new FunctionDockerRegistry("CONTAINER_REGISTRY");
        /// <summary>
        /// Docker images will be stored in regional Artifact Registry repositories. By default, GCF will create and use repositories named `gcf-artifacts` in every region in which a function is deployed. But the repository to use can also be specified by the user using the `docker_repository` field.
        /// </summary>
        public static FunctionDockerRegistry ArtifactRegistry { get; } = new FunctionDockerRegistry("ARTIFACT_REGISTRY");

        public static bool operator ==(FunctionDockerRegistry left, FunctionDockerRegistry right) => left.Equals(right);
        public static bool operator !=(FunctionDockerRegistry left, FunctionDockerRegistry right) => !left.Equals(right);

        public static explicit operator string(FunctionDockerRegistry value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FunctionDockerRegistry other && Equals(other);
        public bool Equals(FunctionDockerRegistry other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The ingress settings for the function, controlling what traffic can reach it.
    /// </summary>
    [EnumType]
    public readonly struct FunctionIngressSettings : IEquatable<FunctionIngressSettings>
    {
        private readonly string _value;

        private FunctionIngressSettings(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified.
        /// </summary>
        public static FunctionIngressSettings IngressSettingsUnspecified { get; } = new FunctionIngressSettings("INGRESS_SETTINGS_UNSPECIFIED");
        /// <summary>
        /// Allow HTTP traffic from public and private sources.
        /// </summary>
        public static FunctionIngressSettings AllowAll { get; } = new FunctionIngressSettings("ALLOW_ALL");
        /// <summary>
        /// Allow HTTP traffic from only private VPC sources.
        /// </summary>
        public static FunctionIngressSettings AllowInternalOnly { get; } = new FunctionIngressSettings("ALLOW_INTERNAL_ONLY");
        /// <summary>
        /// Allow HTTP traffic from private VPC sources and through GCLB.
        /// </summary>
        public static FunctionIngressSettings AllowInternalAndGclb { get; } = new FunctionIngressSettings("ALLOW_INTERNAL_AND_GCLB");

        public static bool operator ==(FunctionIngressSettings left, FunctionIngressSettings right) => left.Equals(right);
        public static bool operator !=(FunctionIngressSettings left, FunctionIngressSettings right) => !left.Equals(right);

        public static explicit operator string(FunctionIngressSettings value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FunctionIngressSettings other && Equals(other);
        public bool Equals(FunctionIngressSettings other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The egress settings for the connector, controlling what traffic is diverted through it.
    /// </summary>
    [EnumType]
    public readonly struct FunctionVpcConnectorEgressSettings : IEquatable<FunctionVpcConnectorEgressSettings>
    {
        private readonly string _value;

        private FunctionVpcConnectorEgressSettings(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified.
        /// </summary>
        public static FunctionVpcConnectorEgressSettings VpcConnectorEgressSettingsUnspecified { get; } = new FunctionVpcConnectorEgressSettings("VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED");
        /// <summary>
        /// Use the VPC Access Connector only for private IP space from RFC1918.
        /// </summary>
        public static FunctionVpcConnectorEgressSettings PrivateRangesOnly { get; } = new FunctionVpcConnectorEgressSettings("PRIVATE_RANGES_ONLY");
        /// <summary>
        /// Force the use of VPC Access Connector for all egress traffic from the function.
        /// </summary>
        public static FunctionVpcConnectorEgressSettings AllTraffic { get; } = new FunctionVpcConnectorEgressSettings("ALL_TRAFFIC");

        public static bool operator ==(FunctionVpcConnectorEgressSettings left, FunctionVpcConnectorEgressSettings right) => left.Equals(right);
        public static bool operator !=(FunctionVpcConnectorEgressSettings left, FunctionVpcConnectorEgressSettings right) => !left.Equals(right);

        public static explicit operator string(FunctionVpcConnectorEgressSettings value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FunctionVpcConnectorEgressSettings other && Equals(other);
        public bool Equals(FunctionVpcConnectorEgressSettings other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The security level for the function.
    /// </summary>
    [EnumType]
    public readonly struct HttpsTriggerSecurityLevel : IEquatable<HttpsTriggerSecurityLevel>
    {
        private readonly string _value;

        private HttpsTriggerSecurityLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified.
        /// </summary>
        public static HttpsTriggerSecurityLevel SecurityLevelUnspecified { get; } = new HttpsTriggerSecurityLevel("SECURITY_LEVEL_UNSPECIFIED");
        /// <summary>
        /// Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect.
        /// </summary>
        public static HttpsTriggerSecurityLevel SecureAlways { get; } = new HttpsTriggerSecurityLevel("SECURE_ALWAYS");
        /// <summary>
        /// Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly.
        /// </summary>
        public static HttpsTriggerSecurityLevel SecureOptional { get; } = new HttpsTriggerSecurityLevel("SECURE_OPTIONAL");

        public static bool operator ==(HttpsTriggerSecurityLevel left, HttpsTriggerSecurityLevel right) => left.Equals(right);
        public static bool operator !=(HttpsTriggerSecurityLevel left, HttpsTriggerSecurityLevel right) => !left.Equals(right);

        public static explicit operator string(HttpsTriggerSecurityLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HttpsTriggerSecurityLevel other && Equals(other);
        public bool Equals(HttpsTriggerSecurityLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
