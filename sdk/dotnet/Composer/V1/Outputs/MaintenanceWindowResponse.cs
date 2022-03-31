// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1.Outputs
{

    /// <summary>
    /// The configuration settings for Cloud Composer maintenance window. The following example: ``` { "startTime":"2019-08-01T01:00:00Z" "endTime":"2019-08-01T07:00:00Z" "recurrence":"FREQ=WEEKLY;BYDAY=TU,WE" } ``` would define a maintenance window between 01 and 07 hours UTC during each Tuesday and Wednesday.
    /// </summary>
    [OutputType]
    public sealed class MaintenanceWindowResponse
    {
        /// <summary>
        /// Maintenance window end time. It is used only to calculate the duration of the maintenance window. The value for end-time must be in the future, relative to `start_time`.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Maintenance window recurrence. Format is a subset of [RFC-5545](https://tools.ietf.org/html/rfc5545) `RRULE`. The only allowed values for `FREQ` field are `FREQ=DAILY` and `FREQ=WEEKLY;BYDAY=...` Example values: `FREQ=WEEKLY;BYDAY=TU,WE`, `FREQ=DAILY`.
        /// </summary>
        public readonly string Recurrence;
        /// <summary>
        /// Start time of the first recurrence of the maintenance window.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private MaintenanceWindowResponse(
            string endTime,

            string recurrence,

            string startTime)
        {
            EndTime = endTime;
            Recurrence = recurrence;
            StartTime = startTime;
        }
    }
}
