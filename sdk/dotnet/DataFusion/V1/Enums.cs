// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.DataFusion.V1
{
    /// <summary>
    /// The type of an accelator for a CDF instance.
    /// </summary>
    [EnumType]
    public readonly struct AcceleratorAcceleratorType : IEquatable<AcceleratorAcceleratorType>
    {
        private readonly string _value;

        private AcceleratorAcceleratorType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value, if unspecified.
        /// </summary>
        public static AcceleratorAcceleratorType AcceleratorTypeUnspecified { get; } = new AcceleratorAcceleratorType("ACCELERATOR_TYPE_UNSPECIFIED");
        /// <summary>
        /// Change Data Capture accelerator for CDF.
        /// </summary>
        public static AcceleratorAcceleratorType Cdc { get; } = new AcceleratorAcceleratorType("CDC");
        /// <summary>
        /// Cloud Healthcare accelerator for CDF. This accelerator is to enable Cloud Healthcare specific CDF plugins developed by Healthcare team.
        /// </summary>
        public static AcceleratorAcceleratorType Healthcare { get; } = new AcceleratorAcceleratorType("HEALTHCARE");
        /// <summary>
        /// Contact Center AI Insights This accelerator is used to enable import and export pipelines custom built to streamline CCAI Insights processing.
        /// </summary>
        public static AcceleratorAcceleratorType CcaiInsights { get; } = new AcceleratorAcceleratorType("CCAI_INSIGHTS");

        public static bool operator ==(AcceleratorAcceleratorType left, AcceleratorAcceleratorType right) => left.Equals(right);
        public static bool operator !=(AcceleratorAcceleratorType left, AcceleratorAcceleratorType right) => !left.Equals(right);

        public static explicit operator string(AcceleratorAcceleratorType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AcceleratorAcceleratorType other && Equals(other);
        public bool Equals(AcceleratorAcceleratorType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The state of the accelerator
    /// </summary>
    [EnumType]
    public readonly struct AcceleratorState : IEquatable<AcceleratorState>
    {
        private readonly string _value;

        private AcceleratorState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value, do not use
        /// </summary>
        public static AcceleratorState StateUnspecified { get; } = new AcceleratorState("STATE_UNSPECIFIED");
        /// <summary>
        /// Indicates that the accelerator is enabled and available to use
        /// </summary>
        public static AcceleratorState Enabled { get; } = new AcceleratorState("ENABLED");
        /// <summary>
        /// Indicates that the accelerator is disabled and not available to use
        /// </summary>
        public static AcceleratorState Disabled { get; } = new AcceleratorState("DISABLED");
        /// <summary>
        /// Indicates that accelerator state is currently unknown. Requests for enable, disable could be retried while in this state
        /// </summary>
        public static AcceleratorState Unknown { get; } = new AcceleratorState("UNKNOWN");

        public static bool operator ==(AcceleratorState left, AcceleratorState right) => left.Equals(right);
        public static bool operator !=(AcceleratorState left, AcceleratorState right) => !left.Equals(right);

        public static explicit operator string(AcceleratorState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AcceleratorState other && Equals(other);
        public bool Equals(AcceleratorState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

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
    /// Required. Instance type.
    /// </summary>
    [EnumType]
    public readonly struct InstanceType : IEquatable<InstanceType>
    {
        private readonly string _value;

        private InstanceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// No type specified. The instance creation will fail.
        /// </summary>
        public static InstanceType TypeUnspecified { get; } = new InstanceType("TYPE_UNSPECIFIED");
        /// <summary>
        /// Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines using point and click UI. However, there are certain limitations, such as fewer number of concurrent pipelines, no support for streaming pipelines, etc.
        /// </summary>
        public static InstanceType Basic { get; } = new InstanceType("BASIC");
        /// <summary>
        /// Enterprise Data Fusion instance. In Enterprise type, the user will have all features available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
        /// </summary>
        public static InstanceType Enterprise { get; } = new InstanceType("ENTERPRISE");
        /// <summary>
        /// Developer Data Fusion instance. In Developer type, the user will have all features available but with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration pipelines at low cost.
        /// </summary>
        public static InstanceType Developer { get; } = new InstanceType("DEVELOPER");

        public static bool operator ==(InstanceType left, InstanceType right) => left.Equals(right);
        public static bool operator !=(InstanceType left, InstanceType right) => !left.Equals(right);

        public static explicit operator string(InstanceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceType other && Equals(other);
        public bool Equals(InstanceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type represents the release availability of the version
    /// </summary>
    [EnumType]
    public readonly struct VersionType : IEquatable<VersionType>
    {
        private readonly string _value;

        private VersionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Version does not have availability yet
        /// </summary>
        public static VersionType TypeUnspecified { get; } = new VersionType("TYPE_UNSPECIFIED");
        /// <summary>
        /// Version is under development and not considered stable
        /// </summary>
        public static VersionType TypePreview { get; } = new VersionType("TYPE_PREVIEW");
        /// <summary>
        /// Version is available for public use
        /// </summary>
        public static VersionType TypeGeneralAvailability { get; } = new VersionType("TYPE_GENERAL_AVAILABILITY");

        public static bool operator ==(VersionType left, VersionType right) => left.Equals(right);
        public static bool operator !=(VersionType left, VersionType right) => !left.Equals(right);

        public static explicit operator string(VersionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VersionType other && Equals(other);
        public bool Equals(VersionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
