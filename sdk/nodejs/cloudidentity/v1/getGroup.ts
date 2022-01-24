// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves a `Group`.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudidentity/v1:getGroup", {
        "groupId": args.groupId,
    }, opts);
}

export interface GetGroupArgs {
    groupId: string;
}

export interface GetGroupResult {
    /**
     * The time when the `Group` was created.
     */
    readonly createTime: string;
    /**
     * An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
     */
    readonly description: string;
    /**
     * The display name of the `Group`.
     */
    readonly displayName: string;
    /**
     * Optional. Dynamic group metadata like queries and status.
     */
    readonly dynamicGroupMetadata: outputs.cloudidentity.v1.DynamicGroupMetadataResponse;
    /**
     * Immutable. The `EntityKey` of the `Group`.
     */
    readonly groupKey: outputs.cloudidentity.v1.EntityKeyResponse;
    /**
     * One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value.
     */
    readonly labels: {[key: string]: string};
    /**
     * The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group}`.
     */
    readonly name: string;
    /**
     * Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source}` for external- identity-mapped groups or `customers/{customer}` for Google Groups. The `customer` must begin with "C" (for example, 'C046psxkn').
     */
    readonly parent: string;
    /**
     * The time when the `Group` was last updated.
     */
    readonly updateTime: string;
}

export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply(a => getGroup(a, opts))
}

export interface GetGroupOutputArgs {
    groupId: pulumi.Input<string>;
}
