// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha
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
    /// Immutable. Specifies CA configuration.
    /// </summary>
    [EnumType]
    public readonly struct FeatureSpecProvisionGoogleCa : IEquatable<FeatureSpecProvisionGoogleCa>
    {
        private readonly string _value;

        private FeatureSpecProvisionGoogleCa(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Disable default Google managed CA.
        /// </summary>
        public static FeatureSpecProvisionGoogleCa GoogleCaProvisioningUnspecified { get; } = new FeatureSpecProvisionGoogleCa("GOOGLE_CA_PROVISIONING_UNSPECIFIED");
        /// <summary>
        /// Disable default Google managed CA.
        /// </summary>
        public static FeatureSpecProvisionGoogleCa Disabled { get; } = new FeatureSpecProvisionGoogleCa("DISABLED");
        /// <summary>
        /// Use default Google managed CA.
        /// </summary>
        public static FeatureSpecProvisionGoogleCa Enabled { get; } = new FeatureSpecProvisionGoogleCa("ENABLED");

        public static bool operator ==(FeatureSpecProvisionGoogleCa left, FeatureSpecProvisionGoogleCa right) => left.Equals(right);
        public static bool operator !=(FeatureSpecProvisionGoogleCa left, FeatureSpecProvisionGoogleCa right) => !left.Equals(right);

        public static explicit operator string(FeatureSpecProvisionGoogleCa value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FeatureSpecProvisionGoogleCa other && Equals(other);
        public bool Equals(FeatureSpecProvisionGoogleCa other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies workload certificate management.
    /// </summary>
    [EnumType]
    public readonly struct MembershipSpecCertificateManagement : IEquatable<MembershipSpecCertificateManagement>
    {
        private readonly string _value;

        private MembershipSpecCertificateManagement(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Disable workload certificate feature.
        /// </summary>
        public static MembershipSpecCertificateManagement CertificateManagementUnspecified { get; } = new MembershipSpecCertificateManagement("CERTIFICATE_MANAGEMENT_UNSPECIFIED");
        /// <summary>
        /// Disable workload certificate feature.
        /// </summary>
        public static MembershipSpecCertificateManagement Disabled { get; } = new MembershipSpecCertificateManagement("DISABLED");
        /// <summary>
        /// Enable workload certificate feature.
        /// </summary>
        public static MembershipSpecCertificateManagement Enabled { get; } = new MembershipSpecCertificateManagement("ENABLED");

        public static bool operator ==(MembershipSpecCertificateManagement left, MembershipSpecCertificateManagement right) => left.Equals(right);
        public static bool operator !=(MembershipSpecCertificateManagement left, MembershipSpecCertificateManagement right) => !left.Equals(right);

        public static explicit operator string(MembershipSpecCertificateManagement value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MembershipSpecCertificateManagement other && Equals(other);
        public bool Equals(MembershipSpecCertificateManagement other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Deprecated: This field will be ignored and should not be set. Customer's billing structure.
    /// </summary>
    [EnumType]
    public readonly struct MultiClusterIngressFeatureSpecBilling : IEquatable<MultiClusterIngressFeatureSpecBilling>
    {
        private readonly string _value;

        private MultiClusterIngressFeatureSpecBilling(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unknown
        /// </summary>
        public static MultiClusterIngressFeatureSpecBilling BillingUnspecified { get; } = new MultiClusterIngressFeatureSpecBilling("BILLING_UNSPECIFIED");
        /// <summary>
        /// User pays a fee per-endpoint.
        /// </summary>
        public static MultiClusterIngressFeatureSpecBilling PayAsYouGo { get; } = new MultiClusterIngressFeatureSpecBilling("PAY_AS_YOU_GO");
        /// <summary>
        /// User is paying for Anthos as a whole.
        /// </summary>
        public static MultiClusterIngressFeatureSpecBilling AnthosLicense { get; } = new MultiClusterIngressFeatureSpecBilling("ANTHOS_LICENSE");

        public static bool operator ==(MultiClusterIngressFeatureSpecBilling left, MultiClusterIngressFeatureSpecBilling right) => left.Equals(right);
        public static bool operator !=(MultiClusterIngressFeatureSpecBilling left, MultiClusterIngressFeatureSpecBilling right) => !left.Equals(right);

        public static explicit operator string(MultiClusterIngressFeatureSpecBilling value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MultiClusterIngressFeatureSpecBilling other && Equals(other);
        public bool Equals(MultiClusterIngressFeatureSpecBilling other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
