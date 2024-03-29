// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Outputs
{

    /// <summary>
    /// Determines whether or no a connection is locked. If locked, a reason must be specified.
    /// </summary>
    [OutputType]
    public sealed class LockConfigResponse
    {
        /// <summary>
        /// Indicates whether or not the connection is locked.
        /// </summary>
        public readonly bool Locked;
        /// <summary>
        /// Describes why a connection is locked.
        /// </summary>
        public readonly string Reason;

        [OutputConstructor]
        private LockConfigResponse(
            bool locked,

            string reason)
        {
            Locked = locked;
            Reason = reason;
        }
    }
}
