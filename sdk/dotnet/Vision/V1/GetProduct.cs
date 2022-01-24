// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Vision.V1
{
    public static class GetProduct
    {
        /// <summary>
        /// Gets information associated with a Product. Possible errors: * Returns NOT_FOUND if the Product does not exist.
        /// </summary>
        public static Task<GetProductResult> InvokeAsync(GetProductArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProductResult>("google-native:vision/v1:getProduct", args ?? new GetProductArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information associated with a Product. Possible errors: * Returns NOT_FOUND if the Product does not exist.
        /// </summary>
        public static Output<GetProductResult> Invoke(GetProductInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProductResult>("google-native:vision/v1:getProduct", args ?? new GetProductInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProductArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetProductArgs()
        {
        }
    }

    public sealed class GetProductInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetProductInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProductResult
    {
        /// <summary>
        /// User-provided metadata to be stored with this product. Must be at most 4096 characters long.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The user-provided name for this Product. Must not be empty. Must be at most 4096 characters long.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The resource name of the product. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID`. This field is ignored when creating a product.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The category for the product identified by the reference image. This should be one of "homegoods-v2", "apparel-v2", "toys-v2", "packagedgoods-v1" or "general-v1". The legacy categories "homegoods", "apparel", and "toys" are still supported, but these should not be used for new products.
        /// </summary>
        public readonly string ProductCategory;
        /// <summary>
        /// Key-value pairs that can be attached to a product. At query time, constraints can be specified based on the product_labels. Note that integer values can be provided as strings, e.g. "1199". Only strings with integer values can match a range-based restriction which is to be supported soon. Multiple values can be assigned to the same key. One product may have up to 500 product_labels. Notice that the total number of distinct product_labels over all products in one ProductSet cannot exceed 1M, otherwise the product search pipeline will refuse to work for that ProductSet.
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyValueResponse> ProductLabels;

        [OutputConstructor]
        private GetProductResult(
            string description,

            string displayName,

            string name,

            string productCategory,

            ImmutableArray<Outputs.KeyValueResponse> productLabels)
        {
            Description = description;
            DisplayName = displayName;
            Name = name;
            ProductCategory = productCategory;
            ProductLabels = productLabels;
        }
    }
}
