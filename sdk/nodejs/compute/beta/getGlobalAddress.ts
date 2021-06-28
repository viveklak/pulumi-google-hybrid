// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified address resource. Gets a list of available addresses by making a list() request.
 */
export function getGlobalAddress(args: GetGlobalAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalAddressResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/beta:getGlobalAddress", {
        "address": args.address,
        "project": args.project,
    }, opts);
}

export interface GetGlobalAddressArgs {
    address: string;
    project: string;
}

export interface GetGlobalAddressResult {
    /**
     * The static IP address represented by this resource.
     */
    readonly address: string;
    /**
     * The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
     */
    readonly addressType: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    readonly description: string;
    /**
     * The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
     */
    readonly ipVersion: string;
    /**
     * Type of the resource. Always compute#address for addresses.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an Address.
     */
    readonly labelFingerprint: string;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    readonly name: string;
    /**
     * The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
     */
    readonly network: string;
    /**
     * This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Global forwarding rules can only be Premium Tier. Regional forwarding rules can be either Premium or Standard Tier. Standard Tier addresses applied to regional forwarding rules can be used with any external load balancer. Regional forwarding rules in Premium Tier can only be used with a network load balancer.
     *
     * If this field is not specified, it is assumed to be PREMIUM.
     */
    readonly networkTier: string;
    /**
     * The prefix length if the resource represents an IP range.
     */
    readonly prefixLength: number;
    /**
     * The purpose of this resource, which can be one of the following values:  
     * - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources. 
     * - `DNS_RESOLVER` for a DNS resolver address in a subnetwork 
     * - `VPC_PEERING` for addresses that are reserved for VPC peer networks. 
     * - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT. 
     * - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration. These addresses are regional resources.
     */
    readonly purpose: string;
    /**
     * The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. This field is not applicable to global addresses.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
     */
    readonly status: string;
    /**
     * The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
     */
    readonly subnetwork: string;
    /**
     * The URLs of the resources that are using this address.
     */
    readonly users: string[];
}
