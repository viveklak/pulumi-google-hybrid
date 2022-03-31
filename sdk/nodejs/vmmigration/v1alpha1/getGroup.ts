// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Group.
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:vmmigration/v1alpha1:getGroup", {
        "groupId": args.groupId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetGroupArgs {
    groupId: string;
    location: string;
    project?: string;
}

export interface GetGroupResult {
    /**
     * The create time timestamp.
     */
    readonly createTime: string;
    /**
     * User-provided description of the group.
     */
    readonly description: string;
    /**
     * Display name is a user defined name for this group which can be updated.
     */
    readonly displayName: string;
    /**
     * The Group name.
     */
    readonly name: string;
    /**
     * The update time timestamp.
     */
    readonly updateTime: string;
}

export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupResult> {
    return pulumi.output(args).apply(a => getGroup(a, opts))
}

export interface GetGroupOutputArgs {
    groupId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
