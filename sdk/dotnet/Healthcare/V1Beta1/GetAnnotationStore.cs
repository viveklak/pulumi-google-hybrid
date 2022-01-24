// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    public static class GetAnnotationStore
    {
        /// <summary>
        /// Gets the specified Annotation store or returns NOT_FOUND if it does not exist.
        /// </summary>
        public static Task<GetAnnotationStoreResult> InvokeAsync(GetAnnotationStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAnnotationStoreResult>("google-native:healthcare/v1beta1:getAnnotationStore", args ?? new GetAnnotationStoreArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified Annotation store or returns NOT_FOUND if it does not exist.
        /// </summary>
        public static Output<GetAnnotationStoreResult> Invoke(GetAnnotationStoreInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAnnotationStoreResult>("google-native:healthcare/v1beta1:getAnnotationStore", args ?? new GetAnnotationStoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAnnotationStoreArgs : Pulumi.InvokeArgs
    {
        [Input("annotationStoreId", required: true)]
        public string AnnotationStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAnnotationStoreArgs()
        {
        }
    }

    public sealed class GetAnnotationStoreInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("annotationStoreId", required: true)]
        public Input<string> AnnotationStoreId { get; set; } = null!;

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAnnotationStoreInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAnnotationStoreResult
    {
        /// <summary>
        /// Optional. User-supplied key-value pairs used to organize Annotation stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name of the Annotation store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetAnnotationStoreResult(
            ImmutableDictionary<string, string> labels,

            string name)
        {
            Labels = labels;
            Name = name;
        }
    }
}
