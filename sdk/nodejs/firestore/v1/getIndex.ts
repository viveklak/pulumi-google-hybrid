// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a composite index.
 */
export function getIndex(args: GetIndexArgs, opts?: pulumi.InvokeOptions): Promise<GetIndexResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:firestore/v1:getIndex", {
        "collectionGroupId": args.collectionGroupId,
        "databaseId": args.databaseId,
        "indexId": args.indexId,
        "project": args.project,
    }, opts);
}

export interface GetIndexArgs {
    collectionGroupId: string;
    databaseId: string;
    indexId: string;
    project?: string;
}

export interface GetIndexResult {
    /**
     * The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
     */
    readonly fields: outputs.firestore.v1.GoogleFirestoreAdminV1IndexFieldResponse[];
    /**
     * A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
     */
    readonly name: string;
    /**
     * Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
     */
    readonly queryScope: string;
    /**
     * The serving state of the index.
     */
    readonly state: string;
}

export function getIndexOutput(args: GetIndexOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIndexResult> {
    return pulumi.output(args).apply(a => getIndex(a, opts))
}

export interface GetIndexOutputArgs {
    collectionGroupId: pulumi.Input<string>;
    databaseId: pulumi.Input<string>;
    indexId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
