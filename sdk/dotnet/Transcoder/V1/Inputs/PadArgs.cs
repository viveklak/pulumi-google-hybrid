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
    /// Pad filter configuration for the input video. The padded input video is scaled after padding with black to match the output resolution.
    /// </summary>
    public sealed class PadArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of pixels to add to the bottom. The default is 0.
        /// </summary>
        [Input("bottomPixels")]
        public Input<int>? BottomPixels { get; set; }

        /// <summary>
        /// The number of pixels to add to the left. The default is 0.
        /// </summary>
        [Input("leftPixels")]
        public Input<int>? LeftPixels { get; set; }

        /// <summary>
        /// The number of pixels to add to the right. The default is 0.
        /// </summary>
        [Input("rightPixels")]
        public Input<int>? RightPixels { get; set; }

        /// <summary>
        /// The number of pixels to add to the top. The default is 0.
        /// </summary>
        [Input("topPixels")]
        public Input<int>? TopPixels { get; set; }

        public PadArgs()
        {
        }
    }
}
