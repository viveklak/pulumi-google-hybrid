// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    /// <summary>
    /// Creates a lake resource.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:Lake")]
    public partial class Lake : Pulumi.CustomResource
    {
        /// <summary>
        /// Aggregated status of the underlying assets of the lake.
        /// </summary>
        [Output("assetStatus")]
        public Output<Outputs.GoogleCloudDataplexV1AssetStatusResponse> AssetStatus { get; private set; } = null!;

        /// <summary>
        /// The time when the lake was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the lake.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. User-defined labels for the lake.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. Settings to manage lake and Dataproc Metastore service instance association.
        /// </summary>
        [Output("metastore")]
        public Output<Outputs.GoogleCloudDataplexV1LakeMetastoreResponse> Metastore { get; private set; } = null!;

        /// <summary>
        /// Metastore status of the lake.
        /// </summary>
        [Output("metastoreStatus")]
        public Output<Outputs.GoogleCloudDataplexV1LakeMetastoreStatusResponse> MetastoreStatus { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of the lake, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Current state of the lake.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the lake was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Lake resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lake(string name, LakeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Lake", name, args ?? new LakeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lake(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Lake", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Lake resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lake Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Lake(name, id, options);
        }
    }

    public sealed class LakeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the lake.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User-defined labels for the lake.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Required. Lake identifier. This ID will be used to generate names such as database and dataset names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the customer project / location.
        /// </summary>
        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. Settings to manage lake and Dataproc Metastore service instance association.
        /// </summary>
        [Input("metastore")]
        public Input<Inputs.GoogleCloudDataplexV1LakeMetastoreArgs>? Metastore { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Only validate the request, but do not perform mutations. The default is false.
        /// </summary>
        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        public LakeArgs()
        {
        }
    }
}
