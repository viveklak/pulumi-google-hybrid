// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2Alpha1.Inputs
{

    /// <summary>
    /// Network related configurations.
    /// </summary>
    public sealed class NetworkConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows the TPU node to send and receive packets with non-matching destination or source IPs. This is required if you plan to use the TPU workers to forward routes.
        /// </summary>
        [Input("canIpForward")]
        public Input<bool>? CanIpForward { get; set; }

        /// <summary>
        /// Indicates that external IP addresses would be associated with the TPU workers. If set to false, the specified subnetwork or network should have Private Google Access enabled.
        /// </summary>
        [Input("enableExternalIps")]
        public Input<bool>? EnableExternalIps { get; set; }

        /// <summary>
        /// The name of the network for the TPU node. It must be a preexisting Google Compute Engine network. If none is provided, "default" will be used.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The name of the subnetwork for the TPU node. It must be a preexisting Google Compute Engine subnetwork. If none is provided, "default" will be used.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        public NetworkConfigArgs()
        {
        }
    }
}
