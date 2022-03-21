// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Memcache.V1
{
    /// <summary>
    /// The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
    /// </summary>
    [EnumType]
    public readonly struct InstanceMemcacheVersion : IEquatable<InstanceMemcacheVersion>
    {
        private readonly string _value;

        private InstanceMemcacheVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InstanceMemcacheVersion MemcacheVersionUnspecified { get; } = new InstanceMemcacheVersion("MEMCACHE_VERSION_UNSPECIFIED");
        /// <summary>
        /// Memcached 1.5 version.
        /// </summary>
        public static InstanceMemcacheVersion Memcache15 { get; } = new InstanceMemcacheVersion("MEMCACHE_1_5");

        public static bool operator ==(InstanceMemcacheVersion left, InstanceMemcacheVersion right) => left.Equals(right);
        public static bool operator !=(InstanceMemcacheVersion left, InstanceMemcacheVersion right) => !left.Equals(right);

        public static explicit operator string(InstanceMemcacheVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceMemcacheVersion other && Equals(other);
        public bool Equals(InstanceMemcacheVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A code that correspond to one type of user-facing message.
    /// </summary>
    [EnumType]
    public readonly struct InstanceMessageCode : IEquatable<InstanceMessageCode>
    {
        private readonly string _value;

        private InstanceMessageCode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Message Code not set.
        /// </summary>
        public static InstanceMessageCode CodeUnspecified { get; } = new InstanceMessageCode("CODE_UNSPECIFIED");
        /// <summary>
        /// Memcached nodes are distributed unevenly.
        /// </summary>
        public static InstanceMessageCode ZoneDistributionUnbalanced { get; } = new InstanceMessageCode("ZONE_DISTRIBUTION_UNBALANCED");

        public static bool operator ==(InstanceMessageCode left, InstanceMessageCode right) => left.Equals(right);
        public static bool operator !=(InstanceMessageCode left, InstanceMessageCode right) => !left.Equals(right);

        public static explicit operator string(InstanceMessageCode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceMessageCode other && Equals(other);
        public bool Equals(InstanceMessageCode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
