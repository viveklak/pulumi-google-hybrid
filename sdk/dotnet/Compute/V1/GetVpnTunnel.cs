// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetVpnTunnel
    {
        /// <summary>
        /// Returns the specified VpnTunnel resource. Gets a list of available VPN tunnels by making a list() request.
        /// </summary>
        public static Task<GetVpnTunnelResult> InvokeAsync(GetVpnTunnelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpnTunnelResult>("google-native:compute/v1:getVpnTunnel", args ?? new GetVpnTunnelArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified VpnTunnel resource. Gets a list of available VPN tunnels by making a list() request.
        /// </summary>
        public static Output<GetVpnTunnelResult> Invoke(GetVpnTunnelInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpnTunnelResult>("google-native:compute/v1:getVpnTunnel", args ?? new GetVpnTunnelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpnTunnelArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("vpnTunnel", required: true)]
        public string VpnTunnel { get; set; } = null!;

        public GetVpnTunnelArgs()
        {
        }
    }

    public sealed class GetVpnTunnelInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("vpnTunnel", required: true)]
        public Input<string> VpnTunnel { get; set; } = null!;

        public GetVpnTunnelInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpnTunnelResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Detailed status message for the VPN tunnel.
        /// </summary>
        public readonly string DetailedStatus;
        /// <summary>
        /// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        /// </summary>
        public readonly int IkeVersion;
        /// <summary>
        /// Type of resource. Always compute#vpnTunnel for VPN tunnels.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        /// </summary>
        public readonly ImmutableArray<string> LocalTrafficSelector;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        /// </summary>
        public readonly string PeerExternalGateway;
        /// <summary>
        /// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        /// </summary>
        public readonly int PeerExternalGatewayInterface;
        /// <summary>
        /// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        /// </summary>
        public readonly string PeerGcpGateway;
        /// <summary>
        /// IP address of the peer VPN gateway. Only IPv4 is supported.
        /// </summary>
        public readonly string PeerIp;
        /// <summary>
        /// URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        /// </summary>
        public readonly ImmutableArray<string> RemoteTrafficSelector;
        /// <summary>
        /// URL of the router resource to be used for dynamic routing.
        /// </summary>
        public readonly string Router;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        /// </summary>
        public readonly string SharedSecret;
        /// <summary>
        /// Hash of the shared secret.
        /// </summary>
        public readonly string SharedSecretHash;
        /// <summary>
        /// The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel. 
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        /// </summary>
        public readonly string TargetVpnGateway;
        /// <summary>
        /// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        /// </summary>
        public readonly string VpnGateway;
        /// <summary>
        /// The interface ID of the VPN gateway with which this VPN tunnel is associated.
        /// </summary>
        public readonly int VpnGatewayInterface;

        [OutputConstructor]
        private GetVpnTunnelResult(
            string creationTimestamp,

            string description,

            string detailedStatus,

            int ikeVersion,

            string kind,

            ImmutableArray<string> localTrafficSelector,

            string name,

            string peerExternalGateway,

            int peerExternalGatewayInterface,

            string peerGcpGateway,

            string peerIp,

            string region,

            ImmutableArray<string> remoteTrafficSelector,

            string router,

            string selfLink,

            string sharedSecret,

            string sharedSecretHash,

            string status,

            string targetVpnGateway,

            string vpnGateway,

            int vpnGatewayInterface)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            DetailedStatus = detailedStatus;
            IkeVersion = ikeVersion;
            Kind = kind;
            LocalTrafficSelector = localTrafficSelector;
            Name = name;
            PeerExternalGateway = peerExternalGateway;
            PeerExternalGatewayInterface = peerExternalGatewayInterface;
            PeerGcpGateway = peerGcpGateway;
            PeerIp = peerIp;
            Region = region;
            RemoteTrafficSelector = remoteTrafficSelector;
            Router = router;
            SelfLink = selfLink;
            SharedSecret = sharedSecret;
            SharedSecretHash = sharedSecretHash;
            Status = status;
            TargetVpnGateway = targetVpnGateway;
            VpnGateway = vpnGateway;
            VpnGatewayInterface = vpnGatewayInterface;
        }
    }
}
