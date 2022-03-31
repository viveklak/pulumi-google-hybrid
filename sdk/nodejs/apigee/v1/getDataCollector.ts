// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a data collector.
 */
export function getDataCollector(args: GetDataCollectorArgs, opts?: pulumi.InvokeOptions): Promise<GetDataCollectorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:apigee/v1:getDataCollector", {
        "datacollectorId": args.datacollectorId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetDataCollectorArgs {
    datacollectorId: string;
    organizationId: string;
}

export interface GetDataCollectorResult {
    /**
     * The time at which the data collector was created in milliseconds since the epoch.
     */
    readonly createdAt: string;
    /**
     * A description of the data collector.
     */
    readonly description: string;
    /**
     * The time at which the Data Collector was last updated in milliseconds since the epoch.
     */
    readonly lastModifiedAt: string;
    /**
     * ID of the data collector. Must begin with `dc_`.
     */
    readonly name: string;
    /**
     * Immutable. The type of data this data collector will collect.
     */
    readonly type: string;
}

export function getDataCollectorOutput(args: GetDataCollectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataCollectorResult> {
    return pulumi.output(args).apply(a => getDataCollector(a, opts))
}

export interface GetDataCollectorOutputArgs {
    datacollectorId: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
