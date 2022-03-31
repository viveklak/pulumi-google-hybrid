// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1
{
    public static class GetCluster
    {
        /// <summary>
        /// Gets the details for a specific cluster.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("google-native:container/v1beta1:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details for a specific cluster.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetClusterResult>("google-native:container/v1beta1:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// Configurations for the various addons available to run in the cluster.
        /// </summary>
        public readonly Outputs.AddonsConfigResponse AddonsConfig;
        /// <summary>
        /// Configuration controlling RBAC group membership information.
        /// </summary>
        public readonly Outputs.AuthenticatorGroupsConfigResponse AuthenticatorGroupsConfig;
        /// <summary>
        /// Autopilot configuration for the cluster.
        /// </summary>
        public readonly Outputs.AutopilotResponse Autopilot;
        /// <summary>
        /// Cluster-level autoscaling configuration.
        /// </summary>
        public readonly Outputs.ClusterAutoscalingResponse Autoscaling;
        /// <summary>
        /// Configuration for Binary Authorization.
        /// </summary>
        public readonly Outputs.BinaryAuthorizationResponse BinaryAuthorization;
        /// <summary>
        /// The IP address range of the container pods in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically chosen or specify a `/14` block in `10.0.0.0/8`.
        /// </summary>
        public readonly string ClusterIpv4Cidr;
        /// <summary>
        /// Telemetry integration for the cluster.
        /// </summary>
        public readonly Outputs.ClusterTelemetryResponse ClusterTelemetry;
        /// <summary>
        /// Which conditions caused the current cluster state.
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusConditionResponse> Conditions;
        /// <summary>
        /// Configuration of Confidential Nodes. All the nodes in the cluster will be Confidential VM once enabled.
        /// </summary>
        public readonly Outputs.ConfidentialNodesResponse ConfidentialNodes;
        /// <summary>
        /// [Output only] The time the cluster was created, in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// [Output only] The current software version of the master endpoint.
        /// </summary>
        public readonly string CurrentMasterVersion;
        /// <summary>
        /// [Output only] The number of nodes currently in the cluster. Deprecated. Call Kubernetes API directly to retrieve node information.
        /// </summary>
        public readonly int CurrentNodeCount;
        /// <summary>
        /// [Output only] Deprecated, use [NodePool.version](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools) instead. The current version of the node software components. If they are currently at multiple versions because they're in the process of being upgraded, this reflects the minimum version of all nodes.
        /// </summary>
        public readonly string CurrentNodeVersion;
        /// <summary>
        /// Configuration of etcd encryption.
        /// </summary>
        public readonly Outputs.DatabaseEncryptionResponse DatabaseEncryption;
        /// <summary>
        /// The default constraint on the maximum number of pods that can be run simultaneously on a node in the node pool of this cluster. Only honored if cluster created with IP Alias support.
        /// </summary>
        public readonly Outputs.MaxPodsConstraintResponse DefaultMaxPodsConstraint;
        /// <summary>
        /// An optional description of this cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Kubernetes alpha features are enabled on this cluster. This includes alpha API groups (e.g. v1beta1) and features that may not be production ready in the kubernetes version of the master and nodes. The cluster has no SLA for uptime and master/node upgrades are disabled. Alpha enabled clusters are automatically deleted thirty days after creation.
        /// </summary>
        public readonly bool EnableKubernetesAlpha;
        /// <summary>
        /// Enable the ability to use Cloud TPUs in this cluster. This field is deprecated, use tpu_config.enabled instead.
        /// </summary>
        public readonly bool EnableTpu;
        /// <summary>
        /// [Output only] The IP address of this cluster's master endpoint. The endpoint can be accessed from the internet at `https://username:password@endpoint/`. See the `masterAuth` property of this resource for username and password information.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// [Output only] The time the cluster will be automatically deleted in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// Configuration for Identity Service component.
        /// </summary>
        public readonly Outputs.IdentityServiceConfigResponse IdentityServiceConfig;
        /// <summary>
        /// The initial Kubernetes version for this cluster. Valid versions are those found in validMasterVersions returned by getServerConfig. The version can be upgraded over time; such upgrades are reflected in currentMasterVersion and currentNodeVersion. Users may specify either explicit versions offered by Kubernetes Engine or version aliases, which have the following behavior: - "latest": picks the highest valid Kubernetes version - "1.X": picks the highest valid patch+gke.N patch in the 1.X version - "1.X.Y": picks the highest valid gke.N patch in the 1.X.Y version - "1.X.Y-gke.N": picks an explicit Kubernetes version - "","-": picks the default Kubernetes version
        /// </summary>
        public readonly string InitialClusterVersion;
        /// <summary>
        /// The number of nodes to create in this cluster. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota. For requests, this field should only be used in lieu of a "node_pool" object, since this configuration (along with the "node_config") will be used to create a "NodePool" object with an auto-generated name. Do not use this and a node_pool at the same time. This field is deprecated, use node_pool.initial_node_count instead.
        /// </summary>
        public readonly int InitialNodeCount;
        /// <summary>
        /// Deprecated. Use node_pools.instance_group_urls.
        /// </summary>
        public readonly ImmutableArray<string> InstanceGroupUrls;
        /// <summary>
        /// Configuration for cluster IP allocation.
        /// </summary>
        public readonly Outputs.IPAllocationPolicyResponse IpAllocationPolicy;
        /// <summary>
        /// The fingerprint of the set of labels for this cluster.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Configuration for the legacy ABAC authorization mode.
        /// </summary>
        public readonly Outputs.LegacyAbacResponse LegacyAbac;
        /// <summary>
        /// [Output only] The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/regions-zones/regions-zones#available) or [region](https://cloud.google.com/compute/docs/regions-zones/regions-zones#available) in which the cluster resides.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the cluster's nodes should be located. This field provides a default value if [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) are not specified during node pool creation. Warning: changing cluster locations will update the [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) of all node pools and will result in nodes being added and/or removed.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// Logging configuration for the cluster.
        /// </summary>
        public readonly Outputs.LoggingConfigResponse LoggingConfig;
        /// <summary>
        /// The logging service the cluster should use to write logs. Currently available options: * `logging.googleapis.com/kubernetes` - The Cloud Logging service with a Kubernetes-native resource model * `logging.googleapis.com` - The legacy Cloud Logging service (no longer available as of GKE 1.15). * `none` - no logs will be exported from the cluster. If left as an empty string,`logging.googleapis.com/kubernetes` will be used for GKE 1.14+ or `logging.googleapis.com` for earlier versions.
        /// </summary>
        public readonly string LoggingService;
        /// <summary>
        /// Configure the maintenance policy for this cluster.
        /// </summary>
        public readonly Outputs.MaintenancePolicyResponse MaintenancePolicy;
        /// <summary>
        /// Configuration for master components.
        /// </summary>
        public readonly Outputs.MasterResponse Master;
        /// <summary>
        /// The authentication information for accessing the master endpoint. If unspecified, the defaults are used: For clusters before v1.12, if master_auth is unspecified, `username` will be set to "admin", a random password will be generated, and a client certificate will be issued.
        /// </summary>
        public readonly Outputs.MasterAuthResponse MasterAuth;
        /// <summary>
        /// The configuration options for master authorized networks feature.
        /// </summary>
        public readonly Outputs.MasterAuthorizedNetworksConfigResponse MasterAuthorizedNetworksConfig;
        /// <summary>
        /// The IP prefix in CIDR notation to use for the hosted master network. This prefix will be used for assigning private IP addresses to the master or set of masters, as well as the ILB VIP. This field is deprecated, use private_cluster_config.master_ipv4_cidr_block instead.
        /// </summary>
        public readonly string MasterIpv4CidrBlock;
        /// <summary>
        /// Configuration for issuance of mTLS keys and certificates to Kubernetes pods.
        /// </summary>
        public readonly Outputs.MeshCertificatesResponse MeshCertificates;
        /// <summary>
        /// Monitoring configuration for the cluster.
        /// </summary>
        public readonly Outputs.MonitoringConfigResponse MonitoringConfig;
        /// <summary>
        /// The monitoring service the cluster should use to write metrics. Currently available options: * "monitoring.googleapis.com/kubernetes" - The Cloud Monitoring service with a Kubernetes-native resource model * `monitoring.googleapis.com` - The legacy Cloud Monitoring service (no longer available as of GKE 1.15). * `none` - No metrics will be exported from the cluster. If left as an empty string,`monitoring.googleapis.com/kubernetes` will be used for GKE 1.14+ or `monitoring.googleapis.com` for earlier versions.
        /// </summary>
        public readonly string MonitoringService;
        /// <summary>
        /// The name of this cluster. The name must be unique within this project and location (e.g. zone or region), and can be up to 40 characters with the following restrictions: * Lowercase letters, numbers, and hyphens only. * Must start with a letter. * Must end with a number or a letter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks) to which the cluster is connected. If left unspecified, the `default` network will be used. On output this shows the network ID instead of the name.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Configuration for cluster networking.
        /// </summary>
        public readonly Outputs.NetworkConfigResponse NetworkConfig;
        /// <summary>
        /// Configuration options for the NetworkPolicy feature.
        /// </summary>
        public readonly Outputs.NetworkPolicyResponse NetworkPolicy;
        /// <summary>
        /// Parameters used in creating the cluster's nodes. For requests, this field should only be used in lieu of a "node_pool" object, since this configuration (along with the "initial_node_count") will be used to create a "NodePool" object with an auto-generated name. Do not use this and a node_pool at the same time. For responses, this field will be populated with the node configuration of the first node pool. (For configuration of each node pool, see `node_pool.config`) If unspecified, the defaults are used. This field is deprecated, use node_pool.config instead.
        /// </summary>
        public readonly Outputs.NodeConfigResponse NodeConfig;
        /// <summary>
        /// [Output only] The size of the address space on each node for hosting containers. This is provisioned from within the `container_ipv4_cidr` range. This field will only be set when cluster is in route-based network mode.
        /// </summary>
        public readonly int NodeIpv4CidrSize;
        /// <summary>
        /// Node pool configs that apply to all auto-provisioned node pools in autopilot clusters and node auto-provisioning enabled clusters.
        /// </summary>
        public readonly Outputs.NodePoolAutoConfigResponse NodePoolAutoConfig;
        /// <summary>
        /// Default NodePool settings for the entire cluster. These settings are overridden if specified on the specific NodePool object.
        /// </summary>
        public readonly Outputs.NodePoolDefaultsResponse NodePoolDefaults;
        /// <summary>
        /// The node pools associated with this cluster. This field should not be set if "node_config" or "initial_node_count" are specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodePoolResponse> NodePools;
        /// <summary>
        /// Notification configuration of the cluster.
        /// </summary>
        public readonly Outputs.NotificationConfigResponse NotificationConfig;
        /// <summary>
        /// Configuration for the PodSecurityPolicy feature.
        /// </summary>
        public readonly Outputs.PodSecurityPolicyConfigResponse PodSecurityPolicyConfig;
        /// <summary>
        /// If this is a private cluster setup. Private clusters are clusters that, by default have no external IP addresses on the nodes and where nodes and the master communicate over private IP addresses. This field is deprecated, use private_cluster_config.enable_private_nodes instead.
        /// </summary>
        public readonly bool PrivateCluster;
        /// <summary>
        /// Configuration for private cluster.
        /// </summary>
        public readonly Outputs.PrivateClusterConfigResponse PrivateClusterConfig;
        /// <summary>
        /// Release channel configuration.
        /// </summary>
        public readonly Outputs.ReleaseChannelResponse ReleaseChannel;
        /// <summary>
        /// The resource labels for the cluster to use to annotate any related Google Compute Engine resources.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ResourceLabels;
        /// <summary>
        /// Configuration for exporting resource usages. Resource usage export is disabled when this config unspecified.
        /// </summary>
        public readonly Outputs.ResourceUsageExportConfigResponse ResourceUsageExportConfig;
        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Output only] The IP address range of the Kubernetes services in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `1.2.3.4/29`). Service addresses are typically put in the last `/16` from the container CIDR.
        /// </summary>
        public readonly string ServicesIpv4Cidr;
        /// <summary>
        /// Shielded Nodes configuration.
        /// </summary>
        public readonly Outputs.ShieldedNodesResponse ShieldedNodes;
        /// <summary>
        /// [Output only] The current status of this cluster.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// [Output only] Deprecated. Use conditions instead. Additional information about the current status of this cluster, if available.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// The name of the Google Compute Engine [subnetwork](https://cloud.google.com/compute/docs/subnetworks) to which the cluster is connected. On output this shows the subnetwork ID instead of the name.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// Configuration for Cloud TPU support;
        /// </summary>
        public readonly Outputs.TpuConfigResponse TpuConfig;
        /// <summary>
        /// [Output only] The IP address range of the Cloud TPUs in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `1.2.3.4/29`).
        /// </summary>
        public readonly string TpuIpv4CidrBlock;
        /// <summary>
        /// Cluster-level Vertical Pod Autoscaling configuration.
        /// </summary>
        public readonly Outputs.VerticalPodAutoscalingResponse VerticalPodAutoscaling;
        /// <summary>
        /// Configuration for direct-path (via ALTS) with workload identity.
        /// </summary>
        public readonly Outputs.WorkloadALTSConfigResponse WorkloadAltsConfig;
        /// <summary>
        /// Configuration for issuance of mTLS keys and certificates to Kubernetes pods.
        /// </summary>
        public readonly Outputs.WorkloadCertificatesResponse WorkloadCertificates;
        /// <summary>
        /// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
        /// </summary>
        public readonly Outputs.WorkloadIdentityConfigResponse WorkloadIdentityConfig;
        /// <summary>
        /// [Output only] The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/zones#available) in which the cluster resides. This field is deprecated, use location instead.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetClusterResult(
            Outputs.AddonsConfigResponse addonsConfig,

            Outputs.AuthenticatorGroupsConfigResponse authenticatorGroupsConfig,

            Outputs.AutopilotResponse autopilot,

            Outputs.ClusterAutoscalingResponse autoscaling,

            Outputs.BinaryAuthorizationResponse binaryAuthorization,

            string clusterIpv4Cidr,

            Outputs.ClusterTelemetryResponse clusterTelemetry,

            ImmutableArray<Outputs.StatusConditionResponse> conditions,

            Outputs.ConfidentialNodesResponse confidentialNodes,

            string createTime,

            string currentMasterVersion,

            int currentNodeCount,

            string currentNodeVersion,

            Outputs.DatabaseEncryptionResponse databaseEncryption,

            Outputs.MaxPodsConstraintResponse defaultMaxPodsConstraint,

            string description,

            bool enableKubernetesAlpha,

            bool enableTpu,

            string endpoint,

            string expireTime,

            Outputs.IdentityServiceConfigResponse identityServiceConfig,

            string initialClusterVersion,

            int initialNodeCount,

            ImmutableArray<string> instanceGroupUrls,

            Outputs.IPAllocationPolicyResponse ipAllocationPolicy,

            string labelFingerprint,

            Outputs.LegacyAbacResponse legacyAbac,

            string location,

            ImmutableArray<string> locations,

            Outputs.LoggingConfigResponse loggingConfig,

            string loggingService,

            Outputs.MaintenancePolicyResponse maintenancePolicy,

            Outputs.MasterResponse master,

            Outputs.MasterAuthResponse masterAuth,

            Outputs.MasterAuthorizedNetworksConfigResponse masterAuthorizedNetworksConfig,

            string masterIpv4CidrBlock,

            Outputs.MeshCertificatesResponse meshCertificates,

            Outputs.MonitoringConfigResponse monitoringConfig,

            string monitoringService,

            string name,

            string network,

            Outputs.NetworkConfigResponse networkConfig,

            Outputs.NetworkPolicyResponse networkPolicy,

            Outputs.NodeConfigResponse nodeConfig,

            int nodeIpv4CidrSize,

            Outputs.NodePoolAutoConfigResponse nodePoolAutoConfig,

            Outputs.NodePoolDefaultsResponse nodePoolDefaults,

            ImmutableArray<Outputs.NodePoolResponse> nodePools,

            Outputs.NotificationConfigResponse notificationConfig,

            Outputs.PodSecurityPolicyConfigResponse podSecurityPolicyConfig,

            bool privateCluster,

            Outputs.PrivateClusterConfigResponse privateClusterConfig,

            Outputs.ReleaseChannelResponse releaseChannel,

            ImmutableDictionary<string, string> resourceLabels,

            Outputs.ResourceUsageExportConfigResponse resourceUsageExportConfig,

            string selfLink,

            string servicesIpv4Cidr,

            Outputs.ShieldedNodesResponse shieldedNodes,

            string status,

            string statusMessage,

            string subnetwork,

            Outputs.TpuConfigResponse tpuConfig,

            string tpuIpv4CidrBlock,

            Outputs.VerticalPodAutoscalingResponse verticalPodAutoscaling,

            Outputs.WorkloadALTSConfigResponse workloadAltsConfig,

            Outputs.WorkloadCertificatesResponse workloadCertificates,

            Outputs.WorkloadIdentityConfigResponse workloadIdentityConfig,

            string zone)
        {
            AddonsConfig = addonsConfig;
            AuthenticatorGroupsConfig = authenticatorGroupsConfig;
            Autopilot = autopilot;
            Autoscaling = autoscaling;
            BinaryAuthorization = binaryAuthorization;
            ClusterIpv4Cidr = clusterIpv4Cidr;
            ClusterTelemetry = clusterTelemetry;
            Conditions = conditions;
            ConfidentialNodes = confidentialNodes;
            CreateTime = createTime;
            CurrentMasterVersion = currentMasterVersion;
            CurrentNodeCount = currentNodeCount;
            CurrentNodeVersion = currentNodeVersion;
            DatabaseEncryption = databaseEncryption;
            DefaultMaxPodsConstraint = defaultMaxPodsConstraint;
            Description = description;
            EnableKubernetesAlpha = enableKubernetesAlpha;
            EnableTpu = enableTpu;
            Endpoint = endpoint;
            ExpireTime = expireTime;
            IdentityServiceConfig = identityServiceConfig;
            InitialClusterVersion = initialClusterVersion;
            InitialNodeCount = initialNodeCount;
            InstanceGroupUrls = instanceGroupUrls;
            IpAllocationPolicy = ipAllocationPolicy;
            LabelFingerprint = labelFingerprint;
            LegacyAbac = legacyAbac;
            Location = location;
            Locations = locations;
            LoggingConfig = loggingConfig;
            LoggingService = loggingService;
            MaintenancePolicy = maintenancePolicy;
            Master = master;
            MasterAuth = masterAuth;
            MasterAuthorizedNetworksConfig = masterAuthorizedNetworksConfig;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
            MeshCertificates = meshCertificates;
            MonitoringConfig = monitoringConfig;
            MonitoringService = monitoringService;
            Name = name;
            Network = network;
            NetworkConfig = networkConfig;
            NetworkPolicy = networkPolicy;
            NodeConfig = nodeConfig;
            NodeIpv4CidrSize = nodeIpv4CidrSize;
            NodePoolAutoConfig = nodePoolAutoConfig;
            NodePoolDefaults = nodePoolDefaults;
            NodePools = nodePools;
            NotificationConfig = notificationConfig;
            PodSecurityPolicyConfig = podSecurityPolicyConfig;
            PrivateCluster = privateCluster;
            PrivateClusterConfig = privateClusterConfig;
            ReleaseChannel = releaseChannel;
            ResourceLabels = resourceLabels;
            ResourceUsageExportConfig = resourceUsageExportConfig;
            SelfLink = selfLink;
            ServicesIpv4Cidr = servicesIpv4Cidr;
            ShieldedNodes = shieldedNodes;
            Status = status;
            StatusMessage = statusMessage;
            Subnetwork = subnetwork;
            TpuConfig = tpuConfig;
            TpuIpv4CidrBlock = tpuIpv4CidrBlock;
            VerticalPodAutoscaling = verticalPodAutoscaling;
            WorkloadAltsConfig = workloadAltsConfig;
            WorkloadCertificates = workloadCertificates;
            WorkloadIdentityConfig = workloadIdentityConfig;
            Zone = zone;
        }
    }
}
