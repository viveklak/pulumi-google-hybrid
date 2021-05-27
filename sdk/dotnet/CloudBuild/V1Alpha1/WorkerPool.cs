// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1Alpha1
{
    /// <summary>
    /// Creates a `WorkerPool` to run the builds, and returns the new worker pool.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudbuild/v1alpha1:WorkerPool")]
    public partial class WorkerPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Time at which the request to create the `WorkerPool` was received.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Time at which the request to delete the `WorkerPool` was received.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// User-defined name of the `WorkerPool`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project ID of the GCP project for which the `WorkerPool` is created.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// List of regions to create the `WorkerPool`. Regions can't be empty. If Cloud Build adds a new GCP region in the future, the existing `WorkerPool` will not be enabled in the new region automatically; you must add the new region to the `regions` field to enable the `WorkerPool` in that region.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// The service account used to manage the `WorkerPool`. The service account must have the Compute Instance Admin (Beta) permission at the project level.
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// WorkerPool Status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Time at which the request to update the `WorkerPool` was received.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Configuration to be used for a creating workers in the `WorkerPool`.
        /// </summary>
        [Output("workerConfig")]
        public Output<Outputs.WorkerConfigResponse> WorkerConfig { get; private set; } = null!;

        /// <summary>
        /// Total number of workers to be created across all requested regions.
        /// </summary>
        [Output("workerCount")]
        public Output<string> WorkerCount { get; private set; } = null!;


        /// <summary>
        /// Create a WorkerPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkerPool(string name, WorkerPoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudbuild/v1alpha1:WorkerPool", name, args ?? new WorkerPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkerPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudbuild/v1alpha1:WorkerPool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WorkerPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkerPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkerPool(name, id, options);
        }
    }

    public sealed class WorkerPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time at which the request to create the `WorkerPool` was received.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Time at which the request to delete the `WorkerPool` was received.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// User-defined name of the `WorkerPool`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID of the GCP project for which the `WorkerPool` is created.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// List of regions to create the `WorkerPool`. Regions can't be empty. If Cloud Build adds a new GCP region in the future, the existing `WorkerPool` will not be enabled in the new region automatically; you must add the new region to the `regions` field to enable the `WorkerPool` in that region.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// The service account used to manage the `WorkerPool`. The service account must have the Compute Instance Admin (Beta) permission at the project level.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// WorkerPool Status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Time at which the request to update the `WorkerPool` was received.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Configuration to be used for a creating workers in the `WorkerPool`.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.WorkerConfigArgs>? WorkerConfig { get; set; }

        /// <summary>
        /// Total number of workers to be created across all requested regions.
        /// </summary>
        [Input("workerCount")]
        public Input<string>? WorkerCount { get; set; }

        public WorkerPoolArgs()
        {
        }
    }
}
