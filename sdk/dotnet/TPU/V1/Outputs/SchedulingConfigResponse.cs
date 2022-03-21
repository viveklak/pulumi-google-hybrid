// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V1.Outputs
{

    /// <summary>
    /// Sets the scheduling options for this node.
    /// </summary>
    [OutputType]
    public sealed class SchedulingConfigResponse
    {
        /// <summary>
        /// Defines whether the node is preemptible.
        /// </summary>
        public readonly bool Preemptible;
        /// <summary>
        /// Whether the node is created under a reservation.
        /// </summary>
        public readonly bool Reserved;

        [OutputConstructor]
        private SchedulingConfigResponse(
            bool preemptible,

            bool reserved)
        {
            Preemptible = preemptible;
            Reserved = reserved;
        }
    }
}
