// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class PacketMirroringForwardingRuleInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource URL to the forwarding rule representing the ILB configured as destination of the mirrored traffic.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public PacketMirroringForwardingRuleInfoArgs()
        {
        }
    }
}
