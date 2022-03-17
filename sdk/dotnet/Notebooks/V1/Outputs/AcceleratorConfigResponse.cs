// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1.Outputs
{

    /// <summary>
    /// Definition of a hardware accelerator. Note that not all combinations of `type` and `core_count` are valid. Check [GPUs on Compute Engine](/compute/docs/gpus/#gpus-list) to find a valid combination. TPUs are not supported.
    /// </summary>
    [OutputType]
    public sealed class AcceleratorConfigResponse
    {
        /// <summary>
        /// Count of cores of this accelerator.
        /// </summary>
        public readonly string CoreCount;
        /// <summary>
        /// Type of this accelerator.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AcceleratorConfigResponse(
            string coreCount,

            string type)
        {
            CoreCount = coreCount;
            Type = type;
        }
    }
}
