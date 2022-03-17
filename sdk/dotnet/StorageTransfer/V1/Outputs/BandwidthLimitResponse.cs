// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Outputs
{

    /// <summary>
    /// Specifies a bandwidth limit for an agent pool.
    /// </summary>
    [OutputType]
    public sealed class BandwidthLimitResponse
    {
        /// <summary>
        /// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
        /// </summary>
        public readonly string LimitMbps;

        [OutputConstructor]
        private BandwidthLimitResponse(string limitMbps)
        {
            LimitMbps = limitMbps;
        }
    }
}
