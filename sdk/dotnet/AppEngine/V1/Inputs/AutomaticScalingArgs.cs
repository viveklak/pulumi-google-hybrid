// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Inputs
{

    /// <summary>
    /// Automatic scaling is based on request rate, response latencies, and other application metrics.
    /// </summary>
    public sealed class AutomaticScalingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time period that the Autoscaler (https://cloud.google.com/compute/docs/autoscaler/) should wait before it starts collecting information from a new instance. This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable. Only applicable in the App Engine flexible environment.
        /// </summary>
        [Input("coolDownPeriod")]
        public Input<string>? CoolDownPeriod { get; set; }

        /// <summary>
        /// Target scaling by CPU usage.
        /// </summary>
        [Input("cpuUtilization")]
        public Input<Inputs.CpuUtilizationArgs>? CpuUtilization { get; set; }

        /// <summary>
        /// Target scaling by disk usage.
        /// </summary>
        [Input("diskUtilization")]
        public Input<Inputs.DiskUtilizationArgs>? DiskUtilization { get; set; }

        /// <summary>
        /// Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.Defaults to a runtime-specific value.
        /// </summary>
        [Input("maxConcurrentRequests")]
        public Input<int>? MaxConcurrentRequests { get; set; }

        /// <summary>
        /// Maximum number of idle instances that should be maintained for this version.
        /// </summary>
        [Input("maxIdleInstances")]
        public Input<int>? MaxIdleInstances { get; set; }

        /// <summary>
        /// Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.
        /// </summary>
        [Input("maxPendingLatency")]
        public Input<string>? MaxPendingLatency { get; set; }

        /// <summary>
        /// Maximum number of instances that should be started to handle requests for this version.
        /// </summary>
        [Input("maxTotalInstances")]
        public Input<int>? MaxTotalInstances { get; set; }

        /// <summary>
        /// Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.
        /// </summary>
        [Input("minIdleInstances")]
        public Input<int>? MinIdleInstances { get; set; }

        /// <summary>
        /// Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.
        /// </summary>
        [Input("minPendingLatency")]
        public Input<string>? MinPendingLatency { get; set; }

        /// <summary>
        /// Minimum number of running instances that should be maintained for this version.
        /// </summary>
        [Input("minTotalInstances")]
        public Input<int>? MinTotalInstances { get; set; }

        /// <summary>
        /// Target scaling by network usage.
        /// </summary>
        [Input("networkUtilization")]
        public Input<Inputs.NetworkUtilizationArgs>? NetworkUtilization { get; set; }

        /// <summary>
        /// Target scaling by request utilization.
        /// </summary>
        [Input("requestUtilization")]
        public Input<Inputs.RequestUtilizationArgs>? RequestUtilization { get; set; }

        /// <summary>
        /// Scheduler settings for standard environment.
        /// </summary>
        [Input("standardSchedulerSettings")]
        public Input<Inputs.StandardSchedulerSettingsArgs>? StandardSchedulerSettings { get; set; }

        public AutomaticScalingArgs()
        {
        }
    }
}
