// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyRuleMatcherConfigResponse
    {
        /// <summary>
        /// CIDR IP address range. Maximum number of src_ip_ranges allowed is 10.
        /// </summary>
        public readonly ImmutableArray<string> SrcIpRanges;

        [OutputConstructor]
        private SecurityPolicyRuleMatcherConfigResponse(ImmutableArray<string> srcIpRanges)
        {
            SrcIpRanges = srcIpRanges;
        }
    }
}
