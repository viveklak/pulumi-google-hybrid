// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1.Inputs
{

    /// <summary>
    /// RouterAppliance represents a Router appliance which is specified by a VM URI and a NIC address.
    /// </summary>
    public sealed class RouterApplianceInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the network interface to use for peering.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("networkInterface")]
        public Input<string>? NetworkInterface { get; set; }

        /// <summary>
        /// The URI of the virtual machine resource
        /// </summary>
        [Input("virtualMachine")]
        public Input<string>? VirtualMachine { get; set; }

        public RouterApplianceInstanceArgs()
        {
        }
    }
}
