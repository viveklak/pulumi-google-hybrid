// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    /// <summary>
    /// The rating of a Product.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaRatingResponse
    {
        /// <summary>
        /// The average rating of the Product. The rating is scaled at 1-5. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly double AverageRating;
        /// <summary>
        /// The total number of ratings. This value is independent of the value of rating_histogram. This value must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly int RatingCount;
        /// <summary>
        /// List of rating counts per rating value (index = rating - 1). The list is empty if there is no rating. If the list is non-empty, its size is always 5. Otherwise, an INVALID_ARGUMENT error is returned. For example, [41, 14, 13, 47, 303]. It means that the Product got 41 ratings with 1 star, 14 ratings with 2 star, and so on.
        /// </summary>
        public readonly ImmutableArray<int> RatingHistogram;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaRatingResponse(
            double averageRating,

            int ratingCount,

            ImmutableArray<int> ratingHistogram)
        {
            AverageRating = averageRating;
            RatingCount = ratingCount;
            RatingHistogram = ratingHistogram;
        }
    }
}
