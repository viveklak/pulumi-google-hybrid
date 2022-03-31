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
    public sealed class DiskResourceStatusResponse
    {
        public readonly Outputs.DiskResourceStatusAsyncReplicationStatusResponse AsyncPrimaryDisk;
        /// <summary>
        /// Key: disk, value: AsyncReplicationStatus message
        /// </summary>
        public readonly ImmutableDictionary<string, string> AsyncSecondaryDisks;

        [OutputConstructor]
        private DiskResourceStatusResponse(
            Outputs.DiskResourceStatusAsyncReplicationStatusResponse asyncPrimaryDisk,

            ImmutableDictionary<string, string> asyncSecondaryDisks)
        {
            AsyncPrimaryDisk = asyncPrimaryDisk;
            AsyncSecondaryDisks = asyncSecondaryDisks;
        }
    }
}
