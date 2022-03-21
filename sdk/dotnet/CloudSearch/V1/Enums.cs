// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.CloudSearch.V1
{
    /// <summary>
    /// The logic operator of the sub filter.
    /// </summary>
    [EnumType]
    public readonly struct CompositeFilterLogicOperator : IEquatable<CompositeFilterLogicOperator>
    {
        private readonly string _value;

        private CompositeFilterLogicOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Logical operators, which can only be applied to sub filters.
        /// </summary>
        public static CompositeFilterLogicOperator And { get; } = new CompositeFilterLogicOperator("AND");
        public static CompositeFilterLogicOperator Or { get; } = new CompositeFilterLogicOperator("OR");
        /// <summary>
        /// NOT can only be applied on a single sub filter.
        /// </summary>
        public static CompositeFilterLogicOperator Not { get; } = new CompositeFilterLogicOperator("NOT");

        public static bool operator ==(CompositeFilterLogicOperator left, CompositeFilterLogicOperator right) => left.Equals(right);
        public static bool operator !=(CompositeFilterLogicOperator left, CompositeFilterLogicOperator right) => !left.Equals(right);

        public static explicit operator string(CompositeFilterLogicOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CompositeFilterLogicOperator other && Equals(other);
        public bool Equals(CompositeFilterLogicOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Ascending is the default sort order
    /// </summary>
    [EnumType]
    public readonly struct SortOptionsSortOrder : IEquatable<SortOptionsSortOrder>
    {
        private readonly string _value;

        private SortOptionsSortOrder(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SortOptionsSortOrder Ascending { get; } = new SortOptionsSortOrder("ASCENDING");
        public static SortOptionsSortOrder Descending { get; } = new SortOptionsSortOrder("DESCENDING");

        public static bool operator ==(SortOptionsSortOrder left, SortOptionsSortOrder right) => left.Equals(right);
        public static bool operator !=(SortOptionsSortOrder left, SortOptionsSortOrder right) => !left.Equals(right);

        public static explicit operator string(SortOptionsSortOrder value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SortOptionsSortOrder other && Equals(other);
        public bool Equals(SortOptionsSortOrder other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Predefined content source for Google Apps.
    /// </summary>
    [EnumType]
    public readonly struct SourcePredefinedSource : IEquatable<SourcePredefinedSource>
    {
        private readonly string _value;

        private SourcePredefinedSource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SourcePredefinedSource None { get; } = new SourcePredefinedSource("NONE");
        /// <summary>
        /// Suggests queries issued by the user in the past. Only valid when used with the suggest API. Ignored when used in the query API.
        /// </summary>
        public static SourcePredefinedSource QueryHistory { get; } = new SourcePredefinedSource("QUERY_HISTORY");
        /// <summary>
        /// Suggests people in the organization. Only valid when used with the suggest API. Results in an error when used in the query API.
        /// </summary>
        public static SourcePredefinedSource Person { get; } = new SourcePredefinedSource("PERSON");
        public static SourcePredefinedSource GoogleDrive { get; } = new SourcePredefinedSource("GOOGLE_DRIVE");
        public static SourcePredefinedSource GoogleGmail { get; } = new SourcePredefinedSource("GOOGLE_GMAIL");
        public static SourcePredefinedSource GoogleSites { get; } = new SourcePredefinedSource("GOOGLE_SITES");
        public static SourcePredefinedSource GoogleGroups { get; } = new SourcePredefinedSource("GOOGLE_GROUPS");
        public static SourcePredefinedSource GoogleCalendar { get; } = new SourcePredefinedSource("GOOGLE_CALENDAR");
        public static SourcePredefinedSource GoogleKeep { get; } = new SourcePredefinedSource("GOOGLE_KEEP");

        public static bool operator ==(SourcePredefinedSource left, SourcePredefinedSource right) => left.Equals(right);
        public static bool operator !=(SourcePredefinedSource left, SourcePredefinedSource right) => !left.Equals(right);

        public static explicit operator string(SourcePredefinedSource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SourcePredefinedSource other && Equals(other);
        public bool Equals(SourcePredefinedSource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Importance of the source.
    /// </summary>
    [EnumType]
    public readonly struct SourceScoringConfigSourceImportance : IEquatable<SourceScoringConfigSourceImportance>
    {
        private readonly string _value;

        private SourceScoringConfigSourceImportance(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SourceScoringConfigSourceImportance Default { get; } = new SourceScoringConfigSourceImportance("DEFAULT");
        public static SourceScoringConfigSourceImportance Low { get; } = new SourceScoringConfigSourceImportance("LOW");
        public static SourceScoringConfigSourceImportance High { get; } = new SourceScoringConfigSourceImportance("HIGH");

        public static bool operator ==(SourceScoringConfigSourceImportance left, SourceScoringConfigSourceImportance right) => left.Equals(right);
        public static bool operator !=(SourceScoringConfigSourceImportance left, SourceScoringConfigSourceImportance right) => !left.Equals(right);

        public static explicit operator string(SourceScoringConfigSourceImportance value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SourceScoringConfigSourceImportance other && Equals(other);
        public bool Equals(SourceScoringConfigSourceImportance other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
