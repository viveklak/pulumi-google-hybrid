// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Commitment for a particular resource (a Commitment is composed of one or more of these).
    /// </summary>
    public sealed class ResourceCommitmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the accelerator type resource. Applicable only when the type is ACCELERATOR.
        /// </summary>
        [Input("acceleratorType")]
        public Input<string>? AcceleratorType { get; set; }

        /// <summary>
        /// The amount of the resource purchased (in a type-dependent unit, such as bytes). For vCPUs, this can just be an integer. For memory, this must be provided in MB. Memory must be a multiple of 256 MB, with up to 6.5GB of memory per every vCPU.
        /// </summary>
        [Input("amount")]
        public Input<string>? Amount { get; set; }

        /// <summary>
        /// Type of resource for which this commitment applies. Possible values are VCPU and MEMORY
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.Beta.ResourceCommitmentType>? Type { get; set; }

        public ResourceCommitmentArgs()
        {
        }
    }
}
