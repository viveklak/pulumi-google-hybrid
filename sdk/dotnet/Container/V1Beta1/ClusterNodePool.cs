// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1
{
    /// <summary>
    /// Creates a node pool for a cluster.
    /// </summary>
    [GoogleNativeResourceType("google-native:container/v1beta1:ClusterNodePool")]
    public partial class ClusterNodePool : Pulumi.CustomResource
    {
        /// <summary>
        /// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        /// </summary>
        [Output("autoscaling")]
        public Output<Outputs.NodePoolAutoscalingResponse> Autoscaling { get; private set; } = null!;

        /// <summary>
        /// Which conditions caused the current node pool state.
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.StatusConditionResponse>> Conditions { get; private set; } = null!;

        /// <summary>
        /// The node configuration of the pool.
        /// </summary>
        [Output("config")]
        public Output<Outputs.NodeConfigResponse> Config { get; private set; } = null!;

        /// <summary>
        /// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        /// </summary>
        [Output("initialNodeCount")]
        public Output<int> InitialNodeCount { get; private set; } = null!;

        /// <summary>
        /// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
        /// </summary>
        [Output("instanceGroupUrls")]
        public Output<ImmutableArray<string>> InstanceGroupUrls { get; private set; } = null!;

        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<string>> Locations { get; private set; } = null!;

        /// <summary>
        /// NodeManagement configuration for this NodePool.
        /// </summary>
        [Output("management")]
        public Output<Outputs.NodeManagementResponse> Management { get; private set; } = null!;

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Output("maxPodsConstraint")]
        public Output<Outputs.MaxPodsConstraintResponse> MaxPodsConstraint { get; private set; } = null!;

        /// <summary>
        /// The name of the node pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        /// </summary>
        [Output("networkConfig")]
        public Output<Outputs.NodeNetworkConfigResponse> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// [Output only] The pod CIDR block size per node in this node pool.
        /// </summary>
        [Output("podIpv4CidrSize")]
        public Output<int> PodIpv4CidrSize { get; private set; } = null!;

        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output only] The status of the nodes in this pool instance.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        [Output("upgradeSettings")]
        public Output<Outputs.UpgradeSettingsResponse> UpgradeSettings { get; private set; } = null!;

        /// <summary>
        /// The version of the Kubernetes of this node.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterNodePool(string name, ClusterNodePoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:container/v1beta1:ClusterNodePool", name, args ?? new ClusterNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterNodePool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:container/v1beta1:ClusterNodePool", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterNodePool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClusterNodePool(name, id, options);
        }
    }

    public sealed class ClusterNodePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
        /// </summary>
        [Input("autoscaling")]
        public Input<Inputs.NodePoolAutoscalingArgs>? Autoscaling { get; set; }

        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("conditions")]
        private InputList<Inputs.StatusConditionArgs>? _conditions;

        /// <summary>
        /// Which conditions caused the current node pool state.
        /// </summary>
        public InputList<Inputs.StatusConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.StatusConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// The node configuration of the pool.
        /// </summary>
        [Input("config")]
        public Input<Inputs.NodeConfigArgs>? Config { get; set; }

        /// <summary>
        /// The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
        /// </summary>
        [Input("initialNodeCount")]
        public Input<int>? InitialNodeCount { get; set; }

        [Input("instanceGroupUrls")]
        private InputList<string>? _instanceGroupUrls;

        /// <summary>
        /// [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
        /// </summary>
        public InputList<string> InstanceGroupUrls
        {
            get => _instanceGroupUrls ?? (_instanceGroupUrls = new InputList<string>());
            set => _instanceGroupUrls = value;
        }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// NodeManagement configuration for this NodePool.
        /// </summary>
        [Input("management")]
        public Input<Inputs.NodeManagementArgs>? Management { get; set; }

        /// <summary>
        /// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
        /// </summary>
        [Input("maxPodsConstraint")]
        public Input<Inputs.MaxPodsConstraintArgs>? MaxPodsConstraint { get; set; }

        /// <summary>
        /// The name of the node pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
        /// </summary>
        [Input("networkConfig")]
        public Input<Inputs.NodeNetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*/locations/*/clusters/*`.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// [Output only] The pod CIDR block size per node in this node pool.
        /// </summary>
        [Input("podIpv4CidrSize")]
        public Input<int>? PodIpv4CidrSize { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output only] The status of the nodes in this pool instance.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Upgrade settings control disruption and speed of the upgrade.
        /// </summary>
        [Input("upgradeSettings")]
        public Input<Inputs.UpgradeSettingsArgs>? UpgradeSettings { get; set; }

        /// <summary>
        /// The version of the Kubernetes of this node.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ClusterNodePoolArgs()
        {
        }
    }
}
