// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// Dialogflow source setting. Supported feature: DIALOGFLOW_ASSIST.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionQueryConfigDialogflowQuerySourceResponse
    {
        /// <summary>
        /// The name of a dialogflow virtual agent used for end user side intent detection and suggestion. Format: `projects//locations//agent`. When multiple agents are allowed in the same Dialogflow project.
        /// </summary>
        public readonly string Agent;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionQueryConfigDialogflowQuerySourceResponse(string agent)
        {
            Agent = agent;
        }
    }
}
