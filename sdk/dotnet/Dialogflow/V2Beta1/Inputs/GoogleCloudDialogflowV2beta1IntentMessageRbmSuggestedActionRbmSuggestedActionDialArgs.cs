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
    /// Opens the user's default dialer app with the specified phone number but does not dial automatically.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestedActionRbmSuggestedActionDialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The phone number to fill in the default dialer app. This field should be in [E.164](https://en.wikipedia.org/wiki/E.164) format. An example of a correctly formatted phone number: +15556767888.
        /// </summary>
        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestedActionRbmSuggestedActionDialArgs()
        {
        }
    }
}
