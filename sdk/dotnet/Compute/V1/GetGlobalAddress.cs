// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetGlobalAddress
    {
        /// <summary>
        /// Returns the specified address resource. Gets a list of available addresses by making a list() request.
        /// </summary>
        public static Task<GetGlobalAddressResult> InvokeAsync(GetGlobalAddressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGlobalAddressResult>("google-native:compute/v1:getGlobalAddress", args ?? new GetGlobalAddressArgs(), options.WithVersion());
    }


    public sealed class GetGlobalAddressArgs : Pulumi.InvokeArgs
    {
        [Input("address", required: true)]
        public string Address { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetGlobalAddressArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGlobalAddressResult
    {
        /// <summary>
        /// The static IP address represented by this resource.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
        /// </summary>
        public readonly string AddressType;
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
        /// </summary>
        public readonly string IpVersion;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#address for addresses.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Global forwarding rules can only be Premium Tier. Regional forwarding rules can be either Premium or Standard Tier. Standard Tier addresses applied to regional forwarding rules can be used with any external load balancer. Regional forwarding rules in Premium Tier can only be used with a network load balancer. If this field is not specified, it is assumed to be PREMIUM.
        /// </summary>
        public readonly string NetworkTier;
        /// <summary>
        /// The prefix length if the resource represents an IP range.
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// The purpose of this resource, which can be one of the following values: - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources. - `DNS_RESOLVER` for a DNS resolver address in a subnetwork - `VPC_PEERING` for addresses that are reserved for VPC peer networks. - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT. - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an *IPsec-encrypted Cloud Interconnect* configuration. These addresses are regional resources. Not currently available publicly. 
        /// </summary>
        public readonly string Purpose;
        /// <summary>
        /// [Output Only] The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. *This field is not applicable to global addresses.*
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// [Output Only] The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// [Output Only] The URLs of the resources that are using this address.
        /// </summary>
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetGlobalAddressResult(
            string address,

            string addressType,

            string creationTimestamp,

            string description,

            string ipVersion,

            string kind,

            string name,

            string network,

            string networkTier,

            int prefixLength,

            string purpose,

            string region,

            string selfLink,

            string status,

            string subnetwork,

            ImmutableArray<string> users)
        {
            Address = address;
            AddressType = addressType;
            CreationTimestamp = creationTimestamp;
            Description = description;
            IpVersion = ipVersion;
            Kind = kind;
            Name = name;
            Network = network;
            NetworkTier = networkTier;
            PrefixLength = prefixLength;
            Purpose = purpose;
            Region = region;
            SelfLink = selfLink;
            Status = status;
            Subnetwork = subnetwork;
            Users = users;
        }
    }
}
