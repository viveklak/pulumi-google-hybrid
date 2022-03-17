// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the ACL entry for the specified entity on the specified bucket.
 */
export function getBucketAccessControl(args: GetBucketAccessControlArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketAccessControlResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:storage/v1:getBucketAccessControl", {
        "bucket": args.bucket,
        "entity": args.entity,
        "provisionalUserProject": args.provisionalUserProject,
        "userProject": args.userProject,
    }, opts);
}

export interface GetBucketAccessControlArgs {
    bucket: string;
    entity: string;
    provisionalUserProject?: string;
    userProject?: string;
}

export interface GetBucketAccessControlResult {
    /**
     * The name of the bucket.
     */
    readonly bucket: string;
    /**
     * The domain associated with the entity, if any.
     */
    readonly domain: string;
    /**
     * The email address associated with the entity, if any.
     */
    readonly email: string;
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
    readonly entity: string;
    /**
     * The ID for the entity, if any.
     */
    readonly entityId: string;
    /**
     * HTTP 1.1 Entity tag for the access-control entry.
     */
    readonly etag: string;
    /**
     * The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
     */
    readonly kind: string;
    /**
     * The project team associated with the entity, if any.
     */
    readonly projectTeam: outputs.storage.v1.BucketAccessControlProjectTeamResponse;
    /**
     * The access permission for the entity.
     */
    readonly role: string;
    /**
     * The link to this access-control entry.
     */
    readonly selfLink: string;
}

export function getBucketAccessControlOutput(args: GetBucketAccessControlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketAccessControlResult> {
    return pulumi.output(args).apply(a => getBucketAccessControl(a, opts))
}

export interface GetBucketAccessControlOutputArgs {
    bucket: pulumi.Input<string>;
    entity: pulumi.Input<string>;
    provisionalUserProject?: pulumi.Input<string>;
    userProject?: pulumi.Input<string>;
}
