// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Redis.V1Beta1.Outputs
{

    /// <summary>
    /// Time window in which disruptive maintenance updates occur. Non-disruptive updates can occur inside or outside this window.
    /// </summary>
    [OutputType]
    public sealed class WeeklyMaintenanceWindowResponse
    {
        /// <summary>
        /// The day of week that maintenance updates occur.
        /// </summary>
        public readonly string Day;
        /// <summary>
        /// Duration of the maintenance window. The current window is fixed at 1 hour.
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Start time of the window in UTC time.
        /// </summary>
        public readonly Outputs.TimeOfDayResponse StartTime;

        [OutputConstructor]
        private WeeklyMaintenanceWindowResponse(
            string day,

            string duration,

            Outputs.TimeOfDayResponse startTime)
        {
            Day = day;
            Duration = duration;
            StartTime = startTime;
        }
    }
}
