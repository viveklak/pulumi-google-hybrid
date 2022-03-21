// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified PublicAdvertisedPrefix resource.
 */
export function getPublicAdvertisedPrefix(args: GetPublicAdvertisedPrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicAdvertisedPrefixResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/v1:getPublicAdvertisedPrefix", {
        "project": args.project,
        "publicAdvertisedPrefix": args.publicAdvertisedPrefix,
    }, opts);
}

export interface GetPublicAdvertisedPrefixArgs {
    project?: string;
    publicAdvertisedPrefix: string;
}

export interface GetPublicAdvertisedPrefixResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * The IPv4 address to be used for reverse DNS verification.
     */
    readonly dnsVerificationIp: string;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicAdvertisedPrefix. An up-to-date fingerprint must be provided in order to update the PublicAdvertisedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicAdvertisedPrefix.
     */
    readonly fingerprint: string;
    /**
     * The IPv4 address range, in CIDR format, represented by this public advertised prefix.
     */
    readonly ipCidrRange: string;
    /**
     * Type of the resource. Always compute#publicAdvertisedPrefix for public advertised prefixes.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The list of public delegated prefixes that exist for this public advertised prefix.
     */
    readonly publicDelegatedPrefixs: outputs.compute.v1.PublicAdvertisedPrefixPublicDelegatedPrefixResponse[];
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * The shared secret to be used for reverse DNS verification.
     */
    readonly sharedSecret: string;
    /**
     * The status of the public advertised prefix. Possible values include: - `INITIAL`: RPKI validation is complete. - `PTR_CONFIGURED`: User has configured the PTR. - `VALIDATED`: Reverse DNS lookup is successful. - `REVERSE_DNS_LOOKUP_FAILED`: Reverse DNS lookup failed. - `PREFIX_CONFIGURATION_IN_PROGRESS`: The prefix is being configured. - `PREFIX_CONFIGURATION_COMPLETE`: The prefix is fully configured. - `PREFIX_REMOVAL_IN_PROGRESS`: The prefix is being removed. 
     */
    readonly status: string;
}

export function getPublicAdvertisedPrefixOutput(args: GetPublicAdvertisedPrefixOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicAdvertisedPrefixResult> {
    return pulumi.output(args).apply(a => getPublicAdvertisedPrefix(a, opts))
}

export interface GetPublicAdvertisedPrefixOutputArgs {
    project?: pulumi.Input<string>;
    publicAdvertisedPrefix: pulumi.Input<string>;
}
