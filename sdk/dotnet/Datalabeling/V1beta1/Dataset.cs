// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Datalabeling.V1beta1
{
    /// <summary>
    /// Creates dataset. If success return a Dataset resource.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:datalabeling/v1beta1:Dataset")]
    public partial class Dataset : Pulumi.CustomResource
    {
        /// <summary>
        /// The names of any related resources that are blocking changes to the dataset.
        /// </summary>
        [Output("blockingResources")]
        public Output<ImmutableArray<string>> BlockingResources { get; private set; } = null!;

        /// <summary>
        /// Time the dataset is created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The number of data items in the dataset.
        /// </summary>
        [Output("dataItemCount")]
        public Output<string> DataItemCount { get; private set; } = null!;

        /// <summary>
        /// Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Required. The display name of the dataset. Maximum of 64 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// This is populated with the original input configs where ImportData is called. It is available only after the clients import data to this dataset.
        /// </summary>
        [Output("inputConfigs")]
        public Output<ImmutableArray<Outputs.GoogleCloudDatalabelingV1beta1InputConfigResponse>> InputConfigs { get; private set; } = null!;

        /// <summary>
        /// Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
        /// </summary>
        [Output("lastMigrateTime")]
        public Output<string> LastMigrateTime { get; private set; } = null!;

        /// <summary>
        /// Dataset resource name, format is: projects/{project_id}/datasets/{dataset_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:datalabeling/v1beta1:Dataset", name, args ?? new DatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:datalabeling/v1beta1:Dataset", name, null, MakeResourceOptions(options, id))
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
        [Input("blockingResources")]
        private InputList<string>? _blockingResources;

        /// <summary>
        /// The names of any related resources that are blocking changes to the dataset.
        /// </summary>
        public InputList<string> BlockingResources
        {
            get => _blockingResources ?? (_blockingResources = new InputList<string>());
            set => _blockingResources = value;
        }

        /// <summary>
        /// Time the dataset is created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The number of data items in the dataset.
        /// </summary>
        [Input("dataItemCount")]
        public Input<string>? DataItemCount { get; set; }

        [Input("datasetsId", required: true)]
        public Input<string> DatasetsId { get; set; } = null!;

        /// <summary>
        /// Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The display name of the dataset. Maximum of 64 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("inputConfigs")]
        private InputList<Inputs.GoogleCloudDatalabelingV1beta1InputConfigArgs>? _inputConfigs;

        /// <summary>
        /// This is populated with the original input configs where ImportData is called. It is available only after the clients import data to this dataset.
        /// </summary>
        public InputList<Inputs.GoogleCloudDatalabelingV1beta1InputConfigArgs> InputConfigs
        {
            get => _inputConfigs ?? (_inputConfigs = new InputList<Inputs.GoogleCloudDatalabelingV1beta1InputConfigArgs>());
            set => _inputConfigs = value;
        }

        /// <summary>
        /// Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
        /// </summary>
        [Input("lastMigrateTime")]
        public Input<string>? LastMigrateTime { get; set; }

        /// <summary>
        /// Dataset resource name, format is: projects/{project_id}/datasets/{dataset_id}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        public DatasetArgs()
        {
        }
    }
}