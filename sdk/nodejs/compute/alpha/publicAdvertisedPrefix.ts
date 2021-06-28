// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a PublicAdvertisedPrefix in the specified project using the parameters that are included in the request.
 */
export class PublicAdvertisedPrefix extends pulumi.CustomResource {
    /**
     * Get an existing PublicAdvertisedPrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PublicAdvertisedPrefix {
        return new PublicAdvertisedPrefix(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:PublicAdvertisedPrefix';

    /**
     * Returns true if the given object is an instance of PublicAdvertisedPrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicAdvertisedPrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicAdvertisedPrefix.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The IPv4 address to be used for reverse DNS verification.
     */
    public readonly dnsVerificationIp!: pulumi.Output<string>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicAdvertisedPrefix. An up-to-date fingerprint must be provided in order to update the PublicAdvertisedPrefix, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a PublicAdvertisedPrefix.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The IPv4 address range, in CIDR format, represented by this public advertised prefix.
     */
    public readonly ipCidrRange!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#publicAdvertisedPrefix for public advertised prefixes.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of public delegated prefixes that exist for this public advertised prefix.
     */
    public /*out*/ readonly publicDelegatedPrefixs!: pulumi.Output<outputs.compute.alpha.PublicAdvertisedPrefixPublicDelegatedPrefixResponse[]>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL with id for the resource.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * The shared secret to be used for reverse DNS verification.
     */
    public /*out*/ readonly sharedSecret!: pulumi.Output<string>;
    /**
     * The status of the public advertised prefix.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a PublicAdvertisedPrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicAdvertisedPrefixArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["dnsVerificationIp"] = args ? args.dnsVerificationIp : undefined;
            inputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["publicDelegatedPrefixs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["sharedSecret"] = undefined /*out*/;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["dnsVerificationIp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["ipCidrRange"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["publicDelegatedPrefixs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["sharedSecret"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(PublicAdvertisedPrefix.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PublicAdvertisedPrefix resource.
 */
export interface PublicAdvertisedPrefixArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The IPv4 address to be used for reverse DNS verification.
     */
    dnsVerificationIp?: pulumi.Input<string>;
    /**
     * The IPv4 address range, in CIDR format, represented by this public advertised prefix.
     */
    ipCidrRange?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The status of the public advertised prefix.
     */
    status?: pulumi.Input<enums.compute.alpha.PublicAdvertisedPrefixStatus>;
}
