// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Inputs
{

    /// <summary>
    /// Represents a weekly schedule.
    /// </summary>
    public sealed class WeeklyScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of the week.
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<Pulumi.GoogleNative.OSConfig.V1.WeeklyScheduleDayOfWeek> DayOfWeek { get; set; } = null!;

        public WeeklyScheduleArgs()
        {
        }
    }
}
