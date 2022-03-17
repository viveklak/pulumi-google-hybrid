// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// A specification of the type and number of accelerator cards attached to the instance.
    /// </summary>
    public sealed class AcceleratorConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of the guest accelerator cards exposed to this instance.
        /// </summary>
        [Input("acceleratorCount")]
        public Input<int>? AcceleratorCount { get; set; }

        /// <summary>
        /// Full or partial URL of the accelerator type resource to attach to this instance. For example: projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100 If you are creating an instance template, specify only the accelerator name. See GPUs on Compute Engine for a full list of accelerator types.
        /// </summary>
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        public AcceleratorConfigArgs()
        {
        }
    }
}
