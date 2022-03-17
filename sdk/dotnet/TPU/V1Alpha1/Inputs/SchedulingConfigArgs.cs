// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V1Alpha1.Inputs
{

    /// <summary>
    /// Sets the scheduling options for this node.
    /// </summary>
    public sealed class SchedulingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether the node is preemptible.
        /// </summary>
        [Input("preemptible")]
        public Input<bool>? Preemptible { get; set; }

        /// <summary>
        /// Whether the node is created under a reservation.
        /// </summary>
        [Input("reserved")]
        public Input<bool>? Reserved { get; set; }

        public SchedulingConfigArgs()
        {
        }
    }
}
