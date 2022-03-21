// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1.Outputs
{

    /// <summary>
    /// For display only. Metadata associated with a Compute Engine VPN tunnel.
    /// </summary>
    [OutputType]
    public sealed class VpnTunnelInfoResponse
    {
        /// <summary>
        /// Name of a VPN tunnel.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// URI of a Compute Engine network where the VPN tunnel is configured.
        /// </summary>
        public readonly string NetworkUri;
        /// <summary>
        /// Name of a Google Cloud region where this VPN tunnel is configured.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// URI of a VPN gateway at remote end of the tunnel.
        /// </summary>
        public readonly string RemoteGateway;
        /// <summary>
        /// Remote VPN gateway's IP address.
        /// </summary>
        public readonly string RemoteGatewayIp;
        /// <summary>
        /// Type of the routing policy.
        /// </summary>
        public readonly string RoutingType;
        /// <summary>
        /// URI of the VPN gateway at local end of the tunnel.
        /// </summary>
        public readonly string SourceGateway;
        /// <summary>
        /// Local VPN gateway's IP address.
        /// </summary>
        public readonly string SourceGatewayIp;
        /// <summary>
        /// URI of a VPN tunnel.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private VpnTunnelInfoResponse(
            string displayName,

            string networkUri,

            string region,

            string remoteGateway,

            string remoteGatewayIp,

            string routingType,

            string sourceGateway,

            string sourceGatewayIp,

            string uri)
        {
            DisplayName = displayName;
            NetworkUri = networkUri;
            Region = region;
            RemoteGateway = remoteGateway;
            RemoteGatewayIp = remoteGatewayIp;
            RoutingType = routingType;
            SourceGateway = sourceGateway;
            SourceGatewayIp = sourceGatewayIp;
            Uri = uri;
        }
    }
}
