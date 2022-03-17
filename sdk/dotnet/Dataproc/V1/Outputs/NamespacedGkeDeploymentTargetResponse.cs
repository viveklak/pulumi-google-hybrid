// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// A full, namespace-isolated deployment target for an existing GKE cluster.
    /// </summary>
    [OutputType]
    public sealed class NamespacedGkeDeploymentTargetResponse
    {
        /// <summary>
        /// Optional. A namespace within the GKE cluster to deploy into.
        /// </summary>
        public readonly string ClusterNamespace;
        /// <summary>
        /// Optional. The target GKE cluster to deploy to. Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'
        /// </summary>
        public readonly string TargetGkeCluster;

        [OutputConstructor]
        private NamespacedGkeDeploymentTargetResponse(
            string clusterNamespace,

            string targetGkeCluster)
        {
            ClusterNamespace = clusterNamespace;
            TargetGkeCluster = targetGkeCluster;
        }
    }
}
