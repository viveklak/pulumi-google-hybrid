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
    public static class GetDicomStore
    {
        /// <summary>
        /// Gets the specified DICOM store.
        /// </summary>
        public static Task<GetDicomStoreResult> InvokeAsync(GetDicomStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDicomStoreResult>("google-native:healthcare/v1beta1:getDicomStore", args ?? new GetDicomStoreArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified DICOM store.
        /// </summary>
        public static Output<GetDicomStoreResult> Invoke(GetDicomStoreInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDicomStoreResult>("google-native:healthcare/v1beta1:getDicomStore", args ?? new GetDicomStoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDicomStoreArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("dicomStoreId", required: true)]
        public string DicomStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDicomStoreArgs()
        {
        }
    }

    public sealed class GetDicomStoreInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("dicomStoreId", required: true)]
        public Input<string> DicomStoreId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDicomStoreInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDicomStoreResult
    {
        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notification destination for new DICOM instances. Supplied by the client.
        /// </summary>
        public readonly Outputs.NotificationConfigResponse NotificationConfig;
        /// <summary>
        /// A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudHealthcareV1beta1DicomStreamConfigResponse> StreamConfigs;

        [OutputConstructor]
        private GetDicomStoreResult(
            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.NotificationConfigResponse notificationConfig,

            ImmutableArray<Outputs.GoogleCloudHealthcareV1beta1DicomStreamConfigResponse> streamConfigs)
        {
            Labels = labels;
            Name = name;
            NotificationConfig = notificationConfig;
            StreamConfigs = streamConfigs;
        }
    }
}
