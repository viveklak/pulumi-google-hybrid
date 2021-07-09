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
    public sealed class GoogleCloudDialogflowV2IntentMessageCarouselSelectResponse
    {
        /// <summary>
        /// Carousel items.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectItemResponse> Items;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageCarouselSelectResponse(ImmutableArray<Outputs.GoogleCloudDialogflowV2IntentMessageCarouselSelectItemResponse> items)
        {
            Items = items;
        }
    }
}
