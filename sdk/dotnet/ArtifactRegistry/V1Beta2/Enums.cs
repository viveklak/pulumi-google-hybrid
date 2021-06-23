// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.ArtifactRegistry.V1Beta2
{
    /// <summary>
    /// The format of packages that are stored in the repository.
    /// </summary>
    [EnumType]
    public readonly struct RepositoryFormat : IEquatable<RepositoryFormat>
    {
        private readonly string _value;

        private RepositoryFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified package format.
        /// </summary>
        public static RepositoryFormat FormatUnspecified { get; } = new RepositoryFormat("FORMAT_UNSPECIFIED");
        /// <summary>
        /// Docker package format.
        /// </summary>
        public static RepositoryFormat Docker { get; } = new RepositoryFormat("DOCKER");
        /// <summary>
        /// Maven package format.
        /// </summary>
        public static RepositoryFormat Maven { get; } = new RepositoryFormat("MAVEN");
        /// <summary>
        /// NPM package format.
        /// </summary>
        public static RepositoryFormat Npm { get; } = new RepositoryFormat("NPM");
        /// <summary>
        /// PyPI package format.
        /// </summary>
        public static RepositoryFormat Pypi { get; } = new RepositoryFormat("PYPI");
        /// <summary>
        /// Python package format.
        /// </summary>
        public static RepositoryFormat Python { get; } = new RepositoryFormat("PYTHON");

        public static bool operator ==(RepositoryFormat left, RepositoryFormat right) => left.Equals(right);
        public static bool operator !=(RepositoryFormat left, RepositoryFormat right) => !left.Equals(right);

        public static explicit operator string(RepositoryFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RepositoryFormat other && Equals(other);
        public bool Equals(RepositoryFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
