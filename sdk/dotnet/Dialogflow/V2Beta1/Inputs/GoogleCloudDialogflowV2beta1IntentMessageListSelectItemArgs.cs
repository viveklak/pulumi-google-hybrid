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
    /// An item in the list.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageListSelectItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The main text describing the item.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. The image to display.
        /// </summary>
        [Input("image")]
        public Input<Inputs.GoogleCloudDialogflowV2beta1IntentMessageImageArgs>? Image { get; set; }

        /// <summary>
        /// Additional information about this option.
        /// </summary>
        [Input("info", required: true)]
        public Input<Inputs.GoogleCloudDialogflowV2beta1IntentMessageSelectItemInfoArgs> Info { get; set; } = null!;

        /// <summary>
        /// The title of the list item.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1IntentMessageListSelectItemArgs()
        {
        }
    }
}
