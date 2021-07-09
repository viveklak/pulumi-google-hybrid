// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets dataset by resource name.
 */
export function getDataset(args: GetDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetDatasetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:datalabeling/v1beta1:getDataset", {
        "datasetId": args.datasetId,
        "project": args.project,
    }, opts);
}

export interface GetDatasetArgs {
    datasetId: string;
    project: string;
}

export interface GetDatasetResult {
    /**
     * The names of any related resources that are blocking changes to the dataset.
     */
    readonly blockingResources: string[];
    /**
     * Time the dataset is created.
     */
    readonly createTime: string;
    /**
     * The number of data items in the dataset.
     */
    readonly dataItemCount: string;
    /**
     * Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
     */
    readonly description: string;
    /**
     * The display name of the dataset. Maximum of 64 characters.
     */
    readonly displayName: string;
    /**
     * This is populated with the original input configs where ImportData is called. It is available only after the clients import data to this dataset.
     */
    readonly inputConfigs: outputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1InputConfigResponse[];
    /**
     * Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
     */
    readonly lastMigrateTime: string;
    /**
     * Dataset resource name, format is: projects/{project_id}/datasets/{dataset_id}
     */
    readonly name: string;
}
