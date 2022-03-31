// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// A routing configuration attached to a network resource. The message includes the list of routers associated with the network, and a flag indicating the type of routing behavior to enforce network-wide.
    /// </summary>
    public sealed class NetworkRoutingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network-wide routing mode to use. If set to REGIONAL, this network's Cloud Routers will only advertise routes with subnets of this network in the same region as the router. If set to GLOBAL, this network's Cloud Routers will advertise routes with all subnets of this network, across regions.
        /// </summary>
        [Input("routingMode")]
        public Input<Pulumi.GoogleNative.Compute.V1.NetworkRoutingConfigRoutingMode>? RoutingMode { get; set; }

        public NetworkRoutingConfigArgs()
        {
        }
    }
}
