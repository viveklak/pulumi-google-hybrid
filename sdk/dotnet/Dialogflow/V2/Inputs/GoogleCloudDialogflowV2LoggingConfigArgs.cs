// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// Defines logging behavior for conversation lifecycle events.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2LoggingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to log conversation events like CONVERSATION_STARTED to Stackdriver in the conversation project as JSON format ConversationEvent protos.
        /// </summary>
        [Input("enableStackdriverLogging")]
        public Input<bool>? EnableStackdriverLogging { get; set; }

        public GoogleCloudDialogflowV2LoggingConfigArgs()
        {
        }
    }
}
