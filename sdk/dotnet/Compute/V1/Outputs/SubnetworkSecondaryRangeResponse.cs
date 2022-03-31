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
    /// Represents a secondary IP range of a subnetwork.
    /// </summary>
    [OutputType]
    public sealed class SubnetworkSecondaryRangeResponse
    {
        /// <summary>
        /// The range of IP addresses belonging to this subnetwork secondary range. Provide this property when you create the subnetwork. Ranges must be unique and non-overlapping with all primary and secondary IP ranges within a network. Only IPv4 is supported. The range can be any range listed in the Valid ranges list.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance. The name must be 1-63 characters long, and comply with RFC1035. The name must be unique within the subnetwork.
        /// </summary>
        public readonly string RangeName;

        [OutputConstructor]
        private SubnetworkSecondaryRangeResponse(
            string ipCidrRange,

            string rangeName)
        {
            IpCidrRange = ipCidrRange;
            RangeName = rangeName;
        }
    }
}
