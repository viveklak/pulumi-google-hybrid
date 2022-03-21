// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetAddress
    {
        /// <summary>
        /// Returns the specified address resource.
        /// </summary>
        public static Task<GetAddressResult> InvokeAsync(GetAddressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAddressResult>("google-native:compute/alpha:getAddress", args ?? new GetAddressArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified address resource.
        /// </summary>
        public static Output<GetAddressResult> Invoke(GetAddressInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAddressResult>("google-native:compute/alpha:getAddress", args ?? new GetAddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddressArgs : Pulumi.InvokeArgs
    {
        [Input("address", required: true)]
        public string Address { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetAddressArgs()
        {
        }
    }

    public sealed class GetAddressInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetAddressInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAddressResult
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
        /// Creation timestamp in RFC3339 text format.
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
        /// The endpoint type of this address, which should be VM. This is used for deciding which endpoint this address will be assigned to during the IPv6 external IP address reservation.
        /// </summary>
        public readonly string Ipv6EndpointType;
        /// <summary>
        /// Type of the resource. Always compute#address for addresses.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Address.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
        /// </summary>
        public readonly string NetworkTier;
        /// <summary>
        /// The prefix length if the resource represents an IP range.
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *IPsec-encrypted Cloud Interconnect* configuration. These addresses are regional resources. Not currently available publicly. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose. 
        /// </summary>
        public readonly string Purpose;
        /// <summary>
        /// The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. *This field is not applicable to global addresses.*
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// The URLs of the resources that are using this address.
        /// </summary>
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetAddressResult(
            string address,

            string addressType,

            string creationTimestamp,

            string description,

            string ipVersion,

            string ipv6EndpointType,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            string networkTier,

            int prefixLength,

            string purpose,

            string region,

            string selfLink,

            string selfLinkWithId,

            string status,

            string subnetwork,

            ImmutableArray<string> users)
        {
            Address = address;
            AddressType = addressType;
            CreationTimestamp = creationTimestamp;
            Description = description;
            IpVersion = ipVersion;
            Ipv6EndpointType = ipv6EndpointType;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            Name = name;
            Network = network;
            NetworkTier = networkTier;
            PrefixLength = prefixLength;
            Purpose = purpose;
            Region = region;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            Status = status;
            Subnetwork = subnetwork;
            Users = users;
        }
    }
}
