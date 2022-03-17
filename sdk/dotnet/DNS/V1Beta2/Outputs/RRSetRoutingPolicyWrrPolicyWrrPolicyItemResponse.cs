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
    /// A routing block which contains the routing information for one WRR item.
    /// </summary>
    [OutputType]
    public sealed class RRSetRoutingPolicyWrrPolicyWrrPolicyItemResponse
    {
        public readonly string Kind;
        public readonly ImmutableArray<string> Rrdatas;
        /// <summary>
        /// DNSSEC generated signatures for all the rrdata within this item. Note that if health checked targets are provided for DNSSEC enabled zones, there's a restriction of 1 ip per item. .
        /// </summary>
        public readonly ImmutableArray<string> SignatureRrdatas;
        /// <summary>
        /// The weight corresponding to this subset of rrdata. When multiple WeightedRoundRobinPolicyItems are configured, the probability of returning an rrset is proportional to its weight relative to the sum of weights configured for all items. This weight should be non-negative.
        /// </summary>
        public readonly double Weight;

        [OutputConstructor]
        private RRSetRoutingPolicyWrrPolicyWrrPolicyItemResponse(
            string kind,

            ImmutableArray<string> rrdatas,

            ImmutableArray<string> signatureRrdatas,

            double weight)
        {
            Kind = kind;
            Rrdatas = rrdatas;
            SignatureRrdatas = signatureRrdatas;
            Weight = weight;
        }
    }
}
