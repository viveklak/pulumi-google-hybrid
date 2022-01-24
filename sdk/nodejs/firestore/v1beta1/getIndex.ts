// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets an index.
 */
export function getIndex(args: GetIndexArgs, opts?: pulumi.InvokeOptions): Promise<GetIndexResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:firestore/v1beta1:getIndex", {
        "databaseId": args.databaseId,
        "indexId": args.indexId,
        "project": args.project,
    }, opts);
}

export interface GetIndexArgs {
    databaseId: string;
    indexId: string;
    project?: string;
}

export interface GetIndexResult {
    /**
     * The collection ID to which this index applies. Required.
     */
    readonly collectionId: string;
    /**
     * The fields to index.
     */
    readonly fields: outputs.firestore.v1beta1.GoogleFirestoreAdminV1beta1IndexFieldResponse[];
    /**
     * The resource name of the index. Output only.
     */
    readonly name: string;
    /**
     * The state of the index. Output only.
     */
    readonly state: string;
}

export function getIndexOutput(args: GetIndexOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIndexResult> {
    return pulumi.output(args).apply(a => getIndex(a, opts))
}

export interface GetIndexOutputArgs {
    databaseId: pulumi.Input<string>;
    indexId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
