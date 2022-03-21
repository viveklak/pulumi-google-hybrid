// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// Compute resources associated with the analyze interactive workloads.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1EnvironmentInfrastructureSpecComputeResourcesResponse
    {
        /// <summary>
        /// Optional. Size in GB of the disk. Default is 100 GB.
        /// </summary>
        public readonly int DiskSizeGb;
        /// <summary>
        /// Optional. Max configurable nodes. If max_node_count &gt; node_count, then auto-scaling is enabled.
        /// </summary>
        public readonly int MaxNodeCount;
        /// <summary>
        /// Optional. Total number of nodes in the sessions created for this environment.
        /// </summary>
        public readonly int NodeCount;

        [OutputConstructor]
        private GoogleCloudDataplexV1EnvironmentInfrastructureSpecComputeResourcesResponse(
            int diskSizeGb,

            int maxNodeCount,

            int nodeCount)
        {
            DiskSizeGb = diskSizeGb;
            MaxNodeCount = maxNodeCount;
            NodeCount = nodeCount;
        }
    }
}
