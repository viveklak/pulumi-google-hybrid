// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.NetworkManagement.V1Beta1
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
    /// Type of the network where the endpoint is located. Applicable only to source endpoint, as destination network type can be inferred from the source.
    /// </summary>
    [EnumType]
    public readonly struct EndpointNetworkType : IEquatable<EndpointNetworkType>
    {
        private readonly string _value;

        private EndpointNetworkType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default type if unspecified.
        /// </summary>
        public static EndpointNetworkType NetworkTypeUnspecified { get; } = new EndpointNetworkType("NETWORK_TYPE_UNSPECIFIED");
        /// <summary>
        /// A network hosted within Google Cloud Platform. To receive more detailed output, specify the URI for the source or destination network.
        /// </summary>
        public static EndpointNetworkType GcpNetwork { get; } = new EndpointNetworkType("GCP_NETWORK");
        /// <summary>
        /// A network hosted outside of Google Cloud Platform. This can be an on-premises network, or a network hosted by another cloud provider.
        /// </summary>
        public static EndpointNetworkType NonGcpNetwork { get; } = new EndpointNetworkType("NON_GCP_NETWORK");

        public static bool operator ==(EndpointNetworkType left, EndpointNetworkType right) => left.Equals(right);
        public static bool operator !=(EndpointNetworkType left, EndpointNetworkType right) => !left.Equals(right);

        public static explicit operator string(EndpointNetworkType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EndpointNetworkType other && Equals(other);
        public bool Equals(EndpointNetworkType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
