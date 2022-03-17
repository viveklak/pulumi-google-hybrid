// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// An alias IP range attached to an instance's network interface.
    /// </summary>
    [OutputType]
    public sealed class AliasIpRangeResponse
    {
        /// <summary>
        /// The IP alias ranges to allocate for this interface. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. This range may be a single IP address (such as 10.2.3.4), a netmask (such as /24) or a CIDR-formatted string (such as 10.1.2.0/24).
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// The name of a subnetwork secondary IP range from which to allocate an IP alias range. If not specified, the primary range of the subnetwork is used.
        /// </summary>
        public readonly string SubnetworkRangeName;

        [OutputConstructor]
        private AliasIpRangeResponse(
            string ipCidrRange,

            string subnetworkRangeName)
        {
            IpCidrRange = ipCidrRange;
            SubnetworkRangeName = subnetworkRangeName;
        }
    }
}
