// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for resources used by Airflow workers.
    /// </summary>
    [OutputType]
    public sealed class WorkerResourceResponse
    {
        /// <summary>
        /// Optional. CPU request and limit for a single Airflow worker replica.
        /// </summary>
        public readonly double Cpu;
        /// <summary>
        /// Optional. Maximum number of workers for autoscaling.
        /// </summary>
        public readonly int MaxCount;
        /// <summary>
        /// Optional. Memory (GB) request and limit for a single Airflow worker replica.
        /// </summary>
        public readonly double MemoryGb;
        /// <summary>
        /// Optional. Minimum number of workers for autoscaling.
        /// </summary>
        public readonly int MinCount;
        /// <summary>
        /// Optional. Storage (GB) request and limit for a single Airflow worker replica.
        /// </summary>
        public readonly double StorageGb;

        [OutputConstructor]
        private WorkerResourceResponse(
            double cpu,

            int maxCount,

            double memoryGb,

            int minCount,

            double storageGb)
        {
            Cpu = cpu;
            MaxCount = maxCount;
            MemoryGb = memoryGb;
            MinCount = minCount;
            StorageGb = storageGb;
        }
    }
}
