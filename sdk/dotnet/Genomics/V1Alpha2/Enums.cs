// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Genomics.V1Alpha2
{
    /// <summary>
    /// Required. The type of the disk to create.
    /// </summary>
    [EnumType]
    public readonly struct DiskType : IEquatable<DiskType>
    {
        private readonly string _value;

        private DiskType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default disk type. Use one of the other options below.
        /// </summary>
        public static DiskType TypeUnspecified { get; } = new DiskType("TYPE_UNSPECIFIED");
        /// <summary>
        /// Specifies a Google Compute Engine persistent hard disk. See https://cloud.google.com/compute/docs/disks/#pdspecs for details.
        /// </summary>
        public static DiskType PersistentHdd { get; } = new DiskType("PERSISTENT_HDD");
        /// <summary>
        /// Specifies a Google Compute Engine persistent solid-state disk. See https://cloud.google.com/compute/docs/disks/#pdspecs for details.
        /// </summary>
        public static DiskType PersistentSsd { get; } = new DiskType("PERSISTENT_SSD");
        /// <summary>
        /// Specifies a Google Compute Engine local SSD. See https://cloud.google.com/compute/docs/disks/local-ssd for details.
        /// </summary>
        public static DiskType LocalSsd { get; } = new DiskType("LOCAL_SSD");

        public static bool operator ==(DiskType left, DiskType right) => left.Equals(right);
        public static bool operator !=(DiskType left, DiskType right) => !left.Equals(right);

        public static explicit operator string(DiskType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiskType other && Equals(other);
        public bool Equals(DiskType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
