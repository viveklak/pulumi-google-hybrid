// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V1.Inputs
{

    /// <summary>
    /// Definition of the types of hardware accelerators that can be used. Definition of the types of hardware accelerators that can be used. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `nvidia-tesla-k80` * `nvidia-tesla-p100` * `nvidia-tesla-v100` * `nvidia-tesla-p4` * `nvidia-tesla-t4` * `nvidia-tesla-a100`
    /// </summary>
    public sealed class RuntimeAcceleratorConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Count of cores of this accelerator.
        /// </summary>
        [Input("coreCount")]
        public Input<string>? CoreCount { get; set; }

        /// <summary>
        /// Accelerator model.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Notebooks.V1.RuntimeAcceleratorConfigType>? Type { get; set; }

        public RuntimeAcceleratorConfigArgs()
        {
        }
    }
}
