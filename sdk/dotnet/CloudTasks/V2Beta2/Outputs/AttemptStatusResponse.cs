// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta2.Outputs
{

    /// <summary>
    /// The status of a task attempt.
    /// </summary>
    [OutputType]
    public sealed class AttemptStatusResponse
    {
        /// <summary>
        /// The time that this attempt was dispatched. `dispatch_time` will be truncated to the nearest microsecond.
        /// </summary>
        public readonly string DispatchTime;
        /// <summary>
        /// The response from the target for this attempt. If the task has not been attempted or the task is currently running then the response status is unset.
        /// </summary>
        public readonly Outputs.StatusResponse ResponseStatus;
        /// <summary>
        /// The time that this attempt response was received. `response_time` will be truncated to the nearest microsecond.
        /// </summary>
        public readonly string ResponseTime;
        /// <summary>
        /// The time that this attempt was scheduled. `schedule_time` will be truncated to the nearest microsecond.
        /// </summary>
        public readonly string ScheduleTime;

        [OutputConstructor]
        private AttemptStatusResponse(
            string dispatchTime,

            Outputs.StatusResponse responseStatus,

            string responseTime,

            string scheduleTime)
        {
            DispatchTime = dispatchTime;
            ResponseStatus = responseStatus;
            ResponseTime = responseTime;
            ScheduleTime = scheduleTime;
        }
    }
}
