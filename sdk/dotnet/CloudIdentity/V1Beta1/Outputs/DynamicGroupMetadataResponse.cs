// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// Dynamic group metadata like queries and status.
    /// </summary>
    [OutputType]
    public sealed class DynamicGroupMetadataResponse
    {
        /// <summary>
        /// Memberships will be the union of all queries. Only one entry with USER resource is currently supported. Customers can create up to 100 dynamic groups.
        /// </summary>
        public readonly ImmutableArray<Outputs.DynamicGroupQueryResponse> Queries;
        /// <summary>
        /// Status of the dynamic group.
        /// </summary>
        public readonly Outputs.DynamicGroupStatusResponse Status;

        [OutputConstructor]
        private DynamicGroupMetadataResponse(
            ImmutableArray<Outputs.DynamicGroupQueryResponse> queries,

            Outputs.DynamicGroupStatusResponse status)
        {
            Queries = queries;
            Status = status;
        }
    }
}
