// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates an InterconnectAttachment in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:InterconnectAttachment")]
    public partial class InterconnectAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
        /// </summary>
        [Output("adminEnabled")]
        public Output<bool> AdminEnabled { get; private set; } = null!;

        /// <summary>
        /// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: 
        /// - BPS_50M: 50 Mbit/s 
        /// - BPS_100M: 100 Mbit/s 
        /// - BPS_200M: 200 Mbit/s 
        /// - BPS_300M: 300 Mbit/s 
        /// - BPS_400M: 400 Mbit/s 
        /// - BPS_500M: 500 Mbit/s 
        /// - BPS_1G: 1 Gbit/s 
        /// - BPS_2G: 2 Gbit/s 
        /// - BPS_5G: 5 Gbit/s 
        /// - BPS_10G: 10 Gbit/s 
        /// - BPS_20G: 20 Gbit/s 
        /// - BPS_50G: 50 Gbit/s
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        /// </summary>
        [Output("candidateSubnets")]
        public Output<ImmutableArray<string>> CandidateSubnets { get; private set; } = null!;

        /// <summary>
        /// [Output Only] IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        /// </summary>
        [Output("cloudRouterIpAddress")]
        public Output<string> CloudRouterIpAddress { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// [Output Only] IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
        /// </summary>
        [Output("customerRouterIpAddress")]
        public Output<string> CustomerRouterIpAddress { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Dataplane version for this InterconnectAttachment.
        /// </summary>
        [Output("dataplaneVersion")]
        public Output<int> DataplaneVersion { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: 
        /// - AVAILABILITY_DOMAIN_ANY 
        /// - AVAILABILITY_DOMAIN_1 
        /// - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
        /// </summary>
        [Output("edgeAvailabilityDomain")]
        public Output<string> EdgeAvailabilityDomain { get; private set; } = null!;

        /// <summary>
        /// Indicates the user-supplied encryption option of this interconnect attachment: 
        /// - NONE is the default value, which means that the attachment carries unencrypted traffic. VMs can send traffic to, or receive traffic from, this type of attachment. 
        /// - IPSEC indicates that the attachment carries only traffic encrypted by an IPsec device such as an HA VPN gateway. VMs cannot directly send traffic to, or receive traffic from, such an attachment. To use IPsec-encrypted Cloud Interconnect, create the attachment using this option. 
        /// Not currently available in all Interconnect locations.
        /// </summary>
        [Output("encryption")]
        public Output<string> Encryption { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues. [Deprecated] This field is not used.
        /// </summary>
        [Output("googleReferenceId")]
        public Output<string> GoogleReferenceId { get; private set; } = null!;

        /// <summary>
        /// URL of the underlying Interconnect object that this attachment's traffic will traverse through.
        /// </summary>
        [Output("interconnect")]
        public Output<string> Interconnect { get; private set; } = null!;

        /// <summary>
        /// URL of addresses that have been reserved for the interconnect attachment, Used only for interconnect attachment that has the encryption option as IPSEC. The addresses must be RFC 1918 IP address ranges. When creating HA VPN gateway over the interconnect attachment, if the attachment is configured to use an RFC 1918 IP address, then the VPN gateway?s IP address will be allocated from the IP address range specified here. For example, if the HA VPN gateway?s interface 0 is paired to this interconnect attachment, then an RFC 1918 IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this interconnect attachment. If this field is not specified for interconnect attachment that has encryption option as IPSEC, later on when creating HA VPN gateway on this interconnect attachment, the HA VPN gateway's IP address will be allocated from regional external IP address pool.
        /// Not currently available in all Interconnect locations.
        /// </summary>
        [Output("ipsecInternalAddresses")]
        public Output<ImmutableArray<string>> IpsecInternalAddresses { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this InterconnectAttachment, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve an InterconnectAttachment.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The current status of whether or not this interconnect attachment is functional, which can take one of the following values: 
        /// - OS_ACTIVE: The attachment has been turned up and is ready to use. 
        /// - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete.
        /// </summary>
        [Output("operationalStatus")]
        public Output<string> OperationalStatus { get; private set; } = null!;

        /// <summary>
        /// [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        /// </summary>
        [Output("pairingKey")]
        public Output<string> PairingKey { get; private set; } = null!;

        /// <summary>
        /// Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
        /// </summary>
        [Output("partnerAsn")]
        public Output<string> PartnerAsn { get; private set; } = null!;

        /// <summary>
        /// Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
        /// </summary>
        [Output("partnerMetadata")]
        public Output<Outputs.InterconnectAttachmentPartnerMetadataResponse> PartnerMetadata { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
        /// </summary>
        [Output("privateInterconnectInfo")]
        public Output<Outputs.InterconnectAttachmentPrivateInfoResponse> PrivateInterconnectInfo { get; private set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network &amp; region within which the Cloud Router is configured.
        /// </summary>
        [Output("router")]
        public Output<string> Router { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: 
        /// - ACTIVE: The attachment has been turned up and is ready to use. 
        /// - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. 
        /// - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. 
        /// - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. 
        /// - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. 
        /// - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of interconnect attachment this is, which can take one of the following values: 
        /// - DEDICATED: an attachment to a Dedicated Interconnect. 
        /// - PARTNER: an attachment to a Partner Interconnect, created by the customer. 
        /// - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. Only specified at creation time.
        /// </summary>
        [Output("vlanTag8021q")]
        public Output<int> VlanTag8021q { get; private set; } = null!;


        /// <summary>
        /// Create a InterconnectAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterconnectAttachment(string name, InterconnectAttachmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:InterconnectAttachment", name, args ?? new InterconnectAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterconnectAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:InterconnectAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InterconnectAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterconnectAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InterconnectAttachment(name, id, options);
        }
    }

    public sealed class InterconnectAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
        /// </summary>
        [Input("adminEnabled")]
        public Input<bool>? AdminEnabled { get; set; }

        /// <summary>
        /// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: 
        /// - BPS_50M: 50 Mbit/s 
        /// - BPS_100M: 100 Mbit/s 
        /// - BPS_200M: 200 Mbit/s 
        /// - BPS_300M: 300 Mbit/s 
        /// - BPS_400M: 400 Mbit/s 
        /// - BPS_500M: 500 Mbit/s 
        /// - BPS_1G: 1 Gbit/s 
        /// - BPS_2G: 2 Gbit/s 
        /// - BPS_5G: 5 Gbit/s 
        /// - BPS_10G: 10 Gbit/s 
        /// - BPS_20G: 20 Gbit/s 
        /// - BPS_50G: 50 Gbit/s
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        [Input("candidateSubnets")]
        private InputList<string>? _candidateSubnets;

        /// <summary>
        /// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        /// </summary>
        public InputList<string> CandidateSubnets
        {
            get => _candidateSubnets ?? (_candidateSubnets = new InputList<string>());
            set => _candidateSubnets = value;
        }

        /// <summary>
        /// [Output Only] IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        /// </summary>
        [Input("cloudRouterIpAddress")]
        public Input<string>? CloudRouterIpAddress { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// [Output Only] IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
        /// </summary>
        [Input("customerRouterIpAddress")]
        public Input<string>? CustomerRouterIpAddress { get; set; }

        /// <summary>
        /// [Output Only] Dataplane version for this InterconnectAttachment.
        /// </summary>
        [Input("dataplaneVersion")]
        public Input<int>? DataplaneVersion { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: 
        /// - AVAILABILITY_DOMAIN_ANY 
        /// - AVAILABILITY_DOMAIN_1 
        /// - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
        /// </summary>
        [Input("edgeAvailabilityDomain")]
        public Input<string>? EdgeAvailabilityDomain { get; set; }

        /// <summary>
        /// Indicates the user-supplied encryption option of this interconnect attachment: 
        /// - NONE is the default value, which means that the attachment carries unencrypted traffic. VMs can send traffic to, or receive traffic from, this type of attachment. 
        /// - IPSEC indicates that the attachment carries only traffic encrypted by an IPsec device such as an HA VPN gateway. VMs cannot directly send traffic to, or receive traffic from, such an attachment. To use IPsec-encrypted Cloud Interconnect, create the attachment using this option. 
        /// Not currently available in all Interconnect locations.
        /// </summary>
        [Input("encryption")]
        public Input<string>? Encryption { get; set; }

        /// <summary>
        /// [Output Only] Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues. [Deprecated] This field is not used.
        /// </summary>
        [Input("googleReferenceId")]
        public Input<string>? GoogleReferenceId { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// URL of the underlying Interconnect object that this attachment's traffic will traverse through.
        /// </summary>
        [Input("interconnect")]
        public Input<string>? Interconnect { get; set; }

        [Input("ipsecInternalAddresses")]
        private InputList<string>? _ipsecInternalAddresses;

        /// <summary>
        /// URL of addresses that have been reserved for the interconnect attachment, Used only for interconnect attachment that has the encryption option as IPSEC. The addresses must be RFC 1918 IP address ranges. When creating HA VPN gateway over the interconnect attachment, if the attachment is configured to use an RFC 1918 IP address, then the VPN gateway?s IP address will be allocated from the IP address range specified here. For example, if the HA VPN gateway?s interface 0 is paired to this interconnect attachment, then an RFC 1918 IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this interconnect attachment. If this field is not specified for interconnect attachment that has encryption option as IPSEC, later on when creating HA VPN gateway on this interconnect attachment, the HA VPN gateway's IP address will be allocated from regional external IP address pool.
        /// Not currently available in all Interconnect locations.
        /// </summary>
        public InputList<string> IpsecInternalAddresses
        {
            get => _ipsecInternalAddresses ?? (_ipsecInternalAddresses = new InputList<string>());
            set => _ipsecInternalAddresses = value;
        }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A fingerprint for the labels being applied to this InterconnectAttachment, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve an InterconnectAttachment.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [Output Only] The current status of whether or not this interconnect attachment is functional, which can take one of the following values: 
        /// - OS_ACTIVE: The attachment has been turned up and is ready to use. 
        /// - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete.
        /// </summary>
        [Input("operationalStatus")]
        public Input<string>? OperationalStatus { get; set; }

        /// <summary>
        /// [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        /// </summary>
        [Input("pairingKey")]
        public Input<string>? PairingKey { get; set; }

        /// <summary>
        /// Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
        /// </summary>
        [Input("partnerAsn")]
        public Input<string>? PartnerAsn { get; set; }

        /// <summary>
        /// Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
        /// </summary>
        [Input("partnerMetadata")]
        public Input<Inputs.InterconnectAttachmentPartnerMetadataArgs>? PartnerMetadata { get; set; }

        /// <summary>
        /// [Output Only] Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
        /// </summary>
        [Input("privateInterconnectInfo")]
        public Input<Inputs.InterconnectAttachmentPrivateInfoArgs>? PrivateInterconnectInfo { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network &amp; region within which the Cloud Router is configured.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

        /// <summary>
        /// [Output Only] The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: 
        /// - ACTIVE: The attachment has been turned up and is ready to use. 
        /// - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. 
        /// - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. 
        /// - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. 
        /// - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. 
        /// - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of interconnect attachment this is, which can take one of the following values: 
        /// - DEDICATED: an attachment to a Dedicated Interconnect. 
        /// - PARTNER: an attachment to a Partner Interconnect, created by the customer. 
        /// - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        /// <summary>
        /// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. Only specified at creation time.
        /// </summary>
        [Input("vlanTag8021q")]
        public Input<int>? VlanTag8021q { get; set; }

        public InterconnectAttachmentArgs()
        {
        }
    }
}
