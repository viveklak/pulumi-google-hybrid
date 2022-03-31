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
    /// The basic card message. Useful for displaying information.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageBasicCardArgs : Pulumi.ResourceArgs
    {
        [Input("buttons")]
        private InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonArgs>? _buttons;

        /// <summary>
        /// Optional. The collection of card buttons.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonArgs> Buttons
        {
            get => _buttons ?? (_buttons = new InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageBasicCardButtonArgs>());
            set => _buttons = value;
        }

        /// <summary>
        /// Required, unless image is present. The body text of the card.
        /// </summary>
        [Input("formattedText")]
        public Input<string>? FormattedText { get; set; }

        /// <summary>
        /// Optional. The image for the card.
        /// </summary>
        [Input("image")]
        public Input<Inputs.GoogleCloudDialogflowV2beta1IntentMessageImageArgs>? Image { get; set; }

        /// <summary>
        /// Optional. The subtitle of the card.
        /// </summary>
        [Input("subtitle")]
        public Input<string>? Subtitle { get; set; }

        /// <summary>
        /// Optional. The title of the card.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public GoogleCloudDialogflowV2beta1IntentMessageBasicCardArgs()
        {
        }
    }
}
