// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2Beta3.Outputs
{

    /// <summary>
    /// Statistics for a queue.
    /// </summary>
    [OutputType]
    public sealed class QueueStatsResponse
    {
        /// <summary>
        /// The number of requests that the queue has dispatched but has not received a reply for yet.
        /// </summary>
        public readonly string ConcurrentDispatchesCount;
        /// <summary>
        /// The current maximum number of tasks per second executed by the queue. The maximum value of this variable is controlled by the RateLimits of the Queue. However, this value could be less to avoid overloading the endpoints tasks in the queue are targeting.
        /// </summary>
        public readonly double EffectiveExecutionRate;
        /// <summary>
        /// The number of tasks that the queue has dispatched and received a reply for during the last minute. This variable counts both successful and non-successful executions.
        /// </summary>
        public readonly string ExecutedLastMinuteCount;
        /// <summary>
        /// An estimation of the nearest time in the future where a task in the queue is scheduled to be executed.
        /// </summary>
        public readonly string OldestEstimatedArrivalTime;
        /// <summary>
        /// An estimation of the number of tasks in the queue, that is, the tasks in the queue that haven't been executed, the tasks in the queue which the queue has dispatched but has not yet received a reply for, and the failed tasks that the queue is retrying.
        /// </summary>
        public readonly string TasksCount;

        [OutputConstructor]
        private QueueStatsResponse(
            string concurrentDispatchesCount,

            double effectiveExecutionRate,

            string executedLastMinuteCount,

            string oldestEstimatedArrivalTime,

            string tasksCount)
        {
            ConcurrentDispatchesCount = concurrentDispatchesCount;
            EffectiveExecutionRate = effectiveExecutionRate;
            ExecutedLastMinuteCount = executedLastMinuteCount;
            OldestEstimatedArrivalTime = oldestEstimatedArrivalTime;
            TasksCount = tasksCount;
        }
    }
}
