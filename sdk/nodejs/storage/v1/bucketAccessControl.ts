// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new ACL entry on the specified bucket.
 * Auto-naming is currently not supported for this resource.
 */
export class BucketAccessControl extends pulumi.CustomResource {
    /**
     * Get an existing BucketAccessControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BucketAccessControl {
        return new BucketAccessControl(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storage/v1:BucketAccessControl';

    /**
     * Returns true if the given object is an instance of BucketAccessControl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketAccessControl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketAccessControl.__pulumiType;
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
     * The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The project team associated with the entity, if any.
     */
    public readonly projectTeam!: pulumi.Output<outputs.storage.v1.BucketAccessControlProjectTeamResponse>;
    /**
     * The access permission for the entity.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The link to this access-control entry.
     */
    public readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a BucketAccessControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketAccessControlArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["entity"] = args ? args.entity : undefined;
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["projectTeam"] = args ? args.projectTeam : undefined;
            resourceInputs["provisionalUserProject"] = args ? args.provisionalUserProject : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["selfLink"] = args ? args.selfLink : undefined;
            resourceInputs["userProject"] = args ? args.userProject : undefined;
        } else {
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["entity"] = undefined /*out*/;
            resourceInputs["entityId"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["projectTeam"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BucketAccessControl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BucketAccessControl resource.
 */
export interface BucketAccessControlArgs {
    /**
     * The name of the bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * The domain associated with the entity, if any.
     */
    domain?: pulumi.Input<string>;
    /**
     * The email address associated with the entity, if any.
     */
    email?: pulumi.Input<string>;
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
    entity?: pulumi.Input<string>;
    /**
     * The ID for the entity, if any.
     */
    entityId?: pulumi.Input<string>;
    /**
     * HTTP 1.1 Entity tag for the access-control entry.
     */
    etag?: pulumi.Input<string>;
    /**
     * The ID of the access-control entry.
     */
    id?: pulumi.Input<string>;
    /**
     * The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
     */
    kind?: pulumi.Input<string>;
    /**
     * The project team associated with the entity, if any.
     */
    projectTeam?: pulumi.Input<inputs.storage.v1.BucketAccessControlProjectTeamArgs>;
    provisionalUserProject?: pulumi.Input<string>;
    /**
     * The access permission for the entity.
     */
    role?: pulumi.Input<string>;
    /**
     * The link to this access-control entry.
     */
    selfLink?: pulumi.Input<string>;
    userProject?: pulumi.Input<string>;
}
