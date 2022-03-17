// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.BigtableAdmin.V2
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
    /// Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
    /// </summary>
    [EnumType]
    public readonly struct ClusterDefaultStorageType : IEquatable<ClusterDefaultStorageType>
    {
        private readonly string _value;

        private ClusterDefaultStorageType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The user did not specify a storage type.
        /// </summary>
        public static ClusterDefaultStorageType StorageTypeUnspecified { get; } = new ClusterDefaultStorageType("STORAGE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Flash (SSD) storage should be used.
        /// </summary>
        public static ClusterDefaultStorageType Ssd { get; } = new ClusterDefaultStorageType("SSD");
        /// <summary>
        /// Magnetic drive (HDD) storage should be used.
        /// </summary>
        public static ClusterDefaultStorageType Hdd { get; } = new ClusterDefaultStorageType("HDD");

        public static bool operator ==(ClusterDefaultStorageType left, ClusterDefaultStorageType right) => left.Equals(right);
        public static bool operator !=(ClusterDefaultStorageType left, ClusterDefaultStorageType right) => !left.Equals(right);

        public static explicit operator string(ClusterDefaultStorageType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ClusterDefaultStorageType other && Equals(other);
        public bool Equals(ClusterDefaultStorageType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The type of the instance. Defaults to `PRODUCTION`.
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
        /// The type of the instance is unspecified. If set when creating an instance, a `PRODUCTION` instance will be created. If set when updating an instance, the type will be left unchanged.
        /// </summary>
        public static InstanceType TypeUnspecified { get; } = new InstanceType("TYPE_UNSPECIFIED");
        /// <summary>
        /// An instance meant for production use. `serve_nodes` must be set on the cluster.
        /// </summary>
        public static InstanceType Production { get; } = new InstanceType("PRODUCTION");
        /// <summary>
        /// DEPRECATED: Prefer PRODUCTION for all use cases, as it no longer enforces a higher minimum node count than DEVELOPMENT.
        /// </summary>
        public static InstanceType Development { get; } = new InstanceType("DEVELOPMENT");

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
    /// Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
    /// </summary>
    [EnumType]
    public readonly struct TableGranularity : IEquatable<TableGranularity>
    {
        private readonly string _value;

        private TableGranularity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The user did not specify a granularity. Should not be returned. When specified during table creation, MILLIS will be used.
        /// </summary>
        public static TableGranularity TimestampGranularityUnspecified { get; } = new TableGranularity("TIMESTAMP_GRANULARITY_UNSPECIFIED");
        /// <summary>
        /// The table keeps data versioned at a granularity of 1ms.
        /// </summary>
        public static TableGranularity Millis { get; } = new TableGranularity("MILLIS");

        public static bool operator ==(TableGranularity left, TableGranularity right) => left.Equals(right);
        public static bool operator !=(TableGranularity left, TableGranularity right) => !left.Equals(right);

        public static explicit operator string(TableGranularity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TableGranularity other && Equals(other);
        public bool Equals(TableGranularity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
