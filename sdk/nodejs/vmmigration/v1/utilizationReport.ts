// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new UtilizationReport.
 * Auto-naming is currently not supported for this resource.
 */
export class UtilizationReport extends pulumi.CustomResource {
    /**
     * Get an existing UtilizationReport resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): UtilizationReport {
        return new UtilizationReport(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:vmmigration/v1:UtilizationReport';

    /**
     * Returns true if the given object is an instance of UtilizationReport.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UtilizationReport {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UtilizationReport.__pulumiType;
    }

    /**
     * The time the report was created (this refers to the time of the request, not the time the report creation completed).
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The report display name, as assigned by the user.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Provides details on the state of the report in case of an error.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.vmmigration.v1.StatusResponse>;
    /**
     * The point in time when the time frame ends. Notice that the time frame is counted backwards. For instance if the "frame_end_time" value is 2021/01/20 and the time frame is WEEK then the report covers the week between 2021/01/20 and 2021/01/14.
     */
    public /*out*/ readonly frameEndTime!: pulumi.Output<string>;
    /**
     * The report unique name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Current state of the report.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The time the state was last set.
     */
    public /*out*/ readonly stateTime!: pulumi.Output<string>;
    /**
     * Time frame of the report.
     */
    public readonly timeFrame!: pulumi.Output<string>;
    /**
     * Total number of VMs included in the report.
     */
    public /*out*/ readonly vmCount!: pulumi.Output<number>;
    /**
     * List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
     */
    public readonly vms!: pulumi.Output<outputs.vmmigration.v1.VmUtilizationInfoResponse[]>;

    /**
     * Create a UtilizationReport resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UtilizationReportArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.sourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceId'");
            }
            if ((!args || args.utilizationReportId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'utilizationReportId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["sourceId"] = args ? args.sourceId : undefined;
            resourceInputs["timeFrame"] = args ? args.timeFrame : undefined;
            resourceInputs["utilizationReportId"] = args ? args.utilizationReportId : undefined;
            resourceInputs["vms"] = args ? args.vms : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["frameEndTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTime"] = undefined /*out*/;
            resourceInputs["vmCount"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["frameEndTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTime"] = undefined /*out*/;
            resourceInputs["timeFrame"] = undefined /*out*/;
            resourceInputs["vmCount"] = undefined /*out*/;
            resourceInputs["vms"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UtilizationReport.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a UtilizationReport resource.
 */
export interface UtilizationReportArgs {
    /**
     * The report display name, as assigned by the user.
     */
    displayName?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
    /**
     * Time frame of the report.
     */
    timeFrame?: pulumi.Input<enums.vmmigration.v1.UtilizationReportTimeFrame>;
    utilizationReportId: pulumi.Input<string>;
    /**
     * List of utilization information per VM. When sent as part of the request, the "vm_id" field is used in order to specify which VMs to include in the report. In that case all other fields are ignored.
     */
    vms?: pulumi.Input<pulumi.Input<inputs.vmmigration.v1.VmUtilizationInfoArgs>[]>;
}
