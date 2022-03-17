// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a big query export.
 */
export function getFolderBigQueryExport(args: GetFolderBigQueryExportArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderBigQueryExportResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:securitycenter/v1:getFolderBigQueryExport", {
        "bigQueryExportId": args.bigQueryExportId,
        "folderId": args.folderId,
    }, opts);
}

export interface GetFolderBigQueryExportArgs {
    bigQueryExportId: string;
    folderId: string;
}

export interface GetFolderBigQueryExportResult {
    /**
     * The time at which the big query export was created. This field is set by the server and will be ignored if provided on export on creation.
     */
    readonly createTime: string;
    /**
     * The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
     */
    readonly dataset: string;
    /**
     * The description of the export (max of 1024 characters).
     */
    readonly description: string;
    /**
     * Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes.
     */
    readonly filter: string;
    /**
     * Email address of the user who last edited the big query export. This field is set by the server and will be ignored if provided on export creation or update.
     */
    readonly mostRecentEditor: string;
    /**
     * The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
     */
    readonly name: string;
    /**
     * The service account that needs permission to create table, upload data to the big query dataset.
     */
    readonly principal: string;
    /**
     * The most recent time at which the big export was updated. This field is set by the server and will be ignored if provided on export creation or update.
     */
    readonly updateTime: string;
}

export function getFolderBigQueryExportOutput(args: GetFolderBigQueryExportOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFolderBigQueryExportResult> {
    return pulumi.output(args).apply(a => getFolderBigQueryExport(a, opts))
}

export interface GetFolderBigQueryExportOutputArgs {
    bigQueryExportId: pulumi.Input<string>;
    folderId: pulumi.Input<string>;
}
