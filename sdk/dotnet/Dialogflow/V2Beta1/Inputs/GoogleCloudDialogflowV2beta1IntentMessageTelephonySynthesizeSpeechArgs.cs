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
    /// Synthesizes speech and plays back the synthesized audio to the caller in Telephony Gateway. Telephony Gateway takes the synthesizer settings from `DetectIntentResponse.output_audio_config` which can either be set at request-level or can come from the agent-level synthesizer config.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeechArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SSML to be synthesized. For more information, see [SSML](https://developers.google.com/actions/reference/ssml).
        /// </summary>
        [Input("ssml")]
        public Input<string>? Ssml { get; set; }

        /// <summary>
        /// The raw text to be synthesized.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public GoogleCloudDialogflowV2beta1IntentMessageTelephonySynthesizeSpeechArgs()
        {
        }
    }
}
