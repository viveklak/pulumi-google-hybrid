// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudTasks.V2.Outputs
{

    [OutputType]
    public sealed class RateLimitsResponse
    {
        /// <summary>
        /// The max burst size. Max burst size limits how fast tasks in queue are processed when many tasks are in the queue and the rate is high. This field allows the queue to have a high rate so processing starts shortly after a task is enqueued, but still limits resource usage when many tasks are enqueued in a short period of time. The [token bucket](https://wikipedia.org/wiki/Token_Bucket) algorithm is used to control the rate of task dispatches. Each queue has a token bucket that holds tokens, up to the maximum specified by `max_burst_size`. Each time a task is dispatched, a token is removed from the bucket. Tasks will be dispatched until the queue's bucket runs out of tokens. The bucket will be continuously refilled with new tokens based on max_dispatches_per_second. Cloud Tasks will pick the value of `max_burst_size` based on the value of max_dispatches_per_second. For queues that were created or updated using `queue.yaml/xml`, `max_burst_size` is equal to [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size). Since `max_burst_size` is output only, if UpdateQueue is called on a queue created by `queue.yaml/xml`, `max_burst_size` will be reset based on the value of max_dispatches_per_second, regardless of whether max_dispatches_per_second is updated. 
        /// </summary>
        public readonly int MaxBurstSize;
        /// <summary>
        /// The maximum number of concurrent tasks that Cloud Tasks allows to be dispatched for this queue. After this threshold has been reached, Cloud Tasks stops dispatching tasks until the number of concurrent requests decreases. If unspecified when the queue is created, Cloud Tasks will pick the default. The maximum allowed value is 5,000. This field has the same meaning as [max_concurrent_requests in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
        /// </summary>
        public readonly int MaxConcurrentDispatches;
        /// <summary>
        /// The maximum rate at which tasks are dispatched from this queue. If unspecified when the queue is created, Cloud Tasks will pick the default. * The maximum allowed value is 500. This field has the same meaning as [rate in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
        /// </summary>
        public readonly double MaxDispatchesPerSecond;

        [OutputConstructor]
        private RateLimitsResponse(
            int maxBurstSize,

            int maxConcurrentDispatches,

            double maxDispatchesPerSecond)
        {
            MaxBurstSize = maxBurstSize;
            MaxConcurrentDispatches = maxConcurrentDispatches;
            MaxDispatchesPerSecond = maxDispatchesPerSecond;
        }
    }
}