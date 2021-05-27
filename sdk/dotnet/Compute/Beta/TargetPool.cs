// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates a target pool in the specified project and region using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:TargetPool")]
    public partial class TargetPool : Pulumi.CustomResource
    {
        /// <summary>
        /// The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1].
        /// 
        /// backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.
        /// 
        /// In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        /// </summary>
        [Output("backupPool")]
        public Output<string> BackupPool { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1].
        /// 
        /// If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.
        /// 
        /// In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        /// </summary>
        [Output("failoverRatio")]
        public Output<double> FailoverRatio { get; private set; } = null!;

        /// <summary>
        /// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
        /// </summary>
        [Output("healthChecks")]
        public Output<ImmutableArray<string>> HealthChecks { get; private set; } = null!;

        /// <summary>
        /// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<string>> Instances { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#targetPool for target pools.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the target pool resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Session affinity option, must be one of the following values:
        /// NONE: Connections from the same client IP may go to any instance in the pool.
        /// CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy.
        /// CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
        /// </summary>
        [Output("sessionAffinity")]
        public Output<string> SessionAffinity { get; private set; } = null!;


        /// <summary>
        /// Create a TargetPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetPool(string name, TargetPoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:TargetPool", name, args ?? new TargetPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:TargetPool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TargetPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TargetPool(name, id, options);
        }
    }

    public sealed class TargetPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The server-defined URL for the resource. This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool, and its failoverRatio field is properly set to a value between [0, 1].
        /// 
        /// backupPool and failoverRatio together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below failoverRatio, traffic arriving at the load-balanced IP will be directed to the backup pool.
        /// 
        /// In case where failoverRatio and backupPool are not set, or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        /// </summary>
        [Input("backupPool")]
        public Input<string>? BackupPool { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This field is applicable only when the containing target pool is serving a forwarding rule as the primary pool (i.e., not as a backup pool to some other target pool). The value of the field must be in [0, 1].
        /// 
        /// If set, backupPool must also be set. They together define the fallback behavior of the primary target pool: if the ratio of the healthy instances in the primary pool is at or below this number, traffic arriving at the load-balanced IP will be directed to the backup pool.
        /// 
        /// In case where failoverRatio is not set or all the instances in the backup pool are unhealthy, the traffic will be directed back to the primary pool in the "force" mode, where traffic will be spread to the healthy instances with the best effort, or to all instances when no instance is healthy.
        /// </summary>
        [Input("failoverRatio")]
        public Input<double>? FailoverRatio { get; set; }

        [Input("healthChecks")]
        private InputList<string>? _healthChecks;

        /// <summary>
        /// The URL of the HttpHealthCheck resource. A member instance in this pool is considered healthy if and only if the health checks pass. Only legacy HttpHealthChecks are supported. Only one health check may be specified.
        /// </summary>
        public InputList<string> HealthChecks
        {
            get => _healthChecks ?? (_healthChecks = new InputList<string>());
            set => _healthChecks = value;
        }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instances")]
        private InputList<string>? _instances;

        /// <summary>
        /// A list of resource URLs to the virtual machine instances serving this pool. They must live in zones contained in the same region as this pool.
        /// </summary>
        public InputList<string> Instances
        {
            get => _instances ?? (_instances = new InputList<string>());
            set => _instances = value;
        }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#targetPool for target pools.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the target pool resides.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Session affinity option, must be one of the following values:
        /// NONE: Connections from the same client IP may go to any instance in the pool.
        /// CLIENT_IP: Connections from the same client IP will go to the same instance in the pool while that instance remains healthy.
        /// CLIENT_IP_PROTO: Connections from the same client IP with the same IP protocol will go to the same instance in the pool while that instance remains healthy.
        /// </summary>
        [Input("sessionAffinity")]
        public Input<string>? SessionAffinity { get; set; }

        public TargetPoolArgs()
        {
        }
    }
}
