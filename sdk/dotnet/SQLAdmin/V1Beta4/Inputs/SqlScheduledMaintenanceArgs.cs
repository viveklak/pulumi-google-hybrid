// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Inputs
{

    /// <summary>
    /// Any scheduled maintenancce for this instance.
    /// </summary>
    public sealed class SqlScheduledMaintenanceArgs : Pulumi.ResourceArgs
    {
        [Input("canDefer")]
        public Input<bool>? CanDefer { get; set; }

        /// <summary>
        /// If the scheduled maintenance can be rescheduled.
        /// </summary>
        [Input("canReschedule")]
        public Input<bool>? CanReschedule { get; set; }

        /// <summary>
        /// Maintenance cannot be rescheduled to start beyond this deadline.
        /// </summary>
        [Input("scheduleDeadlineTime")]
        public Input<string>? ScheduleDeadlineTime { get; set; }

        /// <summary>
        /// The start time of any upcoming scheduled maintenance for this instance.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public SqlScheduledMaintenanceArgs()
        {
        }
    }
}