// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// This reservation type allows to pre allocate specific instance configuration.
    /// </summary>
    public sealed class AllocationSpecificSKUReservationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the number of resources that are allocated.
        /// </summary>
        [Input("count")]
        public Input<string>? Count { get; set; }

        /// <summary>
        /// The instance properties for the reservation.
        /// </summary>
        [Input("instanceProperties")]
        public Input<Inputs.AllocationSpecificSKUAllocationReservedInstancePropertiesArgs>? InstanceProperties { get; set; }

        public AllocationSpecificSKUReservationArgs()
        {
        }
    }
}
