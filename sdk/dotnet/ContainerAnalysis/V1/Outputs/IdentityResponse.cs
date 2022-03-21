// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// The unique identifier of the update.
    /// </summary>
    [OutputType]
    public sealed class IdentityResponse
    {
        /// <summary>
        /// The revision number of the update.
        /// </summary>
        public readonly int Revision;
        /// <summary>
        /// The revision independent identifier of the update.
        /// </summary>
        public readonly string UpdateId;

        [OutputConstructor]
        private IdentityResponse(
            int revision,

            string updateId)
        {
            Revision = revision;
            UpdateId = updateId;
        }
    }
}
