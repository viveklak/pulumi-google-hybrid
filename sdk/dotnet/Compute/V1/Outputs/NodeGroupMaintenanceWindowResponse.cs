// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Time window specified for daily maintenance operations. GCE's internal maintenance will be performed within this window.
    /// </summary>
    [OutputType]
    public sealed class NodeGroupMaintenanceWindowResponse
    {
        /// <summary>
        /// [Output only] A predetermined duration for the window, automatically chosen to be the smallest possible in the given scenario.
        /// </summary>
        public readonly Outputs.DurationResponse MaintenanceDuration;
        /// <summary>
        /// Start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private NodeGroupMaintenanceWindowResponse(
            Outputs.DurationResponse maintenanceDuration,

            string startTime)
        {
            MaintenanceDuration = maintenanceDuration;
            StartTime = startTime;
        }
    }
}
