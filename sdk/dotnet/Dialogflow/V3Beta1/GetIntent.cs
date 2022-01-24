// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1
{
    public static class GetIntent
    {
        /// <summary>
        /// Retrieves the specified intent.
        /// </summary>
        public static Task<GetIntentResult> InvokeAsync(GetIntentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntentResult>("google-native:dialogflow/v3beta1:getIntent", args ?? new GetIntentArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified intent.
        /// </summary>
        public static Output<GetIntentResult> Invoke(GetIntentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIntentResult>("google-native:dialogflow/v3beta1:getIntent", args ?? new GetIntentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIntentArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("intentId", required: true)]
        public string IntentId { get; set; } = null!;

        [Input("languageCode")]
        public string? LanguageCode { get; set; }

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetIntentArgs()
        {
        }
    }

    public sealed class GetIntentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("intentId", required: true)]
        public Input<string> IntentId { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetIntentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntentResult
    {
        /// <summary>
        /// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The human-readable name of the intent, unique within the agent.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        /// </summary>
        public readonly bool IsFallback;
        /// <summary>
        /// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys-contextual" means the intent is a contextual intent.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The collection of parameters associated with the intent.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1IntentParameterResponse> Parameters;
        /// <summary>
        /// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The collection of training phrases the agent is trained on to identify the intent.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponse> TrainingPhrases;

        [OutputConstructor]
        private GetIntentResult(
            string description,

            string displayName,

            bool isFallback,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1IntentParameterResponse> parameters,

            int priority,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1IntentTrainingPhraseResponse> trainingPhrases)
        {
            Description = description;
            DisplayName = displayName;
            IsFallback = isFallback;
            Labels = labels;
            Name = name;
            Parameters = parameters;
            Priority = priority;
            TrainingPhrases = trainingPhrases;
        }
    }
}
