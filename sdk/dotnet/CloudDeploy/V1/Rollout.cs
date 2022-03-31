// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1
{
    /// <summary>
    /// Creates a new Rollout in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:clouddeploy/v1:Rollout")]
    public partial class Rollout : Pulumi.CustomResource
    {
        /// <summary>
        /// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>> Annotations { get; private set; } = null!;

        /// <summary>
        /// Approval state of the `Rollout`.
        /// </summary>
        [Output("approvalState")]
        public Output<string> ApprovalState { get; private set; } = null!;

        /// <summary>
        /// Time at which the `Rollout` was approved.
        /// </summary>
        [Output("approveTime")]
        public Output<string> ApproveTime { get; private set; } = null!;

        /// <summary>
        /// Time at which the `Rollout` was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Time at which the `Rollout` finished deploying.
        /// </summary>
        [Output("deployEndTime")]
        public Output<string> DeployEndTime { get; private set; } = null!;

        /// <summary>
        /// The reason this deploy failed. This will always be unspecified while the deploy in progress.
        /// </summary>
        [Output("deployFailureCause")]
        public Output<string> DeployFailureCause { get; private set; } = null!;

        /// <summary>
        /// Time at which the `Rollout` started deploying.
        /// </summary>
        [Output("deployStartTime")]
        public Output<string> DeployStartTime { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Cloud Build `Build` object that is used to deploy the Rollout. Format is `projects/{project}/locations/{location}/builds/{build}`.
        /// </summary>
        [Output("deployingBuild")]
        public Output<string> DeployingBuild { get; private set; } = null!;

        /// <summary>
        /// Description of the `Rollout` for user purposes. Max length is 255 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Time at which the `Rollout` was enqueued.
        /// </summary>
        [Output("enqueueTime")]
        public Output<string> EnqueueTime { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Reason the build failed. Empty if the build succeeded.
        /// </summary>
        [Output("failureReason")]
        public Output<string> FailureReason { get; private set; } = null!;

        /// <summary>
        /// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Optional. Name of the `Rollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Current state of the `Rollout`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The ID of Target to which this `Rollout` is deploying.
        /// </summary>
        [Output("targetId")]
        public Output<string> TargetId { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the `Rollout`.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;


        /// <summary>
        /// Create a Rollout resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rollout(string name, RolloutArgs args, CustomResourceOptions? options = null)
            : base("google-native:clouddeploy/v1:Rollout", name, args ?? new RolloutArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rollout(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:clouddeploy/v1:Rollout", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Rollout resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rollout Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Rollout(name, id, options);
        }
    }

    public sealed class RolloutArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        [Input("deliveryPipelineId", required: true)]
        public Input<string> DeliveryPipelineId { get; set; } = null!;

        /// <summary>
        /// Description of the `Rollout` for user purposes. Max length is 255 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 128 bytes.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. Name of the `Rollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("releaseId", required: true)]
        public Input<string> ReleaseId { get; set; } = null!;

        /// <summary>
        /// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Required. ID of the `Rollout`.
        /// </summary>
        [Input("rolloutId", required: true)]
        public Input<string> RolloutId { get; set; } = null!;

        /// <summary>
        /// The ID of Target to which this `Rollout` is deploying.
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        /// <summary>
        /// Optional. If set to true, the request is validated and the user is provided with an expected result, but no actual change is made.
        /// </summary>
        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        public RolloutArgs()
        {
        }
    }
}
