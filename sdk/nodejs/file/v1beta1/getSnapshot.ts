// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets the details of a specific snapshot.
 */
export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:file/v1beta1:getSnapshot", {
        "instanceId": args.instanceId,
        "location": args.location,
        "project": args.project,
        "snapshotId": args.snapshotId,
    }, opts);
}

export interface GetSnapshotArgs {
    instanceId: string;
    location: string;
    project?: string;
    snapshotId: string;
}

export interface GetSnapshotResult {
    /**
     * The time when the snapshot was created.
     */
    readonly createTime: string;
    /**
     * A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
     */
    readonly description: string;
    /**
     * The amount of bytes needed to allocate a full copy of the snapshot content
     */
    readonly filesystemUsedBytes: string;
    /**
     * Resource labels to represent user provided metadata.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the snapshot, in the format `projects/{project_id}/locations/{location_id}/instances/{instance_id}/snapshots/{snapshot_id}`.
     */
    readonly name: string;
    /**
     * The snapshot state.
     */
    readonly state: string;
}

export function getSnapshotOutput(args: GetSnapshotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotResult> {
    return pulumi.output(args).apply(a => getSnapshot(a, opts))
}

export interface GetSnapshotOutputArgs {
    instanceId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    snapshotId: pulumi.Input<string>;
}
