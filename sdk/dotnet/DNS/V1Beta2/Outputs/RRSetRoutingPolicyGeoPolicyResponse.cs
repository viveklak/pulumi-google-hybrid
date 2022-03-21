// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1Beta2.Outputs
{

    /// <summary>
    /// Configures a RRSetRoutingPolicy that routes based on the geo location of the querying user.
    /// </summary>
    [OutputType]
    public sealed class RRSetRoutingPolicyGeoPolicyResponse
    {
        /// <summary>
        /// The primary geo routing configuration. If there are multiple items with the same location, an error is returned instead.
        /// </summary>
        public readonly ImmutableArray<Outputs.RRSetRoutingPolicyGeoPolicyGeoPolicyItemResponse> Items;
        public readonly string Kind;

        [OutputConstructor]
        private RRSetRoutingPolicyGeoPolicyResponse(
            ImmutableArray<Outputs.RRSetRoutingPolicyGeoPolicyGeoPolicyItemResponse> items,

            string kind)
        {
            Items = items;
            Kind = kind;
        }
    }
}
