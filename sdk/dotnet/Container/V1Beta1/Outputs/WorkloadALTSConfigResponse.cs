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
    /// Configuration for direct-path (via ALTS) with workload identity.
    /// </summary>
    [OutputType]
    public sealed class WorkloadALTSConfigResponse
    {
        /// <summary>
        /// enable_alts controls whether the alts handshaker should be enabled or not for direct-path. Requires Workload Identity (workload_pool must be non-empty).
        /// </summary>
        public readonly bool EnableAlts;

        [OutputConstructor]
        private WorkloadALTSConfigResponse(bool enableAlts)
        {
            EnableAlts = enableAlts;
        }
    }
}
