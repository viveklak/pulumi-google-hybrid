// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1IntentResponse
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
        private GoogleCloudDialogflowCxV3beta1IntentResponse(
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
