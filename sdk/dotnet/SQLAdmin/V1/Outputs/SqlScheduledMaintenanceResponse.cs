// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1.Outputs
{

    /// <summary>
    /// Any scheduled maintenancce for this instance.
    /// </summary>
    [OutputType]
    public sealed class SqlScheduledMaintenanceResponse
    {
        public readonly bool CanDefer;
        /// <summary>
        /// If the scheduled maintenance can be rescheduled.
        /// </summary>
        public readonly bool CanReschedule;
        /// <summary>
        /// Maintenance cannot be rescheduled to start beyond this deadline.
        /// </summary>
        public readonly string ScheduleDeadlineTime;
        /// <summary>
        /// The start time of any upcoming scheduled maintenance for this instance.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private SqlScheduledMaintenanceResponse(
            bool canDefer,

            bool canReschedule,

            string scheduleDeadlineTime,

            string startTime)
        {
            CanDefer = canDefer;
            CanReschedule = canReschedule;
            ScheduleDeadlineTime = scheduleDeadlineTime;
            StartTime = startTime;
        }
    }
}
