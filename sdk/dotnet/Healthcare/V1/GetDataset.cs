// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Healthcare.V1
{
    public static class GetDataset
    {
        /// <summary>
        /// Gets any metadata associated with a dataset.
        /// </summary>
        public static Task<GetDatasetResult> InvokeAsync(GetDatasetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatasetResult>("google-native:healthcare/v1:getDataset", args ?? new GetDatasetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets any metadata associated with a dataset.
        /// </summary>
        public static Output<GetDatasetResult> Invoke(GetDatasetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatasetResult>("google-native:healthcare/v1:getDataset", args ?? new GetDatasetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatasetArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDatasetArgs()
        {
        }
    }

    public sealed class GetDatasetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatasetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatasetResult
    {
        /// <summary>
        /// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private GetDatasetResult(
            string name,

            string timeZone)
        {
            Name = name;
            TimeZone = timeZone;
        }
    }
}
