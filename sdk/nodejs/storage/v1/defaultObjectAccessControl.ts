// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new default object ACL entry on the specified bucket.
 */
export class DefaultObjectAccessControl extends pulumi.CustomResource {
    /**
     * Get an existing DefaultObjectAccessControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DefaultObjectAccessControl {
        return new DefaultObjectAccessControl(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storage/v1:DefaultObjectAccessControl';

    /**
     * Returns true if the given object is an instance of DefaultObjectAccessControl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultObjectAccessControl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultObjectAccessControl.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The domain associated with the entity, if any.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The email address associated with the entity, if any.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The entity holding the permission, in one of the following forms: 
     * - user-userId 
     * - user-email 
     * - group-groupId 
     * - group-email 
     * - domain-domain 
     * - project-team-projectId 
     * - allUsers 
     * - allAuthenticatedUsers Examples: 
     * - The user liz@example.com would be user-liz@example.com. 
     * - The group example@googlegroups.com would be group-example@googlegroups.com. 
     * - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
     */
    public readonly entity!: pulumi.Output<string>;
    /**
     * The ID for the entity, if any.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * HTTP 1.1 Entity tag for the access-control entry.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The content generation of the object, if applied to an object.
     */
    public readonly generation!: pulumi.Output<string>;
    /**
     * The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The name of the object, if applied to an object.
     */
    public readonly object!: pulumi.Output<string>;
    /**
     * The project team associated with the entity, if any.
     */
    public readonly projectTeam!: pulumi.Output<outputs.storage.v1.DefaultObjectAccessControlProjectTeamResponse>;
    /**
     * The access permission for the entity.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The link to this access-control entry.
     */
    public readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a DefaultObjectAccessControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultObjectAccessControlArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["entity"] = args ? args.entity : undefined;
            inputs["entityId"] = args ? args.entityId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["generation"] = args ? args.generation : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["object"] = args ? args.object : undefined;
            inputs["projectTeam"] = args ? args.projectTeam : undefined;
            inputs["provisionalUserProject"] = args ? args.provisionalUserProject : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["userProject"] = args ? args.userProject : undefined;
        } else {
            inputs["bucket"] = undefined /*out*/;
            inputs["domain"] = undefined /*out*/;
            inputs["email"] = undefined /*out*/;
            inputs["entity"] = undefined /*out*/;
            inputs["entityId"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["generation"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["object"] = undefined /*out*/;
            inputs["projectTeam"] = undefined /*out*/;
            inputs["role"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DefaultObjectAccessControl.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DefaultObjectAccessControl resource.
 */
export interface DefaultObjectAccessControlArgs {
    /**
     * The name of the bucket.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * The domain associated with the entity, if any.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The email address associated with the entity, if any.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The entity holding the permission, in one of the following forms: 
     * - user-userId 
     * - user-email 
     * - group-groupId 
     * - group-email 
     * - domain-domain 
     * - project-team-projectId 
     * - allUsers 
     * - allAuthenticatedUsers Examples: 
     * - The user liz@example.com would be user-liz@example.com. 
     * - The group example@googlegroups.com would be group-example@googlegroups.com. 
     * - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
     */
    readonly entity?: pulumi.Input<string>;
    /**
     * The ID for the entity, if any.
     */
    readonly entityId?: pulumi.Input<string>;
    /**
     * HTTP 1.1 Entity tag for the access-control entry.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The content generation of the object, if applied to an object.
     */
    readonly generation?: pulumi.Input<string>;
    /**
     * The ID of the access-control entry.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The name of the object, if applied to an object.
     */
    readonly object?: pulumi.Input<string>;
    /**
     * The project team associated with the entity, if any.
     */
    readonly projectTeam?: pulumi.Input<inputs.storage.v1.DefaultObjectAccessControlProjectTeamArgs>;
    readonly provisionalUserProject?: pulumi.Input<string>;
    /**
     * The access permission for the entity.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The link to this access-control entry.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly userProject?: pulumi.Input<string>;
}
