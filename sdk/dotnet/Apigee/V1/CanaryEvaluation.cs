// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates a new canary evaluation for an organization.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:CanaryEvaluation")]
    public partial class CanaryEvaluation : Pulumi.CustomResource
    {
        /// <summary>
        /// The stable version that is serving requests.
        /// </summary>
        [Output("control")]
        public Output<string> Control { get; private set; } = null!;

        /// <summary>
        /// Create time of the canary evaluation.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// End time for the evaluation's analysis.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Labels used to filter the metrics used for a canary evaluation.
        /// </summary>
        [Output("metricLabels")]
        public Output<Outputs.GoogleCloudApigeeV1CanaryEvaluationMetricLabelsResponse> MetricLabels { get; private set; } = null!;

        /// <summary>
        /// Name of the canary evalution.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Start time for the canary evaluation's analysis.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The current state of the canary evaluation.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The newer version that is serving requests.
        /// </summary>
        [Output("treatment")]
        public Output<string> Treatment { get; private set; } = null!;

        /// <summary>
        /// The resulting verdict of the canary evaluations: NONE, PASS, or FAIL.
        /// </summary>
        [Output("verdict")]
        public Output<string> Verdict { get; private set; } = null!;


        /// <summary>
        /// Create a CanaryEvaluation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CanaryEvaluation(string name, CanaryEvaluationArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:CanaryEvaluation", name, args ?? new CanaryEvaluationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CanaryEvaluation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:CanaryEvaluation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CanaryEvaluation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CanaryEvaluation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CanaryEvaluation(name, id, options);
        }
    }

    public sealed class CanaryEvaluationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The stable version that is serving requests.
        /// </summary>
        [Input("control", required: true)]
        public Input<string> Control { get; set; } = null!;

        /// <summary>
        /// End time for the evaluation's analysis.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Labels used to filter the metrics used for a canary evaluation.
        /// </summary>
        [Input("metricLabels", required: true)]
        public Input<Inputs.GoogleCloudApigeeV1CanaryEvaluationMetricLabelsArgs> MetricLabels { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Start time for the canary evaluation's analysis.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        /// <summary>
        /// The newer version that is serving requests.
        /// </summary>
        [Input("treatment", required: true)]
        public Input<string> Treatment { get; set; } = null!;

        public CanaryEvaluationArgs()
        {
        }
    }
}
