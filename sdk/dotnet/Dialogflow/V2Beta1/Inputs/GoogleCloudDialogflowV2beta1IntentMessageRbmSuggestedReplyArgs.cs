// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Rich Business Messaging (RBM) suggested reply that the user can click instead of typing in their own response.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestedReplyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Opaque payload that the Dialogflow receives in a user event when the user taps the suggested reply. This data will be also forwarded to webhook to allow performing custom business logic.
        /// </summary>
        [Input("postbackData")]
        public Input<string>? PostbackData { get; set; }

        /// <summary>
        /// Suggested reply text.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestedReplyArgs()
        {
        }
    }
}
