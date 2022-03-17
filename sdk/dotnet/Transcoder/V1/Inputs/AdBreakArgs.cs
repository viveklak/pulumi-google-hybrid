// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Inputs
{

    /// <summary>
    /// Ad break.
    /// </summary>
    public sealed class AdBreakArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Start time in seconds for the ad break, relative to the output file timeline. The default is `0s`.
        /// </summary>
        [Input("startTimeOffset")]
        public Input<string>? StartTimeOffset { get; set; }

        public AdBreakArgs()
        {
        }
    }
}
