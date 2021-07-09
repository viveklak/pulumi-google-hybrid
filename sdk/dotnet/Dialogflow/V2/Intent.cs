// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2
{
    /// <summary>
    /// Creates an intent in the specified agent.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v2:Intent")]
    public partial class Intent : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
        /// </summary>
        [Output("defaultResponsePlatforms")]
        public Output<ImmutableArray<string>> DefaultResponsePlatforms { get; private set; } = null!;

        /// <summary>
        /// The name of this intent.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
        /// </summary>
        [Output("endInteraction")]
        public Output<bool> EndInteraction { get; private set; } = null!;

        /// <summary>
        /// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
        /// </summary>
        [Output("events")]
        public Output<ImmutableArray<string>> Events { get; private set; } = null!;

        /// <summary>
        /// Read-only. Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
        /// </summary>
        [Output("followupIntentInfo")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentFollowupIntentInfoResponse>> FollowupIntentInfo { get; private set; } = null!;

        /// <summary>
        /// Optional. The list of context names required for this intent to be triggered. Format: `projects//agent/sessions/-/contexts/`.
        /// </summary>
        [Output("inputContextNames")]
        public Output<ImmutableArray<string>> InputContextNames { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates whether this is a fallback intent.
        /// </summary>
        [Output("isFallback")]
        public Output<bool> IsFallback { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
        /// </summary>
        [Output("liveAgentHandoff")]
        public Output<bool> LiveAgentHandoff { get; private set; } = null!;

        /// <summary>
        /// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
        /// </summary>
        [Output("messages")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageResponse>> Messages { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
        /// </summary>
        [Output("mlDisabled")]
        public Output<bool> MlDisabled { get; private set; } = null!;

        /// <summary>
        /// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Format: `projects//agent/intents/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
        /// </summary>
        [Output("outputContexts")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowV2ContextResponse>> OutputContexts { get; private set; } = null!;

        /// <summary>
        /// Optional. The collection of parameters associated with the intent.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentParameterResponse>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Read-only after creation. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
        /// </summary>
        [Output("parentFollowupIntentName")]
        public Output<string> ParentFollowupIntentName { get; private set; } = null!;

        /// <summary>
        /// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
        /// </summary>
        [Output("resetContexts")]
        public Output<bool> ResetContexts { get; private set; } = null!;

        /// <summary>
        /// Read-only. The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. We populate this field only in the output. Format: `projects//agent/intents/`.
        /// </summary>
        [Output("rootFollowupIntentName")]
        public Output<string> RootFollowupIntentName { get; private set; } = null!;

        /// <summary>
        /// Optional. The collection of examples that the agent is trained on.
        /// </summary>
        [Output("trainingPhrases")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentTrainingPhraseResponse>> TrainingPhrases { get; private set; } = null!;

        /// <summary>
        /// Optional. Indicates whether webhooks are enabled for the intent.
        /// </summary>
        [Output("webhookState")]
        public Output<string> WebhookState { get; private set; } = null!;


        /// <summary>
        /// Create a Intent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Intent(string name, IntentArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:Intent", name, args ?? new IntentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Intent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v2:Intent", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Intent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Intent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Intent(name, id, options);
        }
    }

    public sealed class IntentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("defaultResponsePlatforms")]
        private InputList<Pulumi.GoogleNative.Dialogflow.V2.IntentDefaultResponsePlatformsItem>? _defaultResponsePlatforms;

        /// <summary>
        /// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
        /// </summary>
        public InputList<Pulumi.GoogleNative.Dialogflow.V2.IntentDefaultResponsePlatformsItem> DefaultResponsePlatforms
        {
            get => _defaultResponsePlatforms ?? (_defaultResponsePlatforms = new InputList<Pulumi.GoogleNative.Dialogflow.V2.IntentDefaultResponsePlatformsItem>());
            set => _defaultResponsePlatforms = value;
        }

        /// <summary>
        /// The name of this intent.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
        /// </summary>
        [Input("endInteraction")]
        public Input<bool>? EndInteraction { get; set; }

        [Input("events")]
        private InputList<string>? _events;

        /// <summary>
        /// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
        /// </summary>
        public InputList<string> Events
        {
            get => _events ?? (_events = new InputList<string>());
            set => _events = value;
        }

        [Input("followupIntentInfo")]
        private InputList<Inputs.GoogleCloudDialogflowV2IntentFollowupIntentInfoArgs>? _followupIntentInfo;

        /// <summary>
        /// Read-only. Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2IntentFollowupIntentInfoArgs> FollowupIntentInfo
        {
            get => _followupIntentInfo ?? (_followupIntentInfo = new InputList<Inputs.GoogleCloudDialogflowV2IntentFollowupIntentInfoArgs>());
            set => _followupIntentInfo = value;
        }

        [Input("inputContextNames")]
        private InputList<string>? _inputContextNames;

        /// <summary>
        /// Optional. The list of context names required for this intent to be triggered. Format: `projects//agent/sessions/-/contexts/`.
        /// </summary>
        public InputList<string> InputContextNames
        {
            get => _inputContextNames ?? (_inputContextNames = new InputList<string>());
            set => _inputContextNames = value;
        }

        [Input("intentView")]
        public Input<string>? IntentView { get; set; }

        /// <summary>
        /// Optional. Indicates whether this is a fallback intent.
        /// </summary>
        [Input("isFallback")]
        public Input<bool>? IsFallback { get; set; }

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
        /// </summary>
        [Input("liveAgentHandoff")]
        public Input<bool>? LiveAgentHandoff { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("messages")]
        private InputList<Inputs.GoogleCloudDialogflowV2IntentMessageArgs>? _messages;

        /// <summary>
        /// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2IntentMessageArgs> Messages
        {
            get => _messages ?? (_messages = new InputList<Inputs.GoogleCloudDialogflowV2IntentMessageArgs>());
            set => _messages = value;
        }

        /// <summary>
        /// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
        /// </summary>
        [Input("mlDisabled")]
        public Input<bool>? MlDisabled { get; set; }

        /// <summary>
        /// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Format: `projects//agent/intents/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outputContexts")]
        private InputList<Inputs.GoogleCloudDialogflowV2ContextArgs>? _outputContexts;

        /// <summary>
        /// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2ContextArgs> OutputContexts
        {
            get => _outputContexts ?? (_outputContexts = new InputList<Inputs.GoogleCloudDialogflowV2ContextArgs>());
            set => _outputContexts = value;
        }

        [Input("parameters")]
        private InputList<Inputs.GoogleCloudDialogflowV2IntentParameterArgs>? _parameters;

        /// <summary>
        /// Optional. The collection of parameters associated with the intent.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2IntentParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.GoogleCloudDialogflowV2IntentParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Read-only after creation. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
        /// </summary>
        [Input("parentFollowupIntentName")]
        public Input<string>? ParentFollowupIntentName { get; set; }

        /// <summary>
        /// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
        /// </summary>
        [Input("resetContexts")]
        public Input<bool>? ResetContexts { get; set; }

        /// <summary>
        /// Read-only. The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. We populate this field only in the output. Format: `projects//agent/intents/`.
        /// </summary>
        [Input("rootFollowupIntentName")]
        public Input<string>? RootFollowupIntentName { get; set; }

        [Input("trainingPhrases")]
        private InputList<Inputs.GoogleCloudDialogflowV2IntentTrainingPhraseArgs>? _trainingPhrases;

        /// <summary>
        /// Optional. The collection of examples that the agent is trained on.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2IntentTrainingPhraseArgs> TrainingPhrases
        {
            get => _trainingPhrases ?? (_trainingPhrases = new InputList<Inputs.GoogleCloudDialogflowV2IntentTrainingPhraseArgs>());
            set => _trainingPhrases = value;
        }

        /// <summary>
        /// Optional. Indicates whether webhooks are enabled for the intent.
        /// </summary>
        [Input("webhookState")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2.IntentWebhookState>? WebhookState { get; set; }

        public IntentArgs()
        {
        }
    }
}
