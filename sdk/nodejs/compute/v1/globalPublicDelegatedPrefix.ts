// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.
 */
export class GlobalPublicDelegatedPrefix extends pulumi.CustomResource {
    /**
     * Get an existing GlobalPublicDelegatedPrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GlobalPublicDelegatedPrefix {
        return new GlobalPublicDelegatedPrefix(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:GlobalPublicDelegatedPrefix';

    /**
     * Returns true if the given object is an instance of GlobalPublicDelegatedPrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalPublicDelegatedPrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalPublicDelegatedPrefix.__pulumiType;
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
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The IPv4 address range, in CIDR format, represented by this public delegated prefix.
     */
    public readonly ipCidrRange!: pulumi.Output<string>;
    /**
     * If true, the prefix will be live migrated.
     */
    public readonly isLiveMigration!: pulumi.Output<boolean>;
    /**
     * Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
     */
    public readonly parentPrefix!: pulumi.Output<string>;
    /**
     * The list of sub public delegated prefixes that exist for this public delegated prefix.
     */
    public readonly publicDelegatedSubPrefixs!: pulumi.Output<outputs.compute.v1.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse[]>;
    /**
     * URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned. 
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a GlobalPublicDelegatedPrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GlobalPublicDelegatedPrefixArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipCidrRange"] = args ? args.ipCidrRange : undefined;
            resourceInputs["isLiveMigration"] = args ? args.isLiveMigration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentPrefix"] = args ? args.parentPrefix : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["publicDelegatedSubPrefixs"] = args ? args.publicDelegatedSubPrefixs : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["ipCidrRange"] = undefined /*out*/;
            resourceInputs["isLiveMigration"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parentPrefix"] = undefined /*out*/;
            resourceInputs["publicDelegatedSubPrefixs"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GlobalPublicDelegatedPrefix.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GlobalPublicDelegatedPrefix resource.
 */
export interface GlobalPublicDelegatedPrefixArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The IPv4 address range, in CIDR format, represented by this public delegated prefix.
     */
    ipCidrRange?: pulumi.Input<string>;
    /**
     * If true, the prefix will be live migrated.
     */
    isLiveMigration?: pulumi.Input<boolean>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
     */
    parentPrefix?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The list of sub public delegated prefixes that exist for this public delegated prefix.
     */
    publicDelegatedSubPrefixs?: pulumi.Input<pulumi.Input<inputs.compute.v1.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs>[]>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
}
