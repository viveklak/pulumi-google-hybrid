// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Sets the time for recurring patch deployments.
    /// </summary>
    public sealed class RecurringScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The end time at which a recurring patch deployment schedule is no longer active.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Required. The frequency unit of this recurring schedule.
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        /// <summary>
        /// Required. Schedule with monthly executions.
        /// </summary>
        [Input("monthly")]
        public Input<Inputs.MonthlyScheduleArgs>? Monthly { get; set; }

        /// <summary>
        /// Optional. The time that the recurring schedule becomes effective. Defaults to `create_time` of the patch deployment.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Required. Time of the day to run a recurring deployment.
        /// </summary>
        [Input("timeOfDay")]
        public Input<Inputs.TimeOfDayArgs>? TimeOfDay { get; set; }

        /// <summary>
        /// Required. Defines the time zone that `time_of_day` is relative to. The rules for daylight saving time are determined by the chosen time zone.
        /// </summary>
        [Input("timeZone")]
        public Input<Inputs.TimeZoneArgs>? TimeZone { get; set; }

        /// <summary>
        /// Required. Schedule with weekly executions.
        /// </summary>
        [Input("weekly")]
        public Input<Inputs.WeeklyScheduleArgs>? Weekly { get; set; }

        public RecurringScheduleArgs()
        {
        }
    }
}