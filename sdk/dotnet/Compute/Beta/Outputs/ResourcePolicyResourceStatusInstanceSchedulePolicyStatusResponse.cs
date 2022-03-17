// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class ResourcePolicyResourceStatusInstanceSchedulePolicyStatusResponse
    {
        /// <summary>
        /// The last time the schedule successfully ran. The timestamp is an RFC3339 string.
        /// </summary>
        public readonly string LastRunStartTime;
        /// <summary>
        /// The next time the schedule is planned to run. The actual time might be slightly different. The timestamp is an RFC3339 string.
        /// </summary>
        public readonly string NextRunStartTime;

        [OutputConstructor]
        private ResourcePolicyResourceStatusInstanceSchedulePolicyStatusResponse(
            string lastRunStartTime,

            string nextRunStartTime)
        {
            LastRunStartTime = lastRunStartTime;
            NextRunStartTime = nextRunStartTime;
        }
    }
}
