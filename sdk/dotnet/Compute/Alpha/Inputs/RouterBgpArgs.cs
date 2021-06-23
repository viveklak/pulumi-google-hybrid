// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class RouterBgpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-specified flag to indicate which mode to use for advertisement. The options are DEFAULT or CUSTOM.
        /// </summary>
        [Input("advertiseMode")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.RouterBgpAdvertiseMode>? AdvertiseMode { get; set; }

        [Input("advertisedGroups")]
        private InputList<Pulumi.GoogleNative.Compute.Alpha.RouterBgpAdvertisedGroupsItem>? _advertisedGroups;

        /// <summary>
        /// User-specified list of prefix groups to advertise in custom mode. This field can only be populated if advertise_mode is CUSTOM and is advertised to all peers of the router. These groups will be advertised in addition to any specified prefixes. Leave this field blank to advertise no custom groups.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Compute.Alpha.RouterBgpAdvertisedGroupsItem> AdvertisedGroups
        {
            get => _advertisedGroups ?? (_advertisedGroups = new InputList<Pulumi.GoogleNative.Compute.Alpha.RouterBgpAdvertisedGroupsItem>());
            set => _advertisedGroups = value;
        }

        [Input("advertisedIpRanges")]
        private InputList<Inputs.RouterAdvertisedIpRangeArgs>? _advertisedIpRanges;

        /// <summary>
        /// User-specified list of individual IP ranges to advertise in custom mode. This field can only be populated if advertise_mode is CUSTOM and is advertised to all peers of the router. These IP ranges will be advertised in addition to any specified groups. Leave this field blank to advertise no custom IP ranges.
        /// </summary>
        public InputList<Inputs.RouterAdvertisedIpRangeArgs> AdvertisedIpRanges
        {
            get => _advertisedIpRanges ?? (_advertisedIpRanges = new InputList<Inputs.RouterAdvertisedIpRangeArgs>());
            set => _advertisedIpRanges = value;
        }

        /// <summary>
        /// Local BGP Autonomous System Number (ASN). Must be an RFC6996 private ASN, either 16-bit or 32-bit. The value will be fixed for this router resource. All VPN tunnels that link to this router will have the same local ASN.
        /// </summary>
        [Input("asn")]
        public Input<int>? Asn { get; set; }

        /// <summary>
        /// The interval in seconds between BGP keepalive messages that are sent to the peer. Hold time is three times the interval at which keepalive messages are sent, and the hold time is the maximum number of seconds allowed to elapse between successive keepalive messages that BGP receives from a peer. BGP will use the smaller of either the local hold time value or the peer's hold time value as the hold time for the BGP connection between the two peers. If set, this value must be between 20 and 60. The default is 20.
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<int>? KeepaliveInterval { get; set; }

        public RouterBgpArgs()
        {
        }
    }
}
