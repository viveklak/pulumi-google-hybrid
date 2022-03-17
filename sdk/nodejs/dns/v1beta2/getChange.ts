// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Fetches the representation of an existing Change.
 */
export function getChange(args: GetChangeArgs, opts?: pulumi.InvokeOptions): Promise<GetChangeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dns/v1beta2:getChange", {
        "changeId": args.changeId,
        "clientOperationId": args.clientOperationId,
        "managedZone": args.managedZone,
        "project": args.project,
    }, opts);
}

export interface GetChangeArgs {
    changeId: string;
    clientOperationId?: string;
    managedZone: string;
    project?: string;
}

export interface GetChangeResult {
    /**
     * Which ResourceRecordSets to add?
     */
    readonly additions: outputs.dns.v1beta2.ResourceRecordSetResponse[];
    /**
     * Which ResourceRecordSets to remove? Must match existing data exactly.
     */
    readonly deletions: outputs.dns.v1beta2.ResourceRecordSetResponse[];
    /**
     * If the DNS queries for the zone will be served.
     */
    readonly isServing: boolean;
    readonly kind: string;
    /**
     * The time that this operation was started by the server (output only). This is in RFC3339 text format.
     */
    readonly startTime: string;
    /**
     * Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
     */
    readonly status: string;
}

export function getChangeOutput(args: GetChangeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChangeResult> {
    return pulumi.output(args).apply(a => getChange(a, opts))
}

export interface GetChangeOutputArgs {
    changeId: pulumi.Input<string>;
    clientOperationId?: pulumi.Input<string>;
    managedZone: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
