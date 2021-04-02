// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Composer.V1beta1.Outputs
{

    [OutputType]
    public sealed class PrivateClusterConfigResponse
    {
        /// <summary>
        /// Optional. If `true`, access to the public endpoint of the GKE cluster is denied.
        /// </summary>
        public readonly bool EnablePrivateEndpoint;
        /// <summary>
        /// Optional. The CIDR block from which IPv4 range for GKE master will be reserved. If left blank, the default value of '172.16.0.0/23' is used.
        /// </summary>
        public readonly string MasterIpv4CidrBlock;
        /// <summary>
        /// The IP range in CIDR notation to use for the hosted master network. This range is used for assigning internal IP addresses to the cluster master or set of masters and to the internal load balancer virtual IP. This range must not overlap with any other ranges in use within the cluster's network.
        /// </summary>
        public readonly string MasterIpv4ReservedRange;

        [OutputConstructor]
        private PrivateClusterConfigResponse(
            bool enablePrivateEndpoint,

            string masterIpv4CidrBlock,

            string masterIpv4ReservedRange)
        {
            EnablePrivateEndpoint = enablePrivateEndpoint;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
            MasterIpv4ReservedRange = masterIpv4ReservedRange;
        }
    }
}