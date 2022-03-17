// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetRegionAutoscaler
    {
        /// <summary>
        /// Returns the specified autoscaler.
        /// </summary>
        public static Task<GetRegionAutoscalerResult> InvokeAsync(GetRegionAutoscalerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionAutoscalerResult>("google-native:compute/v1:getRegionAutoscaler", args ?? new GetRegionAutoscalerArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified autoscaler.
        /// </summary>
        public static Output<GetRegionAutoscalerResult> Invoke(GetRegionAutoscalerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionAutoscalerResult>("google-native:compute/v1:getRegionAutoscaler", args ?? new GetRegionAutoscalerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionAutoscalerArgs : Pulumi.InvokeArgs
    {
        [Input("autoscaler", required: true)]
        public string Autoscaler { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionAutoscalerArgs()
        {
        }
    }

    public sealed class GetRegionAutoscalerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("autoscaler", required: true)]
        public Input<string> Autoscaler { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetRegionAutoscalerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionAutoscalerResult
    {
        /// <summary>
        /// The configuration parameters for the autoscaling algorithm. You can define one or more signals for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
        /// </summary>
        public readonly Outputs.AutoscalingPolicyResponse AutoscalingPolicy;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Type of the resource. Always compute#autoscaler for autoscalers.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
        /// </summary>
        public readonly int RecommendedSize;
        /// <summary>
        /// URL of the region where the instance group resides (for autoscalers living in regional scope).
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Status information of existing scaling schedules.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ScalingScheduleStatus;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The status of the autoscaler configuration. Current set of possible values: - PENDING: Autoscaler backend hasn't read new/updated configuration. - DELETING: Configuration is being deleted. - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field. New values might be added in the future.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
        /// </summary>
        public readonly ImmutableArray<Outputs.AutoscalerStatusDetailsResponse> StatusDetails;
        /// <summary>
        /// URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// URL of the zone where the instance group resides (for autoscalers living in zonal scope).
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetRegionAutoscalerResult(
            Outputs.AutoscalingPolicyResponse autoscalingPolicy,

            string creationTimestamp,

            string description,

            string kind,

            string name,

            int recommendedSize,

            string region,

            ImmutableDictionary<string, string> scalingScheduleStatus,

            string selfLink,

            string status,

            ImmutableArray<Outputs.AutoscalerStatusDetailsResponse> statusDetails,

            string target,

            string zone)
        {
            AutoscalingPolicy = autoscalingPolicy;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            RecommendedSize = recommendedSize;
            Region = region;
            ScalingScheduleStatus = scalingScheduleStatus;
            SelfLink = selfLink;
            Status = status;
            StatusDetails = statusDetails;
            Target = target;
            Zone = zone;
        }
    }
}
