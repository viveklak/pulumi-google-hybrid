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
    /// Exact product price.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecommendationengineV1beta1ProductCatalogItemExactPriceResponse
    {
        /// <summary>
        /// Optional. Display price of the product.
        /// </summary>
        public readonly double DisplayPrice;
        /// <summary>
        /// Optional. Price of the product without any discount. If zero, by default set to be the 'displayPrice'.
        /// </summary>
        public readonly double OriginalPrice;

        [OutputConstructor]
        private GoogleCloudRecommendationengineV1beta1ProductCatalogItemExactPriceResponse(
            double displayPrice,

            double originalPrice)
        {
            DisplayPrice = displayPrice;
            OriginalPrice = originalPrice;
        }
    }
}
