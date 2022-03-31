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
    /// Table card for Actions on Google.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageTableCardResponse
    {
        /// <summary>
        /// Optional. List of buttons for the card.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonResponse> Buttons;
        /// <summary>
        /// Optional. Display properties for the columns in this table.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageColumnPropertiesResponse> ColumnProperties;
        /// <summary>
        /// Optional. Image which should be displayed on the card.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowV2beta1IntentMessageImageResponse Image;
        /// <summary>
        /// Optional. Rows in this table of data.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageTableCardRowResponse> Rows;
        /// <summary>
        /// Optional. Subtitle to the title.
        /// </summary>
        public readonly string Subtitle;
        /// <summary>
        /// Title of the card.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1IntentMessageTableCardResponse(
            ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonResponse> buttons,

            ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageColumnPropertiesResponse> columnProperties,

            Outputs.GoogleCloudDialogflowV2beta1IntentMessageImageResponse image,

            ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageTableCardRowResponse> rows,

            string subtitle,

            string title)
        {
            Buttons = buttons;
            ColumnProperties = columnProperties;
            Image = image;
            Rows = rows;
            Subtitle = subtitle;
            Title = title;
        }
    }
}
