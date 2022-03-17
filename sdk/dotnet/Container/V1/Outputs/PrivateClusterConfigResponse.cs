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
    /// Configuration options for private clusters.
    /// </summary>
    [OutputType]
    public sealed class PrivateClusterConfigResponse
    {
        /// <summary>
        /// Whether the master's internal IP address is used as the cluster endpoint.
        /// </summary>
        public readonly bool EnablePrivateEndpoint;
        /// <summary>
        /// Whether nodes have internal IP addresses only. If enabled, all nodes are given only RFC 1918 private addresses and communicate with the master via private networking.
        /// </summary>
        public readonly bool EnablePrivateNodes;
        /// <summary>
        /// Controls master global access settings.
        /// </summary>
        public readonly Outputs.PrivateClusterMasterGlobalAccessConfigResponse MasterGlobalAccessConfig;
        /// <summary>
        /// The IP range in CIDR notation to use for the hosted master network. This range will be used for assigning internal IP addresses to the master or set of masters, as well as the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network.
        /// </summary>
        public readonly string MasterIpv4CidrBlock;
        /// <summary>
        /// The peering name in the customer VPC used by this cluster.
        /// </summary>
        public readonly string PeeringName;
        /// <summary>
        /// The internal IP address of this cluster's master endpoint.
        /// </summary>
        public readonly string PrivateEndpoint;
        /// <summary>
        /// The external IP address of this cluster's master endpoint.
        /// </summary>
        public readonly string PublicEndpoint;

        [OutputConstructor]
        private PrivateClusterConfigResponse(
            bool enablePrivateEndpoint,

            bool enablePrivateNodes,

            Outputs.PrivateClusterMasterGlobalAccessConfigResponse masterGlobalAccessConfig,

            string masterIpv4CidrBlock,

            string peeringName,

            string privateEndpoint,

            string publicEndpoint)
        {
            EnablePrivateEndpoint = enablePrivateEndpoint;
            EnablePrivateNodes = enablePrivateNodes;
            MasterGlobalAccessConfig = masterGlobalAccessConfig;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
            PeeringName = peeringName;
            PrivateEndpoint = privateEndpoint;
            PublicEndpoint = publicEndpoint;
        }
    }
}
