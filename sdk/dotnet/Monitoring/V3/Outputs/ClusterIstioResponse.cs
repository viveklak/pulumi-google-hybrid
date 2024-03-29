// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// Istio service scoped to a single Kubernetes cluster. Learn more at https://istio.io. Clusters running OSS Istio will have their services ingested as this type.
    /// </summary>
    [OutputType]
    public sealed class ClusterIstioResponse
    {
        /// <summary>
        /// The name of the Kubernetes cluster in which this Istio service is defined. Corresponds to the cluster_name resource label in k8s_cluster resources.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The location of the Kubernetes cluster in which this Istio service is defined. Corresponds to the location resource label in k8s_cluster resources.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Istio service underlying this service. Corresponds to the destination_service_name metric label in Istio metrics.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// The namespace of the Istio service underlying this service. Corresponds to the destination_service_namespace metric label in Istio metrics.
        /// </summary>
        public readonly string ServiceNamespace;

        [OutputConstructor]
        private ClusterIstioResponse(
            string clusterName,

            string location,

            string serviceName,

            string serviceNamespace)
        {
            ClusterName = clusterName;
            Location = location;
            ServiceName = serviceName;
            ServiceNamespace = serviceNamespace;
        }
    }
}
