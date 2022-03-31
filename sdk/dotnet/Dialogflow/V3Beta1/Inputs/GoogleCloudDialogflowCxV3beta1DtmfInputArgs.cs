// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Inputs
{

    /// <summary>
    /// Represents the input for dtmf event.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1DtmfInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dtmf digits.
        /// </summary>
        [Input("digits")]
        public Input<string>? Digits { get; set; }

        /// <summary>
        /// The finish digit (if any).
        /// </summary>
        [Input("finishDigit")]
        public Input<string>? FinishDigit { get; set; }

        public GoogleCloudDialogflowCxV3beta1DtmfInputArgs()
        {
        }
    }
}
