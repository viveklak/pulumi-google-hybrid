// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Source.
 */
export function getSource(args: GetSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:vmmigration/v1alpha1:getSource", {
        "location": args.location,
        "project": args.project,
        "sourceId": args.sourceId,
    }, opts);
}

export interface GetSourceArgs {
    location: string;
    project?: string;
    sourceId: string;
}

export interface GetSourceResult {
    /**
     * The create time timestamp.
     */
    readonly createTime: string;
    /**
     * User-provided description of the source.
     */
    readonly description: string;
    /**
     * Provides details on the state of the Source in case of an error.
     */
    readonly error: outputs.vmmigration.v1alpha1.StatusResponse;
    /**
     * The labels of the source.
     */
    readonly labels: {[key: string]: string};
    /**
     * The Source name.
     */
    readonly name: string;
    /**
     * The update time timestamp.
     */
    readonly updateTime: string;
    /**
     * Vmware type source details.
     */
    readonly vmware: outputs.vmmigration.v1alpha1.VmwareSourceDetailsResponse;
}

export function getSourceOutput(args: GetSourceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceResult> {
    return pulumi.output(args).apply(a => getSource(a, opts))
}

export interface GetSourceOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
}
