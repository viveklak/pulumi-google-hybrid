// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified global PublicDelegatedPrefix resource.
 */
export function getGlobalPublicDelegatedPrefix(args: GetGlobalPublicDelegatedPrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalPublicDelegatedPrefixResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/alpha:getGlobalPublicDelegatedPrefix", {
        "project": args.project,
        "publicDelegatedPrefix": args.publicDelegatedPrefix,
    }, opts);
}

export interface GetGlobalPublicDelegatedPrefixArgs {
    project: string;
    publicDelegatedPrefix: string;
}

export interface GetGlobalPublicDelegatedPrefixResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
     */
    readonly fingerprint: string;
    /**
     * The IPv4 address range, in CIDR format, represented by this public delegated prefix.
     */
    readonly ipCidrRange: string;
    /**
     * If true, the prefix will be live migrated.
     */
    readonly isLiveMigration: boolean;
    /**
     * Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
     */
    readonly parentPrefix: string;
    /**
     * The list of sub public delegated prefixes that exist for this public delegated prefix.
     */
    readonly publicDelegatedSubPrefixs: outputs.compute.alpha.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse[];
    /**
     * URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL with id for the resource.
     */
    readonly selfLinkWithId: string;
    /**
     * The status of the public delegated prefix.
     */
    readonly status: string;
}
