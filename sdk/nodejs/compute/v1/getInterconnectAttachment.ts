// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified interconnect attachment.
 */
export function getInterconnectAttachment(args: GetInterconnectAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetInterconnectAttachmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/v1:getInterconnectAttachment", {
        "interconnectAttachment": args.interconnectAttachment,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetInterconnectAttachmentArgs {
    interconnectAttachment: string;
    project?: string;
    region: string;
}

export interface GetInterconnectAttachmentResult {
    /**
     * Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
     */
    readonly adminEnabled: boolean;
    /**
     * Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: - BPS_50M: 50 Mbit/s - BPS_100M: 100 Mbit/s - BPS_200M: 200 Mbit/s - BPS_300M: 300 Mbit/s - BPS_400M: 400 Mbit/s - BPS_500M: 500 Mbit/s - BPS_1G: 1 Gbit/s - BPS_2G: 2 Gbit/s - BPS_5G: 5 Gbit/s - BPS_10G: 10 Gbit/s - BPS_20G: 20 Gbit/s - BPS_50G: 50 Gbit/s 
     */
    readonly bandwidth: string;
    /**
     * Up to 16 candidate prefixes that control the allocation of cloudRouterIpv6Address and customerRouterIpv6Address for this attachment. Each prefix must be in the Global Unique Address (GUA) space. It is highly recommended that it be in a range owned by the requestor. A GUA in a range owned by Google will cause the request to fail. Google will select an available prefix from the supplied candidates or fail the request. If not supplied, a /125 from a Google-owned GUA block will be selected.
     */
    readonly candidateIpv6Subnets: string[];
    /**
     * Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
     */
    readonly candidateSubnets: string[];
    /**
     * IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
     */
    readonly cloudRouterIpAddress: string;
    /**
     * IPv6 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
     */
    readonly cloudRouterIpv6Address: string;
    /**
     * If supplied, the interface id (index within the subnet) to be used for the cloud router address. The id must be in the range of 1 to 6. If a subnet mask is supplied, it must be /125, and the subnet should either be 0 or match the selected subnet.
     */
    readonly cloudRouterIpv6InterfaceId: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
     */
    readonly customerRouterIpAddress: string;
    /**
     * IPv6 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
     */
    readonly customerRouterIpv6Address: string;
    /**
     * If supplied, the interface id (index within the subnet) to be used for the customer router address. The id must be in the range of 1 to 6. If a subnet mask is supplied, it must be /125, and the subnet should either be 0 or match the selected subnet.
     */
    readonly customerRouterIpv6InterfaceId: string;
    /**
     * [Output only for types PARTNER and DEDICATED. Not present for PARTNER_PROVIDER.] Dataplane version for this InterconnectAttachment. This field is only present for Dataplane version 2 and higher. Absence of this field in the API output indicates that the Dataplane is version 1.
     */
    readonly dataplaneVersion: number;
    /**
     * An optional description of this resource.
     */
    readonly description: string;
    /**
     * Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: - AVAILABILITY_DOMAIN_ANY - AVAILABILITY_DOMAIN_1 - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
     */
    readonly edgeAvailabilityDomain: string;
    /**
     * Indicates the user-supplied encryption option of this VLAN attachment (interconnectAttachment). Can only be specified at attachment creation for PARTNER or DEDICATED attachments. Possible values are: - NONE - This is the default value, which means that the VLAN attachment carries unencrypted traffic. VMs are able to send traffic to, or receive traffic from, such a VLAN attachment. - IPSEC - The VLAN attachment carries only encrypted traffic that is encrypted by an IPsec device, such as an HA VPN gateway or third-party IPsec VPN. VMs cannot directly send traffic to, or receive traffic from, such a VLAN attachment. To use *IPsec-encrypted Cloud Interconnect*, the VLAN attachment must be created with this option. Not currently available publicly. 
     */
    readonly encryption: string;
    /**
     * URL of the underlying Interconnect object that this attachment's traffic will traverse through.
     */
    readonly interconnect: string;
    /**
     * A list of URLs of addresses that have been reserved for the VLAN attachment. Used only for the VLAN attachment that has the encryption option as IPSEC. The addresses must be regional internal IP address ranges. When creating an HA VPN gateway over the VLAN attachment, if the attachment is configured to use a regional internal IP address, then the VPN gateway's IP address is allocated from the IP address range specified here. For example, if the HA VPN gateway's interface 0 is paired to this VLAN attachment, then a regional internal IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this VLAN attachment. If this field is not specified when creating the VLAN attachment, then later on when creating an HA VPN gateway on this VLAN attachment, the HA VPN gateway's IP address is allocated from the regional external IP address pool. Not currently available publicly. 
     */
    readonly ipsecInternalAddresses: string[];
    /**
     * Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
     */
    readonly kind: string;
    /**
     * Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
     */
    readonly mtu: number;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The current status of whether or not this interconnect attachment is functional, which can take one of the following values: - OS_ACTIVE: The attachment has been turned up and is ready to use. - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. 
     */
    readonly operationalStatus: string;
    /**
     * [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
     */
    readonly pairingKey: string;
    /**
     * Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
     */
    readonly partnerAsn: string;
    /**
     * Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
     */
    readonly partnerMetadata: outputs.compute.v1.InterconnectAttachmentPartnerMetadataResponse;
    /**
     * Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
     */
    readonly privateInterconnectInfo: outputs.compute.v1.InterconnectAttachmentPrivateInfoResponse;
    /**
     * URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region within which the Cloud Router is configured.
     */
    readonly router: string;
    /**
     * Set to true if the resource satisfies the zone separation organization policy constraints and false otherwise. Defaults to false if the field is not present.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * The stack type for this interconnect attachment to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used. This field can be both set at interconnect attachments creation and update interconnect attachment operations.
     */
    readonly stackType: string;
    /**
     * The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: - ACTIVE: The attachment has been turned up and is ready to use. - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted. 
     */
    readonly state: string;
    /**
     * The type of interconnect attachment this is, which can take one of the following values: - DEDICATED: an attachment to a Dedicated Interconnect. - PARTNER: an attachment to a Partner Interconnect, created by the customer. - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner. 
     */
    readonly type: string;
    /**
     * The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. Only specified at creation time.
     */
    readonly vlanTag8021q: number;
}

export function getInterconnectAttachmentOutput(args: GetInterconnectAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInterconnectAttachmentResult> {
    return pulumi.output(args).apply(a => getInterconnectAttachment(a, opts))
}

export interface GetInterconnectAttachmentOutputArgs {
    interconnectAttachment: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
