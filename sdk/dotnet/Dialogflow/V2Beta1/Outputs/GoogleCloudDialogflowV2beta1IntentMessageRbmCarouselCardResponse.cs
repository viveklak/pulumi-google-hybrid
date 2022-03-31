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
    /// Carousel Rich Business Messaging (RBM) rich card. Rich cards allow you to respond to users with more vivid content, e.g. with media and suggestions. If you want to show a single card with more control over the layout, please use RbmStandaloneCard instead.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageRbmCarouselCardResponse
    {
        /// <summary>
        /// The cards in the carousel. A carousel must have at least 2 cards and at most 10.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageRbmCardContentResponse> CardContents;
        /// <summary>
        /// The width of the cards in the carousel.
        /// </summary>
        public readonly string CardWidth;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1IntentMessageRbmCarouselCardResponse(
            ImmutableArray<Outputs.GoogleCloudDialogflowV2beta1IntentMessageRbmCardContentResponse> cardContents,

            string cardWidth)
        {
            CardContents = cardContents;
            CardWidth = cardWidth;
        }
    }
}
