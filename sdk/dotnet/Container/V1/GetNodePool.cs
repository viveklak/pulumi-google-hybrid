// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Container.V1
{
    public static class GetNodePool
    {
        /// <summary>
        /// Retrieves the requested node pool.
        /// </summary>
        public static Task<GetNodePoolResult> InvokeAsync(GetNodePoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodePoolResult>("google-native:container/v1:getNodePool", args ?? new GetNodePoolArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the requested node pool.
        /// </summary>
        public static Output<GetNodePoolResult> Invoke(GetNodePoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNodePoolResult>("google-native:container/v1:getNodePool", args ?? new GetNodePoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodePoolArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("nodePoolId", required: true)]
        public string NodePoolId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNodePoolArgs()
        {
        }
    }

    public sealed class GetNodePoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("nodePoolId", required: true)]
        public Input<string> NodePoolId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNodePoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNodePoolResult
    {
        /// <summary>
        /// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        /// </summary>
        public readonly Outputs.NodePoolAutoscalingResponse Autoscaling;
        /// <summary>
        /// Which conditions caused the current node pool state.
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusConditionResponse> Conditions;
        /// <summary>
        /// The node configuration of the pool.
        /// </summary>
        public readonly Outputs.NodeConfigResponse Config;
        /// <summary>
        /// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        /// </summary>
        public readonly int InitialNodeCount;
        /// <summary>
        /// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
        /// </summary>
        public readonly ImmutableArray<string> InstanceGroupUrls;
        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// NodeManagement configuration for this NodePool.
        /// </summary>
        public readonly Outputs.NodeManagementResponse Management;
        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        public readonly Outputs.MaxPodsConstraintResponse MaxPodsConstraint;
        /// <summary>
        /// The name of the node pool.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        /// </summary>
        public readonly Outputs.NodeNetworkConfigResponse NetworkConfig;
        /// <summary>
        /// [Output only] The pod CIDR block size per node in this node pool.
        /// </summary>
        public readonly int PodIpv4CidrSize;
        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Output only] The status of the nodes in this pool instance.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        public readonly Outputs.UpgradeSettingsResponse UpgradeSettings;
        /// <summary>
        /// The version of the Kubernetes of this node.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetNodePoolResult(
            Outputs.NodePoolAutoscalingResponse autoscaling,

            ImmutableArray<Outputs.StatusConditionResponse> conditions,

            Outputs.NodeConfigResponse config,

            int initialNodeCount,

            ImmutableArray<string> instanceGroupUrls,

            ImmutableArray<string> locations,

            Outputs.NodeManagementResponse management,

            Outputs.MaxPodsConstraintResponse maxPodsConstraint,

            string name,

            Outputs.NodeNetworkConfigResponse networkConfig,

            int podIpv4CidrSize,

            string selfLink,

            string status,

            Outputs.UpgradeSettingsResponse upgradeSettings,

            string version)
        {
            Autoscaling = autoscaling;
            Conditions = conditions;
            Config = config;
            InitialNodeCount = initialNodeCount;
            InstanceGroupUrls = instanceGroupUrls;
            Locations = locations;
            Management = management;
            MaxPodsConstraint = maxPodsConstraint;
            Name = name;
            NetworkConfig = networkConfig;
            PodIpv4CidrSize = podIpv4CidrSize;
            SelfLink = selfLink;
            Status = status;
            UpgradeSettings = upgradeSettings;
            Version = version;
        }
    }
}
