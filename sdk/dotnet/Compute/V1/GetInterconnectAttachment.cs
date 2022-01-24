// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetInterconnectAttachment
    {
        /// <summary>
        /// Returns the specified interconnect attachment.
        /// </summary>
        public static Task<GetInterconnectAttachmentResult> InvokeAsync(GetInterconnectAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInterconnectAttachmentResult>("google-native:compute/v1:getInterconnectAttachment", args ?? new GetInterconnectAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified interconnect attachment.
        /// </summary>
        public static Output<GetInterconnectAttachmentResult> Invoke(GetInterconnectAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInterconnectAttachmentResult>("google-native:compute/v1:getInterconnectAttachment", args ?? new GetInterconnectAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInterconnectAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("interconnectAttachment", required: true)]
        public string InterconnectAttachment { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetInterconnectAttachmentArgs()
        {
        }
    }

    public sealed class GetInterconnectAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("interconnectAttachment", required: true)]
        public Input<string> InterconnectAttachment { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetInterconnectAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInterconnectAttachmentResult
    {
        /// <summary>
        /// Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
        /// </summary>
        public readonly bool AdminEnabled;
        /// <summary>
        /// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: - BPS_50M: 50 Mbit/s - BPS_100M: 100 Mbit/s - BPS_200M: 200 Mbit/s - BPS_300M: 300 Mbit/s - BPS_400M: 400 Mbit/s - BPS_500M: 500 Mbit/s - BPS_1G: 1 Gbit/s - BPS_2G: 2 Gbit/s - BPS_5G: 5 Gbit/s - BPS_10G: 10 Gbit/s - BPS_20G: 20 Gbit/s - BPS_50G: 50 Gbit/s 
        /// </summary>
        public readonly string Bandwidth;
        /// <summary>
        /// Up to 16 candidate prefixes that control the allocation of cloudRouterIpv6Address and customerRouterIpv6Address for this attachment. Each prefix must be in the Global Unique Address (GUA) space. It is highly recommended that it be in a range owned by the requestor. A GUA in a range owned by Google will cause the request to fail. Google will select an available prefix from the supplied candidates or fail the request. If not supplied, a /125 from a Google-owned GUA block will be selected.
        /// </summary>
        public readonly ImmutableArray<string> CandidateIpv6Subnets;
        /// <summary>
        /// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        /// </summary>
        public readonly ImmutableArray<string> CandidateSubnets;
        /// <summary>
        /// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        /// </summary>
        public readonly string CloudRouterIpAddress;
        /// <summary>
        /// IPv6 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        /// </summary>
        public readonly string CloudRouterIpv6Address;
        /// <summary>
        /// If supplied, the interface id (index within the subnet) to be used for the cloud router address. The id must be in the range of 1 to 6. If a subnet mask is supplied, it must be /125, and the subnet should either be 0 or match the selected subnet.
        /// </summary>
        public readonly string CloudRouterIpv6InterfaceId;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
        /// </summary>
        public readonly string CustomerRouterIpAddress;
        /// <summary>
        /// IPv6 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
        /// </summary>
        public readonly string CustomerRouterIpv6Address;
        /// <summary>
        /// If supplied, the interface id (index within the subnet) to be used for the customer router address. The id must be in the range of 1 to 6. If a subnet mask is supplied, it must be /125, and the subnet should either be 0 or match the selected subnet.
        /// </summary>
        public readonly string CustomerRouterIpv6InterfaceId;
        /// <summary>
        /// [Output only for types PARTNER and DEDICATED. Not present for PARTNER_PROVIDER.] Dataplane version for this InterconnectAttachment. This field is only present for Dataplane version 2 and higher. Absence of this field in the API output indicates that the Dataplane is version 1.
        /// </summary>
        public readonly int DataplaneVersion;
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: - AVAILABILITY_DOMAIN_ANY - AVAILABILITY_DOMAIN_1 - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
        /// </summary>
        public readonly string EdgeAvailabilityDomain;
        /// <summary>
        /// Indicates the user-supplied encryption option of this VLAN attachment (interconnectAttachment). Can only be specified at attachment creation for PARTNER or DEDICATED attachments. Possible values are: - NONE - This is the default value, which means that the VLAN attachment carries unencrypted traffic. VMs are able to send traffic to, or receive traffic from, such a VLAN attachment. - IPSEC - The VLAN attachment carries only encrypted traffic that is encrypted by an IPsec device, such as an HA VPN gateway or third-party IPsec VPN. VMs cannot directly send traffic to, or receive traffic from, such a VLAN attachment. To use *IPsec-encrypted Cloud Interconnect*, the VLAN attachment must be created with this option. Not currently available publicly. 
        /// </summary>
        public readonly string Encryption;
        /// <summary>
        /// URL of the underlying Interconnect object that this attachment's traffic will traverse through.
        /// </summary>
        public readonly string Interconnect;
        /// <summary>
        /// A list of URLs of addresses that have been reserved for the VLAN attachment. Used only for the VLAN attachment that has the encryption option as IPSEC. The addresses must be regional internal IP address ranges. When creating an HA VPN gateway over the VLAN attachment, if the attachment is configured to use a regional internal IP address, then the VPN gateway's IP address is allocated from the IP address range specified here. For example, if the HA VPN gateway's interface 0 is paired to this VLAN attachment, then a regional internal IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this VLAN attachment. If this field is not specified when creating the VLAN attachment, then later on when creating an HA VPN gateway on this VLAN attachment, the HA VPN gateway's IP address is allocated from the regional external IP address pool. Not currently available publicly. 
        /// </summary>
        public readonly ImmutableArray<string> IpsecInternalAddresses;
        /// <summary>
        /// Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
        /// </summary>
        public readonly int Mtu;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current status of whether or not this interconnect attachment is functional, which can take one of the following values: - OS_ACTIVE: The attachment has been turned up and is ready to use. - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. 
        /// </summary>
        public readonly string OperationalStatus;
        /// <summary>
        /// [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        /// </summary>
        public readonly string PairingKey;
        /// <summary>
        /// Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
        /// </summary>
        public readonly string PartnerAsn;
        /// <summary>
        /// Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
        /// </summary>
        public readonly Outputs.InterconnectAttachmentPartnerMetadataResponse PartnerMetadata;
        /// <summary>
        /// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
        /// </summary>
        public readonly Outputs.InterconnectAttachmentPrivateInfoResponse PrivateInterconnectInfo;
        /// <summary>
        /// URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network &amp; region within which the Cloud Router is configured.
        /// </summary>
        public readonly string Router;
        /// <summary>
        /// Set to true if the resource satisfies the zone separation organization policy constraints and false otherwise. Defaults to false if the field is not present.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The stack type for this interconnect attachment to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used. This field can be both set at interconnect attachments creation and update interconnect attachment operations.
        /// </summary>
        public readonly string StackType;
        /// <summary>
        /// The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: - ACTIVE: The attachment has been turned up and is ready to use. - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted. 
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of interconnect attachment this is, which can take one of the following values: - DEDICATED: an attachment to a Dedicated Interconnect. - PARTNER: an attachment to a Partner Interconnect, created by the customer. - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner. 
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. Only specified at creation time.
        /// </summary>
        public readonly int VlanTag8021q;

        [OutputConstructor]
        private GetInterconnectAttachmentResult(
            bool adminEnabled,

            string bandwidth,

            ImmutableArray<string> candidateIpv6Subnets,

            ImmutableArray<string> candidateSubnets,

            string cloudRouterIpAddress,

            string cloudRouterIpv6Address,

            string cloudRouterIpv6InterfaceId,

            string creationTimestamp,

            string customerRouterIpAddress,

            string customerRouterIpv6Address,

            string customerRouterIpv6InterfaceId,

            int dataplaneVersion,

            string description,

            string edgeAvailabilityDomain,

            string encryption,

            string interconnect,

            ImmutableArray<string> ipsecInternalAddresses,

            string kind,

            int mtu,

            string name,

            string operationalStatus,

            string pairingKey,

            string partnerAsn,

            Outputs.InterconnectAttachmentPartnerMetadataResponse partnerMetadata,

            Outputs.InterconnectAttachmentPrivateInfoResponse privateInterconnectInfo,

            string region,

            string router,

            bool satisfiesPzs,

            string selfLink,

            string stackType,

            string state,

            string type,

            int vlanTag8021q)
        {
            AdminEnabled = adminEnabled;
            Bandwidth = bandwidth;
            CandidateIpv6Subnets = candidateIpv6Subnets;
            CandidateSubnets = candidateSubnets;
            CloudRouterIpAddress = cloudRouterIpAddress;
            CloudRouterIpv6Address = cloudRouterIpv6Address;
            CloudRouterIpv6InterfaceId = cloudRouterIpv6InterfaceId;
            CreationTimestamp = creationTimestamp;
            CustomerRouterIpAddress = customerRouterIpAddress;
            CustomerRouterIpv6Address = customerRouterIpv6Address;
            CustomerRouterIpv6InterfaceId = customerRouterIpv6InterfaceId;
            DataplaneVersion = dataplaneVersion;
            Description = description;
            EdgeAvailabilityDomain = edgeAvailabilityDomain;
            Encryption = encryption;
            Interconnect = interconnect;
            IpsecInternalAddresses = ipsecInternalAddresses;
            Kind = kind;
            Mtu = mtu;
            Name = name;
            OperationalStatus = operationalStatus;
            PairingKey = pairingKey;
            PartnerAsn = partnerAsn;
            PartnerMetadata = partnerMetadata;
            PrivateInterconnectInfo = privateInterconnectInfo;
            Region = region;
            Router = router;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            StackType = stackType;
            State = state;
            Type = type;
            VlanTag8021q = vlanTag8021q;
        }
    }
}
