// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Domains.V1Beta1.Outputs
{

    /// <summary>
    /// Defines a host on your domain that is a DNS name server for your domain and/or other domains. Glue records are a way of making the IP address of a name server known, even when it serves DNS queries for its parent domain. For example, when `ns.example.com` is a name server for `example.com`, the host `ns.example.com` must have a glue record to break the circular DNS reference.
    /// </summary>
    [OutputType]
    public sealed class GlueRecordResponse
    {
        /// <summary>
        /// Domain name of the host in Punycode format.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// List of IPv4 addresses corresponding to this host in the standard decimal format (e.g. `198.51.100.1`). At least one of `ipv4_address` and `ipv6_address` must be set.
        /// </summary>
        public readonly ImmutableArray<string> Ipv4Addresses;
        /// <summary>
        /// List of IPv6 addresses corresponding to this host in the standard hexadecimal format (e.g. `2001:db8::`). At least one of `ipv4_address` and `ipv6_address` must be set.
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Addresses;

        [OutputConstructor]
        private GlueRecordResponse(
            string hostName,

            ImmutableArray<string> ipv4Addresses,

            ImmutableArray<string> ipv6Addresses)
        {
            HostName = hostName;
            Ipv4Addresses = ipv4Addresses;
            Ipv6Addresses = ipv6Addresses;
        }
    }
}
