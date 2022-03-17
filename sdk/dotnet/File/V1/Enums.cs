// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.File.V1
{
    /// <summary>
    /// The service tier of the instance.
    /// </summary>
    [EnumType]
    public readonly struct InstanceTier : IEquatable<InstanceTier>
    {
        private readonly string _value;

        private InstanceTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not set.
        /// </summary>
        public static InstanceTier TierUnspecified { get; } = new InstanceTier("TIER_UNSPECIFIED");
        /// <summary>
        /// STANDARD tier.
        /// </summary>
        public static InstanceTier Standard { get; } = new InstanceTier("STANDARD");
        /// <summary>
        /// PREMIUM tier.
        /// </summary>
        public static InstanceTier Premium { get; } = new InstanceTier("PREMIUM");
        /// <summary>
        /// BASIC instances offer a maximum capacity of 63.9 TB. BASIC_HDD is an alias for STANDARD Tier, offering economical performance backed by HDD.
        /// </summary>
        public static InstanceTier BasicHdd { get; } = new InstanceTier("BASIC_HDD");
        /// <summary>
        /// BASIC instances offer a maximum capacity of 63.9 TB. BASIC_SSD is an alias for PREMIUM Tier, and offers improved performance backed by SSD.
        /// </summary>
        public static InstanceTier BasicSsd { get; } = new InstanceTier("BASIC_SSD");
        /// <summary>
        /// HIGH_SCALE instances offer expanded capacity and performance scaling capabilities.
        /// </summary>
        public static InstanceTier HighScaleSsd { get; } = new InstanceTier("HIGH_SCALE_SSD");
        /// <summary>
        /// ENTERPRISE instances offer the features and availability needed for mission-critical workloads.
        /// </summary>
        public static InstanceTier Enterprise { get; } = new InstanceTier("ENTERPRISE");

        public static bool operator ==(InstanceTier left, InstanceTier right) => left.Equals(right);
        public static bool operator !=(InstanceTier left, InstanceTier right) => !left.Equals(right);

        public static explicit operator string(InstanceTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InstanceTier other && Equals(other);
        public bool Equals(InstanceTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
    /// </summary>
    [EnumType]
    public readonly struct NetworkConfigConnectMode : IEquatable<NetworkConfigConnectMode>
    {
        private readonly string _value;

        private NetworkConfigConnectMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Not set.
        /// </summary>
        public static NetworkConfigConnectMode ConnectModeUnspecified { get; } = new NetworkConfigConnectMode("CONNECT_MODE_UNSPECIFIED");
        /// <summary>
        /// Connect via direct peering to the Filestore service.
        /// </summary>
        public static NetworkConfigConnectMode DirectPeering { get; } = new NetworkConfigConnectMode("DIRECT_PEERING");
        /// <summary>
        /// Connect to your Filestore instance using Private Service Access. Private services access provides an IP address range for multiple Google Cloud services, including Filestore.
        /// </summary>
        public static NetworkConfigConnectMode PrivateServiceAccess { get; } = new NetworkConfigConnectMode("PRIVATE_SERVICE_ACCESS");

        public static bool operator ==(NetworkConfigConnectMode left, NetworkConfigConnectMode right) => left.Equals(right);
        public static bool operator !=(NetworkConfigConnectMode left, NetworkConfigConnectMode right) => !left.Equals(right);

        public static explicit operator string(NetworkConfigConnectMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkConfigConnectMode other && Equals(other);
        public bool Equals(NetworkConfigConnectMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NetworkConfigModesItem : IEquatable<NetworkConfigModesItem>
    {
        private readonly string _value;

        private NetworkConfigModesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Internet protocol not set.
        /// </summary>
        public static NetworkConfigModesItem AddressModeUnspecified { get; } = new NetworkConfigModesItem("ADDRESS_MODE_UNSPECIFIED");
        /// <summary>
        /// Use the IPv4 internet protocol.
        /// </summary>
        public static NetworkConfigModesItem ModeIpv4 { get; } = new NetworkConfigModesItem("MODE_IPV4");

        public static bool operator ==(NetworkConfigModesItem left, NetworkConfigModesItem right) => left.Equals(right);
        public static bool operator !=(NetworkConfigModesItem left, NetworkConfigModesItem right) => !left.Equals(right);

        public static explicit operator string(NetworkConfigModesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkConfigModesItem other && Equals(other);
        public bool Equals(NetworkConfigModesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
    /// </summary>
    [EnumType]
    public readonly struct NfsExportOptionsAccessMode : IEquatable<NfsExportOptionsAccessMode>
    {
        private readonly string _value;

        private NfsExportOptionsAccessMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// AccessMode not set.
        /// </summary>
        public static NfsExportOptionsAccessMode AccessModeUnspecified { get; } = new NfsExportOptionsAccessMode("ACCESS_MODE_UNSPECIFIED");
        /// <summary>
        /// The client can only read the file share.
        /// </summary>
        public static NfsExportOptionsAccessMode ReadOnly { get; } = new NfsExportOptionsAccessMode("READ_ONLY");
        /// <summary>
        /// The client can read and write the file share (default).
        /// </summary>
        public static NfsExportOptionsAccessMode ReadWrite { get; } = new NfsExportOptionsAccessMode("READ_WRITE");

        public static bool operator ==(NfsExportOptionsAccessMode left, NfsExportOptionsAccessMode right) => left.Equals(right);
        public static bool operator !=(NfsExportOptionsAccessMode left, NfsExportOptionsAccessMode right) => !left.Equals(right);

        public static explicit operator string(NfsExportOptionsAccessMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NfsExportOptionsAccessMode other && Equals(other);
        public bool Equals(NfsExportOptionsAccessMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
    /// </summary>
    [EnumType]
    public readonly struct NfsExportOptionsSquashMode : IEquatable<NfsExportOptionsSquashMode>
    {
        private readonly string _value;

        private NfsExportOptionsSquashMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// SquashMode not set.
        /// </summary>
        public static NfsExportOptionsSquashMode SquashModeUnspecified { get; } = new NfsExportOptionsSquashMode("SQUASH_MODE_UNSPECIFIED");
        /// <summary>
        /// The Root user has root access to the file share (default).
        /// </summary>
        public static NfsExportOptionsSquashMode NoRootSquash { get; } = new NfsExportOptionsSquashMode("NO_ROOT_SQUASH");
        /// <summary>
        /// The Root user has squashed access to the anonymous uid/gid.
        /// </summary>
        public static NfsExportOptionsSquashMode RootSquash { get; } = new NfsExportOptionsSquashMode("ROOT_SQUASH");

        public static bool operator ==(NfsExportOptionsSquashMode left, NfsExportOptionsSquashMode right) => left.Equals(right);
        public static bool operator !=(NfsExportOptionsSquashMode left, NfsExportOptionsSquashMode right) => !left.Equals(right);

        public static explicit operator string(NfsExportOptionsSquashMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NfsExportOptionsSquashMode other && Equals(other);
        public bool Equals(NfsExportOptionsSquashMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
