// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Outputs
{

    [OutputType]
    public sealed class RecurringScheduleResponse
    {
        /// <summary>
        /// Optional. The end time at which a recurring patch deployment schedule is no longer active.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The frequency unit of this recurring schedule.
        /// </summary>
        public readonly string Frequency;
        /// <summary>
        /// The time the last patch job ran successfully.
        /// </summary>
        public readonly string LastExecuteTime;
        /// <summary>
        /// Schedule with monthly executions.
        /// </summary>
        public readonly Outputs.MonthlyScheduleResponse Monthly;
        /// <summary>
        /// The time the next patch job is scheduled to run.
        /// </summary>
        public readonly string NextExecuteTime;
        /// <summary>
        /// Optional. The time that the recurring schedule becomes effective. Defaults to `create_time` of the patch deployment.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Time of the day to run a recurring deployment.
        /// </summary>
        public readonly Outputs.TimeOfDayResponse TimeOfDay;
        /// <summary>
        /// Defines the time zone that `time_of_day` is relative to. The rules for daylight saving time are determined by the chosen time zone.
        /// </summary>
        public readonly Outputs.TimeZoneResponse TimeZone;
        /// <summary>
        /// Schedule with weekly executions.
        /// </summary>
        public readonly Outputs.WeeklyScheduleResponse Weekly;

        [OutputConstructor]
        private RecurringScheduleResponse(
            string endTime,

            string frequency,

            string lastExecuteTime,

            Outputs.MonthlyScheduleResponse monthly,

            string nextExecuteTime,

            string startTime,

            Outputs.TimeOfDayResponse timeOfDay,

            Outputs.TimeZoneResponse timeZone,

            Outputs.WeeklyScheduleResponse weekly)
        {
            EndTime = endTime;
            Frequency = frequency;
            LastExecuteTime = lastExecuteTime;
            Monthly = monthly;
            NextExecuteTime = nextExecuteTime;
            StartTime = startTime;
            TimeOfDay = timeOfDay;
            TimeZone = timeZone;
            Weekly = weekly;
        }
    }
}
