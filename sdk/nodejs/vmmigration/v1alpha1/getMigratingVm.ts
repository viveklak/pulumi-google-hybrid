// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single MigratingVm.
 */
export function getMigratingVm(args: GetMigratingVmArgs, opts?: pulumi.InvokeOptions): Promise<GetMigratingVmResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:vmmigration/v1alpha1:getMigratingVm", {
        "location": args.location,
        "migratingVmId": args.migratingVmId,
        "project": args.project,
        "sourceId": args.sourceId,
        "view": args.view,
    }, opts);
}

export interface GetMigratingVmArgs {
    location: string;
    migratingVmId: string;
    project?: string;
    sourceId: string;
    view?: string;
}

export interface GetMigratingVmResult {
    /**
     * Details of the target VM in Compute Engine.
     */
    readonly computeEngineTargetDefaults: outputs.vmmigration.v1alpha1.ComputeEngineTargetDefaultsResponse;
    /**
     * Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_defaults instead.
     *
     * @deprecated Details of the VM in Compute Engine. Deprecated: Use compute_engine_target_defaults instead.
     */
    readonly computeEngineVmDefaults: outputs.vmmigration.v1alpha1.TargetVMDetailsResponse;
    /**
     * The time the migrating VM was created (this refers to this resource and not to the time it was installed in the source).
     */
    readonly createTime: string;
    /**
     * The percentage progress of the current running replication cycle.
     */
    readonly currentSyncInfo: outputs.vmmigration.v1alpha1.ReplicationCycleResponse;
    /**
     * The description attached to the migrating VM by the user.
     */
    readonly description: string;
    /**
     * The display name attached to the MigratingVm by the user.
     */
    readonly displayName: string;
    /**
     * Provides details on the state of the Migrating VM in case of an error in replication.
     */
    readonly error: outputs.vmmigration.v1alpha1.StatusResponse;
    /**
     * The group this migrating vm is included in, if any. The group is represented by the full path of the appropriate Group resource.
     */
    readonly group: string;
    /**
     * The labels of the migrating VM.
     */
    readonly labels: {[key: string]: string};
    /**
     * The most updated snapshot created time in the source that finished replication.
     */
    readonly lastSync: outputs.vmmigration.v1alpha1.ReplicationSyncResponse;
    /**
     * The identifier of the MigratingVm.
     */
    readonly name: string;
    /**
     * The replication schedule policy.
     */
    readonly policy: outputs.vmmigration.v1alpha1.SchedulePolicyResponse;
    /**
     * The recent clone jobs performed on the migrating VM. This field holds the vm's last completed clone job and the vm's running clone job, if one exists. Note: To have this field populated you need to explicitly request it via the "view" parameter of the Get/List request.
     */
    readonly recentCloneJobs: outputs.vmmigration.v1alpha1.CloneJobResponse[];
    /**
     * The recent cutover jobs performed on the migrating VM. This field holds the vm's last completed cutover job and the vm's running cutover job, if one exists. Note: To have this field populated you need to explicitly request it via the "view" parameter of the Get/List request.
     */
    readonly recentCutoverJobs: outputs.vmmigration.v1alpha1.CutoverJobResponse[];
    /**
     * The unique ID of the VM in the source. The VM's name in vSphere can be changed, so this is not the VM's name but rather its moRef id. This id is of the form vm-.
     */
    readonly sourceVmId: string;
    /**
     * State of the MigratingVm.
     */
    readonly state: string;
    /**
     * The last time the migrating VM state was updated.
     */
    readonly stateTime: string;
    /**
     * The default configuration of the target VM that will be created in GCP as a result of the migration. Deprecated: Use compute_engine_target_defaults instead.
     *
     * @deprecated The default configuration of the target VM that will be created in GCP as a result of the migration. Deprecated: Use compute_engine_target_defaults instead.
     */
    readonly targetDefaults: outputs.vmmigration.v1alpha1.TargetVMDetailsResponse;
    /**
     * The last time the migrating VM resource was updated.
     */
    readonly updateTime: string;
}

export function getMigratingVmOutput(args: GetMigratingVmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMigratingVmResult> {
    return pulumi.output(args).apply(a => getMigratingVm(a, opts))
}

export interface GetMigratingVmOutputArgs {
    location: pulumi.Input<string>;
    migratingVmId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
    view?: pulumi.Input<string>;
}
