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
    /// The suggestion chip message that allows the user to jump out to the app or website associated with this agent.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the app or site this chip is linking to.
        /// </summary>
        [Input("destinationName", required: true)]
        public Input<string> DestinationName { get; set; } = null!;

        /// <summary>
        /// The URI of the app or site to open when the user taps the suggestion chip.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1IntentMessageLinkOutSuggestionArgs()
        {
        }
    }
}
