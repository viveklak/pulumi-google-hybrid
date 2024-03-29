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
    /// Category represents catalog item category hierarchy.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyResponse
    {
        /// <summary>
        /// Catalog item categories. Each category should be a UTF-8 encoded string with a length limit of 2 KiB. Note that the order in the list denotes the specificity (from least to most specific).
        /// </summary>
        public readonly ImmutableArray<string> Categories;

        [OutputConstructor]
        private GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyResponse(ImmutableArray<string> categories)
        {
            Categories = categories;
        }
    }
}
