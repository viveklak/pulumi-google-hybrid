// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetTargetPool
    {
        /// <summary>
        /// Returns the specified target pool. Gets a list of available target pools by making a list() request.
        /// </summary>
        public static Task<GetTargetPoolResult> InvokeAsync(GetTargetPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTargetPoolResult>("google-native:compute/beta:getTargetPool", args ?? new GetTargetPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified target pool. Gets a list of available target pools by making a list() request.
        /// </summary>
        public static Output<GetTargetPoolResult> Invoke(GetTargetPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTargetPoolResult>("google-native:compute/beta:getTargetPool", args ?? new GetTargetPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetPoolArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("targetPool", required: true)]
        public string TargetPool { get; set; } = null!;

        public GetTargetPoolArgs()
        {
        }
    }

    public sealed class GetTargetPoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("targetPool", required: true)]
        public Input<string> TargetPool { get; set; } = null!;

        public GetTargetPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTargetPoolResult
    {
        /// <summary>
        /// The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1]. backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool. In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        /// </summary>
        public readonly string BackupPool;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1]. If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool. In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        /// </summary>
        public readonly double FailoverRatio;
        /// <summary>
        /// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
        /// </summary>
        public readonly ImmutableArray<string> HealthChecks;
        /// <summary>
        /// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
        /// </summary>
        public readonly ImmutableArray<string> Instances;
        /// <summary>
        /// Type of the resource. Always compute#targetPool for target pools.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the region where the target pool resides.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Session affinity option, must be one of the following values: NONE: Connections from the same client IP may go to any instance in the pool. CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy. CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
        /// </summary>
        public readonly string SessionAffinity;

        [OutputConstructor]
        private GetTargetPoolResult(
            string backupPool,

            string creationTimestamp,

            string description,

            double failoverRatio,

            ImmutableArray<string> healthChecks,

            ImmutableArray<string> instances,

            string kind,

            string name,

            string region,

            string selfLink,

            string sessionAffinity)
        {
            BackupPool = backupPool;
            CreationTimestamp = creationTimestamp;
            Description = description;
            FailoverRatio = failoverRatio;
            HealthChecks = healthChecks;
            Instances = instances;
            Kind = kind;
            Name = name;
            Region = region;
            SelfLink = selfLink;
            SessionAffinity = sessionAffinity;
        }
    }
}
