// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a single Utilization Report.
 */
export function getUtilizationReport(args: GetUtilizationReportArgs, opts?: pulumi.InvokeOptions): Promise<GetUtilizationReportResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:vmmigration/v1:getUtilizationReport", {
        "location": args.location,
        "project": args.project,
        "sourceId": args.sourceId,
        "utilizationReportId": args.utilizationReportId,
        "view": args.view,
    }, opts);
}

export interface GetUtilizationReportArgs {
    location: string;
    project?: string;
    sourceId: string;
    utilizationReportId: string;
    view?: string;
}

export interface GetUtilizationReportResult {
    /**
     * The time the report was created (this refers to the time of the request, not the time the report creation completed).
     */
    readonly createTime: string;
    /**
     * The report display name, as assigned by the user.
     */
    readonly displayName: string;
    /**
     * Provides details on the state of the report in case of an error.
     */
    readonly error: outputs.vmmigration.v1.StatusResponse;
    /**
     * The point in time when the time frame ends. Notice that the time frame is counted backwards. For instance if the "frame_end_time" value is 2021/01/20 and the time frame is WEEK then the report covers the week between 2021/01/20 and 2021/01/14.
     */
    readonly frameEndTime: string;
    /**
     * The report unique name.
     */
    readonly name: string;
    /**
     * Current state of the report.
     */
    readonly state: string;
    /**
     * The time the state was last set.
     */
    readonly stateTime: string;
    /**
     * Time frame of the report.
     */
    readonly timeFrame: string;
    /**
     * Total number of VMs included in the report.
     */
    readonly vmCount: number;
    /**
     * List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
     */
    readonly vms: outputs.vmmigration.v1.VmUtilizationInfoResponse[];
}

export function getUtilizationReportOutput(args: GetUtilizationReportOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUtilizationReportResult> {
    return pulumi.output(args).apply(a => getUtilizationReport(a, opts))
}

export interface GetUtilizationReportOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
    utilizationReportId: pulumi.Input<string>;
    view?: pulumi.Input<string>;
}
