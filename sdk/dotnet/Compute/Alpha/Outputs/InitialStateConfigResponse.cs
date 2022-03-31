// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Initial State for shielded instance, these are public keys which are safe to store in public
    /// </summary>
    [OutputType]
    public sealed class InitialStateConfigResponse
    {
        /// <summary>
        /// The Key Database (db).
        /// </summary>
        public readonly ImmutableArray<Outputs.FileContentBufferResponse> Dbs;
        /// <summary>
        /// The forbidden key database (dbx).
        /// </summary>
        public readonly ImmutableArray<Outputs.FileContentBufferResponse> Dbxs;
        /// <summary>
        /// The Key Exchange Key (KEK).
        /// </summary>
        public readonly ImmutableArray<Outputs.FileContentBufferResponse> Keks;
        /// <summary>
        /// The Platform Key (PK).
        /// </summary>
        public readonly Outputs.FileContentBufferResponse Pk;

        [OutputConstructor]
        private InitialStateConfigResponse(
            ImmutableArray<Outputs.FileContentBufferResponse> dbs,

            ImmutableArray<Outputs.FileContentBufferResponse> dbxs,

            ImmutableArray<Outputs.FileContentBufferResponse> keks,

            Outputs.FileContentBufferResponse pk)
        {
            Dbs = dbs;
            Dbxs = dbxs;
            Keks = keks;
            Pk = pk;
        }
    }
}
