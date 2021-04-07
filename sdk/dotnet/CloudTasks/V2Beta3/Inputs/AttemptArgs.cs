// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudTasks.V2Beta3.Inputs
{

    /// <summary>
    /// The status of a task attempt.
    /// </summary>
    public sealed class AttemptArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time that this attempt was dispatched. `dispatch_time` will be truncated to the nearest microsecond.
        /// </summary>
        [Input("dispatchTime")]
        public Input<string>? DispatchTime { get; set; }

        /// <summary>
        /// The response from the worker for this attempt. If `response_time` is unset, then the task has not been attempted or is currently running and the `response_status` field is meaningless.
        /// </summary>
        [Input("responseStatus")]
        public Input<Inputs.StatusArgs>? ResponseStatus { get; set; }

        /// <summary>
        /// The time that this attempt response was received. `response_time` will be truncated to the nearest microsecond.
        /// </summary>
        [Input("responseTime")]
        public Input<string>? ResponseTime { get; set; }

        /// <summary>
        /// The time that this attempt was scheduled. `schedule_time` will be truncated to the nearest microsecond.
        /// </summary>
        [Input("scheduleTime")]
        public Input<string>? ScheduleTime { get; set; }

        public AttemptArgs()
        {
        }
    }
}