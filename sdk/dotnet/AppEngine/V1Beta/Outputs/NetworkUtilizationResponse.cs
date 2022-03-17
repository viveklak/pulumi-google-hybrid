// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// Target scaling by network usage. Only applicable in the App Engine flexible environment.
    /// </summary>
    [OutputType]
    public sealed class NetworkUtilizationResponse
    {
        /// <summary>
        /// Target bytes received per second.
        /// </summary>
        public readonly int TargetReceivedBytesPerSecond;
        /// <summary>
        /// Target packets received per second.
        /// </summary>
        public readonly int TargetReceivedPacketsPerSecond;
        /// <summary>
        /// Target bytes sent per second.
        /// </summary>
        public readonly int TargetSentBytesPerSecond;
        /// <summary>
        /// Target packets sent per second.
        /// </summary>
        public readonly int TargetSentPacketsPerSecond;

        [OutputConstructor]
        private NetworkUtilizationResponse(
            int targetReceivedBytesPerSecond,

            int targetReceivedPacketsPerSecond,

            int targetSentBytesPerSecond,

            int targetSentPacketsPerSecond)
        {
            TargetReceivedBytesPerSecond = targetReceivedBytesPerSecond;
            TargetReceivedPacketsPerSecond = targetReceivedPacketsPerSecond;
            TargetSentBytesPerSecond = targetSentBytesPerSecond;
            TargetSentPacketsPerSecond = targetSentPacketsPerSecond;
        }
    }
}
