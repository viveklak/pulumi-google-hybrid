// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// WorkloadMetadataConfig defines the metadata configuration to expose to workloads on the node pool.
    /// </summary>
    public sealed class WorkloadMetadataConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode is the configuration for how to expose metadata to workloads running on the node pool.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// NodeMetadata is the configuration for how to expose metadata to the workloads running on the node.
        /// </summary>
        [Input("nodeMetadata")]
        public Input<string>? NodeMetadata { get; set; }

        public WorkloadMetadataConfigArgs()
        {
        }
    }
}