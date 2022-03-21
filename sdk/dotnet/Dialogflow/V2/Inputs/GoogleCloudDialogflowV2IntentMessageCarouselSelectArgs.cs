// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// The card for presenting a carousel of options to select from.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2IntentMessageCarouselSelectArgs : Pulumi.ResourceArgs
    {
        [Input("items", required: true)]
        private InputList<Inputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectItemArgs>? _items;

        /// <summary>
        /// Carousel items.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectItemArgs>());
            set => _items = value;
        }

        public GoogleCloudDialogflowV2IntentMessageCarouselSelectArgs()
        {
        }
    }
}
