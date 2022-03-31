// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Memcache.V1.Outputs
{

    /// <summary>
    /// Time window specified for weekly operations.
    /// </summary>
    [OutputType]
    public sealed class WeeklyMaintenanceWindowResponse
    {
        /// <summary>
        /// Allows to define schedule that runs specified day of the week.
        /// </summary>
        public readonly string Day;
        /// <summary>
        /// Duration of the time window.
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Start time of the window in UTC.
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
