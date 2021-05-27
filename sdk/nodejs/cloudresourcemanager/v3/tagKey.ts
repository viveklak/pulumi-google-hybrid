// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new TagKey. If another request with the same parameters is sent while the original request is in process, the second request will receive an error. A maximum of 300 TagKeys can exist under a parent at any given time.
 */
export class TagKey extends pulumi.CustomResource {
    /**
     * Get an existing TagKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TagKey {
        return new TagKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudresourcemanager/v3:TagKey';

    /**
     * Returns true if the given object is an instance of TagKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagKey.__pulumiType;
    }

    /**
     * Creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. Namespaced name of the TagKey.
     */
    public /*out*/ readonly namespacedName!: pulumi.Output<string>;
    /**
     * Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
     */
    public readonly shortName!: pulumi.Output<string>;
    /**
     * Update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TagKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TagKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["shortName"] = args ? args.shortName : undefined;
            inputs["validateOnly"] = args ? args.validateOnly : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["namespacedName"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["namespacedName"] = undefined /*out*/;
            inputs["parent"] = undefined /*out*/;
            inputs["shortName"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TagKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagKey resource.
 */
export interface TagKeyArgs {
    /**
     * Optional. User-assigned description of the TagKey. Must not exceed 256 characters. Read-write.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Optional. Entity tag which users can pass to prevent race conditions. This field is always set in server responses. See UpdateTagKeyRequest for details.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Immutable. The resource name for a TagKey. Must be in the format `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for the TagKey.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Immutable. The resource name of the new TagKey's parent. Must be of the form `organizations/{org_id}`.
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Required. Immutable. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace. The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
     */
    readonly shortName?: pulumi.Input<string>;
    readonly validateOnly?: pulumi.Input<string>;
}
