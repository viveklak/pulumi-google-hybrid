// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// AcceleratorConfig represents a Hardware Accelerator request.
    /// </summary>
    [OutputType]
    public sealed class AcceleratorConfigResponse
    {
        /// <summary>
        /// The number of the accelerator cards exposed to an instance.
        /// </summary>
        public readonly string AcceleratorCount;
        /// <summary>
        /// The accelerator type resource name. List of supported accelerators [here](https://cloud.google.com/compute/docs/gpus)
        /// </summary>
        public readonly string AcceleratorType;
        /// <summary>
        /// Size of partitions to create on the GPU. Valid values are described in the NVIDIA [mig user guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
        /// </summary>
        public readonly string GpuPartitionSize;

        [OutputConstructor]
        private AcceleratorConfigResponse(
            string acceleratorCount,

            string acceleratorType,

            string gpuPartitionSize)
        {
            AcceleratorCount = acceleratorCount;
            AcceleratorType = acceleratorType;
            GpuPartitionSize = gpuPartitionSize;
        }
    }
}
