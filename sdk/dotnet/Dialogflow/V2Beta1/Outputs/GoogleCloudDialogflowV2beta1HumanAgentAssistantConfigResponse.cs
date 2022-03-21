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
    /// Defines the Human Agent Assistant to connect to a conversation.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponse
    {
        /// <summary>
        /// Configuration for agent assistance of end user participant. Currently, this feature is not general available, please contact Google to get access.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionConfigResponse EndUserSuggestionConfig;
        /// <summary>
        /// Configuration for agent assistance of human agent participant.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionConfigResponse HumanAgentSuggestionConfig;
        /// <summary>
        /// Configuration for message analysis.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigMessageAnalysisConfigResponse MessageAnalysisConfig;
        /// <summary>
        /// Pub/Sub topic on which to publish new agent assistant events.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1NotificationConfigResponse NotificationConfig;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigResponse(
            Outputs.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionConfigResponse endUserSuggestionConfig,

            Outputs.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionConfigResponse humanAgentSuggestionConfig,

            Outputs.GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigMessageAnalysisConfigResponse messageAnalysisConfig,

            Outputs.GoogleCloudDialogflowV2beta1NotificationConfigResponse notificationConfig)
        {
            EndUserSuggestionConfig = endUserSuggestionConfig;
            HumanAgentSuggestionConfig = humanAgentSuggestionConfig;
            MessageAnalysisConfig = messageAnalysisConfig;
            NotificationConfig = notificationConfig;
        }
    }
}
