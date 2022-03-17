// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1
{
    /// <summary>
    /// Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
    /// </summary>
    [EnumType]
    public readonly struct ConversationMedium : IEquatable<ConversationMedium>
    {
        private readonly string _value;

        private ConversationMedium(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value, if unspecified will default to PHONE_CALL.
        /// </summary>
        public static ConversationMedium MediumUnspecified { get; } = new ConversationMedium("MEDIUM_UNSPECIFIED");
        /// <summary>
        /// The format for conversations that took place over the phone.
        /// </summary>
        public static ConversationMedium PhoneCall { get; } = new ConversationMedium("PHONE_CALL");
        /// <summary>
        /// The format for conversations that took place over chat.
        /// </summary>
        public static ConversationMedium Chat { get; } = new ConversationMedium("CHAT");

        public static bool operator ==(ConversationMedium left, ConversationMedium right) => left.Equals(right);
        public static bool operator !=(ConversationMedium left, ConversationMedium right) => !left.Equals(right);

        public static explicit operator string(ConversationMedium value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConversationMedium other && Equals(other);
        public bool Equals(ConversationMedium other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Medium of conversations used in training data. This field is being deprecated. To specify the medium to be used in training a new issue model, set the `medium` field on `filter`.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium : IEquatable<GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium>
    {
        private readonly string _value;

        private GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value, if unspecified will default to PHONE_CALL.
        /// </summary>
        public static GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium MediumUnspecified { get; } = new GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium("MEDIUM_UNSPECIFIED");
        /// <summary>
        /// The format for conversations that took place over the phone.
        /// </summary>
        public static GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium PhoneCall { get; } = new GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium("PHONE_CALL");
        /// <summary>
        /// The format for conversations that took place over chat.
        /// </summary>
        public static GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium Chat { get; } = new GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium("CHAT");

        public static bool operator ==(GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium left, GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium right) => left.Equals(right);
        public static bool operator !=(GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium left, GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium other && Equals(other);
        public bool Equals(GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The type of this phrase match rule group.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType : IEquatable<GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType>
    {
        private readonly string _value;

        private GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified.
        /// </summary>
        public static GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType PhraseMatchRuleGroupTypeUnspecified { get; } = new GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType("PHRASE_MATCH_RULE_GROUP_TYPE_UNSPECIFIED");
        /// <summary>
        /// Must meet all phrase match rules or there is no match.
        /// </summary>
        public static GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType AllOf { get; } = new GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType("ALL_OF");
        /// <summary>
        /// If any of the phrase match rules are met, there is a match.
        /// </summary>
        public static GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType AnyOf { get; } = new GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType("ANY_OF");

        public static bool operator ==(GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType left, GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType right) => left.Equals(right);
        public static bool operator !=(GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType left, GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType other && Equals(other);
        public bool Equals(GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
    /// </summary>
    [EnumType]
    public readonly struct PhraseMatcherRoleMatch : IEquatable<PhraseMatcherRoleMatch>
    {
        private readonly string _value;

        private PhraseMatcherRoleMatch(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Participant's role is not set.
        /// </summary>
        public static PhraseMatcherRoleMatch RoleUnspecified { get; } = new PhraseMatcherRoleMatch("ROLE_UNSPECIFIED");
        /// <summary>
        /// Participant is a human agent.
        /// </summary>
        public static PhraseMatcherRoleMatch HumanAgent { get; } = new PhraseMatcherRoleMatch("HUMAN_AGENT");
        /// <summary>
        /// Participant is an automated agent.
        /// </summary>
        public static PhraseMatcherRoleMatch AutomatedAgent { get; } = new PhraseMatcherRoleMatch("AUTOMATED_AGENT");
        /// <summary>
        /// Participant is an end user who conversed with the contact center.
        /// </summary>
        public static PhraseMatcherRoleMatch EndUser { get; } = new PhraseMatcherRoleMatch("END_USER");
        /// <summary>
        /// Participant is either a human or automated agent.
        /// </summary>
        public static PhraseMatcherRoleMatch AnyAgent { get; } = new PhraseMatcherRoleMatch("ANY_AGENT");

        public static bool operator ==(PhraseMatcherRoleMatch left, PhraseMatcherRoleMatch right) => left.Equals(right);
        public static bool operator !=(PhraseMatcherRoleMatch left, PhraseMatcherRoleMatch right) => !left.Equals(right);

        public static explicit operator string(PhraseMatcherRoleMatch value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PhraseMatcherRoleMatch other && Equals(other);
        public bool Equals(PhraseMatcherRoleMatch other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The type of this phrase matcher.
    /// </summary>
    [EnumType]
    public readonly struct PhraseMatcherType : IEquatable<PhraseMatcherType>
    {
        private readonly string _value;

        private PhraseMatcherType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified.
        /// </summary>
        public static PhraseMatcherType PhraseMatcherTypeUnspecified { get; } = new PhraseMatcherType("PHRASE_MATCHER_TYPE_UNSPECIFIED");
        /// <summary>
        /// Must meet all phrase match rule groups or there is no match.
        /// </summary>
        public static PhraseMatcherType AllOf { get; } = new PhraseMatcherType("ALL_OF");
        /// <summary>
        /// If any of the phrase match rule groups are met, there is a match.
        /// </summary>
        public static PhraseMatcherType AnyOf { get; } = new PhraseMatcherType("ANY_OF");

        public static bool operator ==(PhraseMatcherType left, PhraseMatcherType right) => left.Equals(right);
        public static bool operator !=(PhraseMatcherType left, PhraseMatcherType right) => !left.Equals(right);

        public static explicit operator string(PhraseMatcherType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PhraseMatcherType other && Equals(other);
        public bool Equals(PhraseMatcherType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
