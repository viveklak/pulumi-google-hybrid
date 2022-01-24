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
    public static class GetProductSet
    {
        /// <summary>
        /// Gets information associated with a ProductSet. Possible errors: * Returns NOT_FOUND if the ProductSet does not exist.
        /// </summary>
        public static Task<GetProductSetResult> InvokeAsync(GetProductSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProductSetResult>("google-native:vision/v1:getProductSet", args ?? new GetProductSetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information associated with a ProductSet. Possible errors: * Returns NOT_FOUND if the ProductSet does not exist.
        /// </summary>
        public static Output<GetProductSetResult> Invoke(GetProductSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProductSetResult>("google-native:vision/v1:getProductSet", args ?? new GetProductSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProductSetArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("productSetId", required: true)]
        public string ProductSetId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetProductSetArgs()
        {
        }
    }

    public sealed class GetProductSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("productSetId", required: true)]
        public Input<string> ProductSetId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetProductSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProductSetResult
    {
        /// <summary>
        /// The user-provided name for this ProductSet. Must not be empty. Must be at most 4096 characters long.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// If there was an error with indexing the product set, the field is populated. This field is ignored when creating a ProductSet.
        /// </summary>
        public readonly Outputs.StatusResponse IndexError;
        /// <summary>
        /// The time at which this ProductSet was last indexed. Query results will reflect all updates before this time. If this ProductSet has never been indexed, this timestamp is the default value "1970-01-01T00:00:00Z". This field is ignored when creating a ProductSet.
        /// </summary>
        public readonly string IndexTime;
        /// <summary>
        /// The resource name of the ProductSet. Format is: `projects/PROJECT_ID/locations/LOC_ID/productSets/PRODUCT_SET_ID`. This field is ignored when creating a ProductSet.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetProductSetResult(
            string displayName,

            Outputs.StatusResponse indexError,

            string indexTime,

            string name)
        {
            DisplayName = displayName;
            IndexError = indexError;
            IndexTime = indexTime;
            Name = name;
        }
    }
}
