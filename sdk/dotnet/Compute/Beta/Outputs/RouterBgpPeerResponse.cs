// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class RouterBgpPeerResponse
    {
        /// <summary>
        /// User-specified flag to indicate which mode to use for advertisement.
        /// </summary>
        public readonly string AdvertiseMode;
        /// <summary>
        /// User-specified list of prefix groups to advertise in custom mode, which can take one of the following options: 
        /// - ALL_SUBNETS: Advertises all available subnets, including peer VPC subnets. 
        /// - ALL_VPC_SUBNETS: Advertises the router's own VPC subnets. Note that this field can only be populated if advertise_mode is CUSTOM and overrides the list defined for the router (in the "bgp" message). These groups are advertised in addition to any specified prefixes. Leave this field blank to advertise no custom groups.
        /// </summary>
        public readonly ImmutableArray<string> AdvertisedGroups;
        /// <summary>
        /// User-specified list of individual IP ranges to advertise in custom mode. This field can only be populated if advertise_mode is CUSTOM and overrides the list defined for the router (in the "bgp" message). These IP ranges are advertised in addition to any specified groups. Leave this field blank to advertise no custom IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterAdvertisedIpRangeResponse> AdvertisedIpRanges;
        /// <summary>
        /// The priority of routes advertised to this BGP peer. Where there is more than one matching route of maximum length, the routes with the lowest priority value win.
        /// </summary>
        public readonly int AdvertisedRoutePriority;
        /// <summary>
        /// BFD configuration for the BGP peering.
        /// Not currently available publicly.
        /// </summary>
        public readonly Outputs.RouterBgpPeerBfdResponse Bfd;
        /// <summary>
        /// The status of the BGP peer connection.
        /// Not currently available publicly.
        /// If set to FALSE, any active session with the peer is terminated and all associated routing information is removed. If set to TRUE, the peer connection can be established with routing information. The default is TRUE.
        /// </summary>
        public readonly string Enable;
        /// <summary>
        /// Name of the interface the BGP peer is associated with.
        /// </summary>
        public readonly string InterfaceName;
        /// <summary>
        /// IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// [Output Only] The resource that configures and manages this BGP peer. 
        /// - MANAGED_BY_USER is the default value and can be managed by you or other users 
        /// - MANAGED_BY_ATTACHMENT is a BGP peer that is configured and managed by Cloud Interconnect, specifically by an InterconnectAttachment of type PARTNER. Google automatically creates, updates, and deletes this type of BGP peer when the PARTNER InterconnectAttachment is created, updated, or deleted.
        /// </summary>
        public readonly string ManagementType;
        /// <summary>
        /// Name of this BGP peer. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Peer BGP Autonomous System Number (ASN). Each BGP interface may use a different value.
        /// </summary>
        public readonly int PeerAsn;
        /// <summary>
        /// IP address of the BGP interface outside Google Cloud Platform. Only IPv4 is supported.
        /// </summary>
        public readonly string PeerIpAddress;
        /// <summary>
        /// URI of the VM instance that is used as third-party router appliances such as Next Gen Firewalls, Virtual Routers, or Router Appliances. The VM instance must be located in zones contained in the same region as this Cloud Router. The VM instance is the peer side of the BGP session.
        /// </summary>
        public readonly string RouterApplianceInstance;

        [OutputConstructor]
        private RouterBgpPeerResponse(
            string advertiseMode,

            ImmutableArray<string> advertisedGroups,

            ImmutableArray<Outputs.RouterAdvertisedIpRangeResponse> advertisedIpRanges,

            int advertisedRoutePriority,

            Outputs.RouterBgpPeerBfdResponse bfd,

            string enable,

            string interfaceName,

            string ipAddress,

            string managementType,

            string name,

            int peerAsn,

            string peerIpAddress,

            string routerApplianceInstance)
        {
            AdvertiseMode = advertiseMode;
            AdvertisedGroups = advertisedGroups;
            AdvertisedIpRanges = advertisedIpRanges;
            AdvertisedRoutePriority = advertisedRoutePriority;
            Bfd = bfd;
            Enable = enable;
            InterfaceName = interfaceName;
            IpAddress = ipAddress;
            ManagementType = managementType;
            Name = name;
            PeerAsn = peerAsn;
            PeerIpAddress = peerIpAddress;
            RouterApplianceInstance = routerApplianceInstance;
        }
    }
}