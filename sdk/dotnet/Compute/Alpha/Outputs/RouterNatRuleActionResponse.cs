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
    public sealed class RouterNatRuleActionResponse
    {
        /// <summary>
        /// A list of URLs of the IP resources used for this NAT rule. These IP addresses must be valid static external IP addresses assigned to the project. This field is used for public NAT.
        /// </summary>
        public readonly ImmutableArray<string> SourceNatActiveIps;
        /// <summary>
        /// A list of URLs of the subnetworks used as source ranges for this NAT Rule. These subnetworks must have purpose set to PRIVATE_NAT. This field is used for private NAT.
        /// </summary>
        public readonly ImmutableArray<string> SourceNatActiveRanges;
        /// <summary>
        /// A list of URLs of the IP resources to be drained. These IPs must be valid static external IPs that have been assigned to the NAT. These IPs should be used for updating/patching a NAT rule only. This field is used for public NAT.
        /// </summary>
        public readonly ImmutableArray<string> SourceNatDrainIps;
        /// <summary>
        /// A list of URLs of subnetworks representing source ranges to be drained. This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule. This field is used for private NAT.
        /// </summary>
        public readonly ImmutableArray<string> SourceNatDrainRanges;

        [OutputConstructor]
        private RouterNatRuleActionResponse(
            ImmutableArray<string> sourceNatActiveIps,

            ImmutableArray<string> sourceNatActiveRanges,

            ImmutableArray<string> sourceNatDrainIps,

            ImmutableArray<string> sourceNatDrainRanges)
        {
            SourceNatActiveIps = sourceNatActiveIps;
            SourceNatActiveRanges = sourceNatActiveRanges;
            SourceNatDrainIps = sourceNatDrainIps;
            SourceNatDrainRanges = sourceNatDrainRanges;
        }
    }
}
