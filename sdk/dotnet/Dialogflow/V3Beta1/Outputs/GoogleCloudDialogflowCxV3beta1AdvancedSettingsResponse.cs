// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// Hierarchical advanced settings for agent/flow/page/fulfillment/parameter. Settings exposed at lower level overrides the settings exposed at higher level. Hierarchy: Agent-&gt;Flow-&gt;Page-&gt;Fulfillment/Parameter.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1AdvancedSettingsResponse
    {
        /// <summary>
        /// Settings for logging. Settings for Dialogflow History, Contact Center messages, StackDriver logs, and speech logging. Exposed at the following levels: - Agent level.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1AdvancedSettingsLoggingSettingsResponse LoggingSettings;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1AdvancedSettingsResponse(Outputs.GoogleCloudDialogflowCxV3beta1AdvancedSettingsLoggingSettingsResponse loggingSettings)
        {
            LoggingSettings = loggingSettings;
        }
    }
}
