// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified InPlaceSnapshot resource in the specified zone.
 */
export function getZoneInPlaceSnapshot(args: GetZoneInPlaceSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneInPlaceSnapshotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/alpha:getZoneInPlaceSnapshot", {
        "inPlaceSnapshot": args.inPlaceSnapshot,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

export interface GetZoneInPlaceSnapshotArgs {
    inPlaceSnapshot: string;
    project: string;
    zone: string;
}

export interface GetZoneInPlaceSnapshotResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Size of the source disk, specified in GB.
     */
    readonly diskSizeGb: string;
    /**
     * Specifies to create an application consistent in-place snapshot by informing the OS to prepare for the snapshot process. Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
     */
    readonly guestFlush: boolean;
    /**
     * Type of the resource. Always compute#inPlaceSnapshot for InPlaceSnapshot resources.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this InPlaceSnapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve a InPlaceSnapshot.
     */
    readonly labelFingerprint: string;
    /**
     * Labels to apply to this InPlaceSnapshot. These can be later modified by the setLabels method. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * URL of the region where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource's resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * URL of the source disk used to create this in-place snapshot. Note that the source disk must be in the same zone/region as the in-place snapshot to be created. This can be a full or valid partial URL. For example, the following are valid values:  
     * - https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk 
     * - projects/project/zones/zone/disks/disk 
     * - zones/zone/disks/disk
     */
    readonly sourceDisk: string;
    /**
     * The ID value of the disk used to create this InPlaceSnapshot. This value may be used to determine whether the InPlaceSnapshot was taken from the current or a previous instance of a given disk name.
     */
    readonly sourceDiskId: string;
    /**
     * The status of the inPlaceSnapshot. This can be CREATING, DELETING, FAILED, or READY.
     */
    readonly status: string;
    /**
     * URL of the zone where the in-place snapshot resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly zone: string;
}
