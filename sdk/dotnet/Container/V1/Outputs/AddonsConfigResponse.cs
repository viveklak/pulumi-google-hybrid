// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Configuration for the addons that can be automatically spun up in the cluster, enabling additional functionality.
    /// </summary>
    [OutputType]
    public sealed class AddonsConfigResponse
    {
        /// <summary>
        /// Configuration for the Cloud Run addon, which allows the user to use a managed Knative service.
        /// </summary>
        public readonly Outputs.CloudRunConfigResponse CloudRunConfig;
        /// <summary>
        /// Configuration for the ConfigConnector add-on, a Kubernetes extension to manage hosted GCP services through the Kubernetes API
        /// </summary>
        public readonly Outputs.ConfigConnectorConfigResponse ConfigConnectorConfig;
        /// <summary>
        /// Configuration for NodeLocalDNS, a dns cache running on cluster nodes
        /// </summary>
        public readonly Outputs.DnsCacheConfigResponse DnsCacheConfig;
        /// <summary>
        /// Configuration for the Compute Engine Persistent Disk CSI driver.
        /// </summary>
        public readonly Outputs.GcePersistentDiskCsiDriverConfigResponse GcePersistentDiskCsiDriverConfig;
        /// <summary>
        /// Configuration for the GCP Filestore CSI driver.
        /// </summary>
        public readonly Outputs.GcpFilestoreCsiDriverConfigResponse GcpFilestoreCsiDriverConfig;
        /// <summary>
        /// Configuration for the horizontal pod autoscaling feature, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods.
        /// </summary>
        public readonly Outputs.HorizontalPodAutoscalingResponse HorizontalPodAutoscaling;
        /// <summary>
        /// Configuration for the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster.
        /// </summary>
        public readonly Outputs.HttpLoadBalancingResponse HttpLoadBalancing;
        /// <summary>
        /// Configuration for the Kubernetes Dashboard. This addon is deprecated, and will be disabled in 1.15. It is recommended to use the Cloud Console to manage and monitor your Kubernetes clusters, workloads and applications. For more information, see: https://cloud.google.com/kubernetes-engine/docs/concepts/dashboards
        /// </summary>
        public readonly Outputs.KubernetesDashboardResponse KubernetesDashboard;
        /// <summary>
        /// Configuration for NetworkPolicy. This only tracks whether the addon is enabled or not on the Master, it does not track whether network policy is enabled for the nodes.
        /// </summary>
        public readonly Outputs.NetworkPolicyConfigResponse NetworkPolicyConfig;

        [OutputConstructor]
        private AddonsConfigResponse(
            Outputs.CloudRunConfigResponse cloudRunConfig,

            Outputs.ConfigConnectorConfigResponse configConnectorConfig,

            Outputs.DnsCacheConfigResponse dnsCacheConfig,

            Outputs.GcePersistentDiskCsiDriverConfigResponse gcePersistentDiskCsiDriverConfig,

            Outputs.GcpFilestoreCsiDriverConfigResponse gcpFilestoreCsiDriverConfig,

            Outputs.HorizontalPodAutoscalingResponse horizontalPodAutoscaling,

            Outputs.HttpLoadBalancingResponse httpLoadBalancing,

            Outputs.KubernetesDashboardResponse kubernetesDashboard,

            Outputs.NetworkPolicyConfigResponse networkPolicyConfig)
        {
            CloudRunConfig = cloudRunConfig;
            ConfigConnectorConfig = configConnectorConfig;
            DnsCacheConfig = dnsCacheConfig;
            GcePersistentDiskCsiDriverConfig = gcePersistentDiskCsiDriverConfig;
            GcpFilestoreCsiDriverConfig = gcpFilestoreCsiDriverConfig;
            HorizontalPodAutoscaling = horizontalPodAutoscaling;
            HttpLoadBalancing = httpLoadBalancing;
            KubernetesDashboard = kubernetesDashboard;
            NetworkPolicyConfig = networkPolicyConfig;
        }
    }
}
