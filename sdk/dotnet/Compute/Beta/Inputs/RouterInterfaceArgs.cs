// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class RouterInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address and range of the interface. The IP range must be in the RFC3927 link-local IP address space. The value must be a CIDR-formatted string, for example: 169.254.0.1/30. NOTE: Do not truncate the address as it represents the IP address of the interface.
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// URI of the linked Interconnect attachment. It must be in the same region as the router. Each interface can have one linked resource, which can be a VPN tunnel, an Interconnect attachment, or a virtual machine instance.
        /// </summary>
        [Input("linkedInterconnectAttachment")]
        public Input<string>? LinkedInterconnectAttachment { get; set; }

        /// <summary>
        /// URI of the linked VPN tunnel, which must be in the same region as the router. Each interface can have one linked resource, which can be a VPN tunnel, an Interconnect attachment, or a virtual machine instance.
        /// </summary>
        [Input("linkedVpnTunnel")]
        public Input<string>? LinkedVpnTunnel { get; set; }

        /// <summary>
        /// Name of this interface entry. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The regional private internal IP address that is used to establish BGP sessions to a VM instance acting as a third-party Router Appliance, such as a Next Gen Firewall, a Virtual Router, or an SD-WAN VM.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// Name of the interface that will be redundant with the current interface you are creating. The redundantInterface must belong to the same Cloud Router as the interface here. To establish the BGP session to a Router Appliance VM, you must create two BGP peers. The two BGP peers must be attached to two separate interfaces that are redundant with each other. The redundant_interface must be 1-63 characters long, and comply with RFC1035. Specifically, the redundant_interface must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("redundantInterface")]
        public Input<string>? RedundantInterface { get; set; }

        /// <summary>
        /// The URI of the subnetwork resource that this interface belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        public RouterInterfaceArgs()
        {
        }
    }
}
