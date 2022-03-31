// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Represents a weekly schedule.
    /// </summary>
    [OutputType]
    public sealed class WeeklyScheduleResponse
    {
        /// <summary>
        /// Day of the week.
        /// </summary>
        public readonly string DayOfWeek;

        [OutputConstructor]
        private WeeklyScheduleResponse(string dayOfWeek)
        {
            DayOfWeek = dayOfWeek;
        }
    }
}
