// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V2.Outputs
{

    [OutputType]
    public sealed class ManagedZoneForwardingConfigNameServerTargetResponse
    {
        /// <summary>
        /// Forwarding path for this NameServerTarget. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on IP address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        /// </summary>
        public readonly string ForwardingPath;
        /// <summary>
        /// IPv4 address of a target name server.
        /// </summary>
        public readonly string Ipv4Address;
        public readonly string Kind;

        [OutputConstructor]
        private ManagedZoneForwardingConfigNameServerTargetResponse(
            string forwardingPath,

            string ipv4Address,

            string kind)
        {
            ForwardingPath = forwardingPath;
            Ipv4Address = ipv4Address;
            Kind = kind;
        }
    }
}
