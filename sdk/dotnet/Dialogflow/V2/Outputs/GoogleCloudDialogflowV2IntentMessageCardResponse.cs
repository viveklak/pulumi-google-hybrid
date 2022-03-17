// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    /// <summary>
    /// The card response message.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2IntentMessageCardResponse
    {
        /// <summary>
        /// Optional. The collection of card buttons.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageCardButtonResponse> Buttons;
        /// <summary>
        /// Optional. The public URI to an image file for the card.
        /// </summary>
        public readonly string ImageUri;
        /// <summary>
        /// Optional. The subtitle of the card.
        /// </summary>
        public readonly string Subtitle;
        /// <summary>
        /// Optional. The title of the card.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageCardResponse(
            ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageCardButtonResponse> buttons,

            string imageUri,

            string subtitle,

            string title)
        {
            Buttons = buttons;
            ImageUri = imageUri;
            Subtitle = subtitle;
            Title = title;
        }
    }
}
