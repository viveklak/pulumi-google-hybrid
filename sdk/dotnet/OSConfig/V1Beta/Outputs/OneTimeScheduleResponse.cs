// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Sets the time for a one time patch deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
    /// </summary>
    [OutputType]
    public sealed class OneTimeScheduleResponse
    {
        /// <summary>
        /// The desired patch job execution time.
        /// </summary>
        public readonly string ExecuteTime;

        [OutputConstructor]
        private OneTimeScheduleResponse(string executeTime)
        {
            ExecuteTime = executeTime;
        }
    }
}