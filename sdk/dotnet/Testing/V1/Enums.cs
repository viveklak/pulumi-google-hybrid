// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Testing.V1
{
    /// <summary>
    /// The option of whether running each test within its own invocation of instrumentation with Android Test Orchestrator or not. ** Orchestrator is only compatible with AndroidJUnitRunner version 1.1 or higher! ** Orchestrator offers the following benefits: - No shared state - Crashes are isolated - Logs are scoped per test See for more information about Android Test Orchestrator. If not set, the test will be run without the orchestrator.
    /// </summary>
    [EnumType]
    public readonly struct AndroidInstrumentationTestOrchestratorOption : IEquatable<AndroidInstrumentationTestOrchestratorOption>
    {
        private readonly string _value;

        private AndroidInstrumentationTestOrchestratorOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value: the server will choose the mode. Currently implies that the test will run without the orchestrator. In the future, all instrumentation tests will be run with the orchestrator. Using the orchestrator is highly encouraged because of all the benefits it offers.
        /// </summary>
        public static AndroidInstrumentationTestOrchestratorOption OrchestratorOptionUnspecified { get; } = new AndroidInstrumentationTestOrchestratorOption("ORCHESTRATOR_OPTION_UNSPECIFIED");
        /// <summary>
        /// Run test using orchestrator. ** Only compatible with AndroidJUnitRunner version 1.1 or higher! ** Recommended.
        /// </summary>
        public static AndroidInstrumentationTestOrchestratorOption UseOrchestrator { get; } = new AndroidInstrumentationTestOrchestratorOption("USE_ORCHESTRATOR");
        /// <summary>
        /// Run test without using orchestrator.
        /// </summary>
        public static AndroidInstrumentationTestOrchestratorOption DoNotUseOrchestrator { get; } = new AndroidInstrumentationTestOrchestratorOption("DO_NOT_USE_ORCHESTRATOR");

        public static bool operator ==(AndroidInstrumentationTestOrchestratorOption left, AndroidInstrumentationTestOrchestratorOption right) => left.Equals(right);
        public static bool operator !=(AndroidInstrumentationTestOrchestratorOption left, AndroidInstrumentationTestOrchestratorOption right) => !left.Equals(right);

        public static explicit operator string(AndroidInstrumentationTestOrchestratorOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AndroidInstrumentationTestOrchestratorOption other && Equals(other);
        public bool Equals(AndroidInstrumentationTestOrchestratorOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The mode in which Robo should run. Most clients should allow the server to populate this field automatically.
    /// </summary>
    [EnumType]
    public readonly struct AndroidRoboTestRoboMode : IEquatable<AndroidRoboTestRoboMode>
    {
        private readonly string _value;

        private AndroidRoboTestRoboMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// This means that the server should choose the mode. Recommended.
        /// </summary>
        public static AndroidRoboTestRoboMode RoboModeUnspecified { get; } = new AndroidRoboTestRoboMode("ROBO_MODE_UNSPECIFIED");
        /// <summary>
        /// Runs Robo in UIAutomator-only mode without app resigning
        /// </summary>
        public static AndroidRoboTestRoboMode RoboVersion1 { get; } = new AndroidRoboTestRoboMode("ROBO_VERSION_1");
        /// <summary>
        /// Runs Robo in standard Espresso with UIAutomator fallback
        /// </summary>
        public static AndroidRoboTestRoboMode RoboVersion2 { get; } = new AndroidRoboTestRoboMode("ROBO_VERSION_2");

        public static bool operator ==(AndroidRoboTestRoboMode left, AndroidRoboTestRoboMode right) => left.Equals(right);
        public static bool operator !=(AndroidRoboTestRoboMode left, AndroidRoboTestRoboMode right) => !left.Equals(right);

        public static explicit operator string(AndroidRoboTestRoboMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AndroidRoboTestRoboMode other && Equals(other);
        public bool Equals(AndroidRoboTestRoboMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Required. The type of action that Robo should perform on the specified element.
    /// </summary>
    [EnumType]
    public readonly struct RoboDirectiveActionType : IEquatable<RoboDirectiveActionType>
    {
        private readonly string _value;

        private RoboDirectiveActionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// DO NOT USE. For proto versioning only.
        /// </summary>
        public static RoboDirectiveActionType ActionTypeUnspecified { get; } = new RoboDirectiveActionType("ACTION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Direct Robo to click on the specified element. No-op if specified element is not clickable.
        /// </summary>
        public static RoboDirectiveActionType SingleClick { get; } = new RoboDirectiveActionType("SINGLE_CLICK");
        /// <summary>
        /// Direct Robo to enter text on the specified element. No-op if specified element is not enabled or does not allow text entry.
        /// </summary>
        public static RoboDirectiveActionType EnterText { get; } = new RoboDirectiveActionType("ENTER_TEXT");
        /// <summary>
        /// Direct Robo to ignore interactions with a specific element.
        /// </summary>
        public static RoboDirectiveActionType Ignore { get; } = new RoboDirectiveActionType("IGNORE");

        public static bool operator ==(RoboDirectiveActionType left, RoboDirectiveActionType right) => left.Equals(right);
        public static bool operator !=(RoboDirectiveActionType left, RoboDirectiveActionType right) => !left.Equals(right);

        public static explicit operator string(RoboDirectiveActionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RoboDirectiveActionType other && Equals(other);
        public bool Equals(RoboDirectiveActionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
