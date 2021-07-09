// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    [OutputType]
    public sealed class GoogleCloudRetailV2alphaImageResponse
    {
        /// <summary>
        /// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly int Height;
        /// <summary>
        /// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly int Width;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaImageResponse(
            int height,

            string uri,

            int width)
        {
            Height = height;
            Uri = uri;
            Width = width;
        }
    }
}
