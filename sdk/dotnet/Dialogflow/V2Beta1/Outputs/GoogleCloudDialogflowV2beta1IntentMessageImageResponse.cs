// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// The image response message.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageImageResponse
    {
        /// <summary>
        /// A text description of the image to be used for accessibility, e.g., screen readers. Required if image_uri is set for CarouselSelect.
        /// </summary>
        public readonly string AccessibilityText;
        /// <summary>
        /// Optional. The public URI to an image file.
        /// </summary>
        public readonly string ImageUri;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1IntentMessageImageResponse(
            string accessibilityText,

            string imageUri)
        {
            AccessibilityText = accessibilityText;
            ImageUri = imageUri;
        }
    }
}
