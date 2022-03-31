// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Outputs
{

    /// <summary>
    /// Specifies the cluster auto-delete schedule configuration.
    /// </summary>
    [OutputType]
    public sealed class LifecycleConfigResponse
    {
        /// <summary>
        /// Optional. The time when cluster will be auto-deleted. (see JSON representation of Timestamp (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        public readonly string AutoDeleteTime;
        /// <summary>
        /// Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of Duration (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        public readonly string AutoDeleteTtl;
        /// <summary>
        /// Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of Duration (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        public readonly string IdleDeleteTtl;
        /// <summary>
        /// The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of Timestamp (https://developers.google.com/protocol-buffers/docs/proto3#json)).
        /// </summary>
        public readonly string IdleStartTime;

        [OutputConstructor]
        private LifecycleConfigResponse(
            string autoDeleteTime,

            string autoDeleteTtl,

            string idleDeleteTtl,

            string idleStartTime)
        {
            AutoDeleteTime = autoDeleteTime;
            AutoDeleteTtl = autoDeleteTtl;
            IdleDeleteTtl = idleDeleteTtl;
            IdleStartTime = idleStartTime;
        }
    }
}
