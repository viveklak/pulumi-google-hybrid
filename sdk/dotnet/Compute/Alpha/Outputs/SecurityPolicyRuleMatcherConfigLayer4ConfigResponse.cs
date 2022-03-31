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
    public sealed class SecurityPolicyRuleMatcherConfigLayer4ConfigResponse
    {
        /// <summary>
        /// The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP protocol number.
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ["22"], ["80","443"], and ["12345-12349"]. This field may only be specified when versioned_expr is set to FIREWALL.
        /// </summary>
        public readonly ImmutableArray<string> Ports;

        [OutputConstructor]
        private SecurityPolicyRuleMatcherConfigLayer4ConfigResponse(
            string ipProtocol,

            ImmutableArray<string> ports)
        {
            IpProtocol = ipProtocol;
            Ports = ports;
        }
    }
}
