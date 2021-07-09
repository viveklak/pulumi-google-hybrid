// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowV2IntentMessageCarouselSelectItemResponse
    {
        /// <summary>
        /// Optional. The body text of the card.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. The image to display.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageImageResponse Image;
        /// <summary>
        /// Additional info about the option item.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2IntentMessageSelectItemInfoResponse Info;
        /// <summary>
        /// Title of the carousel item.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageCarouselSelectItemResponse(
            string description,

            Outputs.GoogleCloudDialogflowV2IntentMessageImageResponse image,

            Outputs.GoogleCloudDialogflowV2IntentMessageSelectItemInfoResponse info,

            string title)
        {
            Description = description;
            Image = image;
            Info = info;
            Title = title;
        }
    }
}
