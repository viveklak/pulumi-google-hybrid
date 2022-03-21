// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Outputs
{

    /// <summary>
    /// Configures a RRSetRoutingPolicy that routes in a weighted round robin fashion.
    /// </summary>
    [OutputType]
    public sealed class RRSetRoutingPolicyWrrPolicyResponse
    {
        public readonly ImmutableArray<Outputs.RRSetRoutingPolicyWrrPolicyWrrPolicyItemResponse> Items;
        public readonly string Kind;

        [OutputConstructor]
        private RRSetRoutingPolicyWrrPolicyResponse(
            ImmutableArray<Outputs.RRSetRoutingPolicyWrrPolicyWrrPolicyItemResponse> items,

            string kind)
        {
            Items = items;
            Kind = kind;
        }
    }
}
