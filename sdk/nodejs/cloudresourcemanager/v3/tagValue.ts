// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a TagValue as a child of the specified TagKey. If a another request with the same parameters is sent while the original request is in process the second request will receive an error. A maximum of 300 TagValues can exist under a TagKey at any given time.
 */
export class TagValue extends pulumi.CustomResource {
    /**
     * Get an existing TagValue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TagValue {
        return new TagValue(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudresourcemanager/v3:TagValue';

    /**
     * Returns true if the given object is an instance of TagValue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagValue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagValue.__pulumiType;
    }

    /**
     * Creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Immutable. Resource name for TagValue in the format `tagValues/456`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Namespaced name of the TagValue. Must be in the format `{organization_id}/{tag_key_short_name}/{short_name}`.
     */
    public /*out*/ readonly namespacedName!: pulumi.Output<string>;
    /**
     * Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
     */
    public readonly shortName!: pulumi.Output<string>;
    /**
     * Update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TagValue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagValueArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.shortName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shortName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["shortName"] = args ? args.shortName : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["namespacedName"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespacedName"] = undefined /*out*/;
            resourceInputs["parent"] = undefined /*out*/;
            resourceInputs["shortName"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TagValue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagValue resource.
 */
export interface TagValueArgs {
    /**
     * Optional. User-assigned description of the TagValue. Must not exceed 256 characters. Read-write.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagValueRequest for details.
     */
    etag?: pulumi.Input<string>;
    /**
     * Immutable. Resource name for TagValue in the format `tagValues/456`.
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. The resource name of the new TagValue's parent TagKey. Must be of the form `tagKeys/{tag_key_id}`.
     */
    parent?: pulumi.Input<string>;
    /**
     * Immutable. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey. The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
     */
    shortName: pulumi.Input<string>;
    validateOnly?: pulumi.Input<string>;
}
