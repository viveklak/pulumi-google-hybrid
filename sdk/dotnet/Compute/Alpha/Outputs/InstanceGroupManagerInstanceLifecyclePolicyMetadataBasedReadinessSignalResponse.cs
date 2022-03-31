// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class InstanceGroupManagerInstanceLifecyclePolicyMetadataBasedReadinessSignalResponse
    {
        /// <summary>
        /// The number of seconds to wait for a readiness signal during initialization before timing out.
        /// </summary>
        public readonly int TimeoutSec;

        [OutputConstructor]
        private InstanceGroupManagerInstanceLifecyclePolicyMetadataBasedReadinessSignalResponse(int timeoutSec)
        {
            TimeoutSec = timeoutSec;
        }
    }
}
