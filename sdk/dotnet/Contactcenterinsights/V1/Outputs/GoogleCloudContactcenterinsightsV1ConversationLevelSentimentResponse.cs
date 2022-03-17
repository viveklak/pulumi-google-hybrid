// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Outputs
{

    /// <summary>
    /// One channel of conversation-level sentiment data.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1ConversationLevelSentimentResponse
    {
        /// <summary>
        /// The channel of the audio that the data applies to.
        /// </summary>
        public readonly int ChannelTag;
        /// <summary>
        /// Data specifying sentiment.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1SentimentDataResponse SentimentData;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1ConversationLevelSentimentResponse(
            int channelTag,

            Outputs.GoogleCloudContactcenterinsightsV1SentimentDataResponse sentimentData)
        {
            ChannelTag = channelTag;
            SentimentData = sentimentData;
        }
    }
}
