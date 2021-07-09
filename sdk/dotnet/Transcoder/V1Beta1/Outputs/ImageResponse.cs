// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1.Outputs
{

    [OutputType]
    public sealed class ImageResponse
    {
        /// <summary>
        /// Target image opacity. Valid values: `1.0` (solid, default) to `0.0` (transparent).
        /// </summary>
        public readonly double Alpha;
        /// <summary>
        /// Normalized image resolution, based on output video resolution. Valid values: `0.0`–`1.0`. To respect the original image aspect ratio, set either `x` or `y` to `0.0`. To use the original image resolution, set both `x` and `y` to `0.0`.
        /// </summary>
        public readonly Outputs.NormalizedCoordinateResponse Resolution;
        /// <summary>
        /// URI of the JPEG image in Cloud Storage. For example, `gs://bucket/inputs/image.jpeg`. JPEG is the only supported image type.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private ImageResponse(
            double alpha,

            Outputs.NormalizedCoordinateResponse resolution,

            string uri)
        {
            Alpha = alpha;
            Resolution = resolution;
            Uri = uri;
        }
    }
}
