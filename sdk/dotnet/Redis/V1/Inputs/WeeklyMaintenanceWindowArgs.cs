// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Redis.V1.Inputs
{

    /// <summary>
    /// Time window in which disruptive maintenance updates occur. Non-disruptive updates can occur inside or outside this window.
    /// </summary>
    public sealed class WeeklyMaintenanceWindowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The day of week that maintenance updates occur.
        /// </summary>
        [Input("day")]
        public Input<Pulumi.GoogleNative.Redis.V1.WeeklyMaintenanceWindowDay>? Day { get; set; }

        /// <summary>
        /// Required. Start time of the window in UTC time.
        /// </summary>
        [Input("startTime")]
        public Input<Inputs.TimeOfDayArgs>? StartTime { get; set; }

        public WeeklyMaintenanceWindowArgs()
        {
        }
    }
}
