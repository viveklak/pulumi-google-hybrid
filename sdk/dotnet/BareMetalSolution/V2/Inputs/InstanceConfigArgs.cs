// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Inputs
{

    /// <summary>
    /// Configuration parameters for a new instance.
    /// </summary>
    public sealed class InstanceConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client network address.
        /// </summary>
        [Input("clientNetwork")]
        public Input<Inputs.NetworkAddressArgs>? ClientNetwork { get; set; }

        /// <summary>
        /// Whether the instance should be provisioned with Hyperthreading enabled.
        /// </summary>
        [Input("hyperthreading")]
        public Input<bool>? Hyperthreading { get; set; }

        /// <summary>
        /// A transient unique identifier to idenfity an instance within an ProvisioningConfig request.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Instance type. [Available types](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// OS image to initialize the instance. [Available images](https://cloud.google.com/bare-metal/docs/bms-planning#server_configurations)
        /// </summary>
        [Input("osImage")]
        public Input<string>? OsImage { get; set; }

        /// <summary>
        /// Private network address, if any.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.NetworkAddressArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// User note field, it can be used by customers to add additional information for the BMS Ops team .
        /// </summary>
        [Input("userNote")]
        public Input<string>? UserNote { get; set; }

        public InstanceConfigArgs()
        {
        }
    }
}
