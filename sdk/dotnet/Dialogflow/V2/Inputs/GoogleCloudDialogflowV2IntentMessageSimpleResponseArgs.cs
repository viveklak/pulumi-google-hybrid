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
    /// The simple response message containing speech or text.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2IntentMessageSimpleResponseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The text to display.
        /// </summary>
        [Input("displayText")]
        public Input<string>? DisplayText { get; set; }

        /// <summary>
        /// One of text_to_speech or ssml must be provided. Structured spoken response to the user in the SSML format. Mutually exclusive with text_to_speech.
        /// </summary>
        [Input("ssml")]
        public Input<string>? Ssml { get; set; }

        /// <summary>
        /// One of text_to_speech or ssml must be provided. The plain text of the speech output. Mutually exclusive with ssml.
        /// </summary>
        [Input("textToSpeech")]
        public Input<string>? TextToSpeech { get; set; }

        public GoogleCloudDialogflowV2IntentMessageSimpleResponseArgs()
        {
        }
    }
}
