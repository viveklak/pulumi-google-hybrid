// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class NetworkPeeringResponse
    {
        /// <summary>
        /// This field will be deprecated soon. Use the exchange_subnet_routes field instead. Indicates whether full mesh connectivity is created and managed automatically between peered networks. Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE.
        /// </summary>
        public readonly bool AutoCreateRoutes;
        /// <summary>
        /// Indicates whether full mesh connectivity is created and managed automatically between peered networks. Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE.
        /// </summary>
        public readonly bool ExchangeSubnetRoutes;
        /// <summary>
        /// Whether to export the custom routes to peer network.
        /// </summary>
        public readonly bool ExportCustomRoutes;
        /// <summary>
        /// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
        /// </summary>
        public readonly bool ExportSubnetRoutesWithPublicIp;
        /// <summary>
        /// Whether to import the custom routes from peer network.
        /// </summary>
        public readonly bool ImportCustomRoutes;
        /// <summary>
        /// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
        /// </summary>
        public readonly bool ImportSubnetRoutesWithPublicIp;
        /// <summary>
        /// Name of this peering. Provided by the client when the peering is created. The name must comply with RFC1035. Specifically, the name must be 1-63 characters long and match regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all the following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the peer network. It can be either full URL or partial URL. The peer network may belong to a different project. If the partial URL does not contain project, it is assumed that the peer network is in the same project as the current network.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Maximum Transmission Unit in bytes.
        /// </summary>
        public readonly int PeerMtu;
        /// <summary>
        /// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is `ACTIVE` when there's a matching configuration in the peer network.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Details about the current state of the peering.
        /// </summary>
        public readonly string StateDetails;

        [OutputConstructor]
        private NetworkPeeringResponse(
            bool autoCreateRoutes,

            bool exchangeSubnetRoutes,

            bool exportCustomRoutes,

            bool exportSubnetRoutesWithPublicIp,

            bool importCustomRoutes,

            bool importSubnetRoutesWithPublicIp,

            string name,

            string network,

            int peerMtu,

            string state,

            string stateDetails)
        {
            AutoCreateRoutes = autoCreateRoutes;
            ExchangeSubnetRoutes = exchangeSubnetRoutes;
            ExportCustomRoutes = exportCustomRoutes;
            ExportSubnetRoutesWithPublicIp = exportSubnetRoutesWithPublicIp;
            ImportCustomRoutes = importCustomRoutes;
            ImportSubnetRoutesWithPublicIp = importSubnetRoutesWithPublicIp;
            Name = name;
            Network = network;
            PeerMtu = peerMtu;
            State = state;
            StateDetails = stateDetails;
        }
    }
}
