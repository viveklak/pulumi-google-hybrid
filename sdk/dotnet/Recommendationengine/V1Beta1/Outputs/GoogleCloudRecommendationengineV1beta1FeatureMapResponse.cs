// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Recommendationengine.V1Beta1.Outputs
{

    /// <summary>
    /// FeatureMap represents extra features that customers want to include in the recommendation model for catalogs/user events as categorical/numerical features.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecommendationengineV1beta1FeatureMapResponse
    {
        /// <summary>
        /// Categorical features that can take on one of a limited number of possible values. Some examples would be the brand/maker of a product, or country of a customer. Feature names and values must be UTF-8 encoded strings. For example: `{ "colors": {"value": ["yellow", "green"]}, "sizes": {"value":["S", "M"]}`
        /// </summary>
        public readonly ImmutableDictionary<string, string> CategoricalFeatures;
        /// <summary>
        /// Numerical features. Some examples would be the height/weight of a product, or age of a customer. Feature names must be UTF-8 encoded strings. For example: `{ "lengths_cm": {"value":[2.3, 15.4]}, "heights_cm": {"value":[8.1, 6.4]} }`
        /// </summary>
        public readonly ImmutableDictionary<string, string> NumericalFeatures;

        [OutputConstructor]
        private GoogleCloudRecommendationengineV1beta1FeatureMapResponse(
            ImmutableDictionary<string, string> categoricalFeatures,

            ImmutableDictionary<string, string> numericalFeatures)
        {
            CategoricalFeatures = categoricalFeatures;
            NumericalFeatures = numericalFeatures;
        }
    }
}
