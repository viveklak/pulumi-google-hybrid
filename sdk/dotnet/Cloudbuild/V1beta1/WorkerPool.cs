// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Cloudbuild.V1beta1
{
    /// <summary>
    /// Creates a `WorkerPool` to run the builds, and returns the new worker pool. NOTE: As of now, this method returns an `Operation` that is always complete.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:cloudbuild/v1beta1:WorkerPool")]
    public partial class WorkerPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a WorkerPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkerPool(string name, WorkerPoolArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:cloudbuild/v1beta1:WorkerPool", name, args ?? new WorkerPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkerPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:cloudbuild/v1beta1:WorkerPool", name, null, MakeResourceOptions(options, id))
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
        /// Output only. Time at which the request to create the `WorkerPool` was received.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Output only. Time at which the request to delete the `WorkerPool` was received.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        /// <summary>
        /// Output only. The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network configuration for the `WorkerPool`.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NetworkConfigArgs>? NetworkConfig { get; set; }

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Output only. `WorkerPool` state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Output only. Time at which the request to update the `WorkerPool` was received.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Worker configuration for the `WorkerPool`.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.WorkerConfigArgs>? WorkerConfig { get; set; }

        [Input("workerPoolsId", required: true)]
        public Input<string> WorkerPoolsId { get; set; } = null!;

        public WorkerPoolArgs()
        {
        }
    }
}
