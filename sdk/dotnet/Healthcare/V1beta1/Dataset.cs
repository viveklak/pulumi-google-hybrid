// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Healthcare.V1beta1
{
    /// <summary>
    /// Creates a new health dataset. Results are returned through the Operation interface which returns either an `Operation.response` which contains a Dataset or `Operation.error`. The metadata field type is OperationMetadata.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:healthcare/v1beta1:Dataset")]
    public partial class Dataset : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:healthcare/v1beta1:Dataset", name, args ?? new DatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:healthcare/v1beta1:Dataset", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Dataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Dataset(name, id, options);
        }
    }

    public sealed class DatasetArgs : Pulumi.ResourceArgs
    {
        [Input("datasetsId", required: true)]
        public Input<string> DatasetsId { get; set; } = null!;

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public DatasetArgs()
        {
        }
    }
}
