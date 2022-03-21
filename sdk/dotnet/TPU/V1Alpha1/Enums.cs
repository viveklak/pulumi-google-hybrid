// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.TPU.V1Alpha1
{
    /// <summary>
    /// The health status of the TPU node.
    /// </summary>
    [EnumType]
    public readonly struct NodeHealth : IEquatable<NodeHealth>
    {
        private readonly string _value;

        private NodeHealth(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Health status is unknown: not initialized or failed to retrieve.
        /// </summary>
        public static NodeHealth HealthUnspecified { get; } = new NodeHealth("HEALTH_UNSPECIFIED");
        /// <summary>
        /// The resource is healthy.
        /// </summary>
        public static NodeHealth Healthy { get; } = new NodeHealth("HEALTHY");
        /// <summary>
        /// The resource is unhealthy.
        /// </summary>
        public static NodeHealth DeprecatedUnhealthy { get; } = new NodeHealth("DEPRECATED_UNHEALTHY");
        /// <summary>
        /// The resource is unresponsive.
        /// </summary>
        public static NodeHealth Timeout { get; } = new NodeHealth("TIMEOUT");
        /// <summary>
        /// The in-guest ML stack is unhealthy.
        /// </summary>
        public static NodeHealth UnhealthyTensorflow { get; } = new NodeHealth("UNHEALTHY_TENSORFLOW");
        /// <summary>
        /// The node is under maintenance/priority boost caused rescheduling and will resume running once rescheduled.
        /// </summary>
        public static NodeHealth UnhealthyMaintenance { get; } = new NodeHealth("UNHEALTHY_MAINTENANCE");

        public static bool operator ==(NodeHealth left, NodeHealth right) => left.Equals(right);
        public static bool operator !=(NodeHealth left, NodeHealth right) => !left.Equals(right);

        public static explicit operator string(NodeHealth value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NodeHealth other && Equals(other);
        public bool Equals(NodeHealth other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
