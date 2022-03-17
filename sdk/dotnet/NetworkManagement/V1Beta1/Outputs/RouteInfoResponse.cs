// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1Beta1.Outputs
{

    /// <summary>
    /// For display only. Metadata associated with a Compute Engine route.
    /// </summary>
    [OutputType]
    public sealed class RouteInfoResponse
    {
        /// <summary>
        /// Destination IP range of the route.
        /// </summary>
        public readonly string DestIpRange;
        /// <summary>
        /// Name of a Compute Engine route.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Instance tags of the route.
        /// </summary>
        public readonly ImmutableArray<string> InstanceTags;
        /// <summary>
        /// URI of a Compute Engine network.
        /// </summary>
        public readonly string NetworkUri;
        /// <summary>
        /// Next hop of the route.
        /// </summary>
        public readonly string NextHop;
        /// <summary>
        /// Type of next hop.
        /// </summary>
        public readonly string NextHopType;
        /// <summary>
        /// Priority of the route.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Type of route.
        /// </summary>
        public readonly string RouteType;
        /// <summary>
        /// URI of a Compute Engine route. Dynamic route from cloud router does not have a URI. Advertised route from Google Cloud VPC to on-premises network also does not have a URI.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private RouteInfoResponse(
            string destIpRange,

            string displayName,

            ImmutableArray<string> instanceTags,

            string networkUri,

            string nextHop,

            string nextHopType,

            int priority,

            string routeType,

            string uri)
        {
            DestIpRange = destIpRange;
            DisplayName = displayName;
            InstanceTags = instanceTags;
            NetworkUri = networkUri;
            NextHop = nextHop;
            NextHopType = nextHopType;
            Priority = priority;
            RouteType = routeType;
            Uri = uri;
        }
    }
}
